package response

import "github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"

type GetOrderRes struct {
	AppID              string         `json:"appid"`                          // 小程序ID
	MchID              string         `json:"mchid"`                          // 微信商户号
	OutTradeNo         string         `json:"out_trade_no"`                   // 商户订单号
	OrderID            string         `json:"order_id"`                       // B2b支付订单号
	PayStatus          string         `json:"pay_status"`                     // 订单状态
	PayTime            string         `json:"pay_time,omitempty"`             // 支付完成时间
	Attach             string         `json:"attach,omitempty"`               // 附加数据
	PayerOpenid        string         `json:"payer_openid"`                   // 支付者
	Amount             GetOrderAmount `json:"amount,omitempty"`               // 订单金额
	WxPayTransactionID string         `json:"wxpay_transaction_id,omitempty"` // 微信支付订单号
	Env                int            `json:"env"`                            // 订单环境

	response.ResponsePayment
}

type GetOrderAmount struct {
	OrderAmount int64  `json:"order_amount"` // 订单总金额
	PayerAmount int64  `json:"payer_amount"` // 用户支付金额
	Currency    string `json:"currency"`     // 货币类型
}

type RefundRes struct {
	RefundID    string `json:"refund_id"`     // B2b支付退款单号
	OutRefundNo string `json:"out_refund_no"` // 商户退款单号
	OrderID     string `json:"order_id"`      // B2b支付订单号
	OutTradeNo  string `json:"out_trade_no"`  // 商户订单号

	response.ResponsePayment
}

type GetRefundRes struct {
	RefundID      string          `json:"refund_id"`       // B2b支付退款单号
	OutRefundNo   string          `json:"out_refund_no"`   // 商户退款单号
	OrderID       string          `json:"order_id"`        // B2b支付订单号
	OutTradeNo    string          `json:"out_trade_no"`    // 商户订单号
	CreateTime    string          `json:"create_time"`     // 退款创建时间
	RefundTime    string          `json:"refund_time"`     // 退款成功时间
	RefundStatus  string          `json:"refund_status"`   // 退款状态
	RefundDesc    string          `json:"refund_desc"`     // 退款说明
	Amount        GetRefundAmount `json:"amount"`          // 金额信息
	WxPayRefundID string          `json:"wxpay_refund_id"` // 微信支付退款单号

	response.ResponsePayment
}

type GetRefundAmount struct {
	OrderAmount  int64  `json:"order_amount"`  // 订单总金额
	RefundAmount int64  `json:"refund_amount"` // 退款金额
	Currency     string `json:"currency"`      // 货币类型
}

type AddProfitSharingAccountRes response.ResponsePayment

type DelProfitSharingAccountRes response.ResponsePayment

type RetailProfitSharingAccountInfo struct {
	SharingAccountType string `json:"sharing_account_type"` // 分账接收方类型
	SharingAccount     string `json:"sharing_account"`      // 分账接收方标识
	AddTime            int    `json:"add_time"`             // 添加时间
	UpdateTime         int    `json:"update_time"`          // 更新时间
	Name               string `json:"name"`                 // 商户名称
}

type QueryProfitSharingAccountRes struct {
	AccountList []RetailProfitSharingAccountInfo `json:"account_list"`

	response.ResponsePayment
}

type CreateProfitSharingOrderRes response.ResponsePayment

type QueryProfitSharingOrderRes struct {
	OrderStatus int `json:"order_status"`
	response.ResponsePayment
}

type QueryProfitSharingRemainAmtRes struct {
	RemainAmt int64 `json:"remain_amt"` // 冻结金额，单位：分
	response.ResponsePayment
}

type FinishProfitSharingOrderRes response.ResponsePayment

type RefundProfitSharingRes response.ResponsePayment

type QueryRefundProfitSharingOrderRes struct {
	OrderStatus int `json:"order_status"`
	response.ResponsePayment
}

type DownloadBillRes struct {
	SuccessBillURL       string `json:"success_bill_url"`        // 支付成功订单下载链接，否
	RefundBillURL        string `json:"refund_bill_url"`         // 退款单下载链接，否
	AllBillURL           string `json:"all_bill_url"`            // 包括支付成功和退款单的下载链接，否
	FundBillURL          string `json:"fund_bill_url"`           // 资金账单下载链接，否
	EndedDayAvailAmt     int64  `json:"ended_day_avail_amt"`     // 日终账户可提现金额，否
	EndedDayFrozenAmt    int64  `json:"ended_day_frozen_amt"`    // 日终账户待结算金额，否
	EndedDayTotalAmt     int64  `json:"ended_day_total_amt"`     // 日终账户总金额，否
	ProfitSharingBillURL string `json:"profit_sharing_bill_url"` // 分账成功订单下载链接，否
	ProfitRefundBillURL  string `json:"profit_refund_bill_url"`  // 分账回退单下载链接，否

	response.ResponsePayment
}
