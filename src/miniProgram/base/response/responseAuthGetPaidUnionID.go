package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
)

type ResponseAuthGetPaidUnionID struct {
	UnionID string `json:"unionid"`

	response.ResponseMiniProgram
}
