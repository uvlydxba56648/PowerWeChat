package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/power"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
)

type ResponsePluginGetDevApplyList struct {
	response.ResponseMiniProgram
	ApplyList []*power.HashMap `json:"apply_list"`
}
