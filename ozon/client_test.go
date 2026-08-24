package ozon

import (
	"os"
	"testing"
)

var (
	client   *Client
	err      error
	clientId = ""
	apiKey   = ""
	baserUrl = "https://api-seller.ozon.ru"
)

func TestMain(m *testing.M) {
	client, err = NewClient(clientId, apiKey, baserUrl, WithProxyUrl("http://127.0.0.1:8888"))
	if err != nil {
		panic(err)
		return
	}

	os.Exit(m.Run())
}
