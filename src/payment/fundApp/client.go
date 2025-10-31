package fundApp

import (
	"context"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/power"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/payment/fundApp/request"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/payment/fundApp/response"
	payment "github.com/uvlydxba56648/PowerWeChat/v3/src/payment/kernel"

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

// 发起转账
// https://pay.weixin.qq.com/doc/v3/merchant/4012716434
func (comp *Client) TransferBills(ctx context.Context, data *request.RequestTransferBills) (*response.ResponseTransferBills, error) {

	result := &response.ResponseTransferBills{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	endpoint := comp.Wrap("/v3/fund-app/mch-transfer/transfer-bills")
	_, err = comp.SafeRequestV3(ctx, endpoint, nil, http.MethodPost, params, nil, result)

	return result, err
}

// 商户单号查询转账单
// https://pay.weixin.qq.com/doc/v3/merchant/4012716437
func (comp *Client) QueryOutBill(ctx context.Context, outBillNO string) (*response.ResponseQueryOutBill, error) {

	result := &response.ResponseQueryOutBill{}

	endpoint := comp.Wrap(fmt.Sprintf("/v3/fund-app/mch-transfer/transfer-bills/out-bill-no/%s", outBillNO))
	_, err := comp.SafeRequestV3(ctx, endpoint, nil, http.MethodGet, &object.HashMap{}, nil, result)

	return result, err
}

// 微信单号查询转账单
// https://pay.weixin.qq.com/doc/v3/merchant/4012716457
func (comp *Client) QueryTransferBill(ctx context.Context, transferBillNO string) (*response.ResponseQueryTransferBill, error) {

	result := &response.ResponseQueryTransferBill{}

	endpoint := comp.Wrap(fmt.Sprintf("/v3/fund-app/mch-transfer/transfer-bills/transfer-bill-no/%s", transferBillNO))
	_, err := comp.SafeRequestV3(ctx, endpoint, nil, http.MethodGet, &object.HashMap{}, nil, result)

	return result, err
}

// 撤销转账
// https://pay.weixin.qq.com/doc/v3/merchant/4012716458
func (comp *Client) Cancel(ctx context.Context, outBillNO string) (*response.ResponseCancelBill, error) {

	result := &response.ResponseCancelBill{}

	endpoint := comp.Wrap(fmt.Sprintf("/v3/fund-app/mch-transfer/transfer-bills/out-bill-no/%s/cancel", outBillNO))
	_, err := comp.SafeRequestV3(ctx, endpoint, nil, http.MethodPost, &object.HashMap{}, nil, result)

	return result, err
}

// 商户单号申请电子回单
// https://pay.weixin.qq.com/doc/v3/merchant/4012716452
func (comp *Client) ApplyForElecSign(ctx context.Context, outBillNO string) (*response.ResponseApplyForElecSign, error) {

	result := &response.ResponseApplyForElecSign{}

	endpoint := "/v3/fund-app/mch-transfer/elecsign/out-bill-no"
	_, err := comp.SafeRequestV3(ctx, endpoint, nil, http.MethodPost, &object.HashMap{
		"out_bill_no": outBillNO,
	}, nil, result)

	return result, err
}

// 商户单号查询电子回单
// https://pay.weixin.qq.com/doc/v3/merchant/4012716436
func (comp *Client) QueryElecSign(ctx context.Context, outBillNO string) (*response.ResponseQueryElecSign, error) {
	result := &response.ResponseQueryElecSign{}
	endpoint := comp.Wrap(fmt.Sprintf("/v3/fund-app/mch-transfer/elecsign/out-bill-no/%s", outBillNO))
	_, err := comp.SafeRequestV3(ctx, endpoint, nil, http.MethodGet, &object.HashMap{}, nil, result)
	return result, err
}

// 微信单号申请电子回单
// https://pay.weixin.qq.com/doc/v3/merchant/4012716456
func (comp *Client) ApplyForElecSignByTransferBillNo(ctx context.Context, transferBillNO string) (*response.ResponseApplyForElecSign, error) {
	result := &response.ResponseApplyForElecSign{}
	endpoint := "/v3/fund-app/mch-transfer/elecsign/transfer-bill-no"
	_, err := comp.SafeRequestV3(ctx, endpoint, nil, http.MethodPost, &object.HashMap{
		"transfer_bill_no": transferBillNO,
	}, nil, result)
	return result, err
}

// 微信单号查询电子回单
// https://pay.weixin.qq.com/doc/v3/merchant/4012716455
func (comp *Client) QueryElecSignByTransferBillNo(ctx context.Context, transferBillNO string) (*response.ResponseQueryElecSign, error) {
	result := &response.ResponseQueryElecSign{}
	endpoint := comp.Wrap(fmt.Sprintf("/v3/fund-app/mch-transfer/elecsign/transfer-bill-no/%s", transferBillNO))
	_, err := comp.SafeRequestV3(ctx, endpoint, nil, http.MethodGet, &object.HashMap{}, nil, result)
	return result, err
}

// 下载电子回单
// https://pay.weixin.qq.com/doc/v3/merchant/4013866774
func (comp *Client) DownloadElecSignToFilePath(ctx context.Context, requestDownload *power.RequestDownload, filePath string) (int64, error) {
	return comp.StreamDownload(ctx, requestDownload, filePath)
}
