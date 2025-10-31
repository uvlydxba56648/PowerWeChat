package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
)

type ResponseAppChatCreate struct {
	response.ResponseWork

	ChatID string `json:"chatid"`
}
