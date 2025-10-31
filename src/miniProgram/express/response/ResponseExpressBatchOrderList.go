package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/power"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
)

type ResponseExpressBatchOrderList struct {
	response.ResponseMiniProgram

	OrderList []*power.HashMap `json:"order_list"`
}
