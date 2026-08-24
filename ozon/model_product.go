package ozon

import "time"

type ProductImportRsp struct {
	BaseRsp
	Result struct {
		TaskId int64 `json:"task_id"`
	} `json:"result"`
}

type ProductImportInfoRsp struct {
	BaseRsp

	Result struct {
		Items []ProductImportInfoItem `json:"items"`
		Total int                     `json:"total"`
	} `json:"result"`
}

type ProductImportInfoItem struct {
	OfferId   string                           `json:"offer_id"`
	ProductId int                              `json:"product_id"`
	Status    string                           `json:"status"`
	Errors    []ProductImportInfoItemErrorItem `json:"errors"`
}

type ProductImportInfoItemErrorItem struct {
	Code          string `json:"code"`
	Message       string `json:"message"`
	State         string `json:"state"`
	Level         string `json:"level"`
	Description   string `json:"description"`
	Field         string `json:"field"`
	AttributeId   int64  `json:"attribute_id"`
	AttributeName string `json:"attribute_name"`
}

type ProductImportBySkuRsp struct {
	BaseRsp
	Result struct {
		TaskId           int     `json:"task_id"`
		UnmatchedSkuList []int64 `json:"unmatched_sku_list"`
	} `json:"result"`
}

type ProductAttributesUpdateRsp struct {
	BaseRsp
	TaskId int64 `json:"task_id"`
}

type ProductPicturesImportRsp struct {
	BaseRsp

	Result struct {
		Pictures []ProductPicturesImportItem `json:"pictures"`
	} `json:"result"`
}

type ProductPicturesImportItem struct {
	IsColor   bool   `json:"is_color"`
	IsPrimary bool   `json:"is_primary"`
	ProductId int64  `json:"product_id"`
	State     string `json:"state"`
	Url       string `json:"url"`
}

type ProductListRsp struct {
	BaseRsp
	Result struct {
		Items  []ProductListItem `json:"items"`
		Total  int64             `json:"total"`
		LastId string            `json:"last_id"`
	} `json:"result"`
}

type ProductListItem struct {
	ProductId    int64  `json:"product_id"`
	OfferId      string `json:"offer_id"`
	Sku          int64  `json:"sku"`
	HasFboStocks bool   `json:"has_fbo_stocks"`
	HasFbsStocks bool   `json:"has_fbs_stocks"`
	Archived     bool   `json:"archived"`
	IsDiscounted bool   `json:"is_discounted"`
	Quants       []struct {
		QuantCode string `json:"quant_code"`
		QuantSize string `json:"quant_size"`
	} `json:"quants"`
}

type ProductRatingBySkuRsp struct {
	BaseRsp
	Products []ProductRatingBySkuItem `json:"products"`
}

type ProductRatingBySkuItem struct {
	Sku    int64                         `json:"sku"`
	Rating float64                       `json:"rating"`
	Groups []ProductRatingBySkuGroupItem `json:"groups"`
}

