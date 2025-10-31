package nfcUrlLink

import (
	"context"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/nfcUrlLink/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/nfcUrlLink/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// 获取用于 NFC 的小程序 scheme 码，用于写入 NFC 标签
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-scheme/urlscheme.generate.html
func (comp *Client) GenerateNFCScheme(ctx context.Context, options *request.NFCUrlLinkGenerate) (*response.ResponseNFCUrlLink, error) {

	result := &response.ResponseNFCUrlLink{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/generatenfcscheme", options, nil, nil, &result)

	return result, err
}
