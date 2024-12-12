package shopee

type GetTotalBalanceRsp struct {
	BaseRsp
	Response struct {
		DataTimestamp int64   `json:"data_timestamp"`
		TotalBalance  float64 `json:"total_balance"`
	} `json:"response"`
}

type GetShopToggleInfoRsp struct {
	BaseRsp
	Response struct {
		DataTimestamp int  `json:"data_timestamp"`
		AutoTopUp     bool `json:"auto_top_up"`
		CampaignSurge bool `json:"campaign_surge"`
	} `json:"response"`
}

type GetRecommendedKeywordListRsp struct {
	BaseRsp
	Response struct {
		ItemId            int64  `json:"item_id"`
		InputKeyword      string `json:"input_keyword"`
		SuggestedKeywords []struct {
			Keyword      string  `json:"keyword"`
			QualityScore int     `json:"quality_score"`
			SearchVolume int     `json:"search_volume"`
			SuggestedBid float64 `json:"suggested_bid"`
		} `json:"suggested_keywords"`
	} `json:"response"`
}

type GetRecommendedItemListRsp struct {
	BaseRsp
	Response []struct {
		ItemId            int64    `json:"item_id"`
		ItemStatusList    []string `json:"item_status_list"`
		SkuTagList        []string `json:"sku_tag_list"`
		OngoingAdTypeList []string `json:"ongoing_ad_type_list"`
	} `json:"response"`
}

type GetAllCpcAdsHourlyPerformanceRsp struct {
	BaseRsp
	Response []struct {
		Hour              int     `json:"hour"`
		Date              string  `json:"date"`
		Impression        int     `json:"impression"`
		Clicks            int     `json:"clicks"`
		Ctr               float64 `json:"ctr"`
		DirectOrder       int     `json:"direct_order"`
		BroadOrder        int     `json:"broad_order"`
		DirectConversions float64 `json:"direct_conversions"`
		BroadConversions  float64 `json:"broad_conversions"`
		DirectItemSold    int     `json:"direct_item_sold"`
		BroadItemSold     int     `json:"broad_item_sold"`
		DirectGmv         float64 `json:"direct_gmv"`
		BroadGmv          float64 `json:"broad_gmv"`
		Expense           float64 `json:"expense"`
		CostPerConversion float64 `json:"cost_per_conversion"`
		DirectRoas        float64 `json:"direct_roas"`
		BroadRoas         float64 `json:"broad_roas"`
	} `json:"response"`
}

type GetAllCpcAdsDailyPerformanceRsp struct {
	BaseRsp
	Response []struct {
		Date              string  `json:"date"`
		Impression        int     `json:"impression"`
		Clicks            int     `json:"clicks"`
		Ctr               float64 `json:"ctr"`
		DirectOrder       int     `json:"direct_order"`
		BroadOrder        int     `json:"broad_order"`
		DirectConversions float64 `json:"direct_conversions"`
		BroadConversions  float64 `json:"broad_conversions"`
		DirectItemSold    int     `json:"direct_item_sold"`
		BroadItemSold     int     `json:"broad_item_sold"`
		DirectGmv         float64 `json:"direct_gmv"`
		BroadGmv          float64 `json:"broad_gmv"`
		Expense           float64 `json:"expense"`
		CostPerConversion float64 `json:"cost_per_conversion"`
		DirectRoas        float64 `json:"direct_roas"`
		BroadRoas         float64 `json:"broad_roas"`
	} `json:"response"`
}

type GetProductCampaignDailyPerformanceRsp struct {
	BaseRsp
	Response []struct {
		ShopId       int64  `json:"shop_id"`
		Region       string `json:"region"`
		CampaignList []struct {
			CampaignId        int    `json:"campaign_id"`
			AdType            string `json:"ad_type"`
			CampaignPlacement string `json:"campaign_placement"`
			AdName            string `json:"ad_name"`
			MetricsList       []struct {
				Date              string  `json:"date"`
				Impression        int     `json:"impression"`
				Clicks            int     `json:"clicks"`
				Ctr               float64 `json:"ctr"`
				Expense           float64 `json:"expense"`
				BroadGmv          float64 `json:"broad_gmv"`
				BroadOrder        int     `json:"broad_order"`
				BroadOrderAmount  int     `json:"broad_order_amount"`
				BroadRoi          float64 `json:"broad_roi"`
				BroadCir          float64 `json:"broad_cir"`
				Cr                float64 `json:"cr"`
				Cpc               float64 `json:"cpc"`
				DirectOrder       int     `json:"direct_order"`
				DirectOrderAmount int     `json:"direct_order_amount"`
				DirectGmv         float64 `json:"direct_gmv"`
				DirectRoi         float64 `json:"direct_roi"`
				DirectCir         float64 `json:"direct_cir"`
				DirectCr          float64 `json:"direct_cr"`
				Cpdc              float64 `json:"cpdc"`
			} `json:"metrics_list"`
		} `json:"campaign_list"`
	} `json:"response"`
}

