package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
)

type ResponseMeetingCreate struct {
	response.ResponseWork

	MeetingID int `json:"meetingid"`
}
