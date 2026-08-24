package ozon

import "time"

type ActionsRsp struct {
	BaseRsp
	Result []ActionItem `json:"result"`
}

type ActionItem struct {
	Id                         float64     `json:"id"`
	Title                      string      `json:"title"`
	DateStart                  time.Time   `json:"date_start"`
	DateEnd                    time.Time   `json:"date_end"`
	AutoAddDates               []time.Time `json:"auto_add_dates"`
	PotentialProductsCount     float64     `json:"potential_products_count"`
	IsParticipating            bool        `json:"is_participating"`
	ParticipatingProductsCount float64     `json:"participating_products_count"`
	Description                string      `json:"description"`
	ActionType                 string      `json:"action_type"`
	BannedProductsCount        float64     `json:"banned_products_count"`
	WithTargeting              bool        `json:"with_targeting"`
	DiscountType               string      `json:"discount_type"`
	DiscountValue              float64     `json:"discount_value"`
	OrderAmount                float64     `json:"order_amount"`
	FreezeDate                 string      `json:"freeze_date"`
	IsVoucherAction            bool        `json:"is_voucher_action"`
}

type ActionsCandidatesRsp struct {
	BaseRsp
	Result ActionsCandidateItem `json:"result"`
}

type ActionsCandidateItem struct {
	Products []ActionsCandidateProductItem `json:"products"`
	Total    int                           `json:"total"`
	LastId   string                        `json:"last_id"`
}

type ActionsCandidateProductItem struct {
	Id                        int64   `json:"id"`
	Price                     float64 `json:"price"`
	ActionPrice               float64 `json:"action_price"`
	AlertMaxActionPriceFailed bool    `json:"alert_max_action_price_failed"`
	AlertMaxActionPrice       float64 `json:"alert_max_action_price"`
	MaxActionPrice            float64 `json:"max_action_price"`
	AddMode                   string  `json:"add_mode"`
	Stock                     int     `json:"stock"`
	MinStock                  int     `json:"min_stock"`
	CurrentBoost              float64 `json:"current_boost"`
	PriceMinElastic           float64 `json:"price_min_elastic"`
	PriceMaxElastic           float64 `json:"price_max_elastic"`
	MinBoost                  float64 `json:"min_boost"`
	MaxBoost                  float64 `json:"max_boost"`
}

type ActionsProductsRsp struct {
	BaseRsp
	Result ActionsProductsResult `json:"result"`
}

type ActionsProductsResult struct {
	Products []ActionsProductsResultProductItem `json:"products"`
	Total    int64                              `json:"total"`
	LastId   string                             `json:"last_id"`
}

type ActionsProductsResultProductItem struct {
	Id                        int64   `json:"id"`
	Price                     float64 `json:"price"`
	ActionPrice               float64 `json:"action_price"`
	AlertMaxActionPriceFailed bool    `json:"alert_max_action_price_failed"`
	AlertMaxActionPrice       float64 `json:"alert_max_action_price"`
	MaxActionPrice            float64 `json:"max_action_price"`
	AddMode                   string  `json:"add_mode"`
	Stock                     int     `json:"stock"`
	MinStock                  int     `json:"min_stock"`
	CurrentBoost              float64 `json:"current_boost"`
	PriceMinElastic           float64 `json:"price_min_elastic"`
	PriceMaxElastic           float64 `json:"price_max_elastic"`
	MinBoost                  float64 `json:"min_boost"`
	MaxBoost                  float64 `json:"max_boost"`
}

type ActionsProductsActivateRsp struct {
	BaseRsp
	Result ActionsProductsActivateResult `json:"result"`
}

type ActionsProductsActivateResult struct {
	ProductIds []int64                                 `json:"product_ids"`
	Rejected   []ActionsProductsActivateResultRejected `json:"rejected"`
}

type ActionsProductsActivateResultRejected struct {
	ProductId int64  `json:"product_id"`
	Reason    string `json:"reason"`
}

type ActionsProductsDeactivateRsp struct {
	BaseRsp
	Result ActionsProductsDeactivateResult `json:"result"`
}