type ProductRatingBySkuGroupItem struct {
	Key        string  `json:"key"`
	Name       string  `json:"name"`
	Rating     float64 `json:"rating"`
	Weight     float64 `json:"weight"`
	Conditions []struct {
		Key         string  `json:"key"`
		Description string  `json:"description"`
		Fulfilled   bool    `json:"fulfilled"`
		Cost        float64 `json:"cost"`
	} `json:"conditions"`
	ImproveAttributes []struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"improve_attributes"`
	ImproveAtLeast int `json:"improve_at_least"`
}

type ProductInfoListRsp struct {
	BaseRsp
	Items []ProductInfoListItem `json:"items"`
}

type ProductInfoListItem struct {
	Availabilities        []ProductInfoListAvailabilitiesItem  `json:"availabilities"`
	Barcodes              []string                             `json:"barcodes"`
	ColorImage            []string                             `json:"color_image"`
	Commissions           []ProductInfoListCommissionsItem     `json:"commissions"`
	CreatedAt             time.Time                            `json:"created_at"`
	CurrencyCode          string                               `json:"currency_code"`
	DescriptionCategoryId int64                                `json:"description_category_id"`
	DiscountedFboStocks   int                                  `json:"discounted_fbo_stocks"`
	Errors                []ProductInfoListErrorsItem          `json:"errors"`
	HasDiscountedFboItem  bool                                 `json:"has_discounted_fbo_item"`
	Id                    int64                                `json:"id"`
	Images                []string                             `json:"images"`
	IsArchived            bool                                 `json:"is_archived"`
	IsAutoarchived        bool                                 `json:"is_autoarchived"`
	IsDiscounted          bool                                 `json:"is_discounted"`
	IsKgt                 bool                                 `json:"is_kgt"`
	IsPrepaymentAllowed   bool                                 `json:"is_prepayment_allowed"`
	IsSuper               bool                                 `json:"is_super"`
	MinPrice              string                               `json:"min_price"`
	ModelInfo             ProductInfoListModelInfoItem         `json:"model_info"`
	Name                  string                               `json:"name"`
	OfferId               string                               `json:"offer_id"`
	OldPrice              string                               `json:"old_price"`
	Price                 string                               `json:"price"`
	PriceIndexes          ProductInfoListPriceIndexesItem      `json:"price_indexes"`
	PrimaryImage          []string                             `json:"primary_image"`
	Promotions            []ProductInfoListPromotionsItem      `json:"promotions"`
	Sku                   int                                  `json:"sku"`
	Sources               []ProductInfoListSourceItem          `json:"sources"`
	Statuses              ProductInfoListStatusesItem          `json:"statuses"`
	Stocks                ProductInfoListStocksItem            `json:"stocks"`
	TypeId                int64                                `json:"type_id"`
	UpdatedAt             time.Time                            `json:"updated_at"`
	Vat                   string                               `json:"vat"`
	VisibilityDetails     ProductInfoListVisibilityDetailsItem `json:"visibility_details"`
	VolumeWeight          float64                              `json:"volume_weight"`
}

type ProductInfoListAvailabilitiesItem struct {
	Availability string `json:"availability"`
	Reasons      []struct {
		HumanText struct {
			Text string `json:"text"`
		} `json:"human_text"`
		Id int64 `json:"id"`
	} `json:"reasons"`
	Sku    int64  `json:"sku"`
	Source string `json:"source"`
}

type ProductInfoListCommissionsItem struct {
	DeliveryAmount float64 `json:"delivery_amount"`
	Percent        float64 `json:"percent"`
	ReturnAmount   float64 `json:"return_amount"`
	SaleSchema     string  `json:"sale_schema"`
	Value          float64 `json:"value"`
}

type ProductInfoListErrorsItem struct {
	AttributeId int64  `json:"attribute_id"`
	Code        string `json:"code"`
	Field       string `json:"field"`
	Level       string `json:"level"`
	State       string `json:"state"`
	Texts       struct {
		AttributeName string `json:"attribute_name"`
		Description   string `json:"description"`
		HintCode      string `json:"hint_code"`
		Message       string `json:"message"`
		Params        []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"params"`
		ShortDescription string `json:"short_description"`
	} `json:"texts"`
}

type ProductInfoListPriceIndexesItem struct {
	ColorIndex        string `json:"color_index"`
	ExternalIndexData struct {
		MinimalPrice         string  `json:"minimal_price"`
		MinimalPriceCurrency string  `json:"minimal_price_currency"`
		PriceIndexValue      float64 `json:"price_index_value"`
	} `json:"external_index_data"`
	OzonIndexData struct {
		MinimalPrice         string  `json:"minimal_price"`
		MinimalPriceCurrency string  `json:"minimal_price_currency"`
		PriceIndexValue      float64 `json:"price_index_value"`
	} `json:"ozon_index_data"`
	SelfMarketplacesIndexData struct {
		MinimalPrice         string  `json:"minimal_price"`
		MinimalPriceCurrency string  `json:"minimal_price_currency"`
		PriceIndexValue      float64 `json:"price_index_value"`
	} `json:"self_marketplaces_index_data"`
}

type ProductInfoListModelInfoItem struct {
	Count   int64 `json:"count"`
	ModelId int64 `json:"model_id"`
}

type ProductInfoListPromotionsItem struct {
	IsEnabled bool   `json:"is_enabled"`
	Type      string `json:"type"`
}

type ProductInfoListSourceItem struct {
	CreatedAt    string `json:"created_at"`
	QuantCode    string `json:"quant_code"`
	ShipmentType string `json:"shipment_type"`
	Sku          int64  `json:"sku"`
	Source       string `json:"source"`
}

