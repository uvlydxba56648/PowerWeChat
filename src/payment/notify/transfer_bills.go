package notify

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/models"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/payment/kernel"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/payment/notify/request"
	"net/http"
)

type TransferBills struct {
	*Handler
}

func NewTransferBillsNotify(app kernel.ApplicationPaymentInterface, request *http.Request) *TransferBills {

	paid := &TransferBills{
		NewHandler(app, request),
	}

	return paid
}

func (comp *TransferBills) Handle(closure func(message *request.RequestNotify, refund *models.TransferBills, fail func(message string)) interface{}) (*http.Response, error) {

	message, err := comp.GetMessage()
	if err != nil {
		return nil, err
	}

	reqInfo, err := comp.reqInfo()
	if err != nil {
		return nil, err
	}

	// struct the content
	transferBills := &models.TransferBills{}
	err = object.JsonDecode([]byte(reqInfo), transferBills)
	if err != nil {
		return nil, err
	}

	result := closure(message, transferBills, comp.Fail)
	comp.Strict(result)

	return comp.ToResponse()

}
