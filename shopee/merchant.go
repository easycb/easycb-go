package shopee

import "github.com/easycb/easycb-go"

func (c *Client) GetMerchantInfo() (*GetMerchantInfoRsp, error) {
	var result GetMerchantInfoRsp
	err := c.doRequest("GET", "/api/v2/merchant/get_merchant_info", nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetShopListByMerchant(query easycb.AnyMap) (*GetShopListByMerchantRsp, error) {
	var result GetShopListByMerchantRsp
	err := c.doRequest("GET", "/api/v2/merchant/get_shop_list_by_merchant", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetMerchantWarehouseLocationList() (*GetMerchantWarehouseLocationListRsp, error) {
	var result GetMerchantWarehouseLocationListRsp
	err := c.doRequest("GET", "/api/v2/merchant/get_merchant_warehouse_location_list", nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetMerchantWarehouseList(body easycb.AnyMap) (*GetMerchantWarehouseListRsp, error) {
	var result GetMerchantWarehouseListRsp
	err := c.doRequest("POST", "/api/v2/merchant/get_merchant_warehouse_list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetWarehouseEligibleShopList(body easycb.AnyMap) (*GetWarehouseEligibleShopListRsp, error) {
	var result GetWarehouseEligibleShopListRsp
	err := c.doRequest("POST", "/api/v2/merchant/get_warehouse_eligible_shop_list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
