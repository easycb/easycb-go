package shopee

import (
	"os"
	"testing"
)

var (
	client          *Client
	err             error
	partnerId       = int64(0)
	partnerKey      = ""
	baseUrl         = "https://partner.shopeemobile.com"
	shopAccessToken = ""
	shopId          = int64(0)
)

func TestMain(m *testing.M) {
	client, err = NewClient(partnerId, partnerKey, baseUrl, WithProxyUrl("http://127.0.0.1:8888"))
	if err != nil {
		return
	}
	client.SetShopId(shopId)
	client.SetAccessToken(shopAccessToken)

	os.Exit(m.Run())
}
