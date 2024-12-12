package tiktok

type GetOrderDetailRsp struct {
	BaseRsp
	Data struct {
		Orders []Order `json:"orders"`
	} `json:"data"`
}

type GetOrderListRsp struct {
	BaseRsp
	Data struct {
		NextPageToken string  `json:"next_page_token"`
		Orders        []Order `json:"orders"`
		TotalCount    int     `json:"total_count"`
	} `json:"data"`
}

type Order struct {
	Id                                 string `json:"id"`
	BuyerEmail                         string `json:"buyer_email"`
	BuyerMessage                       string `json:"buyer_message"`
	CancelOrderSlaTime                 int64  `json:"cancel_order_sla_time"`
	CancelReason                       string `json:"cancel_reason"`
	CancelTime                         int64  `json:"cancel_time"`
	CancellationInitiator              string `json:"cancellation_initiator"`
	CollectionDueTime                  int64  `json:"collection_due_time"`
	CollectionTime                     int64  `json:"collection_time"`
	Cpf                                string `json:"cpf"`
	CreateTime                         int64  `json:"create_time"`
	DeliveryDueTime                    int64  `json:"delivery_due_time"`
	DeliveryOptionId                   string `json:"delivery_option_id"`
	DeliveryOptionName                 string `json:"delivery_option_name"`
	DeliveryOptionRequiredDeliveryTime int64  `json:"delivery_option_required_delivery_time"`
	DeliverySlaTime                    int64  `json:"delivery_sla_time"`
	DeliveryTime                       int64  `json:"delivery_time"`
	FulfillmentType                    string `json:"fulfillment_type"`
	HasUpdatedRecipientAddress         bool   `json:"has_updated_recipient_address"`
	IsBuyerRequestCancel               bool   `json:"is_buyer_request_cancel"`
	IsCod                              bool   `json:"is_cod"`
	IsOnHoldOrder                      bool   `json:"is_on_hold_order"`
	IsReplacementOrder                 bool   `json:"is_replacement_order"`
	IsSampleOrder                      bool   `json:"is_sample_order"`
	LineItems                          []struct {
		Id            string `json:"id"`
		CancelReason  string `json:"cancel_reason"`
		CancelUser    string `json:"cancel_user"`
		Currency      string `json:"currency"`
		DisplayStatus string `json:"display_status"`
		IsGift        bool   `json:"is_gift"`
		ItemTax       []struct {
			TaxAmount string `json:"tax_amount"`
			TaxRate   string `json:"tax_rate"`
			TaxType   string `json:"tax_type"`
		} `json:"item_tax"`
		OriginalPrice        string `json:"original_price"`
		PackageId            string `json:"package_id"`
		PackageStatus        string `json:"package_status"`
		PlatformDiscount     string `json:"platform_discount"`
		ProductId            string `json:"product_id"`
		ProductName          string `json:"product_name"`
		RetailDeliveryFee    string `json:"retail_delivery_fee"`
		RtsTime              int64  `json:"rts_time"`
		SalePrice            string `json:"sale_price"`
		SellerDiscount       string `json:"seller_discount"`
		SellerSku            string `json:"seller_sku"`
		ShippingProviderId   string `json:"shipping_provider_id"`
		ShippingProviderName string `json:"shipping_provider_name"`
		SkuId                string `json:"sku_id"`
		SkuImage             string `json:"sku_image"`
		SkuName              string `json:"sku_name"`
		SkuType              string `json:"sku_type"`
		SmallOrderFee        string `json:"small_order_fee"`
		TrackingNumber       string `json:"tracking_number"`
	} `json:"line_items"`
	NeedUploadInvoice string `json:"need_upload_invoice"`
	Packages          []struct {
		Id string `json:"id"`
	} `json:"packages"`
	PaidTime int64 `json:"paid_time"`
	Payment  struct {
		Currency                    string `json:"currency"`
		OriginalShippingFee         string `json:"original_shipping_fee"`
		OriginalTotalProductPrice   string `json:"original_total_product_price"`
		PlatformDiscount            string `json:"platform_discount"`
		ProductTax                  string `json:"product_tax"`
		RetailDeliveryFee           string `json:"retail_delivery_fee"`
		SellerDiscount              string `json:"seller_discount"`
		ShippingFee                 string `json:"shipping_fee"`
		ShippingFeePlatformDiscount string `json:"shipping_fee_platform_discount"`
		ShippingFeeSellerDiscount   string `json:"shipping_fee_seller_discount"`
		ShippingFeeTax              string `json:"shipping_fee_tax"`
		SmallOrderFee               string `json:"small_order_fee"`
		SubTotal                    string `json:"sub_total"`
		Tax                         string `json:"tax"`
		TotalAmount                 string `json:"total_amount"`
	} `json:"payment"`
	PaymentMethodName string `json:"payment_method_name"`
	RecipientAddress  struct {
		AddressDetail       string `json:"address_detail"`
		AddressLine1        string `json:"address_line1"`
		AddressLine2        string `json:"address_line2"`
		AddressLine3        string `json:"address_line3"`
		AddressLine4        string `json:"address_line4"`
		DeliveryPreferences struct {
			DropOffLocation string `json:"drop_off_location"`
		} `json:"delivery_preferences"`
		DistrictInfo []struct {
			AddressLevel     string `json:"address_level"`
			AddressLevelName string `json:"address_level_name"`
			AddressName      string `json:"address_name"`
		} `json:"district_info"`
		FullAddress string `json:"full_address"`
		Name        string `json:"name"`
		PhoneNumber string `json:"phone_number"`
		PostalCode  string `json:"postal_code"`
		RegionCode  string `json:"region_code"`
	} `json:"recipient_address"`
	ReplacedOrderId    string `json:"replaced_order_id"`
	RequestCancelTime  int64  `json:"request_cancel_time"`
	RtsSlaTime         int64  `json:"rts_sla_time"`
	RtsTime            int64  `json:"rts_time"`
	SellerNote         string `json:"seller_note"`
	ShippingDueTime    int64  `json:"shipping_due_time"`
	ShippingProvider   string `json:"shipping_provider"`
	ShippingProviderId string `json:"shipping_provider_id"`
	ShippingType       string `json:"shipping_type"`
	SplitOrCombineTag  string `json:"split_or_combine_tag"`
	Status             string `json:"status"`
	TrackingNumber     string `json:"tracking_number"`
	TtsSlaTime         int64  `json:"tts_sla_time"`
	UpdateTime         int64  `json:"update_time"`
	UserId             string `json:"user_id"`
	WarehouseId        string `json:"warehouse_id"`
}

