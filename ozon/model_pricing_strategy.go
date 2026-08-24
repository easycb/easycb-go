package ozon

import "time"

type PricingStrategyCompetitorsListRsp struct {
	BaseRsp
	Competitor []PricingStrategyCompetitorsListItem `json:"competitor"`
	Total      int                                  `json:"total"`
}

type PricingStrategyCompetitorsListItem struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type PricingStrategyListRsp struct {
	BaseRsp
	Strategies []PricingStrategyListItem `json:"strategies"`
	Total      int                       `json:"total"`
}

type PricingStrategyListItem struct {
	Id               string    `json:"id"`
	Name             string    `json:"name"`
	Type             string    `json:"type"`
	UpdateType       string    `json:"update_type"`
	UpdatedAt        time.Time `json:"updated_at"`
	ProductsCount    int64     `json:"products_count"`
	CompetitorsCount int64     `json:"competitors_count"`
	Enabled          bool      `json:"enabled"`
}

type PricingStrategyCreateRsp struct {
	BaseRsp
	Result struct {
		Id string `json:"id"`
	} `json:"result"`
}

type PricingStrategyInfoRsp struct {
	BaseRsp
	Result struct {
		StrategyName string `json:"strategy_name"`
		Enabled      bool   `json:"enabled"`
		UpdateType   string `json:"update_type"`
		Type         string `json:"type"`
		Competitors  []struct {
			CompetitorId int64   `json:"competitor_id"`
			Coefficient  float64 `json:"coefficient"`
		} `json:"competitors"`
	} `json:"result"`
}

type PricingStrategyUpdateRsp struct {
	BaseRsp
}

type PricingStrategyProductAddRsp struct {
	BaseRsp
	Result struct {
		FailedProductCount int `json:"failed_product_count"`
		Errors             []struct {
			Code      string `json:"code"`
			Error     string `json:"error"`
			ProductId int64  `json:"product_id"`
		} `json:"errors"`
	} `json:"result"`
}

type PricingStrategyStrategyIdsByProductIdsRsp struct {
	BaseRsp

	Result struct {
		ProductsInfo []struct {
			ProductId  int64  `json:"product_id"`
			StrategyId string `json:"strategy_id"`
		} `json:"products_info"`
	} `json:"result"`
}

type PricingStrategyProductListRsp struct {
	BaseRsp
	Result struct {
		ProductId []int64 `json:"product_id"`
	} `json:"result"`
}

type PricingStrategyProductInfoRsp struct {
	BaseRsp

	Result struct {
		StrategyId                   string    `json:"strategy_id"`
		IsEnabled                    bool      `json:"is_enabled"`
		StrategyProductPrice         float64   `json:"strategy_product_price"`
		PriceDownloadedAt            time.Time `json:"price_downloaded_at"`
		StrategyCompetitorId         string    `json:"strategy_competitor_id"`
		StrategyCompetitorProductUrl string    `json:"strategy_competitor_product_url"`
	} `json:"result"`
}

type PricingStrategyProductDeleteRsp struct {
	BaseRsp
	Result struct {
		FailedProductCount int `json:"failed_product_count"`
	} `json:"result"`
}

type PricingStrategyStatusRsp struct {
	BaseRsp
}

type PricingStrategyDeleteRsp struct {
	BaseRsp
}
