package response

import "github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"

type ResponseGetMomentTaskResult struct {
	response.ResponseWork

	Status int    `json:"status"`
	Type   string `json:"type"`
	Result Result `json:"result"`
}
