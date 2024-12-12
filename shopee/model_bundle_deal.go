package shopee

type AddBundleDealRsp struct {
	BaseRsp
	Response struct {
		BundleDealId int64 `json:"bundle_deal_id"`
	} `json:"response"`
}

type AddBundleDealItemRsp struct {
	BaseRsp
	Response struct {
		FailedList []struct {
			ItemId      int    `json:"item_id"`
			FailError   string `json:"fail_error"`
			FailMessage string `json:"fail_message"`
		} `json:"failed_list"`
		SuccessList []int64 `json:"success_list"`
	} `json:"response"`
}

type GetBundleDealListRsp struct {
	BaseRsp
	Response struct {
		BundleDealList []struct {
			BundleDealId   int64 `json:"bundle_deal_id"`
			BundleDealRule struct {
				AdditionalTiers []struct {
					DiscountPercentage int     `json:"discount_percentage"`
					DiscountValue      float64 `json:"discount_value"`
					FixPrice           float64 `json:"fix_price"`
					MinAmount          int     `json:"min_amount"`
				} `json:"additional_tiers"`
				DiscountPercentage int     `json:"discount_percentage"`
				DiscountValue      float64 `json:"discount_value"`
				FixPrice           float64 `json:"fix_price"`
				MinAmount          float64 `json:"min_amount"`
				RuleType           int     `json:"rule_type"`
			} `json:"bundle_deal_rule"`
			EndTime       int64  `json:"end_time"`
			Name          string `json:"name"`
			PurchaseLimit int    `json:"purchase_limit"`
			StartTime     int64  `json:"start_time"`
		} `json:"bundle_deal_list"`
		More bool `json:"more"`
	} `json:"response"`
}

type GetBundleDealRsp struct {
	BaseRsp
	Response struct {
		BundleDealId   int64  `json:"bundle_deal_id"`
		Name           string `json:"name"`
		StartTime      int64  `json:"start_time"`
		EndTime        int64  `json:"end_time"`
		BundleDealRule struct {
			RuleType           int     `json:"rule_type"`
			DiscountValue      int     `json:"discount_value"`
			FixPrice           float64 `json:"fix_price"`
			DiscountPercentage int     `json:"discount_percentage"`
			MinAmount          float64 `json:"min_amount"`
			AdditionalTiers    struct {
				MinAmount          float64 `json:"min_amount"`
				FixPrice           float64 `json:"fix_price"`
				DiscountValue      float64 `json:"discount_value"`
				DiscountPercentage int     `json:"discount_percentage"`
			} `json:"additional_tiers"`
		} `json:"bundle_deal_rule"`
		PurchaseLimit int `json:"purchase_limit"`
	} `json:"response"`
}

type GetBundleDealItemRsp struct {
	BaseRsp
	Response struct {
		ItemList []struct {
			ItemId int64 `json:"item_id"`
			Status int   `json:"status"`
		} `json:"item_list"`
		TotalCount int `json:"total_count"`
	} `json:"response"`
}

type UpdateBundleDealRsp struct {
	BaseRsp
	Response struct {
		BundleDealId   int64  `json:"bundle_deal_id"`
		Name           string `json:"name"`
		StartTime      int    `json:"start_time"`
		EndTime        int    `json:"end_time"`
		BundleDealRule struct {
			RuleType           int     `json:"rule_type"`
			DiscountValue      float64 `json:"discount_value"`
			FixPrice           float64 `json:"fix_price"`
			DiscountPercentage int     `json:"discount_percentage"`
			MinAmount          float64 `json:"min_amount"`
		} `json:"bundle_deal_rule"`
		PurchaseLimit int `json:"purchase_limit"`
	} `json:"response"`
}

type UpdateBundleDealItemRsp struct {
	BaseRsp
	Response struct {
		FailedList []struct {
			ItemId      int    `json:"item_id"`
			FailError   string `json:"fail_error"`
			FailMessage string `json:"fail_message"`
		} `json:"failed_list"`
		SuccessList []int `json:"success_list"`
	} `json:"response"`
}

type EndBundleDealRsp struct {
	BaseRsp
	Response struct {
		BundleDealId int64 `json:"bundle_deal_id"`
	} `json:"response"`
}

type DeleteBundleDealRsp struct {
	BaseRsp
	Response struct {
		BundleDealId int64 `json:"bundle_deal_id"`
	} `json:"response"`
}

type DeleteBundleDealItemRsp struct {
	BaseRsp
	Response struct {
		FailedList []struct {
			ItemId      int    `json:"item_id"`
			FailError   string `json:"fail_error"`
			FailMessage string `json:"fail_message"`
		} `json:"failed_list"`
		SuccessList []int `json:"success_list"`
	} `json:"response"`
}
