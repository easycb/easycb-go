package shopee

import (
	"github.com/easycb/easycb-go"
)

func (c *Client) GetOrderList(query easycb.AnyMap) (*GetOrderListRsp, error) {
	var result GetOrderListRsp
	err := c.doRequest("GET", "/api/v2/order/get_order_list", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetShipmentList(query easycb.AnyMap) (*GetShipmentListRsp, error) {
	var result GetShipmentListRsp
	err := c.doRequest("GET", "/api/v2/order/get_shipment_list", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetOrderDetail(query easycb.AnyMap) (*GetOrderDetailRsp, error) {
	var result GetOrderDetailRsp
	err := c.doRequest("GET", "/api/v2/order/get_order_detail", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SplitOrder(body easycb.AnyMap) (*SplitOrderRsp, error) {
	var result SplitOrderRsp
	err := c.doRequest("POST", "/api/v2/order/split_order", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UnSplitOrder(body easycb.AnyMap) (*UnSplitOrderRsp, error) {
	var result UnSplitOrderRsp
	err := c.doRequest("POST", "/api/v2/order/unsplit_order", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CancelOrder(body easycb.AnyMap) (*CancelOrderRsp, error) {
	var result CancelOrderRsp
	err := c.doRequest("POST", "/api/v2/order/cancel_order", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) HandleBuyerCancellation(body easycb.AnyMap) (*HandleBuyerCancellationRsp, error) {
	var result HandleBuyerCancellationRsp
	err := c.doRequest("POST", "/api/v2/order/handle_buyer_cancellation", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SetNote(body easycb.AnyMap) (*SetNoteRsp, error) {
	var result SetNoteRsp
	err := c.doRequest("POST", "/api/v2/order/set_note", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetPendingBuyerInvoiceOrderList(query easycb.AnyMap) (*GetPendingBuyerInvoiceOrderListRsp, error) {
	var result GetPendingBuyerInvoiceOrderListRsp
	err := c.doRequest("GET", "/api/v2/order/get_pending_buyer_invoice_order_list", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UploadInvoiceDoc(body easycb.AnyMap) (*UploadInvoiceDocRsp, error) {
	var result UploadInvoiceDocRsp
	err := c.doRequest("POST", "/api/v2/order/upload_invoice_doc", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DownloadInvoiceDoc(query easycb.AnyMap) (*DownloadInvoiceDocRsp, error) {
	var result DownloadInvoiceDocRsp
	err := c.doRequest("GET", "/api/v2/order/download_invoice_doc", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetBuyerInvoiceInfo(query easycb.AnyMap) (*GetBuyerInvoiceInfoRsp, error) {
	var result GetBuyerInvoiceInfoRsp
	err := c.doRequest("GET", "/api/v2/order/get_buyer_invoice_info", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
