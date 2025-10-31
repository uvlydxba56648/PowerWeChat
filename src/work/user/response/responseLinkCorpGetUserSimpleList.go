package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/power"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
)

type ResponseLinkCorpGetUserSimpleList struct {
	response.ResponseWork

	UserList []*power.HashMap `json:"userlist"`
}
