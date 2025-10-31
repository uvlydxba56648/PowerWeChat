package models

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/contract"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/models"
)

const (
	CALLBACK_EVENT_MSGAUDIT_NOTIFY   = "msgaudit_notify"
	CALLBACK_EVENT_UNLICENSED_NOTIFY = "unlicensed_notify"
)

type EventMsgAuditNotify struct {
	contract.EventInterface
	models.CallbackMessageHeader
	AgentID string `xml:"AgentID"`
}
