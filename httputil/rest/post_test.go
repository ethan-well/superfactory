package rest

import (
	"github.com/ItsWewin/superfactory/logger"
	"testing"
)

type _accessTokenRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
	RedirectUrl  string `json:"redirect_url"`
}

type _accessTokenResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

func TestPost(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}
	url := `https://github.com/login/oauth/access_token`

	body := _accessTokenRequest{
		ClientID:     "",
		ClientSecret: "",
		Code:         "",
	}
	var response _accessTokenResponse
	_, err := Post(url, body, &response)
	if err != nil {
		t.Fatal("some error: ", err)
	}

	t.Logf("resp: %s", logger.ToJson(response))
}
