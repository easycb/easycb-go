package lazada

import (
	"github.com/easycb/easycb-go/shein"
)

func NewClientSample(appId, openKeyId, secretKey, baseUrl string) {
	client, err := shein.NewClient(appId, openKeyId, secretKey, baseUrl)
	if err != nil {
		return
	}

	client.SetBaseurl("")
}

func NewClientDefaultSample(appId, appSecretKey string) {
	client, err := shein.NewClientDefault(appId, appSecretKey)
	if err != nil {
		return
	}

	client.SetBaseurl("")
}

func NewClientWithConfigurerSample(appId, openKeyId, secretKey, baseUrl string) {
	// with proxy url or with timeout...
	client, err := shein.NewClient(appId, openKeyId, secretKey, baseUrl, shein.WithProxyUrl("https://127.0.0.1:8888"), shein.WithTimeout(60))
	if err != nil {
		return
	}

	client.SetBaseurl("")
}
