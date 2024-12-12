package lazada

import "github.com/easycb/easycb-go"

func (c *Client) ConfirmDeliveryForDBSRsp(body easycb.AnyMap) (*ConfirmDeliveryForDBSRsp, error) {
	var result ConfirmDeliveryForDBSRsp
	err := c.doRequest("POST", "/order/package/sof/delivered", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeliverDigitalRsp(body easycb.AnyMap) (*DeliverDigitalRsp, error) {
	var result DeliverDigitalRsp
	err := c.doRequest("POST", "/order/digital/delivered", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FailedDeliveryForDBSRsp(body easycb.AnyMap) (*FailedDeliveryForDBSRsp, error) {
	var result FailedDeliveryForDBSRsp
	err := c.doRequest("POST", "/order/package/sof/failed_delivery", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetShipmentProviderRsp(body easycb.AnyMap) (*GetShipmentProviderRsp, error) {
	var result GetShipmentProviderRsp
	err := c.doRequest("POST", "/order/shipment/providers/get", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PackRsp(body easycb.AnyMap) (*PackRsp, error) {
	var result PackRsp
	err := c.doRequest("POST", "/order/fulfill/pack", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PackageStatusUpdateForDBSRsp(body easycb.AnyMap) (*PackageStatusUpdateForDBSRsp, error) {
	var result PackageStatusUpdateForDBSRsp
	err := c.doRequest("POST", "/order/package/sof/status/update", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PrintAWBRsp(body easycb.AnyMap) (*PrintAWBRsp, error) {
	var result PrintAWBRsp
	err := c.doRequest("POST", "/order/package/document/get", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ReadyToShipRsp(body easycb.AnyMap) (*ReadyToShipRsp, error) {
	var result ReadyToShipRsp
	err := c.doRequest("POST", "/order/package/rts", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) RecreatePackageRsp(body easycb.AnyMap) (*RecreatePackageRsp, error) {
	var result RecreatePackageRsp
	err := c.doRequest("POST", "/order/package/repack", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
