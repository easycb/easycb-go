package tiktok

import (
	"fmt"
	"github.com/easycb/easycb-go"
)

func (c *Client) GetOrderList(query easycb.AnyMap, body easycb.AnyMap) (*GetOrderListRsp, error) {
	var result GetOrderListRsp
	path := "/order/202309/orders/search"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetOrderDetail(query easycb.AnyMap) (*GetOrderDetailRsp, error) {
	var result GetOrderDetailRsp
	path := "/order/202309/orders"
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetOrderDetailV202507(query easycb.AnyMap) (*GetOrderDetailV202507Rsp, error) {
	var result GetOrderDetailV202507Rsp
	path := "/order/202507/orders"
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetPriceDetail(orderId string, query easycb.AnyMap) (*GetPriceDetailRsp, error) {
	var result GetPriceDetailRsp
	path := fmt.Sprintf("/order/202407/orders/%s/price_detail", orderId)
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) AddExternalOrderReferences(query easycb.AnyMap, body easycb.AnyMap) (*AddExternalOrderReferencesRsp, error) {
	var result AddExternalOrderReferencesRsp
	path := "/order/202406/orders/external_orders"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetExternalOrderReferences(orderId string, query easycb.AnyMap) (*GetExternalOrderReferencesRsp, error) {
	var result GetExternalOrderReferencesRsp
	path := fmt.Sprintf("/order/202406/orders/%s/external_orders", orderId)
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SearchOrderByExternalOrderReference(query easycb.AnyMap) (*SearchOrderByExternalOrderReferenceRsp, error) {
	var result SearchOrderByExternalOrderReferenceRsp
	path := "/order/202406/orders/external_order_search"
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
