package ozon

type AnalyticsDataRsp struct {
	BaseRsp

	Result struct {
		Data []struct {
			Dimensions []struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"dimensions"`
			Metrics []float64 `json:"metrics"`
		} `json:"data"`
		Totals []float64 `json:"totals"`
	} `json:"result"`
	Timestamp string `json:"timestamp"`
}

type AnalyticsProductQueriesRsp struct {
	BaseRsp

	AnalyticsPeriod struct {
		DateFrom string `json:"date_from"`
		DateTo   string `json:"date_to"`
	} `json:"analytics_period"`
	Items     []AnalyticsProductQueriesItem `json:"items"`
	PageCount int64                         `json:"page_count"`
	Total     int64                         `json:"total"`
}

type AnalyticsProductQueriesItem struct {
	Category          string  `json:"category"`
	Currency          string  `json:"currency"`
	Gmv               float64 `json:"gmv"`
	Name              string  `json:"name"`
	OfferId           string  `json:"offer_id"`
	Position          float64 `json:"position"`
	Sku               int64   `json:"sku"`
	UniqueSearchUsers int64   `json:"unique_search_users"`
	UniqueViewUsers   int64   `json:"unique_view_users"`
	ViewConversion    float64 `json:"view_conversion"`
}

type AnalyticsProductQueriesDetailsRsp struct {
	BaseRsp

	AnalyticsPeriod struct {
		DateFrom string `json:"date_from"`
		DateTo   string `json:"date_to"`
	} `json:"analytics_period"`
	PageCount int64                                       `json:"page_count"`
	Queries   []AnalyticsProductQueriesDetailsQueriesItem `json:"queries"`
	Total     int64                                       `json:"total"`
}

type AnalyticsProductQueriesDetailsQueriesItem struct {
	Currency          string  `json:"currency"`
	Gmv               float64 `json:"gmv"`
	OrderCount        int64   `json:"order_count"`
	Position          float64 `json:"position"`
	Query             string  `json:"query"`
	QueryIndex        int64   `json:"query_index"`
	Sku               int64   `json:"sku"`
	UniqueSearchUsers int64   `json:"unique_search_users"`
	UniqueViewUsers   int64   `json:"unique_view_users"`
	ViewConversion    float64 `json:"view_conversion"`
}

type FinanceRealizationByDayRsp struct {
	BaseRsp

	Rows []FinanceRealizationByDayRowItem `json:"rows"`
}

type FinanceRealizationByDayRowItem struct {
	CommissionRatio    float64 `json:"commission_ratio"`
	DeliveryCommission struct {
		Amount           float64 `json:"amount"`
		Bonus            float64 `json:"bonus"`
		Commission       float64 `json:"commission"`
		Compensation     float64 `json:"compensation"`
		PricePerInstance float64 `json:"price_per_instance"`
		Quantity         int     `json:"quantity"`
		StandardFee      float64 `json:"standard_fee"`
		BankCoinvestment float64 `json:"bank_coinvestment"`
		Stars            float64 `json:"stars"`
		Total            float64 `json:"total"`
	} `json:"delivery_commission"`
	Item struct {
		Barcode string `json:"barcode"`
		Name    string `json:"name"`
		OfferId string `json:"offer_id"`
		Sku     int64  `json:"sku"`
	} `json:"item"`
	ReturnCommission struct {
		Amount           float64 `json:"amount"`
		Bonus            float64 `json:"bonus"`
		Commission       float64 `json:"commission"`
		Compensation     float64 `json:"compensation"`
		PricePerInstance float64 `json:"price_per_instance"`
		Quantity         int     `json:"quantity"`
		StandardFee      float64 `json:"standard_fee"`
		BankCoinvestment float64 `json:"bank_coinvestment"`
		Stars            float64 `json:"stars"`
		Total            float64 `json:"total"`
	} `json:"return_commission"`
	RowNumber              int     `json:"rowNumber"`
	SellerPricePerInstance float64 `json:"seller_price_per_instance"`
}

type SearchQueriesTextRsp struct {
	BaseRsp

	SearchQueries []SearchQueriesTextItem `json:"search_queries"`
	Offset        string                  `json:"offset"`
	Total         string                  `json:"total"`
}

type SearchQueriesTextItem struct {
	AvgPrice         float64 `json:"avg_price"`
	ConversionToCart float64 `json:"conversion_to_cart"`
	ClientCount      float64 `json:"client_count"`
	ItemsViews       float64 `json:"items_views"`
	Query            string  `json:"query"`
	AddToCart        float64 `json:"add_to_cart"`
	SellersCount     float64 `json:"sellers_count"`
}

type SearchQueriesTopRsp struct {
	BaseRsp

	SearchQueries []SearchQueriesTopItem `json:"search_queries"`
	Offset        string                 `json:"offset"`
	Total         string                 `json:"total"`
}

type SearchQueriesTopItem struct {
	AvgPrice         float64 `json:"avg_price"`
	ConversionToCart float64 `json:"conversion_to_cart"`
	ClientCount      float64 `json:"client_count"`
	ItemsViews       float64 `json:"items_views"`
	Query            string  `json:"query"`
	AddToCart        float64 `json:"add_to_cart"`
	SellersCount     float64 `json:"sellers_count"`
}

type ProductPricesDetailsRsp struct {
	BaseRsp

	Prices []ProductPricesDetailItem `json:"prices"`
}

type ProductPricesDetailItem struct {
	CustomerPrice struct {
		Amount   string `json:"amount"`
		Currency string `json:"currency"`
	} `json:"customer_price"`
	DiscountPercent float64 `json:"discount_percent"`
	OfferId         string  `json:"offer_id"`
	Price           struct {
		Amount   string `json:"amount"`
		Currency string `json:"currency"`
	} `json:"price"`
	PriceIndexes []struct {
		ExternalIndexData struct {
			MinPrice struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"min_price"`
			PriceIndex float64 `json:"price_index"`
			Url        string  `json:"url"`
		} `json:"external_index_data"`
		SelfIndexData struct {
			MinPrice struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"min_price"`
			PriceIndex float64 `json:"price_index"`
			Url        string  `json:"url"`
		} `json:"self_index_data"`
	} `json:"price_indexes"`
	Sku int64 `json:"sku"`
}
