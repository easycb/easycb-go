package shein

type OrderListRsp struct {
	BaseRsp
	Info struct {
		Count     int             `json:"count"`
		OrderList []OrderListItem `json:"orderList"`
	} `json:"info"`
}

type OrderListItem struct {
	OrderNo         string `json:"orderNo"`
	OrderStatus     string `json:"orderStatus"`
	OrderCreateTime string `json:"orderCreateTime"`
	OrderUpdateTime string `json:"orderUpdateTime"`
}

type OrderDetailRsp struct {
	BaseRsp
	Info []OrderDetailItem `json:"info"`
}

type OrderDetailItem struct {
	OrderNo            string      `json:"orderNo"`
	OrderType          int         `json:"orderType"`
	PerformanceType    int         `json:"performanceType"`
	OrderStatus        int         `json:"orderStatus"`
	IsCod              int         `json:"isCod"`
	IsOverLimitOrder   int         `json:"isOverLimitOrder"`
	UnpackingStatus    interface{} `json:"unpackingStatus"`
	OrderTag           int         `json:"orderTag"`
	DeliveryType       int         `json:"deliveryType"`
	PrintOrderStatus   int         `json:"printOrderStatus"`
	InvoiceStatus      int         `json:"invoiceStatus"`
	OrderGoodsInfoList []struct {
		GoodsId        int    `json:"goodsId"`
		SkuCode        string `json:"skuCode"`
		Skc            string `json:"skc"`
		GoodsSn        string `json:"goodsSn"`
		SellerSku      string `json:"sellerSku"`
		GoodsStatus    int    `json:"goodsStatus"`
		NewGoodsStatus int    `json:"newGoodsStatus"`
		SkuAttribute   []struct {
			AttrValueId string `json:"attrValueId"`
			AttrName    string `json:"attrName"`
			Language    string `json:"language"`
		} `json:"skuAttribute"`
		GoodsTitle                    string      `json:"goodsTitle"`
		SpuPicURL                     string      `json:"spuPicURL"`
		GoodsWeight                   float64     `json:"goodsWeight"`
		StorageTag                    int         `json:"storageTag"`
		PerformanceTag                int         `json:"performanceTag"`
		GoodsExchangeTag              int         `json:"goodsExchangeTag"`
		UnpackingGroupNo              string      `json:"unpackingGroupNo"`
		UnpackingGroupInvoiceStatus   interface{} `json:"unpackingGroupInvoiceStatus"`
		BeExchangeEntityId            int64       `json:"beExchangeEntityId"`
		OrderCurrency                 string      `json:"orderCurrency"`
		SaleCurrency                  interface{} `json:"saleCurrency"`
		SellerCurrencyPrice           float64     `json:"sellerCurrencyPrice"`
		CostPrice                     float64     `json:"costPrice"`
		OrderCurrencyStoreCouponPrice float64     `json:"orderCurrencyStoreCouponPrice"`
		OrderCurrencyPromotionPrice   float64     `json:"orderCurrencyPromotionPrice"`
		SettleCurrencyPromotionPrice  float64     `json:"settleCurrencyPromotionPrice"`
		Commission                    float64     `json:"commission"`
		CommissionRate                float64     `json:"commissionRate"`
		ServiceCharge                 float64     `json:"serviceCharge"`
		PerformanceServiceCharge      float64     `json:"performanceServiceCharge"`
		EstimatedIncome               float64     `json:"estimatedIncome"`
		SpuName                       string      `json:"spuName"`
		SaleTax                       float64     `json:"saleTax"`
		SellerRealTax                 float64     `json:"sellerRealTax"`
		CommissionSaleTax             float64     `json:"commissionSaleTax"`
		WarehouseCode                 string      `json:"warehouseCode"`
		WarehouseName                 string      `json:"warehouseName"`
		SellerCurrencyDiscountPrice   float64     `json:"sellerCurrencyDiscountPrice"`
		MxTaxPrice                    float64     `json:"mxTaxPrice"`
	} `json:"orderGoodsInfoList"`
	PackageWaybillList []struct {
		PackageNo            string `json:"packageNo"`
		WaybillNo            string `json:"waybillNo"`
		Carrier              string `json:"carrier"`
		CarrierCode          string `json:"carrierCode"`
		ProductInventoryList []struct {
			ProductId string `json:"productId"`
		} `json:"productInventoryList"`
		PackageLabel       string `json:"packageLabel"`
		SortingCode        string `json:"sortingCode"`
		ExpressSortingCode string `json:"expressSortingCode"`
		IsCutOffSeller     int    `json:"isCutOffSeller"`
	} `json:"packageWaybillList"`
	OrderCurrency                     string        `json:"orderCurrency"`
	SaleCurrency                      interface{}   `json:"saleCurrency"`
	ProductTotalPrice                 float64       `json:"productTotalPrice"`
	TotalCostPrice                    interface{}   `json:"totalCostPrice"`
	StoreDiscountTotalPrice           float64       `json:"storeDiscountTotalPrice"`
	PromotionDiscountTotalPrice       float64       `json:"promotionDiscountTotalPrice"`
	TotalSettleCurrencyPromotionPrice interface{}   `json:"totalSettleCurrencyPromotionPrice"`
	TotalCommission                   float64       `json:"totalCommission"`
	TotalServiceCharge                interface{}   `json:"totalServiceCharge"`
	TotalPerformanceServiceCharge     interface{}   `json:"totalPerformanceServiceCharge"`
	EstimatedGrossIncome              float64       `json:"estimatedGrossIncome"`
	TotalSaleTax                      float64       `json:"totalSaleTax"`
	TotalSellerRealTax                float64       `json:"totalSellerRealTax"`
	TotalMxTaxPrice                   float64       `json:"totalMxTaxPrice"`
	SellerShippingFee                 float64       `json:"sellerShippingFee"`
	OrderTime                         string        `json:"orderTime"`
	PaymentTime                       string        `json:"paymentTime"`
	OrderAllocateTime                 string        `json:"orderAllocateTime"`
	RequestDeliveryTime               string        `json:"requestDeliveryTime"`
	RequestPrintTime                  interface{}   `json:"requestPrintTime"`
	RequestHandoverTime               interface{}   `json:"requestHandoverTime"`
	SellerDeliveryTime                string        `json:"sellerDeliveryTime"`
	WarehouseDeliveryTime             string        `json:"warehouseDeliveryTime"`
	PrintingTime                      string        `json:"printingTime"`
	ScheduleDeliveryTime              string        `json:"scheduleDeliveryTime"`
	PickUpTime                        string        `json:"pickUpTime"`
	OrderReceiptTime                  string        `json:"orderReceiptTime"`
	OrderRejectionTime                string        `json:"orderRejectionTime"`
	OrderReportedLossTime             string        `json:"orderReportedLossTime"`
	OrderReturnTime                   string        `json:"orderReturnTime"`
	OrderMsgUpdateTime                string        `json:"orderMsgUpdateTime"`
	BillNo                            string        `json:"billNo"`
	SalesArea                         int           `json:"salesArea"`
	StockMode                         int           `json:"stockMode"`
	SalesSite                         string        `json:"salesSite"`
	StoreCode                         int64         `json:"storeCode"`
	DeclareCollectionPattern          int           `json:"declareCollectionPattern"`
	SettleActuallyPrice               float64       `json:"settleActuallyPrice"`
	UnProcessReason                   []interface{} `json:"unProcessReason"`
	PackageInvoiceProblems            []struct {
		ProblemCode        interface{} `json:"problemCode"`
		ProblemDescEnglish string      `json:"problemDescEnglish"`
		ProblemField       interface{} `json:"problemField"`
		ProposalEnglish    interface{} `json:"proposalEnglish"`
		PackageNo          string      `json:"packageNo"`
	} `json:"packageInvoiceProblems"`
	ExpectedCollectTime string `json:"expectedCollectTime"`
}

