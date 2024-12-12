package shopee

import (
	"fmt"
	"github.com/easycb/easycb-go"
	"testing"
)

func TestGetAddOnDealList(t *testing.T) {
	res, err := client.GetAddOnDealList(easycb.AnyMap{"promotion_status": "all", "page_size": 100})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Message)
}
