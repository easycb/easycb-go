package tiktok

type GetStatementsRsp struct {
	BaseRsp
	Data struct {
		NextPageToken string `json:"next_page_token"`
		Statements    []struct {
			AdjustmentAmount   string `json:"adjustment_amount"`
			Currency           string `json:"currency"`
			FeeAmount          string `json:"fee_amount"`
			Id                 string `json:"id"`
			NetSalesAmount     string `json:"net_sales_amount"`
			PaymentId          string `json:"payment_id"`
			PaymentStatus      string `json:"payment_status"`
			RevenueAmount      string `json:"revenue_amount"`
			SettlementAmount   string `json:"settlement_amount"`
			ShippingCostAmount string `json:"shipping_cost_amount"`
			StatementTime      int64  `json:"statement_time"`
		} `json:"statements"`
	} `json:"data"`
}

type GetStatementTransactionsRsp struct {
	BaseRsp
	Data struct {
		AdjustmentAmount      string `json:"adjustment_amount"`
		Currency              string `json:"currency"`
		FeeAmount             string `json:"fee_amount"`
		NetSalesAmount        string `json:"net_sales_amount"`
		NextPageToken         string `json:"next_page_token"`
		RevenueAmount         string `json:"revenue_amount"`
		SettlementAmount      string `json:"settlement_amount"`
		ShippingCostAmount    string `json:"shipping_cost_amount"`
		StatementId           string `json:"statement_id"`
		StatementTime         int64  `json:"statement_time"`
		StatementTransactions []struct {
			ActualReturnShippingFeeAmount       string `json:"actual_return_shipping_fee_amount"`
			ActualShippingFeeAmount             string `json:"actual_shipping_fee_amount"`
			AdjustmentAmount                    string `json:"adjustment_amount"`
			AdjustmentId                        string `json:"adjustment_id"`
			AdjustmentOrderId                   string `json:"adjustment_order_id"`
			AffiliateAdsCommissionAmount        string `json:"affiliate_ads_commission_amount"`
			AffiliateCommissionAmount           string `json:"affiliate_commission_amount"`
			AffiliateCommissionBeforePit        string `json:"affiliate_commission_before_pit"`
			AffiliatePartnerCommissionAmount    string `json:"affiliate_partner_commission_amount"`
			AfterSellerDiscountsSubtotalAmount  string `json:"after_seller_discounts_subtotal_amount"`
			Currency                            string `json:"currency"`
			CustomerOrderRefundAmount           string `json:"customer_order_refund_amount"`
			CustomerPaidShippingFeeAmount       string `json:"customer_paid_shipping_fee_amount"`
			CustomerPaidShippingFeeRefundAmount string `json:"customer_paid_shipping_fee_refund_amount"`
			CustomerPaymentAmount               string `json:"customer_payment_amount"`
			CustomerRefundAmount                string `json:"customer_refund_amount"`
			CustomerShippingFeeAmount           string `json:"customer_shipping_fee_amount"`
			CustomerShippingFeeOffsetAmount     string `json:"customer_shipping_fee_offset_amount"`
			FbmShippingCostAmount               string `json:"fbm_shipping_cost_amount"`
			FbtFulfillmentFeeAmount             string `json:"fbt_fulfillment_fee_amount"`
			FbtShippingCostAmount               string `json:"fbt_shipping_cost_amount"`
			FeeAmount                           string `json:"fee_amount"`
			GrossSalesAmount                    string `json:"gross_sales_amount"`
			GrossSalesRefundAmount              string `json:"gross_sales_refund_amount"`
			Id                                  string `json:"id"`
			NetSalesAmount                      string `json:"net_sales_amount"`
			OrderCreateTime                     int64  `json:"order_create_time"`
			OrderId                             string `json:"order_id"`
			PitAmount                           string `json:"pit_amount"`
			PlatformCommissionAmount            string `json:"platform_commission_amount"`
			PlatformDiscountAmount              string `json:"platform_discount_amount"`
			PlatformDiscountRefundAmount        string `json:"platform_discount_refund_amount"`
			PlatformRefundSubsidyAmount         string `json:"platform_refund_subsidy_amount"`
			PlatformShippingFeeDiscountAmount   string `json:"platform_shipping_fee_discount_amount"`
			PromoShippingIncentiveAmount        string `json:"promo_shipping_incentive_amount"`
			ReferralFeeAmount                   string `json:"referral_fee_amount"`
			RefundAdministrationFeeAmount       string `json:"refund_administration_fee_amount"`
			RefundShippingCostDiscountAmount    string `json:"refund_shipping_cost_discount_amount"`
			RetailDeliveryFeeAmount             string `json:"retail_delivery_fee_amount"`
			RetailDeliveryFeePaymentAmount      string `json:"retail_delivery_fee_payment_amount"`
			RetailDeliveryFeeRefundAmount       string `json:"retail_delivery_fee_refund_amount"`
			ReturnShippingFeeAmount             string `json:"return_shipping_fee_amount"`
			RevenueAmount                       string `json:"revenue_amount"`
			SalesTaxAmount                      string `json:"sales_tax_amount"`
			SalesTaxPaymentAmount               string `json:"sales_tax_payment_amount"`
			SalesTaxRefundAmount                string `json:"sales_tax_refund_amount"`
			SellerDiscountAmount                string `json:"seller_discount_amount"`
			SellerDiscountRefundAmount          string `json:"seller_discount_refund_amount"`
			SettlementAmount                    string `json:"settlement_amount"`
			ShippingCostAmount                  string `json:"shipping_cost_amount"`
			ShippingCostDiscountAmount          string `json:"shipping_cost_discount_amount"`
			ShippingFeeAmount                   string `json:"shipping_fee_amount"`
			ShippingFeeSubsidyAmount            string `json:"shipping_fee_subsidy_amount"`
			ShippingInsuranceFeeAmount          string `json:"shipping_insurance_fee_amount"`
			SignatureConfirmationFeeAmount      string `json:"signature_confirmation_fee_amount"`
			TransactionFeeAmount                string `json:"transaction_fee_amount"`
			Type                                string `json:"type"`
		} `json:"statement_transactions"`
		TotalCount int `json:"total_count"`
	} `json:"data"`
}