type ActionsProductsDeactivateResult struct {
	ProductIds []int64                                   `json:"product_ids"`
	Rejected   []ActionsProductsDeactivateResultRejected `json:"rejected"`
}

type ActionsProductsDeactivateResultRejected struct {
	ProductId int64  `json:"product_id"`
	Reason    string `json:"reason"`
}

type ActionsDiscountsTaskListRsp struct {
	BaseRsp
	Result []ActionsDiscountsTaskListItem `json:"result"`
}

type ActionsDiscountsTaskListItem struct {
	Id                      int64     `json:"id"`
	CreatedAt               time.Time `json:"created_at"`
	EndAt                   time.Time `json:"end_at"`
	EditedTill              time.Time `json:"edited_till"`
	Status                  string    `json:"status"`
	CustomerName            string    `json:"customer_name"`
	Sku                     int64     `json:"sku"`
	UserComment             string    `json:"user_comment"`
	SellerComment           string    `json:"seller_comment"`
	RequestedPrice          float64   `json:"requested_price"`
	ApprovedPrice           float64   `json:"approved_price"`
	OriginalPrice           float64   `json:"original_price"`
	Discount                float64   `json:"discount"`
	DiscountPercent         float64   `json:"discount_percent"`
	BasePrice               float64   `json:"base_price"`
	MinAutoPrice            float64   `json:"min_auto_price"`
	PrevTaskId              int64     `json:"prev_task_id"`
	IsDamaged               bool      `json:"is_damaged"`
	ModeratedAt             time.Time `json:"moderated_at"`
	ApprovedDiscount        float64   `json:"approved_discount"`
	ApprovedDiscountPercent float64   `json:"approved_discount_percent"`
	IsPurchased             bool      `json:"is_purchased"`
	IsAutoModerated         bool      `json:"is_auto_moderated"`
	OfferId                 string    `json:"offer_id"`
	Email                   string    `json:"email"`
	LastName                string    `json:"last_name"`
	FirstName               string    `json:"first_name"`
	Patronymic              string    `json:"patronymic"`
	ApprovedQuantityMin     int64     `json:"approved_quantity_min"`
	ApprovedQuantityMax     int64     `json:"approved_quantity_max"`
	RequestedQuantityMin    int64     `json:"requested_quantity_min"`
	RequestedQuantityMax    int64     `json:"requested_quantity_max"`
	RequestedPriceWithFee   float64   `json:"requested_price_with_fee"`
	ApprovedPriceWithFee    float64   `json:"approved_price_with_fee"`
	ApprovedPriceFeePercent float64   `json:"approved_price_fee_percent"`
}

type ActionsDiscountsTaskApproveRsp struct {
	BaseRsp
	Result struct {
		FailDetails []struct {
			TaskId       int64  `json:"task_id"`
			ErrorForUser string `json:"error_for_user"`
		} `json:"fail_details"`
		SuccessCount int64 `json:"success_count"`
		FailCount    int64 `json:"fail_count"`
	} `json:"result"`
}

type ActionsDiscountsTaskDeclineRsp struct {
	BaseRsp
	Result struct {
		FailDetails []struct {
			TaskId       int    `json:"task_id"`
			ErrorForUser string `json:"error_for_user"`
		} `json:"fail_details"`
		SuccessCount int `json:"success_count"`
		FailCount    int `json:"fail_count"`
	} `json:"result"`
}

type ActionsDiscountsTaskListV2Rsp struct {
	BaseRsp

	Tasks []ActionsDiscountsTaskListV2TaskItem `json:"tasks"`
}

