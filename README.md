# Easy Cross Border (EasyCb-Go)

Lazada/Tiktok/Shopee Seller Open Platform SDK For Golang

[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

## Overview

* Support [Lazada Open Api](https://open.lazada.com/apps/doc/api) 
* Support [Tiktok Shope Partner Api](https://partner.tiktokshop.com/api/document) `V2`
* Support [Shopee Open Api](https://open.shopee.com/documents) `V2`
* Every feature comes with tests
* Developer Friendly

**EasyCb features are support platform:**
- Shopify
- Amazon
- AliExpress
- ...

## Getting Started

### Prerequisites

EasyCb requires [Go](https://go.dev/) version [1.17](https://go.dev/doc/devel/release#go1.17.0) or above.

```sh
import "github.com/easycb/easycb-go"
```

Alternatively, use `go get`:

```sh
go get -u github.com/easycb/easycb-go
```

### Running EasyCb


A basic example Lazada open api:
```go
package main

import (
	"fmt"
	"time"
	"github.com/easycb/easycb-go"
	"github.com/easycb/easycb-go/lazada"
)

func main() {
	appKey := ""
	appSecret := ""
	baseUrl := "https://api.lazada.com.my/rest"
	accessToken := ""
	client, err := lazada.NewClient(appKey, appSecret, baseUrl)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	client.SetAccessToken(accessToken)

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
```


A basic example Shopee open api:
```go
package main

import (
	"fmt"
	"time"
	"github.com/easycb/easycb-go"
	"github.com/easycb/easycb-go/shopee"
)

func main() {
	partnerId := int64(123456)
	partnerKey := ""
	baseUrl := "https://partner.shopeemobile.com"
	shopId := int64(56789)
	accessToken := ""
	client, err := shopee.NewClient(partnerId, partnerKey, baseUrl)
	if err != nil {
		return
	}
	client.SetShopId(shopId).SetAccessToken(accessToken)
	
	
	query := easycb.AnyMap{
		"offset":      0,
		"page_size":   50,
		"item_status": []string{"NORMAL", "UNLIST"},
	}

	res, err := client.GetProductList(query)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Message)
}
```


A basic example Tiktok shop partner api:

```go
package main

import (
	"fmt"
	"time"
	"github.com/easycb/easycb-go"
	"github.com/easycb/easycb-go/tiktok"
)

func main() {
	appKey := ""
	appSecret := ""
	baseUrl := "https://open-api.tiktokglobalshop.com"
	accessToken := ""
	shopCipher := ""
	client, err := tiktok.NewClient(appKey, appSecret, baseUrl)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	client.SetAccessToken(accessToken).SetShopCipher(shopCipher)

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
	
	fmt.Println(res.Code)
}
```


## Contributors

[Thank you](https://github.com/easycb/easycb-go/graphs/contributors) for contributing to the EasyCb SDK!

## License

Released under the [MIT License](https://github.com/easycb/easycb-go/blob/master/LICENSE)
