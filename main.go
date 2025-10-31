package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	fmt2 "github.com/ArtisanCloud/PowerLibs/v3/fmt"

	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/power"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/miniProgram"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/officialAccount"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/officialAccount/templateMessage/request"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/openPlatform"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/openWork"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/payment"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/work"
)

func GetOfficialConfig() *officialAccount.UserConfig {
	return &officialAccount.UserConfig{
		AppID:  os.Getenv("official_app_id"), // 小程序、公众号或者企业微信的appid
		Secret: os.Getenv("official_secret"), // 商户号 appID

		Token:  "",
		AESKey: "",
		Log: officialAccount.Log{
			Level:  "debug",
			Stdout: false,
			File:   "~/wechat/official-account/info.log",
			Error:  "~/wechat/official-account/error.log",
		},
		Http: officialAccount.Http{
			ProxyURI: "",
		},

		// ResponseType: os.Getenv("response_type"),
		Cache: kernel.NewRedisClient(&kernel.UniversalOptions{
			Addrs:    []string{"127.0.0.1:6379"},
			Password: "",
			DB:       1,
		}),
		HttpDebug: true,
		Debug:     false,
	}
}

func GetWorkConfig() *work.UserConfig {
	agentID, _ := strconv.Atoi(os.Getenv("wecom_agent_id"))
	return &work.UserConfig{
		CorpID:  os.Getenv("corp_id"),
		AgentID: agentID,
		Secret:  os.Getenv("wecom_secret"),

		ResponseType: os.Getenv("array"),
		Log: work.Log{
			Level: "debug",
			//File:  "./wechat/info.log",
			//Error: "./wechat/error.log",
			File:  "~/wechat/work/info.log",
			Error: "~/wechat/work/error.log",
			ENV:   os.Getenv("work.env"),
		},

		Http: work.Http{
			ProxyURI: "",
		},

		OAuth: work.OAuth{
			Callback: os.Getenv("app_oauth_callback_url"),
			Scopes:   []string{},
		},
		Cache: kernel.NewRedisClient(&kernel.UniversalOptions{
			Addrs:    []string{"127.0.0.1:6379"},
			Password: "",
			DB:       1,
		}),

		// HttpDebug: true,
		Debug: true,

		// server config
		Token:  os.Getenv("app_message_token"),
		AESKey: os.Getenv("app_message_aes_key"),
	}
}

func GetPaymentConfig() *payment.UserConfig {
	return &payment.UserConfig{
		//"corp_id":        os.Getenv("corp_id"),
		//"secret":         os.Getenv("secret"),
		AppID:              os.Getenv("app_id"),
		MchID:              os.Getenv("mch_id"),
		MchApiV3Key:        os.Getenv("mch_api_v3_key"),
		Key:                os.Getenv("key"),
		CertPath:           os.Getenv("wx_cert_path"),
		KeyPath:            os.Getenv("wx_key_path"),
		SerialNo:           os.Getenv("serial_no"),
		CertificateKeyPath: os.Getenv("certificate_key_path"),
		WechatPaySerial:    os.Getenv("wechat_pay_serial"),
		RSAPublicKeyPath:   os.Getenv("rsa_public_key_path"),

		ResponseType: os.Getenv("array"),
		Log: payment.Log{
			Level: "debug",
			//File:  "./wechat/info.log",
			//Error: "./wechat/error.log",
			File:  "~/wechat/payment/info.log",
			Error: "~/wechat/payment/error.log",
		},
		Http: payment.Http{
			Timeout:  30.0,
			BaseURI:  "https://api.mch.weixin.qq.com",
			ProxyURI: "",
		},

		Cache: kernel.NewRedisClient(&kernel.UniversalOptions{
			Addrs:    []string{"127.0.0.1:6379"},
			Password: "",
			DB:       1,
		}),
		NotifyURL: os.Getenv("notify_url"),
		HttpDebug: true,
		//Debug: true,
		//"sandbox": true,

		// server config
		// Token:            os.Getenv("token"),
		// AESKey:           os.Getenv("aes_key"),

	}
}

