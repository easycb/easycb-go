package lazada

type SellerVoucheDeleteSelectedProductSKURsp struct {
	ErrorMsg  string `json:"error_msg"`
	Code      string `json:"code"`
	Success   string `json:"success"`
	ErrorCode int    `json:"error_code"`
	RequestId string `json:"request_id"`
}

type SellerVoucherActivateRsp struct {
	Success   bool   `json:"success"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type SellerVoucherAddSelectedProductSKURsp struct {
	ErrorMsg string `json:"error_msg"`
	Code     string `json:"code"`
	Data     struct {
		SkuId string `json:"sku id"`
	} `json:"data"`
	Success   bool   `json:"success"`
	ErrorCode int    `json:"error_code"`
	RequestId string `json:"request_id"`
}

type SellerVoucherCreateRsp struct {
	Data      int64  `json:"data"`
	Success   bool   `json:"success"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type SellerVoucherDeactivateRsp struct {
	Success   bool   `json:"success"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type SellerVoucherDetailQueryRsp struct {
	Data struct {
		PeriodEndTime         int64  `json:"period_end_time"`
		CriteriaOverMoney     string `json:"criteria_over_money"`
		Apply                 string `json:"apply"`
		VoucherName           string `json:"voucher_name"`
		OfferingMoneyValueOff string `json:"offering_money_value_off"`
		OrderUsedBudget       int    `json:"order_used_budget"`
		PeriodStartTime       int64  `json:"period_start_time"`
		DisplayArea           string `json:"display_area"`
		VoucherType           string `json:"voucher_type"`
		Limit                 int    `json:"limit"`
		CollectStart          int64  `json:"collect_start"`
		VoucherDiscountType   string `json:"voucher_discount_type"`
		Currency              string `json:"currency"`
		Id                    int64  `json:"id"`
		Issued                int    `json:"issued"`
		Status                string `json:"status"`
	} `json:"data"`
	Success   bool   `json:"success"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type SellerVoucherListRsp struct {
	ErrorMsg string `json:"error_msg"`
	Code     string `json:"code"`
	Data     struct {
		DataList []struct {
			PeriodEndTime                 int64  `json:"period_end_time"`
			MaxDiscountOfferingMoneyValue string `json:"max_discount_offering_money_value"`
			CriteriaOverMoney             string `json:"criteria_over_money"`
			Apply                         string `json:"apply"`
			VoucherName                   string `json:"voucher_name"`
			VoucherCode                   string `json:"voucher_code"`
			OfferingMoneyValueOff         string `json:"offering_money_value_off"`
			OrderUsedBudget               int    `json:"order_used_budget"`
			OfferingPercentageDiscountOff string `json:"offering_percentage_discount_off"`
			PeriodStartTime               int64  `json:"period_start_time"`
			DisplayArea                   string `json:"display_area"`
			VoucherType                   string `json:"voucher_type"`
			Limit                         int    `json:"limit"`
			CollectStart                  int64  `json:"collect_start"`
			VoucherDiscountType           string `json:"voucher_discount_type"`
			Currency                      string `json:"currency"`
			Id                            int    `json:"id"`
			Issued                        int    `json:"issued"`
			Status                        string `json:"status"`
		} `json:"data_list"`
		Total    int `json:"total"`
		Current  int `json:"current"`
		PageSize int `json:"page_size"`
	} `json:"data"`
	Success   bool   `json:"success"`
	ErrorCode string `json:"error_code"`
	RequestId string `json:"request_id"`
}

type SellerVoucherSelectedProductListRsp struct {
	ErrorMsg string `json:"error_msg"`
	Code     string `json:"code"`
	Data     struct {
		DataList []struct {
			SkuIds    []int64 `json:"sku_ids"`
			ProductId int64   `json:"product_id"`
		} `json:"data_list"`
		Total    string `json:"total"`
		Current  string `json:"current"`
		PageSize string `json:"page_size"`
	} `json:"data"`
	Success   bool   `json:"success"`
	ErrorCode int    `json:"error_code"`
	RequestId string `json:"request_id"`
}

type SellerVoucherUpdateRsp struct {
	ErrorMsg  string `json:"error_msg"`
	Code      string `json:"code"`
	Data      int64  `json:"data"`
	Success   bool   `json:"success"`
	ErrorCode int    `json:"error_code"`
	RequestId string `json:"request_id"`
}
