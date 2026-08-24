package ozon

import "github.com/easycb/easycb-go"

func (c *Client) FbpOrderDirectCancel(body easycb.AnyMap) (*FbpOrderDirectCancelRsp, error) {
	var result FbpOrderDirectCancelRsp
	err := c.doRequest("POST", "/v1/fbp/order/direct/cancel", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbpOrderDirectSellerDlvEdit(body easycb.AnyMap) (*FbpOrderDirectSellerDlvEditRsp, error) {
	var result FbpOrderDirectSellerDlvEditRsp
	err := c.doRequest("POST", "/v1/fbp/order/direct/seller-dlv/edit", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbpOrderDirectTimeslotEdit(body easycb.AnyMap) (*FbpOrderDirectTimeslotEditRsp, error) {
	var result FbpOrderDirectTimeslotEditRsp
	err := c.doRequest("POST", "/v1/fbp/order/direct/timeslot/edit", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbpOrderDirectTimeslotList(body easycb.AnyMap) (*FbpOrderDirectTimeslotListRsp, error) {
	var result FbpOrderDirectTimeslotListRsp
	err := c.doRequest("POST", "/v1/fbp/order/direct/timeslot/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
