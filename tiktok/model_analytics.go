package tiktok

type GetShopPerformanceRsp struct {
	BaseRsp
	Data struct {
		LatestAvailableDate string `json:"latest_available_date"`
		Performance         struct {
			ComparisonIntervals []struct {
				AvgOrderValue struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"avg_order_value"`
				AvgProductPageVisitorBreakdowns []struct {
					Amount int    `json:"amount"`
					Type   string `json:"type"`
				} `json:"avg_product_page_visitor_breakdowns"`
				AvgProductPageVisitors int `json:"avg_product_page_visitors"`
				BuyerBreakdowns        []struct {
					Amount int    `json:"amount"`
					Type   string `json:"type"`
				} `json:"buyer_breakdowns"`
				Buyers                  int    `json:"buyers"`
				CancellationsAndReturns int    `json:"cancellations_and_returns"`
				EndDate                 string `json:"end_date"`
				Gmv                     struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"gmv"`
				GmvBreakdowns []struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
					Type     string `json:"type"`
				} `json:"gmv_breakdowns"`
				Orders                      int `json:"orders"`
				ProductImpressionBreakdowns []struct {
					Amount int    `json:"amount"`
					Type   string `json:"type"`
				} `json:"product_impression_breakdowns"`
				ProductImpressions        int `json:"product_impressions"`
				ProductPageViewBreakdowns []struct {
					Amount int    `json:"amount"`
					Type   string `json:"type"`
				} `json:"product_page_view_breakdowns"`
				ProductPageViews int `json:"product_page_views"`
				Refunds          struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"refunds"`
				SkuOrders int    `json:"sku_orders"`
				StartDate string `json:"start_date"`
				UnitsSold int    `json:"units_sold"`
			} `json:"comparison_intervals"`
			Intervals []struct {
				AvgOrderValue struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"avg_order_value"`
				AvgProductPageVisitorBreakdowns []struct {
					Amount int    `json:"amount"`
					Type   string `json:"type"`
				} `json:"avg_product_page_visitor_breakdowns"`
				AvgProductPageVisitors int `json:"avg_product_page_visitors"`
				BuyerBreakdowns        []struct {
					Amount int    `json:"amount"`
					Type   string `json:"type"`
				} `json:"buyer_breakdowns"`
				Buyers                  int    `json:"buyers"`
				CancellationsAndReturns int    `json:"cancellations_and_returns"`
				EndDate                 string `json:"end_date"`
				Gmv                     struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"gmv"`
				GmvBreakdowns []struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
					Type     string `json:"type"`
				} `json:"gmv_breakdowns"`
				Orders                      int `json:"orders"`
				ProductImpressionBreakdowns []struct {
					Amount int    `json:"amount"`
					Type   string `json:"type"`
				} `json:"product_impression_breakdowns"`
				ProductImpressions        int `json:"product_impressions"`
				ProductPageViewBreakdowns []struct {
					Amount int    `json:"amount"`
					Type   string `json:"type"`
				} `json:"product_page_view_breakdowns"`
				ProductPageViews int `json:"product_page_views"`
				Refunds          struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"refunds"`
				SkuOrders int    `json:"sku_orders"`
				StartDate string `json:"start_date"`
				UnitsSold int    `json:"units_sold"`
			} `json:"intervals"`
		} `json:"performance"`
	} `json:"data"`
}

type GetShopProductPerformanceRsp struct {
	BaseRsp
	Data struct {
		LatestAvailableDate string `json:"latest_available_date"`
		Performance         struct {
			ComparisonIntervals []struct {
				AvgPageVisitorBreakdowns []struct {
					Amount int    `json:"amount"`
					Type   string `json:"type"`
				} `json:"avg_page_visitor_breakdowns"`
				AvgPageVisitors            int    `json:"avg_page_visitors"`
				ClickThroughRate           string `json:"click_through_rate"`
				ClickThroughRateBreakdowns []struct {
					Amount string `json:"amount"`
					Type   string `json:"type"`
				} `json:"click_through_rate_breakdowns"`
				EndDate string `json:"end_date"`
				Gmv     struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"gmv"`
				GmvBreakdowns []struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
					Type     string `json:"type"`
				} `json:"gmv_breakdowns"`
				ImpressionBreakdowns []struct {
					Amount int    `json:"amount"`
					Type   string `json:"type"`
				} `json:"impression_breakdowns"`
				Impressions        int `json:"impressions"`
				Orders             int `json:"orders"`
				PageViewBreakdowns []struct {
					Amount int    `json:"amount"`
					Type   string `json:"type"`
				} `json:"page_view_breakdowns"`
				PageViews          int    `json:"page_views"`
				StartDate          string `json:"start_date"`
				UnitSoldBreakdowns []struct {
					Amount int    `json:"amount"`
					Type   string `json:"type"`
				} `json:"unit_sold_breakdowns"`
				UnitsSold int `json:"units_sold"`
			} `json:"comparison_intervals"`
			Intervals []struct {
				AvgPageVisitorBreakdowns []struct {
					Amount int    `json:"amount"`
					Type   string `json:"type"`
				} `json:"avg_page_visitor_breakdowns"`
				AvgPageVisitors            int    `json:"avg_page_visitors"`
				ClickThroughRate           string `json:"click_through_rate"`
				ClickThroughRateBreakdowns []struct {
					Amount string `json:"amount"`
					Type   string `json:"type"`
				} `json:"click_through_rate_breakdowns"`
				EndDate string `json:"end_date"`
				Gmv     struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"gmv"`
				GmvBreakdowns []struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
					Type     string `json:"type"`
				} `json:"gmv_breakdowns"`
				ImpressionBreakdowns []struct {
					Amount int    `json:"amount"`
					Type   string `json:"type"`
				} `json:"impression_breakdowns"`
				Impressions        int `json:"impressions"`
				Orders             int `json:"orders"`
				PageViewBreakdowns []struct {
					Amount int    `json:"amount"`
					Type   string `json:"type"`
				} `json:"page_view_breakdowns"`
				PageViews          int    `json:"page_views"`
				StartDate          string `json:"start_date"`
				UnitSoldBreakdowns []struct {
					Amount int    `json:"amount"`
					Type   string `json:"type"`
				} `json:"unit_sold_breakdowns"`
				UnitsSold int `json:"units_sold"`
			} `json:"intervals"`
		} `json:"performance"`
	} `json:"data"`
}

