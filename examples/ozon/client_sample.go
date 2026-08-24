package ozon

import (
	"github.com/easycb/easycb-go/ozon"
)

func NewClientSample(clientId, apiKey, baseUrl string) {
	client, err := ozon.NewClient(clientId, apiKey, baseUrl)
	if err != nil {
		return
	}

	// do something
	client.GetRoles(nil)
}

func NewClientDefaultSample(appId, appSecretKey string) {
	client, err := ozon.NewClientDefault()
	if err != nil {
		return
	}

	// set base url
	client.SetBaseurl("https://api-seller.ozon.ru")
}

func NewClientWithConfigurerSample(clientId, apiKey, baseUrl string) {
	// with proxy url or with timeout...
	client, err := ozon.NewClient(clientId, apiKey, baseUrl, ozon.WithProxyUrl("https://127.0.0.1:8888"), ozon.WithTimeout(60))
	if err != nil {
		return
	}

	// do something
	client.GetRoles(nil)
}
