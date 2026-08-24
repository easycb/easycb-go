package ozon

import "github.com/easycb/easycb-go"

func (c *Client) FbpOrderPickUpCancel(body easycb.AnyMap) (*FbpOrderPickUpCancelRsp, error) {
	var result FbpOrderPickUpCancelRsp

	err := c.doRequest("POST", "/v1/fbp/order/pick-up/cancel", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbpOrderPickUpDlvEdit(body easycb.AnyMap) (*FbpOrderPickUpDlvEditRsp, error) {
	var result FbpOrderPickUpDlvEditRsp

	err := c.doRequest("POST", "/v1/fbp/order/pick-up/dlv/edit", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
