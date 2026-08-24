package ozon

type FinanceRealizationRsp struct {
	BaseRsp
	Result FinanceRealizationItem `json:"result"`
}

type FinanceRealizationItem struct {
	Header FinanceRealizationHeaderItem `json:"header"`
	Rows   []FinanceRealizationRowItem  `json:"rows"`
}

type FinanceRealizationHeaderItem struct {
	ContractDate    string `json:"contract_date"`
	ContractNumber  string `json:"contract_number"`
	CurrencySysName string `json:"currency_sys_name"`
	DocDate         string `json:"doc_date"`
	Number          string `json:"number"`
	PayerInn        string `json:"payer_inn"`
	PayerKpp        string `json:"payer_kpp"`
	PayerName       string `json:"payer_name"`
	ReceiverInn     string `json:"receiver_inn"`
	ReceiverKpp     string `json:"receiver_kpp"`
	ReceiverName    string `json:"receiver_name"`
	StartDate       string `json:"start_date"`
	StopDate        string `json:"stop_date"`
}

type FinanceRealizationRowItem struct {
	CommissionRatio    float64 `json:"commission_ratio"`
	DeliveryCommission struct {
		Amount           float64 `json:"amount"`
		Bonus            float64 `json:"bonus"`
		Commission       float64 `json:"commission"`
		Compensation     float64 `json:"compensation"`
		PricePerInstance float64 `json:"price_per_instance"`
		Quantity         int     `json:"quantity"`
		StandardFee      float64 `json:"standard_fee"`
		BankCoinvestment float64 `json:"bank_coinvestment"`
		Stars            float64 `json:"stars"`
		Total            float64 `json:"total"`
	} `json:"delivery_commission"`
	Item struct {
		Barcode string `json:"barcode"`
		Name    string `json:"name"`
		OfferId string `json:"offer_id"`
		Sku     int64  `json:"sku"`
	} `json:"item"`
	ReturnCommission struct {
		Amount           float64 `json:"amount"`
		Bonus            float64 `json:"bonus"`
		Commission       float64 `json:"commission"`
		Compensation     float64 `json:"compensation"`
		PricePerInstance float64 `json:"price_per_instance"`
		Quantity         int     `json:"quantity"`
		StandardFee      float64 `json:"standard_fee"`
		BankCoinvestment float64 `json:"bank_coinvestment"`
		Stars            float64 `json:"stars"`
		Total            float64 `json:"total"`
	} `json:"return_commission"`
	RowNumber              int     `json:"rowNumber"`
	SellerPricePerInstance float64 `json:"seller_price_per_instance"`
	Order                  struct {
		PostingNumber string `json:"posting_number"`
		CreatedDate   string `json:"created_date"`
	} `json:"order"`
	LegalEntityDocument struct {
		Number   string `json:"number"`
		SaleDate string `json:"sale_date"`
	} `json:"legal_entity_document"`
}

type FinanceRealizationPostingRsp struct {
	BaseRsp
	Header FinanceRealizationPostingRspHeaderItem `json:"header"`
	Rows   []FinanceRealizationPostingRspRowItem  `json:"rows"`
}

type FinanceRealizationPostingRspHeaderItem struct {
	ContractDate    string `json:"contract_date"`
	ContractNumber  string `json:"contract_number"`
	CurrencySysName string `json:"currency_sys_name"`
	DocDate         string `json:"doc_date"`
	Number          string `json:"number"`
	PayerInn        string `json:"payer_inn"`
	PayerKpp        string `json:"payer_kpp"`
	PayerName       string `json:"payer_name"`
	ReceiverInn     string `json:"receiver_inn"`
	ReceiverKpp     string `json:"receiver_kpp"`
	ReceiverName    string `json:"receiver_name"`
	StartDate       string `json:"start_date"`
	StopDate        string `json:"stop_date"`
}

