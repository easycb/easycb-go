package shopee

import "github.com/easycb/easycb-go"

func (c *Client) GetShippingParameter(query easycb.AnyMap) (*GetShippingParameterRsp, error) {
	var result GetShippingParameterRsp
	err := c.doRequest("GET", "/api/v2/logistics/get_shipping_parameter", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetTrackingNumber(query easycb.AnyMap) (*GetTrackingNumberRsp, error) {
	var result GetTrackingNumberRsp
	err := c.doRequest("GET", "/api/v2/logistics/get_tracking_number", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ShipOrder(body easycb.AnyMap) (*ShipOrderRsp, error) {
	var result ShipOrderRsp
	err := c.doRequest("POST", "/api/v2/logistics/ship_order", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateShippingOrder(body easycb.AnyMap) (*UpdateShippingOrderRsp, error) {
	var result UpdateShippingOrderRsp
	err := c.doRequest("POST", "/api/v2/logistics/update_shipping_order", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetShippingDocumentParameter(body easycb.AnyMap) (*GetShippingDocumentParameterRsp, error) {
	var result GetShippingDocumentParameterRsp
	err := c.doRequest("POST", "/api/v2/logistics/get_shipping_document_parameter", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CreateShippingDocument(body easycb.AnyMap) (*CreateShippingDocumentRsp, error) {
	var result CreateShippingDocumentRsp
	err := c.doRequest("POST", "/api/v2/logistics/create_shipping_document", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetShippingDocumentResult(body easycb.AnyMap) (*GetShippingDocumentResultRsp, error) {
	var result GetShippingDocumentResultRsp
	err := c.doRequest("POST", "/api/v2/logistics/get_shipping_document_result", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DownloadShippingDocument(body easycb.AnyMap) (*DownloadShippingDocumentRsp, error) {
	var result DownloadShippingDocumentRsp
	err := c.doRequest("POST", "/api/v2/logistics/download_shipping_document", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetTrackingInfo(query easycb.AnyMap) (*GetTrackingInfoRsp, error) {
	var result GetTrackingInfoRsp
	err := c.doRequest("GET", "/api/v2/logistics/get_tracking_info", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetAddressList(query easycb.AnyMap) (*GetAddressListRsp, error) {
	var result GetAddressListRsp
	err := c.doRequest("GET", "/api/v2/logistics/get_address_list", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SetAddressConfig(body easycb.AnyMap) (*SetAddressConfigRsp, error) {
	var result SetAddressConfigRsp
	err := c.doRequest("POST", "/api/v2/logistics/set_address_config", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteAddress(body easycb.AnyMap) (*DeleteAddressRsp, error) {
	var result DeleteAddressRsp
	err := c.doRequest("POST", "/api/v2/logistics/delete_address", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetChannelList(query easycb.AnyMap) (*GetChannelListRsp, error) {
	var result GetChannelListRsp
	err := c.doRequest("GET", "/api/v2/logistics/get_channel_list", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateChannel(body easycb.AnyMap) (*UpdateChannelRsp, error) {
	var result UpdateChannelRsp
	err := c.doRequest("POST", "/api/v2/logistics/update_channel", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) BatchShipOrder(body easycb.AnyMap) (*BatchShipOrderRsp, error) {
	var result BatchShipOrderRsp
	err := c.doRequest("POST", "/api/v2/logistics/batch_ship_order", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetShippingDocumentDataInfo(body easycb.AnyMap) (*GetShippingDocumentDataInfoRsp, error) {
	var result GetShippingDocumentDataInfoRsp
	err := c.doRequest("POST", "/api/v2/logistics/get_shipping_document_data_info", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateTrackingStatus(body easycb.AnyMap) (*UpdateTrackingStatusRsp, error) {
	var result UpdateTrackingStatusRsp
	err := c.doRequest("POST", "/api/v2/logistics/update_tracking_status", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
