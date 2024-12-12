package shopee

type GetEscrowDetailRsp struct {
	BaseRsp
	Response struct {
		BuyerUserName string `json:"buyer_user_name"`
		OrderIncome   struct {
			ActualShippingFee                        float64 `json:"actual_shipping_fee"`
			BuyerPaidShippingFee                     float64 `json:"buyer_paid_shipping_fee"`
			BuyerPaymentMethod                       string  `json:"buyer_payment_method"`
			BuyerTotalAmount                         float64 `json:"buyer_total_amount"`
			BuyerTransactionFee                      float64 `json:"buyer_transaction_fee"`
			CampaignFee                              float64 `json:"campaign_fee"`
			Coins                                    float64 `json:"coins"`
			CommissionFee                            float64 `json:"commission_fee"`
			CostOfGoodsSold                          float64 `json:"cost_of_goods_sold"`
			CreditCardPromotion                      float64 `json:"credit_card_promotion"`
			CreditCardTransactionFee                 float64 `json:"credit_card_transaction_fee"`
			CrossBorderTax                           float64 `json:"cross_border_tax"`
			DeliverySellerProtectionFeePremiumAmount float64 `json:"delivery_seller_protection_fee_premium_amount"`
			DrcAdjustableRefund                      float64 `json:"drc_adjustable_refund"`
			EscrowAmount                             float64 `json:"escrow_amount"`
			EscrowAmountAfterAdjustment              float64 `json:"escrow_amount_after_adjustment"`
			EscrowTax                                float64 `json:"escrow_tax"`
			EstimatedShippingFee                     float64 `json:"estimated_shipping_fee"`
			FinalEscrowProductGst                    float64 `json:"final_escrow_product_gst"`
			FinalEscrowShippingGst                   float64 `json:"final_escrow_shipping_gst"`
			FinalProductProtection                   float64 `json:"final_product_protection"`
			FinalProductVatTax                       float64 `json:"final_product_vat_tax"`
			FinalReturnToSellerShippingFee           float64 `json:"final_return_to_seller_shipping_fee"`
			FinalShippingFee                         float64 `json:"final_shipping_fee"`
			FinalShippingVatTax                      float64 `json:"final_shipping_vat_tax"`
			FsfSellerProtectionFeeClaimAmount        float64 `json:"fsf_seller_protection_fee_claim_amount"`
			InstalmentPlan                           string  `json:"instalment_plan"`
			Items                                    []struct {
				ActivityId                int64   `json:"activity_id"`
				ActivityType              string  `json:"activity_type"`
				AmsCommissionFee          float64 `json:"ams_commission_fee"`
				DiscountFromCoin          float64 `json:"discount_from_coin"`
				DiscountFromVoucherSeller float64 `json:"discount_from_voucher_seller"`
				DiscountFromVoucherShopee float64 `json:"discount_from_voucher_shopee"`
				DiscountedPrice           float64 `json:"discounted_price"`
				IsB2CShopItem             bool    `json:"is_b2c_shop_item"`
				IsMainItem                bool    `json:"is_main_item"`
				ItemId                    int64   `json:"item_id"`
				ItemName                  string  `json:"item_name"`
				ItemSku                   string  `json:"item_sku"`
				ModelId                   int64   `json:"model_id"`
				ModelName                 string  `json:"model_name"`
				ModelSku                  string  `json:"model_sku"`
				OriginalPrice             float64 `json:"original_price"`
				QuantityPurchased         int     `json:"quantity_purchased"`
				SellerDiscount            float64 `json:"seller_discount"`
				SellingPrice              float64 `json:"selling_price"`
				ShopeeDiscount            float64 `json:"shopee_discount"`
			} `json:"items"`
			OrderAmsCommissionFee                  float64       `json:"order_ams_commission_fee"`
			OrderChargeableWeight                  float64       `json:"order_chargeable_weight"`
			OrderDiscountedPrice                   float64       `json:"order_discounted_price"`
			OrderOriginalPrice                     float64       `json:"order_original_price"`
			OrderSellerDiscount                    float64       `json:"order_seller_discount"`
			OrderSellingPrice                      float64       `json:"order_selling_price"`
			OriginalCostOfGoodsSold                float64       `json:"original_cost_of_goods_sold"`
			OriginalPrice                          float64       `json:"original_price"`
			OriginalShopeeDiscount                 float64       `json:"original_shopee_discount"`
			PaymentPromotion                       float64       `json:"payment_promotion"`
			ProratedCoinsValueOffsetReturnItems    int           `json:"prorated_coins_value_offset_return_items"`
			ProratedShopeeVoucherOffsetReturnItems int           `json:"prorated_shopee_voucher_offset_return_items"`
			ReverseShippingFee                     float64       `json:"reverse_shipping_fee"`
			RsfSellerProtectionFeeClaimAmount      float64       `json:"rsf_seller_protection_fee_claim_amount"`
			SalesTaxOnLvg                          float64       `json:"sales_tax_on_lvg"`
			SellerCoinCashBack                     float64       `json:"seller_coin_cash_back"`
			SellerDiscount                         float64       `json:"seller_discount"`
			SellerLostCompensation                 float64       `json:"seller_lost_compensation"`
			SellerReturnRefund                     float64       `json:"seller_return_refund"`
			SellerShippingDiscount                 float64       `json:"seller_shipping_discount"`
			SellerTransactionFee                   float64       `json:"seller_transaction_fee"`
			SellerVoucherCode                      []interface{} `json:"seller_voucher_code"`
			ServiceFee                             float64       `json:"service_fee"`
			ShippingFeeDiscountFrom3Pl             float64       `json:"shipping_fee_discount_from_3pl"`
			ShippingSellerProtectionFeeAmount      float64       `json:"shipping_seller_protection_fee_amount"`
			ShopeeDiscount                         float64       `json:"shopee_discount"`
			ShopeeShippingRebate                   float64       `json:"shopee_shipping_rebate"`
			TotalAdjustmentAmount                  float64       `json:"total_adjustment_amount"`
			VoucherFromSeller                      float64       `json:"voucher_from_seller"`
			VoucherFromShopee                      float64       `json:"voucher_from_shopee"`
		} `json:"order_income"`
		OrderSn           string        `json:"order_sn"`
		ReturnOrderSnList []interface{} `json:"return_order_sn_list"`
	} `json:"response"`
}