type FinanceRealizationPostingRspRowItem struct {
	CommissionRatio    float64 `json:"commission_ratio"`
	DeliveryCommission struct {
		Amount           float64 `json:"amount"`
		Bonus            float64 `json:"bonus"`
		Commission       float64 `json:"commission"`
		Compensation     float64 `json:"compensation"`
		PricePerInstance float64 `json:"price_per_instance"`
		Quantity         int     `json:"quantity"`
		StandardFee      float64 `json:"standard_fee"`
		BankCoinvestment float64 `json:"bank_coinvestment"`
		Stars            float64 `json:"stars"`
		Total            float64 `json:"total"`
	} `json:"delivery_commission"`
	Item struct {
		Barcode string `json:"barcode"`
		Name    string `json:"name"`
		OfferId string `json:"offer_id"`
		Sku     int64  `json:"sku"`
	} `json:"item"`
	ReturnCommission struct {
		Amount           float64 `json:"amount"`
		Bonus            float64 `json:"bonus"`
		Commission       float64 `json:"commission"`
		Compensation     float64 `json:"compensation"`
		PricePerInstance float64 `json:"price_per_instance"`
		Quantity         int     `json:"quantity"`
		StandardFee      float64 `json:"standard_fee"`
		BankCoinvestment float64 `json:"bank_coinvestment"`
		Stars            float64 `json:"stars"`
		Total            float64 `json:"total"`
	} `json:"return_commission"`
	RowNumber              int     `json:"row_number"`
	SellerPricePerInstance float64 `json:"seller_price_per_instance"`
	Order                  struct {
		PostingNumber string `json:"posting_number"`
		CreatedDate   string `json:"created_date"`
	} `json:"order"`
	LegalEntityDocument struct {
		Number   string `json:"number"`
		SaleDate string `json:"sale_date"`
	} `json:"legal_entity_document"`
}

type FinanceTransactionListRsp struct {
	BaseRsp
	Result FinanceTransactionListItem `json:"result"`
}

type FinanceTransactionListItem struct {
	Operations []FinanceTransactionListOperationItem `json:"operations"`
	PageCount  int64                                 `json:"page_count"`
	RowCount   int64                                 `json:"row_count"`
}

type FinanceTransactionListOperationItem struct {
	OperationId          int64   `json:"operation_id"`
	OperationType        string  `json:"operation_type"`
	OperationDate        string  `json:"operation_date"`
	OperationTypeName    string  `json:"operation_type_name"`
	DeliveryCharge       float64 `json:"delivery_charge"`
	ReturnDeliveryCharge float64 `json:"return_delivery_charge"`
	AccrualsForSale      float64 `json:"accruals_for_sale"`
	SaleCommission       float64 `json:"sale_commission"`
	Amount               float64 `json:"amount"`
	Type                 string  `json:"type"`
	Posting              struct {
		DeliverySchema string `json:"delivery_schema"`
		OrderDate      string `json:"order_date"`
		PostingNumber  string `json:"posting_number"`
		WarehouseId    int    `json:"warehouse_id"`
	} `json:"posting"`
	Items []struct {
		Name string `json:"name"`
		Sku  int64  `json:"sku"`
	} `json:"items"`
	Services []struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	} `json:"services"`
}

type FinanceTransactionTotalsRsp struct {
	BaseRsp
	Result struct {
		AccrualsForSale         float64 `json:"accruals_for_sale"`
		SaleCommission          float64 `json:"sale_commission"`
		ProcessingAndDelivery   float64 `json:"processing_and_delivery"`
		RefundsAndCancellations float64 `json:"refunds_and_cancellations"`
		ServicesAmount          float64 `json:"services_amount"`
		CompensationAmount      float64 `json:"compensation_amount"`
		MoneyTransfer           float64 `json:"money_transfer"`
		OthersAmount            float64 `json:"others_amount"`
	} `json:"result"`
}

type FinanceCompensationRsp struct {
	BaseRsp
	Result struct {
		Code string `json:"code"`
	} `json:"result"`
}

type FinanceDecompensationRsp struct {
	BaseRsp
	Result struct {
		Code string `json:"code"`
	} `json:"result"`
}

