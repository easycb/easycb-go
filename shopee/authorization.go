package shopee

import "github.com/easycb/easycb-go"

func (c *Client) RefreshAccessToken(body easycb.AnyMap) (*RefreshAccessTokenRsp, error) {
	var result RefreshAccessTokenRsp
	err := c.doRequest("POST", "/api/v2/auth/access_token/get", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetAccessToken(body easycb.AnyMap) (*GetAccessTokenRsp, error) {
	var result GetAccessTokenRsp
	err := c.doRequest("POST", "/api/v2/auth/token/get", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
