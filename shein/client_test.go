package shein

import (
	"os"
	"testing"
)

var (
	client       *Client
	err          error
	appId        = ""
	appSecretKey = ""
	openKeyId    = ""
	secretKey    = ""
	baserUrl     = "https://openapi.sheincorp.com"
)

func TestMain(m *testing.M) {
	client, err = NewClient(appId, openKeyId, secretKey, baserUrl, WithProxyUrl("http://127.0.0.1:8888"))
	if err != nil {
		panic(err)
		return
	}

	os.Exit(m.Run())
}