type SetShopInstallmentStatusRsp struct {
	BaseRsp
	Response struct {
		InstallmentStatus int `json:"installment_status"`
	} `json:"response"`
}

type GetShopInstallmentStatusRsp struct {
	BaseRsp
	Response struct {
		InstallmentStatus int `json:"installment_status"`
	} `json:"response"`
}

type GetPayoutDetailRsp struct {
	BaseRsp
	Response struct {
		More       bool `json:"more"`
		PayoutList []struct {
			PayoutInfo struct {
				FromCurrency   string  `json:"from_currency"`
				PayoutCurrency string  `json:"payout_currency"`
				FromAmount     float64 `json:"from_amount"`
				PayoutAmount   float64 `json:"payout_amount"`
				ExchangeRate   string  `json:"exchange_rate"`
				PayoutTime     int64   `json:"payout_time"`
				PayService     string  `json:"pay_service"`
				PayeeId        string  `json:"payee_id"`
			} `json:"payout_info"`
			EscrowList []struct {
				EscrowAmount float64 `json:"escrow_amount"`
				Currency     string  `json:"currency"`
				OrderSn      string  `json:"order_sn"`
			} `json:"escrow_list"`
			OfflineAdjustmentList []struct {
				AdjustmentAmount float64 `json:"adjustment_amount"`
				Module           string  `json:"module"`
				Remark           string  `json:"remark"`
				Scenario         string  `json:"scenario"`
				AdjustmentLevel  string  `json:"adjustment_level"`
				OrderSn          string  `json:"order_sn"`
			} `json:"offline_adjustment_list"`
		} `json:"payout_list"`
	} `json:"response"`
}

type SetItemInstallmentStatusRsp struct {
	BaseRsp
	Response struct {
		ItemInstallmentList []struct {
			ItemId      int    `json:"item_id,omitempty"`
			TenureList  []int  `json:"tenure_list,omitempty"`
			FailMessage string `json:"fail_message,omitempty"`
			FailError   string `json:"fail_error,omitempty"`
		} `json:"item_installment_list"`
	} `json:"response"`
}

