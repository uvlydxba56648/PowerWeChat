package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
)

type ResponseBroadcastGoodsAudit struct {
	response.ResponseMiniProgram

	AuditID int `json:"auditId"`
}
