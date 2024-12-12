package tiktok

type GetAfterSaleEligibilityRsp struct {
	BaseRsp
	Data struct {
		SkuEligibility []struct {
			LineItemEligibility []struct {
				Eligible          bool     `json:"eligible"`
				IneligibleCode    int      `json:"ineligible_code"`
				IneligibleReason  string   `json:"ineligible_reason"`
				OrderLineItemsIds []string `json:"order_line_items_ids"`
				RequestType       string   `json:"request_type"`
			} `json:"line_item_eligibility"`
			SkuId string `json:"sku_id"`
		} `json:"sku_eligibility"`
	} `json:"data"`
}

type GetRejectReasonsRsp struct {
	BaseRsp
	Data struct {
		Reasons []struct {
			Name string `json:"name"`
			Text string `json:"text"`
		} `json:"reasons"`
	} `json:"data"`
}

type CreateReturnRsp struct {
	BaseRsp
	Data struct {
		ReturnId     string `json:"return_id"`
		ReturnStatus string `json:"return_status"`
	} `json:"data"`
}

type ApproveReturnRsp struct {
	BaseRsp
}

type RejectReturnRsp struct {
	BaseRsp
}

type SearchReturnsRsp struct {
	BaseRsp
	Data struct {
		NextPageToken string `json:"next_page_token"`
		ReturnOrders  []struct {
			ArbitrationStatus          string `json:"arbitration_status"`
			BuyerRejectedPartialRefund bool   `json:"buyer_rejected_partial_refund"`
			CanBuyerKeepItem           bool   `json:"can_buyer_keep_item"`
			CombinedReturnId           string `json:"combined_return_id"`
			CreateTime                 int64  `json:"create_time"`
			DiscountAmount             []struct {
				Currency                    string `json:"currency"`
				ProductPlatformDiscount     string `json:"product_platform_discount"`
				ProductSellerDiscount       string `json:"product_seller_discount"`
				ShippingFeePlatformDiscount string `json:"shipping_fee_platform_discount"`
				ShippingFeeSellerDiscount   string `json:"shipping_fee_seller_discount"`
			} `json:"discount_amount"`
			HandoverMethod   string `json:"handover_method"`
			IsCombinedReturn string `json:"is_combined_return"`
			NextReturnId     string `json:"next_return_id"`
			OrderId          string `json:"order_id"`
			PartialRefund    struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"partial_refund"`
			PreReturnId  string `json:"pre_return_id"`
			RefundAmount struct {
				BuyerServiceFee   string `json:"buyer_service_fee"`
				Currency          string `json:"currency"`
				RefundShippingFee string `json:"refund_shipping_fee"`
				RefundSubtotal    string `json:"refund_subtotal"`
				RefundTax         string `json:"refund_tax"`
				RefundTotal       string `json:"refund_total"`
				RetailDeliveryFee string `json:"retail_delivery_fee"`
			} `json:"refund_amount"`
			ReturnId        string `json:"return_id"`
			ReturnLineItems []struct {
				OrderLineItemId string `json:"order_line_item_id"`
				ProductImage    struct {
					Height int    `json:"height"`
					Url    string `json:"url"`
					Width  int    `json:"width"`
				} `json:"product_image"`
				ProductName  string `json:"product_name"`
				RefundAmount struct {
					BuyerServiceFee   string `json:"buyer_service_fee"`
					Currency          string `json:"currency"`
					RefundShippingFee string `json:"refund_shipping_fee"`
					RefundSubtotal    string `json:"refund_subtotal"`
					RefundTax         string `json:"refund_tax"`
					RefundTotal       string `json:"refund_total"`
					RetailDeliveryFee string `json:"retail_delivery_fee"`
				} `json:"refund_amount"`
				ReturnLineItemId string `json:"return_line_item_id"`
				SellerSku        string `json:"seller_sku"`
				SkuId            string `json:"sku_id"`
				SkuName          string `json:"sku_name"`
			} `json:"return_line_items"`
			ReturnMethod               string `json:"return_method"`
			ReturnProviderId           string `json:"return_provider_id"`
			ReturnProviderName         string `json:"return_provider_name"`
			ReturnReason               string `json:"return_reason"`
			ReturnReasonText           string `json:"return_reason_text"`
			ReturnShippingDocumentType string `json:"return_shipping_document_type"`
			ReturnStatus               string `json:"return_status"`
			ReturnTrackingNumber       string `json:"return_tracking_number"`
			ReturnType                 string `json:"return_type"`
			ReturnWarehouseAddress     struct {
				FullAddress string `json:"full_address"`
			} `json:"return_warehouse_address"`
			Role                     string `json:"role"`
			SellerNextActionResponse []struct {
				Action   string `json:"action"`
				Deadline int64  `json:"deadline"`
			} `json:"seller_next_action_response"`
			SellerProposedReturnType string `json:"seller_proposed_return_type"`
			ShipmentType             string `json:"shipment_type"`
			ShippingFeeAmount        []struct {
				BuyerPaidReturnShippingFee    string `json:"buyer_paid_return_shipping_fee"`
				Currency                      string `json:"currency"`
				PlatformPaidReturnShippingFee string `json:"platform_paid_return_shipping_fee"`
				SellerPaidReturnShippingFee   string `json:"seller_paid_return_shipping_fee"`
			} `json:"shipping_fee_amount"`
			UpdateTime int64 `json:"update_time"`
		} `json:"return_orders"`
		TotalCount int `json:"total_count"`
	} `json:"data"`
}

type GetReturnRecordsRsp struct {
	BaseRsp
	Data struct {
		Records []struct {
			CreateTime  int64  `json:"create_time"`
			Description string `json:"description"`
			Event       string `json:"event"`
			Images      []struct {
				Height int    `json:"height"`
				Url    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"images"`
			Note       string `json:"note"`
			ReasonText string `json:"reason_text"`
			Role       string `json:"role"`
			Videos     []struct {
				Cover          string `json:"cover"`
				DurationMillis int    `json:"duration_millis"`
				Height         int    `json:"height"`
				Url            string `json:"url"`
				Width          int    `json:"width"`
			} `json:"videos"`
		} `json:"records"`
	} `json:"data"`
}

