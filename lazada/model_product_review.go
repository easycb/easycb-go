package lazada

type GetHistoryReviewIdListRsp struct {
	ErrorMsg string `json:"error_msg"`
	Code     string `json:"code"`
	Data     struct {
		Current  int   `json:"current"`
		Total    int   `json:"total"`
		IdList   []int `json:"id_list"`
		PageSize int   `json:"page_size"`
	} `json:"data"`
	Success   bool   `json:"success"`
	ErrorCode string `json:"error_code"`
	RequestId string `json:"request_id"`
}

type GetReviewListByIdListRsp struct {
	ErrorMsg string `json:"error_msg"`
	Code     string `json:"code"`
	Data     struct {
		OutdatedReviews []int `json:"outdated_reviews"`
		ReviewList      []struct {
			ReviewImages  []interface{} `json:"review_images"`
			CanReply      bool          `json:"can_reply"`
			CreateTime    int64         `json:"create_time"`
			SubmitTime    int64         `json:"submit_time"`
			ReviewContent string        `json:"review_content"`
			Ratings       struct {
				SellerRating    int `json:"seller_rating"`
				OverallRating   int `json:"overall_rating"`
				LogisticsRating int `json:"logistics_rating"`
				ProductRating   int `json:"product_rating"`
			} `json:"ratings"`
			ProductId    int64 `json:"product_id"`
			ReviewVideos []struct {
				VideoUrl      string `json:"video_url"`
				VideoCoverUrl string `json:"video_cover_url"`
			} `json:"review_videos"`
			Id          int    `json:"id"`
			SellerReply string `json:"seller_reply"`
			OrderId     int64  `json:"order_id"`
			ReviewType  string `json:"review_type"`
		} `json:"review_list"`
	} `json:"data"`
	Success   bool   `json:"success"`
	ErrorCode string `json:"error_code"`
	RequestId string `json:"request_id"`
}

type SubmitSellerReplyRsp struct {
	ErrorMsg  string `json:"error_msg"`
	Code      string `json:"code"`
	Data      bool   `json:"data"`
	Success   bool   `json:"success"`
	ErrorCode string `json:"error_code"`
	RequestId string `json:"request_id"`
}
