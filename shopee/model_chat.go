package shopee

type GetMessagesRsp struct {
	BaseRsp
	Response struct {
		Messages   []MessageItem `json:"messages"`
		PageResult struct {
			NextOffset string `json:"next_offset"`
			PageSize   int    `json:"page_size"`
		} `json:"page_result"`
	} `json:"response"`
}

type MessageItem struct {
	MessageId        string      `json:"message_id"`
	RequestId        string      `json:"request_id"`
	FromId           int         `json:"from_id"`
	ToId             int         `json:"to_id"`
	FromShopId       int64       `json:"from_shop_id"`
	ToShopId         int64       `json:"to_shop_id"`
	MessageType      string      `json:"message_type"`
	Content          interface{} `json:"content"`
	ConversationId   string      `json:"conversation_id"`
	CreatedTimestamp int64       `json:"created_timestamp"`
	Region           string      `json:"region"`
	Status           string      `json:"status"`
	MessageOption    int         `json:"message_option"`
	Source           string      `json:"source"`
	SourceContent    interface{} `json:"source_content"`
	QuotedMsg        interface{} `json:"quoted_msg"`
}

type SendMessageRsp struct {
	BaseRsp
	Response struct {
		MessageId        string                 `json:"message_id"`
		RequestId        string                 `json:"request_id"`
		ToId             int                    `json:"to_id"`
		MessageType      string                 `json:"message_type"`
		Content          map[string]interface{} `json:"content"`
		ConversationId   int64                  `json:"conversation_id"`
		CreatedTimestamp int64                  `json:"created_timestamp"`
		MessageOption    int                    `json:"message_option"`
		SourceContent    struct{}               `json:"source_content"`
	} `json:"response"`
}

type ConversationItem struct {
	ConversationId           string      `json:"conversation_id"`
	ToId                     int         `json:"to_id"`
	ToName                   string      `json:"to_name"`
	ToAvatar                 string      `json:"to_avatar"`
	ShopId                   int64       `json:"shop_id"`
	UnreadCount              int         `json:"unread_count"`
	Pinned                   bool        `json:"pinned"`
	LastReadMessageId        string      `json:"last_read_message_id"`
	LatestMessageId          string      `json:"latest_message_id"`
	LatestMessageType        string      `json:"latest_message_type"`
	LatestMessageContent     interface{} `json:"latest_message_content"`
	LatestMessageFromId      int         `json:"latest_message_from_id"`
	LastMessageTimestamp     int64       `json:"last_message_timestamp"`
	LastMessageOption        int         `json:"last_message_option"`
	MaxGeneralOptionHideTime string      `json:"max_general_option_hide_time"`
	Mute                     bool        `json:"mute"`
}

type GetConversationListRsp struct {
	BaseRsp
	Response struct {
		PageResult struct {
			PageSize   int `json:"page_size"`
			NextCursor struct {
				NextMessageTimeNano string `json:"next_message_time_nano"`
				ConversationId      string `json:"conversation_id"`
			} `json:"next_cursor"`
			More bool `json:"more"`
		} `json:"page_result"`
		Conversations []ConversationItem `json:"conversations"`
	} `json:"response"`
}

type GetOneConversationRsp struct {
	BaseRsp
	Response ConversationItem `json:"response"`
}

type ReadConversationRsp struct {
	BaseRsp
}

type SellerChatUploadImageRsp struct {
	BaseRsp
	Response struct {
		Url           string `json:"url"`
		Thumbnail     string `json:"thumbnail"`
		ThumbnailHash string `json:"thumbnail_hash"`
		FileServerId  int    `json:"file_server_id"`
		ThumbHeight   int    `json:"thumb_height"`
		ThumbWidth    int    `json:"thumb_width"`
		UrlHash       string `json:"url_hash"`
	} `json:"response"`
}

type GetOfferToggleStatusRsp struct {
	BaseRsp
	Response struct {
		ShopId          int64  `json:"shop_id"`
		MakeOfferStatus string `json:"make_offer_status"`
	} `json:"response"`
}

type GetUnreadConversationCountRsp struct {
	BaseRsp
	Response struct {
		TotalUnreadCount int `json:"total_unread_count"`
	} `json:"response"`
}

type DeleteConversationRsp struct {
	BaseRsp
}

type PinConversationRsp struct {
	BaseRsp
}

type UnpinConversationRsp struct {
	BaseRsp
}
