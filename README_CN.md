[English](./README.md) | 简体中文

# Easy Cross Border (EasyCb-Go)

Lazada/TikTok/Shopee/Shein卖家开放平台Golang SDK

[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/easycb/easycb-go)

## 概述

* 支持 [Lazada 卖家开放平台](https://open.lazada.com/apps/doc/api)
* 支持 [Tiktok 合作伙伴](https://partner.tiktokshop.com/api/document) `V2`
* 支持 [Shopee 卖家开放平台](https://open.shopee.com/documents) `V2`
* 支持 [Shein 开放平台](https://open.sheincorp.com)
* 每个功能都附带测试
* 对开发者友好

**EasyCb 未来将支持以下平台:**
- Shopify
- Amazon
- AliExpress
- ...

## 开始使用

### 准备工作

EasyCb 需要 [Go](https://go.dev/) 版本>= [1.17](https://go.dev/doc/devel/release#go1.17.0) 及更高版本

```sh
import "github.com/easycb/easycb-go"
```

或者使用 `go get`:

```sh
go get -u github.com/easycb/easycb-go
```

### 运行 EasyCb


Lazada 基础使用示例:
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


Shopee 基础使用示例:
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


Tiktok 基础使用示例:
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


Shein 基础使用示例:
```go
package main

import (
	"fmt"
	"time"
	"github.com/easycb/easycb-go"
	"github.com/easycb/easycb-go/shein"
)

func main() {
	baseUrl := "https://openapi.sheincorp.com"
	appId := ""
	openKeyId := ""
	secretKey := ""
	client, err := shein.NewClient(appId, openKeyId, secretKey, baseUrl)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

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
```


## 贡献者

[感谢](https://github.com/easycb/easycb-go/graphs/contributors) 对EasyCb贡献!

## 许可

协议 [MIT License](https://github.com/easycb/easycb-go/blob/master/LICENSE)
