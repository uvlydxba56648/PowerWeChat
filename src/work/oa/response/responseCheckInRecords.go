package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/power"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
)

type ResponseCheckInRecords struct {
	response.ResponseWork

	CheckInData []*power.HashMap `json:"checkindata"`
}
