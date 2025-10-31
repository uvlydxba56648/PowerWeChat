package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/power"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
)

type ResponseGroupChatTransfer struct {
	response.ResponseWork

	Customer []*power.HashMap `json:"failed_chat_list"`
}
