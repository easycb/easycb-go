package shopee

type GetOrderListRsp struct {
	BaseRsp
	Response struct {
		More       bool   `json:"more"`
		NextCursor string `json:"next_cursor"`
		OrderList  []struct {
			OrderSn     string `json:"order_sn"`
			OrderStatus string `json:"order_status"`
		} `json:"order_list"`
	} `json:"response"`
}

type GetShipmentListRsp struct {
	BaseRsp
	Response struct {
		OrderList []struct {
			OrderSn       string `json:"order_sn"`
			PackageNumber string `json:"package_number"`
		} `json:"order_list"`
		More       bool   `json:"more"`
		NextCursor string `json:"next_cursor"`
	} `json:"response"`
}

type GetOrderDetailRsp struct {
	BaseRsp
	Response struct {
		OrderList []OrderDetailItem `json:"order_list"`
	} `json:"response"`
}

type OrderDetailItem struct {
	ActualShippingFeeConfirmed bool        `json:"actual_shipping_fee_confirmed"`
	BuyerCancelReason          string      `json:"buyer_cancel_reason"`
	BuyerCpfId                 interface{} `json:"buyer_cpf_id"`
	BuyerUserId                int         `json:"buyer_user_id"`
	BuyerUsername              string      `json:"buyer_username"`
	CancelBy                   string      `json:"cancel_by"`
	CancelReason               string      `json:"cancel_reason"`
	CheckoutShippingCarrier    string      `json:"checkout_shipping_carrier"`
	Cod                        bool        `json:"cod"`
	CreateTime                 int64       `json:"create_time"`
	Currency                   string      `json:"currency"`
	DaysToShip                 int         `json:"days_to_ship"`
	Dropshipper                interface{} `json:"dropshipper"`
	DropshipperPhone           interface{} `json:"dropshipper_phone"`
	EstimatedShippingFee       float64     `json:"estimated_shipping_fee"`
	FulfillmentFlag            string      `json:"fulfillment_flag"`
	GoodsToDeclare             bool        `json:"goods_to_declare"`
	InvoiceData                interface{} `json:"invoice_data"`
	ItemList                   []struct {
		ItemId                 int64   `json:"item_id"`
		ItemName               string  `json:"item_name"`
		ItemSku                string  `json:"item_sku"`
		ModelId                int64   `json:"model_id"`
		ModelName              string  `json:"model_name"`
		ModelSku               string  `json:"model_sku"`
		ModelQuantityPurchased int     `json:"model_quantity_purchased"`
		ModelOriginalPrice     float64 `json:"model_original_price"`
		ModelDiscountedPrice   float64 `json:"model_discounted_price"`
		Wholesale              bool    `json:"wholesale"`
		Weight                 float64 `json:"weight"`
		AddOnDeal              bool    `json:"add_on_deal"`
		MainItem               bool    `json:"main_item"`
		AddOnDealId            int     `json:"add_on_deal_id"`
		PromotionType          string  `json:"promotion_type"`
		PromotionId            int     `json:"promotion_id"`
		OrderItemId            int64   `json:"order_item_id"`
		PromotionGroupId       int     `json:"promotion_group_id"`
		ImageInfo              struct {
			ImageUrl string `json:"image_url"`
		} `json:"image_info"`
		ProductLocationId  []string `json:"product_location_id"`
		IsPrescriptionItem bool     `json:"is_prescription_item"`
		IsB2COwnedItem     bool     `json:"is_b2c_owned_item"`
	} `json:"item_list"`
	MessageToSeller string `json:"message_to_seller"`
	Note            string `json:"note"`
	NoteUpdateTime  int    `json:"note_update_time"`
	OrderSn         string `json:"order_sn"`
	OrderStatus     string `json:"order_status"`
	PackageList     []struct {
		PackageNumber   string      `json:"package_number"`
		GroupShipmentId interface{} `json:"group_shipment_id"`
		LogisticsStatus string      `json:"logistics_status"`
		ShippingCarrier string      `json:"shipping_carrier"`
		ItemList        []struct {
			ItemId            int64  `json:"item_id"`
			ModelId           int64  `json:"model_id"`
			ModelQuantity     int    `json:"model_quantity"`
			OrderItemId       int64  `json:"order_item_id"`
			PromotionGroupId  int    `json:"promotion_group_id"`
			ProductLocationId string `json:"product_location_id"`
		} `json:"item_list"`
		ParcelChargeableWeightGram int `json:"parcel_chargeable_weight_gram"`
	} `json:"package_list"`
	PayTime          int    `json:"pay_time"`
	PaymentMethod    string `json:"payment_method"`
	PickupDoneTime   int    `json:"pickup_done_time"`
	RecipientAddress struct {
		Name        string `json:"name"`
		Phone       string `json:"phone"`
		Town        string `json:"town"`
		District    string `json:"district"`
		City        string `json:"city"`
		State       string `json:"state"`
		Region      string `json:"region"`
		Zipcode     string `json:"zipcode"`
		FullAddress string `json:"full_address"`
	} `json:"recipient_address"`
	Region             string      `json:"region"`
	ReverseShippingFee int         `json:"reverse_shipping_fee"`
	ShipByDate         int         `json:"ship_by_date"`
	ShippingCarrier    string      `json:"shipping_carrier"`
	SplitUp            bool        `json:"split_up"`
	TotalAmount        float64     `json:"total_amount"`
	UpdateTime         int         `json:"update_time"`
	EscrowDetail       interface{} `json:"escrow_detail"` // 额外的字段, 源数据没有此值, 仅订单上报使用
}