type GetOrderStatementTransactionsRsp struct {
	BaseRsp
	Data struct {
		OrderCreateTime       int64  `json:"order_create_time"`
		OrderId               string `json:"order_id"`
		StatementTransactions []struct {
			ActualReturnShippingFeeAmount       string `json:"actual_return_shipping_fee_amount"`
			ActualShippingFeeAmount             string `json:"actual_shipping_fee_amount"`
			AdjustmentAmount                    string `json:"adjustment_amount"`
			AffiliateAdsCommissionAmount        string `json:"affiliate_ads_commission_amount"`
			AffiliateCommissionAmount           string `json:"affiliate_commission_amount"`
			AffiliateCommissionBeforePit        string `json:"affiliate_commission_before_pit"`
			AffiliatePartnerCommissionAmount    string `json:"affiliate_partner_commission_amount"`
			AfterSellerDiscountsSubtotalAmount  string `json:"after_seller_discounts_subtotal_amount"`
			Currency                            string `json:"currency"`
			CustomerOrderRefundAmount           string `json:"customer_order_refund_amount"`
			CustomerPaidShippingFeeAmount       string `json:"customer_paid_shipping_fee_amount"`
			CustomerPaidShippingFeeRefundAmount string `json:"customer_paid_shipping_fee_refund_amount"`
			CustomerPaymentAmount               string `json:"customer_payment_amount"`
			CustomerRefundAmount                string `json:"customer_refund_amount"`
			CustomerShippingFeeAmount           string `json:"customer_shipping_fee_amount"`
			CustomerShippingFeeOffsetAmount     string `json:"customer_shipping_fee_offset_amount"`
			FbmShippingCostAmount               string `json:"fbm_shipping_cost_amount"`
			FbtFulfillmentFeeAmount             string `json:"fbt_fulfillment_fee_amount"`
			FbtShippingCostAmount               string `json:"fbt_shipping_cost_amount"`
			FeeAmount                           string `json:"fee_amount"`
			GrossSalesAmount                    string `json:"gross_sales_amount"`
			GrossSalesRefundAmount              string `json:"gross_sales_refund_amount"`
			Id                                  string `json:"id"`
			NetSalesAmount                      string `json:"net_sales_amount"`
			PitAmount                           string `json:"pit_amount"`
			PlatformCommissionAmount            string `json:"platform_commission_amount"`
			PlatformDiscountAmount              string `json:"platform_discount_amount"`
			PlatformDiscountRefundAmount        string `json:"platform_discount_refund_amount"`
			PlatformRefundSubsidyAmount         string `json:"platform_refund_subsidy_amount"`
			PlatformShippingFeeDiscountAmount   string `json:"platform_shipping_fee_discount_amount"`
			PromoShippingIncentiveAmount        string `json:"promo_shipping_incentive_amount"`
			ReferralFeeAmount                   string `json:"referral_fee_amount"`
			RefundAdministrationFeeAmount       string `json:"refund_administration_fee_amount"`
			RefundShippingCostDiscountAmount    string `json:"refund_shipping_cost_discount_amount"`
			RetailDeliveryFeeAmount             string `json:"retail_delivery_fee_amount"`
			RetailDeliveryFeePaymentAmount      string `json:"retail_delivery_fee_payment_amount"`
			RetailDeliveryFeeRefundAmount       string `json:"retail_delivery_fee_refund_amount"`
			ReturnShippingFeeAmount             string `json:"return_shipping_fee_amount"`
			RevenueAmount                       string `json:"revenue_amount"`
			SalesTaxAmount                      string `json:"sales_tax_amount"`
			SalesTaxPaymentAmount               string `json:"sales_tax_payment_amount"`
			SalesTaxRefundAmount                string `json:"sales_tax_refund_amount"`
			SellerDiscountAmount                string `json:"seller_discount_amount"`
			SellerDiscountRefundAmount          string `json:"seller_discount_refund_amount"`
			SettlementAmount                    string `json:"settlement_amount"`
			ShippingCostAmount                  string `json:"shipping_cost_amount"`
			ShippingCostDiscountAmount          string `json:"shipping_cost_discount_amount"`
			ShippingFeeAmount                   string `json:"shipping_fee_amount"`
			ShippingFeeSubsidyAmount            string `json:"shipping_fee_subsidy_amount"`
			ShippingInsuranceFeeAmount          string `json:"shipping_insurance_fee_amount"`
			SignatureConfirmationFeeAmount      string `json:"signature_confirmation_fee_amount"`
			SkuStatementTransactions            []struct {
				ActualReturnShippingFeeAmount       string `json:"actual_return_shipping_fee_amount"`
				ActualShippingFeeAmount             string `json:"actual_shipping_fee_amount"`
				AdjustmentAmount                    string `json:"adjustment_amount"`
				AffiliateAdsCommissionAmount        string `json:"affiliate_ads_commission_amount"`
				AffiliateCommissionAmount           string `json:"affiliate_commission_amount"`
				AffiliateCommissionBeforePit        string `json:"affiliate_commission_before_pit"`
				AffiliatePartnerCommissionAmount    string `json:"affiliate_partner_commission_amount"`
				AfterSellerDiscountsSubtotalAmount  string `json:"after_seller_discounts_subtotal_amount"`
				Currency                            string `json:"currency"`
				CustomerOrderRefundAmount           string `json:"customer_order_refund_amount"`
				CustomerPaidShippingFeeAmount       string `json:"customer_paid_shipping_fee_amount"`
				CustomerPaidShippingFeeRefundAmount string `json:"customer_paid_shipping_fee_refund_amount"`
				CustomerPaymentAmount               string `json:"customer_payment_amount"`
				CustomerRefundAmount                string `json:"customer_refund_amount"`
				CustomerShippingFeeAmount           string `json:"customer_shipping_fee_amount"`
				CustomerShippingFeeOffsetAmount     string `json:"customer_shipping_fee_offset_amount"`
				FbmShippingCostAmount               string `json:"fbm_shipping_cost_amount"`
				FbtFulfillmentFeeAmount             string `json:"fbt_fulfillment_fee_amount"`
				FbtShippingCostAmount               string `json:"fbt_shipping_cost_amount"`
				FeeAmount                           string `json:"fee_amount"`
				GrossSalesAmount                    string `json:"gross_sales_amount"`
				GrossSalesRefundAmount              string `json:"gross_sales_refund_amount"`
				NetSalesAmount                      string `json:"net_sales_amount"`
				PitAmount                           string `json:"pit_amount"`
				PlatformCommissionAmount            string `json:"platform_commission_amount"`
				PlatformDiscountAmount              string `json:"platform_discount_amount"`
				PlatformDiscountRefundAmount        string `json:"platform_discount_refund_amount"`
				PlatformRefundSubsidyAmount         string `json:"platform_refund_subsidy_amount"`
				PlatformShippingFeeDiscountAmount   string `json:"platform_shipping_fee_discount_amount"`
				ProductName                         string `json:"product_name"`
				PromoShippingIncentiveAmount        string `json:"promo_shipping_incentive_amount"`
				Quantity                            int    `json:"quantity"`
				ReferralFeeAmount                   string `json:"referral_fee_amount"`
				RefundAdministrationFeeAmount       string `json:"refund_administration_fee_amount"`
				RefundShippingCostDiscountAmount    string `json:"refund_shipping_cost_discount_amount"`
				RetailDeliveryFeeAmount             string `json:"retail_delivery_fee_amount"`
				RetailDeliveryFeePaymentAmount      string `json:"retail_delivery_fee_payment_amount"`
				RetailDeliveryFeeRefundAmount       string `json:"retail_delivery_fee_refund_amount"`
				ReturnShippingFeeAmount             string `json:"return_shipping_fee_amount"`
				RevenueAmount                       string `json:"revenue_amount"`
				SalesTaxAmount                      string `json:"sales_tax_amount"`
				SalesTaxPaymentAmount               string `json:"sales_tax_payment_amount"`
				SalesTaxRefundAmount                string `json:"sales_tax_refund_amount"`
				SellerDiscountAmount                string `json:"seller_discount_amount"`
				SellerDiscountRefundAmount          string `json:"seller_discount_refund_amount"`
				SettlementAmount                    string `json:"settlement_amount"`
				ShippingCostAmount                  string `json:"shipping_cost_amount"`
				ShippingCostDiscountAmount          string `json:"shipping_cost_discount_amount"`
				ShippingFeeAmount                   string `json:"shipping_fee_amount"`
				ShippingFeeSubsidyAmount            string `json:"shipping_fee_subsidy_amount"`
				ShippingInsuranceFeeAmount          string `json:"shipping_insurance_fee_amount"`
				SignatureConfirmationFeeAmount      string `json:"signature_confirmation_fee_amount"`
				SkuId                               string `json:"sku_id"`
				SkuName                             string `json:"sku_name"`
				TransactionFeeAmount                string `json:"transaction_fee_amount"`
			} `json:"sku_statement_transactions"`
			StatementId          string `json:"statement_id"`
			StatementTime        int64  `json:"statement_time"`
			Status               string `json:"status"`
			TransactionFeeAmount string `json:"transaction_fee_amount"`
		} `json:"statement_transactions"`
	} `json:"data"`
}

