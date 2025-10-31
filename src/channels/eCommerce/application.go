package eCommerce

import (
	"github.com/ArtisanCloud/PowerLibs/v3/logger"
	"github.com/ArtisanCloud/PowerLibs/v3/logger/contract"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/channels/eCommerce/store"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/providers"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/miniProgram/auth"
	"net/http"
)

type ECommerce struct {
	*kernel.ServiceContainer

	Config      *kernel.Config
	AccessToken *auth.AccessToken

	Store *store.Client

	Logger *logger.Logger
}

type UserConfig struct {
	AppID  string
	Secret string

	StableTokenMode   bool
	ForceRefresh      bool
	RefreshToken      string
	ComponentAppID    string
	ComponentAppToken string
	Token             string
	AESKey            string

	ResponseType string
	Log          Log
	OAuth        OAuth
	Cache        kernel.CacheInterface

	Http Http

	HttpDebug bool
	Debug     bool
	NotifyURL string
	Sandbox   bool
}

type Http struct {
	Timeout   float64
	BaseURI   string
	ProxyURI  string
	Transport http.RoundTripper
}

type Log struct {
	Driver contract.LoggerInterface
	Level  string
	File   string
	Error  string
	ENV    string
	Stdout bool
}

type OAuth struct {
	Callback string
	Scopes   []string
}

func NewECommerce(config *UserConfig, extraInfos ...*kernel.ExtraInfo) (*ECommerce, error) {
	var err error

	userConfig, err := MapUserConfig(config)
	if err != nil {
		return nil, err
	}

	var extraInfo, _ = kernel.NewExtraInfo()
	if len(extraInfos) > 0 {
		extraInfo = extraInfos[0]
	}

	// init an app container
	container, err := kernel.NewServiceContainer(userConfig, extraInfo)
	if err != nil {
		return nil, err
	}
	container.GetConfig()

	// init app
	app := &ECommerce{
		ServiceContainer: container,
	}

	//-------------- global app config --------------
	// global app config
	app.Config = providers.RegisterConfigProvider(app)

	app.Logger, err = logger.NewLogger(app.Config.Get("log.driver", nil), &object.HashMap{
		"level":      app.Config.GetString("log.level", "info"),
		"env":        app.Config.GetString("log.env", "develop"),
		"outputPath": app.Config.GetString("log.file", "./wechat/info.log"),
		"errorPath":  app.Config.GetString("log.error", "./wechat/error.log"),
		"stdout":     app.Config.GetBool("log.stdout", false),
	})
	if err != nil {
		return nil, err
	}

	//-------------- register auth,AccessToken --------------
	app.AccessToken, err = auth.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	return app, err
}

func (app *ECommerce) GetContainer() *kernel.ServiceContainer {
	return app.ServiceContainer
}

func (app *ECommerce) GetAccessToken() *kernel.AccessToken {
	return app.AccessToken.AccessToken
}

func (app *ECommerce) GetConfig() *kernel.Config {
	return app.Config
}

func (app *ECommerce) GetComponent(name string) interface{} {

	switch name {
	case "AccessToken":
		return app.AccessToken
	case "Config":
		return app.Config

	case "Store":
		return app.Store

	case "Logger":
		return app.Logger

	default:
		return nil
	}

}

func MapUserConfig(userConfig *UserConfig) (*object.HashMap, error) {

	baseURI := "https://api.weixin.qq.com/"
	if userConfig.Http.BaseURI != "" {
		baseURI = userConfig.Http.BaseURI
	}

	config := &object.HashMap{

		"app_id": userConfig.AppID,
		"secret": userConfig.Secret,

		"token":               userConfig.Token,
		"aes_key":             userConfig.AESKey,
		"component_app_id":    userConfig.ComponentAppID,
		"component_app_token": userConfig.ComponentAppToken,
		"stable_token_mode":   userConfig.StableTokenMode,
		"refresh_token":       userConfig.RefreshToken,

		"response_type": userConfig.ResponseType,
		"http": &object.HashMap{
			"timeout":   userConfig.Http.Timeout,
			"base_uri":  baseURI,
			"proxy_uri": userConfig.Http.ProxyURI,
			"transport": userConfig.Http.Transport,
		},
		"log": &object.HashMap{
			"driver": userConfig.Log.Driver,
			"level":  userConfig.Log.Level,
			"file":   userConfig.Log.File,
			"error":  userConfig.Log.Error,
			"env":    userConfig.Log.ENV,
			"stdout": userConfig.Log.Stdout,
		},
		"cache": userConfig.Cache,

		"http_debug": userConfig.HttpDebug,
		"debug":      userConfig.Debug,
	}

	return config, nil

}
