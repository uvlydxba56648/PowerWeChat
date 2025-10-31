package featureUnit

import (
	"context"
	"testing"
	"time"

	"github.com/ArtisanCloud/PowerLibs/v3/fmt"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/miniProgram"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/miniProgram/wxa/sec/order/request"
)

var app *miniProgram.MiniProgram
var err error

func init() {
	app, err = miniProgram.NewMiniProgram(&miniProgram.UserConfig{
		AppID:     "", // 小程序、公众号或者企业微信的appid
		Secret:    "", // 商户号 appID
		HttpDebug: true,
		Debug:     true,
	})
	if err != nil {
		panic(err)
	}
}

func TestUploadingShippingOrderInfo(t *testing.T) {
	req := &request.RequestUploadShippingInfo{
		OrderKey: request.OrderKey{
			OrderNumberType: 2,
			TransactionID:   "",
			MchID:           "",
			OutTradeNo:      "",
		},
		LogisticsType:  4,
		DeliveryMode:   1,
		IsAllDelivered: true,
		ShippingList:   make([]request.ShippingList, 0),
		// 2022-12-15T13:29:35.120+08:00
		UploadTime: time.Now().Format(time.RFC3339Nano),
		Payer: struct {
			OpenID string `json:"openid"`
		}{
			OpenID: "",
		},
	}
	req.ShippingList = append(req.ShippingList, request.ShippingList{
		TrackingNo:     "",
		ExpressCompany: "STO",
		ItemDesc:       "到店自提",
		Contact: struct {
			ConsignorContact string `json:"consignor_contact,omitempty"`
			ReceiverContact  string `json:"receiver_contact,omitempty"`
		}{
			ConsignorContact: "河马很忙宝峰店",
			ReceiverContact:  "罗雄丰",
		},
	})
	responseUploadShippingInfo, err := app.WXASecOrder.UploadShippingInfo(context.Background(), req)
	if err != nil {
		panic(err)
	}
	fmt.Dump(responseUploadShippingInfo)
}

func TestIsTradeManaged(t *testing.T) {
	appID := ""
	responseIsTradeManaged, err := app.WXASecOrder.IsTradeManaged(context.Background(), &request.RequestIsTradeManaged{AppID: appID})
	if err != nil {
		panic(err)
	}
	fmt.Dump(responseIsTradeManaged)
}

func TestOpSpecialOrder(t *testing.T) {
	responseOpSpecialOrder, err := app.WXASecOrder.OpSpecialOrder(context.Background(), &request.RequestOpSpecialOrder{
		OrderID: "",
		Types:   2,
	})
	if err != nil {
		panic(err)
	}
	fmt.Dump(responseOpSpecialOrder)
}

func TestGetOrder(t *testing.T) {
	req := &request.RequestGetOrder{
		TransactionID:   "",
		MerchantID:      "",
		SubMerchantID:   "",
		MerchantTradeNo: "",
	}
	responseGetOrder, err := app.WXASecOrder.GetOrder(context.Background(), req)
	if err != nil {
		panic(err)
	}
	fmt.Dump(responseGetOrder)
}

func TestGetOrderList(t *testing.T) {
	req := &request.RequestGetOrderList{
		PayTimeRange: nil,
		OrderState:   0,
		OpenID:       "",
		LastIndex:    "",
		PageSize:     1,
	}
	responseGetOrderList, err := app.WXASecOrder.GetOrderList(context.Background(), req)
	if err != nil {
		panic(err)
	}
	fmt.Dump(responseGetOrderList)
}

func TestIsTradeManagementConfirmationCompleted(t *testing.T) {
	responseIsTradeManagementConfirmationCompleted, err := app.WXASecOrder.IsTradeManagementConfirmationCompleted(context.Background(), &request.RequestIsTradeManagementConfirmationCompleted{
		AppID: "",
	})
	if err != nil {
		panic(err)
	}
	fmt.Dump(responseIsTradeManagementConfirmationCompleted)
}

func TestNotifyConfirmReceive(t *testing.T) {
	responseNotifyConfirmReceive, err := app.WXASecOrder.NotifyConfirmReceive(context.Background(), &request.RequestNotifyConfirmReceive{
		TransactionID:   "",
		MerchantID:      "",
		SubMerchantID:   "",
		MerchantTradeNo: "",
		ReceivedTime:    0,
	})
	if err != nil {
		panic(err)
	}
	fmt.Dump(responseNotifyConfirmReceive)
}

func TestSet(t *testing.T) {
	responseSetMsgJumpPath, err := app.WXASecOrder.SetMsgJumpPath(context.Background(), &request.RequestSetMsgJumpPath{
		Path: "",
	})
	if err != nil {
		panic(err)
	}
	fmt.Dump(responseSetMsgJumpPath)
}
