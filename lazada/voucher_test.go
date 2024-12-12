package lazada

import (
	"fmt"
	"github.com/easycb/easycb-go"
	"testing"
)

func TestSellerVoucherDetailQuery(t *testing.T) {
	res, err := client.SellerVoucherDetailQuery(easycb.AnyMap{"id": 900000032474075, "voucher_type": "COLLECTIBLE_VOUCHER"})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}
