package tiktok

import (
	"fmt"
	"github.com/easycb/easycb-go"
	"testing"
	"time"
)

func TestGetOrderList(t *testing.T) {
	query := easycb.AnyMap{
		"page_size": 20,
	}

	body := easycb.AnyMap{
		"create_time_ge": time.Now().Add(-15 * 24 * time.Hour).Unix(),
		"create_time_lt": time.Now().Unix(),
	}

	res, err := client.GetOrderList(query, body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Message)
}
