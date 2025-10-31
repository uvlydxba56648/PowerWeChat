package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
)

type ResponseSecurityMediaCheckASync struct {
	response.ResponseMiniProgram
	TraceID string `json:"trace_id"`
}
