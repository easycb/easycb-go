package tiktok

type GetMessageTemplatesRsp struct {
	BaseRsp
	Data struct {
		MessageTemplates []struct {
			CouponCardRules struct {
				CouponType []string `json:"coupon_type"`
				MaxCount   int      `json:"max_count"`
				MinCount   int      `json:"min_count"`
			} `json:"coupon_card_rules"`
			Id               string `json:"id"`
			MessageBody      string `json:"message_body"`
			MessageTitle     string `json:"message_title"`
			ProductCardRules struct {
				MaxCount int `json:"max_count"`
				MinCount int `json:"min_count"`
			} `json:"product_card_rules"`
		} `json:"message_templates"`
	} `json:"data"`
}

type CreateEngagementTaskRsp struct {
	BaseRsp
	Data struct {
		TaskId string `json:"task_id"`
	} `json:"data"`
}

type SendEngagementMessageRsp struct {
	BaseRsp
	Data struct {
		Errors []struct {
			Code   int `json:"code"`
			Detail struct {
				BuyerEmail string `json:"buyer_email"`
			} `json:"detail"`
			Message string `json:"message"`
		} `json:"errors"`
	} `json:"data"`
}

type GetTaskPerformancesRsp struct {
	BaseRsp
	Data struct {
		TaskPerformances []struct {
			ClaimedCouponsCount int    `json:"claimed_coupons_count"`
			GmvAmount           string `json:"gmv_amount"`
			Id                  string `json:"id"`
			OrderCount          int    `json:"order_count"`
			ReadRecipientCount  int    `json:"read_recipient_count"`
			SentRecipientCount  int    `json:"sent_recipient_count"`
			Status              string `json:"status"`
		} `json:"task_performances"`
	} `json:"data"`
}
