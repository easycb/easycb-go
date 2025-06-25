package shein

import "github.com/easycb/easycb-go"

func (c *Client) WarehouseList(query easycb.AnyMap, headers map[string]string) (*WarehouseListRsp, error) {
	var result WarehouseListRsp
	err := c.doRequest("GET", "/open-api/msc/warehouse/list", query, nil, headers, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) StockQuery(body easycb.AnyMap) (*StockQueryRsp, error) {
	var result StockQueryRsp
	err := c.doRequest("POST", "/open-api/stock/stock-query", nil, body, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ChangeInventory(body easycb.AnyMap, headers map[string]string) (*ChangeInventoryRsp, error) {
	var result ChangeInventoryRsp
	err := c.doRequest("POST", "/open-api/gsp/goods/change-inventory", nil, body, headers, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) StockUpdate(body easycb.AnyMap) (*StockUpdateRsp, error) {
	var result StockUpdateRsp
	err := c.doRequest("POST", "/open-api/goods/stock-update", nil, body, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) QuerySkuSales(body easycb.AnyMap) (*QuerySkuSalesRsp, error) {
	var result QuerySkuSalesRsp
	err := c.doRequest("POST", "/open-api/goods/query-sku-sales", nil, body, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
