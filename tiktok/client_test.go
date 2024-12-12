package tiktok

import (
	"fmt"
	"os"
	"testing"
)

var (
	client      *Client
	err         error
	appKey      = ""
	appSecret   = ""
	baseUrl     = "https://open-api.tiktokglobalshop.com"
	accessToken = ""
	shopCipher  = ""
)

func TestMain(m *testing.M) {
	client, err = NewClient(appKey, appSecret, baseUrl, WithProxyUrl("http://127.0.0.1:8888"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	client.SetAccessToken(accessToken)
	client.SetShopCipher(shopCipher)
	// client.SetShopCipher(shopCipher)
	os.Exit(m.Run())
}
