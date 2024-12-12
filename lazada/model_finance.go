package lazada

type GetPayoutStatusRsp struct {
	Code string `json:"code"`
	Data []struct {
		Subtotal2          string `json:"subtotal2"`
		Subtotal1          string `json:"subtotal1"`
		ShipmentFeeCredit  string `json:"shipment_fee_credit"`
		Payout             string `json:"payout"`
		ItemRevenue        string `json:"item_revenue"`
		CreatedAt          string `json:"created_at"`
		OtherRevenueTotal  string `json:"other_revenue_total"`
		FeesTotal          string `json:"fees_total"`
		Refunds            string `json:"refunds"`
		GuaranteeDeposit   string `json:"guarantee_deposit"`
		UpdatedAt          string `json:"updated_at"`
		FeesOnRefundsTotal string `json:"fees_on_refunds_total"`
		ClosingBalance     string `json:"closing_balance"`
		Paid               string `json:"paid"`
		OpeningBalance     string `json:"opening_balance"`
		StatementNumber    string `json:"statement_number"`
		ShipmentFee        string `json:"shipment_fee"`
	} `json:"data"`
	RequestId string `json:"request_id"`
}

type QueryAccountTransactionsRsp struct {
	Msg  string `json:"msg"`
	Code string `json:"code"`
	Data struct {
		PageInfo struct {
			TotalCount int `json:"total_count"`
			TotalPage  int `json:"total_page"`
			PageNum    int `json:"page_num"`
			PageSize   int `json:"page_size"`
		} `json:"page_info"`
		Transactions []struct {
			PmtReference string `json:"pmt_reference"`
			PayeeAccount struct {
				Description string `json:"description"`
				Account     string `json:"account"`
			} `json:"payee_account"`
			Amount            string `json:"amount"`
			SubType           string `json:"sub_type"`
			TransactionNumber string `json:"transaction_number"`
			TransactionTime   string `json:"transaction_time"`
			Currency          string `json:"currency"`
			TrackingList      []struct {
				UpdateTime string `json:"update_time"`
				Name       string `json:"name"`
				Remark     string `json:"remark"`
				Status     string `json:"status"`
			} `json:"tracking_list"`
			Type    string `json:"type"`
			Remarks string `json:"remarks"`
		} `json:"transactions"`
	} `json:"data"`
	Success   bool   `json:"success"`
	ErrorCode string `json:"error_code"`
	RequestId string `json:"request_id"`
}

