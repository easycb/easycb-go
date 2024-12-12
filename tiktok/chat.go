package tiktok

import (
	"fmt"
	"github.com/easycb/easycb-go"
)

func (c *Client) CreateConversation(query easycb.AnyMap, body easycb.AnyMap) (*CreateConversationRsp, error) {
	var result CreateConversationRsp
	path := "/customer_service/202309/conversations"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetConversations(query easycb.AnyMap) (*GetConversationsRsp, error) {
	var result GetConversationsRsp
	path := "/customer_service/202309/conversations"
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetConversationMessages(conversationId string, query easycb.AnyMap) (*GetConversationMessagesRsp, error) {
	var result GetConversationMessagesRsp
	path := fmt.Sprintf("/customer_service/202309/conversations/%s/messages", conversationId)
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UploadBuyerMessagesImage(body easycb.AnyMap, fileType string) (*UploadBuyerMessagesImageRsp, error) {
	var result UploadBuyerMessagesImageRsp
	path := "/customer_service/202309/images/upload"
	err := c.doFileRequest("POST", path, nil, body, fileType, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SendMessage(conversationId string, body easycb.AnyMap) (*SendMessageRsp, error) {
	var result SendMessageRsp
	path := fmt.Sprintf("/customer_service/202309/conversations/%s/messages", conversationId)
	err := c.doRequest("POST", path, nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ReadMessage(conversationId string) (*ReadMessageRsp, error) {
	var result ReadMessageRsp
	path := fmt.Sprintf("/customer_service/202309/conversations/%s/messages/read", conversationId)
	err := c.doRequest("POST", path, nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetAgentSettings(query easycb.AnyMap) (*GetAgentSettingsRsp, error) {
	var result GetAgentSettingsRsp
	path := "/customer_service/202309/agents/settings"
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateAgentSettings(query easycb.AnyMap, body easycb.AnyMap) (*UpdateAgentSettingsRsp, error) {
	var result UpdateAgentSettingsRsp
	path := "/customer_service/202309/agents/settings"
	err := c.doRequest("PUT", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetCustomerServicePerformance(query easycb.AnyMap) (*GetCustomerServicePerformanceRsp, error) {
	var result GetCustomerServicePerformanceRsp
	path := "/customer_service/202407/performance"
	err := c.doRequest("PUT", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
