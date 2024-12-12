package lazada

import (
	"fmt"
	"github.com/easycb/easycb-go"
	"os"
	"testing"
)

func TestUploadImage(t *testing.T) {
	f, err := os.Open("../examples/demo.png")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer f.Close()

	body := easycb.AnyMap{
		"image": f,
	}

	res, err := client.UploadImage(body, "png")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
