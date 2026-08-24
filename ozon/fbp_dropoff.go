package ozon

import "github.com/easycb/easycb-go"

func (c *Client) FbpOrderDropOffCancel(body easycb.AnyMap) (*FbpOrderDropOffCancelRsp, error) {
	var result FbpOrderDropOffCancelRsp
	err := c.doRequest("POST", "/v1/fbp/order/drop-off/cancel", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbpOrderDropOffDlvEdit(body easycb.AnyMap) (*FbpOrderDropOffDlvEditRsp, error) {
	var result FbpOrderDropOffDlvEditRsp
	err := c.doRequest("POST", "/v1/fbp/order/drop-off/dlv/edit", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbpOrderDropOffTimetable(body easycb.AnyMap) (*FbpOrderDropOffTimetableRsp, error) {
	var result FbpOrderDropOffTimetableRsp
	err := c.doRequest("POST", "/v1/fbp/order/drop-off/timetable", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
