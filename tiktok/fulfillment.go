package tiktok

import "fmt"

func (c *Client) GetTracking(orderId string) (*GetTrackingRsp, error) {
	var result GetTrackingRsp
	path := fmt.Sprintf("/fulfillment/202309/orders/%s/tracking", orderId)
	err := c.doRequest("GET", path, nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
