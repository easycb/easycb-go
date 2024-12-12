package lazada

import (
	"github.com/easycb/easycb-go"
)

func (c *Client) GetDocument(query easycb.AnyMap) (*GetDocumentRsp, error) {
	var result GetDocumentRsp
	err := c.doRequest("GET", "/order/document/get", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetMultipleOrderItems(query easycb.AnyMap) (*GetMultipleOrderItemsRsp, error) {
	var result GetMultipleOrderItemsRsp
	err := c.doRequest("GET", "/orders/items/get", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetOVOOrders(query easycb.AnyMap) (*GetOVOOrdersRsp, error) {
	var result GetOVOOrdersRsp
	err := c.doRequest("GET", "/orders/ovo/get", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetOrder(query easycb.AnyMap) (*GetOrderRsp, error) {
	var result GetOrderRsp
	err := c.doRequest("GET", "/order/get", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetOrderItems(query easycb.AnyMap) (*GetOrderItemsRsp, error) {
	var result GetOrderItemsRsp
	err := c.doRequest("GET", "/order/items/get", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetOrders(query easycb.AnyMap) (*GetOrdersRsp, error) {
	var result GetOrdersRsp
	err := c.doRequest("GET", "/orders/get", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) OrderCancelValidate(query easycb.AnyMap) (*OrderCancelValidateRsp, error) {
	var result OrderCancelValidateRsp
	err := c.doRequest("GET", "/order/reverse/cancel/validate", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SetInvoiceNumber(body easycb.AnyMap) (*SetInvoiceNumberRsp, error) {
	var result SetInvoiceNumberRsp
	err := c.doRequest("POST", "", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SetStatusToCanceled(body easycb.AnyMap) (*SetStatusToCanceledRsp, error) {
	var result SetStatusToCanceledRsp
	err := c.doRequest("POST", "/order/cancel", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
