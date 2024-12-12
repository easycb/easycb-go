package lazada

import "github.com/easycb/easycb-go"

func (c *Client) GetHistoryReviewIdList(query easycb.AnyMap) (*GetHistoryReviewIdListRsp, error) {
	var result GetHistoryReviewIdListRsp
	err := c.doRequest("GET", "/review/seller/history/list", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetReviewListByIdList(query easycb.AnyMap) (*GetReviewListByIdListRsp, error) {
	var result GetReviewListByIdListRsp
	err := c.doRequest("GET", "/review/seller/list/v2", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SubmitSellerReply(query easycb.AnyMap) (*SubmitSellerReplyRsp, error) {
	var result SubmitSellerReplyRsp
	err := c.doRequest("GET", "/review/seller/reply/add", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
