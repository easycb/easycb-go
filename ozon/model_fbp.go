package ozon

import "time"

type PostingFbpGetRsp struct {
	BaseRsp

	Posting struct {
		AnalyticsData struct {
			City              string    `json:"city"`
			DeliveryDateBegin time.Time `json:"delivery_date_begin"`
			DeliveryDateEnd   time.Time `json:"delivery_date_end"`
			DeliveryType      string    `json:"delivery_type"`
			Region            string    `json:"region"`
			WarehouseId       int64     `json:"warehouse_id"`
		} `json:"analytics_data"`
		Cancellation struct {
			CancelReason          string `json:"cancel_reason"`
			CancelReasonId        int64  `json:"cancel_reason_id"`
			CancellationInitiator string `json:"cancellation_initiator"`
			CancellationType      string `json:"cancellation_type"`
		} `json:"cancellation"`
		FinancialData struct {
			ClusterFrom    string  `json:"cluster_from"`
			ClusterTo      string  `json:"cluster_to"`
			DeliveryAmount float64 `json:"delivery_amount"`
			Products       []struct {
				Actions []struct {
					ActionId        int64     `json:"action_id"`
					ActionType      string    `json:"action_type"`
					DateFrom        time.Time `json:"date_from"`
					DateTo          time.Time `json:"date_to"`
					Description     string    `json:"description"`
					DiscountPercent float64   `json:"discount_percent"`
					DiscountValue   float64   `json:"discount_value"`
				} `json:"actions"`
				CommissionsPrice struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"commissions_price"`
				CustomerPrice struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"customer_price"`
				OldPrice          float64 `json:"old_price"`
				PostingCommission struct {
					Amount  float64 `json:"amount"`
					Payout  float64 `json:"payout"`
					Percent float64 `json:"percent"`
				} `json:"posting_commission"`
				Quantity         int `json:"quantity"`
				ReturnCommission struct {
					Amount  float64 `json:"amount"`
					Payout  float64 `json:"payout"`
					Percent float64 `json:"percent"`
				} `json:"return_commission"`
				SellerPrice struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"seller_price"`
				Sku                  int64   `json:"sku"`
				TotalDiscountPercent float64 `json:"total_discount_percent"`
				TotalDiscountValue   float64 `json:"total_discount_value"`
			} `json:"products"`
		} `json:"financial_data"`
		InProcessAt   time.Time `json:"in_process_at"`
		OrderDate     time.Time `json:"order_date"`
		OrderId       int64     `json:"order_id"`
		OrderNumber   string    `json:"order_number"`
		PostingNumber string    `json:"posting_number"`
		Products      []struct {
			HasImei                bool `json:"has_imei"`
			MarketplaceSellerPrice struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"marketplace_seller_price"`
			Name      string  `json:"name"`
			OfferId   string  `json:"offer_id"`
			Quantity  int     `json:"quantity"`
			Sku       int64   `json:"sku"`
			WeightMax float64 `json:"weight_max"`
		} `json:"products"`
		Status        int64  `json:"status"`
		Substatus     string `json:"substatus"`
		TplProviderId int64  `json:"tpl_provider_id"`
	} `json:"posting"`
}

type PostingFbpListRsp struct {
	BaseRsp

	Cursor   string               `json:"cursor"`
	Postings []PostingFbpListItem `json:"postings"`
}

type PostingFbpListItem struct {
	FinancialData struct {
		ClusterFrom    string  `json:"cluster_from"`
		ClusterTo      string  `json:"cluster_to"`
		DeliveryAmount float64 `json:"delivery_amount"`
		Products       []struct {
			Actions []struct {
				ActionId        string    `json:"action_id"`
				DateFrom        time.Time `json:"date_from"`
				DateTo          time.Time `json:"date_to"`
				DiscountPercent float64   `json:"discount_percent"`
				DiscountValue   float64   `json:"discount_value"`
				IsFromSeller    bool      `json:"is_from_seller"`
				Description     string    `json:"description"`
			} `json:"actions"`
			CommissionsCurrencyCode string  `json:"commissions_currency_code"`
			OldPrice                float64 `json:"old_price"`
			Price                   float64 `json:"price"`
			ProductId               int64   `json:"product_id"`
			Quantity                int     `json:"quantity"`
			TotalDiscountPercent    float64 `json:"total_discount_percent"`
			PostingCommission       struct {
				Amount  float64 `json:"amount"`
				Payout  float64 `json:"payout"`
				Percent float64 `json:"percent"`
			} `json:"posting_commission"`
			ReturnCommission struct {
				Amount  float64 `json:"amount"`
				Payout  float64 `json:"payout"`
				Percent float64 `json:"percent"`
			} `json:"return_commission"`
			TotalDiscountValue int `json:"total_discount_value"`
		} `json:"products"`
	} `json:"financial_data"`
	InProcessAt   time.Time `json:"in_process_at"`
	OrderDate     time.Time `json:"order_date"`
	OrderId       int64     `json:"order_id"`
	OrderNumber   string    `json:"order_number"`
	PostingNumber string    `json:"posting_number"`
	Products      []struct {
		CustomerPrice struct {
			Amount   string `json:"amount"`
			Currency string `json:"currency"`
		} `json:"customer_price"`
		Name    string `json:"name"`
		OfferId string `json:"offer_id"`
		Price   struct {
			Amount   string `json:"amount"`
			Currency string `json:"currency"`
		} `json:"price"`
		Quantity    int `json:"quantity"`
		SellerPrice struct {
			Amount   string `json:"amount"`
			Currency string `json:"currency"`
		} `json:"seller_price"`
		Sku int64 `json:"sku"`
	} `json:"products"`
	ProviderId int64  `json:"provider_id"`
	Status     string `json:"status"`
}

type FbpActFromCreateRsp struct {
	BaseRsp

	Errors    []string `json:"errors"`
	FileUuid  string   `json:"file_uuid"`
	IsSuccess bool     `json:"is_success"`
}

type FbpActFromGetRsp struct {
	BaseRsp

	CdnUrl string `json:"cdn_url"`
	Error  string `json:"error"`
	Status string `json:"status"`
}

type FbpActToCreateRsp struct {
	BaseRsp

	Code string `json:"code"`
}

type FbpActToGetRsp struct {
	BaseRsp

	ErrorMessage string `json:"error_message"`
	LabelUrl     string `json:"label_url"`
	State        string `json:"state"`
}

type FbpArchiveGetRsp struct {
	BaseRsp

	ActFileUuid      string `json:"act_file_uuid"`
	BundleId         string `json:"bundle_id"`
	BundleSkuSummary struct {
		RoundedTotalVolumeInLitres float64 `json:"rounded_total_volume_in_litres"`
		TotalItemsCount            int     `json:"total_items_count"`
		TotalQuantity              int     `json:"total_quantity"`
	} `json:"bundle_sku_summary"`
	BusinessFlowTypeId int64     `json:"business_flow_type_id"`
	CreatedDate        time.Time `json:"created_date"`
	DeclineReason      struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"decline_reason"`
	DeliveryDetails struct {
		DirectDetails struct {
			BySellerDetails struct {
				DriverName                string `json:"driver_name"`
				VehicleRegistrationNumber string `json:"vehicle_registration_number"`
				VehicleType               string `json:"vehicle_type"`
			} `json:"by_seller_details"`
			ByTplDetails struct {
				TrackingNumber       string `json:"tracking_number"`
				TransportCompanyName string `json:"transport_company_name"`
			} `json:"by_tpl_details"`
			TimeslotDetails struct {
				Timeslot struct {
					TimeslotEnd   time.Time `json:"timeslot_end"`
					TimeslotStart time.Time `json:"timeslot_start"`
				} `json:"timeslot"`
				TimeslotReservationId string `json:"timeslot_reservation_id"`
			} `json:"timeslot_details"`
		} `json:"direct_details"`
		DropOffPoint struct {
			Id           int64  `json:"id"`
			ProvinceUuid string `json:"province_uuid"`
			Timeslot     struct {
				TimeslotEnd   time.Time `json:"timeslot_end"`
				TimeslotStart time.Time `json:"timeslot_start"`
			} `json:"timeslot"`
		} `json:"drop_off_point"`
		PickupDetails struct {
			Address     string    `json:"address"`
			Comment     string    `json:"comment"`
			Date        time.Time `json:"date"`
			SenderName  string    `json:"sender_name"`
			SenderPhone string    `json:"sender_phone"`
		} `json:"pickup_details"`
		SupplyType string `json:"supply_type"`
	} `json:"delivery_details"`
	HasAct            bool      `json:"has_act"`
	HasLabel          bool      `json:"has_label"`
	Id                int64     `json:"id"`
	OrderDraftId      int64     `json:"order_draft_id"`
	OrderNumber       string    `json:"order_number"`
	PackageUnitsCount int       `json:"package_units_count"`
	ReceiveDate       time.Time `json:"receive_date"`
	RowVersion        int64     `json:"row_version"`
	Status            string    `json:"status"`
	SupplyId          string    `json:"supply_id"`
	WarehouseId       int64     `json:"warehouse_id"`
}

type FbpArchiveListRsp struct {
	BaseRsp

	HasNext bool                 `json:"has_next"`
	Items   []FbpArchiveListItem `json:"items"`
	LastId  int64                `json:"last_id"`
}

type FbpArchiveListItem struct {
	ActFileUuid      string `json:"act_file_uuid"`
	BundleId         string `json:"bundle_id"`
	BundleSkuSummary struct {
		RoundedTotalVolumeInLitres float64 `json:"rounded_total_volume_in_litres"`
		TotalItemsCount            int     `json:"total_items_count"`
		TotalQuantity              int     `json:"total_quantity"`
	} `json:"bundle_sku_summary"`
	CreatedDate   time.Time `json:"created_date"`
	DeclineReason struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"decline_reason"`
	DeliveryDetails struct {
		DirectDetails struct {
			BySellerDetails struct {
				DriverName                string `json:"driver_name"`
				VehicleRegistrationNumber string `json:"vehicle_registration_number"`
				VehicleType               string `json:"vehicle_type"`
			} `json:"by_seller_details"`
			ByTplDetails struct {
				TrackingNumber       string `json:"tracking_number"`
				TransportCompanyName string `json:"transport_company_name"`
			} `json:"by_tpl_details"`
			TimeslotDetails struct {
				Timeslot struct {
					TimeslotEnd   time.Time `json:"timeslot_end"`
					TimeslotStart time.Time `json:"timeslot_start"`
				} `json:"timeslot"`
				TimeslotReservationId string `json:"timeslot_reservation_id"`
			} `json:"timeslot_details"`
		} `json:"direct_details"`
		DropOffPoint struct {
			Id           int64  `json:"id"`
			ProvinceUuid string `json:"province_uuid"`
			Timeslot     struct {
				TimeslotEnd   time.Time `json:"timeslot_end"`
				TimeslotStart time.Time `json:"timeslot_start"`
			} `json:"timeslot"`
		} `json:"drop_off_point"`
		PickupDetails struct {
			Address     string    `json:"address"`
			Comment     string    `json:"comment"`
			Date        time.Time `json:"date"`
			SenderName  string    `json:"sender_name"`
			SenderPhone string    `json:"sender_phone"`
		} `json:"pickup_details"`
		SupplyType string `json:"supply_type"`
	} `json:"delivery_details"`
	ExternalOrderId   string    `json:"external_order_id"`
	HasAct            bool      `json:"has_act"`
	HasLabel          bool      `json:"has_label"`
	OrderDraftId      int64     `json:"order_draft_id"`
	PackageUnitsCount int       `json:"package_units_count"`
	ReceiveDate       time.Time `json:"receive_date"`
	RowVersion        int       `json:"row_version"`
	Status            string    `json:"status"`
	SupplyId          string    `json:"supply_id"`
	WarehouseId       int64     `json:"warehouse_id"`
	WhcOrderId        int64     `json:"whc_order_id"`
}

type FbpLabelCreateRsp struct {
	BaseRsp

	Code string `json:"code"`
}

type FbpLabelGetRsp struct {
	BaseRsp

	LabelUrl string `json:"label_url"`
	State    string `json:"state"`
}

type FbpOrderGetRsp struct {
	BaseRsp

	AttentionReasons  string `json:"attention_reasons"`
	BundleUuid        string `json:"bundle_uuid"`
	CanBeCancelled    bool   `json:"can_be_cancelled"`
	CancellationState struct {
		CancellationError struct {
			ErrorCode string `json:"error_code"`
			Message   string `json:"message"`
		} `json:"cancellation_error"`
		CancellationStatus string `json:"cancellation_status"`
	} `json:"cancellation_state"`
	CreatedDate     time.Time `json:"created_date"`
	DeliveryDetails struct {
		DirectDetails struct {
			BySellerDetails struct {
				DriverName                string `json:"driver_name"`
				VehicleRegistrationNumber string `json:"vehicle_registration_number"`
				VehicleType               string `json:"vehicle_type"`
			} `json:"by_seller_details"`
			ByTplDetails struct {
				TrackingNumber       string `json:"tracking_number"`
				TransportCompanyName string `json:"transport_company_name"`
			} `json:"by_tpl_details"`
			TimeslotDetails struct {
				Timeslot struct {
					TimeslotEnd   time.Time `json:"timeslot_end"`
					TimeslotStart time.Time `json:"timeslot_start"`
				} `json:"timeslot"`
				TimeslotReservationId string `json:"timeslot_reservation_id"`
			} `json:"timeslot_details"`
		} `json:"direct_details"`
		DropOffPoint struct {
			Id           int64  `json:"id"`
			ProvinceUuid string `json:"province_uuid"`
			Timeslot     struct {
				TimeslotEnd   time.Time `json:"timeslot_end"`
				TimeslotStart time.Time `json:"timeslot_start"`
			} `json:"timeslot"`
		} `json:"drop_off_point"`
		PickupDetails struct {
			Address     string    `json:"address"`
			Comment     string    `json:"comment"`
			Date        time.Time `json:"date"`
			SenderName  string    `json:"sender_name"`
			SenderPhone string    `json:"sender_phone"`
		} `json:"pickup_details"`
		SupplyType string `json:"supply_type"`
	} `json:"delivery_details"`
	DraftId            int64     `json:"draft_id"`
	HasConsignmentNote bool      `json:"has_consignment_note"`
	HasLabel           bool      `json:"has_label"`
	Id                 int64     `json:"id"`
	Locked             bool      `json:"locked"`
	OrderNumber        string    `json:"order_number"`
	PackageUnitsCount  int64     `json:"package_units_count"`
	ReceiveDate        time.Time `json:"receive_date"`
	RowVersion         int64     `json:"row_version"`
	Status             string    `json:"status"`
	SupplyId           string    `json:"supply_id"`
	WarehouseId        int64     `json:"warehouse_id"`
}

type FbpOrderListRsp struct {
	BaseRsp

	HasNext bool               `json:"has_next"`
	Items   []FbpOrderListItem `json:"items"`
	LastId  int64              `json:"last_id"`
}

type FbpOrderListItem struct {
	AttentionReasons string `json:"attention_reasons"`
	BundleSummary    struct {
		RoundedTotalVolumeInLitres float64 `json:"rounded_total_volume_in_litres"`
		TotalItemCount             int     `json:"total_item_count"`
		TotalQuantity              int     `json:"total_quantity"`
	} `json:"bundle_summary"`
	CanBeCancelled    bool `json:"can_be_cancelled"`
	CancellationState struct {
		CancellationError struct {
			ErrorCode string `json:"error_code"`
			Message   string `json:"message"`
		} `json:"cancellation_error"`
		CancellationStatus string `json:"cancellation_status"`
	} `json:"cancellation_state"`
	CreatedDate     time.Time `json:"created_date"`
	DeliveryDetails struct {
		DirectDetails struct {
			BySellerDetails struct {
				DriverName                string `json:"driver_name"`
				VehicleRegistrationNumber string `json:"vehicle_registration_number"`
				VehicleType               string `json:"vehicle_type"`
			} `json:"by_seller_details"`
			ByTplDetails struct {
				TrackingNumber       string `json:"tracking_number"`
				TransportCompanyName string `json:"transport_company_name"`
			} `json:"by_tpl_details"`
			TimeslotDetails struct {
				Timeslot struct {
					TimeslotEnd   time.Time `json:"timeslot_end"`
					TimeslotStart time.Time `json:"timeslot_start"`
				} `json:"timeslot"`
				TimeslotReservationId string `json:"timeslot_reservation_id"`
			} `json:"timeslot_details"`
		} `json:"direct_details"`
		DropOffPoint struct {
			Id           int64  `json:"id"`
			ProvinceUuid string `json:"province_uuid"`
			Timeslot     struct {
				TimeslotEnd   time.Time `json:"timeslot_end"`
				TimeslotStart time.Time `json:"timeslot_start"`
			} `json:"timeslot"`
		} `json:"drop_off_point"`
		PickupDetails struct {
			Address     string    `json:"address"`
			Comment     string    `json:"comment"`
			Date        time.Time `json:"date"`
			SenderName  string    `json:"sender_name"`
			SenderPhone string    `json:"sender_phone"`
		} `json:"pickup_details"`
		SupplyType string `json:"supply_type"`
	} `json:"delivery_details"`
	HasConsignmentNote bool      `json:"has_consignment_note"`
	HasLabel           bool      `json:"has_label"`
	Id                 int64     `json:"id"`
	Locked             bool      `json:"locked"`
	OrderNumber        string    `json:"order_number"`
	PackageUnitsCount  int       `json:"package_units_count"`
	ReceiveDate        time.Time `json:"receive_date"`
	Status             string    `json:"status"`
	SupplyId           string    `json:"supply_id"`
	WarehouseId        int64     `json:"warehouse_id"`
}
