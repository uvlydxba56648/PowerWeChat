package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/power"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
)

type ResponseSubscribeMessageGetPubTemplateTitleList struct {
	response.ResponseMiniProgram
	Count int              `json:"count"`
	Data  []*power.HashMap `json:"data"`
}
