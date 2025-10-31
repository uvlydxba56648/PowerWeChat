package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/openWork/license/model"
)

type ResponseGetAppLicenseInfo struct {
	response.ResponseWork
	model.LicenseInfo
}
