package ozon

import "time"

type CarriageCreateRsp struct {
	BaseRsp

	CarriageId int64 `json:"carriage_id"`
}

type CarriageApproveRsp struct {
	BaseRsp
}

type CarriageGetRsp struct {
	BaseRsp

	ActType            string   `json:"act_type"`
	ArrivalPassIds     []int64  `json:"arrival_pass_ids"`
	AvailableActions   []string `json:"available_actions"`
	CancelAvailability struct {
		IsCancelAvailable bool   `json:"is_cancel_available"`
		Reason            string `json:"reason"`
	} `json:"cancel_availability"`
	CarriageId                 int64     `json:"carriage_id"`
	CompanyId                  int64     `json:"company_id"`
	ContainersCount            int       `json:"containers_count"`
	CreatedAt                  time.Time `json:"created_at"`
	DeliveryMethodId           int64     `json:"delivery_method_id"`
	DepartureDate              string    `json:"departure_date"`
	FirstMileType              string    `json:"first_mile_type"`
	HasPostingsForNextCarriage bool      `json:"has_postings_for_next_carriage"`
	IntegrationType            string    `json:"integration_type"`
	IsContainerLabelPrinted    bool      `json:"is_container_label_printed"`
	IsPartial                  bool      `json:"is_partial"`
	PartialNum                 int64     `json:"partial_num"`
	RetryCount                 int       `json:"retry_count"`
	Status                     string    `json:"status"`
	TplProviderId              int64     `json:"tpl_provider_id"`
	UpdatedAt                  time.Time `json:"updated_at"`
	WarehouseId                int64     `json:"warehouse_id"`
}

type CarriageCourierContactSetRsp struct {
	BaseRsp
}

type CarriageCourierContactGetRsp struct {
	BaseRsp

	Contact []CarriageCourierContactGetContactItem `json:"contact"`
}

type CarriageCourierContactGetContactItem struct {
	CarriageId     int64     `json:"carriage_id"`
	Phone          string    `json:"phone"`
	WechatNickname string    `json:"wechat_nickname"`
	Comment        string    `json:"comment"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type PostingCarriageAvailableListRsp struct {
	BaseRsp

	Result []struct {
		CarriageId            int64     `json:"carriage_id"`
		CarriagePostingsCount int       `json:"carriage_postings_count"`
		CarriageStatus        string    `json:"carriage_status"`
		CutoffAt              time.Time `json:"cutoff_at"`
		DeliveryMethodId      int64     `json:"delivery_method_id"`
		DeliveryMethodName    string    `json:"delivery_method_name"`
		Errors                []struct {
			Code   string `json:"code"`
			Status string `json:"status"`
		} `json:"errors"`
		FirstMileType                     string `json:"first_mile_type"`
		HasEntrustedAcceptance            bool   `json:"has_entrusted_acceptance"`
		MandatoryPostingsCount            int    `json:"mandatory_postings_count"`
		MandatoryPackagedCount            int    `json:"mandatory_packaged_count"`
		RecommendedTimeLocal              string `json:"recommended_time_local"`
		RecommendedTimeUtcOffsetInMinutes int    `json:"recommended_time_utc_offset_in_minutes"`
		TplProviderIconUrl                string `json:"tpl_provider_icon_url"`
		TplProviderName                   string `json:"tpl_provider_name"`
		WarehouseCity                     string `json:"warehouse_city"`
		WarehouseId                       int64  `json:"warehouse_id"`
		WarehouseName                     string `json:"warehouse_name"`
		WarehouseTimezone                 string `json:"warehouse_timezone"`
	} `json:"result"`
}

type CarriageSetPostingsRsp struct {
	BaseRsp
	Result []struct {
		Error         string `json:"error"`
		PostingNumber string `json:"posting_number"`
		Result        bool   `json:"result"`
	} `json:"result"`
}

type CarriageCancelRsp struct {
	BaseRsp

	Error          string `json:"error"`
	CarriageStatus string `json:"carriage_status"`
}

type PostingFbsActGetPostingsRsp struct {
	BaseRsp
	Result []PostingFbsActGetPostingsResultItem `json:"result"`
}

type PostingFbsActGetPostingsResultItem struct {
	Id            int64     `json:"id"`
	MultiBoxQty   int       `json:"multi_box_qty"`
	PostingNumber string    `json:"posting_number"`
	Status        string    `json:"status"`
	SellerError   string    `json:"seller_error"`
	UpdatedAt     time.Time `json:"updated_at"`
	CreatedAt     time.Time `json:"created_at"`
	Products      []struct {
		Name     string `json:"name"`
		OfferId  string `json:"offer_id"`
		Price    string `json:"price"`
		Quantity int    `json:"quantity"`
		Sku      int64  `json:"sku"`
	} `json:"products"`
}

type PostingFbsActGetContainerLabelsRsp struct {
	BaseRsp
	ContentType string `json:"content_type"`
	FileName    string `json:"file_name"`
	FileContent string `json:"file_content"`
}

type AssemblyCarriagePostingListRsp struct {
	BaseRsp

	CanPrintMassLabel bool                                     `json:"can_print_mass_label"`
	Cursor            string                                   `json:"cursor"`
	Postings          []AssemblyCarriagePostingListPostingItem `json:"postings"`
}

type AssemblyCarriagePostingListPostingItem struct {
	AssemblyCode  string `json:"assembly_code"`
	CanPrintLabel bool   `json:"can_print_label"`
	PostingNumber string `json:"posting_number"`
	Products      []struct {
		OfferId     string `json:"offer_id"`
		PictureUrl  string `json:"picture_url"`
		ProductName string `json:"product_name"`
		Quantity    int    `json:"quantity"`
		Sku         int64  `json:"sku"`
	} `json:"products"`
}

type AssemblyCarriageProductListRsp struct {
	BaseRsp

	Cursor   string                           `json:"cursor"`
	Products []AssemblyCarriageProductLisItem `json:"products"`
}

type AssemblyCarriageProductLisItem struct {
	OfferId        string   `json:"offer_id"`
	PictureUrl     string   `json:"picture_url"`
	PostingNumbers []string `json:"posting_numbers"`
	ProductName    string   `json:"product_name"`
	Quantity       int      `json:"quantity"`
	Sku            int64    `json:"sku"`
}

type AssemblyFbsPostingListRsp struct {
	BaseRsp

	Cutoff   time.Time                           `json:"cutoff"`
	Postings []AssemblyFbsPostingListPostingItem `json:"postings"`
	Cursor   string                              `json:"cursor"`
}

type AssemblyFbsPostingListPostingItem struct {
	PostingNumber string `json:"posting_number"`
	AssemblyCode  string `json:"assembly_code"`
	Products      []struct {
		Sku         int64  `json:"sku"`
		OfferId     string `json:"offer_id"`
		ProductName string `json:"product_name"`
		PictureUrl  string `json:"picture_url"`
		Quantity    int    `json:"quantity"`
	} `json:"products"`
}

type AssemblyFbsProductListRsp struct {
	BaseRsp

	Products      []AssemblyFbsProductListProductItem `json:"products"`
	ProductsCount int                                 `json:"products_count"`
	HasNext       bool                                `json:"has_next"`
}

type AssemblyFbsProductListProductItem struct {
	Sku         int64  `json:"sku"`
	OfferId     string `json:"offer_id"`
	ProductName string `json:"product_name"`
	PictureUrl  string `json:"picture_url"`
	Quantity    int    `json:"quantity"`
	Postings    []struct {
		PostingNumber string `json:"posting_number"`
		Quantity      int    `json:"quantity"`
	} `json:"postings"`
}
