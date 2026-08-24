package ozon

import "github.com/easycb/easycb-go"

func (c *Client) GetRoles(body easycb.AnyMap) (*GetRolesRsp, error) {
	var result GetRolesRsp
	err := c.doRequest("POST", "/v1/roles", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
