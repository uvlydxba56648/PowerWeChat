package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/power"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
)

type ResponseLivingGetWatchStat struct {
	response.ResponseWork

	NextKey  string         `json:"next_key"`
	StatInfo *power.HashMap `json:"stat_info"`
}