type CancelOrderRsp struct {
	BaseRsp
	Data struct {
		CancelId     string `json:"cancel_id"`
		CancelStatus string `json:"cancel_status"`
	} `json:"data"`
}

type ApproveCancellationRsp struct {
	BaseRsp
}

type RejectCancellationRsp struct {
	BaseRsp
}

type SearchCancellationsRsp struct {
	BaseRsp
	Data struct {
		Cancellations []struct {
			CancelId        string `json:"cancel_id"`
			CancelLineItems []struct {
				CancelLineItemId string `json:"cancel_line_item_id"`
				OrderLineItemId  string `json:"order_line_item_id"`
				ProductImage     struct {
					Height int    `json:"height"`
					Url    string `json:"url"`
					Width  int    `json:"width"`
				} `json:"product_image"`
				ProductName  string `json:"product_name"`
				RefundAmount struct {
					BuyerServiceFee   string `json:"buyer_service_fee"`
					Currency          string `json:"currency"`
					RefundShippingFee string `json:"refund_shipping_fee"`
					RefundSubtotal    string `json:"refund_subtotal"`
					RefundTax         string `json:"refund_tax"`
					RefundTotal       string `json:"refund_total"`
					RetailDeliveryFee string `json:"retail_delivery_fee"`
				} `json:"refund_amount"`
				SellerSku string `json:"seller_sku"`
				SkuId     string `json:"sku_id"`
				SkuName   string `json:"sku_name"`
			} `json:"cancel_line_items"`
			CancelReason     string `json:"cancel_reason"`
			CancelReasonText string `json:"cancel_reason_text"`
			CancelStatus     string `json:"cancel_status"`
			CancelType       string `json:"cancel_type"`
			CreateTime       int64  `json:"create_time"`
			OrderId          string `json:"order_id"`
			RefundAmount     struct {
				BuyerServiceFee   string `json:"buyer_service_fee"`
				Currency          string `json:"currency"`
				RefundShippingFee string `json:"refund_shipping_fee"`
				RefundSubtotal    string `json:"refund_subtotal"`
				RefundTax         string `json:"refund_tax"`
				RefundTotal       string `json:"refund_total"`
				RetailDeliveryFee string `json:"retail_delivery_fee"`
			} `json:"refund_amount"`
			Role                     string `json:"role"`
			SellerNextActionResponse []struct {
				Action   string `json:"action"`
				Deadline int64  `json:"deadline"`
			} `json:"seller_next_action_response"`
			UpdateTime int64 `json:"update_time"`
		} `json:"cancellations"`
		NextPageToken string `json:"next_page_token"`
		TotalCount    int    `json:"total_count"`
	} `json:"data"`
}

type CalculateRefundRsp struct {
	BaseRsp
	Data struct {
		OrderRefundAmount struct {
			Currency          string `json:"currency"`
			RefundShippingFee string `json:"refund_shipping_fee"`
			RefundSubtotal    string `json:"refund_subtotal"`
			RefundTax         string `json:"refund_tax"`
			RefundTotal       string `json:"refund_total"`
			RetailDeliveryFee string `json:"retail_delivery_fee"`
		} `json:"order_refund_amount"`
	} `json:"data"`
}
