package shopee

import "github.com/easycb/easycb-go"

func (c *Client) AddAddOnDeal(body easycb.AnyMap) (*AddAddOnDealRsp, error) {
	var result AddAddOnDealRsp
	err := c.doRequest("POST", "/api/v2/add_on_deal/add_add_on_deal", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) AddAddOnDealMainItem(body easycb.AnyMap) (*AddAddOnDealMainItemRsp, error) {
	var result AddAddOnDealMainItemRsp
	err := c.doRequest("POST", "/api/v2/add_on_deal/add_add_on_deal_main_item", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) AddAddOnDealSubItem(body easycb.AnyMap) (*AddAddOnDealSubItemRsp, error) {
	var result AddAddOnDealSubItemRsp
	err := c.doRequest("POST", "/api/v2/add_on_deal/add_add_on_deal_sub_item", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteAddOnDeal(body easycb.AnyMap) (*DeleteAddOnDealRsp, error) {
	var result DeleteAddOnDealRsp
	err := c.doRequest("POST", "/api/v2/add_on_deal/delete_add_on_deal", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteAddOnDealMainItem(body easycb.AnyMap) (*DeleteAddOnDealMainItemRsp, error) {
	var result DeleteAddOnDealMainItemRsp
	err := c.doRequest("POST", "/api/v2/add_on_deal/delete_add_on_deal_main_item", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteAddOnDealSubItem(body easycb.AnyMap) (*DeleteAddOnDealSubItemRsp, error) {
	var result DeleteAddOnDealSubItemRsp
	err := c.doRequest("POST", "/api/v2/add_on_deal/delete_add_on_deal_sub_item", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetAddOnDealList(query easycb.AnyMap) (*GetAddOnDealListRsp, error) {
	var result GetAddOnDealListRsp
	err := c.doRequest("GET", "/api/v2/add_on_deal/get_add_on_deal_list", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetAddOnDeal(query easycb.AnyMap) (*GetAddOnDealRsp, error) {
	var result GetAddOnDealRsp
	err := c.doRequest("GET", "/api/v2/add_on_deal/get_add_on_deal", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetAddOnDealMainItem(query easycb.AnyMap) (*GetAddOnDealMainItemRsp, error) {
	var result GetAddOnDealMainItemRsp
	err := c.doRequest("GET", "/api/v2/add_on_deal/get_add_on_deal_main_item", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetAddOnDealSubItem(query easycb.AnyMap) (*GetAddOnDealSubItemRsp, error) {
	var result GetAddOnDealSubItemRsp
	err := c.doRequest("GET", "/api/v2/add_on_deal/get_add_on_deal_sub_item", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateAddOnDeal(body easycb.AnyMap) (*UpdateAddOnDealRsp, error) {
	var result UpdateAddOnDealRsp
	err := c.doRequest("POST", "/api/v2/add_on_deal/update_add_on_deal", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateAddOnDealMainItem(body easycb.AnyMap) (*UpdateAddOnDealMainItemRsp, error) {
	var result UpdateAddOnDealMainItemRsp
	err := c.doRequest("POST", "/api/v2/add_on_deal/update_add_on_deal_main_item", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateAddOnDealSubItem(body easycb.AnyMap) (*UpdateAddOnDealSubItemRsp, error) {
	var result UpdateAddOnDealSubItemRsp
	err := c.doRequest("POST", "/api/v2/add_on_deal/update_add_on_deal_sub_item", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) EndAddOnDeal(body easycb.AnyMap) (*EndAddOnDealRsp, error) {
	var result EndAddOnDealRsp
	err := c.doRequest("POST", "/api/v2/add_on_deal/end_add_on_deal", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
