package device

import (
	"context"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/miniProgram/device/request"
	response2 "github.com/uvlydxba56648/PowerWeChat/v3/src/miniProgram/device/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// 发送设备消息
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/hardware-device/sendHardwareDeviceMessage.html
func (comp *Client) SendHardwareDeviceMessage(ctx context.Context, data *request.RequestSendHardwareDeviceMessage) (*response.ResponseMiniProgram, error) {

	var rs = &response.ResponseMiniProgram{}

	_, err := comp.BaseClient.HttpPost(ctx, "cgi-bin/message/device/subscribe/send", data, nil, nil)

	return rs, err
}

// 获取设备票据
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/hardware-device/getSnTicket.html
func (comp *Client) GetSnTicket(ctx context.Context, data *request.RequestGetSnTicket) (*response.ResponseMiniProgram, error) {
	var rs = &response.ResponseMiniProgram{}
	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/getsnticket", data, nil, nil, rs)
	return rs, err
}

// 创建设备组
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/hardware-device/createIotGroupId.html
func (comp *Client) CreateIotGroupId(ctx context.Context, data *request.RequestCreateIotGroupId) (*response2.ResponseCreateIotGroupId, error) {
	var rs = &response2.ResponseCreateIotGroupId{}
	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/business/group/createid", data, nil, nil, rs)
	return rs, err
}

// 查询设备组信息
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/hardware-device/getIotGroupInfo.html
func (comp *Client) GetIotGroupInfo(ctx context.Context, data *request.RequestGetIotGroupInfo) (*response2.ResponseGetIotGroupInfo, error) {
	var rs = &response2.ResponseGetIotGroupInfo{}
	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/business/group/getinfo", data, nil, nil, rs)
	return rs, err
}

// 设备组添加设备
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/hardware-device/addIotGroupDevice.html
func (comp *Client) AddIotGroupDevice(ctx context.Context, data *request.RequestAddIotGroupDevice) (*response2.ResponseAddIotGroupDevice, error) {
	var rs = &response2.ResponseAddIotGroupDevice{}
	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/business/group/adddevice", data, nil, nil, rs)
	return rs, err
}

// 设备组删除设备
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/hardware-device/removeIotGroupDevice.html
func (comp *Client) RemoveIotGroupDevice(ctx context.Context, data *request.RequestRemoveIotGroupDevice) (*response2.ResponseRemoveIotGroupDevice, error) {
	var rs = &response2.ResponseRemoveIotGroupDevice{}
	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/business/group/removedevice", data, nil, nil, rs)
	return rs, err
}

// 查询license资源包列表
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/hardware-device/getLicensePkgList.html
func (comp *Client) GetLicensePkgList(ctx context.Context, data *request.RequestGetLicensePkgList) (*response2.ResponseGetLicensePkgList, error) {
	var rs = &response2.ResponseGetLicensePkgList{}
	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/business/license/getpkglist", data, nil, nil, rs)
	return rs, err
}

// 激活设备license
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/hardware-device/activeLicenseDevice.html
func (comp *Client) ActiveLicenseDevice(ctx context.Context, data *request.RequestActiveLicenseDevice) (*response2.ResponseActiveLicenseDevice, error) {
	var rs = &response2.ResponseActiveLicenseDevice{}
	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/business/license/activedevice", data, nil, nil, rs)
	return rs, err
}

// 查询设备激活详情
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/hardware-device/getLicenseDeviceInfo.html
func (comp *Client) GetLicenseDeviceInfo(ctx context.Context, data *request.RequestGetLicenseDeviceInfo) (*response2.ResponseGetLicenseDeviceInfo, error) {
	var rs = &response2.ResponseGetLicenseDeviceInfo{}
	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/business/license/getdeviceinfo", data, nil, nil, rs)
	return rs, err
}
