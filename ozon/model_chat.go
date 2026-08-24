package ozon

import "time"

type ChatSendMessageRsp struct {
	BaseRsp

	Result string `json:"result"`
}

type ChatSendFileRsp struct {
	BaseRsp

	Result string `json:"result"`
}

type ChatStartRsp struct {
	BaseRsp
	Result struct {
		ChatId string `json:"chat_id"`
	} `json:"result"`
}

type ChatListRsp struct {
	BaseRsp
	Filter struct {
		ChatStatus string `json:"chat_status"`
		UnreadOnly bool   `json:"unread_only"`
	} `json:"filter"`
	Limit  int64  `json:"limit"`
	Cursor string `json:"cursor"`
}

type ChatHistoryRsp struct {
	BaseRsp
	HasNext  bool                     `json:"has_next"`
	Messages []ChatHistoryMessageItem `json:"messages"`
}

type ChatHistoryMessageItem struct {
	Context struct {
		OrderNumber string `json:"order_number"`
		Sku         string `json:"sku"`
	} `json:"context"`
	CreatedAt           time.Time `json:"created_at"`
	Data                []string  `json:"data"`
	IsImage             bool      `json:"is_image"`
	IsRead              bool      `json:"is_read"`
	MessageId           string    `json:"message_id"`
	ModerateImageStatus string    `json:"moderate_image_status"`
	User                struct {
		Id   string `json:"id"`
		Type string `json:"type"`
	} `json:"user"`
}

type ChatReadRsp struct {
	BaseRsp
	UnreadCount int64 `json:"unread_count"`
}
