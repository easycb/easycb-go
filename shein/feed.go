package shein

import "github.com/easycb/easycb-go"

func (c *Client) CreateFeedDocument(body easycb.AnyMap) (*CreateFeedDocumentRsp, error) {
	var result CreateFeedDocumentRsp
	err := c.doRequest("POST", "/open-api/sem/feed/createFeedDocument", nil, body, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetFeedDocument(query easycb.AnyMap) (*GetFeedDocumentRsp, error) {
	var result GetFeedDocumentRsp
	err := c.doRequest("GET", "/open-api/sem/feed/getFeedDocument", query, nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UploadDocumentContent(body easycb.AnyMap) (*UploadDocumentContentRsp, error) {
	var result UploadDocumentContentRsp
	err := c.doRequest("POST", "/open-api/sem/feed/uploadDocumentContent", nil, body, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CreateFeed(body easycb.AnyMap) (*CreateFeedRsp, error) {
	var result CreateFeedRsp
	err := c.doRequest("POST", "/open-api/sem/feed/createFeed", nil, body, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetFeed(query easycb.AnyMap) (*GetFeedRsp, error) {
	var result GetFeedRsp
	err := c.doRequest("GET", "/open-api/sem/feed/getFeed", query, nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CancelFeed(body easycb.AnyMap) (*CancelFeedRsp, error) {
	var result CancelFeedRsp
	err := c.doRequest("POST", "/open-api/sem/feed/cancelFeed", nil, body, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
