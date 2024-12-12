package tiktok

import (
	"github.com/easycb/easycb-go/tiktok"
)

func NewClientSample(appKey, appSecret, baseUrl string) {
	client, err := tiktok.NewClient(appKey, appSecret, baseUrl)
	if err != nil {
		return
	}

	client.SetAccessToken("")
	client.SetShopCipher("")
}

func NewClientWithConfigurerSample(appKey, appSecret, baseUrl string) {
	// with proxy url or with timeout...
	client, err := tiktok.NewClient(appKey, appSecret, baseUrl, tiktok.WithProxyUrl("https://127.0.0.1:8888"), tiktok.WithTimeout(60))
	if err != nil {
		return
	}

	client.SetAccessToken("")
}
