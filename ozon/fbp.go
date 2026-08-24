package ozon

import "github.com/easycb/easycb-go"

func (c *Client) PostingFbpGet(body easycb.AnyMap) (*PostingFbpGetRsp, error) {
	var result PostingFbpGetRsp
	err := c.doRequest("POST", "/v1/posting/fbp/get", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PostingFbpList(body easycb.AnyMap) (*PostingFbpListRsp, error) {
	var result PostingFbpListRsp
	err := c.doRequest("POST", "/v1/posting/fbp/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbpActFromCreate(body easycb.AnyMap) (*PostingFbpGetRsp, error) {
	var result PostingFbpGetRsp
	err := c.doRequest("POST", "/v1/fbp/act-from/create", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbpActFromGet(body easycb.AnyMap) (*PostingFbpListRsp, error) {
	var result PostingFbpListRsp
	err := c.doRequest("POST", "/v1/fbp/act-from/get", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbpActToCreate(body easycb.AnyMap) (*FbpActFromCreateRsp, error) {
	var result FbpActFromCreateRsp
	err := c.doRequest("POST", "/v1/fbp/act-to/create", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbpActToGet(body easycb.AnyMap) (*FbpActFromGetRsp, error) {
	var result FbpActFromGetRsp
	err := c.doRequest("POST", "/v1/fbp/act-to/get", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbpArchiveGet(body easycb.AnyMap) (*FbpActToCreateRsp, error) {
	var result FbpActToCreateRsp
	err := c.doRequest("POST", "/v1/fbp/archive/get", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbpArchiveList(body easycb.AnyMap) (*FbpActToGetRsp, error) {
	var result FbpActToGetRsp
	err := c.doRequest("POST", "/v1/fbp/archive/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbpLabelCreate(body easycb.AnyMap) (*FbpArchiveGetRsp, error) {
	var result FbpArchiveGetRsp
	err := c.doRequest("POST", "/v1/fbp/label/create", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbpLabelGet(body easycb.AnyMap) (*FbpArchiveListRsp, error) {
	var result FbpArchiveListRsp
	err := c.doRequest("POST", "/v1/fbp/label/get", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbpOrderGet(body easycb.AnyMap) (*FbpLabelCreateRsp, error) {
	var result FbpLabelCreateRsp
	err := c.doRequest("POST", "/v1/fbp/order/get", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbpOrderList(body easycb.AnyMap) (*FbpLabelGetRsp, error) {
	var result FbpLabelGetRsp
	err := c.doRequest("POST", "/v1/fbp/order/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
