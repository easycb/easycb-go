package lazada

import "github.com/easycb/easycb-go"

func (c *Client) CreateGlobalProduct(body easycb.AnyMap) (*CreateGlobalProductRsp, error) {
	var result CreateGlobalProductRsp
	err := c.doRequest("POST", "/product/global/create", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetGlobalProductExtension(query easycb.AnyMap) (*GetGlobalProductExtensionRsp, error) {
	var result GetGlobalProductExtensionRsp
	err := c.doRequest("GET", "/product/global/extension", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetGlobalProductStatus(query easycb.AnyMap) (*GetGlobalProductStatusRsp, error) {
	var result GetGlobalProductStatusRsp
	err := c.doRequest("GET", "/product/global/status/get", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetRecommendPrice(query easycb.AnyMap) (*GetRecommendPriceRsp, error) {
	var result GetRecommendPriceRsp
	err := c.doRequest("GET", "/product/global/semi/recommend/price/get", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetUnfilledAttribute(query easycb.AnyMap) (*GetUnfilledAttributeRsp, error) {
	var result GetUnfilledAttributeRsp
	err := c.doRequest("GET", "/product/global/unfilled/attribute/get", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetUpgradableGlobalPlusProductList(query easycb.AnyMap) (*GetUpgradableGlobalPlusProductListRsp, error) {
	var result GetUpgradableGlobalPlusProductListRsp
	err := c.doRequest("GET", "/product/global/semi/avaible/get", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SemiProductUpdate(body easycb.AnyMap) (*SemiProductUpdateRsp, error) {
	var result SemiProductUpdateRsp
	err := c.doRequest("POST", "/product/global/semi/update", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SemiProductUpgrade(body easycb.AnyMap) (*SemiProductUpgradeRsp, error) {
	var result SemiProductUpgradeRsp
	err := c.doRequest("POST", "/product/global/semi/upgrade", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateGlobalProductAttribute(body easycb.AnyMap) (*UpdateGlobalProductAttributeRsp, error) {
	var result UpdateGlobalProductAttributeRsp
	err := c.doRequest("POST", "/product/global/attribute/update", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteMerchantProduct(body easycb.AnyMap) (*DeleteMerchantProductRsp, error) {
	var result DeleteMerchantProductRsp
	err := c.doRequest("POST", "/product/global/delete", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateProductStatus(body easycb.AnyMap) (*UpdateProductStatusRsp, error) {
	var result UpdateProductStatusRsp
	err := c.doRequest("POST", "/product/global/update/status", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
