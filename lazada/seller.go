package lazada

import "github.com/easycb/easycb-go"

func (c *Client) BatchQueryFollowStatus(query easycb.AnyMap) (*BatchQueryFollowStatusRsp, error) {
	var result BatchQueryFollowStatusRsp
	err := c.doRequest("GET", "/shop/follow/status/batch/query", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetPickUpStoreList(query easycb.AnyMap) (*GetPickUpStoreListRsp, error) {
	var result GetPickUpStoreListRsp
	err := c.doRequest("GET", "/rc/store/list/get", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetSeller(query easycb.AnyMap) (*GetSellerRsp, error) {
	var result GetSellerRsp
	err := c.doRequest("GET", "/seller/get", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetSellerMetricsById(query easycb.AnyMap) (*GetSellerMetricsByIdRsp, error) {
	var result GetSellerMetricsByIdRsp
	err := c.doRequest("GET", "/seller/metrics/get", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetSellerPerformance(query easycb.AnyMap) (*GetSellerPerformanceRsp, error) {
	var result GetSellerPerformanceRsp
	err := c.doRequest("GET", "/seller/performance/get", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetWarehouseBySellerId(query easycb.AnyMap) (*GetWarehouseBySellerIdRsp, error) {
	var result GetWarehouseBySellerIdRsp
	err := c.doRequest("GET", "/rc/warehouse/get", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) QueryWarehouseDetailInfoBySellerId(query easycb.AnyMap) (*QueryWarehouseDetailInfoBySellerIdRsp, error) {
	var result QueryWarehouseDetailInfoBySellerIdRsp
	err := c.doRequest("GET", "/rc/warehouse/detail/get", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SellerCenterMsgList(query easycb.AnyMap) (*SellerCenterMsgListRsp, error) {
	var result SellerCenterMsgListRsp
	err := c.doRequest("GET", "/sellercenter/msg/list", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SellerPolicyFetch(query easycb.AnyMap) (*SellerPolicyFetchRsp, error) {
	var result SellerPolicyFetchRsp
	err := c.doRequest("GET", "/seller/policy/fetch", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SynchronizeSellerItemArConfig(query easycb.AnyMap) (*SynchronizeSellerItemArConfigRsp, error) {
	var result SynchronizeSellerItemArConfigRsp
	err := c.doRequest("GET", "/seller/ar/config/syn", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetCountryInfo(query easycb.AnyMap) (*GetCountryInfoRsp, error) {
	var result GetCountryInfoRsp
	err := c.doRequest("GET", "/seller/cb/country/get", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetSellerRegisterInfo(query easycb.AnyMap) (*GetSellerRegisterInfoRsp, error) {
	var result GetSellerRegisterInfoRsp
	err := c.doRequest("GET", "/seller/cb/register/info", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetSubAddress(query easycb.AnyMap) (*GetSubAddressRsp, error) {
	var result GetSubAddressRsp
	err := c.doRequest("GET", "/seller/cb/country/location/get", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PaymentBinding(body easycb.AnyMap) (*PaymentBindingRsp, error) {
	var result PaymentBindingRsp
	err := c.doRequest("POST", "/seller/cb/payment/config", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) QueryBuyboxHuntingInfo(body easycb.AnyMap) (*QueryBuyboxHuntingInfoRsp, error) {
	var result QueryBuyboxHuntingInfoRsp
	err := c.doRequest("POST", "/hunting/buybox/get", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SaveSellerWarehouseInfo(body easycb.AnyMap) (*SaveSellerWarehouseInfoRsp, error) {
	var result SaveSellerWarehouseInfoRsp
	err := c.doRequest("POST", "/rc/sellerWarehouse/saveWarehouseInfo", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SellerFieldVerify(body easycb.AnyMap) (*SellerFieldVerifyRsp, error) {
	var result SellerFieldVerifyRsp
	err := c.doRequest("POST", "/seller/cb/register/fieldcheck", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
