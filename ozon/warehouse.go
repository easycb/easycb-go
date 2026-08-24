package ozon

import "github.com/easycb/easycb-go"

func (c *Client) WarehouseList(body easycb.AnyMap) (*WarehouseListRsp, error) {
	var result WarehouseListRsp
	err := c.doRequest("POST", "/v2/warehouse/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeliveryMethodList(body easycb.AnyMap) (*DeliveryMethodListRsp, error) {
	var result DeliveryMethodListRsp
	err := c.doRequest("POST", "/v2/delivery-method/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) WarehouseOperationStatus(body easycb.AnyMap) (*WarehouseOperationStatusRsp, error) {
	var result WarehouseOperationStatusRsp
	err := c.doRequest("POST", "/v1/warehouse/operation/status", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) WarehouseArchive(body easycb.AnyMap) (*WarehouseArchiveRsp, error) {
	var result WarehouseArchiveRsp
	err := c.doRequest("POST", "/v1/warehouse/archive", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) WarehouseUnArchive(body easycb.AnyMap) (*WarehouseUnArchiveRsp, error) {
	var result WarehouseUnArchiveRsp
	err := c.doRequest("POST", "/v1/warehouse/unarchive", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) WarehouseInvalidProductsGet(body easycb.AnyMap) (*WarehouseInvalidProductsGetRsp, error) {
	var result WarehouseInvalidProductsGetRsp
	err := c.doRequest("POST", "/v1/warehouse/invalid-products/get", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) WarehouseWarehouseWithInvalidProducts(body easycb.AnyMap) (*WarehouseWarehouseWithInvalidProductsRsp, error) {
	var result WarehouseWarehouseWithInvalidProductsRsp
	err := c.doRequest("POST", "/v1/warehouse/warehouses-with-invalid-products", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) WarehouseFbsCreateDropOffList(body easycb.AnyMap) (*WarehouseFbsCreateDropOffListRsp, error) {
	var result WarehouseFbsCreateDropOffListRsp
	err := c.doRequest("POST", "/v1/warehouse/fbs/create/drop-off/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) WarehouseFbsUpdateDropOffList(body easycb.AnyMap) (*WarehouseFbsUpdateDropOffListRsp, error) {
	var result WarehouseFbsUpdateDropOffListRsp
	err := c.doRequest("POST", "/v1/warehouse/fbs/update/drop-off/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) WarehouseFbsCreateDropOffTimeslotList(body easycb.AnyMap) (*WarehouseFbsCreateDropOffTimeslotListRsp, error) {
	var result WarehouseFbsCreateDropOffTimeslotListRsp
	err := c.doRequest("POST", "/v1/warehouse/fbs/create/drop-off/timeslot/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) WarehouseFbsUpdateDropOffTimeslotList(body easycb.AnyMap) (*WarehouseFbsUpdateDropOffTimeslotListRsp, error) {
	var result WarehouseFbsUpdateDropOffTimeslotListRsp
	err := c.doRequest("POST", "/v1/warehouse/fbs/update/drop-off/timeslot/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) WarehouseFbsCreatePickupTimeslotList(body easycb.AnyMap) (*WarehouseFbsCreatePickupTimeslotListRsp, error) {
	var result WarehouseFbsCreatePickupTimeslotListRsp
	err := c.doRequest("POST", "/v1/warehouse/fbs/create/pick-up/timeslot/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) WarehouseFbsUpdatePickupTimeslotList(body easycb.AnyMap) (*WarehouseFbsUpdatePickupTimeslotListRsp, error) {
	var result WarehouseFbsUpdatePickupTimeslotListRsp
	err := c.doRequest("POST", "/v1/warehouse/fbs/update/pick-up/timeslot/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) WarehouseFbsCreate(body easycb.AnyMap) (*WarehouseFbsCreateRsp, error) {
	var result WarehouseFbsCreateRsp
	err := c.doRequest("POST", "/v1/warehouse/fbs/create", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) WarehouseFbsUpdate(body easycb.AnyMap) (*WarehouseFbsUpdateRsp, error) {
	var result WarehouseFbsUpdateRsp
	err := c.doRequest("POST", "/v1/warehouse/fbs/update", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) WarehouseFbsFirstMileUpdate(body easycb.AnyMap) (*WarehouseFbsFirstMileUpdateRsp, error) {
	var result WarehouseFbsFirstMileUpdateRsp
	err := c.doRequest("POST", "/v1/warehouse/fbs/first-mile/update", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
