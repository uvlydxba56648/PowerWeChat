package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
)

type ResponseTagCreate struct {
	response.ResponseWork

	TagID int64 `json:"tagid"`
}
