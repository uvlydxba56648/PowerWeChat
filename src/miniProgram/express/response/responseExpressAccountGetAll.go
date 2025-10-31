package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/power"
	response2 "github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
)

type ResponseExpressAccountGetAll struct {
	*response2.ResponseMiniProgram

	Count int              `json:"count"`
	List  []*power.HashMap `json:"list"`
}
