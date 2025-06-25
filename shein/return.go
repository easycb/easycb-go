package shein

import "github.com/easycb/easycb-go"

func (c *Client) ReturnOrderList(body easycb.AnyMap) (*ReturnOrderListRsp, error) {
	var result ReturnOrderListRsp
	err := c.doRequest("POST", "/open-api/return-order/list", nil, body, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ReturnOrderDetail(body easycb.AnyMap) (*ReturnOrderDetailRsp, error) {
	var result ReturnOrderDetailRsp
	err := c.doRequest("POST", "/open-api/return-order/details", nil, body, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SignReturnOrder(body easycb.AnyMap) (*SignReturnOrderRsp, error) {
	var result SignReturnOrderRsp
	err := c.doRequest("POST", "/open-api/return-order/sign-return-order", nil, body, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
