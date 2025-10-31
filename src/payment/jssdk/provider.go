package jssdk

import "github.com/uvlydxba56648/PowerWeChat/v3/src/kernel"

func RegisterProvider(app kernel.ApplicationInterface) (*Client, error) {

	return NewClient(app)

}
