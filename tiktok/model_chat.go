package tiktok

type CreateConversationRsp struct {
	BaseRsp
	Data struct {
		ConversationId string `json:"conversation_id"`
	} `json:"data"`
}

type GetConversationsRsp struct {
	BaseRsp
	Data struct {
		Conversations []Conversation `json:"conversations"`
		NextPageToken string         `json:"next_page_token"`
	} `json:"data"`
}

type GetConversationRsp struct {
	BaseRsp
	Data struct {
		Conversations []Conversation `json:"conversations"`
		NextPageToken string         `json:"next_page_token"`
	} `json:"data"`
}

type Conversation struct {
	Id             string `json:"id"`
	CanSendMessage bool   `json:"can_send_message"`
	CreateTime     int64  `json:"create_time"`
	LatestMessage  struct {
		Content    string `json:"content"`
		CreateTime int    `json:"create_time"`
		Id         string `json:"id"`
		IsVisible  bool   `json:"is_visible"`
		Sender     struct {
			Avatar   string `json:"avatar"`
			ImUserId string `json:"im_user_id"`
			Nickname string `json:"nickname"`
			Role     string `json:"role"`
		} `json:"sender"`
		Type string `json:"type"`
	} `json:"latest_message"`
	ParticipantCount int `json:"participant_count"`
	Participants     []struct {
		Avatar   string `json:"avatar"`
		ImUserId string `json:"im_user_id"`
		Nickname string `json:"nickname"`
		Role     string `json:"role"`
		UserId   string `json:"user_id"`
	} `json:"participants"`
	UnreadCount int `json:"unread_count"`
}

type UploadBuyerMessagesImageRsp struct {
	BaseRsp
	Data struct {
		Height int    `json:"height"`
		Url    string `json:"url"`
		Width  int    `json:"width"`
	} `json:"data"`
}

type GetConversationMessagesRsp struct {
	BaseRsp
	Data struct {
		Messages           []Message `json:"messages"`
		NextPageToken      string    `json:"next_page_token"`
		UnsupportedMsgTips string    `json:"unsupported_msg_tips"`
	} `json:"data"`
}

type Message struct {
	Id         string `json:"id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
	IsVisible  bool   `json:"is_visible"`
	Sender     struct {
		Avatar   string `json:"avatar"`
		ImUserId string `json:"im_user_id"`
		Nickname string `json:"nickname"`
		Role     string `json:"role"`
	} `json:"sender"`
	Type string `json:"type"`
}

type SendMessageRsp struct {
	BaseRsp
	Data struct {
		MessageId string `json:"message_id"`
	} `json:"data"`
}

type ReadMessageRsp struct {
	BaseRsp
	Data struct{} `json:"data"`
}

type GetAgentSettingsRsp struct {
	BaseRsp
	Data struct {
		CanAcceptChat bool `json:"can_accept_chat"`
	} `json:"data"`
}

type UpdateAgentSettingsRsp struct {
	BaseRsp
}

type GetCustomerServicePerformanceRsp struct {
	BaseRsp
	Data struct {
		Performance struct {
			ResponsePercentage     string `json:"response_percentage"`
			ResponseTimeMins       string `json:"response_time_mins"`
			SatisfactionPercentage string `json:"satisfaction_percentage"`
			SupportSessionCount    int    `json:"support_session_count"`
		} `json:"performance"`
	} `json:"data"`
}
