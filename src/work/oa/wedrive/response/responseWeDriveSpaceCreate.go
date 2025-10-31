package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
)

type ResponseWeDriveSpaceCreate struct {
	response.ResponseWork

	SpaceID string `json:"spaceid"`
}
