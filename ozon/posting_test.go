package ozon

import (
	"fmt"
	"testing"
	"time"

	"github.com/easycb/easycb-go"
)

func TestPostingFbsList(T *testing.T) {
	body := easycb.AnyMap{
		"with": easycb.AnyMap{
			"analytics_data": true,
			"barcodes":       true,
			"financial_data": true,
			"legal_info":     true,
			"translit":       true,
		},
		"sort_dir": "DESC",
		"limit":    100,
		"filter": easycb.AnyMap{
			"since": time.Now().UTC().Add(-7 * 24 * time.Hour).Format("2006-01-02T15:04:05Z"),
			"to":    time.Now().UTC().Format("2006-01-02T15:04:05Z"),
		},
	}

	res, err := client.PostingFbsList(body)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
