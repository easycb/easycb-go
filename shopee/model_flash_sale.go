package shopee

type GetTimeSlotIdRsp struct {
	BaseRsp
	Response []struct {
		EndTime    int64 `json:"end_time"`
		StartTime  int64 `json:"start_time"`
		TimeslotId int64 `json:"timeslot_id"`
	} `json:"response"`
}

type CreateShopFlashSaleRsp struct {
	BaseRsp
	Response struct {
		FlashSaleId int64 `json:"flash_sale_id"`
		Status      int   `json:"status"`
		TimeslotId  int64 `json:"timeslot_id"`
	} `json:"response"`
}

type GetItemCriteriaRsp struct {
	BaseRsp
	Response struct {
		Criteria []struct {
			CriteriaId       int64   `json:"criteria_id"`
			MinProductRating float64 `json:"min_product_rating"`
			MinLikes         int     `json:"min_likes"`
			MustNotPreOrder  bool    `json:"must_not_pre_order"`
			MinOrderTotal    int     `json:"min_order_total"`
			MaxDaysToShip    int     `json:"max_days_to_ship"`
			MinRepetitionDay int     `json:"min_repetition_day"`
			MinPromoStock    int     `json:"min_promo_stock"`
			MaxPromoStock    int     `json:"max_promo_stock"`
			MinDiscount      int64   `json:"min_discount"`
			MaxDiscount      int64   `json:"max_discount"`
			MinDiscountPrice int64   `json:"min_discount_price"`
			MaxDiscountPrice int64   `json:"max_discount_price"`
			NeedLowestPrice  bool    `json:"need_lowest_price"`
		} `json:"criteria"`
		PariIds []struct {
			CriteriaId   int `json:"criteria_id"`
			CategoryList []struct {
				CategoryId int    `json:"category_id"`
				Name       string `json:"name"`
				ParentId   int64  `json:"parent_id"`
			} `json:"category_list"`
		} `json:"pair_ids"`
		OverlapBlockCategoryIds []int64 `json:"overlap_block_category_ids"`
	} `json:"response"`
}

type AddShopFlashSaleItemsRsp struct {
	BaseRsp
	Response struct {
		FailedItems []struct {
			ErrCode int    `json:"err_code"`
			ErrMsg  string `json:"err_msg"`
			ItemId  int64  `json:"item_id"`
			ModelId int64  `json:"model_id"`
		} `json:"failed_items"`
	} `json:"response"`
}

type GetShopFlashSaleListRsp struct {
	BaseRsp
	Response struct {
		FlashSaleList []struct {
			ClickCount       int   `json:"click_count"`
			EnabledItemCount int   `json:"enabled_item_count"`
			EndTime          int64 `json:"end_time"`
			FlashSaleId      int64 `json:"flash_sale_id"`
			ItemCount        int   `json:"item_count"`
			RemindmeCount    int   `json:"remindme_count"`
			StartTime        int64 `json:"start_time"`
			Status           int   `json:"status"`
			TimeslotId       int64 `json:"timeslot_id"`
			Type             int   `json:"type"`
		} `json:"flash_sale_list"`
		TotalCount int `json:"total_count"`
	} `json:"response"`
}

type GetShopFlashSaleRsp struct {
	BaseRsp
	Response struct {
		EnabledItemCount int   `json:"enabled_item_count"`
		EndTime          int64 `json:"end_time"`
		FlashSaleId      int64 `json:"flash_sale_id"`
		ItemCount        int   `json:"item_count"`
		StartTime        int64 `json:"start_time"`
		Status           int   `json:"status"`
		TimeslotId       int64 `json:"timeslot_id"`
		Type             int   `json:"type"`
	}
}

type GetShopFlashSaleItemsRsp struct {
	BaseRsp
	Response struct {
		ItemInfo []struct {
			Image    string `json:"image"`
			ItemId   int64  `json:"item_id"`
			ItemName string `json:"item_name"`
			Status   int    `json:"status"`
		} `json:"item_info"`
		Models []struct {
			CampaignStock         int     `json:"campaign_stock"`
			InputPromotionPrice   float64 `json:"input_promotion_price"`
			ItemId                int64   `json:"item_id"`
			ModelId               int64   `json:"model_id"`
			ModelName             string  `json:"model_name"`
			OriginalPrice         float64 `json:"original_price"`
			PromotionPriceWithTax float64 `json:"promotion_price_with_tax"`
			PurchaseLimit         int     `json:"purchase_limit"`
			RejectReason          string  `json:"reject_reason"`
			Status                int     `json:"status"`
			Stock                 int     `json:"stock"`
		} `json:"models"`
		TotalCount int `json:"total_count"`
	} `json:"response"`
}

type UpdateShopFlashSaleRsp struct {
	BaseRsp
	Response struct {
		FlashSaleId int64 `json:"flash_sale_id"`
		Status      int   `json:"status"`
		TimeslotId  int64 `json:"timeslot_id"`
	} `json:"response"`
}

type UpdateShopFlashSaleItemsRsp struct {
	BaseRsp
	Response struct {
		FailedItems []struct {
			ErrCode               int    `json:"err_code"`
			ErrMsg                string `json:"err_msg"`
			ItemId                int64  `json:"item_id"`
			ModelId               int    `json:"model_id"`
			UnqualifiedConditions []struct {
				UnqualifiedCode int    `json:"unqualified_code"`
				UnqualifiedMsg  string `json:"unqualified_msg"`
			} `json:"unqualified_conditions"`
		} `json:"failed_items"`
	} `json:"response"`
}

type DeleteShopFlashSaleRsp struct {
	BaseRsp
	Response struct {
		FlashSaleId int64 `json:"flash_sale_id"`
		Status      int   `json:"status"`
		TimeslotId  int64 `json:"timeslot_id"`
	} `json:"response"`
}

type DeleteShopFlashSaleItemsRsp struct {
	BaseRsp
}
