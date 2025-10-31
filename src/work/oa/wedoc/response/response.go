package response

import "github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"

type ResponseWeDocCreateForm struct {
	response.ResponseWork
	FormId string `json:"formid"`
}
