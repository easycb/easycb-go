package shopee

import (
	"fmt"
	"github.com/easycb/easycb-go"
	"io/ioutil"
	"testing"
)

func TestDownloadShippingDocument(t *testing.T) {
	body := easycb.AnyMap{
		"shipping_document_type": "THERMAL_AIR_WAYBILL",
		"order_list": []easycb.AnyMap{
			{
				"order_sn": "241203U5MW1YPX",
			},
		},
	}
	res, err := client.DownloadShippingDocument(body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = ioutil.WriteFile("output.pdf", res.RawBytes, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}
}
