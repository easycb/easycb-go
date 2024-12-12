package shopee

type AddAddOnDealRsp struct {
	BaseRsp
	Response struct {
		AddOnDealId int64 `json:"add_on_deal_id"`
	} `json:"response"`
}

type AddAddOnDealMainItemRsp struct {
	BaseRsp
	Response struct {
		MainItemList []struct {
			ItemId int64 `json:"item_id"`
			Status int   `json:"status"`
		} `json:"main_item_list"`
		AddOnDealId int `json:"add_on_deal_id"`
	} `json:"response"`
}

type AddAddOnDealSubItemRsp struct {
	BaseRsp
	Response struct {
		AddOnDealId int64 `json:"add_on_deal_id"`
		SubItemList []struct {
			ItemId      int64  `json:"item_id"`
			ModelId     int    `json:"model_id"`
			FailMessage string `json:"fail_message"`
			FailError   string `json:"fail_error"`
		} `json:"sub_item_list"`
	} `json:"response"`
}

type DeleteAddOnDealRsp struct {
	BaseRsp
	Response struct {
		AddOnDealId int64 `json:"add_on_deal_id"`
	} `json:"response"`
}

type DeleteAddOnDealMainItemRsp struct {
	BaseRsp
	Response struct {
		MainItemList []struct {
			ItemId int `json:"item_id"`
		} `json:"main_item_list"`
		AddOnDealId int64 `json:"add_on_deal_id"`
	} `json:"response"`
}

type DeleteAddOnDealSubItemRsp struct {
	BaseRsp
	Response struct {
		AddOnDealId int64 `json:"add_on_deal_id"`
		SubItemList []struct {
			ItemId      int64  `json:"item_id"`
			ModelId     int    `json:"model_id"`
			FailMessage string `json:"fail_message"`
			FailError   string `json:"fail_error"`
		} `json:"sub_item_list"`
	} `json:"response"`
}

type GetAddOnDealListRsp struct {
	BaseRsp
	Response struct {
		AddOnDealList []struct {
			AddOnDealId            int64       `json:"add_on_deal_id"`
			AddOnDealName          string      `json:"add_on_deal_name"`
			PerGiftNum             int         `json:"per_gift_num"`
			PromotionPurchaseLimit int         `json:"promotion_purchase_limit"`
			PromotionType          int         `json:"promotion_type"`
			PurchaseMinSpend       float64     `json:"purchase_min_spend"`
			Source                 int         `json:"source"`
			EndTime                int64       `json:"end_time"`
			StartTime              int64       `json:"start_time"`
			SubItemPriority        interface{} `json:"sub_item_priority"`
		} `json:"add_on_deal_list"`
		More bool `json:"more"`
	} `json:"response"`
}

type GetAddOnDealRsp struct {
	BaseRsp
	Response struct {
		StartTime              int64         `json:"start_time"`
		PurchaseMinSpend       float64       `json:"purchase_min_spend"`
		Source                 int           `json:"source"`
		AddOnDealId            int64         `json:"add_on_deal_id"`
		PromotionPurchaseLimit int           `json:"promotion_purchase_limit"`
		EndTime                int64         `json:"end_time"`
		AddOnDealName          string        `json:"add_on_deal_name"`
		PerGiftNum             int           `json:"per_gift_num"`
		PromotionType          int           `json:"promotion_type"`
		SubItemPriority        []interface{} `json:"sub_item_priority"`
	} `json:"response"`
}

type GetAddOnDealMainItemRsp struct {
	BaseRsp
	Response struct {
		MainItemList []struct {
			Status int `json:"status"`
			ItemId int `json:"item_id"`
		} `json:"main_item_list"`
		AddOnDealId int64 `json:"add_on_deal_id"`
	} `json:"response"`
}

type GetAddOnDealSubItemRsp struct {
	BaseRsp
	Response struct {
		AddOnDealId int64 `json:"add_on_deal_id"`
		SubItemList []struct {
			Status       int `json:"status"`
			ItemId       int `json:"item_id"`
			SubItemLimit int `json:"sub_item_limit"`
			Price        struct {
				PromoInputPrice float64 `json:"promo_input_price"`
				PromoPrice      float64 `json:"promo_price"`
			} `json:"price"`
		} `json:"sub_item_list"`
	} `json:"response"`
}

type UpdateAddOnDealRsp struct {
	BaseRsp
	Response struct {
		StartTime              int64   `json:"start_time"`
		PurchaseMinSpend       float64 `json:"purchase_min_spend"`
		Source                 int     `json:"source"`
		AddOnDealId            int64   `json:"add_on_deal_id"`
		PromotionPurchaseLimit int     `json:"promotion_purchase_limit"`
		EndTime                int64   `json:"end_time"`
		AddOnDealName          string  `json:"add_on_deal_name"`
		PerGiftNum             int     `json:"per_gift_num"`
		PromotionType          int     `json:"promotion_type"`
		SubItemPriority        []int   `json:"sub_item_priority"`
	} `json:"response"`
}

type UpdateAddOnDealMainItemRsp struct {
	BaseRsp
	Response struct {
		MainItemList []struct {
			ItemId int `json:"item_id"`
			Status int `json:"status"`
		} `json:"main_item_list"`
		AddOnDealId int64 `json:"add_on_deal_id"`
	} `json:"response"`
}

type UpdateAddOnDealSubItemRsp struct {
	BaseRsp
	Response struct {
		AddOnDealId int64 `json:"add_on_deal_id"`
		SubItemList []struct {
			Status            int     `json:"status"`
			ItemId            int     `json:"item_id"`
			SubItemLimit      int     `json:"sub_item_limit"`
			SubItemInputPrice float64 `json:"sub_item_input_price"`
		} `json:"sub_item_list"`
	} `json:"response"`
}

type EndAddOnDealRsp struct {
	BaseRsp
	Response struct {
		AddOnDealId int64 `json:"add_on_deal_id"`
	} `json:"response"`
}
