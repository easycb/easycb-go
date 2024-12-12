package shopee

import "github.com/easycb/easycb-go"

func (c *Client) AddFollowPrize(body easycb.AnyMap) (*AddFollowPrizeRsp, error) {
	var result AddFollowPrizeRsp
	err := c.doRequest("POST", "/api/v2/follow_prize/add_follow_prize", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteFollowPrize(body easycb.AnyMap) (*DeleteFollowPrizeRsp, error) {
	var result DeleteFollowPrizeRsp
	err := c.doRequest("POST", "/api/v2/follow_prize/delete_follow_prize", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) EndFollowPrize(body easycb.AnyMap) (*EndFollowPrizeRsp, error) {
	var result EndFollowPrizeRsp
	err := c.doRequest("POST", "/api/v2/follow_prize/end_follow_prize", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateFollowPrize(body easycb.AnyMap) (*UpdateFollowPrizeRsp, error) {
	var result UpdateFollowPrizeRsp
	err := c.doRequest("POST", "/api/v2/follow_prize/update_follow_prize", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetFollowPrizeDetail(query easycb.AnyMap) (*GetFollowPrizeDetailRsp, error) {
	var result GetFollowPrizeDetailRsp
	err := c.doRequest("GET", "/api/v2/follow_prize/get_follow_prize_detail", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetFollowPrizeList(query easycb.AnyMap) (*GetFollowPrizeListRsp, error) {
	var result GetFollowPrizeListRsp
	err := c.doRequest("GET", "/api/v2/follow_prize/get_follow_prize_list", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