type ActionsDiscountsTaskListV2TaskItem struct {
	ApprovedDiscount    float64 `json:"approved_discount"`
	ApprovedPrice       float64 `json:"approved_price"`
	ApprovedQuantityMax int64   `json:"approved_quantity_max"`
	AutoModeratedInfo   struct {
		MaxPercent float64 `json:"max_percent"`
		MaxPrice   float64 `json:"max_price"`
		MinPercent float64 `json:"min_percent"`
		MinPrice   float64 `json:"min_price"`
	} `json:"auto_moderated_info"`
	CreatedAt            time.Time `json:"created_at"`
	EditedTill           time.Time `json:"edited_till"`
	EditedTillDuration   int64     `json:"edited_till_duration"`
	Email                string    `json:"email"`
	EndAt                time.Time `json:"end_at"`
	EndAtDuration        int64     `json:"end_at_duration"`
	FirstName            string    `json:"first_name"`
	Id                   int64     `json:"id"`
	IsAutoModerated      bool      `json:"is_auto_moderated"`
	LastName             string    `json:"last_name"`
	MinAutoPrice         float64   `json:"min_auto_price"`
	ModeratedAt          time.Time `json:"moderated_at"`
	Name                 string    `json:"name"`
	OriginalPrice        float64   `json:"original_price"`
	Patronymic           string    `json:"patronymic"`
	ReductionFactor      float64   `json:"reduction_factor"`
	RequestedDiscount    float64   `json:"requested_discount"`
	RequestedPrice       float64   `json:"requested_price"`
	RequestedQuantityMax int       `json:"requested_quantity_max"`
	Sku                  int64     `json:"sku"`
	Status               string    `json:"status"`
}

type ActionsAutoAddProductsListRsp struct {
	BaseRsp

	Products []ActionsAutoAddProductsListProductItem `json:"products"`
	Total    string                                  `json:"total"`
}

type ActionsAutoAddProductsListProductItem struct {
	ProduсtId              int64   `json:"produсt_id"`
	OfferId                string  `json:"offer_id"`
	Sku                    int64   `json:"sku"`
	Name                   string  `json:"name"`
	Price                  float64 `json:"price"`
	MaxDiscountPrice       float64 `json:"max_discount_price"`
	MinSellerPrice         float64 `json:"min_seller_price:"`
	MarketplaceSellerPrice float64 `json:"marketplace_seller_price"`
	ActionPriceToAutoAdd   float64 `json:"action_price_to_auto_add"`
	MinActionQuantity      string  `json:"min_action_quantity"`
	QuantityToAutoAdd      string  `json:"quantity_to_auto_add"`
	Currency               string  `json:"currency"`
	AddMode                string  `json:"add_mode"`
}

type ActionsAutoAddProductsCandidatesRsp struct {
	BaseRsp

	Products []ActionsAutoAddProductsCandidatesProductItem `json:"products"`
	Total    string                                        `json:"total"`
}

type ActionsAutoAddProductsCandidatesProductItem struct {
	ProductId              int64   `json:"produсt_id"`
	OfferId                string  `json:"offer_id"`
	Sku                    int64   `json:"sku"`
	Name                   string  `json:"name"`
	Price                  float64 `json:"price"`
	BasePrice              float64 `json:"base_price"`
	MaxDiscountPrice       float64 `json:"max_discount_price"`
	MinSellerPrice         float64 `json:"min_seller_price:"`
	MarketplaceSellerPrice float64 `json:"marketplace_seller_price"`
	ActionPriceToAutoAdd   float64 `json:"action_price_to_auto_add"`
	MinActionQuantity      string  `json:"min_action_quantity"`
	QuantityToAutoAdd      string  `json:"quantity_to_auto_add"`
	Currency               string  `json:"currency"`
}

type ActionsAutoAddProductsDeleteRsp struct {
	BaseRsp

	ProductIds []int64 `json:"product_ids"`
}

type ActionsAutoAddProductsUpdateRsp struct {
	BaseRsp

	UpdatedIds []int64 `json:"updated_ids"`
	Rejected   []struct {
		Code      string `json:"code"`
		ProductId int64  `json:"product_id"`
		Reason    string `json:"reason"`
	} `json:"rejected"`
	BelowMinPrice []struct {
		Key   string  `json:"key"`
		Value float64 `json:"value"`
	} `json:"below_min_price"`
	ExtremelyLowPrice []struct {
		Key   string  `json:"key"`
		Value float64 `json:"value"`
	} `json:"extremely_low_price"`
	FailedPrice []struct {
		Key   string  `json:"key"`
		Value float64 `json:"value"`
	} `json:"failed_price"`
}
