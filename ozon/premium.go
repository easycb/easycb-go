package ozon

import "github.com/easycb/easycb-go"

func (c *Client) AnalyticsData(body easycb.AnyMap) (*AnalyticsDataRsp, error) {
	var result AnalyticsDataRsp
	err := c.doRequest("POST", "/v1/analytics/data", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) AnalyticsProductQueries(body easycb.AnyMap) (*AnalyticsProductQueriesRsp, error) {
	var result AnalyticsProductQueriesRsp
	err := c.doRequest("POST", "/v1/analytics/product-queries", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) AnalyticsProductQueriesDetails(body easycb.AnyMap) (*AnalyticsProductQueriesDetailsRsp, error) {
	var result AnalyticsProductQueriesDetailsRsp
	err := c.doRequest("POST", "/v1/analytics/product-queries/details", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FinanceRealizationByDay(body easycb.AnyMap) (*FinanceRealizationByDayRsp, error) {
	var result FinanceRealizationByDayRsp
	err := c.doRequest("POST", "/v1/finance/realization/by-day", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SearchQueriesText(body easycb.AnyMap) (*SearchQueriesTextRsp, error) {
	var result SearchQueriesTextRsp
	err := c.doRequest("POST", "/v1/search-queries/text", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SearchQueriesTop(body easycb.AnyMap) (*SearchQueriesTopRsp, error) {
	var result SearchQueriesTopRsp
	err := c.doRequest("POST", "/v1/search-queries/top", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ProductPricesDetails(body easycb.AnyMap) (*ProductPricesDetailsRsp, error) {
	var result ProductPricesDetailsRsp
	err := c.doRequest("POST", "/v1/product/prices/details", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
