package shopee

type GetMerchantInfoRsp struct {
	BaseRsp
	MerchantName     string `json:"merchant_name"`
	IsCnsc           bool   `json:"is_cnsc"`
	AuthTime         int    `json:"auth_time"`
	ExpireTime       int    `json:"expire_time"`
	MerchantCurrency string `json:"merchant_currency"`
	MerchantRegion   string `json:"merchant_region"`
	IsUpgradedCbsc   bool   `json:"is_upgraded_cbsc"`
}

type GetShopListByMerchantRsp struct {
	BaseRsp
	IsCnsc   bool `json:"is_cnsc"`
	ShopList []struct {
		ShopId int64 `json:"shop_id"`
	} `json:"shop_list"`
	More bool `json:"more"`
}

type GetMerchantWarehouseLocationListRsp struct {
	BaseRsp
	Response []struct {
		LocationId    string `json:"location_id"`
		WarehouseName string `json:"warehouse_name"`
	} `json:"response"`
}

type GetMerchantWarehouseListRsp struct {
	BaseRsp
	Response struct {
		Cursor struct {
			NextId   interface{} `json:"next_id"`
			PageSize int         `json:"page_size"`
			PrevId   interface{} `json:"prev_id"`
		} `json:"cursor"`
		TotalCount    int `json:"total_count"`
		WarehouseList []struct {
			Address struct {
				Address     string `json:"address"`
				AddressName string `json:"address_name"`
				City        string `json:"city"`
				District    string `json:"district"`
				Region      string `json:"region"`
				State       string `json:"state"`
				Town        string `json:"town"`
				ZipCode     string `json:"zip_code"`
			} `json:"address"`
			EnterpriseInfo *struct {
				Cnpj                    string `json:"cnpj"`
				CompanyName             string `json:"company_name"`
				IsFreightPayer          bool   `json:"is_freight_payer"`
				StateRegistrationNumber string `json:"state_registration_number"`
			} `json:"enterprise_info"`
			LocationId      string `json:"location_id"`
			WarehouseId     int    `json:"warehouse_id"`
			WarehouseName   string `json:"warehouse_name"`
			WarehouseRegion string `json:"warehouse_region"`
			WarehouseType   int    `json:"warehouse_type"`
		} `json:"warehouse_list"`
	} `json:"response"`
}

type GetWarehouseEligibleShopListRsp struct {
	BaseRsp
	Response struct {
		Cursor struct {
			NextId   int         `json:"next_id"`
			PageSize int         `json:"page_size"`
			PrevId   interface{} `json:"prev_id"`
		} `json:"cursor"`
		ShopList []struct {
			ShopId   int64  `json:"shop_id"`
			ShopName string `json:"shop_name"`
		} `json:"shop_list"`
	} `json:"response"`
}