type GetItemInstallmentStatusRsp struct {
	BaseRsp
	Response struct {
		ItemInstallmentList []struct {
			ItemId      int    `json:"item_id,omitempty"`
			TenureList  []int  `json:"tenure_list,omitempty"`
			FailMessage string `json:"fail_message,omitempty"`
			FailError   string `json:"fail_error,omitempty"`
		} `json:"item_installment_list"`
	} `json:"response"`
}

type GetPaymentMethodListRsp struct {
	BaseRsp
	Response []struct {
		PaymentMethod []string `json:"payment_method"`
		Region        string   `json:"region"`
	} `json:"response"`
}

type GetWalletTransactionListRsp struct {
	BaseRsp
	Response struct {
		More            bool `json:"more"`
		TransactionList []struct {
			Status          string  `json:"status"`
			TransactionType string  `json:"transaction_type"`
			Amount          float64 `json:"amount"`
			CurrentBalance  float64 `json:"current_balance"`
			CreateTime      int64   `json:"create_time"`
			OrderSn         string  `json:"order_sn"`
			RefundSn        string  `json:"refund_sn"`
			WithdrawalType  string  `json:"withdrawal_type"`
			TransactionFee  float64 `json:"transaction_fee"`
			Description     string  `json:"description"`
			BuyerName       string  `json:"buyer_name"`
			PayOrderList    []struct {
				OrderSn  string `json:"order_sn"`
				ShopName string `json:"shop_name"`
			} `json:"pay_order_list"`
			ShopName           string `json:"shop_name"`
			WithdrawalId       int64  `json:"withdrawal_id"`
			Reason             string `json:"reason"`
			RootWithdrawalId   int64  `json:"root_withdrawal_id"`
			TransactionTabType string `json:"transaction_tab_type"`
			MoneyFlow          string `json:"money_flow"`
		} `json:"transaction_list"`
	}
}

type GetEscrowListRsp struct {
	BaseRsp
	Response struct {
		More       bool `json:"more"`
		EscrowList []struct {
			OrderSn           string `json:"order_sn"`
			PayoutAmount      int    `json:"payout_amount"`
			EscrowReleaseTime int    `json:"escrow_release_time"`
		} `json:"escrow_list"`
	} `json:"response"`
}

type GetPayoutInfoRsp struct {
	BaseRsp
	Response struct {
		PayoutList []struct {
			FromCurrency      string  `json:"from_currency"`
			PayoutCurrency    string  `json:"payout_currency"`
			FromAmount        float64 `json:"from_amount"`
			PayoutAmount      float64 `json:"payout_amount"`
			ExchangeRate      string  `json:"exchange_rate"`
			PayoutTime        int64   `json:"payout_time"`
			PayService        string  `json:"pay_service"`
			PayeeId           string  `json:"payee_id"`
			EncryptedPayoutId string  `json:"encrypted_payout_id"`
		} `json:"payout_list"`
		More       bool   `json:"more"`
		NextCursor string `json:"next_cursor"`
	} `json:"response"`
}

type GetBillingTransactionInfoRsp struct {
	BaseRsp
	Response struct {
		Transactions []struct {
			Amount                   float64 `json:"amount"`
			Currency                 string  `json:"currency"`
			OrderSn                  string  `json:"order_sn"`
			CostHeader               string  `json:"cost_header"`
			Scenario                 string  `json:"scenario"`
			Remark                   string  `json:"remark"`
			Level                    string  `json:"level"`
			BillingTransactionType   string  `json:"billing_transaction_type"`
			BillingTransactionStatus string  `json:"billing_transaction_status"`
		} `json:"transactions"`
		More       bool   `json:"more"`
		NextCursor string `json:"next_cursor"`
	} `json:"response"`
}

