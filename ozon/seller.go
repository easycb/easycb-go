package ozon

import "github.com/easycb/easycb-go"

func (c *Client) GetSellerAccount(body easycb.AnyMap) (*GetSellerAccountRsp, error) {
	var result GetSellerAccountRsp
	err := c.doRequest("POST", "/v1/seller/info", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetSellerOzonLogisticsInfo(body easycb.AnyMap) (*GetSellerOzonLogisticsInfoRsp, error) {
	var result GetSellerOzonLogisticsInfoRsp
	err := c.doRequest("POST", "/v1/seller/ozon-logistics/info", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
