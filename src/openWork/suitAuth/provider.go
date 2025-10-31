package suit

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) (*SuiteTicket, *AccessToken, error) {
	ticket, err := NewSuiteTicket(app)
	if err != nil {
		return nil, nil, err
	}
	accessToken, err := NewAccessToken(app)
	if err != nil {
		return nil, nil, err
	}

	return ticket, accessToken, nil
}
