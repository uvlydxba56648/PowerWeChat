package response

import "github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"

// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_8.shtml

type ResponseProfitSharingAddReceiver struct {
	response.ResponsePayment

	SubMchID       string `json:"sub_mchid,omitempty"`
	Type           string `json:"type"`
	Account        string `json:"account"`
	Name           string `json:"name,omitempty"`
	RelationType   string `json:"relation_type,omitempty"`
	CustomRelation string `json:"custom_relation,omitempty"`
}
