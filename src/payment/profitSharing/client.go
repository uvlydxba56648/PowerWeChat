package profitSharing

import (
	"context"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	payment "github.com/uvlydxba56648/PowerWeChat/v3/src/payment/kernel"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/payment/profitSharing/request"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/payment/profitSharing/response"
	"net/http"
)

type Client struct {
	*payment.BaseClient
}

func NewClient(app payment.ApplicationPaymentInterface) (*Client, error) {
	baseClient, err := payment.NewBaseClient(app)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
	}, nil
}

// 请求分账.
// 普通商户：https://pay.weixin.qq.com/doc/v3/merchant/4012524936
// 服务商：https://pay.weixin.qq.com/doc/v3/partner/4012690683

func (comp *Client) Share(ctx context.Context, param *request.RequestShare) (*response.ResponseProfitSharingOrder, error) {

	result := &response.ResponseProfitSharingOrder{}

	if param.AppID == "" {
		config := comp.App.GetConfig()
		param.AppID = config.GetString("app_id", "")
	}

	//options, err := object.StructToHashMapWithTag(param,"json")
	options, err := object.StructToHashMap(param)

	endpoint := comp.Wrap("/v3/profitsharing/orders")
	_, err = comp.Request(ctx, endpoint, nil, http.MethodPost, options, false, nil, result)

	return result, err
}

// 查询分账结果
// 普通商户：https://pay.weixin.qq.com/doc/v3/merchant/4012525210
// 服务商：https://pay.weixin.qq.com/doc/v3/partner/4012466850

func (comp *Client) Query(ctx context.Context, transactionID, outOrderNO string) (*response.ResponseProfitSharingOrder, error) {

	result := &response.ResponseProfitSharingOrder{}

	params := &object.StringMap{
		"transaction_id": transactionID,
	}
	config := comp.App.GetConfig()
	subMchID := config.GetString("sub_mch_id", "")
	if subMchID != "" {
		// 服务商模式下需要此参数
		(*params)["sub_mchid"] = subMchID
	}

	endpoint := comp.Wrap(fmt.Sprintf("/v3/profitsharing/orders/%s", outOrderNO))
	_, err := comp.Request(ctx, endpoint, params, http.MethodGet, &object.HashMap{}, false, nil, result)

	return result, err
}

// 请求分账回退（ApiV3）.
// 普通商户：https://pay.weixin.qq.com/doc/v3/merchant/4012525287
// 服务商：https://pay.weixin.qq.com/doc/v3/partner/4012466854

func (comp *Client) ReturnOrders(ctx context.Context, param *request.RequestShareReturns) (*response.ResponseProfitSharingReturnOrders, error) {

	result := &response.ResponseProfitSharingReturnOrders{}

	//options, err := object.StructToHashMapWithTag(param,"json")
	options, err := object.StructToHashMap(param)

	endpoint := comp.Wrap("/v3/profitsharing/return-orders")
	_, err = comp.Request(ctx, endpoint, nil, http.MethodPost, options, false, nil, result)

	return result, err
}

// Share Return.
// https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_7&index=8
func (comp *Client) Return(ctx context.Context, data *request.RequestShareReturn) (*response.ResponseProfitSharingReturn, error) {

	result := &response.ResponseProfitSharingReturn{}

	config := comp.App.GetConfig()

	params, err := object.StructToHashMapWithXML(data)
	if err != nil {
		return nil, err
	}
	base := &object.HashMap{
		"return_amount": 1,
		"appid":         config.GetString("app_id", ""),
		"mch_id":        config.GetString("mch_id", ""),
	}
	params = object.MergeHashMap(params, base)

	endpoint := comp.Wrap("/secapi/pay/profitsharingreturn")
	_, err = comp.SafeRequest(ctx, endpoint, params, http.MethodPost, &object.HashMap{}, nil, result)

	return result, err
}

// 查询分账回退结果
// 普通商户：https://pay.weixin.qq.com/doc/v3/merchant/4012526279
// 服务商：https://pay.weixin.qq.com/doc/v3/partner/4012466858

func (comp *Client) QueryReturn(ctx context.Context, outOrderNO, outReturnNO string) (*response.ResponseProfitSharingReturnOrders, error) {

	result := &response.ResponseProfitSharingReturnOrders{}

	params := &object.StringMap{
		"out_order_no": outOrderNO,
	}
	config := comp.App.GetConfig()
	subMchID := config.GetString("sub_mch_id", "")
	if subMchID != "" {
		// 服务商模式下需要此参数
		(*params)["sub_mchid"] = subMchID
	}

	endpoint := comp.Wrap(fmt.Sprintf("/v3/profitsharing/return-orders/%s", outReturnNO))
	_, err := comp.Request(ctx, endpoint, params, http.MethodGet, &object.HashMap{}, false, nil, result)

	return result, err
}

