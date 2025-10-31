package auth

// #reference: https://open.work.weixin.qq.com/api/doc/90000/90135/91039

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel"
)

type AccessToken struct {
	*kernel.AccessToken
}

func NewAccessToken(app kernel.ApplicationInterface) (*AccessToken, error) {
	kernelToken, err := kernel.NewAccessToken(app)
	token := &AccessToken{
		kernelToken,
	}

	// Override fields and functions
	cfg := app.GetConfig()
	useStableToken := cfg.GetBool("stable_token_mode", false)
	forceRefresh := cfg.GetBool("force_refresh", false)
	baseUrl := cfg.GetString("http.base_uri", "https://api.weixin.qq.com/")
	if useStableToken {
		token.EndpointToGetToken = baseUrl + "cgi-bin/stable_token"
		token.StableTokenMode = true
		token.ForceRefresh = forceRefresh
	} else {
		token.EndpointToGetToken = baseUrl + "cgi-bin/token"
	}
	token.OverrideGetCredentials()

	return token, err
}

// Override GetCredentials
func (accessToken *AccessToken) OverrideGetCredentials() {
	config := (accessToken.App).GetConfig()

	accessToken.GetCredentials = func() *object.StringMap {
		return &object.StringMap{
			"grant_type": "client_credential",
			"appid":      config.GetString("app_id", ""),
			"secret":     config.GetString("secret", ""),
			"neededText": "",
		}
	}
}
