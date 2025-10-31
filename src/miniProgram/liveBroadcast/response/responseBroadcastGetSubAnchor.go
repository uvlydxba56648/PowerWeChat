package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
)

type ResponseBroadcastGetSubAnchor struct {
	response.ResponseMiniProgram

	UserName string `json:"username"`
}
