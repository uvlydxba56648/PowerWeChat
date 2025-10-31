package wxaSecOrder

import (
	"context"

	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/miniProgram/wxa/sec/order/request"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/miniProgram/wxa/sec/order/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// 发货信息录入接口
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/business-capabilities/order-shipping/order-shipping.html#%E4%B8%80%E3%80%81%E5%8F%91%E8%B4%A7%E4%BF%A1%E6%81%AF%E5%BD%95%E5%85%A5%E6%8E%A5%E5%8F%A3
func (wxa *Client) UploadShippingInfo(ctx context.Context, options *request.RequestUploadShippingInfo) (*response.ResponseUploadShippingInfo, error) {
	result := &response.ResponseUploadShippingInfo{}

	_, err := wxa.BaseClient.HttpPostJson(ctx, "wxa/sec/order/upload_shipping_info", options, nil, nil, result)

	return result, err
}

// 发货信息合单录入接口
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/business-capabilities/order-shipping/order-shipping.html#%E4%BA%8C%E3%80%81%E5%8F%91%E8%B4%A7%E4%BF%A1%E6%81%AF%E5%90%88%E5%8D%95%E5%BD%95%E5%85%A5%E6%8E%A5%E5%8F%A3
func (wxa *Client) UploadCombinedShippingInfo(ctx context.Context, options *request.RequestUploadCombinedShippingInfo) (*response.ResponseUploadCombinedShippingInfo, error) {
	result := &response.ResponseUploadCombinedShippingInfo{}

	_, err := wxa.BaseClient.HttpPostJson(ctx, "wxa/sec/order/upload_combined_shipping_info", options, nil, nil, result)

	return result, err
}

// 查询订单发货状态
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/business-capabilities/order-shipping/order-shipping.html#%E4%B8%89%E3%80%81%E6%9F%A5%E8%AF%A2%E8%AE%A2%E5%8D%95%E5%8F%91%E8%B4%A7%E7%8A%B6%E6%80%81
func (wxa *Client) GetOrder(ctx context.Context, options *request.RequestGetOrder) (*response.ResponseGetOrder, error) {
	result := &response.ResponseGetOrder{}

	_, err := wxa.BaseClient.HttpPostJson(ctx, "wxa/sec/order/get_order", options, nil, nil, result)

	return result, err
}

// 查询订单列表
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/business-capabilities/order-shipping/order-shipping.html#%E5%9B%9B%E3%80%81%E6%9F%A5%E8%AF%A2%E8%AE%A2%E5%8D%95%E5%88%97%E8%A1%A8
func (wxa *Client) GetOrderList(ctx context.Context, options *request.RequestGetOrderList) (*response.ResponseGetOrderList, error) {
	result := &response.ResponseGetOrderList{}

	_, err := wxa.BaseClient.HttpPostJson(ctx, "wxa/sec/order/get_order_list", options, nil, nil, result)

	return result, err
}

// 确认收货提醒接口
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/business-capabilities/order-shipping/order-shipping.html#%E4%BA%94%E3%80%81%E7%A1%AE%E8%AE%A4%E6%94%B6%E8%B4%A7%E6%8F%90%E9%86%92%E6%8E%A5%E5%8F%A3
func (wxa *Client) NotifyConfirmReceive(ctx context.Context, options *request.RequestNotifyConfirmReceive) (*response.ResponseNotifyConfirmReceive, error) {
	result := &response.ResponseNotifyConfirmReceive{}

	_, err := wxa.BaseClient.HttpPostJson(ctx, "wxa/sec/order/notify_confirm_receive", options, nil, nil, result)

	return result, err
}

// 消息跳转路径设置接口
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/business-capabilities/order-shipping/order-shipping.html#%E5%85%AD%E3%80%81%E6%B6%88%E6%81%AF%E8%B7%B3%E8%BD%AC%E8%B7%AF%E5%BE%84%E8%AE%BE%E7%BD%AE%E6%8E%A5%E5%8F%A3
func (wxa *Client) SetMsgJumpPath(ctx context.Context, options *request.RequestSetMsgJumpPath) (*response.ResponseSetMsgJumpPath, error) {
	result := &response.ResponseSetMsgJumpPath{}

	_, err := wxa.BaseClient.HttpPostJson(ctx, "wxa/sec/order/set_msg_jump_path", options, nil, nil, result)

	return result, err
}

// 查询小程序是否已开通发货信息管理服务
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/business-capabilities/order-shipping/order-shipping.html#%E4%B8%83%E3%80%81%E6%9F%A5%E8%AF%A2%E5%B0%8F%E7%A8%8B%E5%BA%8F%E6%98%AF%E5%90%A6%E5%B7%B2%E5%BC%80%E9%80%9A%E5%8F%91%E8%B4%A7%E4%BF%A1%E6%81%AF%E7%AE%A1%E7%90%86%E6%9C%8D%E5%8A%A1
func (wxa *Client) IsTradeManaged(ctx context.Context, options *request.RequestIsTradeManaged) (*response.ResponseIsTradeManaged, error) {

	result := &response.ResponseIsTradeManaged{}

	_, err := wxa.BaseClient.HttpPostJson(ctx, "wxa/sec/order/is_trade_managed", options, nil, nil, result)

	return result, err
}

// 查询小程序是否已完成交易结算管理确认
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/business-capabilities/order-shipping/order-shipping.html#%E5%85%AB%E3%80%81%E6%9F%A5%E8%AF%A2%E5%B0%8F%E7%A8%8B%E5%BA%8F%E6%98%AF%E5%90%A6%E5%B7%B2%E5%AE%8C%E6%88%90%E4%BA%A4%E6%98%93%E7%BB%93%E7%AE%97%E7%AE%A1%E7%90%86%E7%A1%AE%E8%AE%A4
func (wxa *Client) IsTradeManagementConfirmationCompleted(ctx context.Context, options *request.RequestIsTradeManagementConfirmationCompleted) (*response.ResponseIsTradeManagementConfirmationCompleted, error) {
	result := &response.ResponseIsTradeManagementConfirmationCompleted{}

	_, err := wxa.BaseClient.HttpPostJson(ctx, "wxa/sec/order/is_trade_management_confirmation_completed", options, nil, nil, result)

	return result, err
}

// 特殊发货报备
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/business-capabilities/order-shipping/order-shipping.html#%E5%8D%81%E3%80%81%E7%89%B9%E6%AE%8A%E5%8F%91%E8%B4%A7%E6%8A%A5%E5%A4%87
func (wxa *Client) OpSpecialOrder(ctx context.Context, options *request.RequestOpSpecialOrder) (*response.ResponseOpSpecialOrder, error) {
	result := &response.ResponseOpSpecialOrder{}

	_, err := wxa.BaseClient.HttpPostJson(ctx, "wxa/sec/order/opspecialorder", options, nil, nil, result)

	return result, err
}
