package lazada

import "github.com/easycb/easycb-go"

func (c *Client) AddOrUpdatePickupStop(body easycb.AnyMap) (*AddOrUpdatePickupStopRsp, error) {
	var result AddOrUpdatePickupStopRsp
	err := c.doRequest("POST", "/logistics/tps/runsheets/stops", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) Create3PLStation(body easycb.AnyMap) (*Create3PLStationRsp, error) {
	var result Create3PLStationRsp
	err := c.doRequest("POST", "/logistics/tps/stations/create", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetOrderTrace(body easycb.AnyMap) (*GetOrderTraceRsp, error) {
	var result GetOrderTraceRsp
	err := c.doRequest("POST", "/logistic/order/trace", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ScanParcel(body easycb.AnyMap) (*ScanParcelRsp, error) {
	var result ScanParcelRsp
	err := c.doRequest("POST", "/dop/scan", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) StationDopScan(body easycb.AnyMap) (*StationDopScanRsp, error) {
	var result StationDopScanRsp
	err := c.doRequest("POST", "/stations/dop/scan", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) Update3PLStation(body easycb.AnyMap) (*Update3PLStationRsp, error) {
	var result Update3PLStationRsp
	err := c.doRequest("POST", "/logistics/tps/stations/update", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdatePickupTimeSlot(body easycb.AnyMap) (*UpdatePickupTimeSlotRsp, error) {
	var result UpdatePickupTimeSlotRsp
	err := c.doRequest("POST", "/logistics/tps/sellers/pickup_timeslot", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CreateConsolidationService(body easycb.AnyMap) (*CreateConsolidationServiceRsp, error) {
	var result CreateConsolidationServiceRsp
	err := c.doRequest("POST", "/logistics/ldp/createConsolidationService", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateLastMile(body easycb.AnyMap) (*UpdateLastMileRsp, error) {
	var result UpdateLastMileRsp
	err := c.doRequest("POST", "/logistics/ldp/updateLastmile", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
