package ozon

import "github.com/easycb/easycb-go"

func (c *Client) FbsPostingProductExemplarSet(body easycb.AnyMap) (*FbsPostingProductExemplarSetRsp, error) {
	var result FbsPostingProductExemplarSetRsp
	err := c.doRequest("POST", "/v6/fbs/posting/product/exemplar/set", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbsPostingProductExemplarCreateOrGet(body easycb.AnyMap) (*FbsPostingProductExemplarCreateOrGetRsp, error) {
	var result FbsPostingProductExemplarCreateOrGetRsp
	err := c.doRequest("POST", "/v6/fbs/posting/product/exemplar/create-or-get", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbsPostingProductExemplarStatus(body easycb.AnyMap) (*FbsPostingProductExemplarStatusRsp, error) {
	var result FbsPostingProductExemplarStatusRsp
	err := c.doRequest("POST", "/v5/fbs/posting/product/exemplar/status", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbsPostingProductExemplarValidate(body easycb.AnyMap) (*FbsPostingProductExemplarValidateRsp, error) {
	var result FbsPostingProductExemplarValidateRsp
	err := c.doRequest("POST", "/v5/fbs/posting/product/exemplar/validate", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbsPostingProductExemplarUpdate(body easycb.AnyMap) (*FbsPostingProductExemplarUpdateRsp, error) {
	var result FbsPostingProductExemplarUpdateRsp
	err := c.doRequest("POST", "/v1/fbs/posting/product/exemplar/update", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
