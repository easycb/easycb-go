package lazada

type BatchQueryFollowStatusRsp struct {
	Result struct {
		Result  []interface{} `json:"result"`
		Success bool          `json:"success"`
		Error   struct {
		} `json:"error"`
	} `json:"result"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type GetPickUpStoreListRsp struct {
	Result struct {
		BizExtMap struct {
		} `json:"biz_ext_map"`
		Headers struct {
		} `json:"headers"`
		MsgCode        string `json:"msg_code"`
		HttpStatusCode int    `json:"http_status_code"`
		Success        bool   `json:"success"`
		MsgInfo        string `json:"msg_info"`
		Model          struct {
		} `json:"model"`
		MappingCode string `json:"mapping_code"`
	} `json:"result"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type GetSellerRsp struct {
	Data struct {
		NameCompany         string `json:"name_company"`
		Name                string `json:"name"`
		Verified            bool   `json:"verified"`
		SellerId            int64  `json:"seller_id"`
		Email               string `json:"email"`
		ShortCode           string `json:"short_code"`
		Cb                  bool   `json:"cb"`
		Status              string `json:"status"`
		MarketplaceEaseMode bool   `json:"marketplaceEaseMode"`
	} `json:"data"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type GetSellerMetricsByIdRsp struct {
	Code string `json:"code"`
	Data struct {
		MainCategoryName     string `json:"main_category_name"`
		ShipOnTime           string `json:"ship_on_time"`
		PositiveSellerRating string `json:"positive_seller_rating"`
		ResponseTime         string `json:"response_time"`
		SellerId             int64  `json:"seller_id"`
		ResponseRate         string `json:"response_rate"`
		MainCategoryId       int    `json:"main_category_id"`
	} `json:"data"`
	RequestId string `json:"request_id"`
}

type GetSellerPerformanceRsp struct {
	Code string `json:"code"`
	Data struct {
		MainCategoryName string `json:"main_category_name"`
		Indicators       []struct {
			ActionUrl       string `json:"action_url"`
			Score           int    `json:"score"`
			ScoreFormat     string `json:"score_format"`
			FormattedScore  string `json:"formatted_score"`
			Name            string `json:"name"`
			Tip             string `json:"tip"`
			Type            string `json:"type"`
			FormattedTarget string `json:"formatted_target"`
			Target          int    `json:"target"`
			TargetFormat    string `json:"target_format"`
			TargetRespected bool   `json:"target_respected"`
		} `json:"indicators"`
		SellerId       int64 `json:"seller_id"`
		MainCategoryId int   `json:"main_category_id"`
	} `json:"data"`
	Success   bool   `json:"success"`
	ErrorCode string `json:"error_code"`
	RequestId string `json:"request_id"`
}

type GetWarehouseBySellerIdRsp struct {
	Result struct {
		NotSuccess bool `json:"not_success"`
		Success    bool `json:"success"`
		Module     struct {
		} `json:"module"`
		ErrorCode string `json:"error_code"`
		Repeated  bool   `json:"repeated"`
		Retry     bool   `json:"retry"`
	} `json:"result"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type QueryWarehouseDetailInfoBySellerIdRsp struct {
	Result struct {
		NotSuccess bool `json:"not_success"`
		Success    bool `json:"success"`
		Module     struct {
			Country        string `json:"country"`
			DefaultAddress bool   `json:"default_address"`
			Province       string `json:"province"`
			City           string `json:"city"`
			DetailAddress  string `json:"detail_address"`
			WarehouseCode  string `json:"warehouse_code"`
			District       string `json:"district"`
			PostCode       string `json:"post_code"`
			Name           string `json:"name"`
			Status         string `json:"status"`
		} `json:"module"`
		ErrorCode string `json:"error_code"`
		Repeated  bool   `json:"repeated"`
		ClassName string `json:"class_name"`
		Retry     bool   `json:"retry"`
	} `json:"result"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type SellerCenterMsgListRsp struct {
	Result struct {
		Data struct {
			PageInfo struct {
				Current  int `json:"current"`
				Total    int `json:"total"`
				PageSize int `json:"pageSize"`
			} `json:"pageInfo"`
			DataSource []struct {
				MessageContent struct {
					AppLink      string `json:"appLink"`
					WebLink      string `json:"webLink"`
					Description  string `json:"description"`
					Title        string `json:"title"`
					CategoryName string `json:"categoryName"`
					Picture      string `json:"picture"`
				} `json:"message_content"`
				Id   string `json:"id"`
				Time string `json:"time"`
			} `json:"dataSource"`
		} `json:"data"`
		Success   interface{} `json:"success"`
		ErrorCode string      `json:"errorCode"`
		Type      string      `json:"type"`
		Error     string      `json:"error"`
	} `json:"result"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type SellerPolicyFetchRsp struct {
	Code      string `json:"code"`
	Data      string `json:"data"`
	Success   string `json:"success"`
	RequestId string `json:"request_id"`
	// todo 待补充具体数据
}

type SynchronizeSellerItemArConfigRsp struct {
	Code      string `json:"code"`
	Success   bool   `json:"success"`
	ErrorCode string `json:"errorCode"`
	Model     struct {
		Uid string `json:"uid"`
	} `json:"model"`
	RequestId string `json:"request_id"`
	ErrorMsg  string `json:"errorMsg"`
}

type GetCountryInfoRsp struct {
	Code string `json:"code"`
	Data []struct {
		Label string `json:"label"`
		Value string `json:"value"`
	} `json:"data"`
	Success   string `json:"success"`
	RequestId string `json:"request_id"`
}

type GetSellerRegisterInfoRsp struct {
	Code string `json:"code"`
	Data []struct {
		BaseInfoList []struct {
			RegisterCountry string `json:"registerCountry"`
			Phone           string `json:"phone"`
			ShopName        string `json:"shopName"`
			ReqNo           string `json:"reqNo"`
			Email           string `json:"email"`
			Status          string `json:"status"`
		} `json:"baseInfoList"`
		CompanyName   string `json:"companyName"`
		LicenseNumber string `json:"licenseNumber"`
	} `json:"data"`
	Success   string `json:"success"`
	RequestId string `json:"request_id"`
}

type GetSubAddressRsp struct {
	Code string `json:"code"`
	Data []struct {
		Label string `json:"label"`
		Value string `json:"value"`
	} `json:"data"`
	Success   bool   `json:"success"`
	RequestId string `json:"request_id"`
}

type PaymentBindingRsp struct {
	Code string `json:"code"`
	Data []struct {
		Result    bool   `json:"result"`
		Reason    string `json:"reason"`
		ShortCode string `json:"shortCode"`
	} `json:"data"`
	Success   bool   `json:"success"`
	RequestId string `json:"request_id"`
}

type QueryBuyboxHuntingInfoRsp struct {
	Result struct {
		Data struct {
			ItemId    string `json:"itemId"`
			IsValid   string `json:"isValid"`
			Venture   string `json:"venture"`
			SkuId     string `json:"skuId"`
			PriceRank string `json:"priceRank"`
		} `json:"data"`
		RetSuccess bool `json:"retSuccess"`
	} `json:"result"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type SaveSellerWarehouseInfoRsp struct {
	Result struct {
		NotSuccess bool `json:"not_success"`
		Success    bool `json:"success"`
		Module     bool `json:"module"`
		Repeated   bool `json:"repeated"`
		Retry      bool `json:"retry"`
	} `json:"result"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type SellerFieldVerifyRsp struct {
	Code string `json:"code"`
	Data []struct {
		Result   string `json:"result"`
		ErrorMsg string `json:"error_msg"`
		Name     string `json:"name"`
		ErrCode  string `json:"err_code"`
	} `json:"data"`
	Success   string `json:"success"`
	RequestId string `json:"request_id"`
}
