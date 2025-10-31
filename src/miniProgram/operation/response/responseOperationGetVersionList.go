package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/power"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
)

type ResponseOperationGetVersionList struct {
	response.ResponseMiniProgram
	CVList []*power.HashMap `json:"cvlist"`
}