type SplitOrderRsp struct {
	BaseRsp
	Response struct {
		OrderSn     string `json:"order_sn"`
		PackageList []struct {
			ItemList []struct {
				ItemId           int64 `json:"item_id"`
				ModelId          int64 `json:"model_id"`
				ModelQuantity    int   `json:"model_quantity"`
				OrderItemId      int64 `json:"order_item_id"`
				PromotionGroupId int   `json:"promotion_group_id"`
			} `json:"item_list"`
			PackageNumber string `json:"package_number"`
		} `json:"package_list"`
	} `json:"response"`
}

type UnSplitOrderRsp struct {
	BaseRsp
}

type CancelOrderRsp struct {
	BaseRsp
	Response struct {
		UpdateTime int64 `json:"update_time"`
	} `json:"response"`
}

type HandleBuyerCancellationRsp struct {
	BaseRsp
	Response struct {
		UpdateTime int64 `json:"update_time"`
	} `json:"response"`
}

type SetNoteRsp struct {
	BaseRsp
}

type GetPendingBuyerInvoiceOrderListRsp struct {
	BaseRsp
	Response struct {
		More       bool   `json:"more"`
		NextCursor string `json:"next_cursor"`
		OrderList  []struct {
			OrderSn string `json:"order_sn"`
		} `json:"order_list"`
	} `json:"response"`
}

type UploadInvoiceDocRsp struct {
	BaseRsp
}

type DownloadInvoiceDocRsp struct {
	BaseRsp
}

type GetBuyerInvoiceInfoRsp struct {
	BaseRsp
	InvoiceInfoList []struct {
		OrderSn       string `json:"order_sn"`
		InvoiceType   string `json:"invoice_type"`
		InvoiceDetail struct {
			Name             string `json:"name"`
			Email            string `json:"email"`
			Address          string `json:"address"`
			PhoneNumber      string `json:"phone_number"`
			TaxId            string `json:"tax_id"`
			AddressBreakdown struct {
				Region          string `json:"region"`
				State           string `json:"state"`
				City            string `json:"city"`
				District        string `json:"district"`
				Town            string `json:"town"`
				Postcode        string `json:"postcode"`
				DetailedAddress string `json:"detailed_address"`
				AdditionalInfo  string `json:"additional_info"`
				FullAddress     string `json:"full_address"`
			} `json:"address_breakdown"`
		} `json:"invoice_detail"`
		IsRequested bool   `json:"is_requested"`
		Error       string `json:"error"`
	} `json:"invoice_info_list"`
}
