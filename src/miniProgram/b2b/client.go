package b2b

import (
	"context"
	"encoding/json"

	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/miniProgram/b2b/request"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/miniProgram/b2b/response"
)

type Client struct {
	BaseClient *kernel.BaseClient

	appKey        string
	sandboxAppKey string
}

// CreatePayment 创建 B2b 支付 https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/B2b_store_assistant.html#_3-3-%E6%94%AF%E4%BB%98%E4%B8%8E%E9%80%80%E6%AC%BE
func (c *Client) CreatePayment(sessionKey string, data any, overAppKey ...string) (*object.StringMap, error) {
	appKey := c.appKey
	if len(overAppKey) > 0 && overAppKey[0] != "" {
		appKey = overAppKey[0]
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	strData := string(jsonData)
	return &object.StringMap{
		"sign_data": strData,
		"mode":      "retail_pay_goods",
		"signature": kernel.CalcSignature(strData, sessionKey),
		"pay_sig":   kernel.CalcSignature("requestCommonPayment"+"&"+strData, appKey),
	}, nil
}

func postJson[I, O any](ctx context.Context, c *Client, path string, in *I) (out *O, err error) {
	postBody, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}
	signPost := string(postBody)

	paySign := kernel.CalcSignature(path+"&"+signPost, c.appKey)
	query := &object.StringMap{
		"pay_sig": paySign,
	}
	out = new(O)
	_, err = c.BaseClient.HttpPostJson(ctx, path, signPost, query, nil, out)
	return out, err
}

// GetOrder 查询订单 https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/B2b_store_assistant.html#_3-3-%E6%94%AF%E4%BB%98%E4%B8%8E%E9%80%80%E6%AC%BE
func (c *Client) GetOrder(ctx context.Context, in *request.GetOrderReq) (*response.GetOrderRes, error) {
	return postJson[request.GetOrderReq, response.GetOrderRes](ctx, c, "/retail/B2b/getorder", in)
}

// Refund 退款
func (c *Client) Refund(ctx context.Context, in *request.RefundReq) (*response.RefundRes, error) {
	if in.AppKey != "" {
		c.appKey = in.AppKey
	}
	return postJson[request.RefundReq, response.RefundRes](ctx, c, "/retail/B2b/refund", in)
}

// GetRefund 查询退款
func (c *Client) GetRefund(ctx context.Context, in *request.GetRefundReq) (*response.GetRefundRes, error) {
	return postJson[request.GetRefundReq, response.GetRefundRes](ctx, c, "/retail/B2b/getrefund", in)
}

// AddProfitSharingAccount 添加分账方
func (c *Client) AddProfitSharingAccount(ctx context.Context, in *request.AddProfitSharingAccountReq) (*response.AddProfitSharingAccountRes, error) {
	return postJson[request.AddProfitSharingAccountReq, response.AddProfitSharingAccountRes](ctx, c, "/retail/B2b/addprofitsharingaccount", in)
}

// DelProfitSharingAccount 删除分账方
func (c *Client) DelProfitSharingAccount(ctx context.Context, in *request.DelProfitSharingAccountReq) (*response.DelProfitSharingAccountRes, error) {
	return postJson[request.DelProfitSharingAccountReq, response.DelProfitSharingAccountRes](ctx, c, "/retail/B2b/delprofitsharingaccount", in)
}

// QueryProfitSharingAccount 查询分账方
func (c *Client) QueryProfitSharingAccount(ctx context.Context, in *request.QueryProfitSharingAccountReq) (*response.QueryProfitSharingAccountRes, error) {
	return postJson[request.QueryProfitSharingAccountReq, response.QueryProfitSharingAccountRes](ctx, c, "/retail/B2b/queryprofitsharingaccount", in)
}

// CreateProfitSharingOrder 请求分账
func (c *Client) CreateProfitSharingOrder(ctx context.Context, in *request.CreateProfitSharingOrderReq) (*response.CreateProfitSharingOrderRes, error) {
	return postJson[request.CreateProfitSharingOrderReq, response.CreateProfitSharingOrderRes](ctx, c, "/retail/B2b/createprofitsharingorder", in)
}

// QueryProfitSharingOrder 查询分账结果
func (c *Client) QueryProfitSharingOrder(ctx context.Context, in *request.QueryProfitSharingOrderReq) (*response.QueryProfitSharingOrderRes, error) {
	return postJson[request.QueryProfitSharingOrderReq, response.QueryProfitSharingOrderRes](ctx, c, "/retail/B2b/queryprofitsharingorder?", in)
}

// QueryProfitSharingRemainAmt 查询分账剩余金额
func (c *Client) QueryProfitSharingRemainAmt(ctx context.Context, in *request.QueryProfitSharingRemainAmtReq) (*response.QueryProfitSharingRemainAmtRes, error) {
	return postJson[request.QueryProfitSharingRemainAmtReq, response.QueryProfitSharingRemainAmtRes](ctx, c, "/retail/B2b/queryprofitsharingremainamt", in)
}

// FinishProfitSharingOrder 完成分账
func (c *Client) FinishProfitSharingOrder(ctx context.Context, in *request.FinishProfitSharingOrderReq) (*response.FinishProfitSharingOrderRes, error) {
	return postJson[request.FinishProfitSharingOrderReq, response.FinishProfitSharingOrderRes](ctx, c, "/retail/B2b/finishprofitsharingorder", in)
}

// RefundProfitSharing 请求分账回退
func (c *Client) RefundProfitSharing(ctx context.Context, in *request.RefundProfitSharingReq) (*response.RefundProfitSharingRes, error) {
	return postJson[request.RefundProfitSharingReq, response.RefundProfitSharingRes](ctx, c, "/retail/B2b/refundprofitsharing", in)
}

// QueryRefundProfitSharingOrder 查询分账回退结果
func (c *Client) QueryRefundProfitSharingOrder(ctx context.Context, in *request.QueryRefundProfitSharingOrderReq) (*response.QueryRefundProfitSharingOrderRes, error) {
	return postJson[request.QueryRefundProfitSharingOrderReq, response.QueryRefundProfitSharingOrderRes](ctx, c, "/retail/B2b/queryrefundprofitsharingorder", in)
}

// DownloadBill 查看账单
func (c *Client) DownloadBill(ctx context.Context, in *request.DownloadBillReq) (*response.DownloadBillRes, error) {
	return postJson[request.DownloadBillReq, response.DownloadBillRes](ctx, c, "/retail/B2b/downloadbill", in)
}
