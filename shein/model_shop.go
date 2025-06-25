package shein

type QueryStoreInfoRsp struct {
	BaseRsp
	Info struct {
		StoreProductQuota struct {
			TotalLimit     int `json:"totalLimit"`
			AvailableLimit int `json:"availableLimit"`
			UsedQuota      int `json:"usedQuota"`
		} `json:"storeProductQuota"`
		StoreInfo struct {
			SupplierId  int         `json:"supplierId"`
			StoreName   interface{} `json:"storeName"`
			StoreStatus interface{} `json:"storeStatus"`
		} `json:"storeInfo"`
	} `json:"info"`
}

type GetAnnoListRsp struct {
	BaseRsp
	Info struct {
		SuccessList struct {
			Data []struct {
				AnnouncementId string `json:"announcementId"`
			} `json:"data"`
			Mate struct {
				Count     int         `json:"count"`
				CustomObj interface{} `json:"customObj"`
			} `json:"mate"`
		} `json:"successList"`
	} `json:"info"`
}

type GetAnnoDetailRsp struct {
	BaseRsp
	Info struct {
		AnnouncementId string `json:"announcementId"`
		AttachmentList []struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"attachmentList"`
		Content    string `json:"content"`
		ImportType int    `json:"importantType"`
		StartTime  string `json:"startTime"`
		TagCode    string `json:"tagCode"`
		TagDesc    string `json:"tagDesc"`
		Title      string `json:"title"`
		TypeCode   string `json:"typeCode"`
		TypeDesc   string `json:"typeDesc"`
	}
}
