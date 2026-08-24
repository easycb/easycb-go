package ozon

import "github.com/easycb/easycb-go"

func (c *Client) FbpDraftPickUpCreate(body easycb.AnyMap) (*FbpDraftPickUpCreateRsp, error) {
	var result FbpDraftPickUpCreateRsp
	err := c.doRequest("POST", "/v1/fbp/draft/pick-up/create", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbpDraftPickUpDelete(body easycb.AnyMap) (*FbpDraftPickUpDeleteRsp, error) {
	var result FbpDraftPickUpDeleteRsp
	err := c.doRequest("POST", "/v1/fbp/draft/pick-up/delete", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbpDraftPickUpDlvEdit(body easycb.AnyMap) (*FbpDraftPickUpDlvEditRsp, error) {
	var result FbpDraftPickUpDlvEditRsp
	err := c.doRequest("POST", "/v1/fbp/draft/pick-up/dlv/edit", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbpDraftPickUpProductValidate(body easycb.AnyMap) (*FbpDraftPickUpProductValidateRsp, error) {
	var result FbpDraftPickUpProductValidateRsp
	err := c.doRequest("POST", "/v1/fbp/draft/pick-up/product/validate", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbpDraftPickUpRegistrate(body easycb.AnyMap) (*FbpDraftPickUpRegistrateRsp, error) {
	var result FbpDraftPickUpRegistrateRsp
	err := c.doRequest("POST", "/v1/fbp/draft/pick-up/registrate", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
