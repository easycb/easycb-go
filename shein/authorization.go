package shein

import "github.com/easycb/easycb-go"

func (c *Client) AuthGetByToken(body easycb.AnyMap, headers map[string]string) (*AuthGetByTokenRsp, error) {
	var result AuthGetByTokenRsp
	err := c.doRequest("POST", "/open-api/auth/get-by-token", nil, body, headers, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
