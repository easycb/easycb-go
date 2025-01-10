package tiktok

import "github.com/easycb/easycb-go"

func (c *Client) GetFBTMerchantOnboardedRegions(query easycb.AnyMap) (*GetFBTMerchantOnboardedRegionsRsp, error) {
	var result GetFBTMerchantOnboardedRegionsRsp
	path := "/fbt/202409/merchants/onboarded_regions"
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetFBTWarehouseList(query easycb.AnyMap) (*GetFBTWarehouseListRsp, error) {
	var result GetFBTWarehouseListRsp
	path := "/fbt/202408/warehouses"
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetInboundOrder(query easycb.AnyMap) (*GetInboundOrderRsp, error) {
	var result GetInboundOrderRsp
	path := "/fbt/202409/inbound_orders"
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SearchFBTInventory(query easycb.AnyMap) (*SearchFBTInventoryRsp, error) {
	var result SearchFBTInventoryRsp
	path := "/fbt/202408/inventory/search"
	err := c.doRequest("GET", path, nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SearchFBTInventoryRecord(query easycb.AnyMap, body easycb.AnyMap) (*SearchFBTInventoryRecordRsp, error) {
	var result SearchFBTInventoryRecordRsp
	path := "/fbt/202410/inventory_records/search"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SearchGoodsInfo(query easycb.AnyMap, body easycb.AnyMap) (*SearchGoodsInfoRsp, error) {
	var result SearchGoodsInfoRsp
	path := "/fbt/202409/goods/search"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
