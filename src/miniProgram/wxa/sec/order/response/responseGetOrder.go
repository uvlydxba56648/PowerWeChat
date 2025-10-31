package response

import "github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"

type ResponseGetOrder struct {
	response.ResponseMiniProgram
	Order Order `json:"order"`
}

type Order struct {
	TransactionID   string   `json:"transaction_id"`
	MerchantID      string   `json:"merchant_id"`
	SubMerchantID   string   `json:"sub_merchant_id"`
	MerchantTradeNo string   `json:"merchant_trade_no"`
	Description     string   `json:"description"`
	PaidAmount      int      `json:"paid_amount"`
	OpenID          string   `json:"openid"`
	TradeCreateTime int      `json:"trade_create_time"`
	PayTime         int      `json:"pay_time"`
	OrderState      int8     `json:"order_state"`
	InComplaint     bool     `json:"in_complaint"`
	Shipping        Shipping `json:"shipping"`
}

type Shipping struct {
	DeliveryMode        int8           `json:"delivery_mode"`
	LogisticsType       int8           `json:"logistics_type"`
	FinishShipping      bool           `json:"finish_shipping"`
	GoodsDesc           string         `json:"goods_desc"`
	FinishShippingCount int8           `json:"finish_shipping_count"`
	ShippingList        []ShippingList `json:"shipping_list"`
}

type ShippingList struct {
	TrackingNo     string  `json:"tracking_no"`
	ExpressCompany string  `json:"express_company"`
	GoodsDesc      string  `json:"goods_desc"`
	UploadTime     int     `json:"upload_time"`
	Contact        Contact `json:"contact"`
}

type Contact struct {
	ConsignorContact string `json:"consignor_contact"`
	ReceiverContact  string `json:"receiver_contact"`
}