func GetMiniProgramConfig() *miniProgram.UserConfig {
	return &miniProgram.UserConfig{
		AppID:  os.Getenv("miniprogram_app_id"), // 小程序、公众号或者企业微信的appid
		Secret: os.Getenv("miniprogram_secret"), // 商户号 appID

		ResponseType: os.Getenv("array"),
		Log: miniProgram.Log{
			Level: "debug",
			File:  "~/wechat/mini-program/info.log",
			Error: "~/wechat/mini-program/error.log",
		},
		Cache: kernel.NewRedisClient(&kernel.UniversalOptions{
			Addrs:    []string{"127.0.0.1:6379"},
			Password: "",
			DB:       1,
		}),
		Http: miniProgram.Http{
			ProxyURI: "",
		},
		HttpDebug: true,
		//Debug: true,
		//"sandbox": true,

	}
}

func GetOpenPlatformConfig() *openPlatform.UserConfig {
	return &openPlatform.UserConfig{
		AppID:        "123",
		Secret:       "321",
		AuthCode:     "123",
		AESKey:       "321",
		ResponseType: os.Getenv("array"),
		Log: openPlatform.Log{
			Level: "debug",
			//File:  "./wechat/info.log",
			//Error: "./wechat/error.log",
			File:  "~/wechat/platform/info.log",
			Error: "~/wechat/platform/error.log",
		},
		Cache: kernel.NewRedisClient(&kernel.UniversalOptions{
			Addrs:    []string{"127.0.0.1:6379"},
			Password: "",
			DB:       1,
		}),
		Http: openPlatform.Http{
			ProxyURI: "",
		},
		// OAuth:        "",
		// HttpDebug:    "",
		// Debug:        "",
		// NotifyURL:    "",
		// Sandbox:      "",
	}
}

func GetOpenWorkConfig() *openWork.UserConfig {
	return &openWork.UserConfig{
		AppID:        "123",
		Secret:       "321",
		AuthCode:     "123",
		AESKey:       "321",
		ResponseType: os.Getenv("array"),

		Cache: kernel.NewRedisClient(&kernel.UniversalOptions{
			Addrs:    []string{"127.0.0.1:6379"},
			Password: "",
			DB:       1,
		}),

		// OAuth:        "",
		// HttpDebug:    "",
		// Debug:        "",
		// NotifyURL:    "",
		// Sandbox:      "",
	}
}

func main() {
	fmt.Printf("hello Wechat! \n")

	// init officialAccount app
	configOfficialAccount := GetOfficialConfig()
	officialAccountApp, err := officialAccount.NewOfficialAccount(configOfficialAccount)
	if err != nil {
		fmt.Println(err.Error())
	}
	ctx := context.Background()
	officialAccountApp.TemplateMessage.Send(ctx, &request.RequestTemlateMessage{
		ToUser:     "",
		TemplateID: "",
		MiniProgram: &request.MiniProgram{ // 上线后的小程序才可以使用
			AppID:    "",
			PagePath: fmt.Sprintf("pages/order/details?id=%v", 111111),
		},
		Data: &power.HashMap{
			"character_string1": &power.HashMap{
				"value": "111111111",
				"color": "#173177",
			},
			"thing3": &power.HashMap{
				"value": "可口可乐大瓶装",
				"color": "#173177",
			},
			"const2": &power.HashMap{
				"value": "订单超时未支付",
				"color": "#173177",
			},
		},
	})
	fmt2.Dump("officialAccount config:", officialAccountApp.GetConfig().All())
	// init wecom app
	configWecom := GetWorkConfig()
	wecomApp, err := work.NewWork(configWecom)
	if err != nil {
		panic(err.Error())
	}
	fmt2.Dump("wecom config:", wecomApp.GetConfig().All())

	// init payment app
	configPayment := GetPaymentConfig()
	paymentApp, err := payment.NewPayment(configPayment)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt2.Dump("payment config:", paymentApp.GetConfig().All())

	// init miniProgram app
	configMiniProgram := GetMiniProgramConfig()
	miniProgramApp, err := miniProgram.NewMiniProgram(configMiniProgram)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt2.Dump("miniprogram config:", miniProgramApp.GetConfig().All())

	// init openplatform app
	configOpenPlatform := GetOpenPlatformConfig()
	openPlatform, err := openPlatform.NewOpenPlatform(configOpenPlatform)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt2.Dump("openPlatform config:", openPlatform.GetConfig().All())

	// init openwork app
	configOpenWork := GetOpenWorkConfig()
	openWork, err := openWork.NewOpenWork(configOpenWork)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt2.Dump("openPlatform config:", openWork.GetConfig().All())

}
