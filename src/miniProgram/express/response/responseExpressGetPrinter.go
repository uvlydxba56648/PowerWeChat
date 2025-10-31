package response

import "github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"

type ResponseExpressGetPrinter struct {
	response.ResponseMiniProgram

	Count     string   `json:"count"`
	OpenID    []string `json:"openid"`
	TagIDList []string `json:"tagid_list"`
}
