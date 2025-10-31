package provider

import (
	"context"

	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(app, nil)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
	}, nil
}

// 明文corpid转换为加密corpid
// https://developer.work.weixin.qq.com/document/path/95604
func (clt *Client) CorpIDToOpenCorpID(ctx context.Context, corpID string) (string, error) {
	var result struct {
		response.ResponseWork
		OpenCorpID string `json:"open_corpid,omitempty"`
	}
	req := object.HashMap{
		"corpid": corpID,
	}
	_, err := clt.BaseClient.HttpPostJson(ctx, "cgi-bin/service/corpid_to_opencorpid", &req, nil, nil, &result)

	return result.OpenCorpID, err
}

// GetCustomizedAuthURL 获取带参授权链接
// https://developer.work.weixin.qq.com/document/path/98744
func (clt *Client) GetCustomizedAuthURL(ctx context.Context, state string, templateIDList []string) (string, error) {
	var result struct {
		response.ResponseWork
		QRCodeURL string `json:"qrcode_url,omitempty"`
	}
	req := object.HashMap{
		"state":           state,
		"templateid_list": templateIDList,
	}

	_, err := clt.BaseClient.HttpPostJson(ctx, "cgi-bin/service/get_customized_auth_url", &req, nil, nil, &result)

	return result.QRCodeURL, err
}