// 解冻剩余资金
// 普通上商户：https://pay.weixin.qq.com/doc/v3/merchant/4012526374
// 服务商：https://pay.weixin.qq.com/doc/v3/partner/4012466860

func (comp *Client) UnfreezeOrders(ctx context.Context, transactionID, outOrderNO, description string) (*response.ResponseProfitSharingOrder, error) {

	result := &response.ResponseProfitSharingOrder{}

	options := &object.HashMap{
		"transaction_id": transactionID,
		"out_order_no":   outOrderNO,
		"description":    description,
	}

	config := comp.App.GetConfig()
	subMchID := config.GetString("sub_mch_id", "")
	if subMchID != "" {
		// 服务商模式下需要此参数
		(*options)["sub_mchid"] = subMchID
	}

	endpoint := comp.Wrap("/v3/profitsharing/orders/unfreeze")
	_, err := comp.Request(ctx, endpoint, nil, http.MethodPost, options, false, nil, result)

	return result, err
}

// 查询剩余待分金额.
// 普通商户：https://pay.weixin.qq.com/doc/v3/merchant/4012457939
// 服务商：https://pay.weixin.qq.com/doc/v3/partner/4012457927

func (comp *Client) QueryTransactions(ctx context.Context, transactionID string) (*response.ResponseProfitSharingTransaction, error) {

	result := &response.ResponseProfitSharingTransaction{}

	endpoint := comp.Wrap(fmt.Sprintf("/v3/profitsharing/transactions/%s/amounts", transactionID))
	_, err := comp.Request(ctx, endpoint, nil, http.MethodGet, &object.HashMap{}, false, nil, result)

	return result, err
}

// 添加分账接收方
// 普通商户：https://pay.weixin.qq.com/doc/v3/merchant/4012528995
// 服务商：https://pay.weixin.qq.com/doc/v3/partner/4012690944
func (comp *Client) AddReceiver(
	ctx context.Context, subMchID, subAppID string,
	receiverType string, account string, name string,
	relationType string, customRelation string) (*response.ResponseProfitSharingAddReceiver, error) {

	result := &response.ResponseProfitSharingAddReceiver{}

	config := comp.App.GetConfig()
	options := &object.HashMap{
		"appid":           config.GetString("app_id", ""),
		"type":            receiverType,
		"account":         account,
		"name":            name,
		"relation_type":   relationType,
		"custom_relation": customRelation,
	}

	if subMchID != "" {
		// 服务商模式下需要此参数
		(*options)["sub_mchid"] = subMchID
	}
	if subAppID != "" {
		// 服务商模式下需要此参数
		(*options)["sub_appid"] = subAppID
	}

	endpoint := comp.Wrap("/v3/profitsharing/receivers/add")
	_, err := comp.Request(ctx, endpoint, nil, http.MethodPost, options, false, nil, result)

	return result, err
}

// 删除分账接收方
// 普通商户：https://pay.weixin.qq.com/doc/v3/merchant/4012529590
// 服务商：https://pay.weixin.qq.com/doc/v3/partner/4012466868
func (comp *Client) DeleteReceiver(
	ctx context.Context, subMchID, subAppID string,
	receiverType string, account string) (*response.ResponseProfitSharingDeleteReceiver, error) {

	result := &response.ResponseProfitSharingDeleteReceiver{}

	config := comp.App.GetConfig()
	options := &object.HashMap{
		"appid":   config.GetString("app_id", ""),
		"type":    receiverType,
		"account": account,
	}

	if subMchID != "" {
		// 服务商模式下需要此参数
		(*options)["sub_mchid"] = subMchID
	}
	if subAppID != "" {
		// 服务商模式下需要此参数
		(*options)["sub_appid"] = subAppID
	}

	endpoint := comp.Wrap("/v3/profitsharing/receivers/delete")
	_, err := comp.Request(ctx, endpoint, nil, http.MethodPost, options, false, nil, result)

	return result, err
}

// 申请分账账单
// 普通商户：https://pay.weixin.qq.com/doc/v3/merchant/4012529628
// 服务商：https://pay.weixin.qq.com/doc/v3/partner/4012761140

func (comp *Client) GetBills(ctx context.Context, subMchID string, billDate string, tarType string) (*response.ResponseProfitSharingGetBills, error) {

	result := &response.ResponseProfitSharingGetBills{}

	params := &object.StringMap{
		"bill_date": billDate,
		"tar_type":  tarType,
	}

	if subMchID != "" {
		// 服务商模式下需要此参数
		(*params)["sub_mchid"] = subMchID
	}

	endpoint := comp.Wrap("/v3/profitsharing/bills")
	_, err := comp.Request(ctx, endpoint, params, http.MethodGet, &object.HashMap{}, false, nil, result)

	return result, err
}
