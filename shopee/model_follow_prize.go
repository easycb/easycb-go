package shopee

type AddFollowPrizeRsp struct {
	BaseRsp
	Response struct {
		CampaginId int64 `json:"campagin_id"`
	} `json:"response"`
}

type DeleteFollowPrizeRsp struct {
	BaseRsp
	Response struct {
		CampaignId int64 `json:"campaign_id"`
	} `json:"response"`
}

type EndFollowPrizeRsp struct {
	BaseRsp
	Response struct {
		CampaignId int64 `json:"campaign_id"`
	} `json:"response"`
}

type UpdateFollowPrizeRsp struct {
	BaseRsp
	Response struct {
		CampaignId int64 `json:"campaign_id"`
	} `json:"response"`
}

type GetFollowPrizeDetailRsp struct {
	BaseRsp
	Response struct {
		CampaignStatus  string  `json:"campaign_status"`
		CampaignId      int64   `json:"campaign_id"`
		UsageQuantity   int     `json:"usage_quantity"`
		StartTime       int64   `json:"start_time"`
		EndTime         int64   `json:"end_time"`
		MinSpend        int     `json:"min_spend"`
		RewardType      int     `json:"reward_type"`
		FollowPrizeName string  `json:"follow_prize_name"`
		Percentage      int     `json:"percentage"`
		MaxPrice        float64 `json:"max_price"`
	} `json:"response"`
}

type GetFollowPrizeListRsp struct {
	BaseRsp
	Response struct {
		More            bool `json:"more"`
		FollowPrizeList []struct {
			CampaignId      int64  `json:"campaign_id"`
			CampaignStatus  string `json:"campaign_status"`
			FollowPrizeName string `json:"follow_prize_name"`
			StartTime       int64  `json:"start_time"`
			EndTime         int64  `json:"end_time"`
			UsageQuantity   int    `json:"usage_quantity"`
			Claimed         int    `json:"claimed"`
		} `json:"follow_prize_list"`
	} `json:"response"`
}
