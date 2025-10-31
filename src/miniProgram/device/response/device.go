package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
)

type ResponseCreateIotGroupId struct {
	response.ResponseMiniProgram
	GroupId string `json:"group_id"`
}

type Device struct {
	ModelId string `json:"model_id"`
	Sn      string `json:"sn"`
}

type ResponseGetIotGroupInfo struct {
	response.ResponseMiniProgram
	GroupName  string   `json:"group_name"`
	ModelId    string   `json:"model_id"`
	ModelType  string   `json:"model_type"`
	DeviceList []Device `json:"device_list"`
}

type ResponseAddIotGroupDevice struct {
	response.ResponseMiniProgram
	DeviceList []struct {
		ModelId string `json:"model_id"`
		Sn      string `json:"sn"`
		Errcode int    `json:"errcode"`
	} `json:"device_list"`
}
type ResponseRemoveIotGroupDevice struct {
	response.ResponseMiniProgram
	DeviceList []struct {
		ModelId string `json:"model_id"`
		Sn      string `json:"sn"`
		Errcode int    `json:"errcode"`
	} `json:"device_list"`
}

type ResponseGetLicensePkgList struct {
	response.ResponseMiniProgram
	PkgList []struct {
		PkgId     string `json:"pkg_id"`
		PkgType   int    `json:"pkg_type"`
		StartTime int    `json:"start_time"`
		EndTime   int    `json:"end_time"`
		PkgStatus int    `json:"pkg_status"`
		Used      int    `json:"used"`
		All       int    `json:"all"`
	} `json:"pkg_list"`
	MaxActiveNumber int `json:"max_active_number"`
}

type ResponseActiveLicenseDevice struct {
	response.ResponseMiniProgram
	DeviceList []Device `json:"device_list"`
}

type ResponseGetLicenseDeviceInfo struct {
	response.ResponseMiniProgram
	DeviceList []struct {
		ModelId    string `json:"model_id"`
		Sn         string `json:"sn"`
		ExpireTime int    `json:"expire_time"`
	} `json:"device_list"`
}
