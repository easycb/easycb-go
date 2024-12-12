package lazada

import (
	"fmt"
	"github.com/easycb/easycb-go"
	"testing"
	"time"
)

func TestGetOrders(t *testing.T) {
	query := easycb.AnyMap{
		"sort_direction": "DESC",
		"offset":         0,
		"limit":          100,
		"sort_by":        "created_at",
		"created_before": time.Now().Format(time.RFC3339),
		"created_after":  time.Now().Add(-15 * 24 * time.Hour).Format(time.RFC3339),
	}

	res, err := client.GetOrders(query)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Code)
}
