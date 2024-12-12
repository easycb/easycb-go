package shopee

import (
	"github.com/easycb/easycb-go/shopee"
)

func NewClientForShopSample(partnerId int64, partnerKey, baseUrl string) {
	client, err := shopee.NewClient(partnerId, partnerKey, baseUrl)
	if err != nil {
		return
	}

	// set shop id and shop access token
	shopId := int64(0)
	client.SetShopId(shopId)
	client.SetAccessToken("")
}

func NewClientForMerchantSample(partnerId int64, partnerKey, baseUrl string) {
	client, err := shopee.NewClient(partnerId, partnerKey, baseUrl)
	if err != nil {
		return
	}

	// also you can set merchant id and merchant access token
	merchantId := int64(0)
	client.SetMerchantId(merchantId)
	client.SetAccessToken("")
}

func NewClientWithConfigurerSample(partnerId int64, partnerKey, baseUrl string) {
	// with proxy url or with timeout...
	client, err := shopee.NewClient(partnerId, partnerKey, baseUrl, shopee.WithProxyUrl("https://127.0.0.1:8888"), shopee.WithTimeout(60))
	if err != nil {
		return
	}

	// set shop id and shop access token
	shopId := int64(0)
	client.SetShopId(shopId)
	client.SetAccessToken("")
}
