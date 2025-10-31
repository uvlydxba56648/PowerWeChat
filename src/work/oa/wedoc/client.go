package wedoc

import (
	"context"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/work/oa/wedoc/request"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/work/oa/wedoc/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
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

// 创建收集表
// https://developer.work.weixin.qq.com/document/path/97462
func (comp *Client) CreateForm(ctx context.Context, options *request.RequestWeDocCreateForm) (*response.ResponseWeDocCreateForm, error) {

	result := &response.ResponseWeDocCreateForm{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/wedoc/create_form", options, nil, nil, result)

	return result, err
}