type QueryLogisticsFeeDetailRsp struct {
	Code string `json:"code"`
	Data []struct {
		TenantId string `json:"tenant_id"`
		Amount   struct {
		} `json:"amount"`
		SkuInfo struct {
			ItemDetails string `json:"item_details"`
			SellerSku   string `json:"seller_sku"`
			LazadaSku   string `json:"lazada_sku"`
		} `json:"sku_info"`
		SellerShortCode string `json:"seller_short_code"`
		TradeOrderId    string `json:"trade_order_id"`
		FeeCreationDate struct {
			Offset struct {
				TotalSeconds int64 `json:"total_seconds"`
				Rules        struct {
					FixedOffset string `json:"fixed_offset"`
				} `json:"rules"`
				Id string `json:"id"`
			} `json:"offset"`
			Year       int `json:"year"`
			DayOfYear  int `json:"day_of_year"`
			Nano       int `json:"nano"`
			Chronology struct {
				CalendarType string `json:"calendar_type"`
				Id           string `json:"id"`
			} `json:"chronology"`
			MonthValue int    `json:"month_value"`
			DayOfMonth int    `json:"day_of_month"`
			Minute     int    `json:"minute"`
			Second     int    `json:"second"`
			Month      string `json:"month"`
			Hour       int    `json:"hour"`
			Zone       struct {
				Rules struct {
					FixedOffset     bool        `json:"fixed_offset"`
					TransitionRules interface{} `json:"transition_rules"`
					Transitions     interface{} `json:"transitions"`
				} `json:"rules"`
				Id string `json:"id"`
			} `json:"zone"`
			DayOfWeek string `json:"day_of_week"`
		} `json:"fee_creation_date"`
		TradeOrderLineId string `json:"trade_order_line_id"`
		StatementId      string `json:"statement_id"`
		OrderInfo        struct {
			OrderItemStatus   string `json:"order_item_status"`
			OrderCreationDate struct {
				Offset     string `json:"offset"`
				Year       int    `json:"year"`
				DayOfYear  int    `json:"day_of_year"`
				Nano       int    `json:"nano"`
				Chronology string `json:"chronology"`
				MonthValue int    `json:"month_value"`
				DayOfMonth int    `json:"day_of_month"`
				Minute     int    `json:"minute"`
				Second     int    `json:"second"`
				Month      string `json:"month"`
				Hour       int    `json:"hour"`
				Zone       struct {
					Rules struct {
						FixedOffset     bool        `json:"fixed_offset"`
						TransitionRules interface{} `json:"transition_rules"`
						Transitions     interface{} `json:"transitions"`
					} `json:"rules"`
					Id string `json:"id"`
				} `json:"zone"`
				DayOfWeek string `json:"day_of_week"`
			} `json:"order_creation_date"`
		} `json:"order_info"`
		FeeName     string `json:"fee_name"`
		FeeCode     string `json:"fee_code"`
		Currency    string `json:"currency"`
		PackageInfo struct {
			DeliveryDate struct {
				Offset     string `json:"offset"`
				Year       int    `json:"year"`
				DayOfYear  int    `json:"day_of_year"`
				Nano       int    `json:"nano"`
				Chronology string `json:"chronology"`
				MonthValue int    `json:"month_value"`
				DayOfMonth int    `json:"day_of_month"`
				Minute     int    `json:"minute"`
				Second     int    `json:"second"`
				Month      string `json:"month"`
				Hour       int    `json:"hour"`
				Zone       struct {
					Rules struct {
						FixedOffset     bool        `json:"fixed_offset"`
						TransitionRules interface{} `json:"transition_rules"`
						Transitions     interface{} `json:"transitions"`
					} `json:"rules"`
					Id string `json:"id"`
				} `json:"zone"`
				DayOfWeek string `json:"day_of_week"`
			} `json:"delivery_date"`
			DestinationAddress string `json:"destination_address"`
			OriginAddress      string `json:"origin_address"`
			TrackingNumber     string `json:"tracking_number"`
			BillingDate        struct {
				Offset     string `json:"offset"`
				Year       string `json:"year"`
				DayOfYear  string `json:"day_of_year"`
				Nano       string `json:"nano"`
				Chronology string `json:"chronology"`
				MonthValue string `json:"month_value"`
				DayOfMonth string `json:"day_of_month"`
				Minute     string `json:"minute"`
				Second     string `json:"second"`
				Month      string `json:"month"`
				Hour       string `json:"hour"`
				Zone       struct {
					Rules struct {
						FixedOffset string `json:"fixed_offset"`
					} `json:"rules"`
					Id string `json:"id"`
				} `json:"zone"`
				DayOfWeek string `json:"day_of_week"`
			} `json:"billing_date"`
			PackageChargeableWeight string `json:"package_chargeable_weight"`
		} `json:"package_info"`
		TaxInAmount struct {
		} `json:"tax_in_amount"`
		SellerId        string `json:"seller_id"`
		StatementPeriod string `json:"statement_period"`
	} `json:"data"`
	Success   bool   `json:"success"`
	Remark    string `json:"remark"`
	RequestId string `json:"request_id"`
}

type QueryTransactionDetailsRsp struct {
	Code string `json:"code"`
	Data []struct {
		OrderNo             string `json:"order_no"`
		TransactionDate     string `json:"transaction_date"`
		Amount              string `json:"amount"`
		PaidStatus          string `json:"paid_status"`
		ShippingProvider    string `json:"shipping_provider"`
		WHTIncludedInAmount string `json:"WHT_included_in_amount"`
		PaymentRefId        string `json:"payment_ref_id"`
		LazadaSku           string `json:"lazada_sku"`
		FeeType             string `json:"fee_type"`
		TransactionType     string `json:"transaction_type"`
		OrderItemNo         string `json:"orderItem_no"`
		OrderItemStatus     string `json:"orderItem_status"`
		Reference           string `json:"reference"`
		FeeName             string `json:"fee_name"`
		ShippingSpeed       string `json:"shipping_speed"`
		WHTAmount           string `json:"WHT_amount"`
		TransactionNumber   string `json:"transaction_number"`
		SellerSku           string `json:"seller_sku"`
		Statement           string `json:"statement"`
		Details             string `json:"details"`
		Comment             string `json:"comment"`
		VATInAmount         string `json:"VAT_in_amount"`
		ShipmentType        string `json:"shipment_type"`
	} `json:"data"`
	RequestId string `json:"request_id"`
}
