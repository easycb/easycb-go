package shopee

import (
	"github.com/easycb/easycb-go"
)

func (c *Client) GetConversationList(query easycb.AnyMap) (*GetConversationListRsp, error) {
	var result GetConversationListRsp
	err := c.doRequest("GET", "/api/v2/sellerchat/get_conversation_list", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetConversationDetail(query easycb.AnyMap) (*GetOneConversationRsp, error) {
	var result GetOneConversationRsp
	err := c.doRequest("GET", "/api/v2/sellerchat/get_one_conversation", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ReadConversation(body easycb.AnyMap) (*ReadConversationRsp, error) {
	var result ReadConversationRsp
	err := c.doRequest("POST", "/api/v2/sellerchat/read_conversation", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetMessages(query easycb.AnyMap) (*GetMessagesRsp, error) {
	var result GetMessagesRsp
	err := c.doRequest("GET", "/api/v2/sellerchat/get_message", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SendMessage(body easycb.AnyMap) (*SendMessageRsp, error) {
	var result SendMessageRsp
	err := c.doRequest("POST", "/api/v2/sellerchat/send_message", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetOfferToggleStatus() (*GetOfferToggleStatusRsp, error) {
	var result GetOfferToggleStatusRsp
	err := c.doRequest("GET", "/api/v2/sellerchat/get_offer_toggle_status", nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SetOfferToggleStatus(body easycb.AnyMap) (*GetOfferToggleStatusRsp, error) {
	var result GetOfferToggleStatusRsp
	err := c.doRequest("POST", "/api/v2/sellerchat/set_offer_toggle_status", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetUnreadConversationCount() (*GetUnreadConversationCountRsp, error) {
	var result GetUnreadConversationCountRsp
	err := c.doRequest("GET", "/api/v2/sellerchat/get_unread_conversation_count", nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteConversation(body easycb.AnyMap) (*DeleteConversationRsp, error) {
	var result DeleteConversationRsp
	err := c.doRequest("POST", "/api/v2/sellerchat/delete_conversation", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PinConversation(body easycb.AnyMap) (*PinConversationRsp, error) {
	var result PinConversationRsp
	err := c.doRequest("POST", "/api/v2/sellerchat/pin_conversation", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UnpinConversation(body easycb.AnyMap) (*UnpinConversationRsp, error) {
	var result UnpinConversationRsp
	err := c.doRequest("POST", "/api/v2/sellerchat/unpin_conversation", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SellerChatUploadImage(body easycb.AnyMap, fileType string) (*SellerChatUploadImageRsp, error) {
	var result SellerChatUploadImageRsp
	err := c.doFileRequest("POST", "/api/v2/media_space/upload_image", nil, body, fileType, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
