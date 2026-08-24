package ozon

import (
	"fmt"
	"testing"
	"time"

	"github.com/easycb/easycb-go"
)

func TestReturnsList(t *testing.T) {
	body := easycb.AnyMap{
		"limit": 100,
		"filter": map[string]interface{}{
			"logistic_return_date": map[string]interface{}{
				"time_from": time.Now().Add(-180 * 24 * time.Hour).UTC().Format("2006-01-02T15:04:05Z"),
				"time_to":   time.Now().Add(-1 * 24 * time.Hour).UTC().Format("2006-01-02T15:04:05Z"),
			},
		},
	}

	res, err := client.ReturnsList(body)
	if err != nil {

		panic(err)
	}

	fmt.Println(res)
}
