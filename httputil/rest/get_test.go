package rest

import (
	"github.com/ItsWewin/superfactory/httputil"
	"github.com/ItsWewin/superfactory/logger"
	"testing"
)

type _response struct {
	By string `json:"by"`
	ID int64 `json:"id"`
	Kids []int64 `json:"kids"`
	Parent int64 `json:"parent"`
	Text string `json:"text"`
	Time int64 `json:"time"`
	MsgType string `json:"type"`
}

type _userProfileResponse struct {
	ID int64 `json:"id"`
	Login string `json:"login"`
}

func TestGet(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}
	//url := `https://hacker-news.firebaseio.com/v0/item/2921983.json?print=pretty`
	url := `https://api.github.com/user`
	accessToken := "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"
	
	var dest _userProfileResponse
	opts := func(opts *Options) {
		opts.Header = map[string]string{
			httputil.HeaderContent: httputil.JsonHeaderContent,
			httputil.HeaderAuthorization: "token " + accessToken,
		}
	}
	
	err := Get(url, &dest, opts)
	if err != nil {
		t.Fatal(err)
	}
	
	t.Logf("resp: %s", logger.ToJson(dest))
}