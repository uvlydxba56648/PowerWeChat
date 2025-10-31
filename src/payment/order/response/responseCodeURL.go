package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
)

type ResponseCodeURL struct {
	response.ResponsePayment

	CodeURL string `json:"code_url"`
}
