package ozon

import "github.com/easycb/easycb-go"

func (c *Client) PolygonCreate(body easycb.AnyMap) (*PolygonCreateRsp, error) {
	var result PolygonCreateRsp
	err := c.doRequest("POST", "/v1/polygon/create", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PolygonBind(body easycb.AnyMap) (*PolygonBindRsp, error) {
	var result PolygonBindRsp
	err := c.doRequest("POST", "/v1/polygon/bind", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
