package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/power"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
)

type ResponseCustomerStrategyGet struct {
	response.ResponseWork

	Strategy *power.HashMap `json:"momentStrategy"`
}
