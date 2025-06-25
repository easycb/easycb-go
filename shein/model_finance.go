package shein

type FinanceReportListRsp struct {
	BaseRsp
	Info struct {
		ReportOrderInfos []FinanceReportListItem `json:"reportOrderInfos"`
		Count            int                     `json:"count"`
	} `json:"info"`
}

type FinanceReportListItem struct {
	ReportOrderNo            string  `json:"reportOrderNo"`
	SalesTotal               int     `json:"salesTotal"`
	ReplenishTotal           int     `json:"replenishTotal"`
	AddTime                  string  `json:"addTime"`
	LastUpdateTime           string  `json:"lastUpdateTime"`
	SettlementStatus         int     `json:"settlementStatus"`
	SettlementStatusName     string  `json:"settlementStatusName"`
	EstimatePayTime          string  `json:"estimatePayTime"`
	CompletedPayTime         string  `json:"completedPayTime"`
	CompanyName              string  `json:"companyName"`
	EstimateIncomeMoneyTotal float64 `json:"estimateIncomeMoneyTotal"`
	CurrencyCode             string  `json:"currencyCode"`
}

type FinanceReportSalesDetailRsp struct {
	BaseRsp
	Info struct {
		ReportSalesDetails []FinanceReportSalesDetailItem `json:"reportSalesDetails"`
		Count              int                            `json:"count"`
		Query              interface{}                    `json:"query"`
	} `json:"info"`
}

type FinanceReportSalesDetailItem struct {
	SecondOrderType     int     `json:"secondOrderType"`
	SecondOrderTypeName string  `json:"secondOrderTypeName"`
	InAndOut            int     `json:"inAndOut"`
	InAndOutName        string  `json:"inAndOutName"`
	BzOrderNo           string  `json:"bzOrderNo"`
	Id                  string  `json:"id"`
	SkcName             string  `json:"skcName"`
	SkuCode             string  `json:"skuCode"`
	GoodsCount          int     `json:"goodsCount"`
	SettleCurrencyCode  string  `json:"settleCurrencyCode"`
	Amount              float64 `json:"amount"`
	CompanyName         string  `json:"companyName"`
	AddTime             string  `json:"addTime"`
	UnitPrice           float64 `json:"unitPrice"`
}

type FinanceAdjustmentDetailRsp struct {
	BaseRsp
	Info struct {
		ReportReplenishDetail []FinanceReportReplenishDetailItem `json:"reportReplenishDetail"`
		Count                 int                                `json:"count"`
		Query                 interface{}                        `json:"query"`
	} `json:"info"`
}

type FinanceReportReplenishDetailItem struct {
	ReplenishNo        string      `json:"replenishNo"`
	ReplenishType      int         `json:"replenishType"`
	ReplenishTypeName  string      `json:"replenishTypeName"`
	ReplenishCategory  string      `json:"replenishCategory"`
	Id                 string      `json:"id"`
	BzOrderNo          string      `json:"bzOrderNo"`
	UnitPrice          float64     `json:"unitPrice"`
	SkcName            string      `json:"skcName"`
	SkuCode            interface{} `json:"skuCode"`
	GoodsCount         int         `json:"goodsCount"`
	SettleCurrencyCode string      `json:"settleCurrencyCode"`
	Amount             float64     `json:"amount"`
	CompanyName        string      `json:"companyName"`
	AddTime            string      `json:"addTime"`
}

type FinanceGetCheckOrderListRsp struct {
	BaseRsp
	Info struct {
		Count int                            `json:"count"`
		List  []FinanceGetCheckOrderListItem `json:"list"`
	} `json:"info"`
}

type FinanceGetCheckOrderListItem struct {
	CheckOrderNo             string      `json:"checkOrderNo"`
	CheckStatus              int         `json:"checkStatus"`
	SecondOrderType          int         `json:"secondOrderType"`
	IncomeExpenditureType    int         `json:"incomeExpenditureType"`
	BzOrderNo                string      `json:"bzOrderNo"`
	BusinessCompletedTime    string      `json:"businessCompletedTime"`
	CompletedPayTime         string      `json:"completedPayTime"`
	Site                     string      `json:"site"`
	CurrencyCode             string      `json:"currencyCode"`
	EstimateIncomeMoneyTotal float64     `json:"estimateIncomeMoneyTotal"`
	ReportOrderNo            interface{} `json:"reportOrderNo"`
	EstimatePayTime          string      `json:"estimatePayTime"`
}

type FinanceGetCheckOrderDetailRsp struct {
	BaseRsp
	Info struct {
		FinanceGetCheckOrderDetailItem
	} `json:"info"`
}

type FinanceGetCheckOrderDetailItem struct {
	CheckOrderNo             string      `json:"checkOrderNo"`
	CheckStatus              int         `json:"checkStatus"`
	SecondOrderType          int         `json:"secondOrderType"`
	IncomeExpenditureType    int         `json:"incomeExpenditureType"`
	BzOrderNo                string      `json:"bzOrderNo"`
	BusinessCompletedTime    string      `json:"businessCompletedTime"`
	CompletedPayTime         string      `json:"completedPayTime"`
	Site                     string      `json:"site"`
	CurrencyCode             string      `json:"currencyCode"`
	EstimateIncomeMoneyTotal float64     `json:"estimateIncomeMoneyTotal"`
	ReportOrderNo            interface{} `json:"reportOrderNo"`
	EstimatePayTime          string      `json:"estimatePayTime"`
	ItemList                 []struct {
		SkuCode                      string        `json:"skuCode"`
		SellerCurrencyPrice          float64       `json:"sellerCurrencyPrice"`
		CostPrice                    float64       `json:"costPrice"`
		CouponAmount                 float64       `json:"couponAmount"`
		SettleCurrencyPromotionPrice float64       `json:"settleCurrencyPromotionPrice"`
		SellerCurrencyPromotionPrice float64       `json:"sellerCurrencyPromotionPrice"`
		SellerRealTax                float64       `json:"sellerRealTax"`
		CommissionSaleTax            float64       `json:"commissionSaleTax"`
		WhtTotalAmount               float64       `json:"whtTotalAmount"`
		SubjectAmountList            []interface{} `json:"subjectAmountList"`
		SubsidyTotalAmount           float64       `json:"subsidyTotalAmount"`
		CommissionRate               float64       `json:"commissionRate"`
		CommissionAmount             float64       `json:"commissionAmount"`
		PerformanceCost              float64       `json:"performanceCost"`
		StockExpense                 float64       `json:"stockExpense"`
		ReturnExpense                float64       `json:"returnExpense"`
		ReturnFreightSubsidy         float64       `json:"returnFreightSubsidy"`
		ServiceFee                   float64       `json:"serviceFee"`
		RefundRatio                  string        `json:"refundRatio"`
		IncomeAmount                 float64       `json:"incomeAmount"`
	} `json:"itemList"`
}
