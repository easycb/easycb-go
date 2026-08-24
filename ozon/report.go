package ozon

import "github.com/easycb/easycb-go"

func (c *Client) ReportInfo(body easycb.AnyMap) (*ReportInfoRsp, error) {
	var result ReportInfoRsp
	err := c.doRequest("POST", "/v1/report/info", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ReportList(body easycb.AnyMap) (*ReportListRsp, error) {
	var result ReportListRsp
	err := c.doRequest("POST", "/v1/report/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ReportProductsCreate(body easycb.AnyMap) (*ReportProductsCreateRsp, error) {
	var result ReportProductsCreateRsp
	err := c.doRequest("POST", "/v1/report/products/create", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ReportPostingCreate(body easycb.AnyMap) (*ReportPostingCreateRsp, error) {
	var result ReportPostingCreateRsp
	err := c.doRequest("POST", "/v1/report/postings/create", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FinanceCashFlowStatementList(body easycb.AnyMap) (*FinanceCashFlowStatementListRsp, error) {
	var result FinanceCashFlowStatementListRsp
	err := c.doRequest("POST", "/v1/finance/cash-flow-statement/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ReportDiscountedCreate(body easycb.AnyMap) (*ReportDiscountedCreateRsp, error) {
	var result ReportDiscountedCreateRsp
	err := c.doRequest("POST", "/v1/report/discounted/create", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ReportWarehouseStock(body easycb.AnyMap) (*ReportWarehouseStockRsp, error) {
	var result ReportWarehouseStockRsp
	err := c.doRequest("POST", "/v1/report/warehouse/stock", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ReportMarkedProductsSalesCreate(body easycb.AnyMap) (*ReportMarkedProductsSalesCreateRsp, error) {
	var result ReportMarkedProductsSalesCreateRsp
	err := c.doRequest("POST", "/v1/report/marked-products-sales/create", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ReportRealizationPostingCreate(body easycb.AnyMap) (*ReportRealizationPostingCreateRsp, error) {
	var result ReportRealizationPostingCreateRsp
	err := c.doRequest("POST", "/v1/report/realization/posting/create", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
