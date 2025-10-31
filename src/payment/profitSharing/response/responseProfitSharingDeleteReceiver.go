package response

import "github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"

// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_9.shtml

type ResponseProfitSharingDeleteReceiver struct {
	response.ResponsePayment

	SubMchID string `json:"sub_mchid,omitempty"`
	Type     string `json:"type"`
	Account  string `json:"accountService"`
}
