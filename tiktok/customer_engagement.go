package tiktok

import "github.com/easycb/easycb-go"

func (c *Client) GetMessageTemplates(query easycb.AnyMap) (*GetMessageTemplatesRsp, error) {
	var result GetMessageTemplatesRsp
	path := "/customer_engagement/202412/message_templates"
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CreateEngagementTask(query easycb.AnyMap, body easycb.AnyMap) (*CreateEngagementTaskRsp, error) {
	var result CreateEngagementTaskRsp
	path := "/customer_engagement/202412/engagement_tasks"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SendEngagementMessage(query easycb.AnyMap, body easycb.AnyMap) (*SendEngagementMessageRsp, error) {
	var result SendEngagementMessageRsp
	path := "/customer_engagement/202412/messages"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetTaskPerformances(query easycb.AnyMap, body easycb.AnyMap) (*GetTaskPerformancesRsp, error) {
	var result GetTaskPerformancesRsp
	path := "/customer_engagement/202412/performances"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
