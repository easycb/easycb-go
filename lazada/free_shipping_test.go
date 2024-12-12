package lazada

import (
	"fmt"
	"github.com/easycb/easycb-go"
	"testing"
)

func TestFreeShippingRegionsQuery(t *testing.T) {
	res, err := client.FreeShippingRegionsQuery()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}

func TestFreeShippingDeliveryOptionsQuery(t *testing.T) {
	res, err := client.FreeShippingDeliveryOptionsQuery()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}

func TestFreeShippingGet(t *testing.T) {
	res, err := client.FreeShippingGet(easycb.AnyMap{"id": 900000032412090})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
