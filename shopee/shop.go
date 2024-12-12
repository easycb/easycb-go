package shopee

import "github.com/easycb/easycb-go"

func (c *Client) GetShopInfo() (*GetShopInfoRsp, error) {
	var result GetShopInfoRsp
	err := c.doRequest("GET", "/api/v2/shop/get_shop_info", nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetProfile() (*GetProfileRsp, error) {
	var result GetProfileRsp
	err := c.doRequest("GET", "/api/v2/shop/get_profile", nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateProfile(body easycb.AnyMap) (*UpdateProfileRsp, error) {
	var result UpdateProfileRsp
	err := c.doRequest("POST", "/api/v2/shop/update_profile", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetWarehouseDetail() (*GetWarehouseDetailRsp, error) {
	var result GetWarehouseDetailRsp
	err := c.doRequest("GET", "/api/v2/shop/get_warehouse_detail", nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetShopNotification(query easycb.AnyMap) (*GetShopNotificationRsp, error) {
	var result GetShopNotificationRsp
	err := c.doRequest("GET", "/api/v2/shop/get_shop_notification", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetAuthorisedResellerBrand(query easycb.AnyMap) (*GetAuthorisedResellerBrandRsp, error) {
	var result GetAuthorisedResellerBrandRsp
	err := c.doRequest("GET", "/api/v2/shop/get_authorised_reseller_brand", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
