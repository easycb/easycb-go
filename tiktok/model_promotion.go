package tiktok

type CreateActivityRsp struct {
	BaseRsp
	Data struct {
		ActivityId string `json:"activity_id"`
		CreateTime int64  `json:"create_time"`
		Status     string `json:"status"`
		UpdateTime int64  `json:"update_time"`
	} `json:"data"`
}

type UpdateActivityRsp struct {
	BaseRsp
	Data struct {
		ActivityId string `json:"activity_id"`
		Title      string `json:"title"`
		UpdateTime int64  `json:"update_time"`
	} `json:"data"`
}

type DeactivateActivityRsp struct {
	BaseRsp
	Data struct {
		ActivityId string `json:"activity_id"`
		Status     string `json:"status"`
		Title      string `json:"title"`
		UpdateTime int64  `json:"update_time"`
	} `json:"data"`
}

type GetActivityRsp struct {
	BaseRsp
	Data struct {
		ActivityId   string `json:"activity_id"`
		ActivityType string `json:"activity_type"`
		BeginTime    int64  `json:"begin_time"`
		CreateTime   int64  `json:"create_time"`
		EndTime      int64  `json:"end_time"`
		ProductLevel string `json:"product_level"`
		Products     []struct {
			ActivityPrice struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"activity_price"`
			Discount        string `json:"discount"`
			Id              string `json:"id"`
			QuantityLimit   int    `json:"quantity_limit"`
			QuantityPerUser int    `json:"quantity_per_user"`
			Skus            []struct {
				ActivityPrice struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"activity_price"`
				Discount        string `json:"discount"`
				Id              string `json:"id"`
				QuantityLimit   int    `json:"quantity_limit"`
				QuantityPerUser int    `json:"quantity_per_user"`
			} `json:"skus"`
		} `json:"products"`
		Status     string `json:"status"`
		Title      string `json:"title"`
		UpdateTime int    `json:"update_time"`
	} `json:"data"`
}

type SearchActivitiesRsp struct {
	BaseRsp
	Data struct {
		Activities []struct {
			ActivityType string `json:"activity_type"`
			BeginTime    int64  `json:"begin_time"`
			CreateTime   int64  `json:"create_time"`
			EndTime      int64  `json:"end_time"`
			Id           string `json:"id"`
			ProductLevel string `json:"product_level"`
			Status       string `json:"status"`
			Title        string `json:"title"`
			UpdateTime   int64  `json:"update_time"`
		} `json:"activities"`
		NextPageToken string `json:"next_page_token"`
		TotalCount    int    `json:"total_count"`
	} `json:"data"`
}

type UpdateActivityProductRsp struct {
	BaseRsp
	Data struct {
		ActivityId string `json:"activity_id"`
		Status     string `json:"status"`
		Title      string `json:"title"`
		TotalCount int    `json:"total_count"`
		UpdateTime int64  `json:"update_time"`
	} `json:"data"`
}

type RemoveActivityProductRsp struct {
	BaseRsp
	Data struct {
		ActivityId string `json:"activity_id"`
		Status     string `json:"status"`
		UpdateTime int64  `json:"update_time"`
	} `json:"data"`
}

type GetCouponRsp struct {
	BaseRsp
	Data struct {
		Coupon struct {
			ClaimDuration struct {
				EndTime   int64 `json:"end_time"`
				StartTime int64 `json:"start_time"`
			} `json:"claim_duration"`
			CreateTime     int64  `json:"create_time"`
			CreationSource string `json:"creation_source"`
			Discount       struct {
				MaxDiscount struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"max_discount"`
				Percentage      string `json:"percentage"`
				ReductionAmount struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"reduction_amount"`
				Type string `json:"type"`
			} `json:"discount"`
			DisplayChannels []string `json:"display_channels"`
			DisplayType     string   `json:"display_type"`
			Id              string   `json:"id"`
			LiveTasks       []struct {
				MinWatchTime string `json:"min_watch_time"`
				Type         string `json:"type"`
			} `json:"live_tasks"`
			ProductIds         []string `json:"product_ids"`
			ProductScope       string   `json:"product_scope"`
			PromoCode          string   `json:"promo_code"`
			RedemptionDuration struct {
				EndTime      int64  `json:"end_time"`
				RelativeTime int64  `json:"relative_time"`
				StartTime    int64  `json:"start_time"`
				Type         string `json:"type"`
			} `json:"redemption_duration"`
			SellerTnc          string `json:"seller_tnc"`
			Status             string `json:"status"`
			TargetBuyerSegment string `json:"target_buyer_segment"`
			Threshold          struct {
				MinSpend struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"min_spend"`
				Type string `json:"type"`
			} `json:"threshold"`
			Title       string `json:"title"`
			UpdateTime  int64  `json:"update_time"`
			UsageLimits struct {
				RedemptionLimit       int `json:"redemption_limit"`
				SingleBuyerClaimLimit int `json:"single_buyer_claim_limit"`
				TotalClaimLimit       int `json:"total_claim_limit"`
			} `json:"usage_limits"`
			UsageStats struct {
				ClaimedCount  int `json:"claimed_count"`
				RedeemedCount int `json:"redeemed_count"`
			} `json:"usage_stats"`
		} `json:"coupon"`
	} `json:"data"`
}

type SearchCouponsRsp struct {
	BaseRsp
	Data struct {
		Coupons []struct {
			ClaimDuration struct {
				EndTime   int64 `json:"end_time"`
				StartTime int64 `json:"start_time"`
			} `json:"claim_duration"`
			CreateTime     int64  `json:"create_time"`
			CreationSource string `json:"creation_source"`
			Discount       struct {
				MaxDiscount struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"max_discount"`
				Percentage      string `json:"percentage"`
				ReductionAmount struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"reduction_amount"`
				Type string `json:"type"`
			} `json:"discount"`
			DisplayType        string `json:"display_type"`
			Id                 string `json:"id"`
			ProductScope       string `json:"product_scope"`
			PromoCode          string `json:"promo_code"`
			RedemptionDuration struct {
				EndTime      int64  `json:"end_time"`
				RelativeTime int64  `json:"relative_time"`
				StartTime    int64  `json:"start_time"`
				Type         string `json:"type"`
			} `json:"redemption_duration"`
			Status             string `json:"status"`
			TargetBuyerSegment string `json:"target_buyer_segment"`
			Threshold          struct {
				MinSpend struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"min_spend"`
				Type string `json:"type"`
			} `json:"threshold"`
			Title       string `json:"title"`
			UpdateTime  int64  `json:"update_time"`
			UsageLimits struct {
				RedemptionLimit       int `json:"redemption_limit"`
				SingleBuyerClaimLimit int `json:"single_buyer_claim_limit"`
				TotalClaimLimit       int `json:"total_claim_limit"`
			} `json:"usage_limits"`
		} `json:"coupons"`
		NextPageToken string `json:"next_page_token"`
		TotalCount    int    `json:"total_count"`
	} `json:"data"`
}
