package shein

import "github.com/easycb/easycb-go"

func (c *Client) GetOrderList(body easycb.AnyMap) (*OrderListRsp, error) {
	var result OrderListRsp
	err := c.doRequest("POST", "/open-api/order/order-list", nil, body, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetOrderDetail(body easycb.AnyMap) (*OrderDetailRsp, error) {
	var result OrderDetailRsp
	err := c.doRequest("POST", "/open-api/order/order-detail", nil, body, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ExportAddress(body easycb.AnyMap) (*ExportAddressRsp, error) {
	var result ExportAddressRsp
	err := c.doRequest("POST", "/open-api/order/export-address", nil, body, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetExpressChannel(body easycb.AnyMap) (*GetExpressChannelRsp, error) {
	var result GetExpressChannelRsp
	err := c.doRequest("POST", "/open-api/order/express-channel", nil, body, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ImportBatchMultipleExpress(body easycb.AnyMap) (*ImportBatchMultipleExpressRsp, error) {
	var result ImportBatchMultipleExpressRsp
	err := c.doRequest("POST", "/open-api/order/import-batch-multiple-express", nil, body, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SyncInvoiceInfo(body easycb.AnyMap) (*SyncInvoiceInfoRsp, error) {
	var result SyncInvoiceInfoRsp
	err := c.doRequest("POST", "/open-api/order/sync-invoice-info", nil, body, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PrintExpressInfo(body easycb.AnyMap) (*PrintExpressInfoRsp, error) {
	var result PrintExpressInfoRsp
	err := c.doRequest("POST", "/open-api/order/print-express-info", nil, body, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) LogisticsTrack(query easycb.AnyMap, headers map[string]string) (*LogisticsTrackRsp, error) {
	var result LogisticsTrackRsp
	err := c.doRequest("GET", "/open-api/gsp/logistics-track", query, nil, headers, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ConfirmNoStock(body easycb.AnyMap) (*ConfirmNoStockRsp, error) {
	var result ConfirmNoStockRsp
	err := c.doRequest("POST", "/open-api/order/confirm-no-stock", nil, body, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UnpackingGroupRemove(body easycb.AnyMap) (*UnpackingGroupRemoveRsp, error) {
	var result UnpackingGroupRemoveRsp
	err := c.doRequest("POST", "/open-api/order/unpacking-group-remove", nil, body, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UnpackingGroupConfirm(body easycb.AnyMap) (*UnpackingGroupConfirmRsp, error) {
	var result UnpackingGroupConfirmRsp
	err := c.doRequest("POST", "/open-api/order/unpacking-group-confirm", nil, body, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
