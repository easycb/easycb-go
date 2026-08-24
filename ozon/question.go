package ozon

import "github.com/easycb/easycb-go"

func (c *Client) QuestionAnswerCreate(body easycb.AnyMap) (*QuestionAnswerCreateRsp, error) {
	var result QuestionAnswerCreateRsp
	err := c.doRequest("POST", "/v1/question/answer/create", nil, body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) QuestionAnswerDelete(body easycb.AnyMap) (*QuestionAnswerDeleteRsp, error) {
	var result QuestionAnswerDeleteRsp
	err := c.doRequest("POST", "/v1/question/answer/delete", nil, body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) QuestionAnswerList(body easycb.AnyMap) (*QuestionAnswerListRsp, error) {
	var result QuestionAnswerListRsp
	err := c.doRequest("POST", "/v1/question/answer/list", nil, body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) QuestionChangeStatus(body easycb.AnyMap) (*QuestionChangeStatusRsp, error) {
	var result QuestionChangeStatusRsp
	err := c.doRequest("POST", "/v1/question/change_status", nil, body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) QuestionChangeCount(body easycb.AnyMap) (*QuestionChangeCountRsp, error) {
	var result QuestionChangeCountRsp
	err := c.doRequest("POST", "/v1/question/count", nil, body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) QuestionInfo(body easycb.AnyMap) (*QuestionInfoRsp, error) {
	var result QuestionInfoRsp
	err := c.doRequest("POST", "/v1/question/info", nil, body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) QuestionList(body easycb.AnyMap) (*QuestionListRsp, error) {
	var result QuestionListRsp
	err := c.doRequest("POST", "/v1/question/list", nil, body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) QuestionTopSku(body easycb.AnyMap) (*QuestionTopSkuRsp, error) {
	var result QuestionTopSkuRsp
	err := c.doRequest("POST", "/v1/question/top-sku", nil, body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
