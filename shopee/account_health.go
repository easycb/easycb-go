package shopee

import "github.com/easycb/easycb-go"

func (c *Client) ShopPenalty() (*ShopPenaltyRsp, error) {
	var result ShopPenaltyRsp
	err := c.doRequest("GET", "/api/v2/account_health/shop_penalty", nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetShopPerformance() (*GetShopPerformanceRsp, error) {
	var result GetShopPerformanceRsp
	err := c.doRequest("GET", "/api/v2/account_health/get_shop_performance", nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetMetricSourceDetail(query easycb.AnyMap) (*GetMetricSourceDetailRsp, error) {
	var result GetMetricSourceDetailRsp
	err := c.doRequest("GET", "/api/v2/account_health/get_metric_source_detail", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetPenaltyPointHistory(query easycb.AnyMap) (*GetPenaltyPointHistoryRsp, error) {
	var result GetPenaltyPointHistoryRsp
	err := c.doRequest("GET", "/api/v2/account_health/get_penalty_point_history", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetPunishmentHistory(query easycb.AnyMap) (*GetPunishmentHistoryRsp, error) {
	var result GetPunishmentHistoryRsp
	err := c.doRequest("GET", "/api/v2/account_health/get_punishment_history", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetListingsWithIssues(query easycb.AnyMap) (*GetListingsWithIssuesRsp, error) {
	var result GetListingsWithIssuesRsp
	err := c.doRequest("GET", "/api/v2/account_health/get_listings_with_issues", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetLateOrders(query easycb.AnyMap) (*GetLateOrdersRsp, error) {
	var result GetLateOrdersRsp
	err := c.doRequest("GET", "/api/v2/account_health/get_late_orders", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
