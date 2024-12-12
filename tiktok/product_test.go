package tiktok

import (
	"fmt"
	"github.com/easycb/easycb-go"
	"os"
	"testing"
)

func TestUploadProductImage(t *testing.T) {
	f, err := os.Open("../examples/demo.png")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer f.Close()

	body := easycb.AnyMap{
		"use_case": "MAIN_IMAGE",
		"data":     f,
	}
	res, err := client.UploadProductImage(body, "png")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Data)
}
