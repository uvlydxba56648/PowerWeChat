package response

import "github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"

type ResponseIsTradeManagementConfirmationCompleted struct {
	response.ResponseMiniProgram
	Completed bool `json:"completed"`
}
