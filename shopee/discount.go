package shopee

import "github.com/easycb/easycb-go"

func (c *Client) AddDiscount(body easycb.AnyMap) (*AddDiscountRsp, error) {
	var result AddDiscountRsp
	err := c.doRequest("POST", "/api/v2/discount/add_discount", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) AddDiscountItem(body easycb.AnyMap) (*AddDiscountItemRsp, error) {
	var result AddDiscountItemRsp
	err := c.doRequest("POST", "/api/v2/discount/add_discount_item", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteDiscount(body easycb.AnyMap) (*DeleteDiscountRsp, error) {
	var result DeleteDiscountRsp
	err := c.doRequest("POST", "/api/v2/discount/delete_discount", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteDiscountItem(body easycb.AnyMap) (*DeleteDiscountItemRsp, error) {
	var result DeleteDiscountItemRsp
	err := c.doRequest("POST", "/api/v2/discount/delete_discount_item", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetDiscount(query easycb.AnyMap) (*GetDiscountRsp, error) {
	var result GetDiscountRsp
	err := c.doRequest("GET", "/api/v2/discount/get_discount", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetDiscountList(query easycb.AnyMap) (*GetDiscountListRsp, error) {
	var result GetDiscountListRsp
	err := c.doRequest("GET", "/api/v2/discount/get_discount_list", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateDiscount(body easycb.AnyMap) (*UpdateDiscountRsp, error) {
	var result UpdateDiscountRsp
	err := c.doRequest("POST", "/api/v2/discount/update_discount", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateDiscountItem(body easycb.AnyMap) (*UpdateDiscountItemRsp, error) {
	var result UpdateDiscountItemRsp
	err := c.doRequest("POST", "/api/v2/discount/update_discount_item", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) EndDiscount(body easycb.AnyMap) (*EndDiscountRsp, error) {
	var result EndDiscountRsp
	err := c.doRequest("POST", "/api/v2/discount/end_discount", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