type GetEscrowDetailBatchRsp struct {
	BaseRsp
	Response []struct {
		EscrowDetail struct {
			BuyerPaymentInfo struct {
				BuyerPaymentMethod   string  `json:"buyer_payment_method"`
				BuyerServiceFee      float64 `json:"buyer_service_fee"`
				BuyerTaxAmount       float64 `json:"buyer_tax_amount"`
				BuyerTotalAmount     float64 `json:"buyer_total_amount"`
				CreditCardPromotion  float64 `json:"credit_card_promotion"`
				IcmsTaxAmount        float64 `json:"icms_tax_amount"`
				ImportTaxAmount      float64 `json:"import_tax_amount"`
				InitialBuyerTxnFee   float64 `json:"initial_buyer_txn_fee"`
				InsurancePremium     float64 `json:"insurance_premium"`
				IofTaxAmount         float64 `json:"iof_tax_amount"`
				IsPaidByCreditCard   bool    `json:"is_paid_by_credit_card"`
				MerchantSubtotal     float64 `json:"merchant_subtotal"`
				SellerVoucher        float64 `json:"seller_voucher"`
				ShippingFee          float64 `json:"shipping_fee"`
				ShippingFeeSstAmount float64 `json:"shipping_fee_sst_amount"`
				ShopeeCoinsRedeemed  float64 `json:"shopee_coins_redeemed"`
				ShopeeVoucher        float64 `json:"shopee_voucher"`
			} `json:"buyer_payment_info"`
			BuyerUserName string `json:"buyer_user_name"`
			OrderIncome   struct {
				ActualShippingFee                        float64 `json:"actual_shipping_fee"`
				BuyerPaidShippingFee                     float64 `json:"buyer_paid_shipping_fee"`
				BuyerPaymentMethod                       string  `json:"buyer_payment_method"`
				BuyerTotalAmount                         float64 `json:"buyer_total_amount"`
				BuyerTransactionFee                      float64 `json:"buyer_transaction_fee"`
				CampaignFee                              float64 `json:"campaign_fee"`
				Coins                                    float64 `json:"coins"`
				CommissionFee                            float64 `json:"commission_fee"`
				CostOfGoodsSold                          float64 `json:"cost_of_goods_sold"`
				CreditCardPromotion                      float64 `json:"credit_card_promotion"`
				CreditCardTransactionFee                 float64 `json:"credit_card_transaction_fee"`
				CrossBorderTax                           float64 `json:"cross_border_tax"`
				DeliverySellerProtectionFeePremiumAmount float64 `json:"delivery_seller_protection_fee_premium_amount"`
				DrcAdjustableRefund                      float64 `json:"drc_adjustable_refund"`
				EscrowAmount                             float64 `json:"escrow_amount"`
				EscrowTax                                float64 `json:"escrow_tax"`
				EstimatedShippingFee                     float64 `json:"estimated_shipping_fee"`
				FinalEscrowProductGst                    float64 `json:"final_escrow_product_gst"`
				FinalEscrowShippingGst                   float64 `json:"final_escrow_shipping_gst"`
				FinalProductProtection                   float64 `json:"final_product_protection"`
				FinalProductVatTax                       float64 `json:"final_product_vat_tax"`
				FinalReturnToSellerShippingFee           float64 `json:"final_return_to_seller_shipping_fee"`
				FinalShippingFee                         float64 `json:"final_shipping_fee"`
				FinalShippingVatTax                      float64 `json:"final_shipping_vat_tax"`
				FsfSellerProtectionFeeClaimAmount        float64 `json:"fsf_seller_protection_fee_claim_amount"`
				InstalmentPlan                           string  `json:"instalment_plan"`
				Items                                    []struct {
					ActivityId                int64   `json:"activity_id"`
					ActivityType              string  `json:"activity_type"`
					AmsCommissionFee          float64 `json:"ams_commission_fee"`
					DiscountFromCoin          float64 `json:"discount_from_coin"`
					DiscountFromVoucherSeller float64 `json:"discount_from_voucher_seller"`
					DiscountFromVoucherShopee float64 `json:"discount_from_voucher_shopee"`
					DiscountedPrice           float64 `json:"discounted_price"`
					IsB2CShopItem             bool    `json:"is_b2c_shop_item"`
					IsMainItem                bool    `json:"is_main_item"`
					ItemId                    int64   `json:"item_id"`
					ItemName                  string  `json:"item_name"`
					ItemSku                   string  `json:"item_sku"`
					ModelId                   int64   `json:"model_id"`
					ModelName                 string  `json:"model_name"`
					ModelSku                  string  `json:"model_sku"`
					OriginalPrice             float64 `json:"original_price"`
					QuantityPurchased         float64 `json:"quantity_purchased"`
					SellerDiscount            float64 `json:"seller_discount"`
					SellingPrice              float64 `json:"selling_price"`
					ShopeeDiscount            float64 `json:"shopee_discount"`
				} `json:"items"`
				OrderAmsCommissionFee                              float64       `json:"order_ams_commission_fee"`
				OrderChargeableWeight                              float64       `json:"order_chargeable_weight"`
				OrderDiscountedPrice                               float64       `json:"order_discounted_price"`
				OrderOriginalPrice                                 float64       `json:"order_original_price"`
				OrderSellerDiscount                                float64       `json:"order_seller_discount"`
				OrderSellingPrice                                  float64       `json:"order_selling_price"`
				OriginalCostOfGoodsSold                            float64       `json:"original_cost_of_goods_sold"`
				OriginalPrice                                      float64       `json:"original_price"`
				OriginalShopeeDiscount                             float64       `json:"original_shopee_discount"`
				OverseasReturnServiceFee                           float64       `json:"overseas_return_service_fee"`
				PaymentPromotion                                   float64       `json:"payment_promotion"`
				ProratedCoinsValueOffsetReturnItems                float64       `json:"prorated_coins_value_offset_return_items"`
				ProratedPaymentChannelPromoBankOffsetReturnItems   float64       `json:"prorated_payment_channel_promo_bank_offset_return_items"`
				ProratedPaymentChannelPromoShopeeOffsetReturnItems float64       `json:"prorated_payment_channel_promo_shopee_offset_return_items"`
				ProratedSellerVoucherOffsetReturnItems             float64       `json:"prorated_seller_voucher_offset_return_items"`
				ProratedShopeeVoucherOffsetReturnItems             float64       `json:"prorated_shopee_voucher_offset_return_items"`
				ReverseShippingFee                                 float64       `json:"reverse_shipping_fee"`
				ReverseShippingFeeSst                              float64       `json:"reverse_shipping_fee_sst"`
				RsfSellerProtectionFeeClaimAmount                  float64       `json:"rsf_seller_protection_fee_claim_amount"`
				SalesTaxOnLvg                                      float64       `json:"sales_tax_on_lvg"`
				SellerCoinCashBack                                 float64       `json:"seller_coin_cash_back"`
				SellerDiscount                                     float64       `json:"seller_discount"`
				SellerLostCompensation                             float64       `json:"seller_lost_compensation"`
				SellerReturnRefund                                 float64       `json:"seller_return_refund"`
				SellerShippingDiscount                             float64       `json:"seller_shipping_discount"`
				SellerTransactionFee                               float64       `json:"seller_transaction_fee"`
				SellerVoucherCode                                  []interface{} `json:"seller_voucher_code"`
				ServiceFee                                         float64       `json:"service_fee"`
				ShippingFeeDiscountFrom3Pl                         float64       `json:"shipping_fee_discount_from_3pl"`
				ShippingFeeSst                                     float64       `json:"shipping_fee_sst"`
				ShippingSellerProtectionFeeAmount                  float64       `json:"shipping_seller_protection_fee_amount"`
				ShopeeDiscount                                     float64       `json:"shopee_discount"`
				ShopeeShippingRebate                               float64       `json:"shopee_shipping_rebate"`
				VatOnImportedGoods                                 float64       `json:"vat_on_imported_goods"`
				VoucherFromSeller                                  float64       `json:"voucher_from_seller"`
				VoucherFromShopee                                  float64       `json:"voucher_from_shopee"`
				WithholdingTax                                     float64       `json:"withholding_tax"`
			} `json:"order_income"`
			OrderSn           string   `json:"order_sn"`
			ReturnOrderSnList []string `json:"return_order_sn_list"`
		} `json:"escrow_detail"`
	} `json:"response"`
}
