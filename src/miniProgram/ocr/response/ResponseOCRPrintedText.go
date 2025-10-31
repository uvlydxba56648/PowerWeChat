package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/power"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
)

type ResponseOCRPrintedText struct {
	response.ResponseMiniProgram
	Items   []*power.HashMap `json:"items"`
	ImgSize *power.HashMap   `json:"img_size"`
}
