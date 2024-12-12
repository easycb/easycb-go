package tiktok

import (
	"fmt"
	"github.com/easycb/easycb-go"
)

func (c *Client) GetWarehouseList(query easycb.AnyMap) (*GetWarehouseListRsp, error) {
	var result GetWarehouseListRsp
	path := "/logistics/202309/warehouses"
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetGlobalSellerWarehouse(query easycb.AnyMap) (*GetGlobalSellerWarehouseRsp, error) {
	var result GetGlobalSellerWarehouseRsp
	path := "/logistics/202309/global_warehouses"
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetWarehouseDeliveryOptions(warehouseId string) (*GetWarehouseDeliveryOptionsRsp, error) {
	var result GetWarehouseDeliveryOptionsRsp
	path := fmt.Sprintf("/logistics/202309/warehouses/%s/delivery_options", warehouseId)
	err := c.doRequest("GET", path, nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetShippingProviders(deliveryOptionId string) (*GetShippingProvidersRsp, error) {
	var result GetShippingProvidersRsp
	path := fmt.Sprintf("/logistics/202309/delivery_options/%s/shipping_providers", deliveryOptionId)
	err := c.doRequest("GET", path, nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
