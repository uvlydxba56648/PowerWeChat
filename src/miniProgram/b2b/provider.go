package b2b

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(app, nil)
	config := app.GetConfig()
	appKey := config.GetString("app_key", "")
	sandboxAppKey := config.GetString("sandbox_app_key", "")

	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
		appKey,
		sandboxAppKey,
	}, nil

}