type GetShopProductPerformanceListRsp struct {
	BaseRsp
	Data struct {
		LatestAvailableDate string `json:"latest_available_date"`
		NextPageToken       string `json:"next_page_token"`
		Products            []struct {
			ClickThroughRate string `json:"click_through_rate"`
			Gmv              struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"gmv"`
			Id        string `json:"id"`
			Orders    int    `json:"orders"`
			UnitsSold int    `json:"units_sold"`
		} `json:"products"`
		TotalCount int `json:"total_count"`
	} `json:"data"`
}

type GetShopSKUPerformanceRsp struct {
	BaseRsp
	Data struct {
		LatestAvailableDate string `json:"latest_available_date"`
		Performance         struct {
			ComparisonIntervals []struct {
				EndDate string `json:"end_date"`
				Gmv     struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"gmv"`
				GmvBreakdown []struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
					Type     string `json:"type"`
				} `json:"gmv_breakdown"`
				SkuOrders          int    `json:"sku_orders"`
				StartDate          string `json:"start_date"`
				UnitsSold          int    `json:"units_sold"`
				UnitsSoldBreakdown []struct {
					Amount int    `json:"amount"`
					Type   string `json:"type"`
				} `json:"units_sold_breakdown"`
			} `json:"comparison_intervals"`
			Intervals []struct {
				EndDate string `json:"end_date"`
				Gmv     struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"gmv"`
				GmvBreakdown []struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
					Type     string `json:"type"`
				} `json:"gmv_breakdown"`
				SkuOrders          int    `json:"sku_orders"`
				StartDate          string `json:"start_date"`
				UnitsSold          int    `json:"units_sold"`
				UnitsSoldBreakdown []struct {
					Amount int    `json:"amount"`
					Type   string `json:"type"`
				} `json:"units_sold_breakdown"`
			} `json:"intervals"`
			ProductId int `json:"product_id"`
		} `json:"performance"`
	} `json:"data"`
}

type GetShopSKUPerformanceListRsp struct {
	BaseRsp
	Data struct {
		LatestAvailableDate string `json:"latest_available_date"`
		NextPageToken       string `json:"next_page_token"`
		Skus                []struct {
			Gmv struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"gmv"`
			Id        int64 `json:"id"`
			ProductId int64 `json:"product_id"`
			SkuOrders int64 `json:"sku_orders"`
			UnitsSold int   `json:"units_sold"`
		} `json:"skus"`
		TotalCount int `json:"total_count"`
	} `json:"data"`
}

