package lazada

type GetStoreCustomPageRsp struct {
	Code string `json:"code"`
	Data struct {
		Result struct {
			PageList []struct {
				DecoratePageUrl        string `json:"decorate_page_url"`
				WirelessPagePreviewUrl string `json:"wireless_page_preview_url"`
				WirelessEndTime        string `json:"wireless_end_time"`
				TimedPublishTime       string `json:"timed_publish_time"`
				RelatePageId           int    `json:"relate_page_id"`
				ClientType             string `json:"client_type"`
				PcEndTime              string `json:"pc_end_time"`
				PcPagePreviewUrl       string `json:"pc_page_preview_url"`
				PageId                 int    `json:"page_id"`
				Path                   string `json:"path"`
				WirelessPageViewUrl    string `json:"wireless_page_view_url"`
				PageViewUrl            string `json:"page_view_url"`
				LastEditTime           string `json:"last_edit_time"`
				PublishTime            string `json:"publish_time"`
				QrUrl                  string `json:"qr_url"`
				PageName               string `json:"page_name"`
				StatusKey              string `json:"status_key"`
			} `json:"page_list"`
			PageInfo struct {
				TotalCount  string `json:"total_count"`
				CurrentPage string `json:"current_page"`
			} `json:"page_info"`
		} `json:"result"`
		ErrorMessage string `json:"error_message"`
		Success      bool   `json:"success"`
		Error        string `json:"error"`
	} `json:"data"`
	RequestId string `json:"request_id"`
}
