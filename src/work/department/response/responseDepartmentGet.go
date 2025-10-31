package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/models"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
)

type ResponseDepartmentGet struct {
	response.ResponseWork
	Department *models.Department `json:"department"`
}
