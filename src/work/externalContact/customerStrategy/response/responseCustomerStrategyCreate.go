package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/power"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
)

type ResponseCustomerStrategyCreate struct {
	response.ResponseWork

	StrategyID *power.HashMap `json:"strategy_id"`
}
