package handlers

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/contract"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/openPlatform/auth"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/openPlatform/response"
	"net/http"
)

type VerifyTicketRefreshed struct {
	contract.EventHandlerInterface

	App kernel.ApplicationInterface
}

func NewVerifyTicketRefreshed(app kernel.ApplicationInterface) *VerifyTicketRefreshed {
	handler := &VerifyTicketRefreshed{
		App: app,
	}
	return handler
}

func (handler *VerifyTicketRefreshed) Handle(request *http.Request, header contract.EventInterface, content interface{}) interface{} {

	verifyTicket := content.(*response.ResponseVerifyTicket)
	ticket := verifyTicket.ComponentVerifyTicket

	if ticket != "" {
		handler.App.GetComponent("VerifyTicket").(*auth.VerifyTicket).SetTicket(ticket)
	}

	return nil
}
