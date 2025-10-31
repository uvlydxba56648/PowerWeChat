package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
)

type ResponseWeDriveFileUpload struct {
	response.ResponseWork

	FileID string `json:"fileid"`
}
