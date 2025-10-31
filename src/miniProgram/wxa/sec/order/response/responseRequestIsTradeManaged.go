package response

import "github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"

type ResponseIsTradeManaged struct {
	response.ResponseMiniProgram
	IsTradeManaged bool `json:"is_trade_managed"`
}
