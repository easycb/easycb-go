package lazada

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
	baseUrl     = "https://api.lazada.com/rest"
	accessToken = ""
)

func TestMain(m *testing.M) {
	client, err = NewClient(appKey, appSecret, baseUrl, WithProxyUrl("http://127.0.0.1:8888"), WithSslVerify(false))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	client.SetAccessToken(accessToken)

	os.Exit(m.Run())
}
