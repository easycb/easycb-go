package ozon

import "github.com/easycb/easycb-go"

func (c *Client) ChatSendMessage(body easycb.AnyMap) (*ChatSendMessageRsp, error) {
	var result ChatSendMessageRsp
	err := c.doRequest("POST", "/v1/chat/send/message", body, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ChatSendFile(body easycb.AnyMap) (*ChatSendFileRsp, error) {
	var result ChatSendFileRsp
	err := c.doRequest("POST", "/v1/chat/send/file", body, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ChatStart(body easycb.AnyMap) (*ChatStartRsp, error) {
	var result ChatStartRsp
	err := c.doRequest("POST", "/v1/chat/start", body, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ChatList(body easycb.AnyMap) (*ChatListRsp, error) {
	var result ChatListRsp
	err := c.doRequest("POST", "/v3/chat/list", body, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ChatHistory(body easycb.AnyMap) (*ChatHistoryRsp, error) {
	var result ChatHistoryRsp
	err := c.doRequest("POST", "/v3/chat/history", body, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ChatRead(body easycb.AnyMap) (*ChatReadRsp, error) {
	var result ChatReadRsp
	err := c.doRequest("POST", "/v2/chat/read", body, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
