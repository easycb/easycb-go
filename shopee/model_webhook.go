package shopee

type BaseWebhook struct {
	Code      int         `json:"code"`
	Timestamp int64       `json:"timestamp"`
	ShopId    int64       `json:"shop_id"`
	Data      interface{} `json:"data"`
}

// Product

type WebhookReservedStockChangePush struct {
	BaseWebhook
	Data struct {
		ShopId        int   `json:"shop_id"`
		ItemId        int64 `json:"item_id"`
		VariationId   int64 `json:"variation_id"`
		ChangedValues []struct {
			Name string `json:"name"`
			Old  int    `json:"old"`
			New  int    `json:"new"`
		} `json:"changed_values"`
		PromotionType string `json:"promotion_type"`
		PromotionId   int64  `json:"promotion_id"`
		Action        string `json:"action"`
		Ordersn       string `json:"ordersn"`
		UpdateTime    int    `json:"update_time"`
	} `json:"data"`
}

type WebhookVideoUploadPush struct {
	BaseWebhook
	Data struct {
		VideoUploadId string `json:"video_upload_id"`
		Status        string `json:"status"`
		Message       string `json:"message"`
		VideoInfo     struct {
			VideoId  string `json:"video_id"`
			VideoUrl []struct {
				VideoUrlRegion string `json:"video_url_region"`
				VideoUrl       string `json:"video_url"`
			} `json:"video_url"`
			ThumbnailUrl []struct {
				ImageUrlRegion string `json:"image_url_region"`
				ImageUrl       string `json:"image_url"`
			} `json:"thumbnail_url"`
			Duration int `json:"duration"`
		} `json:"video_info"`
	} `json:"data"`
}

type WebhookBrandRegisterResult struct {
	BaseWebhook
	Data struct {
		ShopId        int `json:"shop_id"`
		RegisterBrand struct {
			BrandId   int    `json:"brand_id"`
			BrandName string `json:"brand_name"`
		} `json:"register_brand"`
		RegisterResult struct {
			Result string `json:"result"`
			Reason string `json:"reason"`
		} `json:"register_result"`
		CombinedBrand struct {
			BrandId   int    `json:"brand_id"`
			BrandName string `json:"brand_name"`
		} `json:"combined_brand"`
	} `json:"data"`
}

type WebhookViolationItemPushResult struct {
	BaseWebhook
	Data struct {
		ItemId            int64  `json:"item_id"`
		ItemName          string `json:"item_name"`
		ItemStatus        string `json:"item_status"`
		Deboost           bool   `json:"deboost"`
		ItemStatusDetails []struct {
			ViolationType   string `json:"violation_type"`
			ViolationReason string `json:"violation_reason"`
			Suggestion      string `json:"suggestion"`
			FixDeadlineTime int64  `json:"fix_deadline_time"`
			UpdateTime      int64  `json:"update_time"`
		} `json:"item_status_details"`
		DeboostedDetails []struct {
			ViolationType     string `json:"violation_type"`
			ViolationReason   string `json:"violation_reason"`
			Suggestion        string `json:"suggestion"`
			FixDeadlineTime   int    `json:"fix_deadline_time"`
			UpdateTime        int    `json:"update_time"`
			SuggestedCategory []struct {
				CategoryId   int    `json:"category_id"`
				CategoryName string `json:"category_name"`
			} `json:"suggested_category,omitempty"`
		} `json:"deboosted_details"`
	} `json:"data"`
}
