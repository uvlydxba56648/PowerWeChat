package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseNFCUrlLink struct {
	response.ResponseMiniProgram
	OpenLink string `json:"openlink"` // NFC 专用 Scheme (形如 weixin://...)
}