type ExportAddressRsp struct {
	BaseRsp
	Info struct {
		ReceiveMsgList  []ExportAddressItem `json:"receiveMsgList"`
		UnProcessReason []interface{}       `json:"unProcessReason"`
	} `json:"info"`
}

type ExportAddressItem struct {
	OrderNo    string      `json:"orderNo"`
	LastName   string      `json:"lastName"`
	MiddleName interface{} `json:"middleName"`
	FirstName  string      `json:"firstName"`
	Country    string      `json:"country"`
	Province   string      `json:"province"`
	City       string      `json:"city"`
	District   string      `json:"district"`
	Street     string      `json:"street"`
	Address    string      `json:"address"`
	AddressExt string      `json:"addressExt"`
	Phone      string      `json:"phone"`
	PostCode   string      `json:"postCode"`
	TaxNo      string      `json:"taxNo"`
}

type GetExpressChannelRsp struct {
	BaseRsp
	Info struct {
		ExpressChannels []ExpressChannelItem `json:"expressChannels"`
	} `json:"info"`
}

type ExpressChannelItem struct {
	Site               string  `json:"site"`
	ExpressIdCode      string  `json:"expressIdCode"`
	ExpressChannelCode *string `json:"expressChannelCode"`
}

type ImportBatchMultipleExpressRsp struct {
	BaseRsp
	Info []struct {
		GoodsId       interface{} `json:"goodsId"`
		ExpressCode   string      `json:"expressCode"`
		ExpressIdCode string      `json:"expressIdCode"`
		ErrorMsg      string      `json:"errorMsg"`
		Status        int         `json:"status"`
	} `json:"info"`
}

