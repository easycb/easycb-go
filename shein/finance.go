package shein

import "github.com/easycb/easycb-go"

func (c *Client) FinanceReportList(body easycb.AnyMap, headers map[string]string) (*FinanceReportListRsp, error) {
	var result FinanceReportListRsp
	err := c.doRequest("POST", "/open-api/finance/report-list", nil, body, headers, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FinanceReportSalesDetail(body easycb.AnyMap, headers map[string]string) (*FinanceReportSalesDetailRsp, error) {
	var result FinanceReportSalesDetailRsp
	err := c.doRequest("POST", "/open-api/finance/report-sales-detail", nil, body, headers, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FinanceAdjustmentDetail(body easycb.AnyMap, headers map[string]string) (*FinanceAdjustmentDetailRsp, error) {
	var result FinanceAdjustmentDetailRsp
	err := c.doRequest("POST", "/open-api/finance/report-adjustment-detail", nil, body, headers, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FinanceGetCheckOrderList(body easycb.AnyMap) (*FinanceGetCheckOrderListRsp, error) {
	var result FinanceGetCheckOrderListRsp
	err := c.doRequest("POST", "/open-api/finance/get-check-order-list", nil, body, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FinanceGetCheckOrderDetail(query easycb.AnyMap) (*FinanceGetCheckOrderDetailRsp, error) {
	var result FinanceGetCheckOrderDetailRsp
	err := c.doRequest("GET", "/open-api/finance/get-check-order-detail", query, nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
