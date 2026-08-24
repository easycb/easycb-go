package ozon

import "github.com/easycb/easycb-go"

func (c *Client) FbpWarehouseList(body easycb.AnyMap) (*FbpWarehouseListRsp, error) {
	var result FbpWarehouseListRsp
	err := c.doRequest("POST", "/v1/fbp/warehouse/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbpDraftGet(body easycb.AnyMap) (*FbpDraftGetRsp, error) {
	var result FbpDraftGetRsp
	err := c.doRequest("POST", "/v1/fbp/draft/get", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbpDraftList(body easycb.AnyMap) (*FbpDraftListRsp, error) {
	var result FbpDraftListRsp
	err := c.doRequest("POST", "/v1/fbp/draft/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
