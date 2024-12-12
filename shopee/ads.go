package shopee

import "github.com/easycb/easycb-go"

func (c *Client) GetTotalBalance() (*GetTotalBalanceRsp, error) {
	var result GetTotalBalanceRsp
	err := c.doRequest("GET", "/api/v2/ads/get_total_balance", nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetShopToggleInfo() (*GetShopToggleInfoRsp, error) {
	var result GetShopToggleInfoRsp
	err := c.doRequest("GET", "/api/v2/ads/get_shop_toggle_info", nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetRecommendedKeywordList(query easycb.AnyMap) (*GetRecommendedKeywordListRsp, error) {
	var result GetRecommendedKeywordListRsp
	err := c.doRequest("GET", "/api/v2/ads/get_recommended_keyword_list", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetRecommendedItemList() (*GetRecommendedItemListRsp, error) {
	var result GetRecommendedItemListRsp
	err := c.doRequest("GET", "/api/v2/ads/get_recommended_item_list", nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetAllCpcAdsHourlyPerformance(query easycb.AnyMap) (*GetAllCpcAdsHourlyPerformanceRsp, error) {
	var result GetAllCpcAdsHourlyPerformanceRsp
	err := c.doRequest("GET", "/api/v2/ads/get_all_cpc_ads_hourly_performance", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetAllCpcAdsDailyPerformance(query easycb.AnyMap) (*GetAllCpcAdsDailyPerformanceRsp, error) {
	var result GetAllCpcAdsDailyPerformanceRsp
	err := c.doRequest("GET", "/api/v2/ads/get_all_cpc_ads_daily_performance", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetProductCampaignDailyPerformance(query easycb.AnyMap) (*GetProductCampaignDailyPerformanceRsp, error) {
	var result GetProductCampaignDailyPerformanceRsp
	err := c.doRequest("GET", "/api/v2/ads/get_product_campaign_daily_performance", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetProductCampaignHourlyPerformance(query easycb.AnyMap) (*GetProductCampaignHourlyPerformanceRsp, error) {
	var result GetProductCampaignHourlyPerformanceRsp
	err := c.doRequest("GET", "/api/v2/ads/get_product_campaign_hourly_performance", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetProductLevelCampaignIdList(query easycb.AnyMap) (*GetProductLevelCampaignIdListRsp, error) {
	var result GetProductLevelCampaignIdListRsp
	err := c.doRequest("GET", "/api/v2/ads/get_product_level_campaign_id_list", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetProductLevelCampaignSett(query easycb.AnyMap) (*GetProductLevelCampaignSettRsp, error) {
	var result GetProductLevelCampaignSettRsp
	err := c.doRequest("GET", "/api/v2/ads/get_product_level_campaign_setting_info", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
