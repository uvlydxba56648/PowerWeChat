package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
)

type ResponseMaterialAddMaterial struct {
	response.ResponseOfficialAccount

	MediaID string `json:"media_id"`
	URL     string `json:"url"`
}
