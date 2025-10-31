package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
)

type ResponseAccountAdd struct {
	response.ResponseWork

	OpenKFID string `json:"open_kfid"`
}
