package lazada

import (
	"fmt"
	"github.com/easycb/easycb-go"
	"testing"
)

func TestGetSeller(t *testing.T) {
	res, err := client.GetSeller()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}

func TestSellerCenterMsgList(t *testing.T) {
	res, err := client.SellerCenterMsgList(easycb.AnyMap{"page": 1, "pageSize": 20, "language": "zh"})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
