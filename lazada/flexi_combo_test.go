package lazada

import (
	"fmt"
	"github.com/easycb/easycb-go"
	"testing"
)

func TestGetFlexiComboDetails(t *testing.T) {
	res, err := client.GetFlexiComboDetails(easycb.AnyMap{"id": 900000032478060})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}

func TestListFlexiComboProducts(t *testing.T) {
	res, err := client.ListFlexiComboProducts(easycb.AnyMap{"id": 900000021005160, "page_size": 10, "cur_page": 1})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}
