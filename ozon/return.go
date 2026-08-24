package ozon

import "github.com/easycb/easycb-go"

func (c *Client) ReturnsList(body easycb.AnyMap) (*ReturnsListRsp, error) {
	var result ReturnsListRsp
	err := c.doRequest("POST", "/v1/returns/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ReturnsRFbsList(body easycb.AnyMap) (*ReturnsRFbsListRsp, error) {
	var result ReturnsRFbsListRsp
	err := c.doRequest("POST", "/v2/returns/rfbs/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ReturnsRFbsGet(body easycb.AnyMap) (*ReturnsRFbsGetRsp, error) {
	var result ReturnsRFbsGetRsp
	err := c.doRequest("POST", "/v2/returns/rfbs/get", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ReturnsRFbsActionSet(body easycb.AnyMap) (*ReturnsRFbsActionSetRsp, error) {
	var result ReturnsRFbsActionSetRsp
	err := c.doRequest("POST", "/v1/returns/rfbs/action/set", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
