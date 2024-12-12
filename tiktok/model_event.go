package tiktok

type GetShopWebhooksRsp struct {
	BaseRsp
	Data struct {
		TotalCount int `json:"total_count"`
		Webhooks   []struct {
			Address    string `json:"address"`
			CreateTime int64  `json:"create_time"`
			EventType  string `json:"event_type"`
			UpdateTime int64  `json:"update_time"`
		} `json:"webhooks"`
	} `json:"data"`
}

type UpdateShopWebhooksRsp struct {
	BaseRsp
}

type DeleteShopWebhooksRsp struct {
	BaseRsp
}
