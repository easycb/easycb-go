package lazada

type FreeShippingActivateRsp struct {
	Success   bool   `json:"success"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type FreeShippingAddSelectedProductSKURsp struct {
	Success   bool   `json:"success"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type FreeShippingCreateRsp struct {
	Data      int64  `json:"data"`
	Success   bool   `json:"success"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type FreeShippingDeactivateRsp struct {
	Success   bool   `json:"success"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type FreeShippingDeleteSelectedProductSKURsp struct {
	Success   bool   `json:"success"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type FreeShippingDeliveryOptionsQueryRsp struct {
	Data []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"data"`
	Success   bool   `json:"success"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type FreeShippingGetRsp struct {
	Data struct {
		Apply      string `json:"apply"`
		RegionType string `json:"region_type"`
		PromoTier  struct {
			Tiers []struct {
				Filter string `json:"filter"`
				Result string `json:"result"`
			} `json:"tiers"`
			DealCriteria string `json:"deal_criteria"`
			DiscountType string `json:"discount_type"`
		} `json:"promo_tier"`
		PromotionName   string `json:"promotion_name"`
		PlatformChannel string `json:"platform_channel"`
		TemplateType    string `json:"template_type"`
		Currency        string `json:"currency"`
		Id              int64  `json:"id"`
		BudgetType      string `json:"budget_type"`
		PeriodType      string `json:"period_type"`
		DeliveryOption  string `json:"delivery_option"`
		Status          string `json:"status"`
	} `json:"data"`
	Success   bool   `json:"success"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type FreeShippingListRsp struct {
	Data struct {
		DataList []struct {
			Apply       string `json:"apply"`
			BudgetValue string `json:"budget_value,omitempty"`
			RegionType  string `json:"region_type"`
			PromoTier   struct {
				Tiers []struct {
					Filter string `json:"filter"`
					Result string `json:"result"`
				} `json:"tiers"`
				DealCriteria string `json:"deal_criteria"`
				DiscountType string `json:"discount_type"`
			} `json:"promo_tier"`
			PromotionName   string   `json:"promotion_name"`
			UsedBudgetValue string   `json:"used_budget_value,omitempty"`
			PlatformChannel string   `json:"platform_channel"`
			TemplateType    string   `json:"template_type"`
			Currency        string   `json:"currency"`
			Id              int64    `json:"id"`
			BudgetType      string   `json:"budget_type"`
			PeriodType      string   `json:"period_type"`
			DeliveryOption  string   `json:"delivery_option"`
			Status          string   `json:"status"`
			RegionValue     []string `json:"region_value,omitempty"`
		} `json:"data_list"`
		Total    int `json:"total"`
		Current  int `json:"current"`
		PageSize int `json:"page_size"`
	} `json:"data"`
	Success   bool   `json:"success"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type FreeShippingRegionsQueryRsp struct {
	Data []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"data"`
	Success   bool   `json:"success"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type FreeShippingSelectedProductListRsp struct {
	Data struct {
		DataList []struct {
			SkuIds    []string `json:"sku_ids"`
			ProductId int64    `json:"product_id"`
		} `json:"data_list"`
		Total    int `json:"total"`
		Current  int `json:"current"`
		PageSize int `json:"page_size"`
	} `json:"data"`
	Success   bool   `json:"success"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type FreeShippingUpdateRsp struct {
	ErrorMsg  string `json:"error_msg"`
	Code      string `json:"code"`
	Data      int64  `json:"data"`
	Success   string `json:"success"`
	ErrorCode int    `json:"error_code"`
	RequestId string `json:"request_id"`
}
