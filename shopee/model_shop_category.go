package shopee

type AddShopCategoryRsp struct {
	BaseRsp
	Response struct {
		ShopCategoryId int `json:"shop_category_id"`
	} `json:"response"`
}

type GetShopCategoryListRsp struct {
	BaseRsp
	Response struct {
		ShopCategorys []struct {
			ShopCategoryId int    `json:"shop_category_id"`
			Status         int    `json:"status"`
			Name           string `json:"name"`
			SortWeight     int    `json:"sort_weight"`
			CreatedBy      string `json:"created_by"`
		} `json:"shop_categorys"`
		More       bool `json:"more"`
		TotalCount int  `json:"total_count"`
	} `json:"response"`
}

type DeleteShopCategoryRsp struct {
	BaseRsp
	Response struct {
		ShopCategoryId int `json:"shop_category_id"`
	} `json:"response"`
}

type UpdateShopCategoryRsp struct {
	BaseRsp
	Response struct {
		Status         string `json:"status"`
		ShopCategoryId int    `json:"shop_category_id"`
		SortWeight     int    `json:"sort_weight"`
		Name           string `json:"name"`
	} `json:"response"`
}

type ShopCategoryAddItemListRsp struct {
	BaseRsp
	Response struct {
		ShopCategoryId int `json:"shop_category_id"`
		CurrentCount   int `json:"current_count"`
	} `json:"response"`
}

type ShopCategoryGetItemListRsp struct {
	BaseRsp
	Response struct {
		ItemList   []int64 `json:"item_list"`
		More       bool    `json:"more"`
		TotalCount int     `json:"total_count"`
	} `json:"response"`
}

type ShopCategoryDeleteItemListRsp struct {
	BaseRsp
	Response struct {
		ShopCategoryId    int `json:"shop_category_id"`
		CurrentCount      int `json:"current_count"`
		InvalidItemIdList []struct {
			ItemId      int64  `json:"item_id"`
			FailMessage string `json:"fail_message"`
			FailError   string `json:"fail_error"`
		} `json:"invalid_item_id_list"`
	} `json:"response"`
}
