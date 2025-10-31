package response

import "github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"

type ResponseCreateTemplate struct {
	response.ResponseWork

	TemplateId string `json:"template_id"`
}
