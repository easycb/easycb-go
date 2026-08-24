package ozon

import "time"

type PostingFbsUnfulfilledListRsp struct {
	BaseRsp

	HasNext  bool                 `json:"has_next"`
	Cursor   string               `json:"cursor"`
	Postings []PostingFbsListItem `json:"postings"`
	Count    int64                `json:"count"`
}

type PostingFbsListRsp struct {
	BaseRsp
	HasNext  bool                 `json:"has_next"`
	Cursor   string               `json:"cursor"`
	Postings []PostingFbsListItem `json:"postings"`
}

type PostingFbsListItem struct {
	PostingNumber        string    `json:"posting_number"`
	OrderId              int64     `json:"order_id"`
	OrderNumber          string    `json:"order_number"`
	PickupCodeVerifiedAt time.Time `json:"pickup_code_verified_at"`
	Status               string    `json:"status"`
	Substatus            string    `json:"substatus"`
	DeliveryMethod       struct {
		Id            int    `json:"id"`
		Name          string `json:"name"`
		WarehouseId   int64  `json:"warehouse_id"`
		Warehouse     string `json:"warehouse"`
		TplProviderId int64  `json:"tpl_provider_id"`
		TplProvider   string `json:"tpl_provider"`
	} `json:"delivery_method"`
	DeliverySchema           string    `json:"delivery_schema"`
	TrackingNumber           string    `json:"tracking_number"`
	TplIntegrationType       string    `json:"tpl_integration_type"`
	InProcessAt              time.Time `json:"in_process_at"`
	ShipmentDate             time.Time `json:"shipment_date"`
	ShipmentDateWithoutDelay time.Time `json:"shipment_date_without_delay"`
	Optional                 struct {
		ProductsWithPossibleMandatoryMark []int64 `json:"products_with_possible_mandatory_mark"`
	} `json:"optional"`
	Cancellation struct {
		CancelReasonId           int64  `json:"cancel_reason_id"`
		CancelReason             string `json:"cancel_reason"`
		CancellationType         string `json:"cancellation_type"`
		CancelledAfterShip       bool   `json:"cancelled_after_ship"`
		AffectCancellationRating bool   `json:"affect_cancellation_rating"`
		CancellationInitiator    string `json:"cancellation_initiator"`
	} `json:"cancellation"`
	Customer struct {
		CustomerId    int64  `json:"customer_id"`
		CustomerEmail string `json:"customer_email"`
		Phone         string `json:"phone"`
		Name          string `json:"name"`
		Address       struct {
			AddressTail     string  `json:"address_tail"`
			City            string  `json:"city"`
			Comment         string  `json:"comment"`
			Country         string  `json:"country"`
			District        string  `json:"district"`
			Region          string  `json:"region"`
			ZipCode         string  `json:"zip_code"`
			Latitude        float64 `json:"latitude"`
			Longitude       float64 `json:"longitude"`
			PvzCode         int64   `json:"pvz_code"`
			ProviderPvzCode string  `json:"provider_pvz_code"`
		} `json:"address"`
	} `json:"customer"`
	Products []struct {
		IsBlrTraceable      bool     `json:"is_blr_traceable"`
		IsMarketplaceBuyout bool     `json:"is_marketplace_buyout"`
		OfferId             string   `json:"offer_id"`
		Name                string   `json:"name"`
		Sku                 int64    `json:"sku"`
		Quantity            int      `json:"quantity"`
		Imei                []string `json:"imei"`
		Weight              float64  `json:"weight"`
		ProductColor        string   `json:"product_color"`
		Price               struct {
			Amount   string `json:"amount"`
			Currency string `json:"currency"`
		} `json:"price"`
	} `json:"products"`
	Addressee struct {
		Name string `json:"name"`
	} `json:"addressee"`
	Barcodes struct {
		UpperBarcode string `json:"upper_barcode"`
		LowerBarcode string `json:"lower_barcode"`
	} `json:"barcodes"`
	AnalyticsData struct {
		Region                  string    `json:"region"`
		City                    string    `json:"city"`
		DeliveryType            string    `json:"delivery_type"`
		IsPremium               bool      `json:"is_premium"`
		PaymentTypeGroupName    string    `json:"payment_type_group_name"`
		WarehouseId             int64     `json:"warehouse_id"`
		Warehouse               string    `json:"warehouse"`
		TplProviderId           int64     `json:"tpl_provider_id"`
		TplProvider             string    `json:"tpl_provider"`
		DeliveryDateBegin       time.Time `json:"delivery_date_begin"`
		DeliveryDateEnd         time.Time `json:"delivery_date_end"`
		IsLegal                 bool      `json:"is_legal"`
		ClientDeliveryDateBegin time.Time `json:"client_delivery_date_begin"`
		ClientDeliveryDateEnd   time.Time `json:"client_delivery_date_end"`
	} `json:"analytics_data"`
	DestinationPlaceId   int64  `json:"destination_place_id"`
	DestinationPlaceName string `json:"destination_place_name"`
	FinancialData        struct {
		Products []struct {
			Payout               float64 `json:"payout"`
			ProductId            int64   `json:"product_id"`
			OldPrice             float64 `json:"old_price"`
			Price                float64 `json:"price"`
			TotalDiscountValue   float64 `json:"total_discount_value"`
			TotalDiscountPercent float64 `json:"total_discount_percent"`
			Quantity             int     `json:"quantity"`
			CustomerPrice        struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"customer_price"`
			Actions    []string `json:"actions"`
			Commission struct {
				Amount   float64 `json:"amount"`
				Percent  float64 `json:"percent"`
				Currency string  `json:"currency"`
			} `json:"commission"`
		} `json:"products"`
		ClusterFrom string `json:"cluster_from"`
		ClusterTo   string `json:"cluster_to"`
	} `json:"financial_data"`
	IsExpress bool `json:"is_express"`
	LegalInfo struct {
		CompanyName string `json:"company_name"`
		Inn         string `json:"inn"`
		Kpp         string `json:"kpp"`
	} `json:"legal_info"`
	QuantumId                int  `json:"quantum_id"`
	RequireBlrTraceableAttrs bool `json:"require_blr_traceable_attrs"`
	Requirements             struct {
		ProductsRequiringGtd           []int64 `json:"products_requiring_gtd"`
		ProductsRequiringCountry       []int64 `json:"products_requiring_country"`
		ProductsRequiringMandatoryMark []int64 `json:"products_requiring_mandatory_mark"`
		ProductsRequiringRnpt          []int64 `json:"products_requiring_rnpt"`
		ProductsRequiringJwUin         []int64 `json:"products_requiring_jw_uin"`
		ProductsRequiringChangeCountry []int64 `json:"products_requiring_change_country"`
		ProductsRequiringImei          []int64 `json:"products_requiring_imei"`
		ProductsRequiringWeight        []int64 `json:"products_requiring_weight"`
	} `json:"requirements"`
	Tariffication struct {
		CurrentTariffRate   float64 `json:"current_tariff_rate"`
		CurrentTariffType   string  `json:"current_tariff_type"`
		CurrentTariffCharge struct {
			Amount   string `json:"amount"`
			Currency string `json:"currency"`
		} `json:"current_tariff_charge"`
		CurrentTariffMinCharge struct {
			Amount   string `json:"amount"`
			Currency string `json:"currency"`
		} `json:"current_tariff_min_charge"`
		NextTariffRate   float64 `json:"next_tariff_rate"`
		NextTariffType   string  `json:"next_tariff_type"`
		NextTariffCharge struct {
			Amount   string `json:"amount"`
			Currency string `json:"currency"`
		} `json:"next_tariff_charge"`
		NextTariffMinCharge struct {
			Amount   string `json:"amount"`
			Currency string `json:"currency"`
		} `json:"next_tariff_min_charge"`
		NextTariffStartsAt time.Time `json:"next_tariff_starts_at"`
	} `json:"tariffication"`
	ExternalOrder struct {
		IsExternal   bool   `json:"is_external"`
		PlatformName string `json:"platform_name"`
	} `json:"external_order"`
	VolumeWeight        float64   `json:"volume_weight"`
	IsClickAndCollect   bool      `json:"is_click_and_collect"`
	DeliveringDate      time.Time `json:"delivering_date"`
	IsMultibox          bool      `json:"is_multibox"`
	MultiBoxQty         int       `json:"multi_box_qty"`
	IsPresortable       bool      `json:"is_presortable"`
	PrrOption           string    `json:"prr_option"`
	ParentPostingNumber string    `json:"parent_posting_number"`
	AvailableActions    []string  `json:"available_actions"`
	TarifficationSteps  []struct {
		MinCharge struct {
			Amount   string `json:"amount"`
			Currency string `json:"currency"`
		} `json:"min_charge"`
		TariffCharge struct {
			Amount   string `json:"amount"`
			Currency string `json:"currency"`
		} `json:"tariff_charge"`
		TariffDeadlineAt time.Time `json:"tariff_deadline_at"`
		TariffRate       float64   `json:"tariff_rate"`
		TariffType       string    `json:"tariff_type"`
	} `json:"tariffication_steps"`
	ContainerSortType   string      `json:"container_sort_type"`
	Container           interface{} `json:"container"`
	IntegrationTypeFlow string      `json:"integration_type_flow"`
	SortingCenter       struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"sorting_center"`
	ReceivedAtSortingCenter time.Time `json:"received_at_sorting_center"`
}

type PostingFbsGetRsp struct {
	Result struct {
		PostingNumber  string `json:"posting_number"`
		OrderId        int64  `json:"order_id"`
		OrderNumber    string `json:"order_number"`
		Status         string `json:"status"`
		DeliveryMethod struct {
			Id            int64  `json:"id"`
			Name          string `json:"name"`
			WarehouseId   int64  `json:"warehouse_id"`
			Warehouse     string `json:"warehouse"`
			TplProviderId int64  `json:"tpl_provider_id"`
			TplProvider   string `json:"tpl_provider"`
		} `json:"delivery_method"`
		TrackingNumber     string    `json:"tracking_number"`
		TplIntegrationType string    `json:"tpl_integration_type"`
		InProcessAt        time.Time `json:"in_process_at"`
		ShipmentDate       time.Time `json:"shipment_date"`
		DeliveringDate     time.Time `json:"delivering_date"`
		ProviderStatus     string    `json:"provider_status"`
		DeliveryPrice      string    `json:"delivery_price"`
		Cancellation       struct {
			CancelReasonId           int64  `json:"cancel_reason_id"`
			CancelReason             string `json:"cancel_reason"`
			CancellationType         string `json:"cancellation_type"`
			CancelledAfterShip       bool   `json:"cancelled_after_ship"`
			AffectCancellationRating bool   `json:"affect_cancellation_rating"`
			CancellationInitiator    string `json:"cancellation_initiator"`
		} `json:"cancellation"`
		Customer struct {
			CustomerId    int64  `json:"customer_id"`
			CustomerEmail string `json:"customer_email"`
			Phone         string `json:"phone"`
			Address       struct {
				AddressTail     string  `json:"address_tail"`
				City            string  `json:"city"`
				Comment         string  `json:"comment"`
				Country         string  `json:"country"`
				District        string  `json:"district"`
				Region          string  `json:"region"`
				ZipCode         string  `json:"zip_code"`
				Latitude        float64 `json:"latitude"`
				Longitude       float64 `json:"longitude"`
				PvzCode         int64   `json:"pvz_code"`
				ProviderPvzCode string  `json:"provider_pvz_code"`
			} `json:"address"`
			Name string `json:"name"`
		} `json:"customer"`
		Addressee struct {
			Name  string `json:"name"`
			Phone string `json:"phone"`
			Pin   string `json:"pin"`
		} `json:"addressee"`
		Products []struct {
			Price      string `json:"price"`
			OfferId    string `json:"offer_id"`
			Name       string `json:"name"`
			Sku        int64  `json:"sku"`
			Quantity   int    `json:"quantity"`
			Dimensions struct {
				Height string `json:"height"`
				Length string `json:"length"`
				Weight string `json:"weight"`
				Width  string `json:"width"`
			} `json:"dimensions"`
			CurrencyCode        string  `json:"currency_code"`
			IsBlrTraceable      bool    `json:"is_blr_traceable"`
			IsMarketplaceBuyout bool    `json:"is_marketplace_buyout"`
			HasImei             bool    `json:"has_imei"`
			WeightMax           float64 `json:"weight_max"`
			WeightMin           float64 `json:"weight_min"`
			IsWeightNeeded      bool    `json:"is_weight_needed"`
		} `json:"products"`
		Barcodes struct {
			UpperBarcode string `json:"upper_barcode"`
			LowerBarcode string `json:"lower_barcode"`
		} `json:"barcodes"`
		AnalyticsData struct {
			Region                  string    `json:"region"`
			City                    string    `json:"city"`
			DeliveryType            string    `json:"delivery_type"`
			IsPremium               bool      `json:"is_premium"`
			PaymentTypeGroupName    string    `json:"payment_type_group_name"`
			WarehouseId             int64     `json:"warehouse_id"`
			Warehouse               string    `json:"warehouse"`
			TplProviderId           int64     `json:"tpl_provider_id"`
			TplProvider             string    `json:"tpl_provider"`
			DeliveryDateBegin       time.Time `json:"delivery_date_begin"`
			DeliveryDateEnd         time.Time `json:"delivery_date_end"`
			IsLegal                 bool      `json:"is_legal"`
			ClientDeliveryDateBegin time.Time `json:"client_delivery_date_begin"`
			ClientDeliveryDateEnd   time.Time `json:"client_delivery_date_end"`
		} `json:"analytics_data"`
		FinancialData struct {
			Products []struct {
				CommissionAmount     float64  `json:"commission_amount"`
				CommissionPercent    float64  `json:"commission_percent"`
				Payout               float64  `json:"payout"`
				ProductId            int64    `json:"product_id"`
				OldPrice             float64  `json:"old_price"`
				Price                float64  `json:"price"`
				TotalDiscountValue   float64  `json:"total_discount_value"`
				TotalDiscountPercent float64  `json:"total_discount_percent"`
				Actions              []string `json:"actions"`
				Quantity             int      `json:"quantity"`
				CurrencyCode         string   `json:"currency_code"`
				CustomerCurrencyCode string   `json:"customer_currency_code"`
				CustomerPrice        float64  `json:"customer_price"`
			} `json:"products"`
			ClusterFrom string `json:"cluster_from"`
			ClusterTo   string `json:"cluster_to"`
		} `json:"financial_data"`
		AdditionalData []interface{} `json:"additional_data"`
		IsExpress      bool          `json:"is_express"`
		Requirements   struct {
			ProductsRequiringGtd           []int64 `json:"products_requiring_gtd"`
			ProductsRequiringCountry       []int64 `json:"products_requiring_country"`
			ProductsRequiringMandatoryMark []int64 `json:"products_requiring_mandatory_mark"`
			ProductsRequiringRnpt          []int64 `json:"products_requiring_rnpt"`
			ProductsRequiringJwUin         []int64 `json:"products_requiring_jw_uin"`
			ProductsRequiringChangeCountry []int64 `json:"products_requiring_change_country"`
			ProductsRequiringImei          []int64 `json:"products_requiring_imei"`
			ProductsRequiringWeight        []int64 `json:"products_requiring_weight"`
		} `json:"requirements"`
		ProductExemplars struct {
			Products []struct {
				Sku       int64 `json:"sku"`
				Exemplars []struct {
					MandatoryMark string   `json:"mandatory_mark"`
					Gtd           string   `json:"gtd"`
					IsGtdAbsent   bool     `json:"is_gtd_absent"`
					Rnpt          string   `json:"rnpt"`
					IsRnptAbsent  bool     `json:"is_rnpt_absent"`
					JwUin         string   `json:"jw_uin"`
					ExemplarId    int64    `json:"exemplar_id"`
					Imei          []string `json:"imei"`
					Weight        int      `json:"weight"`
				} `json:"exemplars"`
			} `json:"products"`
		} `json:"product_exemplars"`
		Courier struct {
			CarModel  string `json:"car_model"`
			CarNumber string `json:"car_number"`
			Name      string `json:"name"`
			Phone     string `json:"phone"`
		} `json:"courier"`
		ParentPostingNumber string `json:"parent_posting_number"`
		RelatedPostings     struct {
			RelatedPostingNumbers []string `json:"related_posting_numbers"`
		} `json:"related_postings"`
		AvailableActions  []string `json:"available_actions"`
		MultiBoxQty       int      `json:"multi_box_qty"`
		IsMultibox        bool     `json:"is_multibox"`
		Substatus         string   `json:"substatus"`
		PreviousSubstatus string   `json:"previous_substatus"`
		PrrOption         struct {
			Code         string `json:"code"`
			Price        string `json:"price"`
			CurrencyCode string `json:"currency_code"`
			Floor        string `json:"floor"`
		} `json:"prr_option"`
		Tariffication struct {
			CurrentTariffRate               float64   `json:"current_tariff_rate"`
			CurrentTariffType               string    `json:"current_tariff_type"`
			CurrentTariffCharge             string    `json:"current_tariff_charge"`
			CurrentTariffChargeCurrencyCode string    `json:"current_tariff_charge_currency_code"`
			NextTariffRate                  float64   `json:"next_tariff_rate"`
			NextTariffType                  string    `json:"next_tariff_type"`
			NextTariffCharge                string    `json:"next_tariff_charge"`
			NextTariffStartsAt              time.Time `json:"next_tariff_starts_at"`
			NextTariffChargeCurrencyCode    string    `json:"next_tariff_charge_currency_code"`
		} `json:"tariffication"`
		DestinationPlaceId   int       `json:"destination_place_id"`
		DestinationPlaceName string    `json:"destination_place_name"`
		IsPresortable        bool      `json:"is_presortable"`
		PickupCodeVerifiedAt time.Time `json:"pickup_code_verified_at"`
		Optional             struct {
			ProductsWithPossibleMandatoryMark []int64 `json:"products_with_possible_mandatory_mark"`
		} `json:"optional"`
		FactDeliveryDate time.Time `json:"fact_delivery_date"`
		LegalInfo        struct {
			CompanyName string `json:"company_name"`
			Inn         string `json:"inn"`
			Kpp         string `json:"kpp"`
		} `json:"legal_info"`
		RelatedWeightPostings struct {
			RelatedPostingNumbers []string `json:"related_posting_numbers"`
		} `json:"related_weight_postings"`
		ShipmentDateWithoutDelay time.Time `json:"shipment_date_without_delay"`
		RequireBlrTraceableAttrs bool      `json:"require_blr_traceable_attrs"`
		IsClickAndCollect        bool      `json:"is_click_and_collect"`
		TarifficationSteps       []struct {
			MinCharge struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"min_charge"`
			TariffCharge struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"tariff_charge"`
			TariffDeadlineAt time.Time `json:"tariff_deadline_at"`
			TariffRate       float64   `json:"tariff_rate"`
			TariffType       string    `json:"tariff_type"`
		}
		ContainerSortType   string      `json:"container_sort_type"`
		Container           interface{} `json:"container"`
		IntegrationTypeFlow string      `json:"integration_type_flow"`
		SortingCenter       struct {
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"sorting_center"`
		ExternalOrder struct {
			IsExternal   bool   `json:"is_external"`
			PlatformName string `json:"platform_name"`
		} `json:"external_order"`
		ReceivedAtSortingCenter time.Time `json:"received_at_sorting_center"`
	} `json:"result"`
}

type PostingFbsGetByBarcodeRsp struct {
	BaseRsp

	Result PostingFbsGetByBarcodeResult `json:"result"`
}

type PostingFbsGetByBarcodeResult struct {
	OrderId        int64     `json:"order_id"`
	OrderNumber    string    `json:"order_number"`
	PostingNumber  string    `json:"posting_number"`
	Status         string    `json:"status"`
	CancelReasonId int64     `json:"cancel_reason_id"`
	CreatedAt      time.Time `json:"created_at"`
	InProcessAt    time.Time `json:"in_process_at"`
	ShipmentDate   time.Time `json:"shipment_date"`
	Products       []struct {
		Sku      int64  `json:"sku"`
		Name     string `json:"name"`
		Quantity int    `json:"quantity"`
		OfferId  string `json:"offer_id"`
		Price    string `json:"price"`
	} `json:"products"`
	Barcodes struct {
		UpperBarcode string `json:"upper_barcode"`
		LowerBarcode string `json:"lower_barcode"`
	} `json:"barcodes"`
}

type PostingFbsProductCountryListRsp struct {
	BaseRsp
	Result []struct {
		Name           string `json:"name"`
		CountryIsoCode string `json:"country_iso_code"`
	} `json:"result"`
}

type PostingFbsProductCountrySetRsp struct {
	BaseRsp

	ProductId   int64 `json:"product_id"`
	IsGtdNeeded bool  `json:"is_gtd_needed"`
}

type PostingFbsPackageLabelRsp struct {
	BaseRsp

	ContentType string `json:"content_type"`
	FileName    string `json:"file_name"`
	FileContent string `json:"file_content"`
}

type PostingFbsAwaitingDeliveryRsp struct {
	BaseRsp

	Result bool `json:"result"`
}

type PostingFbsCancelReasonListRsp struct {
	BaseRsp

	Result []PostingFbsCancelReasonListItem `json:"result"`
}

type PostingFbsCancelReasonListItem struct {
	Id                         int64  `json:"id"`
	Title                      string `json:"title"`
	TypeId                     string `json:"type_id"`
	IsAvailableForCancellation bool   `json:"is_available_for_cancellation"`
}

type PostingFbsCancelReasonRsp struct {
	BaseRsp

	Result []PostingFbsCancelReasonItem `json:"result"`
}

type PostingFbsCancelReasonItem struct {
	PostingNumber string `json:"posting_number"`
	Reasons       []struct {
		Id     int64  `json:"id"`
		Title  string `json:"title"`
		TypeId string `json:"type_id"`
	} `json:"reasons"`
}

type PostingFbsCancelRsp struct {
	BaseRsp

	Result bool `json:"result"`
}

type PostingFbsProductCancelRsp struct {
	BaseRsp

	Result string `json:"result"`
}

type PostingFbsShipRsp struct {
	BaseRsp

	AdditionalData []PostingFbsShipAdditionalDataItem `json:"additional_data"`
	Result         []string                           `json:"result"`
}

type PostingFbsShipAdditionalDataItem struct {
	PostingNumber string                                    `json:"posting_number"`
	Products      []PostingFbsShipAdditionalDataProductItem `json:"products"`
}

type PostingFbsShipAdditionalDataProductItem struct {
	CurrencyCode  string   `json:"currency_code"`
	MandatoryMark []string `json:"mandatory_mark"`
	Name          string   `json:"name"`
	OfferId       string   `json:"offer_id"`
	Price         string   `json:"price"`
	Quantity      int      `json:"quantity"`
	Sku           int64    `json:"sku"`
}

type PostingFbsShipPackageRsp struct {
	BaseRsp
	Result string `json:"result"`
}

type PostingFbsSplitRsp struct {
	BaseRsp

	ParentPosting struct {
		PostingNumber string `json:"posting_number"`
		Products      []struct {
			ProductId int `json:"product_id"`
			Quantity  int `json:"quantity"`
		} `json:"products"`
	} `json:"parent_posting"`
	Postings []struct {
		PostingNumber string `json:"posting_number"`
		Products      []struct {
			ProductId int `json:"product_id"`
			Quantity  int `json:"quantity"`
		} `json:"products"`
	} `json:"postings"`
}
