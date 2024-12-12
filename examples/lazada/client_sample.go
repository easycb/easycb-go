package lazada

import "github.com/easycb/easycb-go/lazada"

func NewClientSample(appKey, appSecret, baseUrl string) {
	client, err := lazada.NewClient(appKey, appSecret, baseUrl)
	if err != nil {
		return
	}

	client.SetAccessToken("")
}

func NewClientDefaultSample(appKey, appSecret, region string) {
	client, err := lazada.NewClientDefault(appKey, appSecret)
	if err != nil {
		return
	}

	// set region to server url
	client.SetRegion(region)

	// set access token
	client.SetAccessToken("")
}

func NewClientWithConfigurerSample(appKey, appSecret, baseUrl string) {
	// with proxy url or with timeout...
	client, err := lazada.NewClient(appKey, appSecret, baseUrl, lazada.WithProxyUrl("https://127.0.0.1:8888"), lazada.WithTimeout(60))
	if err != nil {
		return
	}

	client.SetAccessToken("")
}
