package shopee

import "github.com/easycb/easycb-go"

func (c *Client) FirstMileGetUnbindOrderList(query easycb.AnyMap) (*FirstMileGetUnbindOrderListRsp, error) {
	var result FirstMileGetUnbindOrderListRsp
	err := c.doRequest("GET", "/api/v2/first_mile/get_unbind_order_list", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FirstMileGetDetail(query easycb.AnyMap) (*FirstMileGetDetailRsp, error) {
	var result FirstMileGetDetailRsp
	err := c.doRequest("GET", "/api/v2/first_mile/get_detail", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FirstMileGenerateFirstMileTrackingNumber(query easycb.AnyMap) (*FirstMileGenerateFirstMileTrackingNumberRsp, error) {
	var result FirstMileGenerateFirstMileTrackingNumberRsp
	err := c.doRequest("GET", "/api/v2/first_mile/generate_first_mile_tracking_number", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FirstMileBindFirstMileTrackingNumber(body easycb.AnyMap) (*FirstMileBindFirstMileTrackingNumberRsp, error) {
	var result FirstMileBindFirstMileTrackingNumberRsp
	err := c.doRequest("POST", "/api/v2/first_mile/bind_first_mile_tracking_number", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FirstMileUnbindFirstMileTrackingNumber(query easycb.AnyMap) (*FirstMileUnbindFirstMileTrackingNumberRsp, error) {
	var result FirstMileUnbindFirstMileTrackingNumberRsp
	err := c.doRequest("POST", "/api/v2/first_mile/unbind_first_mile_tracking_number", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FirstMileGetTrackingNumberList(query easycb.AnyMap) (*FirstMileGetTrackingNumberListRsp, error) {
	var result FirstMileGetTrackingNumberListRsp
	err := c.doRequest("GET", "/api/v2/first_mile/get_tracking_number_list", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FirstMileGetWallBill(body easycb.AnyMap) (*FirstMileGetWallBillRsp, error) {
	var result FirstMileGetWallBillRsp
	err := c.doRequest("POST", "/api/v2/first_mile/get_waybill", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FirstMileGetChannelList(query easycb.AnyMap) (*GetChannelListRsp, error) {
	var result GetChannelListRsp
	err := c.doRequest("GET", "/api/v2/first_mile/get_channel_list", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
