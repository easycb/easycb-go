package ozon

import "github.com/easycb/easycb-go"

func (c *Client) SupplyOrderBundleRsp(body easycb.AnyMap) (*SupplyOrderBundleRsp, error) {
	var result SupplyOrderBundleRsp
	err := c.doRequest("POST", "/v1/supply-order/bundle", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