type GetPriceDetailRsp struct {
	BaseRsp
	Data struct {
		CodFee          string `json:"cod_fee"`
		CodFeeNetAmount string `json:"cod_fee_net_amount"`
		Currency        string `json:"currency"`
		LineItems       []struct {
			CodFee                              string `json:"cod_fee"`
			CodFeeAmount                        string `json:"cod_fee_amount"`
			Currency                            string `json:"currency"`
			Id                                  string `json:"id"`
			NetPriceAmount                      string `json:"net_price_amount"`
			Payment                             string `json:"payment"`
			ShippingFeeDeductionPlatform        string `json:"shipping_fee_deduction_platform"`
			ShippingFeeDeductionPlatformVoucher string `json:"shipping_fee_deduction_platform_voucher"`
			ShippingFeeDeductionSeller          string `json:"shipping_fee_deduction_seller"`
			ShippingListPrice                   string `json:"shipping_list_price"`
			ShippingSalePrice                   string `json:"shipping_sale_price"`
			SkuGiftNetPrice                     string `json:"sku_gift_net_price"`
			SkuGiftOriginalPrice                string `json:"sku_gift_original_price"`
			SkuListPrice                        string `json:"sku_list_price"`
			SkuSalePrice                        string `json:"sku_sale_price"`
			Subtotal                            string `json:"subtotal"`
			SubtotalDeductionPlatform           string `json:"subtotal_deduction_platform"`
			SubtotalDeductionSeller             string `json:"subtotal_deduction_seller"`
			SubtotalTaxAmount                   string `json:"subtotal_tax_amount"`
			TaxAmount                           string `json:"tax_amount"`
			TaxRate                             string `json:"tax_rate"`
			Total                               string `json:"total"`
			VoucherDeductionPlatform            string `json:"voucher_deduction_platform"`
			VoucherDeductionSeller              string `json:"voucher_deduction_seller"`
		} `json:"line_items"`
		NetPriceAmount                      string `json:"net_price_amount"`
		Payment                             string `json:"payment"`
		ShippingFeeDeductionPlatform        string `json:"shipping_fee_deduction_platform"`
		ShippingFeeDeductionPlatformVoucher string `json:"shipping_fee_deduction_platform_voucher"`
		ShippingFeeDeductionSeller          string `json:"shipping_fee_deduction_seller"`
		ShippingListPrice                   string `json:"shipping_list_price"`
		ShippingSalePrice                   string `json:"shipping_sale_price"`
		SkuGiftNetPrice                     string `json:"sku_gift_net_price"`
		SkuGiftOriginalPrice                string `json:"sku_gift_original_price"`
		SkuListPrice                        string `json:"sku_list_price"`
		SkuSalePrice                        string `json:"sku_sale_price"`
		Subtotal                            string `json:"subtotal"`
		SubtotalDeductionPlatform           string `json:"subtotal_deduction_platform"`
		SubtotalDeductionSeller             string `json:"subtotal_deduction_seller"`
		SubtotalTaxAmount                   string `json:"subtotal_tax_amount"`
		TaxAmount                           string `json:"tax_amount"`
		TaxRate                             string `json:"tax_rate"`
		Total                               string `json:"total"`
		VoucherDeductionPlatform            string `json:"voucher_deduction_platform"`
		VoucherDeductionSeller              string `json:"voucher_deduction_seller"`
	} `json:"data"`
}

type AddExternalOrderReferencesRsp struct {
	BaseRsp
	Data struct {
		Errors []struct {
			Code   string `json:"code"`
			Detail struct {
				ExternalOrder struct {
					Id       string `json:"id"`
					Platform string `json:"platform"`
				} `json:"external_order"`
				OrderId string `json:"order_id"`
			} `json:"detail"`
			Message string `json:"message"`
		} `json:"errors"`
	} `json:"data"`
}

type GetExternalOrderReferencesRsp struct {
	BaseRsp
	Data struct {
		ExternalOrders []struct {
			Id        string `json:"id"`
			LineItems []struct {
				Id       string `json:"id"`
				OriginId string `json:"origin_id"`
			} `json:"line_items"`
			Platform string `json:"platform"`
		} `json:"external_orders"`
	} `json:"data"`
}

type SearchOrderByExternalOrderReferenceRsp struct {
	BaseRsp
	Data struct {
		Orders []struct {
			ExternalOrder struct {
				Id        string `json:"id"`
				LineItems []struct {
					Id       string `json:"id"`
					OriginId string `json:"origin_id"`
				} `json:"line_items"`
				Platform string `json:"platform"`
			} `json:"external_order"`
			Id string `json:"id"`
		} `json:"orders"`
	} `json:"data"`
}