type ProductInfoListStatusesItem struct {
	IsCreated         bool   `json:"is_created"`
	ModerateStatus    string `json:"moderate_status"`
	Status            string `json:"status"`
	StatusDescription string `json:"status_description"`
	StatusFailed      string `json:"status_failed"`
	StatusName        string `json:"status_name"`
	StatusTooltip     string `json:"status_tooltip"`
	StatusUpdatedAt   string `json:"status_updated_at"`
	ValidationStatus  string `json:"validation_status"`
}

type ProductInfoListStocksItem struct {
	HasStock bool `json:"has_stock"`
	Stocks   []struct {
		Present  int    `json:"present"`
		Reserved int    `json:"reserved"`
		Sku      int64  `json:"sku"`
		Source   string `json:"source"`
	} `json:"stocks"`
}

type ProductInfoListVisibilityDetailsItem struct {
	HasPrice bool `json:"has_price"`
	HasStock bool `json:"has_stock"`
}

type ProductInfoAttributesRsp struct {
	Result []ProductInfoAttribute `json:"result"`
	Total  int                    `json:"total"`
	LastId string                 `json:"last_id"`
}

type ProductInfoAttribute struct {
	Id                     int64                             `json:"id"`
	Barcode                string                            `json:"barcode"`
	Barcodes               []string                          `json:"barcodes"`
	Name                   string                            `json:"name"`
	OfferId                string                            `json:"offer_id"`
	TypeId                 int64                             `json:"type_id"`
	Height                 float64                           `json:"height"`
	Depth                  float64                           `json:"depth"`
	Width                  float64                           `json:"width"`
	DimensionUnit          string                            `json:"dimension_unit"`
	Weight                 float64                           `json:"weight"`
	WeightUnit             string                            `json:"weight_unit"`
	PrimaryImage           string                            `json:"primary_image"`
	Sku                    int64                             `json:"sku"`
	ModelInfo              ProductInfoAttributeModelInfoItem `json:"model_info"`
	Images                 []string                          `json:"images"`
	PdfList                []ProductInfoAttributePdfListItem `json:"pdf_list"`
	Attributes             []ProductInfoAttributeItem        `json:"attributes"`
	AttributesWithDefaults []int64                           `json:"attributes_with_defaults"`
	ComplexAttributes      []ProductInfoAttributeItem        `json:"complex_attributes"`
	ColorImage             string                            `json:"color_image"`
	DescriptionCategoryId  int64                             `json:"description_category_id"`
}

type ProductInfoAttributeItem struct {
	Id        int64 `json:"id"`
	ComplexId int64 `json:"complex_id"`
	Values    []struct {
		DictionaryValueId int    `json:"dictionary_value_id"`
		Value             string `json:"value"`
	} `json:"values"`
}

type ProductInfoAttributePdfListItem struct {
	FileName string `json:"file_name"`
	Name     string `json:"name"`
}

type ProductInfoAttributeModelInfoItem struct {
	ModelId int64 `json:"model_id"`
	Count   int64 `json:"count"`
}

