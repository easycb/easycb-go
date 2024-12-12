package shopee

type SetAppPushConfigRsp struct {
	BaseRsp
	Response struct {
		Result string `json:"result"`
	} `json:"response"`
}

type GetAppPushConfigRsp struct {
	BaseRsp
	Response struct {
		CallbackUrl       string  `json:"callback_url"`
		LivePushStatus    string  `json:"live_push_status"`
		SuspendedTime     int64   `json:"suspended_time"`
		BlockedShopId     []int64 `json:"blocked_shop_id"`
		PushConfigOnList  []int   `json:"push_config_on_list"`
		PushConfigOffList []int   `json:"push_config_off_list"`
	} `json:"response"`
}

type GetLostPushMessageRsp struct {
	BaseRsp
	Warning  string `json:"warning"`
	Response struct {
		PushMessageList []struct {
			ShopId    int64  `json:"shop_id"`
			Code      int    `json:"code"`
			Timestamp int64  `json:"timestamp"`
			Data      string `json:"data"`
		} `json:"push_message_list"`
		HasNextPage   bool `json:"has_next_page"`
		LastMessageId int  `json:"last_message_id"`
	} `json:"response"`
}

type ConfirmConsumedLostPushRsp struct {
	BaseRsp
}
