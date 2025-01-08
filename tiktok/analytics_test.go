package tiktok

import (
	"fmt"
	"github.com/easycb/easycb-go"
	"testing"
)

func TestGetShopPerformance(t *testing.T) {
	query := easycb.AnyMap{
		"start_date_ge": "2024-12-01",
		"end_date_lt":   "2024-12-31",
	}
	res, err := client.GetShopPerformance(query)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Data)
}
