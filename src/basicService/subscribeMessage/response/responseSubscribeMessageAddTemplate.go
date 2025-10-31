package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
)

type ResponseSubscribeMessageAddTemplate struct {
	response.ResponseMiniProgram
	PriTmplID string `json:"priTmplId"`
}
