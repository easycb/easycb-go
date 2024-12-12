package shopee

type GetShopInfoRsp struct {
	BaseRsp
	ShopName            string `json:"shop_name"`
	Region              string `json:"region"`
	Status              string `json:"status"`
	IsSip               bool   `json:"is_sip"`
	IsCb                bool   `json:"is_cb"`
	IsCnsc              bool   `json:"is_cnsc"`
	AuthTime            int64  `json:"auth_time"`
	ExpireTime          int64  `json:"expire_time"`
	ShopCbsc            string `json:"shop_cbsc"`
	Is3Pf               bool   `json:"is_3pf"`
	IsUpgradedCbsc      bool   `json:"is_upgraded_cbsc"`
	MerchantId          int64  `json:"merchant_id"`
	ShopFulfillmentFlag string `json:"shop_fulfillment_flag"`
	MtskuUpgradedStatus string `json:"mtsku_upgraded_status"`
}

type GetProfileRsp struct {
	BaseRsp
	Response struct {
		ShopLogo    string `json:"shop_logo"`
		Description string `json:"description"`
		ShopName    string `json:"shop_name"`
	} `json:"response"`
}

type UpdateProfileRsp struct {
	BaseRsp
	Response struct {
		ShopLogo    string `json:"shop_logo"`
		Description string `json:"description"`
		ShopName    string `json:"shop_name"`
	} `json:"response"`
}

type GetWarehouseDetailRsp struct {
	BaseRsp
	Response []struct {
		WarehouseId      int    `json:"warehouse_id"`
		WarehouseName    string `json:"warehouse_name"`
		LocationId       string `json:"location_id"`
		AddressId        int    `json:"address_id"`
		Region           string `json:"region"`
		State            string `json:"state"`
		City             string `json:"city"`
		District         string `json:"district"`
		Town             string `json:"town"`
		Address          string `json:"address"`
		Zipcode          string `json:"zipcode"`
		StateCode        string `json:"state_code"`
		HolidayModeState int    `json:"holiday_mode_state"`
	} `json:"response"`
}

type GetShopNotificationRsp struct {
	BaseRsp
	Cursor int `json:"cursor"`
	Data   struct {
		CreateTime int64  `json:"create_time"`
		Content    string `json:"content"`
		Title      string `json:"title"`
		Url        string `json:"url"`
	} `json:"data"`
}

type GetAuthorisedResellerBrandRsp struct {
	BaseRsp
	Response struct {
		AuthorisedBrandList []struct {
			BrandId   int    `json:"brand_id"`
			BrandName string `json:"brand_name"`
		} `json:"authorised_brand_list"`
		IsAuthorisedReseller bool `json:"is_authorised_reseller"`
		More                 bool `json:"more"`
		TotalCount           int  `json:"total_count"`
	} `json:"response"`
}