type GetProductCampaignHourlyPerformanceRsp struct {
	BaseRsp
	Response []struct {
		ShopId       int64  `json:"shop_id"`
		Region       string `json:"region"`
		CampaignList []struct {
			CampaignId        int    `json:"campaign_id"`
			AdType            string `json:"ad_type"`
			CampaignPlacement string `json:"campaign_placement"`
			AdName            string `json:"ad_name"`
			MetricsList       []struct {
				Hour              int     `json:"hour"`
				Date              string  `json:"date"`
				Impression        int     `json:"impression"`
				Clicks            int     `json:"clicks"`
				Ctr               float64 `json:"ctr"`
				Expense           float64 `json:"expense"`
				BroadGmv          float64 `json:"broad_gmv"`
				BroadOrder        int     `json:"broad_order"`
				BroadOrderAmount  int     `json:"broad_order_amount"`
				BroadRoi          float64 `json:"broad_roi"`
				BroadCir          float64 `json:"broad_cir"`
				Cr                float64 `json:"cr"`
				Cpc               float64 `json:"cpc"`
				DirectOrder       int     `json:"direct_order"`
				DirectOrderAmount int     `json:"direct_order_amount"`
				DirectGmv         float64 `json:"direct_gmv"`
				DirectRoi         float64 `json:"direct_roi"`
				DirectCir         float64 `json:"direct_cir"`
				DirectCr          float64 `json:"direct_cr"`
				Cpdc              float64 `json:"cpdc"`
			} `json:"metrics_list"`
		} `json:"campaign_list"`
	} `json:"response"`
}

type GetProductLevelCampaignIdListRsp struct {
	BaseRsp
	Response struct {
		ShopId       int64  `json:"shop_id"`
		Region       string `json:"region"`
		HasNextPage  bool   `json:"has_next_page"`
		CampaignList []struct {
			AdType     string `json:"ad_type"`
			CampaignId int    `json:"campaign_id"`
		} `json:"campaign_list"`
	} `json:"response"`
}

type GetProductLevelCampaignSettRsp struct {
	BaseRsp
	Response struct {
		ShopId       int64  `json:"shop_id"`
		Region       string `json:"region"`
		CampaignList []struct {
			CampaignId int `json:"campaign_id"`
			CommonInfo struct {
				AdType            string  `json:"ad_type"`
				AdName            string  `json:"ad_name"`
				CampaignStatus    string  `json:"campaign_status"`
				BiddingMethod     string  `json:"bidding_method"`
				CampaignPlacement string  `json:"campaign_placement"`
				CampaignBudget    float64 `json:"campaign_budget"`
				CampaignDuration  struct {
					StartTime int64 `json:"start_time"`
					EndTime   int64 `json:"end_time"`
				} `json:"campaign_duration"`
			} `json:"common_info"`
			ManualBiddingInfo struct {
				EnhancedCpc      bool `json:"enhanced_cpc"`
				SelectedKeywords []struct {
					Keyword          string  `json:"keyword"`
					Status           string  `json:"status"`
					MatchType        string  `json:"match_type"`
					BidPricePerClick float64 `json:"bid_price_per_click"`
				} `json:"selected_keywords"`
				DiscoveryAdsLocations []struct {
					Location string  `json:"location"`
					Status   string  `json:"status"`
					BidPrice float64 `json:"bid_price"`
				} `json:"discovery_ads_locations"`
			} `json:"manual_bidding_info"`
			AutoBiddingInfo struct {
				RoasTarget float64 `json:"roas_target"`
			} `json:"auto_bidding_info"`
			AutoProductAdsInfo []struct {
				ProductName string `json:"product_name"`
				Status      string `json:"status"`
			} `json:"auto_product_ads_info"`
		} `json:"campaign_list"`
	} `json:"response"`
}
