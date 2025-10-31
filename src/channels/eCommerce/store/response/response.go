package response

import "github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"

type Info struct {
	Nickname    string `json:"nickname"`
	HeadImgUrl  string `json:"headimg_url"`
	SubjectType string `json:"subject_type"`
	Status      string `json:"status"`
	Username    string `json:"username"`
}

type ResponseGetBasicInfo struct {
	response.ResponseECommerce
	Info Info `json:"info"`
}

type ResponseGetShopQRCode struct {
	response.ResponseECommerce
	ShopQrcode string `json:"shop_qrcode"`
}
