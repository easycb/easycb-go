package ozon

import "time"

type SellerActionsCreateDiscountRsp struct {
	BaseRsp

	ActionId int64 `json:"action_id"`
}

type SellerActionsCreateDiscountWithConditionRsp struct {
	BaseRsp

	ActionId int64 `json:"action_id"`
}

type SellerActionsCreateInstallmentRsp struct {
	BaseRsp

	ActionId int64 `json:"action_id"`
}

type SellerActionsCreateMultiLevelDiscountRsp struct {
	BaseRsp

	ActionId int64 `json:"action_id"`
}

type SellerActionsCreateVoucherRsp struct {
	BaseRsp

	ActionId int64 `json:"action_id"`
}

type SellerActionsUpdateDiscountRsp struct {
	BaseRsp
}

type SellerActionsUpdateDiscountWithConditionRsp struct {
	BaseRsp
}

type SellerActionsUpdateInstallmentRsp struct {
	BaseRsp
}

type SellerActionsUpdateMultiLevelDiscountRsp struct {
	BaseRsp
}

type SellerActionsUpdateVoucherRsp struct {
	BaseRsp
}

type SellerActionsProductsAddRsp struct {
	BaseRsp
}

type SellerActionsProductsCandidatesRsp struct {
	BaseRsp

	Cursor   int64                                 `json:"cursor"`
	HasNext  bool                                  `json:"has_next"`
	Products []SellerActionsProductsCandidatesItem `json:"products"`
}

type SellerActionsProductsCandidatesItem struct {
	ActionPrice     float64 `json:"action_price"`
	BasePrice       float64 `json:"base_price"`
	Currency        string  `json:"currency"`
	DiscountPercent float64 `json:"discount_percent"`
	IsActive        bool    `json:"is_active"`
	MinSellerPrice  float64 `json:"min_seller_price"`
	Name            string  `json:"name"`
	OfferId         string  `json:"offer_id"`
	Price           float64 `json:"price"`
	ProductId       int64   `json:"product_id"`
	QuantSize       int64   `json:"quant_size"`
	QuantType       string  `json:"quant_type"`
	Sku             []int64 `json:"sku"`
}

type SellerActionsProductsDeleteRsp struct {
	BaseRsp
}

type SellerActionsProductsListRsp struct {
	BaseRsp

	Cursor   int64                           `json:"cursor"`
	HasNext  bool                            `json:"has_next"`
	Products []SellerActionsProductsListItem `json:"products"`
}

type SellerActionsProductsListItem struct {
	ActionPrice     float64 `json:"action_price"`
	BasePrice       float64 `json:"base_price"`
	Currency        string  `json:"currency"`
	DiscountPercent float64 `json:"discount_percent"`
	IsActive        bool    `json:"is_active"`
	MinSellerPrice  float64 `json:"min_seller_price"`
	Name            string  `json:"name"`
	OfferId         string  `json:"offer_id"`
	Price           float64 `json:"price"`
	ProductId       int64   `json:"product_id"`
	QuantSize       int64   `json:"quant_size"`
	QuantType       string  `json:"quant_type"`
	Sku             []int64 `json:"sku"`
}

type SellerActionsArchiveRsp struct {
	BaseRsp
}

type SellerActionsChangeActivityRsp struct {
	BaseRsp
}

type SellerActionsListRsp struct {
	BaseRsp

	Actions []SellerActionsListItem `json:"actions"`
	Total   int64                   `json:"total"`
}

type SellerActionsListItem struct {
	ActionId         int64 `json:"action_id"`
	ActionParameters struct {
		Addresses            []string  `json:"addresses"`
		AutoStopActionReason string    `json:"auto_stop_action_reason"`
		Budget               float64   `json:"budget"`
		BudgetSpent          float64   `json:"budget_spent"`
		DateEnd              time.Time `json:"date_end"`
		DateStart            time.Time `json:"date_start"`
		DiscountLevels       []struct {
			DiscountValue float64 `json:"discount_value"`
			OrderAmount   float64 `json:"order_amount"`
		} `json:"discount_levels"`
		DiscountType           string  `json:"discount_type"`
		DiscountValue          float64 `json:"discount_value"`
		IsLegalEntitiesSegment bool    `json:"is_legal_entities_segment"`
		MinActionPercent       float64 `json:"min_action_percent"`
		MinOrderAmount         float64 `json:"min_order_amount"`
		PickedSegments         []struct {
			Segments []struct {
				Description string `json:"description"`
				Id          int64  `json:"id"`
				Name        string `json:"name"`
				Type        string `json:"type"`
			} `json:"segments"`
		} `json:"picked_segments"`
		Status            string `json:"status"`
		Title             string `json:"title"`
		Type              string `json:"type"`
		VoucherParameters struct {
			CountCodes int64  `json:"count_codes"`
			IsPrivate  bool   `json:"is_private"`
			Type       string `json:"type"`
		} `json:"voucher_parameters"`
		Warehouses []string `json:"warehouses"`
	} `json:"action_parameters"`
	AllowDelete    bool   `json:"allow_delete"`
	HighlightUrl   string `json:"highlight_url"`
	IsEditable     bool   `json:"is_editable"`
	IsParticipated bool   `json:"is_participated"`
	IsTurnOn       bool   `json:"is_turn_on"`
	SkuCount       int64  `json:"sku_count"`
}

type SellerActionsVoucherGetRsp struct {
	BaseRsp

	File string `json:"file"`
}
