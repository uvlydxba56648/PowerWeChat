package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
)

type ResponseAgentList struct {
	response.ResponseWork
	AgentList []ResponseAgentGet `json:"agentlist"`
}
