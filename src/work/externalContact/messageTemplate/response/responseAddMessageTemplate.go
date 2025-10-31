package response

import "github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"

type ResponseAddMessageTemplate struct {
	response.ResponseWork
	FailList []string `json:"fail_list"`
	MsgID    string   `json:"msgid"`
}
