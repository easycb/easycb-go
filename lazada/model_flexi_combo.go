package lazada

type ActivateFlexiComboRsp struct {
	Success   bool   `json:"success"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type AddFlexiComboProductsRsp struct {
	Success   bool   `json:"success"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type CreateFlexiComboRsp struct {
	Data      int64  `json:"data"`
	Success   bool   `json:"success"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type DeactivateFlexiComboRsp struct {
	Success   bool   `json:"success"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type DeleteFlexiComboProductsRsp struct {
	Success   bool   `json:"success"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type GetFlexiComboDetailsRsp struct {
	Data struct {
		Stackable        bool     `json:"stackable"`
		Apply            string   `json:"apply"`
		EndTime          int64    `json:"end_time"`
		DiscountValue    []string `json:"discount_value"`
		Type             string   `json:"type"`
		DiscountType     string   `json:"discount_type"`
		OrderUsedNumbers int      `json:"order_used_numbers"`
		StartTime        int64    `json:"start_time"`
		Name             string   `json:"name"`
		Id               int64    `json:"id"`
		CriteriaType     string   `json:"criteria_type"`
		CriteriaValue    []string `json:"criteria_value"`
		OrderNumbers     int      `json:"order_numbers"`
		Status           string   `json:"status"`
	} `json:"data"`
	Success   bool   `json:"success"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type ListFlexiComboRsp struct {
	Data struct {
		DataList []struct {
			Stackable        bool     `json:"stackable"`
			Apply            string   `json:"apply"`
			EndTime          int64    `json:"end_time"`
			DiscountValue    []string `json:"discount_value,omitempty"`
			DiscountType     string   `json:"discount_type"`
			Type             string   `json:"type"`
			StartTime        int64    `json:"start_time"`
			OrderUsedNumbers int      `json:"order_used_numbers"`
			Name             string   `json:"name"`
			Id               int64    `json:"id"`
			CriteriaType     string   `json:"criteria_type"`
			CriteriaValue    []string `json:"criteria_value"`
			OrderNumbers     int      `json:"order_numbers"`
			Status           string   `json:"status"`
			GiftSkus         []struct {
				ProductId int64 `json:"product_id"`
				SkuId     int64 `json:"sku_id"`
			} `json:"gift_skus,omitempty"`
		} `json:"data_list"`
		Total    int `json:"total"`
		Current  int `json:"current"`
		PageSize int `json:"page_size"`
	} `json:"data"`
	Success   bool   `json:"success"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type ListFlexiComboProductsRsp struct {
	Data struct {
		DataList []string `json:"data_list"`
		Total    int      `json:"total"`
		Current  int      `json:"current"`
		PageSize int      `json:"page_size"`
	} `json:"data"`
	Success   bool   `json:"success"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type UpdateFlexiComboRsp struct {
	ErrorMsg  string `json:"error_msg"`
	Code      string `json:"code"`
	Success   bool   `json:"success"`
	ErrorCode string `json:"error_code"`
	RequestId string `json:"request_id"`
}
