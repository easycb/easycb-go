package tiktok

import "github.com/easycb/easycb-go"

func (c *Client) GetShopWebhooks(query easycb.AnyMap) (*GetShopWebhooksRsp, error) {
	var result GetShopWebhooksRsp
	path := "/event/202309/webhooks"
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateShopWebhooks(query easycb.AnyMap, body easycb.AnyMap) (*UpdateShopWebhooksRsp, error) {
	var result UpdateShopWebhooksRsp
	path := "/event/202309/webhooks"
	err := c.doRequest("PUT", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteShopWebhooks(query easycb.AnyMap, body easycb.AnyMap) (*DeleteShopWebhooksRsp, error) {
	var result DeleteShopWebhooksRsp
	path := "/event/202309/webhooks"
	err := c.doRequest("DELETE", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
