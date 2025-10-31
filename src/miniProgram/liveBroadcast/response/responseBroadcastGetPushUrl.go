package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
)

type ResponseBroadcastGetPushUrl struct {
	response.ResponseMiniProgram

	PushAddr string `json:"pushAddr"`
}
