package lazada

import "github.com/easycb/easycb-go"

func (c *Client) FreeShippingActivate(body easycb.AnyMap) (*FreeShippingActivateRsp, error) {
	var result FreeShippingActivateRsp
	err := c.doRequest("POST", "/promotion/freeshipping/activate", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FreeShippingAddSelectedProductSKU(body easycb.AnyMap) (*FreeShippingAddSelectedProductSKURsp, error) {
	var result FreeShippingAddSelectedProductSKURsp
	err := c.doRequest("POST", "/promotion/freeshipping/product/sku/add", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FreeShippingCreate(body easycb.AnyMap) (*FreeShippingCreateRsp, error) {
	var result FreeShippingCreateRsp
	err := c.doRequest("POST", "/promotion/freeshipping/create", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FreeShippingDeactivate(body easycb.AnyMap) (*FreeShippingDeactivateRsp, error) {
	var result FreeShippingDeactivateRsp
	err := c.doRequest("POST", "/promotion/freeshipping/deactivate", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FreeShippingDeleteSelectedProductSKU(body easycb.AnyMap) (*FreeShippingDeleteSelectedProductSKURsp, error) {
	var result FreeShippingDeleteSelectedProductSKURsp
	err := c.doRequest("POST", "/promotion/freeshipping/product/sku/remove", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FreeShippingDeliveryOptionsQuery() (*FreeShippingDeliveryOptionsQueryRsp, error) {
	var result FreeShippingDeliveryOptionsQueryRsp
	err := c.doRequest("GET", "/promotion/freeshipping/deliveryoptions/get", nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FreeShippingGet(query easycb.AnyMap) (*FreeShippingGetRsp, error) {
	var result FreeShippingGetRsp
	err := c.doRequest("GET", "/promotion/freeshipping/get", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FreeShippingList(query easycb.AnyMap) (*FreeShippingListRsp, error) {
	var result FreeShippingListRsp
	err := c.doRequest("GET", "/promotion/freeshippings/get", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FreeShippingRegionsQuery() (*FreeShippingRegionsQueryRsp, error) {
	var result FreeShippingRegionsQueryRsp
	err := c.doRequest("GET", "/promotion/freeshipping/regions/get", nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FreeShippingSelectedProductList(query easycb.AnyMap) (*FreeShippingSelectedProductListRsp, error) {
	var result FreeShippingSelectedProductListRsp
	err := c.doRequest("GET", "/promotion/freeshipping/products/get", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FreeShippingUpdate(body easycb.AnyMap) (*FreeShippingUpdateRsp, error) {
	var result FreeShippingUpdateRsp
	err := c.doRequest("POST", "/promotion/freeshipping/update", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
