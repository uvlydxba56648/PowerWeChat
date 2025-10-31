package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/power"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
)

type ResponseLinkCorpGetUser struct {
	response.ResponseWork

	UserInfo *power.HashMap `json:"user_info"`
}
