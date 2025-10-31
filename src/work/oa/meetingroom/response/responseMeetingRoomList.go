package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/power"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
)

type ResponseMeetingRoomList struct {
	response.ResponseWork

	MeetingRoomList []*power.HashMap `json:"meetingroom_list"`
}
