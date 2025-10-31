package costom

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/payment/notify"
	"io"
	"net/http"

	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/support"
	base2 "github.com/uvlydxba56648/PowerWeChat/v3/src/payment/base"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/payment/kernel"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/payment/notify/request"
)

type CustomHandler struct {
	App        kernel.ApplicationPaymentInterface
	Message    *request.RequestNotify
	fail       string
	Attributes *object.StringMap
	Check      bool
	Sign       bool

	ExternalBody io.Reader

	// Handle func(closure func(message *request.RequestNotify, transaction *models.Transaction, fail func(groupWelcomeTemplate string)) interface{}) *http.Response
}

func NewCustomHandler(app kernel.ApplicationPaymentInterface, body io.Reader) *CustomHandler {

	// -------------- external request --------------
	var req io.Reader = &bytes.Buffer{}
	if body != nil {
		req = body
	}

	return &CustomHandler{
		App:          app,
		Check:        true,
		Sign:         false,
		ExternalBody: req,
	}
}

func (handler *CustomHandler) Fail(message string) {
	handler.fail = message
}

func (handler *CustomHandler) RespondWith(attributes *object.StringMap, sign bool) *CustomHandler {

	handler.Attributes = attributes
	handler.Sign = sign

	return handler
}
func (handler *CustomHandler) ToResponse() (response *http.Response, err error) {

	returnCode := notify.SUCCESS
	returnMsg := "成功"
	if handler.fail != "" {
		returnCode = notify.FAIL
		returnMsg = handler.fail
		err = errors.New(handler.fail)
	}
	base := &object.StringMap{
		"code":           returnCode,
		"uniformMessage": returnMsg,
	}

	attributes := object.MergeStringMap(base, handler.Attributes)
	baseClient := (handler.App).GetComponent("Base").(*base2.Client)
	if handler.Sign {
		(*attributes)["sign"], err = baseClient.BaseClient.Signer.GenerateSign("")
		if err != nil {
			return nil, err
		}
	}

	bodyBuffer, _ := json.Marshal(attributes)
	rs := &http.Response{
		StatusCode: http.StatusOK,
	}
	rs.Body = io.NopCloser(bytes.NewBuffer(bodyBuffer))

	return rs, err
}

func (handler *CustomHandler) GetMessage() (notify *request.RequestNotify, err error) {

	if handler.Message != nil {
		return handler.Message, nil
	}

	body := handler.ExternalBody

	requestBody, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}
	handler.Message = &request.RequestNotify{}
	err = object.JsonDecode(requestBody, handler.Message)
	if err != nil {
		return nil, err
	}

	// note: 自定义的不保存 http request
	// handler.Message.RawRequest = body

	return handler.Message, nil
}

func (handler *CustomHandler) DecryptMessage() (string, error) {
	message, err := handler.GetMessage()
	if err != nil {
		return "", err
	}
	if message.Resource == nil {
		return "", errors.New("uniformMessage doesn't have the key value")
	}

	config := (handler.App).GetConfig()
	wxKey := config.GetString("mch_api_v3_key", "")
	nonce := message.Resource.Nonce
	associatedData := message.Resource.AssociatedData
	cipherText := message.Resource.Ciphertext
	return support.DecryptAES256GCM(
		wxKey,
		associatedData,
		nonce,
		cipherText,
	)

}

func (handler *CustomHandler) Strict(result interface{}) {

	bResult := true
	strResult := ""
	switch result.(type) {
	case bool:
		bResult = result.(bool)
		strResult = fmt.Sprintf("%t", result)
	case string:
		strResult = result.(string)
		if strResult != "" {
			bResult = false
		}
	default:
		return
	}

	if bResult != true && handler.fail == "" {
		handler.Fail(strResult)
	}
}

func (handler *CustomHandler) reqInfo() (content string, err error) {

	content, err = handler.DecryptMessage()
	if err != nil {
		return "", err
	}

	// save the decoded content to message resource
	handler.Message.Resource.Plaintext = content

	return content, nil
}
