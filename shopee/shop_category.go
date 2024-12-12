package shopee

import "github.com/easycb/easycb-go"

func (c *Client) AddShopCategory(body easycb.AnyMap) (*AddShopCategoryRsp, error) {
	var result AddShopCategoryRsp
	err := c.doRequest("POST", "/api/v2/shop_category/add_shop_category", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetShopCategoryList(query easycb.AnyMap) (*GetShopCategoryListRsp, error) {
	var result GetShopCategoryListRsp
	err := c.doRequest("GET", "/api/v2/shop_category/get_shop_category_list", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteShopCategory(body easycb.AnyMap) (*DeleteShopCategoryRsp, error) {
	var result DeleteShopCategoryRsp
	err := c.doRequest("POST", "/api/v2/shop_category/delete_shop_category", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateShopCategory(body easycb.AnyMap) (*UpdateShopCategoryRsp, error) {
	var result UpdateShopCategoryRsp
	err := c.doRequest("POST", "/api/v2/shop_category/update_shop_category", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ShopCategoryAddItemList(body easycb.AnyMap) (*ShopCategoryAddItemListRsp, error) {
	var result ShopCategoryAddItemListRsp
	err := c.doRequest("POST", "/api/v2/shop_category/add_item_list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ShopCategoryGetItemList(query easycb.AnyMap) (*ShopCategoryGetItemListRsp, error) {
	var result ShopCategoryGetItemListRsp
	err := c.doRequest("GET", "/api/v2/shop_category/get_item_list", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ShopCategoryDeleteItemList(body easycb.AnyMap) (*ShopCategoryDeleteItemListRsp, error) {
	var result ShopCategoryDeleteItemListRsp
	err := c.doRequest("POST", "/api/v2/shop_category/delete_item_list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
