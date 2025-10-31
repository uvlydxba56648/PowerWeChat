package request

import "github.com/uvlydxba56648/PowerWeChat/v3/src/miniProgram/device/response"

type Data map[string]struct {
	Value string `json:"value"`
}

type RequestSendHardwareDeviceMessage struct {
	ToOpenidList     []string `json:"to_openid_list"`
	Sn               string   `json:"sn"`
	TemplateId       string   `json:"template_id"`
	Page             string   `json:"page"`
	MiniProgramState string   `json:"miniprogram_state"`
	Lang             string   `json:"lang"`
	Data             Data     `json:"data"`
}

type RequestGetSnTicket struct {
	ModelId string `json:"model_id"`
	Sn      string `json:"sn"`
}

type RequestCreateIotGroupId struct {
	GroupName string `json:"group_name"`
	ModelId   string `json:"model_id"`
}

type RequestGetIotGroupInfo struct {
	GroupId string `json:"group_id"`
}

type RequestAddIotGroupDevice struct {
	GroupId    string `json:"group_id"`
	DeviceList []struct {
		ModelId string `json:"model_id"`
		Sn      string `json:"sn"`
	} `json:"device_list"`
}

type RequestRemoveIotGroupDevice struct {
	GroupId    string            `json:"group_id"`
	DeviceList []response.Device `json:"device_list"`
}

type RequestGetLicensePkgList struct {
	PkgType int `json:"pkg_type"`
}

type RequestActiveLicenseDevice struct {
	PkgType    int               `json:"pkg_type"`
	DeviceList []response.Device `json:"device_list"`
}

type RequestGetLicenseDeviceInfo struct {
	DeviceList []response.Device `json:"device_list"`
}
