package tiktok

import (
	"fmt"
	"github.com/easycb/easycb-go"
)

func (c *Client) GetOrderSplitAttributes(query easycb.AnyMap) (*GetOrderSplitAttributesRsp, error) {
	var result GetOrderSplitAttributesRsp
	path := "/fulfillment/202309/orders/split_attributes"
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SplitOrders(orderId string, query easycb.AnyMap, body easycb.AnyMap) (*SplitOrdersRsp, error) {
	var result SplitOrdersRsp
	path := fmt.Sprintf("/fulfillment/202309/orders/%s/split", orderId)
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetEligibleShippingService(orderId string, query easycb.AnyMap, body easycb.AnyMap) (*GetEligibleShippingServiceRsp, error) {
	var result GetEligibleShippingServiceRsp
	path := fmt.Sprintf("/fulfillment/202309/orders/%s/shipping_services/query", orderId)
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CreateFirstMileBundle(body easycb.AnyMap) (*CreateFirstMileBundleRsp, error) {
	var result CreateFirstMileBundleRsp
	path := "/fulfillment/202407/bundles"
	err := c.doRequest("POST", path, nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CreatePackages(query easycb.AnyMap, body easycb.AnyMap) (*CreatePackagesRsp, error) {
	var result CreatePackagesRsp
	path := "/fulfillment/202309/packages"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SearchPackage(query easycb.AnyMap, body easycb.AnyMap) (*SearchPackageRsp, error) {
	var result SearchPackageRsp
	path := "/fulfillment/202309/packages/search"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SearchCombinablePackages(query easycb.AnyMap) (*SearchCombinablePackagesRsp, error) {
	var result SearchCombinablePackagesRsp
	path := "/fulfillment/202309/combinable_packages/search"
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CombinePackage(query easycb.AnyMap, body easycb.AnyMap) (*CombinePackageRsp, error) {
	var result CombinePackageRsp
	path := "/fulfillment/202309/packages/combine"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UnCombinePackages(packageId string, query easycb.AnyMap, body easycb.AnyMap) (*UnCombinePackagesRsp, error) {
	var result UnCombinePackagesRsp
	path := fmt.Sprintf("/fulfillment/202309/packages/%s/uncombine", packageId)
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetPackageHandoverTimeSlots(packageId string, query easycb.AnyMap) (*GetPackageHandoverTimeSlotsRsp, error) {
	var result GetPackageHandoverTimeSlotsRsp
	path := fmt.Sprintf("/fulfillment/202309/packages/%s/handover_time_slots", packageId)
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ShipPackage(packageId string, query easycb.AnyMap, body easycb.AnyMap) (*ShipPackageRsp, error) {
	var result ShipPackageRsp
	path := fmt.Sprintf("/fulfillment/202309/packages/%s/ship", packageId)
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) BatchShipPackages(query easycb.AnyMap, body easycb.AnyMap) (*BatchShipPackagesRsp, error) {
	var result BatchShipPackagesRsp
	path := "/fulfillment/202309/packages/ship"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) MarkPackageAsShipped(orderId string, query easycb.AnyMap, body easycb.AnyMap) (*MarkPackageAsShippedRsp, error) {
	var result MarkPackageAsShippedRsp
	path := fmt.Sprintf("/fulfillment/202309/orders/%s/packages", orderId)
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetPackageShippingDocument(packageId string, query easycb.AnyMap) (*GetPackageShippingDocumentRsp, error) {
	var result GetPackageShippingDocumentRsp
	path := fmt.Sprintf("/fulfillment/202309/packages/%s/shipping_documents", packageId)
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetPackageDetail(packageId string, query easycb.AnyMap) (*GetPackageDetailRsp, error) {
	var result GetPackageDetailRsp
	path := fmt.Sprintf("/fulfillment/202309/packages/%s", packageId)
	err := c.doRequest("GET", path, nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetTracking(orderId string) (*GetTrackingRsp, error) {
	var result GetTrackingRsp
	path := fmt.Sprintf("/fulfillment/202309/orders/%s/tracking", orderId)
	err := c.doRequest("GET", path, nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateShippingInfo(orderId string, query easycb.AnyMap, body easycb.AnyMap) (*UpdateShippingInfoRsp, error) {
	var result UpdateShippingInfoRsp
	path := fmt.Sprintf("/fulfillment/202309/orders/%s/shipping_info/update", orderId)
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdatePackageShippingInfo(packageId string, query easycb.AnyMap, body easycb.AnyMap) (*UpdatePackageShippingInfoRsp, error) {
	var result UpdatePackageShippingInfoRsp
	path := fmt.Sprintf("/fulfillment/202309/packages/%s/shipping_info/update", packageId)
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FulfillmentUploadDeliveryFile(query easycb.AnyMap, body easycb.AnyMap) (*FulfillmentUploadDeliveryFileRsp, error) {
	var result FulfillmentUploadDeliveryFileRsp
	path := "/fulfillment/202309/files/upload"
	err := c.doFileRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FulfillmentUploadDeliveryImage(query easycb.AnyMap, body easycb.AnyMap) (*FulfillmentUploadDeliveryImageRsp, error) {
	var result FulfillmentUploadDeliveryImageRsp
	path := "/fulfillment/202309/images/upload"
	err := c.doFileRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdatePackageDeliveryStatus(query easycb.AnyMap, body easycb.AnyMap) (*UpdatePackageDeliveryStatusRsp, error) {
	var result UpdatePackageDeliveryStatusRsp
	path := "/fulfillment/202309/packages/deliver"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
