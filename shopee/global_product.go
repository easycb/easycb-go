package shopee

import "github.com/easycb/easycb-go"

func (c *Client) GlobalGetCategory(query easycb.AnyMap) (*GlobalGetCategoryRsp, error) {
	var result GlobalGetCategoryRsp
	err := c.doRequest("GET", "/api/v2/global_product/get_category", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GlobalGetAttributeTree(query easycb.AnyMap) (*GlobalGetAttributeTreeRsp, error) {
	var result GlobalGetAttributeTreeRsp
	err := c.doRequest("GET", "/api/v2/global_product/get_attribute_tree", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GlobalGetBrandList(query easycb.AnyMap) (*GlobalGetBrandListRsp, error) {
	var result GlobalGetBrandListRsp
	err := c.doRequest("GET", "/api/v2/global_product/get_brand_list", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetGlobalItemLimit(query easycb.AnyMap) (*GetGlobalItemLimitRsp, error) {
	var result GetGlobalItemLimitRsp
	err := c.doRequest("GET", "/api/v2/global_product/get_global_item_limit", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetGlobalItemList(query easycb.AnyMap) (*GetGlobalItemListRsp, error) {
	var result GetGlobalItemListRsp
	err := c.doRequest("GET", "/api/v2/global_product/get_global_item_list", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetGlobalItemInfo(query easycb.AnyMap) (*GetGlobalItemInfoRsp, error) {
	var result GetGlobalItemInfoRsp
	err := c.doRequest("GET", "/api/v2/global_product/get_global_item_info", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) AddGlobalItem(body easycb.AnyMap) (*AddGlobalItemRsp, error) {
	var result AddGlobalItemRsp
	err := c.doRequest("POST", "/api/v2/global_product/add_global_item", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateGlobalItem(body easycb.AnyMap) (*UpdateGlobalItemRsp, error) {
	var result UpdateGlobalItemRsp
	err := c.doRequest("POST", "/api/v2/global_product/update_global_item", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteGlobalItem(body easycb.AnyMap) (*DeleteGlobalItemRsp, error) {
	var result DeleteGlobalItemRsp
	err := c.doRequest("POST", "/api/v2/global_product/delete_global_item", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GlobalInitTierVariation(body easycb.AnyMap) (*GlobalInitTierVariationRsp, error) {
	var result GlobalInitTierVariationRsp
	err := c.doRequest("POST", "/api/v2/global_product/init_tier_variation", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GlobalUpdateTierVariation(body easycb.AnyMap) (*GlobalUpdateTierVariationRsp, error) {
	var result GlobalUpdateTierVariationRsp
	err := c.doRequest("POST", "/api/v2/global_product/update_tier_variation", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) AddGlobalModel(body easycb.AnyMap) (*AddGlobalModelRsp, error) {
	var result AddGlobalModelRsp
	err := c.doRequest("POST", "/api/v2/global_product/add_global_model", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateGlobalModel(body easycb.AnyMap) (*UpdateGlobalModelRsp, error) {
	var result UpdateGlobalModelRsp
	err := c.doRequest("POST", "/api/v2/global_product/update_global_model", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteGlobalModel(body easycb.AnyMap) (*DeleteGlobalModelRsp, error) {
	var result DeleteGlobalModelRsp
	err := c.doRequest("POST", "/api/v2/global_product/delete_global_model", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetGlobalModelList(query easycb.AnyMap) (*GetGlobalModelListRsp, error) {
	var result GetGlobalModelListRsp
	err := c.doRequest("GET", "/api/v2/global_product/get_global_model_list", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GlobalSupportSizeChart(query easycb.AnyMap) (*GlobalSupportSizeChartRsp, error) {
	var result GlobalSupportSizeChartRsp
	err := c.doRequest("POST", "/api/v2/global_product/support_size_chart", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GlobalUpdateSizeChart(body easycb.AnyMap) (*GlobalUpdateSizeChartRsp, error) {
	var result GlobalUpdateSizeChartRsp
	err := c.doRequest("POST", "/api/v2/global_product/update_size_chart", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CreatePublishTask(body easycb.AnyMap) (*CreatePublishTaskRsp, error) {
	var result CreatePublishTaskRsp
	err := c.doRequest("POST", "/api/v2/global_product/create_publish_task", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetPublishableShop(query easycb.AnyMap) (*GetPublishableShopRsp, error) {
	var result GetPublishableShopRsp
	err := c.doRequest("GET", "/api/v2/global_product/get_publishable_shop", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetPublishTaskResult(query easycb.AnyMap) (*GetPublishTaskResultRsp, error) {
	var result GetPublishTaskResultRsp
	err := c.doRequest("GET", "/api/v2/global_product/get_publish_task_result", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetPublishedList(query easycb.AnyMap) (*GetPublishedListRsp, error) {
	var result GetPublishedListRsp
	err := c.doRequest("GET", "/api/v2/global_product/get_published_list", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GlobalUpdatePrice(body easycb.AnyMap) (*GlobalUpdatePriceRsp, error) {
	var result GlobalUpdatePriceRsp
	err := c.doRequest("POST", "/api/v2/global_product/update_price", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GlobalUpdateStock(body easycb.AnyMap) (*GlobalUpdateStockRsp, error) {
	var result GlobalUpdateStockRsp
	err := c.doRequest("POST", "/api/v2/global_product/update_stock", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GlobalSetSyncField(body easycb.AnyMap) (*GlobalSetSyncFieldRsp, error) {
	var result GlobalSetSyncFieldRsp
	err := c.doRequest("POST", "/api/v2/global_product/set_sync_field", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetGlobalItemId(query easycb.AnyMap) (*GetGlobalItemIdRsp, error) {
	var result GetGlobalItemIdRsp
	err := c.doRequest("GET", "/api/v2/global_product/get_global_item_id", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GlobalCategoryRecommend(query easycb.AnyMap) (*GlobalCategoryRecommendRsp, error) {
	var result GlobalCategoryRecommendRsp
	err := c.doRequest("GET", "/api/v2/global_product/category_recommend", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GlobalGetRecommendAttribute(query easycb.AnyMap) (*GlobalGetRecommendAttributeRsp, error) {
	var result GlobalGetRecommendAttributeRsp
	err := c.doRequest("GET", "/api/v2/global_product/get_recommend_attribute", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GlobalGetShopPublishableStatus(query easycb.AnyMap) (*GlobalGetShopPublishableStatusRsp, error) {
	var result GlobalGetShopPublishableStatusRsp
	err := c.doRequest("GET", "/api/v2/global_product/get_shop_publishable_status", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GlobalGetVariation(query easycb.AnyMap) (*GlobalGetVariationRsp, error) {
	var result GlobalGetVariationRsp
	err := c.doRequest("GET", "/api/v2/global_product/get_variations", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GlobalGetSizeChartDetail(query easycb.AnyMap) (*GlobalGetSizeChartDetailRsp, error) {
	var result GlobalGetSizeChartDetailRsp
	err := c.doRequest("GET", "/api/v2/global_product/get_size_chart_detail", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GlobalGetSizeChartList(query easycb.AnyMap) (*GlobalGetSizeChartListRsp, error) {
	var result GlobalGetSizeChartListRsp
	err := c.doRequest("GET", "/api/v2/global_product/get_size_chart_list", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
