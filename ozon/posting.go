package ozon

import "github.com/easycb/easycb-go"

func (c *Client) PostingFbsUnfulfilledList(body easycb.AnyMap) (*PostingFbsUnfulfilledListRsp, error) {
	var result PostingFbsUnfulfilledListRsp
	err := c.doRequest("POST", "/v4/posting/fbs/unfulfilled/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PostingFbsList(body easycb.AnyMap) (*PostingFbsListRsp, error) {
	var result PostingFbsListRsp
	err := c.doRequest("POST", "/v4/posting/fbs/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PostingFbsGet(body easycb.AnyMap) (*PostingFbsGetRsp, error) {
	var result PostingFbsGetRsp
	err := c.doRequest("POST", "/v3/posting/fbs/get", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PostingFbsGetByBarcode(body easycb.AnyMap) (*PostingFbsGetByBarcodeRsp, error) {
	var result PostingFbsGetByBarcodeRsp
	err := c.doRequest("POST", "/v2/posting/fbs/get-by-barcode", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PostingFbsProductCountryList(body easycb.AnyMap) (*PostingFbsProductCountryListRsp, error) {
	var result PostingFbsProductCountryListRsp
	err := c.doRequest("POST", "/v2/posting/fbs/product/country/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PostingFbsProductCountrySet(body easycb.AnyMap) (*PostingFbsProductCountrySetRsp, error) {
	var result PostingFbsProductCountrySetRsp
	err := c.doRequest("POST", "/v2/posting/fbs/product/country/set", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PostingFbsPackageLabel(body easycb.AnyMap) (*PostingFbsPackageLabelRsp, error) {
	var result PostingFbsPackageLabelRsp
	err := c.doRequest("POST", "/v2/posting/fbs/package-label", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PostingFbsAwaitingDelivery(body easycb.AnyMap) (*PostingFbsAwaitingDeliveryRsp, error) {
	var result PostingFbsAwaitingDeliveryRsp
	err := c.doRequest("POST", "/v2/posting/fbs/awaiting-delivery", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PostingFbsCancelReasonList(body easycb.AnyMap) (*PostingFbsCancelReasonListRsp, error) {
	var result PostingFbsCancelReasonListRsp
	err := c.doRequest("POST", "/v2/posting/fbs/cancel-reason/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PostingFbsCancelReason(body easycb.AnyMap) (*PostingFbsCancelReasonRsp, error) {
	var result PostingFbsCancelReasonRsp
	err := c.doRequest("POST", "/v1/posting/fbs/cancel-reason", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PostingFbsCancel(body easycb.AnyMap) (*PostingFbsCancelRsp, error) {
	var result PostingFbsCancelRsp
	err := c.doRequest("POST", "/v2/posting/fbs/cancel", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PostingFbsProductCancel(body easycb.AnyMap) (*PostingFbsProductCancelRsp, error) {
	var result PostingFbsProductCancelRsp
	err := c.doRequest("POST", "/v2/posting/fbs/product/cancel", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PostingFbsShip(body easycb.AnyMap) (*PostingFbsShipRsp, error) {
	var result PostingFbsShipRsp
	err := c.doRequest("POST", "/v4/posting/fbs/ship", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PostingFbsShipPackage(body easycb.AnyMap) (*PostingFbsShipPackageRsp, error) {
	var result PostingFbsShipPackageRsp
	err := c.doRequest("POST", "/v4/posting/fbs/ship/package", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PostingFbsSplit(body easycb.AnyMap) (*PostingFbsSplitRsp, error) {
	var result PostingFbsSplitRsp
	err := c.doRequest("POST", "/v1/posting/fbs/split", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
