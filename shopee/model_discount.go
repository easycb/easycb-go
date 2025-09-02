package shopee

type AddDiscountRsp struct {
	BaseRsp
	Response struct {
		DiscountId int64 `json:"discount_id"`
	} `json:"response"`
}

type AddDiscountItemRsp struct {
	BaseRsp
	Response struct {
		DiscountId int64         `json:"discount_id"`
		Count      int           `json:"count"`
		ErrorList  []interface{} `json:"error_list"`
		Warning    string        `json:"warning"`
	} `json:"response"`
}

type DeleteDiscountRsp struct {
	BaseRsp
	Response struct {
		DiscountId int64 `json:"discount_id"`
		ModifyTime int64 `json:"modify_time"`
	} `json:"response"`
}

type DeleteDiscountItemRsp struct {
	BaseRsp
	Response struct {
		DiscountId int64         `json:"discount_id"`
		ErrorList  []interface{} `json:"error_list"`
	} `json:"response"`
}

type GetDiscountRsp struct {
	BaseRsp
	Response struct {
		Status   string `json:"status"`
		ItemList []struct {
			ItemPromotionPrice float64 `json:"item_promotion_price"`
			ItemName           string  `json:"item_name"`
			ModelList          []struct {
				ModelId             int64   `json:"model_id"`
				ModelPromotionPrice float64 `json:"model_promotion_price"`
				ModelOriginalPrice  float64 `json:"model_original_price"`
				ModelNormalStock    int     `json:"model_normal_stock"`
				ModelName           string  `json:"model_name"`
				ModelPromotionStock int     `json:"model_promotion_stock"`
			} `json:"model_list"`
			ItemPromotionStock int     `json:"item_promotion_stock"`
			NormalStock        int     `json:"normal_stock"`
			ItemId             int     `json:"item_id"`
			PurchaseLimit      int     `json:"purchase_limit"`
			ItemOriginalPrice  float64 `json:"item_original_price"`
		} `json:"item_list"`
		DiscountName string `json:"discount_name"`
		StartTime    int64  `json:"start_time"`
		DiscountId   int64  `json:"discount_id"`
		Source       int    `json:"source"`
		EndTime      int64  `json:"end_time"`
		More         bool   `json:"more"`
	} `json:"response"`
}

type GetDiscountListRsp struct {
	BaseRsp
	Response struct {
		DiscountList []struct {
			DiscountId   int64  `json:"discount_id"`
			DiscountName string `json:"discount_name"`
			EndTime      int64  `json:"end_time"`
			Source       int    `json:"source"`
			StartTime    int64  `json:"start_time"`
			Status       string `json:"status"`
		} `json:"discount_list"`
		More bool `json:"more"`
	} `json:"response"`
}

type UpdateDiscountRsp struct {
	BaseRsp
	Response struct {
		DiscountId int64 `json:"discount_id"`
		ModifyTime int64 `json:"modify_time"`
	} `json:"response"`
}

type UpdateDiscountItemRsp struct {
	BaseRsp
	Response struct {
		DiscountId int64 `json:"discount_id"`
		Count      int   `json:"count"`
		ErrorList  []struct {
			ItemId      int    `json:"item_id"`
			ModelId     int    `json:"model_id"`
			FailMessage string `json:"fail_message"`
			FailError   string `json:"fail_error"`
		} `json:"error_list"`
	} `json:"response"`
}

type EndDiscountRsp struct {
	BaseRsp
	Response struct {
		DiscountId int64 `json:"discount_id"`
		ModifyTime int64 `json:"modify_time"`
	} `json:"response"`
}
