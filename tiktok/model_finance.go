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

type GetStatementTransactionsV202501Rsp struct {
	BaseRsp
	Data struct {
		CreateTime               int64  `json:"create_time"`
		Currency                 string `json:"currency"`
		Id                       string `json:"id"`
		NextPageToken            string `json:"next_page_token"`
		PayableAmount            string `json:"payable_amount"`
		Status                   string `json:"status"`
		TotalCount               int    `json:"total_count"`
		TotalReserveAmount       string `json:"total_reserve_amount"`
		TotalSettlementAmount    string `json:"total_settlement_amount"`
		TotalSettlementBreakdown struct {
			TotalAdjustmentAmount   string `json:"total_adjustment_amount"`
			TotalFeeTaxAmount       string `json:"total_fee_tax_amount"`
			TotalRevenueAmount      string `json:"total_revenue_amount"`
			TotalShippingCostAmount string `json:"total_shipping_cost_amount"`
		} `json:"total_settlement_breakdown"`
		Transactions []struct {
			AdjustmentAmount     string `json:"adjustment_amount"`
			AdjustmentId         string `json:"adjustment_id"`
			AdjustmentOrderId    string `json:"adjustment_order_id"`
			AssociatedOrderId    string `json:"associated_order_id"`
			EstimatedReleaseTime string `json:"estimated_release_time"`
			FeeTaxAmount         string `json:"fee_tax_amount"`
			FeeTaxBreakdown      struct {
				Fee struct {
					AffiliateAdsCommissionAmount       string `json:"affiliate_ads_commission_amount"`
					AffiliateCommissionAmount          string `json:"affiliate_commission_amount"`
					AffiliateCommissionAmountBeforePit string `json:"affiliate_commission_amount_before_pit"`
					AffiliatePartnerCommissionAmount   string `json:"affiliate_partner_commission_amount"`
					BonusCashbackServiceFeeAmount      string `json:"bonus_cashback_service_fee_amount"`
					CreditCardHandlingFeeAmount        string `json:"credit_card_handling_fee_amount"`
					LiveSpecialsFeeAmount              string `json:"live_specials_fee_amount"`
					MallServiceFeeAmount               string `json:"mall_service_fee_amount"`
					PlatformCommissionAmount           string `json:"platform_commission_amount"`
					ReferralFeeAmount                  string `json:"referral_fee_amount"`
					RefundAdministrationFeeAmount      string `json:"refund_administration_fee_amount"`
					SfpServiceFeeAmount                string `json:"sfp_service_fee_amount"`
					TransactionFeeAmount               string `json:"transaction_fee_amount"`
				} `json:"fee"`
				Tax struct {
					CustomsClearanceAmount string `json:"customs_clearance_amount"`
					CustomsDutyAmount      string `json:"customs_duty_amount"`
					GstAmount              string `json:"gst_amount"`
					ImportVatAmount        string `json:"import_vat_amount"`
					SstAmount              string `json:"sst_amount"`
					VatAmount              string `json:"vat_amount"`
				} `json:"tax"`
			} `json:"fee_tax_breakdown"`
			Id               string `json:"id"`
			OrderCreateTime  int64  `json:"order_create_time"`
			OrderId          string `json:"order_id"`
			ReserveAmount    string `json:"reserve_amount"`
			ReserveId        string `json:"reserve_id"`
			ReserveStatus    string `json:"reserve_status"`
			RevenueAmount    string `json:"revenue_amount"`
			RevenueBreakdown struct {
				CodServiceFeeAmount                string `json:"cod_service_fee_amount"`
				RefundCodServiceFeeAmount          string `json:"refund_cod_service_fee_amount"`
				RefundSubtotalBeforeDiscountAmount string `json:"refund_subtotal_before_discount_amount"`
				SellerDiscountAmount               string `json:"seller_discount_amount"`
				SellerDiscountRefundAmount         string `json:"seller_discount_refund_amount"`
				SubtotalBeforeDiscountAmount       string `json:"subtotal_before_discount_amount"`
			} `json:"revenue_breakdown"`
			SettlementAmount      string `json:"settlement_amount"`
			ShippingCostAmount    string `json:"shipping_cost_amount"`
			ShippingCostBreakdown struct {
				ActualShippingFeeAmount        string `json:"actual_shipping_fee_amount"`
				CustomerPaidShippingFeeAmount  string `json:"customer_paid_shipping_fee_amount"`
				ExchangeShippingFeeAmount      string `json:"exchange_shipping_fee_amount"`
				ReplacementShippingFeeAmount   string `json:"replacement_shipping_fee_amount"`
				ReturnShippingFeeAmount        string `json:"return_shipping_fee_amount"`
				ShippingFeeDiscountAmount      string `json:"shipping_fee_discount_amount"`
				ShippingInsuranceFeeAmount     string `json:"shipping_insurance_fee_amount"`
				SignatureConfirmationFeeAmount string `json:"signature_confirmation_fee_amount"`
				SupplementaryComponent         struct {
					CustomerShippingFeeOffsetAmount   string `json:"customer_shipping_fee_offset_amount"`
					FbmShippingCostAmount             string `json:"fbm_shipping_cost_amount"`
					FbtFulfillmentFeeAmount           string `json:"fbt_fulfillment_fee_amount"`
					FbtShippingCostAmount             string `json:"fbt_shipping_cost_amount"`
					PlatformShippingFeeDiscountAmount string `json:"platform_shipping_fee_discount_amount"`
					PromoShippingIncentiveAmount      string `json:"promo_shipping_incentive_amount"`
					SellerShippingFeeDiscountAmount   string `json:"seller_shipping_fee_discount_amount"`
					ShippingFeeSubsidyAmount          string `json:"shipping_fee_subsidy_amount"`
				} `json:"supplementary_component"`
			} `json:"shipping_cost_breakdown"`
			SupplementaryComponent struct {
				CustomerPaymentAmount                string `json:"customer_payment_amount"`
				CustomerRefundAmount                 string `json:"customer_refund_amount"`
				PlatformCofundedDiscountAmount       string `json:"platform_cofunded_discount_amount"`
				PlatformCofundedDiscountRefundAmount string `json:"platform_cofunded_discount_refund_amount"`
				PlatformDiscountAmount               string `json:"platform_discount_amount"`
				PlatformDiscountRefundAmount         string `json:"platform_discount_refund_amount"`
				RetailDeliveryFeeAmount              string `json:"retail_delivery_fee_amount"`
				RetailDeliveryFeePaymentAmount       string `json:"retail_delivery_fee_payment_amount"`
				RetailDeliveryFeeRefundAmount        string `json:"retail_delivery_fee_refund_amount"`
				SalesTaxAmount                       string `json:"sales_tax_amount"`
				SalesTaxPaymentAmount                string `json:"sales_tax_payment_amount"`
				SalesTaxRefundAmount                 string `json:"sales_tax_refund_amount"`
				SellerCofundedDiscountAmount         string `json:"seller_cofunded_discount_amount"`
				SellerCofundedDiscountRefundAmount   string `json:"seller_cofunded_discount_refund_amount"`
			} `json:"supplementary_component"`
			Type string `json:"type"`
		} `json:"transactions"`
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

type GetOrderStatementTransactionsV202501Rsp struct {
	BaseRsp
	Data struct {
		Currency           string `json:"currency"`
		FeeAndTaxAmount    string `json:"fee_and_tax_amount"`
		OrderCreateTime    int64  `json:"order_create_time"`
		OrderId            string `json:"order_id"`
		RevenueAmount      string `json:"revenue_amount"`
		SettlementAmount   string `json:"settlement_amount"`
		ShippingCostAmount string `json:"shipping_cost_amount"`
		SkuTransactions    []struct {
			FeeTaxAmount    string `json:"fee_tax_amount"`
			FeeTaxBreakdown struct {
				Fee struct {
					AffiliateAdsCommissionAmount       string `json:"affiliate_ads_commission_amount"`
					AffiliateCommissionAmount          string `json:"affiliate_commission_amount"`
					AffiliateCommissionAmountBeforePit string `json:"affiliate_commission_amount_before_pit"`
					AffiliatePartnerCommissionAmount   string `json:"affiliate_partner_commission_amount"`
					BonusCashbackServiceFeeAmount      string `json:"bonus_cashback_service_fee_amount"`
					CreditCardHandlingFeeAmount        string `json:"credit_card_handling_fee_amount"`
					LiveSpecialsFeeAmount              string `json:"live_specials_fee_amount"`
					MallServiceFeeAmount               string `json:"mall_service_fee_amount"`
					PlatformCommissionAmount           string `json:"platform_commission_amount"`
					ReferralFeeAmount                  string `json:"referral_fee_amount"`
					RefundAdministrationFeeAmount      string `json:"refund_administration_fee_amount"`
					SfpServiceFeeAmount                string `json:"sfp_service_fee_amount"`
					TransactionFeeAmount               string `json:"transaction_fee_amount"`
				} `json:"fee"`
				Tax struct {
					CustomsClearanceAmount string `json:"customs_clearance_amount"`
					CustomsDutyAmount      string `json:"customs_duty_amount"`
					GstAmount              string `json:"gst_amount"`
					ImportVatAmount        string `json:"import_vat_amount"`
					SstAmount              string `json:"sst_amount"`
					VatAmount              string `json:"vat_amount"`
				} `json:"tax"`
			} `json:"fee_tax_breakdown"`
			ProductName      string `json:"product_name"`
			Quantity         string `json:"quantity"`
			RevenueAmount    string `json:"revenue_amount"`
			RevenueBreakdown struct {
				CodServiceFeeAmount                string `json:"cod_service_fee_amount"`
				RefundCodServiceFeeAmount          string `json:"refund_cod_service_fee_amount"`
				RefundSubtotalBeforeDiscountAmount string `json:"refund_subtotal_before_discount_amount"`
				SellerDiscountAmount               string `json:"seller_discount_amount"`
				SellerDiscountRefundAmount         string `json:"seller_discount_refund_amount"`
				SubtotalBeforeDiscountAmount       string `json:"subtotal_before_discount_amount"`
			} `json:"revenue_breakdown"`
			SettlementAmount      string `json:"settlement_amount"`
			ShippingCostAmount    string `json:"shipping_cost_amount"`
			ShippingCostBreakdown struct {
				ActualShippingFeeAmount        string `json:"actual_shipping_fee_amount"`
				CustomerPaidShippingFeeAmount  string `json:"customer_paid_shipping_fee_amount"`
				ExchangeShippingFeeAmount      string `json:"exchange_shipping_fee_amount"`
				ReplacementShippingFeeAmount   string `json:"replacement_shipping_fee_amount"`
				ReturnShippingFeeAmount        string `json:"return_shipping_fee_amount"`
				ShippingFeeDiscountAmount      string `json:"shipping_fee_discount_amount"`
				ShippingInsuranceFeeAmount     string `json:"shipping_insurance_fee_amount"`
				SignatureConfirmationFeeAmount string `json:"signature_confirmation_fee_amount"`
				SupplementaryComponent         struct {
					CustomerShippingFeeOffsetAmount   string `json:"customer_shipping_fee_offset_amount"`
					FbmShippingCostAmount             string `json:"fbm_shipping_cost_amount"`
					FbtFulfillmentFeeAmount           string `json:"fbt_fulfillment_fee_amount"`
					FbtShippingCostAmount             string `json:"fbt_shipping_cost_amount"`
					PlatformShippingFeeDiscountAmount string `json:"platform_shipping_fee_discount_amount"`
					PromoShippingIncentiveAmount      string `json:"promo_shipping_incentive_amount"`
					SellerShippingFeeDiscountAmount   string `json:"seller_shipping_fee_discount_amount"`
					ShippingFeeSubsidyAmount          string `json:"shipping_fee_subsidy_amount"`
				} `json:"supplementary_component"`
			} `json:"shipping_cost_breakdown"`
			SkuId   string `json:"sku_id"`
			SkuName string `json:"sku_name"`
		} `json:"sku_transactions"`
		TotalCount int `json:"total_count"`
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

type GetUnsettledTransactionsRsp struct {
	BaseRsp
	Data struct {
		NextPageToken          string                 `json:"next_page_token"`
		TotalCount             int                    `json:"total_count"`
		SumEstSettlementAmount string                 `json:"sum_est_settlement_amount"`
		SumEstRevenueAmount    string                 `json:"sum_est_revenue_amount"`
		SumEstAdjustmentAmount string                 `json:"sum_est_adjustment_amount"`
		SumEstFeeAmount        string                 `json:"sum_est_fee_amount"`
		Transactions           []UnsettledTransaction `json:"transactions"`
	} `json:"data"`
}

type UnsettledTransaction struct {
	Type                string `json:"type"`
	Id                  string `json:"id"`
	Status              string `json:"status"`
	Currency            string `json:"currency"`
	EstimatedSettlement string `json:"estimated_settlement"`
	UnsettledReason     string `json:"unsettled_reason"`
	OrderCreateTime     int64  `json:"order_create_time"`
	OrderDeliveryTime   int64  `json:"order_delivery_time"`
	OrderId             string `json:"order_id"`
	AdjustmentId        string `json:"adjustment_id"`
	AdjustmentOrderId   string `json:"adjustment_order_id"`
	EstAdjustmentAmount string `json:"est_adjustment_amount"`
	EstSettlementAmount string `json:"est_settlement_amount"`
	EstRevenueAmount    string `json:"est_revenue_amount"`
	RevenueBreakdown    struct {
		SubtotalBeforeDiscountAmount       string `json:"subtotal_before_discount_amount"`
		SellerDiscountAmount               string `json:"seller_discount_amount"`
		RefundSubtotalBeforeDiscountAmount string `json:"refund_subtotal_before_discount_amount"`
		SellerDiscountRefundAmount         string `json:"seller_discount_refund_amount"`
		CodServiceFeeAmount                string `json:"cod_service_fee_amount"`
		RefundCodServiceFeeAmount          string `json:"refund_cod_service_fee_amount"`
	} `json:"revenue_breakdown"`
	EstShippingCostAmount string `json:"est_shipping_cost_amount"`
	ShippingCostBreakdown struct {
		ActualShippingFeeAmount        string `json:"actual_shipping_fee_amount"`
		ShippingFeeDiscountAmount      string `json:"shipping_fee_discount_amount"`
		CustomerPaidShippingFeeAmount  string `json:"customer_paid_shipping_fee_amount"`
		ReturnShippingFeeAmount        string `json:"return_shipping_fee_amount"`
		ReplacementShippingFeeAmount   string `json:"replacement_shipping_fee_amount"`
		ExchangeShippingFeeAmount      string `json:"exchange_shipping_fee_amount"`
		SignatureConfirmationFeeAmount string `json:"signature_confirmation_fee_amount"`
		ShippingInsuranceFeeAmount     string `json:"shipping_insurance_fee_amount"`
		SupplementaryComponent         struct {
			CustomerShippingFeeOffsetAmount   string `json:"customer_shipping_fee_offset_amount"`
			ShippingFeeSubsidyAmount          string `json:"shipping_fee_subsidy_amount"`
			PlatformShippingFeeDiscountAmount string `json:"platform_shipping_fee_discount_amount"`
			FbmShippingCostAmount             string `json:"fbm_shipping_cost_amount"`
			FbtShippingCostAmount             string `json:"fbt_shipping_cost_amount"`
			PromoShippingIncentiveAmount      string `json:"promo_shipping_incentive_amount"`
			FbtFulfillmentFeeAmount           string `json:"fbt_fulfillment_fee_amount"`
			SellerShippingFeeDiscountAmount   string `json:"seller_shipping_fee_discount_amount"`
		} `json:"supplementary_component"`
	} `json:"shipping_cost_breakdown"`
	EstFeeTaxAmount string `json:"est_fee_tax_amount"`
	FeeTaxBreakdown struct {
		Fee struct {
			PlatformCommissionAmount           string `json:"platform_commission_amount"`
			ReferralFeeAmount                  string `json:"referral_fee_amount"`
			RefundAdministrationFeeAmount      string `json:"refund_administration_fee_amount"`
			TransactionFeeAmount               string `json:"transaction_fee_amount"`
			CreditCardHandlingFeeAmount        string `json:"credit_card_handling_fee_amount"`
			AffiliateCommissionAmount          string `json:"affiliate_commission_amount"`
			AffiliateCommissionBeforePitAmount string `json:"affiliate_commission_before_pit_amount"`
			PitWithheldFromAdsCommissionAmount string `json:"pit_withheld_from_ads_commission_amount"`
			AffiliatePartnerCommissionAmount   string `json:"affiliate_partner_commission_amount"`
			AffiliateAdsCommissionAmount       string `json:"affiliate_ads_commission_amount"`
			SfpServiceFeeAmount                string `json:"sfp_service_fee_amount"`
			LiveSpecialsFeeAmount              string `json:"live_specials_fee_amount"`
			BonusCashbackServiceFeeAmount      string `json:"bonus_cashback_service_fee_amount"`
			MallServiceFeeAmount               string `json:"mall_service_fee_amount"`
			RetailDeliveryFeeAmount            string `json:"retail_delivery_fee_amount"`
			RetailDeliveryFeePaymentAmount     string `json:"retail_delivery_fee_payment_amount"`
			RetailDeliveryFeeRefundAmount      string `json:"retail_delivery_fee_refund_amount"`
		} `json:"fee"`
		Tax struct {
			SalesTaxAmount         string `json:"sales_tax_amount"`
			SalesTaxPaymentAmount  string `json:"sales_tax_payment_amount"`
			SalesTaxRefundAmount   string `json:"sales_tax_refund_amount"`
			VatAmount              string `json:"vat_amount"`
			ImportVatAmount        string `json:"import_vat_amount"`
			CustomsDutyAmount      string `json:"customs_duty_amount"`
			CustomsClearanceAmount string `json:"customs_clearance_amount"`
			SstAmount              string `json:"sst_amount"`
			GstAmount              string `json:"gst_amount"`
		} `json:"tax"`
	} `json:"fee_tax_breakdown"`
}
