package shopee

import "github.com/easycb/easycb-go"

func (c *Client) AddBundleDeal(body easycb.AnyMap) (*AddBundleDealRsp, error) {
	var result AddBundleDealRsp
	err := c.doRequest("POST", "/api/v2/bundle_deal/add_bundle_deal", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) AddBundleDealItem(body easycb.AnyMap) (*AddBundleDealItemRsp, error) {
	var result AddBundleDealItemRsp
	err := c.doRequest("POST", "/api/v2/bundle_deal/add_bundle_deal_item", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetBundleDealList(query easycb.AnyMap) (*GetBundleDealListRsp, error) {
	var result GetBundleDealListRsp
	err := c.doRequest("GET", "/api/v2/bundle_deal/get_bundle_deal_list", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetBundleDeal(query easycb.AnyMap) (*GetBundleDealRsp, error) {
	var result GetBundleDealRsp
	err := c.doRequest("GET", "/api/v2/bundle_deal/get_bundle_deal", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetBundleDealItem(query easycb.AnyMap) (*GetBundleDealItemRsp, error) {
	var result GetBundleDealItemRsp
	err := c.doRequest("GET", "/api/v2/bundle_deal/get_bundle_deal_item", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateBundleDeal(body easycb.AnyMap) (*UpdateBundleDealRsp, error) {
	var result UpdateBundleDealRsp
	err := c.doRequest("POST", "/api/v2/bundle_deal/update_bundle_deal", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateBundleDealItem(body easycb.AnyMap) (*UpdateBundleDealItemRsp, error) {
	var result UpdateBundleDealItemRsp
	err := c.doRequest("POST", "/api/v2/bundle_deal/update_bundle_deal_item", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) EndBundleDeal(body easycb.AnyMap) (*EndBundleDealRsp, error) {
	var result EndBundleDealRsp
	err := c.doRequest("POST", "/api/v2/bundle_deal/end_bundle_deal", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteBundleDeal(body easycb.AnyMap) (*DeleteBundleDealRsp, error) {
	var result DeleteBundleDealRsp
	err := c.doRequest("POST", "/api/v2/bundle_deal/delete_bundle_deal", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteBundleDealItem(body easycb.AnyMap) (*DeleteBundleDealItemRsp, error) {
	var result DeleteBundleDealItemRsp
	err := c.doRequest("POST", "/api/v2/bundle_deal/delete_bundle_deal_item", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
