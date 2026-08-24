package ozon

import "github.com/easycb/easycb-go"

func (c *Client) FbsPostingDelivering(body easycb.AnyMap) (*FbsPostingDeliveringRsp, error) {
	var result FbsPostingDeliveringRsp
	err := c.doRequest("POST", "/v2/fbs/posting/delivering", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbsPostingTrackingNumberSet(body easycb.AnyMap) (*FbsPostingTrackingNumberSetRsp, error) {
	var result FbsPostingTrackingNumberSetRsp
	err := c.doRequest("POST", "/v2/fbs/posting/tracking-number/set", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbsPostingLastMile(body easycb.AnyMap) (*FbsPostingLastMileRsp, error) {
	var result FbsPostingLastMileRsp
	err := c.doRequest("POST", "/v2/fbs/posting/last-mile", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbsPostingDelivered(body easycb.AnyMap) (*FbsPostingDeliveredRsp, error) {
	var result FbsPostingDeliveredRsp
	err := c.doRequest("POST", "/v2/fbs/posting/delivered", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PostingCutoffSet(body easycb.AnyMap) (*PostingCutoffSetRsp, error) {
	var result PostingCutoffSetRsp
	err := c.doRequest("POST", "/v1/posting/cutoff/set", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
