package ozon

import "time"

type RatingIndexFbsInfoRsp struct {
	BaseRsp

	CurrencyCode string `json:"currency_code"`
	Defects      []struct {
		Date                     string  `json:"date"`
		IndexByDate              float64 `json:"index_by_date"`
		ProcessingCostsSumByDate float64 `json:"processing_costs_sum_by_date"`
	} `json:"defects"`
	Index              float64 `json:"index"`
	PeriodFrom         string  `json:"period_from"`
	PeriodTo           string  `json:"period_to"`
	ProcessingCostsSum float64 `json:"processing_costs_sum"`
}

type RatingIndexFbsPostingListRsp struct {
	BaseRsp
	Cursor string `json:"cursor"`
	Errors []struct {
		ChargePercent            float64   `json:"charge_percent"`
		ChargePrice              float64   `json:"charge_price"`
		ChargePriceCurrencyCode  string    `json:"charge_price_currency_code"`
		DeliverySchema           string    `json:"delivery_schema"`
		ErrorAt                  time.Time `json:"error_at"`
		HasGraceStatus           bool      `json:"has_grace_status"`
		Index                    float64   `json:"index"`
		PostingErrorType         string    `json:"posting_error_type"`
		PostingNumber            string    `json:"posting_number"`
		ProductPrice             float64   `json:"product_price"`
		ProductPriceCurrencyCode string    `json:"product_price_currency_code"`
	} `json:"errors"`
	HasNext bool `json:"has_next"`
}
