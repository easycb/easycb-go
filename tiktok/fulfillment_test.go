package tiktok

import (
	"fmt"
	"github.com/easycb/easycb-go"
	"os"
	"testing"
	"time"
)

func TestFulfillmentUploadDeliveryFile(t *testing.T) {
	f, err := os.Open("../examples/demo.pdf")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer f.Close()

	body := easycb.AnyMap{
		"name": fmt.Sprintf("%d.pdf", time.Now().Unix()),
		"data": f,
	}
	res, err := client.FulfillmentUploadDeliveryFile(nil, body, "pdf")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Data)
}

func TestFulfillmentUploadDeliveryImage(t *testing.T) {
	f, err := os.Open("../examples/demo.png")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer f.Close()

	body := easycb.AnyMap{
		"data": f,
	}
	res, err := client.FulfillmentUploadDeliveryImage(nil, body, "png")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Data)
}
