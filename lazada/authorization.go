package lazada

import "github.com/easycb/easycb-go"

func (c *Client) GetAccessToken(body easycb.AnyMap) (*GetAccessTokenRsp, error) {
	var result GetAccessTokenRsp
	err := c.doRequest("POST", "/auth/token/create", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) RefreshAccessToken(body easycb.AnyMap) (*RefreshAccessTokenRsp, error) {
	var result RefreshAccessTokenRsp
	err := c.doRequest("POST", "/auth/token/refresh", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GenerateAccessTokenWithOpenId(body easycb.AnyMap) (*GenerateAccessTokenWithOpenIdRsp, error) {
	var result GenerateAccessTokenWithOpenIdRsp
	err := c.doRequest("POST", "/auth/token/createWithOpenId", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
