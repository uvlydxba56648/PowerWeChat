package store

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/channels/eCommerce/store/request"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/channels/eCommerce/store/response"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// 获取店铺基本信息
// https://developers.weixin.qq.com/doc/store/shop/API/basics/getbasicinfo.html
func (comp *Client) GetBasicInfo(ctx context.Context) (*response.ResponseGetBasicInfo, error) {

	result := &response.ResponseGetBasicInfo{}

	_, err := comp.BaseClient.HttpGet(ctx, "channels/ec/basics/info/get", nil, nil, result)

	return result, err
}

// 获取店铺二维码
// https://developers.weixin.qq.com/doc/store/shop/API/basics/getshopqrcode.html
func (comp *Client) GetShopQRCode(ctx context.Context, data *request.RequestGetShopQRCode) (*response.ResponseGetShopQRCode, error) {

	result := &response.ResponseGetShopQRCode{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPost(ctx, "channels/ec/basics/shop/qrcode/get", params, nil, result)

	return result, err
}
