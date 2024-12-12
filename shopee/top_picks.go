package shopee

import "github.com/easycb/easycb-go"

func (c *Client) GetTopPicksList(query easycb.AnyMap) (*GetTopPicksListRsp, error) {
	var result GetTopPicksListRsp
	err := c.doRequest("GET", "/api/v2/top_picks/get_top_picks_list", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) AddTopPicks(body easycb.AnyMap) (*AddTopPicksRsp, error) {
	var result AddTopPicksRsp
	err := c.doRequest("POST", "/api/v2/top_picks/add_top_picks", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateTopPicks(body easycb.AnyMap) (*UpdateTopPicksRsp, error) {
	var result UpdateTopPicksRsp
	err := c.doRequest("POST", "/api/v2/top_picks/update_top_picks", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteTopPicks(body easycb.AnyMap) (*DeleteTopPicksRsp, error) {
	var result DeleteTopPicksRsp
	err := c.doRequest("POST", "/api/v2/top_picks/delete_top_picks", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
