package ozon

import "time"

type ReturnsListRsp struct {
	BaseRsp

	Returns []ReturnsListReturnItem `json:"returns"`
	HasNext bool                    `json:"has_next"`
}

type ReturnsListReturnItem struct {
	Exemplars []struct {
		Id string `json:"id"`
	} `json:"exemplars"`
	Id               int64  `json:"id"`
	CompanyId        int64  `json:"company_id"`
	ReturnReasonName string `json:"return_reason_name"`
	Type             string `json:"type"`
	Schema           string `json:"schema"`
	OrderId          int64  `json:"order_id"`
	OrderNumber      string `json:"order_number"`
	Place            struct {
		Id      int64  `json:"id"`
		Name    string `json:"name"`
		Address string `json:"address"`
	} `json:"place"`
	TargetPlace struct {
		Id      int64  `json:"id"`
		Name    string `json:"name"`
		Address string `json:"address"`
	} `json:"target_place"`
	Storage struct {
		Sum struct {
			CurrencyCode string `json:"currency_code"`
			Price        string `json:"price"`
		} `json:"sum"`
		TarifficationFirstDate time.Time `json:"tariffication_first_date"`
		TarifficationStartDate time.Time `json:"tariffication_start_date"`
		ArrivedMoment          time.Time `json:"arrived_moment"`
		Days                   string    `json:"days"`
		UtilizationSum         struct {
			CurrencyCode string `json:"currency_code"`
			Price        string `json:"price"`
		} `json:"utilization_sum"`
		UtilizationForecastDate time.Time `json:"utilization_forecast_date"`
	} `json:"storage"`
	Product struct {
		Sku     string `json:"sku"`
		OfferId string `json:"offer_id"`
		Name    string `json:"name"`
		Price   struct {
			CurrencyCode string `json:"currency_code"`
			Price        string `json:"price"`
		} `json:"price"`
		PriceWithoutCommission struct {
			CurrencyCode string `json:"currency_code"`
			Price        string `json:"price"`
		} `json:"price_without_commission"`
		CommissionPercent string `json:"commission_percent"`
		Commission        struct {
			CurrencyCode string `json:"currency_code"`
			Price        string `json:"price"`
		} `json:"commission"`
	} `json:"product"`
	Logistic struct {
		TechnicalReturnMoment           time.Time `json:"technical_return_moment"`
		FinalMoment                     time.Time `json:"final_moment"`
		CancelledWithCompensationMoment time.Time `json:"cancelled_with_compensation_moment"`
		ReturnDate                      time.Time `json:"return_date"`
		Barcode                         string    `json:"barcode"`
	} `json:"logistic"`
	Visual struct {
		Status struct {
			Id          int    `json:"id"`
			DisplayName string `json:"display_name"`
			SysName     string `json:"sys_name"`
		} `json:"status"`
		ChangeMoment time.Time `json:"change_moment"`
	} `json:"visual"`
	AdditionalInfo struct {
		IsOpened      bool `json:"is_opened"`
		IsSuperEconom bool `json:"is_super_econom"`
	} `json:"additional_info"`
	SourceId           string `json:"source_id"`
	PostingNumber      string `json:"posting_number"`
	ClearingId         string `json:"clearing_id"`
	ReturnClearingId   int64  `json:"return_clearing_id"`
	CompensationStatus struct {
		Status struct {
			Id          int    `json:"id"`
			DisplayName string `json:"display_name"`
			SysName     string `json:"sys_name"`
		} `json:"status"`
		ChangeMoment time.Time `json:"change_moment"`
	} `json:"compensation_status"`
}

type ReturnsRFbsListRsp struct {
	BaseRsp
	Returns []ReturnRFbsListReturnItem `json:"returns"`
}

type ReturnRFbsListReturnItem struct {
	ReturnId      int64     `json:"return_id"`
	ReturnNumber  string    `json:"return_number"`
	PostingNumber string    `json:"posting_number"`
	OrderNumber   string    `json:"order_number"`
	CreatedAt     time.Time `json:"created_at"`
	Product       struct {
		Sku          int64   `json:"sku"`
		OfferId      string  `json:"offer_id"`
		Name         string  `json:"name"`
		Price        float64 `json:"price"`
		CurrencyCode string  `json:"currency_code"`
	} `json:"product"`
	State struct {
		GroupState string `json:"group_state"`
		State      string `json:"state"`
		StateName  string `json:"state_name"`
	} `json:"state"`
}

type ReturnsRFbsGetRsp struct {
	BaseRsp
	Returns ReturnsRFbsGetReturnItem `json:"returns"`
}

type ReturnsRFbsGetReturnItem struct {
	AvailableActions []struct {
		Action struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		} `json:"action"`
	} `json:"available_actions"`
	ClientName             string   `json:"client_name"`
	ClientPhoto            []string `json:"client_photo"`
	ClientReturnMethodType struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"client_return_method_type"`
	Comment       string    `json:"comment"`
	CreatedAt     time.Time `json:"created_at"`
	OrderNumber   string    `json:"order_number"`
	PostingNumber string    `json:"posting_number"`
	Product       struct {
		CurrencyCode string  `json:"currency_code"`
		Name         string  `json:"name"`
		OfferId      string  `json:"offer_id"`
		Price        float64 `json:"price"`
		Sku          int64   `json:"sku"`
	} `json:"product"`
	RejectionComment string `json:"rejection_comment"`
	RejectionReason  []struct {
		Hint              string `json:"hint"`
		Id                int    `json:"id"`
		IsCommentRequired bool   `json:"is_comment_required"`
		Name              string `json:"name"`
	} `json:"rejection_reason"`
	ReturnMethodDescription string `json:"return_method_description"`
	ReturnNumber            string `json:"return_number"`
	ReturnReason            struct {
		Id       int    `json:"id"`
		IsDefect bool   `json:"is_defect"`
		Name     string `json:"name"`
	} `json:"return_reason"`
	RuPostTrackingNumber string `json:"ru_post_tracking_number"`
	State                struct {
		State     string `json:"state"`
		StateName string `json:"state_name"`
	} `json:"state"`
	WarehouseId int64 `json:"warehouse_id"`
}

type ReturnsRFbsActionSetRsp struct {
	BaseRsp
}
