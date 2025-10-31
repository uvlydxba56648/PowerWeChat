package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/payment/profitSharing/request"
)

type ResponseProfitSharingOrder struct {
	response.ResponsePayment

	SubMchID      string              `json:"sub_mchid,omitempty"` // 服务商模式下返回
	TransactionID string              `json:"transaction_id"`
	OutOrderNO    string              `json:"out_order_no"`
	OrderID       string              `json:"order_id"`
	State         string              `json:"state"`
	Receivers     []*request.Receiver `json:"receivers"`
}
