package tiktok

import (
	"fmt"
	"github.com/easycb/easycb-go"
)

func (c *Client) GetShopPerformance(query easycb.AnyMap) (*GetShopPerformanceRsp, error) {
	var result GetShopPerformanceRsp
	path := "/analytics/202405/shop/performance"
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetShopProductPerformance(productId string, query easycb.AnyMap) (*GetShopProductPerformanceRsp, error) {
	var result GetShopProductPerformanceRsp
	path := fmt.Sprintf("/analytics/202405/shop_products/%s/performance", productId)
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetShopProductPerformanceList(query easycb.AnyMap) (*GetShopProductPerformanceListRsp, error) {
	var result GetShopProductPerformanceListRsp
	path := "/analytics/202405/shop_products/performance"
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetShopSKUPerformance(skuId string, query easycb.AnyMap) (*GetShopSKUPerformanceRsp, error) {
	var result GetShopSKUPerformanceRsp
	path := fmt.Sprintf("/analytics/202406/shop_skus/%s/performance", skuId)
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetShopSKUPerformanceList(query easycb.AnyMap) (*GetShopSKUPerformanceListRsp, error) {
	var result GetShopSKUPerformanceListRsp
	path := "/analytics/202406/shop_skus/performance"
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetShopVideoPerformanceList(query easycb.AnyMap) (*GetShopVideoProductPerformanceListRsp, error) {
	var result GetShopVideoProductPerformanceListRsp
	path := "/analytics/202409/shop_videos/performance"
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetShopVideoPerformanceOverview(query easycb.AnyMap) (*GetShopVideoPerformanceOverviewRsp, error) {
	var result GetShopVideoPerformanceOverviewRsp
	path := "/analytics/202409/shop_videos/overview_performance"
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetShopVideoPerformanceDetails(videoId string, query easycb.AnyMap) (*GetShopVideoPerformanceDetailsRsp, error) {
	var result GetShopVideoPerformanceDetailsRsp
	path := fmt.Sprintf("/analytics/202409/shop_videos/%s/performance", videoId)
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetShopVideoProductPerformanceList(videoId string, query easycb.AnyMap) (*GetShopVideoProductPerformanceListRsp, error) {
	var result GetShopVideoProductPerformanceListRsp
	path := fmt.Sprintf("/analytics/202409/shop_videos/%s/products/performance", videoId)
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
