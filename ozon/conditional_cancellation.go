package ozon

import "github.com/easycb/easycb-go"

func (c *Client) ConditionalCancellationList(body easycb.AnyMap) (*ConditionalCancellationListRsp, error) {
	var result ConditionalCancellationListRsp
	err := c.doRequest("POST", "/v2/conditional-cancellation/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ConditionalCancellationApprove(body easycb.AnyMap) (*ConditionalCancellationApproveRsp, error) {
	var result ConditionalCancellationApproveRsp
	err := c.doRequest("POST", "/v2/conditional-cancellation/approve", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ConditionalCancellationReject(body easycb.AnyMap) (*ConditionalCancellationRejectRsp, error) {
	var result ConditionalCancellationRejectRsp
	err := c.doRequest("POST", "/v2/conditional-cancellation/reject", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
