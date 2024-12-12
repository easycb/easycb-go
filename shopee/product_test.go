package shopee

import (
	"fmt"
	"github.com/easycb/easycb-go"
	"testing"
)

func TestGetCategory(t *testing.T) {
	query := easycb.AnyMap{
		"language": "zh-hans",
	}

	res, err := client.GetCategory(query)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Message)
}

func TestGetAttributeTree(t *testing.T) {
	query := easycb.AnyMap{
		"category_id_list": []int64{101944, 101945},
		"language":         "zh-hans",
	}

	res, err := client.GetAttributeTree(query)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}

func TestGetProductList(t *testing.T) {
	query := easycb.AnyMap{
		"offset":      0,
		"page_size":   50,
		"item_status": []string{"NORMAL", "UNLIST"},
	}

	res, err := client.GetItemList(query)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Message)
}

func TestGetProductBaseInfo(t *testing.T) {
	query := easycb.AnyMap{
		"item_id_list": []int64{28204531813, 25233088606},
	}

	res, err := client.GetItemBaseInfo(query)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Message)
}

func TestGetProductItemExtraInfo(t *testing.T) {
	query := easycb.AnyMap{
		"item_id_list": []int64{28204531813, 25233088606},
	}

	res, err := client.GetItemExtraInfo(query)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Message)
}
