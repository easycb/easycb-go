package ozon

import "time"

type ReportInfoRsp struct {
	BaseRsp
	Result struct {
		AdditionalData []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"additional_data"`
		Code       string    `json:"code"`
		Status     string    `json:"status"`
		Error      string    `json:"error"`
		ExpiresAt  time.Time `json:"expires_at"`
		File       string    `json:"file"`
		ReportType string    `json:"report_type"`
		Params     struct {
		} `json:"params"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"result"`
}

type ReportListRsp struct {
	Result struct {
		Reports []ReportListReportItem `json:"reports"`
		Total   int                    `json:"total"`
	} `json:"result"`
}

type ReportListReportItem struct {
	AdditionalData []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"additional_data,omitempty"`
	Code       string    `json:"code"`
	Status     string    `json:"status"`
	Error      string    `json:"error"`
	ExpiresAt  time.Time `json:"expires_at,omitempty"`
	File       string    `json:"file"`
	ReportType string    `json:"report_type"`
	Params     struct {
		Visibility string `json:"visibility"`
	} `json:"params"`
	CreatedAt time.Time `json:"created_at"`
}

type ReportProductsCreateRsp struct {
	BaseRsp
	Result struct {
		Code string `json:"code"`
	} `json:"result"`
}

type ReportPostingCreateRsp struct {
	BaseRsp

	Result struct {
		Code string `json:"code"`
	} `json:"result"`
}

type FinanceCashFlowStatementListRsp struct {
	BaseRsp
	Result    FinanceCashFlowStatementListItem `json:"result"`
	PageCount int64                            `json:"page_count"`
}

type FinanceCashFlowStatementListItem struct {
	CashFlows []FinanceCashFlowStatementListCashFlowItem `json:"cash_flows"`
	Details   FinanceCashFlowStatementListCashDetailItem `json:"details"`
}

type FinanceCashFlowStatementListCashFlowItem struct {
	CommissionAmount            float64 `json:"commission_amount"`
	CurrencyCode                string  `json:"currency_code"`
	ItemDeliveryAndReturnAmount float64 `json:"item_delivery_and_return_amount"`
	OrdersAmount                float64 `json:"orders_amount"`
	Period                      struct {
		Begin time.Time `json:"begin"`
		End   time.Time `json:"end"`
		Id    int64     `json:"id"`
	} `json:"period"`
	ReturnsAmount  float64 `json:"returns_amount"`
	ServicesAmount float64 `json:"services_amount"`
}

type FinanceCashFlowStatementListCashDetailItem struct {
	Period struct {
		Begin time.Time `json:"begin"`
		End   time.Time `json:"end"`
		Id    int64     `json:"id"`
	} `json:"period"`
	Payments []struct {
		Payment      float64 `json:"payment"`
		CurrencyCode string  `json:"currency_code"`
	} `json:"payments"`
	BeginBalanceAmount float64 `json:"begin_balance_amount"`
	Delivery           struct {
		Total            float64 `json:"total"`
		Amount           float64 `json:"amount"`
		DeliveryServices struct {
			Total float64 `json:"total"`
			Items []struct {
				Name  string  `json:"name"`
				Price float64 `json:"price"`
			} `json:"items"`
		} `json:"delivery_services"`
	} `json:"delivery"`
	Return struct {
		Total          float64 `json:"total"`
		Amount         float64 `json:"amount"`
		ReturnServices struct {
			Total int `json:"total"`
			Items []struct {
				Name  string  `json:"name"`
				Price float64 `json:"price"`
			} `json:"items"`
		} `json:"return_services"`
	} `json:"return"`
	Loan            float64 `json:"loan"`
	InvoiceTransfer float64 `json:"invoice_transfer"`
	Rfbs            struct {
		Total                      float64 `json:"total"`
		TransferDelivery           float64 `json:"transfer_delivery"`
		TransferDeliveryReturn     float64 `json:"transfer_delivery_return"`
		CompensationDeliveryReturn float64 `json:"compensation_delivery_return"`
		PartialCompensation        float64 `json:"partial_compensation"`
		PartialCompensationReturn  float64 `json:"partial_compensation_return"`
	} `json:"rfbs"`
	Services struct {
		Total float64 `json:"total"`
		Items []struct {
			Name  string  `json:"name"`
			Price float64 `json:"price"`
		} `json:"items"`
	} `json:"services"`
	Others struct {
		Total float64 `json:"total"`
		Items []struct {
			Name  string  `json:"name"`
			Price float64 `json:"price"`
		} `json:"items"`
	} `json:"others"`
	EndBalanceAmount float64 `json:"end_balance_amount"`
}

type ReportDiscountedCreateRsp struct {
	BaseRsp
	Code string `json:"code"`
}

type ReportWarehouseStockRsp struct {
	BaseRsp
	Code string `json:"code"`
}

type ReportMarkedProductsSalesCreateRsp struct {
	BaseRsp
	Result struct {
		Code string `json:"code"`
	} `json:"result"`
}

type ReportRealizationPostingCreateRsp struct {
	BaseRsp

	Code string `json:"code"`
}
