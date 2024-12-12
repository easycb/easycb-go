package lazada

type CreateGlobalProductRsp struct {
	Code string `json:"code"`
	Data struct {
		SkuList []struct {
			SellerSku string `json:"seller_sku"`
		} `json:"sku_list"`
	} `json:"data"`
	RequestId string `json:"request_id"`
}

type GetGlobalProductExtensionRsp struct {
	ErrorMsg string `json:"error_msg"`
	Code     string `json:"code"`
	Data     []struct {
		GlobalItemId int64 `json:"global_item_id"`
		ItemId       int64 `json:"item_id"`
		Products     []struct {
			Market     string `json:"market"`
			SemiStatus int    `json:"semi_status"`
			Abs        string `json:"abs"`
			Skus       []struct {
				SpecialPrice struct {
					Amount   float64 `json:"amount"`
					Currency string  `json:"currency"`
				} `json:"special_price"`
				Price struct {
					Amount   float64 `json:"amount"`
					Currency string  `json:"currency"`
				} `json:"price"`
				SellerSku    string `json:"seller_sku"`
				NoPostageFee struct {
					Amount   float64 `json:"amount"`
					Currency string  `json:"currency"`
				} `json:"no_postage_fee"`
				SkuId int64 `json:"sku_id"`
			} `json:"skus"`
			ItemId int64 `json:"item_id"`
		} `json:"products"`
	} `json:"data"`
	Success   bool   `json:"success"`
	ErrorCode string `json:"error_code"`
	RequestId string `json:"request_id"`
}

type GetGlobalProductStatusRsp struct {
	ErrorMsg  string `json:"error_msg"`
	Code      string `json:"code"`
	Data      string `json:"data"`
	Success   bool   `json:"success"`
	ErrorCode string `json:"error_code"`
	RequestId string `json:"request_id"`
	// todo 待补充逻辑
}

type GetRecommendPriceRsp struct {
	ErrorMsg string `json:"error_msg"`
	Code     string `json:"code"`
	Data     struct {
		GlobalItemId int64 `json:"global_item_id"`
		Skus         []struct {
			SellerSku    string `json:"seller_sku"`
			SkuId        int64  `json:"sku_id"`
			CountryPrice []struct {
				Market         string `json:"market"`
				NoPostagePrice string `json:"no_postage_price"`
				Currency       string `json:"currency"`
			} `json:"country_price"`
		} `json:"skus"`
		ItemId int64 `json:"item_id"`
	} `json:"data"`
	Success   bool   `json:"success"`
	ErrorCode string `json:"error_code"`
	RequestId string `json:"request_id"`
}

type GetUnfilledAttributeRsp struct {
	ErrorMsg string `json:"error_msg"`
	Code     string `json:"code"`
	Data     struct {
		TotalProducts int `json:"total_products"`
		Products      []struct {
			ItemId          int64  `json:"item_id"`
			PrimaryCategory int    `json:"primary_category"`
			SellerSku       string `json:"seller_sku"`
			Attributes      []struct {
				Advanced struct {
					IsKeyProp int `json:"is_key_prop"`
				} `json:"advanced"`
				InputType     string   `json:"input_type"`
				Options       []string `json:"options"`
				Name          string   `json:"name"`
				IsMandatory   int      `json:"is_mandatory"`
				AttributeType string   `json:"attribute_type"`
				Label         string   `json:"label"`
			} `json:"attributes"`
		} `json:"products"`
	} `json:"data"`
	Success     bool   `json:"success"`
	ErrorDetail string `json:"error_detail"`
	ErrorCode   string `json:"error_code"`
	RequestId   string `json:"request_id"`
	Errors      string `json:"errors"`
}

type GetUpgradableGlobalPlusProductListRsp struct {
	Code string `json:"code"`
	Data struct {
		Type          string `json:"type"`
		TotalProducts int    `json:"total_products"`
		CurrentPage   int    `json:"current_page"`
		PageSize      int    `json:"page_size"`
		CurrentIndex  int    `json:"current_index"`
		Products      []struct {
			GlobalItemId int64 `json:"global_item_id"`
			Skus         []struct {
				PackageWidth  string `json:"package_width"`
				PackageHeight string `json:"package_height"`
				ItemId        int64  `json:"item_id"`
				PackageLength string `json:"package_length"`
				SellerSku     string `json:"seller_sku"`
				PackageWeight string `json:"package_weight"`
				SkuId         int64  `json:"sku_id"`
				CountryInfo   []struct {
					Market       string `json:"market"`
					Quantity     int    `json:"quantity"`
					Abs          string `json:"abs"`
					SpecialPrice string `json:"special_price"`
					ItemId       int64  `json:"item_id"`
					Price        string `json:"price"`
					Currency     string `json:"currency"`
					SkuId        int64  `json:"sku_id"`
				} `json:"country_info"`
			} `json:"skus"`
			ItemId string `json:"item_id"`
		} `json:"products"`
	} `json:"data"`
	Success   bool   `json:"success"`
	RequestId string `json:"request_id"`
}

type SemiProductUpdateRsp struct {
	ErrorMsg string `json:"error_msg"`
	Code     string `json:"code"`
	Data     struct {
		ProductId int64 `json:"product_id"`
	} `json:"data"`
	Success   bool   `json:"success"`
	ErrorCode string `json:"error_code"`
	RequestId string `json:"request_id"`
}

type SemiProductUpgradeRsp struct {
	ErrorMsg string `json:"error_msg"`
	Code     string `json:"code"`
	Data     struct {
		ProductId int64 `json:"product_id"`
	} `json:"data"`
	Success   bool   `json:"success"`
	ErrorCode string `json:"error_code"`
	RequestId string `json:"request_id"`
}

type UpdateGlobalProductAttributeRsp struct {
	ErrorMsg    string `json:"error_msg"`
	Code        string `json:"code"`
	Success     bool   `json:"success"`
	ErrorDetail string `json:"error_detail"`
	ErrorCode   string `json:"error_code"`
	RequestId   string `json:"request_id"`
	Errors      string `json:"errors"`
}

type DeleteMerchantProductRsp struct {
	ErrorMsg string `json:"error_msg"`
	Code     string `json:"code"`
	Data     struct {
		DeleteICProductResult         bool `json:"deleteICProductResult"`
		DeleteIcProductFailResultList []struct {
			Market       string `json:"market"`
			ProductId    int64  `json:"productId"`
			UpdateMsg    string `json:"updateMsg"`
			UpdateResult bool   `json:"updateResult"`
		} `json:"deleteIcProductFailResultList"`
		DeleteGspProductResult bool `json:"deleteGspProductResult"`
	} `json:"data"`
	Success   bool   `json:"success"`
	ErrorCode string `json:"error_code"`
	RequestId string `json:"request_id"`
}

type UpdateProductStatusRsp struct {
	ErrorMsg string `json:"error_msg"`
	Code     string `json:"code"`
	Data     struct {
		UpdateIcProductResult         bool `json:"update_ic_product_result"`
		UpdateGspProductResult        bool `json:"update_gsp_product_result"`
		UpdateIcProductFailResultList []struct {
			Market       string `json:"market"`
			ProductId    int64  `json:"product_id"`
			UpdateResult bool   `json:"update_result"`
			UpdateMsg    string `json:"update_msg"`
		} `json:"update_ic_product_fail_result_list"`
	} `json:"data"`
	Success   bool   `json:"success"`
	ErrorCode string `json:"error_code"`
	RequestId string `json:"request_id"`
}
