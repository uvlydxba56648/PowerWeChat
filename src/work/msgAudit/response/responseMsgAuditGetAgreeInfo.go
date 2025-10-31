package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/power"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
)

type ResponseMsgAuditGetAgreeInfo struct {
	response.ResponseWork
	AgreeInfo []*power.HashMap `json:"agreeinfo"`
}