type ImportBatchExpressItem struct {
	GoodsId       interface{} `json:"goodsId"`
	ExpressCode   string      `json:"expressCode"`
	ExpressIdCode string      `json:"expressIdCode"`
	ErrorMsg      string      `json:"errorMsg"`
	Status        int         `json:"status"`
}

type SyncInvoiceInfoRsp struct {
	BaseRsp
	Info struct {
		SyncInvoiceInfoItem
	} `json:"info"`
}

type SyncInvoiceInfoItem struct {
	OrderInvoiceInfos []struct {
		OrderNo             string  `json:"orderNo"`
		Ie                  string  `json:"ie"`
		Icms                string  `json:"icms"`
		InvoiceNo           string  `json:"invoiceNo"`
		InvoiceKey          string  `json:"invoiceKey"`
		InvoiceSn           string  `json:"invoiceSn"`
		Amount              float64 `json:"amount"`
		TaxNo               string  `json:"taxNo"`
		Currency            string  `json:"currency"`
		InvoiceType         string  `json:"invoiceType"`
		AuthorizationNumber string  `json:"authorizationNumber"`
		AuthorizationTime   string  `json:"authorizationTime"`
		InvoiceIssueTime    string  `json:"invoiceIssueTime"`
		Quantity            string  `json:"quantity"`
		SendMsg             struct {
			Name              string `json:"name"`
			TaxNo             string `json:"taxNo"`
			Ie                string `json:"ie"`
			StateProvinceCode string `json:"stateProvinceCode"`
			CityCode          string `json:"cityCode"`
			City              string `json:"city"`
			Neighborhood      string `json:"neighborhood"`
			Street            string `json:"street"`
			HouseNumber       string `json:"houseNumber"`
			ZipCode           string `json:"zipCode"`
		} `json:"sendMsg"`
		ReceiveMsg struct {
			Name              string `json:"name"`
			TaxNo             string `json:"taxNo"`
			Ie                string `json:"ie"`
			StateProvinceCode string `json:"stateProvinceCode"`
			CityCode          string `json:"cityCode"`
			City              string `json:"city"`
			Neighborhood      string `json:"neighborhood"`
			Street            string `json:"street"`
			HouseNumber       string `json:"houseNumber"`
			ZipCode           string `json:"zipCode"`
		} `json:"receiveMsg"`
	} `json:"orderInvoiceInfos"`
}

type PrintExpressInfoRsp struct {
	BaseRsp
	Info []PrintExpressInfoItem `json:"info"`
}

type PrintExpressInfoItem struct {
	OrderNo    string `json:"orderNo"`
	PackageNo  string `json:"packageNo"`
	FilePdfUrl string `json:"filePdfUrl"`
}

type LogisticsTrackRsp struct {
	BaseRsp
	Info struct {
		TrackInfo []LogisticsTrackItem `json:"trackInfo"`
	} `json:"info"`
}

type LogisticsTrackItem struct {
	WaybillNo   string `json:"waybillNo"`
	Carrier     string `json:"carrier"`
	CarrierCode string `json:"carrierCode"`
	WaybillType int    `json:"waybillType"`
	Tracking    []struct {
		NodeCode         string `json:"nodeCode"`
		NodeCodeName     string `json:"nodeCodeName"`
		Description      string `json:"description"`
		UpdateTimeMillis int64  `json:"updateTimeMillis"`
	} `json:"tracking"`
}

type ConfirmNoStockRsp struct {
	BaseRsp
}

type UnpackingGroupRemoveRsp struct {
	BaseRsp
}

type UnpackingGroupConfirmRsp struct {
	BaseRsp
}
