package ozon

import "time"

type ReviewCommentCreateRsp struct {
	BaseRsp
	CommentId string `json:"comment_id"`
}

type ReviewCommentDeleteRsp struct {
	BaseRsp
}

type ReviewCommentListRsp struct {
	BaseRsp
	Comments []ReviewCommentListComment `json:"comments"`
	Offset   int                        `json:"offset"`
}

type ReviewCommentListComment struct {
	Id              string    `json:"id"`
	IsOfficial      bool      `json:"is_official"`
	IsOwner         bool      `json:"is_owner"`
	ParentCommentId string    `json:"parent_comment_id"`
	PublishedAt     time.Time `json:"published_at"`
	Text            string    `json:"text"`
}

type ReviewChangeStatusRsp struct {
	BaseRsp
}

type ReviewCountRsp struct {
	BaseRsp

	Processed   int `json:"processed"`
	Total       int `json:"total"`
	Unprocessed int `json:"unprocessed"`
}

type ReviewInfoRsp struct {
	BaseRsp

	CommentsAmount      int    `json:"comments_amount"`
	DislikesAmount      int    `json:"dislikes_amount"`
	Id                  string `json:"id"`
	IsRatingParticipant bool   `json:"is_rating_participant"`
	LikesAmount         int    `json:"likes_amount"`
	OrderStatus         string `json:"order_status"`
	Photos              []struct {
		Height int    `json:"height"`
		Url    string `json:"url"`
		Width  int    `json:"width"`
	} `json:"photos"`
	PhotosAmount int       `json:"photos_amount"`
	PublishedAt  time.Time `json:"published_at"`
	Rating       int       `json:"rating"`
	Sku          int       `json:"sku"`
	Status       string    `json:"status"`
	Text         string    `json:"text"`
	Videos       []struct {
		Height               int    `json:"height"`
		PreviewUrl           string `json:"preview_url"`
		ShortVideoPreviewUrl string `json:"short_video_preview_url"`
		Url                  string `json:"url"`
		Width                int    `json:"width"`
	} `json:"videos"`
	VideosAmount int `json:"videos_amount"`
}

type ReviewListRsp struct {
	BaseRsp
	Reviews []ReviewListItem `json:"reviews"`
	HasNext bool             `json:"has_next"`
	LastId  string           `json:"last_id"`
}

type ReviewListItem struct {
	Id                  string    `json:"id"`
	Sku                 string    `json:"sku"`
	Text                string    `json:"text"`
	PublishedAt         time.Time `json:"published_at"`
	Rating              int       `json:"rating"`
	Status              string    `json:"status"`
	CommentsAmount      int       `json:"comments_amount"`
	PhotosAmount        int       `json:"photos_amount"`
	VideosAmount        int       `json:"videos_amount"`
	OrderStatus         string    `json:"order_status"`
	IsRatingParticipant bool      `json:"is_rating_participant"`
}