type ProductInfoDescriptionRsp struct {
	BaseRsp
	Result struct {
		Id          int64  `json:"id"`
		OfferId     string `json:"offer_id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"result"`
}

type ProductInfoLimitRsp struct {
	BaseRsp
	DailyCreate struct {
		Limit   int64     `json:"limit"`
		ResetAt time.Time `json:"reset_at"`
		Usage   int64     `json:"usage"`
	} `json:"daily_create"`
	DailyUpdate struct {
		Limit   int64     `json:"limit"`
		ResetAt time.Time `json:"reset_at"`
		Usage   int64     `json:"usage"`
	} `json:"daily_update"`
	OperationLimits struct {
		Limit     int64  `json:"limit"`
		LimitType string `json:"limit_type"`
	} `json:"operation_limits"`
	Total struct {
		Limit           int64 `json:"limit"`
		Usage           int64 `json:"usage"`
		QuotaByCategory struct {
			CategoryId int64 `json:"category_id"`
			Limit      int64 `json:"limit"`
			Usage      int64 `json:"usage"`
		} `json:"quota_by_category"`
	} `json:"total"`
}

type ProductUpdateOfferIdRsp struct {
	Errors []struct {
		OfferId string `json:"offer_id"`
		Message string `json:"message"`
	} `json:"errors"`
}

type ProductArchiveRsp struct {
	BaseRsp
	Result bool `json:"result"`
}

type ProductUnarchiveRsp struct {
	BaseRsp
	Result bool `json:"result"`
}

type ProductDeleteRsp struct {
	BaseRsp
	Status []struct {
		OfferId   string `json:"offer_id"`
		IsDeleted bool   `json:"is_deleted"`
		Error     string `json:"error"`
	} `json:"status"`
}

type ProductInfoSubscriptionRsp struct {
	BaseRsp
	Result []struct {
		Count int64 `json:"count"`
		Sku   int64 `json:"sku"`
	} `json:"result"`
}

type ProductRelatedSkuGetRsp struct {
	BaseRsp
	Items []struct {
		Availability   string    `json:"availability"`
		DeletedAt      time.Time `json:"deleted_at"`
		DeliverySchema string    `json:"delivery_schema"`
		ProductId      int64     `json:"product_id"`
		Sku            int64     `json:"sku"`
	} `json:"items"`
	Errors []struct {
		Code    string `json:"code"`
		Sku     int64  `json:"sku"`
		Message string `json:"message"`
	} `json:"errors"`
}

type ProductPicturesInfoRsp struct {
	BaseRsp
	Items []struct {
		ProductId    int64    `json:"product_id"`
		PrimaryPhoto []string `json:"primary_photo"`
		Photo        []string `json:"photo"`
		ColorPhoto   []string `json:"color_photo"`
		Errors       []struct {
			Message string `json:"message"`
			Url     string `json:"url"`
		} `json:"errors"`
	} `json:"items"`
}

type ProductInfoWrongVolumeRsp struct {
	BaseRsp
	Cursor   string                       `json:"cursor"`
	Products []ProductInfoWrongVolumeItem `json:"products"`
}

type ProductInfoWrongVolumeItem struct {
	Height    int64  `json:"height"`
	Length    int64  `json:"length"`
	Name      string `json:"name"`
	OfferId   string `json:"offer_id"`
	ProductId int64  `json:"product_id"`
	Sku       int64  `json:"sku"`
	Weight    int64  `json:"weight"`
	Width     int64  `json:"width"`
}

type ProductStocksRsp struct {
	BaseRsp
	Result []ProductStockItem `json:"result"`
}

type ProductStockItem struct {
	WarehouseId int64  `json:"warehouse_id"`
	ProductId   int64  `json:"product_id"`
	OfferId     string `json:"offer_id"`
	Updated     bool   `json:"updated"`
	Errors      []struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"errors"`
}

type ProductInfoStockRsp struct {
	BaseRsp
	Items  []ProductInfoStockItem `json:"items"`
	Cursor string                 `json:"cursor"`
	Total  int                    `json:"total"`
}

type ProductInfoStockItem struct {
	OfferId   string `json:"offer_id"`
	ProductId int    `json:"product_id"`
	Stocks    []struct {
		Sku          int    `json:"sku"`
		Type         string `json:"type"`
		Present      int    `json:"present"`
		Reserved     int    `json:"reserved"`
		ShipmentType string `json:"shipment_type"`
	} `json:"stocks"`
}

type ProductInfoStocksByWarehouseFbsRsp struct {
	BaseRsp
	Cursor   string                                       `json:"cursor"`
	HasNext  bool                                         `json:"has_next"`
	Products []ProductInfoStocksByWarehouseFbsProductItem `json:"products"`
}

type ProductInfoStocksByWarehouseFbsProductItem struct {
	FreeStock     int64  `json:"free_stock"`
	OfferId       string `json:"offer_id"`
	Present       int64  `json:"present"`
	ProductId     int64  `json:"product_id"`
	Reserved      int64  `json:"reserved"`
	Sku           int64  `json:"sku"`
	WarehouseId   int64  `json:"warehouse_id"`
	WarehouseName string `json:"warehouse_name"`
}

type ProductInfoStocksByWarehouseFbsV1Rsp struct {
	BaseRsp
	Result []ProductInfoStocksByWarehouseFbsV1Item `json:"result"`
}

type ProductInfoStocksByWarehouseFbsV1Item struct {
	Sku           int64  `json:"sku"`
	OfferId       string `json:"offer_id"`
	Present       int64  `json:"present"`
	ProductId     int64  `json:"product_id"`
	Reserved      int64  `json:"reserved"`
	WarehouseId   int64  `json:"warehouse_id"`
	WarehouseName string `json:"warehouse_name"`
}

type ProductImportPricesRsp struct {
	BaseRsp
	Result []ProductImportPricesItem `json:"result"`
}

type ProductImportPricesItem struct {
	ProductId int64  `json:"product_id"`
	OfferId   string `json:"offer_id"`
	Updated   bool   `json:"updated"`
	Errors    []struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"errors"`
}

type ProductActionTimerUpdateRsp struct {
	BaseRsp
}

type ProductActionTimerStatusRsp struct {
	BaseRsp
	Statuses []struct {
		ExpiredAt                     time.Time `json:"expired_at"`
		MinPriceForAutoActionsEnabled bool      `json:"min_price_for_auto_actions_enabled"`
		ProductId                     int64     `json:"product_id"`
	} `json:"statuses"`
}

type ProductInfoPricesRsp struct {
	BaseRsp
	Items  []ProductInfoPricesItem `json:"items"`
	Cursor string                  `json:"cursor"`
	Total  int64                   `json:"total"`
}

type ProductInfoPricesItem struct {
	ProductId int64  `json:"product_id"`
	OfferId   string `json:"offer_id"`
	Price     struct {
		Price                           float64 `json:"price"`
		OldPrice                        float64 `json:"old_price"`
		MinPrice                        float64 `json:"min_price"`
		NetPrice                        float64 `json:"net_price"`
		CurrencyCode                    string  `json:"currency_code"`
		Vat                             float64 `json:"vat"`
		AutoActionEnabled               bool    `json:"auto_action_enabled"`
		AutoAddToOzonActionsListEnabled bool    `json:"auto_add_to_ozon_actions_list_enabled"`
	} `json:"price"`
	Commissions struct {
		SalesPercentFbo             float64 `json:"sales_percent_fbo"`
		SalesPercentFbs             float64 `json:"sales_percent_fbs"`
		SalesPercentFbp             float64 `json:"sales_percent_fbp"`
		SalesPercentRFbp            float64 `json:"sales_percent_rfbs"`
		FboDelivToCustomerAmount    float64 `json:"fbo_deliv_to_customer_amount"`
		FbsDelivToCustomerAmount    float64 `json:"fbs_deliv_to_customer_amount"`
		FboReturnFlowAmount         float64 `json:"fbo_return_flow_amount"`
		FbsReturnFlowAmount         float64 `json:"fbs_return_flow_amount"`
		FboDirectFlowTransMaxAmount float64 `json:"fbo_direct_flow_trans_max_amount"`
		FboDirectFlowTransMinAmount float64 `json:"fbo_direct_flow_trans_min_amount"`
		FbsDirectFlowTransMaxAmount float64 `json:"fbs_direct_flow_trans_max_amount"`
		FbsDirectFlowTransMinAmount float64 `json:"fbs_direct_flow_trans_min_amount"`
		FbsFirstMileMaxAmount       float64 `json:"fbs_first_mile_max_amount"`
		FbsFirstMileMinAmount       float64 `json:"fbs_first_mile_min_amount"`
	} `json:"commissions"`
	PriceIndexes struct {
		ColorIndex    string `json:"color_index"`
		OzonIndexData struct {
			MinPrice         float64 `json:"min_price"`
			MinPriceCurrency string  `json:"min_price_currency"`
			PriceIndexValue  float64 `json:"price_index_value"`
		} `json:"ozon_index_data"`
		ExternalIndexData struct {
			MinPrice         float64 `json:"min_price"`
			MinPriceCurrency string  `json:"min_price_currency"`
			PriceIndexValue  float64 `json:"price_index_value"`
		} `json:"external_index_data,omitempty"`
	} `json:"price_indexes"`
	MarketingActions struct {
		OzonActionsExist  bool   `json:"ozon_actions_exist"`
		CurrentPeriodFrom string `json:"current_period_from,omitempty"`
		CurrentPeriodTo   string `json:"current_period_to,omitempty"`
		Actions           []struct {
			Title    string    `json:"title"`
			Value    int       `json:"value"`
			DateFrom time.Time `json:"date_from"`
			DateTo   time.Time `json:"date_to"`
		} `json:"actions,omitempty"`
	} `json:"marketing_actions"`
	Acquiring    float64 `json:"acquiring"`
	VolumeWeight float64 `json:"volume_weight"`
}

type ProductInfoDiscountedRsp struct {
	BaseRsp
	Items []ProductInfoDiscountedItem `json:"items"`
}

type ProductInfoDiscountedItem struct {
	DiscountedSku        int64  `json:"discounted_sku"`
	Sku                  int64  `json:"sku"`
	ConditionEstimation  string `json:"condition_estimation"`
	PackagingViolation   string `json:"packaging_violation"`
	WarrantyType         string `json:"warranty_type"`
	ReasonDamaged        string `json:"reason_damaged"`
	CommentReasonDamaged string `json:"comment_reason_damaged"`
	Defects              string `json:"defects"`
	MechanicalDamage     string `json:"mechanical_damage"`
	PackageDamage        string `json:"package_damage"`
	Shortage             string `json:"shortage"`
	Repair               string `json:"repair"`
	Condition            string `json:"condition"`
}

type ProductUpdateDiscountRsp struct {
	BaseRsp
	Result bool `json:"result"`
}

type ProductInfoWarehouseStocksRsp struct {
	BaseRsp
	Stocks []struct {
		Sku         int64  `json:"sku"`
		ProductId   int64  `json:"product_id"`
		OfferId     string `json:"offer_id"`
		WarehouseId int64  `json:"warehouse_id"`
		Present     int64  `json:"present"`
		Reserved    int64  `json:"reserved"`
		FreeStock   int64  `json:"free_stock"`
		UpdatedAt   string `json:"updated_at"`
	} `json:"stocks"`
	HasNext bool   `json:"has_next"`
	Cursor  string `json:"cursor"`
}

type ProductStairwayDiscountByQuantitySetRsp struct {
	BaseRsp
	Accepted bool                                               `json:"accepted"`
	Errors   []ProductStairwayDiscountByQuantitySetErrorItem    `json:"errors"`
	Warnings []ProductStairwayDiscountByQuantitySetWarningsItem `json:"warnings"`
}

type ProductStairwayDiscountByQuantitySetErrorItem struct {
	Data []struct {
		Code    string `json:"code"`
		Field   string `json:"field"`
		Message string `json:"message"`
		Step    int    `json:"step"`
		Value   string `json:"value"`
	} `json:"data"`
	Sku int `json:"sku"`
}

type ProductStairwayDiscountByQuantitySetWarningsItem struct {
	Data []struct {
		Code    string `json:"code"`
		Field   string `json:"field"`
		Message string `json:"message"`
		Step    int    `json:"step"`
		Value   string `json:"value"`
	} `json:"data"`
	Sku int64 `json:"sku"`
}

type ProductStairwayDiscountByQuantityGetRsp struct {
	BaseRsp
	Stairways []ProductStairwayDiscountByQuantityGetStairwayItem `json:"stairways"`
}

type ProductStairwayDiscountByQuantityGetStairwayItem struct {
	Enabled  bool  `json:"enabled"`
	Sku      int64 `json:"sku"`
	Stairway struct {
		Steps []struct {
			Discount int64 `json:"discount"`
			Quantity int64 `json:"quantity"`
			Step     int64 `json:"step"`
		} `json:"steps"`
	} `json:"stairway"`
	Status string `json:"status"`
}

type ProductVisibilitySetRsp struct {
	BaseRsp

	Items []struct {
		SelectPermission        string   `json:"select_permission"`
		SellerItemPlacement     string   `json:"seller_item_placement"`
		SellerItemPlacementList []string `json:"seller_item_placement_list"`
		ShowcasesVisibility     string   `json:"showcases_visibility"`
		ShowcasesVisibilityList []string `json:"showcases_visibility_list"`
		Sku                     int64    `json:"sku"`
		Warnings                []string `json:"warnings"`
	} `json:"items"`
	ItemsErrors []struct {
		Code string `json:"code"`
		Sku  int64  `json:"sku"`
	} `json:"items_errors"`
}

type ProductVisibilityInfoRsp struct {
	BaseRsp
	Items []struct {
		ShowcasesVisibility string `json:"showcases_visibility"`
		Sku                 int64  `json:"sku"`
	} `json:"items"`
}