type GetShopVideoPerformanceListRsp struct {
	BaseRsp
	Data struct {
		LatestAvailableDate string `json:"latest_available_date"`
		NextPageToken       string `json:"next_page_token"`
		TotalCount          int    `json:"total_count"`
		Videos              []struct {
			ClickThroughRate string `json:"click_through_rate"`
			Gmv              struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"gmv"`
			Id       string `json:"id"`
			Products []struct {
				Id   string `json:"id"`
				Name string `json:"name"`
			} `json:"products"`
			SkuOrders int    `json:"sku_orders"`
			Title     string `json:"title"`
			UnitsSold int    `json:"units_sold"`
			Username  string `json:"username"`
			Views     int    `json:"views"`
		} `json:"videos"`
	} `json:"data"`
}

type GetShopVideoPerformanceOverviewRsp struct {
	BaseRsp
	Data struct {
		LatestAvailableDate string `json:"latest_available_date"`
		Performance         struct {
			ComparisonIntervals []struct {
				ClickThroughRate string `json:"click_through_rate"`
				EndDate          string `json:"end_date"`
				Gmv              struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"gmv"`
				SkuOrders int    `json:"sku_orders"`
				StartDate string `json:"start_date"`
				UnitsSold int    `json:"units_sold"`
			} `json:"comparison_intervals"`
			Intervals []struct {
				ClickThroughRate string `json:"click_through_rate"`
				EndDate          string `json:"end_date"`
				Gmv              struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"gmv"`
				SkuOrders int    `json:"sku_orders"`
				StartDate string `json:"start_date"`
				UnitsSold int    `json:"units_sold"`
			} `json:"intervals"`
		} `json:"performance"`
	} `json:"data"`
}

type GetShopVideoPerformanceDetailsRsp struct {
	BaseRsp
	Data struct {
		EngagementData struct {
			TotalComments int `json:"total_comments"`
			TotalLikes    int `json:"total_likes"`
			TotalShares   int `json:"total_shares"`
			TotalViews    int `json:"total_views"`
		} `json:"engagement_data"`
		LatestAvailableDate string `json:"latest_available_date"`
		Performance         struct {
			ComparisonIntervals []struct {
				ClickThroughRate string `json:"click_through_rate"`
				DailyAvgBuyers   string `json:"daily_avg_buyers"`
				EndDate          string `json:"end_date"`
				Gmv              struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"gmv"`
				StartDate string `json:"start_date"`
				Views     int    `json:"views"`
			} `json:"comparison_intervals"`
			Intervals []struct {
				ClickThroughRate string `json:"click_through_rate"`
				DailyAvgBuyers   string `json:"daily_avg_buyers"`
				EndDate          string `json:"end_date"`
				Gmv              struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"gmv"`
				StartDate string `json:"start_date"`
				Views     int    `json:"views"`
			} `json:"intervals"`
		} `json:"performance"`
	} `json:"data"`
}

type GetShopVideoProductPerformanceListRsp struct {
	BaseRsp
	Data struct {
		LatestAvailableDate string `json:"latest_available_date"`
		NextPageToken       string `json:"next_page_token"`
		Products            []struct {
			DailyAvgBuyers string `json:"daily_avg_buyers"`
			Gmv            struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"gmv"`
			Id        string `json:"id"`
			Name      string `json:"name"`
			UnitsSold int    `json:"units_sold"`
		} `json:"products"`
		TotalCount int `json:"total_count"`
	} `json:"data"`
}

type GetShopLivePerformanceListRsp struct {
	BaseRsp
	Data *struct {
		LiveStreamSessions []struct {
			Id               string `json:"id"`
			Title            string `json:"title"`
			Username         string `json:"username"`
			StartTime        string `json:"start_time"`
			EndTime          string `json:"end_time"`
			SalesPerformance struct {
				Gmv struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"gmv"`
				ProductsAdded         int `json:"products_added"`
				DifferentProductsSold int `json:"different_products_sold"`
				CreatedSkuOrders      int `json:"created_sku_orders"`
				SkuOrders             int `json:"sku_orders"`
				UnitSold              int `json:"unit_sold"`
				Customers             int `json:"customers"`
				AvgPrice              struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"avg_price"`
				ClickToOrderRate string `json:"click_to_order_rate"`
				HLiveGmv         struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"24h_live_gmv"`
			} `json:"sales_performance"`
			InteractionPerformance struct {
				Acu                int    `json:"acu"`
				Pcu                int    `json:"pcu"`
				Viewers            int    `json:"viewers"`
				Views              int    `json:"views"`
				AvgViewingDuration string `json:"avg_viewing_duration"`
				Comments           int    `json:"comments"`
				Shares             int    `json:"shares"`
				Likes              int    `json:"likes"`
				NewFollowers       int    `json:"new_followers"`
				ProductImpressions int    `json:"product_impressions"`
				ProductClicks      int    `json:"product_clicks"`
				ClickThroughRate   string `json:"click_through_rate"`
			} `json:"interaction_performance"`
		} `json:"live_stream_sessions"`
		LatestAvailableDate string `json:"latest_available_date"`
		NextPageToken       string `json:"next_page_token"`
		TotalCount          int    `json:"total_count"`
	} `json:"data,omitempty"`
}

type GetShopLivePerformanceOverviewRsp struct {
	BaseRsp
	Data *struct {
		Performance struct {
			Intervals []struct {
				StartDate string `json:"start_date"`
				EndDate   string `json:"end_date"`
				Gmv       struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"gmv"`
				SkuOrders        int    `json:"sku_orders"`
				Customers        int    `json:"customers"`
				UnitsSold        int    `json:"units_sold"`
				ClickToOrderRate string `json:"click_to_order_rate"`
				ClickThroughRate string `json:"click_through_rate"`
			} `json:"intervals"`
			ComparisonIntervals []struct {
				StartDate string `json:"start_date"`
				EndDate   string `json:"end_date"`
				Gmv       struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"gmv"`
				SkuOrders        int    `json:"sku_orders"`
				Customers        int    `json:"customers"`
				UnitsSold        int    `json:"units_sold"`
				ClickToOrderRate string `json:"click_to_order_rate"`
				ClickThroughRate string `json:"click_through_rate"`
			} `json:"comparison_intervals"`
		} `json:"performance"`
		LatestAvailableDate string `json:"latest_available_date"`
	} `json:"data,omitempty"`
}