type FinanceBalanceRsp struct {
	BaseRsp
	CashFlows struct {
		Returns struct {
			Amount struct {
				CurrencyCode string  `json:"currency_code"`
				Value        float64 `json:"value"`
			} `json:"amount"`
			AmountDetails struct {
				PartnerPrograms struct {
					CurrencyCode string  `json:"currency_code"`
					Value        float64 `json:"value"`
				} `json:"partner_programs"`
				PointsForDiscounts string `json:"points_for_discounts"`
				Revenue            struct {
					CurrencyCode string  `json:"currency_code"`
					Value        float64 `json:"value"`
				} `json:"revenue"`
			} `json:"amount_details"`
			Fee struct {
				CurrencyCode string  `json:"currency_code"`
				Value        float64 `json:"value"`
			} `json:"fee"`
		} `json:"returns"`
		Sales struct {
			Amount struct {
				CurrencyCode string  `json:"currency_code"`
				Value        float64 `json:"value"`
			} `json:"amount"`
			AmountDetails struct {
				PartnerPrograms struct {
					CurrencyCode string `json:"currency_code"`
					Value        int    `json:"value"`
				} `json:"partner_programs"`
				PointsForDiscounts string `json:"points_for_discounts"`
				Revenue            struct {
					CurrencyCode string  `json:"currency_code"`
					Value        float64 `json:"value"`
				} `json:"revenue"`
			} `json:"amount_details"`
			Fee struct {
				CurrencyCode string  `json:"currency_code"`
				Value        float64 `json:"value"`
			} `json:"fee"`
		} `json:"sales"`
		Services []struct {
			Amount struct {
				CurrencyCode string  `json:"currency_code"`
				Value        float64 `json:"value"`
			} `json:"amount"`
			Name string `json:"name"`
		} `json:"services"`
	} `json:"cashflows"`
	Total struct {
		Accrued struct {
			CurrencyCode string  `json:"currency_code"`
			Value        float64 `json:"value"`
		} `json:"accrued"`
		ClosingBalance struct {
			CurrencyCode string  `json:"currency_code"`
			Value        float64 `json:"value"`
		} `json:"closing_balance"`
		OpeningBalance struct {
			CurrencyCode string  `json:"currency_code"`
			Value        float64 `json:"value"`
		} `json:"opening_balance"`
		Payments []struct {
			CurrencyCode string  `json:"currency_code"`
			Value        float64 `json:"value"`
		} `json:"payments"`
	} `json:"total"`
}

type FinanceAccrualPostingsRsp struct {
	BaseRsp

	PostingAccruals []FinanceAccrualPostingsItem `json:"posting_accruals"`
}

type FinanceAccrualPostingsItem struct {
	Accruals []struct {
		AccrualDate string `json:"accrual_date"`
		Accrued     struct {
			Amount   string `json:"amount"`
			Currency string `json:"currency"`
		} `json:"accrued"`
		Quantity    int `json:"quantity"`
		SellerPrice struct {
			Amount   string `json:"amount"`
			Currency string `json:"currency"`
		} `json:"seller_price"`
		Sku    int64 `json:"sku"`
		TypeId int   `json:"type_id"`
	} `json:"accruals"`
	PostingNumber string `json:"posting_number"`
}

type FinanceAccrualTypesRsp struct {
	BaseRsp

	AccrualTypes []struct {
		Description string `json:"description"`
		Id          int    `json:"id"`
		Name        string `json:"name"`
	} `json:"accrual_types"`
}

type FinanceAccrualByDayRsp struct {
	BaseRsp

	Accruals []FinanceAccrualByDayItem `json:"accruals"`
	LastId   string                    `json:"last_id"`
}

type FinanceAccrualByDayItem struct {
	AccruedCategory string `json:"accrued_category"`
	ContainerFees   struct {
		Fees []struct {
			Accrued struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"accrued"`
			TypeId int `json:"type_id"`
		} `json:"fees"`
	} `json:"container_fees"`
	Date     string `json:"date"`
	ItemFees struct {
		Fees []struct {
			Fees []struct {
				Accrued struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"accrued"`
				TypeId int `json:"type_id"`
			} `json:"fees"`
			Sku int64 `json:"sku"`
		} `json:"fees"`
	} `json:"item_fees"`
	NonItemFee struct {
		Accrued struct {
			Amount   string `json:"amount"`
			Currency string `json:"currency"`
		} `json:"accrued"`
		TypeId int `json:"type_id"`
	} `json:"non_item_fee"`
	Posting struct {
		DeliverySchema string `json:"delivery_schema"`
		DeliverySpeed  int    `json:"delivery_speed"`
		Products       []struct {
			Commission struct {
				Bonus struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"bonus"`
				Coinvestment struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"coinvestment"`
				Commission struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"commission"`
				CommissionRatio string `json:"commission_ratio"`
				SaleAmount      struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"sale_amount"`
				SaleCommission struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"sale_commission"`
				SalePrice struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"sale_price"`
				SellerPrice struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"seller_price"`
			} `json:"commission"`
			Delivery struct {
				Services []struct {
					Accrued struct {
					} `json:"accrued"`
					TypeId int `json:"type_id"`
				} `json:"services"`
				TotalAccrued struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"total_accrued"`
			} `json:"delivery"`
			Sku int64 `json:"sku"`
		} `json:"products"`
	} `json:"posting"`
	TotalAmount struct {
		Amount   string `json:"amount"`
		Currency string `json:"currency"`
	} `json:"total_amount"`
	AccrualId  int64  `json:"accrual_id"`
	UnitNumber string `json:"unit_number"`
}
