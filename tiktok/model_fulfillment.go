package tiktok

type GetTrackingRsp struct {
	BaseRsp
	Data struct {
		Tracking []struct {
			Description      string `json:"description"`
			UpdateTimeMillis int64  `json:"update_time_millis"`
		} `json:"tracking"`
	} `json:"data"`
}
