package shopee

import "github.com/easycb/easycb-go"

func (c *Client) GetTimeSlotId(query easycb.AnyMap) (*GetTimeSlotIdRsp, error) {
	var result GetTimeSlotIdRsp
	err := c.doRequest("GET", "/api/v2/shop_flash_sale/get_time_slot_id", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CreateShopFlashSale(body easycb.AnyMap) (*CreateShopFlashSaleRsp, error) {
	var result CreateShopFlashSaleRsp
	err := c.doRequest("POST", "/api/v2/shop_flash_sale/create_shop_flash_sale", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetItemCriteria(query easycb.AnyMap) (*GetItemCriteriaRsp, error) {
	var result GetItemCriteriaRsp
	err := c.doRequest("GET", "/api/v2/shop_flash_sale/get_item_criteria", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) AddShopFlashSaleItems(body easycb.AnyMap) (*AddShopFlashSaleItemsRsp, error) {
	var result AddShopFlashSaleItemsRsp
	err := c.doRequest("POST", "/api/v2/shop_flash_sale/add_shop_flash_sale_items", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetShopFlashSaleList(query easycb.AnyMap) (*GetShopFlashSaleListRsp, error) {
	var result GetShopFlashSaleListRsp
	err := c.doRequest("GET", "/api/v2/shop_flash_sale/get_shop_flash_sale_list", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetShopFlashSale(query easycb.AnyMap) (*GetShopFlashSaleRsp, error) {
	var result GetShopFlashSaleRsp
	err := c.doRequest("GET", "/api/v2/shop_flash_sale/get_shop_flash_sale", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetShopFlashSaleItems(query easycb.AnyMap) (*GetShopFlashSaleItemsRsp, error) {
	var result GetShopFlashSaleItemsRsp
	err := c.doRequest("GET", "/api/v2/shop_flash_sale/get_shop_flash_sale_items", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateShopFlashSale(body easycb.AnyMap) (*UpdateShopFlashSaleRsp, error) {
	var result UpdateShopFlashSaleRsp
	err := c.doRequest("POST", "/api/v2/shop_flash_sale/update_shop_flash_sale", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateShopFlashSaleItems(body easycb.AnyMap) (*UpdateShopFlashSaleItemsRsp, error) {
	var result UpdateShopFlashSaleItemsRsp
	err := c.doRequest("POST", "/api/v2/shop_flash_sale/update_shop_flash_sale_items", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteShopFlashSale(body easycb.AnyMap) (*DeleteShopFlashSaleRsp, error) {
	var result DeleteShopFlashSaleRsp
	err := c.doRequest("POST", "/api/v2/shop_flash_sale/delete_shop_flash_sale", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteShopFlashSaleItems(body easycb.AnyMap) (*DeleteShopFlashSaleItemsRsp, error) {
	var result DeleteShopFlashSaleItemsRsp
	err := c.doRequest("POST", "/api/v2/shop_flash_sale/delete_shop_flash_sale_items", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
