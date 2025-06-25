package shein

import (
	"fmt"
	"github.com/easycb/easycb-go"
	"os"
	"testing"
)

func TestUploadPic(t *testing.T) {
	f, err := os.Open("../examples/demo.png")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	body := easycb.AnyMap{
		"file":       f,
		"image_type": 7,
	}

	headers := map[string]string{
		"language": "en",
	}

	res, err := client.UploadPic(body, headers)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

func TestQuerySiteList(t *testing.T) {
	res, err := client.QuerySiteList(nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
