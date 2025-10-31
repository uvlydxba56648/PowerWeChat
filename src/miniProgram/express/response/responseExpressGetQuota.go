package response

import "github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"

type ResponseExpressGetQuota struct {
	response.ResponseMiniProgram

	QuotaNum string `json:"quota_num"`
}
