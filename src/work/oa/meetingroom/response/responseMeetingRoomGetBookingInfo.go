package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/power"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
)

type ResponseMeetingRoomGetBookingInfo struct {
	response.ResponseWork

	BookingList []*power.HashMap `json:"booking_list"`
}