type GetPaymentsRsp struct {
	BaseRsp
	Data struct {
		NextPageToken string `json:"next_page_token"`
		Payments      []struct {
			Amount struct {
				Currency string `json:"currency"`
				Value    string `json:"value"`
			} `json:"amount"`
			BankAccount                 string `json:"bank_account"`
			CreateTime                  int64  `json:"create_time"`
			ExchangeRate                string `json:"exchange_rate"`
			Id                          string `json:"id"`
			PaidTime                    int64  `json:"paid_time"`
			PaymentAmountBeforeExchange struct {
				Currency string `json:"currency"`
				Value    string `json:"value"`
			} `json:"payment_amount_before_exchange"`
			ReserveAmount struct {
				Currency string `json:"currency"`
				Value    string `json:"value"`
			} `json:"reserve_amount"`
			SettlementAmount struct {
				Currency string `json:"currency"`
				Value    string `json:"value"`
			} `json:"settlement_amount"`
			Status string `json:"status"`
		} `json:"payments"`
	} `json:"data"`
}

type GetWithdrawalsRsp struct {
	BaseRsp
	Data struct {
		NextPageToken string `json:"next_page_token"`
		TotalCount    int    `json:"total_count"`
		Withdrawals   []struct {
			Amount     string `json:"amount"`
			CreateTime int    `json:"create_time"`
			Currency   string `json:"currency"`
			Id         string `json:"id"`
			Status     string `json:"status"`
			Type       string `json:"type"`
		} `json:"withdrawals"`
	} `json:"data"`
}
