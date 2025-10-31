package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
)

type ResponseSoterVerifySignature struct {
	response.ResponseMiniProgram

	IsOK bool `json:"is_ok"`
}
