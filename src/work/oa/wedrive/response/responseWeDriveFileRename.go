package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/power"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
)

type ResponseWeDriveFileRename struct {
	response.ResponseWork

	File *power.HashMap `json:"file"`
}
