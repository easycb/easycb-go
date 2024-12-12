package lazada

type GetDocumentRsp struct {
	Code string `json:"code"`
	Data struct {
		Document struct {
			File         string `json:"file"`
			MimeType     string `json:"mime_type"`
			DocumentType string `json:"document_type"`
		} `json:"document"`
	} `json:"data"`
	RequestId string `json:"request_id"`
}

type GetMultipleOrderItemsRsp struct {
	Code      string                 `json:"code"`
	RequestId string                 `json:"request_id"`
	Data      []GetMultipleOrderItem `json:"data"`
}

type GetMultipleOrderItem struct {
	OrderNumber int64       `json:"order_number"`
	OrderId     int64       `json:"order_id"`
	OrderItems  []OrderItem `json:"order_items"`
}

type OrderItem struct {
	PickUpStoreInfo             interface{} `json:"pick_up_store_info"`
	TaxAmount                   float64     `json:"tax_amount"`
	Reason                      string      `json:"reason"`
	SlaTimeStamp                string      `json:"sla_time_stamp"`
	VoucherSeller               float64     `json:"voucher_seller"`
	PurchaseOrderId             string      `json:"purchase_order_id"`
	PaymentTime                 int64       `json:"payment_time"`
	VoucherCodeSeller           string      `json:"voucher_code_seller"`
	VoucherCode                 string      `json:"voucher_code"`
	PackageId                   string      `json:"package_id"`
	BuyerId                     int64       `json:"buyer_id"`
	Variation                   string      `json:"variation"`
	ProductId                   string      `json:"product_id"`
	VoucherCodePlatform         string      `json:"voucher_code_platform"`
	PurchaseOrderNumber         string      `json:"purchase_order_number"`
	Sku                         string      `json:"sku"`
	GiftWrapping                string      `json:"gift_wrapping"`
	OrderType                   string      `json:"order_type"`
	InvoiceNumber               string      `json:"invoice_number"`
	CancelReturnInitiator       string      `json:"cancel_return_initiator"`
	ShopSku                     string      `json:"shop_sku"`
	IsReroute                   int         `json:"is_reroute"`
	StagePayStatus              string      `json:"stage_pay_status"`
	SkuId                       string      `json:"sku_id"`
	TrackingCodePre             string      `json:"tracking_code_pre"`
	OrderItemId                 int64       `json:"order_item_id"`
	ShopId                      string      `json:"shop_id"`
	OrderFlag                   string      `json:"order_flag"`
	IsFbl                       int         `json:"is_fbl"`
	Name                        string      `json:"name"`
	DeliveryOptionSof           int         `json:"delivery_option_sof"`
	OrderId                     int64       `json:"order_id"`
	FulfillmentSla              string      `json:"fulfillment_sla"`
	Status                      string      `json:"status"`
	ProductMainImage            string      `json:"product_main_image"`
	VoucherPlatform             float64     `json:"voucher_platform"`
	PaidPrice                   float64     `json:"paid_price"`
	ProductDetailUrl            string      `json:"product_detail_url"`
	WarehouseCode               string      `json:"warehouse_code"`
	PromisedShippingTime        string      `json:"promised_shipping_time"`
	ShippingType                string      `json:"shipping_type"`
	CreatedAt                   string      `json:"created_at"`
	VoucherSellerLpi            float64     `json:"voucher_seller_lpi"`
	ShippingFeeDiscountPlatform float64     `json:"shipping_fee_discount_platform"`
	Personalization             string      `json:"personalization"`
	WalletCredits               float64     `json:"wallet_credits"`
	UpdatedAt                   string      `json:"updated_at"`
	Currency                    string      `json:"currency"`
	ShippingProviderType        string      `json:"shipping_provider_type"`
	VoucherPlatformLpi          float64     `json:"voucher_platform_lpi"`
	ShippingFeeOriginal         float64     `json:"shipping_fee_original"`
	ItemPrice                   float64     `json:"item_price"`
	IsDigital                   float64     `json:"is_digital"`
	ShippingServiceCost         float64     `json:"shipping_service_cost"`
	TrackingCode                string      `json:"tracking_code"`
	ShippingFeeDiscountSeller   float64     `json:"shipping_fee_discount_seller"`
	ShippingAmount              float64     `json:"shipping_amount"`
	ReasonDetail                string      `json:"reason_detail"`
	ReturnStatus                string      `json:"return_status"`
	ShipmentProvider            string      `json:"shipment_provider"`
	PriorityFulfillmentTag      string      `json:"priority_fulfillment_tag"`
	VoucherAmount               float64     `json:"voucher_amount"`
	DigitalDeliveryInfo         string      `json:"digital_delivery_info"`
	ExtraAttributes             string      `json:"extra_attributes"`
}

