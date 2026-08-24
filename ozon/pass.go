package ozon

import "github.com/easycb/easycb-go"

func (c *Client) PassList(body easycb.AnyMap) (*PassListRsp, error) {
	var result PassListRsp
	err := c.doRequest("POST", "/v1/pass/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CarriagePassCreate(body easycb.AnyMap) (*CarriagePassCreateRsp, error) {
	var result CarriagePassCreateRsp
	err := c.doRequest("POST", "/v1/carriage/pass/create", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CarriagePassUpdate(body easycb.AnyMap) (*CarriagePassUpdateRsp, error) {
	var result CarriagePassUpdateRsp
	err := c.doRequest("POST", "/v1/carriage/pass/update", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CarriagePassDelete(body easycb.AnyMap) (*CarriagePassDeleteRsp, error) {
	var result CarriagePassDeleteRsp
	err := c.doRequest("POST", "/v1/carriage/pass/delete", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ReturnPassCreate(body easycb.AnyMap) (*ReturnPassCreateRsp, error) {
	var result ReturnPassCreateRsp
	err := c.doRequest("POST", "/v1/return/pass/create", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ReturnPassUpdate(body easycb.AnyMap) (*ReturnPassUpdateRsp, error) {
	var result ReturnPassUpdateRsp
	err := c.doRequest("POST", "/v1/return/pass/update", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ReturnPassDelete(body easycb.AnyMap) (*ReturnPassDeleteRsp, error) {
	var result ReturnPassDeleteRsp
	err := c.doRequest("POST", "/v1/return/pass/delete", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ReturnsCompanyFbsInfo(body easycb.AnyMap) (*ReturnsCompanyFbsInfoRsp, error) {
	var result ReturnsCompanyFbsInfoRsp
	err := c.doRequest("POST", "/v1/returns/company/fbs/info", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
