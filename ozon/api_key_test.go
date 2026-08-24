package ozon

import (
	"fmt"
	"testing"
)

func TestGetRoles(t *testing.T) {
	res, err := client.GetRoles(nil)
	if err != nil {

		panic(err)
	}

	fmt.Println(res)
}
