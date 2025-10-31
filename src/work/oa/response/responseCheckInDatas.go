package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/power"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
)

type ResponseCheckInDatas struct {
	response.ResponseWork

	Datas []*power.HashMap `json:"datas"`
}
