package shopee

type InitVideoUploadRsp struct {
	BaseRsp
	Response struct {
		VideoUploadId string `json:"video_upload_id"`
	} `json:"response"`
}

type UploadVideoPartRsp struct {
	BaseRsp
}

type CompleteVideoUploadRsp struct {
	BaseRsp
}

type GetVideoUploadResultRsp struct {
	BaseRsp
	Response struct {
		Status    string `json:"status"`
		Message   string `json:"message"`
		VideoInfo struct {
			VideoUrlList []struct {
				VideoUrlRegion string `json:"video_url_region"`
				VideoUrl       string `json:"video_url"`
			} `json:"video_url_list"`
			ThumbnailUrlList []struct {
				ImageUrlRegion string `json:"image_url_region"`
				ImageUrl       string `json:"image_url"`
			} `json:"thumbnail_url_list"`
			Duration int `json:"duration"`
		} `json:"video_info"`
	} `json:"response"`
}

type CancelVideoUploadRsp struct {
	BaseRsp
}

type UploadImageRsp struct {
	BaseRsp
}
