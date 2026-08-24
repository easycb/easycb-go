package ozon

import "github.com/easycb/easycb-go"

func (c *Client) PricingStrategyCompetitorsList(body easycb.AnyMap) (*PricingStrategyCompetitorsListRsp, error) {
	var result PricingStrategyCompetitorsListRsp
	err := c.doRequest("POST", "/v1/pricing-strategy/competitors/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PricingStrategyList(body easycb.AnyMap) (*PricingStrategyListRsp, error) {
	var result PricingStrategyListRsp
	err := c.doRequest("POST", "/v1/pricing-strategy/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PricingStrategyCreate(body easycb.AnyMap) (*PricingStrategyCreateRsp, error) {
	var result PricingStrategyCreateRsp
	err := c.doRequest("POST", "/v1/pricing-strategy/create", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PricingStrategyInfo(body easycb.AnyMap) (*PricingStrategyInfoRsp, error) {
	var result PricingStrategyInfoRsp
	err := c.doRequest("POST", "/v1/pricing-strategy/info", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PricingStrategyUpdate(body easycb.AnyMap) (*PricingStrategyUpdateRsp, error) {
	var result PricingStrategyUpdateRsp
	err := c.doRequest("POST", "/v1/pricing-strategy/update", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PricingStrategyProductAdd(body easycb.AnyMap) (*PricingStrategyProductAddRsp, error) {
	var result PricingStrategyProductAddRsp
	err := c.doRequest("POST", "/v1/pricing-strategy/products/add", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PricingStrategyStrategyIdsByProductIds(body easycb.AnyMap) (*PricingStrategyStrategyIdsByProductIdsRsp, error) {
	var result PricingStrategyStrategyIdsByProductIdsRsp
	err := c.doRequest("POST", "/v1/pricing-strategy/strategy-ids-by-product-ids", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PricingStrategyProductList(body easycb.AnyMap) (*PricingStrategyProductListRsp, error) {
	var result PricingStrategyProductListRsp
	err := c.doRequest("POST", "/v1/pricing-strategy/products/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PricingStrategyProductInfo(body easycb.AnyMap) (*PricingStrategyProductInfoRsp, error) {
	var result PricingStrategyProductInfoRsp
	err := c.doRequest("POST", "/v1/pricing-strategy/product/info", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PricingStrategyProductDelete(body easycb.AnyMap) (*PricingStrategyProductDeleteRsp, error) {
	var result PricingStrategyProductDeleteRsp
	err := c.doRequest("POST", "/v1/pricing-strategy/products/delete", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PricingStrategyStatus(body easycb.AnyMap) (*PricingStrategyStatusRsp, error) {
	var result PricingStrategyStatusRsp
	err := c.doRequest("POST", "/v1/pricing-strategy/status", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PricingStrategyDelete(body easycb.AnyMap) (*PricingStrategyDeleteRsp, error) {
	var result PricingStrategyDeleteRsp
	err := c.doRequest("POST", "/v1/pricing-strategy/status", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
