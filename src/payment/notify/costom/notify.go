package costom

import (
	"io"
	"net/http"

	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/payment/kernel"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/payment/notify/request"
)

// CustomNotify 框架内没有实现的，可以使用这个自己实现
type CustomNotify[T any] struct {
	*CustomHandler
}

func NewCustomNotify[T any](app kernel.ApplicationPaymentInterface, body io.Reader) *CustomNotify[T] {
	return &CustomNotify[T]{
		NewCustomHandler(app, body),
	}
}

func (comp *CustomNotify[T]) Handle(closure func(message *request.RequestNotify, transaction *T, fail func(message string)) interface{}) (*http.Response, error) {
	message, err := comp.GetMessage()
	if err != nil {
		return nil, err
	}

	reqInfo, err := comp.reqInfo()
	if err != nil {
		return nil, err
	}

	// struct the content
	var transaction = new(T)
	err = object.JsonDecode([]byte(reqInfo), transaction)
	if err != nil {
		return nil, err
	}

	result := closure(message, transaction, comp.Fail)
	comp.Strict(result)

	return comp.ToResponse()
}
