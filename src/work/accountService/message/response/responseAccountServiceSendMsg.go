package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
)

type ResponseAccountServiceSendMsg struct {
	response.ResponseWork

	MsgID string `json:"msgid"`
}
