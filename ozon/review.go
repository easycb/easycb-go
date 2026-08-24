package ozon

import "github.com/easycb/easycb-go"

func (c *Client) ReviewCommentCreate(body easycb.AnyMap) (*ReviewCommentCreateRsp, error) {
	var result ReviewCommentCreateRsp
	err := c.doRequest("POST", "/v1/review/comment/create", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ReviewCommentDelete(body easycb.AnyMap) (*ReviewCommentDeleteRsp, error) {
	var result ReviewCommentDeleteRsp
	err := c.doRequest("POST", "/v1/review/comment/delete", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ReviewCommentList(body easycb.AnyMap) (*ReviewCommentListRsp, error) {
	var result ReviewCommentListRsp
	err := c.doRequest("POST", "/v1/review/comment/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ReviewChangeStatus(body easycb.AnyMap) (*ReviewChangeStatusRsp, error) {
	var result ReviewChangeStatusRsp
	err := c.doRequest("POST", "/v1/review/change-status", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ReviewCount(body easycb.AnyMap) (*ReviewCountRsp, error) {
	var result ReviewCountRsp
	err := c.doRequest("POST", "/v1/review/count", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ReviewInfo(body easycb.AnyMap) (*ReviewInfoRsp, error) {
	var result ReviewInfoRsp
	err := c.doRequest("POST", "/v1/review/info", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ReviewList(body easycb.AnyMap) (*ReviewListRsp, error) {
	var result ReviewListRsp
	err := c.doRequest("POST", "/v2/review/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
