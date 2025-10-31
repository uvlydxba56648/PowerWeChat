package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
)

type ResponseCorpGroupTransferSession struct {
	response.ResponseWork
	Userid     string `json:"userid"`
	SessionKey string `json:"session_key"`
}
