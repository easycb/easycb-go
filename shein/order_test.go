package shein

import (
	"fmt"
	"github.com/easycb/easycb-go"
	"testing"
)

func TestGetOrderList(t *testing.T) {
	body := easycb.AnyMap{
		"queryType": 1,
		"startTime": "2025-05-06 00:00:00",
		"endTime":   "2025-05-08 00:00:00",
		"page":      1,
		"pageSize":  30,
	}
	res, err := client.GetOrderList(body)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

func TestGetOrderDetail(t *testing.T) {
	body := easycb.AnyMap{
		"orderNoList": []string{"GSU12515200MPC5"},
	}
	res, err := client.GetOrderDetail(body)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
