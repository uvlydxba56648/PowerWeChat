package response

import "github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"

type ResponseGetOrderList struct {
	response.ResponseMiniProgram
	LastIndex string  `json:"last_index"`
	HasMore   bool    `json:"has_more"`
	OrderList []Order `json:"order_list"`
}
