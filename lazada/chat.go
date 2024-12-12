package lazada

import "github.com/easycb/easycb-go"

func (c *Client) GetConversationList(query easycb.AnyMap) (*GetConversationListRsp, error) {
	var result GetConversationListRsp
	err := c.doRequest("GET", "/im/session/list", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetConversationDetail(query easycb.AnyMap) (*GetOneConversationRsp, error) {
	var result GetOneConversationRsp
	err := c.doRequest("GET", "/im/session/get", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ReadConversation(body easycb.AnyMap) (*ReadConversationRsp, error) {
	var result ReadConversationRsp
	err := c.doRequest("POST", "/im/session/read", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetMessages(query easycb.AnyMap) (*GetMessagesRsp, error) {
	var result GetMessagesRsp
	err := c.doRequest("GET", "/im/message/list", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SendMessage(body easycb.AnyMap) (*SendMessageRsp, error) {
	var result SendMessageRsp
	err := c.doRequest("POST", "/im/message/send", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) RecallMessage(body easycb.AnyMap) (*RecallMessageRsp, error) {
	var result RecallMessageRsp
	err := c.doRequest("POST", "/im/message/recall", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
