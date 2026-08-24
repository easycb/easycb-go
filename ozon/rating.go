package ozon

import "github.com/easycb/easycb-go"

func (c *Client) RatingIndexFbsInfo(body easycb.AnyMap) (*RatingIndexFbsInfoRsp, error) {
	var result RatingIndexFbsInfoRsp
	err := c.doRequest("POST", "/v1/rating/index/fbs/info", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) RatingIndexFbsPostingList(body easycb.AnyMap) (*RatingIndexFbsPostingListRsp, error) {
	var result RatingIndexFbsPostingListRsp
	err := c.doRequest("POST", "/v1/rating/index/fbs/posting/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
