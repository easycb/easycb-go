package lazada

type ServiceMarketAppKeyOrderQueryRsp struct {
	Result struct {
		Data struct {
			TotalItem        string `json:"totalItem"`
			ArticleBizOrders []struct {
				OrderCycleStart string `json:"orderCycleStart"`
				RefundFee       string `json:"refundFee"`
				ArticleItemName string `json:"articleItemName"`
				BizType         string `json:"bizType"`
				ArticleName     string `json:"articleName"`
				TotalPayFee     string `json:"totalPayFee"`
				OrderId         string `json:"orderId"`
				OrderCycleEnd   string `json:"orderCycleEnd"`
				ItemCode        string `json:"itemCode"`
				Fee             string `json:"fee"`
				UserId          string `json:"userId"`
				Nick            string `json:"nick"`
				ActivityCode    string `json:"activityCode"`
				ItemName        string `json:"itemName"`
				OrderCycle      string `json:"orderCycle"`
				BizOrderId      string `json:"bizOrderId"`
				PromFee         string `json:"promFee"`
				Create          string `json:"create"`
				ArticleCode     string `json:"articleCode"`
			} `json:"articleBizOrders"`
		} `json:"data"`
		Success    bool   `json:"success"`
		ResultCode string `json:"resultCode"`
		Remark     string `json:"remark"`
	} `json:"result"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type ServiceMarketAppKeySubQueryRsp struct {
	Result struct {
		Data []struct {
			Nick         string `json:"nick"`
			ItemCode     string `json:"item_code"`
			ExpireNotice bool   `json:"expire_notice"`
			EndTime      int64  `json:"end_time"`
			ArticleName  string `json:"article_name"`
			ItemName     string `json:"item_name"`
			Autosub      bool   `json:"autosub"`
			ArticleCode  string `json:"article_code"`
			Status       int    `json:"status"`
		} `json:"data"`
		Success bool `json:"success"`
	} `json:"result"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}
