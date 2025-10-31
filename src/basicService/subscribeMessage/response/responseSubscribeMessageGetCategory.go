package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/power"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
)

type ResponseSubscribeMessageGetCategory struct {
	response.ResponseMiniProgram
	Data []*power.HashMap `json:"data"`
}
