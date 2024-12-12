package lazada

type GetReverseOrderDetailRsp struct {
	Code string `json:"code"`
	Data struct {
		ReverseOrderId          int64  `json:"reverse_order_id"`
		RequestType             string `json:"request_type"`
		ReverseOrderLineDTOList []struct {
			ReturnOrderLineGmtCreate   int64   `json:"return_order_line_gmt_create"`
			PlatformSkuId              string  `json:"platform_sku_id"`
			IsNeedRefund               string  `json:"is_need_refund"`
			TradeOrderGmtCreate        int64   `json:"trade_order_gmt_create"`
			ReasonText                 string  `json:"reason_text"`
			ItemUnitPrice              float64 `json:"item_unit_price"`
			TradeOrderLineId           int64   `json:"trade_order_line_id"`
			ReturnOrderLineGmtModified int64   `json:"return_order_line_gmt_modified"`
			OfcStatus                  string  `json:"ofc_status"`
			SellerSkuId                string  `json:"seller_sku_id"`
			ProductDTO                 struct {
				ProductId int64  `json:"product_id"`
				Sku       string `json:"sku"`
			} `json:"productDTO"`
			RefundPaymentMethod string `json:"refund_payment_method"`
			Buyer               struct {
				UserId int `json:"user_id"`
			} `json:"buyer"`
			ReasonCode         int     `json:"reason_code"`
			WhqcDecision       string  `json:"whqc_decision"`
			ReverseStatus      string  `json:"reverse_status"`
			RefundAmount       float64 `json:"refund_amount"`
			TrackingNumber     string  `json:"tracking_number"`
			IsDispute          bool    `json:"is_dispute"`
			ReverseOrderLineId int64   `json:"reverse_order_line_id"`
		} `json:"reverseOrderLineDTOList"`
		ShippingType string `json:"shipping_type"`
		IsRtm        bool   `json:"is_rtm"`
		TradeOrderId int64  `json:"trade_order_id"`
	} `json:"data"`
	RequestId string `json:"request_id"`
}

type GetReverseOrderHistoryListRsp struct {
	Code string `json:"code"`
	Data struct {
		PageInfo struct {
			Total             int `json:"total"`
			PageSize          int `json:"page_size"`
			CurrentPageNumber int `json:"current_page_number"`
		} `json:"page_info"`
		List []struct {
			Time     int64         `json:"time"`
			Operator string        `json:"operator"`
			Picture  []interface{} `json:"picture"`
		} `json:"list"`
	} `json:"data"`
	RequestId string `json:"request_id"`
}

type GetReverseOrderReasonListRsp struct {
	Code string `json:"code"`
	Data []struct {
		MutiLanguageText string `json:"muti_language_text"`
		Text             string `json:"text"`
		ReasonId         int    `json:"reason_id"`
	} `json:"data"`
	RequestId string `json:"request_id"`
}

type GetReverseOrdersForSellerRsp struct {
	Result struct {
		Total   int  `json:"total"`
		Success bool `json:"success"`
		PageNo  int  `json:"page_no"`
		Items   []struct {
			ReverseOrderLines []struct {
				Product struct {
					ProductSku string `json:"product_sku"`
					ProductId  string `json:"product_id"`
				} `json:"product"`
				ReturnOrderLineGmtCreate   int64   `json:"return_order_line_gmt_create"`
				PlatformSkuId              string  `json:"platform_sku_id"`
				IsNeedRefund               string  `json:"is_need_refund"`
				TradeOrderGmtCreate        int64   `json:"trade_order_gmt_create"`
				ReasonText                 string  `json:"reason_text"`
				ItemUnitPrice              float64 `json:"item_unit_price"`
				TradeOrderLineId           int64   `json:"trade_order_line_id"`
				ReturnOrderLineGmtModified int64   `json:"return_order_line_gmt_modified"`
				OfcStatus                  string  `json:"ofc_status"`
				SellerSkuId                string  `json:"seller_sku_id"`
				RefundPaymentMethod        string  `json:"refund_payment_method"`
				Buyer                      struct {
					BuyerId int `json:"buyer_id"`
				} `json:"buyer"`
				ReasonCode         int     `json:"reason_code"`
				WhqcDecision       string  `json:"whqc_decision"`
				ReverseStatus      string  `json:"reverse_status"`
				RefundAmount       float64 `json:"refund_amount"`
				TrackingNumber     string  `json:"tracking_number"`
				ReceiverAddress    string  `json:"receiver_address"`
				IsDispute          bool    `json:"is_dispute"`
				ReverseOrderLineId int64   `json:"reverse_order_line_id"`
			} `json:"reverse_order_lines"`
			ReverseOrderId string `json:"reverse_order_id"`
			RequestType    string `json:"request_type"`
			IsRtm          bool   `json:"is_rtm"`
			ShippingType   string `json:"shipping_type"`
			TradeOrderId   int64  `json:"trade_order_id"`
		} `json:"items"`
		PageSize string `json:"page_size"`
	} `json:"result"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type InitReverseOrderCancelRsp struct {
	Code string `json:"code"`
	Data struct {
		TipContent string `json:"tip_content"`
		TipType    string `json:"tip_type"`
	} `json:"data"`
	RequestId string `json:"request_id"`
}

type InitReverseOrderCancelDecideRsp struct {
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type ReverseOrderReturnUpdateRsp struct {
	Code string `json:"code"`
	Data struct {
		ReasonInfo []struct {
			ReasonName string `json:"reason_name"`
			ReasonId   int    `json:"reason_id"`
		} `json:"reason_info"`
		ReverseOrderId   int64  `json:"reverse_order_id"`
		TotalRefund      string `json:"total_refund"`
		ReverseOrderLine []struct {
			PaidPrice          float64 `json:"paid_price"`
			IsCancel           bool    `json:"is_cancel"`
			ReasonId           int     `json:"reason_id"`
			ReasonSource       string  `json:"reason_source"`
			ReasonDesc         string  `json:"reason_desc"`
			ApplyReason        string  `json:"apply_reason"`
			ReasonType         string  `json:"reason_type"`
			SellerSku          string  `json:"seller_sku"`
			RefundAmount       float64 `json:"refund_amount"`
			OrderLineId        int64   `json:"order_line_id"`
			ReasonName         string  `json:"reason_name"`
			OrderId            int64   `json:"order_id"`
			ReverseOrderLineId int64   `json:"reverse_order_line_id"`
		} `json:"reverse_order_line"`
	} `json:"data"`
	RequestId string `json:"request_id"`
}
