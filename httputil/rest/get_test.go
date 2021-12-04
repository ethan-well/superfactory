package rest

import (
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

func TestGet(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}
	url := `https://hacker-news.firebaseio.com/v0/item/2921983.json?print=pretty`
	
	var dest _response
	err := Get(url, &dest)
	if err != nil {
		t.Fatal(err)
	}
	
	t.Logf("resp: %s", logger.ToJson(dest))
}