type GetOVOOrdersRsp struct {
	Result struct {
		Success     string `json:"success"`
		ErrorCode   string `json:"errorCode"`
		TradeOrders []struct {
			TradeOrderId    int64  `json:"tradeOrderId"`
			PaymentMethod   string `json:"paymentMethod"`
			PaidTime        string `json:"paidTime"`
			TradeOrderLines []struct {
				DeliveredTime    string `json:"deliveredTime"`
				TradeOrderLineId int64  `json:"tradeOrderLineId"`
				DeliveryStatus   string `json:"deliveryStatus"`
				ReverseStatus    string `json:"reverseStatus"`
			} `json:"tradeOrderLines"`
		} `json:"tradeOrders"`
	} `json:"result"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type GetOrderRsp struct {
	Data      OrderListItem `json:"data"`
	Code      string        `json:"code"`
	RequestId string        `json:"request_id"`
}

type OrderListItem struct {
	VoucherPlatform             string   `json:"voucher_platform"`
	Voucher                     float64  `json:"voucher"`
	WarehouseCode               string   `json:"warehouse_code"`
	OrderNumber                 int64    `json:"order_number"`
	CreatedAt                   string   `json:"created_at"`
	VoucherCode                 string   `json:"voucher_code"`
	GiftOption                  bool     `json:"gift_option"`
	ShippingFeeDiscountPlatform float64  `json:"shipping_fee_discount_platform"`
	CustomerLastName            string   `json:"customer_last_name"`
	UpdatedAt                   string   `json:"updated_at"`
	PromisedShippingTimes       string   `json:"promised_shipping_times"`
	Price                       string   `json:"price"`
	NationalRegistrationNumber  string   `json:"national_registration_number"`
	ShippingFeeOriginal         float64  `json:"shipping_fee_original"`
	PaymentMethod               string   `json:"payment_method"`
	BuyerNote                   string   `json:"buyer_note"`
	CustomerFirstName           string   `json:"customer_first_name"`
	ShippingFeeDiscountSeller   float64  `json:"shipping_fee_discount_seller"`
	ShippingFee                 float64  `json:"shipping_fee"`
	BranchNumber                string   `json:"branch_number"`
	TaxCode                     string   `json:"tax_code"`
	ItemsCount                  int      `json:"items_count"`
	DeliveryInfo                string   `json:"delivery_info"`
	Statuses                    []string `json:"statuses"`
	AddressBilling              struct {
		Country   string `json:"country"`
		Address3  string `json:"address3"`
		Address2  string `json:"address2"`
		City      string `json:"city"`
		Phone     string `json:"phone"`
		Address1  string `json:"address1"`
		PostCode  string `json:"post_code"`
		Phone2    string `json:"phone2"`
		LastName  string `json:"last_name"`
		Address5  string `json:"address5"`
		Address4  string `json:"address4"`
		FirstName string `json:"first_name"`
	} `json:"address_billing"`
	ExtraAttributes string `json:"extra_attributes"`
	OrderId         int64  `json:"order_id"`
	GiftMessage     string `json:"gift_message"`
	Remarks         string `json:"remarks"`
	AddressShipping struct {
		Country   string `json:"country"`
		Address3  string `json:"address3"`
		Address2  string `json:"address2"`
		City      string `json:"city"`
		Phone     string `json:"phone"`
		Address1  string `json:"address1"`
		PostCode  string `json:"post_code"`
		Phone2    string `json:"phone2"`
		LastName  string `json:"last_name"`
		Address5  string `json:"address5"`
		Address4  string `json:"address4"`
		FirstName string `json:"first_name"`
	} `json:"address_shipping"`
	Items interface{} `json:"items"` // 额外的字段, 源数据没有此值, 仅订单上报使用
}

type GetOrderItemsRsp struct {
	Code      string      `json:"code"`
	RequestId string      `json:"request_id"`
	Data      []OrderItem `json:"data"`
}

type GetOrdersRsp struct {
	Code string `json:"code"`
	Data struct {
		Count      int             `json:"count"`
		CountTotal int             `json:"countTotal"`
		Orders     []OrderListItem `json:"orders"`
	} `json:"data"`
	RequestId string `json:"request_id"`
}

type OrderCancelValidateRsp struct {
	Code string `json:"code"`
	Data struct {
		TipContent    string `json:"tip_content"`
		ReasonOptions []struct {
			ReasonName string `json:"reason_name"`
			ReasonId   string `json:"reason_id"`
		} `json:"reason_options"`
		TipType string `json:"tip_type"`
	} `json:"data"`
	RequestId string `json:"request_id"`
}

type SetInvoiceNumberRsp struct {
	Code string `json:"code"`
	Data struct {
		OrderItemId   int64  `json:"order_item_id"`
		InvoiceNumber string `json:"invoice_number"`
	} `json:"data"`
	RequestId string `json:"request_id"`
}

type SetStatusToCanceledRsp struct {
	Code      string `json:"code"`
	Success   bool   `json:"success"`
	RequestId string `json:"request_id"`
}
