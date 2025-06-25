package shein

type ReturnOrderListRsp struct {
	BaseRsp
	Info struct {
		Count           int                   `json:"count"`
		ReturnOrderList []ReturnOrderListItem `json:"returnOrderList"`
	} `json:"info"`
}

type ReturnOrderListItem struct {
	ReturnOrderNo     string `json:"returnOrderNo"`
	ReturnOrderStatus int    `json:"returnOrderStatus"`
	AddTime           string `json:"addTime"`
	RequestReturnTime string `json:"requestReturnTime"`
	UpdateTime        string `json:"updateTime"`
}

type ReturnOrderDetailRsp struct {
	BaseRsp
	Info []ReturnOrderDetailItem `json:"info"`
}

type ReturnOrderDetailItem struct {
	ReturnOrderNo            string        `json:"returnOrderNo"`
	ReturnOrderStatus        int           `json:"returnOrderStatus"`
	NoReturnGoodsSign        int           `json:"noReturnGoodsSign"`
	ReturnOrderTagCode       interface{}   `json:"returnOrderTagCode"`
	OrderNo                  string        `json:"orderNo"`
	Site                     string        `json:"site"`
	ShippingCode             string        `json:"shippingCode"`
	PlatformExpressNo        string        `json:"platformExpressNo"`
	MemberExpressNo          string        `json:"memberExpressNo"`
	ExpressCompanyName       string        `json:"expressCompanyName"`
	RefundOrderNos           []interface{} `json:"refundOrderNos"`
	RefundWaybill            string        `json:"refundWaybill"`
	RefundExpressCompanyName string        `json:"refundExpressCompanyName"`
	PerformanceCost          int           `json:"performanceCost"`
	InvoiceStatus            int           `json:"invoiceStatus"`
	RequestReturnTime        string        `json:"requestReturnTime"`
	AllocateTime             string        `json:"allocateTime"`
	LastUpdateTime           string        `json:"lastUpdateTime"`
	SellerSignedTime         string        `json:"sellerSignedTime"`
	CancelTime               string        `json:"cancelTime"`
	CompletedTime            string        `json:"completedTime"`
	CheckStatus              int           `json:"checkStatus"`
	StockMode                int           `json:"stockMode"`
	ReceiveType              int           `json:"receiveType"`
	ReturnGoodsInfoList      []struct {
		GoodsId                int         `json:"goodsId"`
		Sku                    string      `json:"sku"`
		Skc                    string      `json:"skc"`
		GoodsSn                string      `json:"goodsSn"`
		Currency               string      `json:"currency"`
		SaleCurrency           interface{} `json:"saleCurrency"`
		SkuSn                  string      `json:"skuSn"`
		CommodityAttributeList []struct {
			AttrValueId string `json:"attrValueId"`
			AttrName    string `json:"attrName"`
			Language    string `json:"language"`
		} `json:"commodityAttributeList"`
		GoodsTitle      string `json:"goodsTitle"`
		ImageUrl        string `json:"imageUrl"`
		GoodsStatus     int    `json:"goodsStatus"`
		ReturnImageList []struct {
			Type int    `json:"type"`
			Link string `json:"link"`
		} `json:"returnImageList"`
		SellerCurrencyPrice            float64     `json:"sellerCurrencyPrice"`
		CostPrice                      interface{} `json:"costPrice"`
		SellerCurrencyStoreCouponPrice int         `json:"sellerCurrencyStoreCouponPrice"`
		SellerCurrencyPromotionPrice   int         `json:"sellerCurrencyPromotionPrice"`
		SettleCurrencyPromotionPrice   interface{} `json:"settleCurrencyPromotionPrice"`
		EstimateCommission             float64     `json:"estimateCommission"`
		CommissionSaleTax              int         `json:"commissionSaleTax"`
		PerformancePrice               int         `json:"performancePrice"`
		ReturnExpense                  int         `json:"returnExpense"`
		SellerRealTax                  int         `json:"sellerRealTax"`
		EstimateIncomeMoney            float64     `json:"estimateIncomeMoney"`
		EstimateTaxIncomeMoney         float64     `json:"estimateTaxIncomeMoney"`
		ReturnReasonList               []struct {
			Language string `json:"language"`
			Reason   string `json:"reason"`
		} `json:"returnReasonList"`
	} `json:"returnGoodsInfoList"`
}

type SignReturnOrderRsp struct {
	BaseRsp
	Info struct {
		ReturnOrderNo string  `json:"returnOrderNo"`
		GoodsIdList   []int64 `json:"goodsIdList"`
	} `json:"info"`
}
