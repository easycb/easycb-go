package shein

type WarehouseListRsp struct {
	BaseRsp
	Info struct {
		List []WarehouseListItem `json:"list"`
	} `json:"info"`
}

type WarehouseListItem struct {
	WarehouseCode   string   `json:"warehouseCode"`
	WarehouseName   string   `json:"warehouseName"`
	SaleCountryList []string `json:"saleCountryList"`
	CreateType      int      `json:"createType"`
	WarehouseType   int      `json:"warehouseType"`
	AuthServiceCode string   `json:"authServiceCode"`
	AuthServiceName string   `json:"authServiceName"`
}

type StockQueryRsp struct {
	BaseRsp
	Info []struct {
		GoodsInventory []GoodsInventoryItem `json:"goodsInventory"`
	} `json:"info"`
}

type GoodsInventoryItem struct {
	SpuName string `json:"spuName"`
	SkcName string `json:"skcName"`
	SkuList []struct {
		SkuCode                string `json:"skuCode"`
		TotalInventoryQuantity int    `json:"totalInventoryQuantity"`
		TotalLockedQuantity    int    `json:"totalLockedQuantity"`
		TotalUsableInventory   int    `json:"totalUsableInventory"`
		TotalTempLockQuantity  int    `json:"totalTempLockQuantity"`
		WarehouseInventoryList []struct {
			WarehouseCode     string `json:"warehouseCode"`
			WarehouseType     string `json:"warehouseType"`
			InventoryQuantity int    `json:"inventoryQuantity"`
			LockedQuantity    int    `json:"lockedQuantity"`
			UsableInventory   int    `json:"usableInventory"`
			TempLockQuantity  int    `json:"tempLockQuantity"`
		} `json:"warehouseInventoryList"`
	} `json:"skuList"`
}

type ChangeInventoryRsp struct {
	BaseRsp
	Info struct {
		SuccessList []struct {
			SkuCode string      `json:"skuCode"`
			Reason  interface{} `json:"reason"`
			Code    interface{} `json:"code"`
		} `json:"successList"`
		FailedList []struct {
			SkuCode string `json:"skuCode"`
			Reason  string `json:"reason"`
			Code    string `json:"code"`
		} `json:"failedList"`
	} `json:"info"`
}

type StockUpdateRsp struct {
	BaseRsp
	Info struct {
		BatchNumber string        `json:"batch_number"`
		FailNum     int           `json:"fail_num"`
		FailList    []interface{} `json:"fail_list"`
	} `json:"info"`
}

type QuerySkuSalesRsp struct {
	BaseRsp
	Info struct {
		DataList []QuerySkuSaleItem `json:"dataList"`
	} `json:"info"`
}

type QuerySkuSaleItem struct {
	SkuCode         string `json:"skuCode"`
	RealTimeSaleCnt int    `json:"realTimeSaleCnt"`
	C7DSaleCnt      int    `json:"c7dSaleCnt"`
	C30DSaleCnt     int    `json:"c30dSaleCnt"`
	Dt              string `json:"dt"`
}
