package lazada

type GetConversationListRsp struct {
	Code string `json:"code"`
	Data struct {
		SessionList   []ConversationItem `json:"session_list"`
		NextStartTime int64              `json:"next_start_time"`
		HasMore       bool               `json:"has_more"`
		LastSessionId string             `json:"last_session_id"`
	} `json:"data"`
	Success    bool   `json:"success"`
	ErrCode    string `json:"err_code"`
	RequestId  string `json:"request_id"`
	ErrMessage string `json:"err_message"`
}

type GetOneConversationRsp struct {
	Code       string           `json:"code"`
	Data       ConversationItem `json:"data"`
	Success    bool             `json:"success"`
	ErrCode    string           `json:"err_code"`
	RequestId  string           `json:"request_id"`
	ErrMessage string           `json:"err_message"`
}

type ConversationItem struct {
	Summary         string   `json:"summary"`
	UnreadCount     int      `json:"unread_count"`
	LastMessageId   string   `json:"last_message_id"`
	HeadUrl         string   `json:"head_url"`
	SelfPosition    int64    `json:"self_position"`
	LastMessageTime int64    `json:"last_message_time"`
	SiteId          string   `json:"site_id"`
	SessionId       string   `json:"session_id"`
	Title           string   `json:"title"`
	BuyerId         int      `json:"buyer_id"`
	ToPosition      int64    `json:"to_position"`
	Tags            []string `json:"tags"`
}

type ReadConversationRsp struct {
	Code       string `json:"code"`
	Success    bool   `json:"success"`
	ErrCode    string `json:"err_code"`
	RequestId  string `json:"request_id"`
	ErrMessage string `json:"err_message"`
}

type GetMessagesRsp struct {
	Code string `json:"code"`
	Data struct {
		LastMessageId string        `json:"last_message_id"`
		MessageList   []MessageItem `json:"message_list"`
		NextStartTime int64         `json:"next_start_time"`
		HasMore       bool          `json:"has_more"`
	} `json:"data"`
	Success    bool   `json:"success"`
	ErrCode    string `json:"err_code"`
	RequestId  string `json:"request_id"`
	ErrMessage string `json:"err_message"`
}

type MessageItem struct {
	FromAccountType int    `json:"from_account_type"`
	ProcessMsg      string `json:"process_msg"`
	SessionId       string `json:"session_id"`
	MessageId       string `json:"message_id"`
	Type            int    `json:"type"`
	Content         string `json:"content"`
	ToAccountId     string `json:"to_account_id"`
	SendTime        int64  `json:"send_time"`
	AutoReply       bool   `json:"auto_reply"`
	ToAccountType   int    `json:"to_account_type"`
	SiteId          string `json:"site_id"`
	TemplateId      int    `json:"template_id"`
	FromAccountId   string `json:"from_account_id"`
	Status          int    `json:"status"`
}

type SendMessageRsp struct {
	Code string `json:"code"`
	Data struct {
		MessageId   string `json:"message_id"`
		TemplateId  int    `json:"template_id"`
		CurrentTime int64  `json:"current_time"`
	} `json:"data"`
	Success    bool   `json:"success"`
	ErrCode    string `json:"err_code"`
	RequestId  string `json:"request_id"`
	ErrMessage string `json:"err_message"`
}

type RecallMessageRsp struct {
	Code       string `json:"code"`
	RequestId  string `json:"request_id"`
	ErrMessage string `json:"err_message"`
	ErrorCode  string `json:"error_code"`
	Success    bool   `json:"success"`
}
