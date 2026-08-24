package ozon

import "github.com/easycb/easycb-go"

func (c *Client) FbpDraftDirectSellerDlvCreate(body easycb.AnyMap) (*FbpDraftDirectSellerDlvCreateRsp, error) {
	var result FbpDraftDirectSellerDlvCreateRsp
	err := c.doRequest("POST", "/v1/fbp/draft/direct/seller-dlv/create", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbpDraftDirectSellerDlvEdit(body easycb.AnyMap) (*FbpDraftDirectSellerDlvEditRsp, error) {
	var result FbpDraftDirectSellerDlvEditRsp
	err := c.doRequest("POST", "/v1/fbp/draft/direct/seller-dlv/edit", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbpDraftDirectTimeslotEdit(body easycb.AnyMap) (*FbpDraftDirectTimeslotEditRsp, error) {
	var result FbpDraftDirectTimeslotEditRsp
	err := c.doRequest("POST", "/v1/fbp/draft/direct/timeslot/edit", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbpDraftDirectTimeslotGet(body easycb.AnyMap) (*FbpDraftDirectTimeslotGetRsp, error) {
	var result FbpDraftDirectTimeslotGetRsp
	err := c.doRequest("POST", "/v1/fbp/draft/direct/timeslot/get", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbpDraftDirectCreate(body easycb.AnyMap) (*FbpDraftDirectCreateRsp, error) {
	var result FbpDraftDirectCreateRsp
	err := c.doRequest("POST", "/v1/fbp/draft/direct/create", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbpDraftDirectDelete(body easycb.AnyMap) (*FbpDraftDirectDeleteRsp, error) {
	var result FbpDraftDirectDeleteRsp
	err := c.doRequest("POST", "/v1/fbp/draft/direct/delete", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbpDraftDirectProductValidate(body easycb.AnyMap) (*FbpDraftDirectProductValidateRsp, error) {
	var result FbpDraftDirectProductValidateRsp
	err := c.doRequest("POST", "/v1/fbp/draft/direct/product/validate", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbpDraftDirectRegistrate(body easycb.AnyMap) (*FbpDraftDirectRegistrateRsp, error) {
	var result FbpDraftDirectRegistrateRsp
	err := c.doRequest("POST", "/v1/fbp/draft/direct/registrate", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbpDraftDirectTplDlvCreate(body easycb.AnyMap) (*FbpDraftDirectTplDlvCreateRsp, error) {
	var result FbpDraftDirectTplDlvCreateRsp
	err := c.doRequest("POST", "/v1/fbp/draft/direct/tpl-dlv/create", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbpDraftDirectTplDlvEdit(body easycb.AnyMap) (*FbpDraftDirectTplDlvEditRsp, error) {
	var result FbpDraftDirectTplDlvEditRsp
	err := c.doRequest("POST", "/v1/fbp/draft/direct/tpl-dlv/edit", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
