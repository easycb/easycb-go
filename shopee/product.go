package shopee

import (
	"github.com/easycb/easycb-go"
)

func (c *Client) GetCategory(query easycb.AnyMap) (*GetCategoryRsp, error) {
	var result GetCategoryRsp
	err := c.doRequest("GET", "/api/v2/product/get_category", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetAttributeTree(query easycb.AnyMap) (*GetAttributeTreeRsp, error) {
	var result GetAttributeTreeRsp
	err := c.doRequest("GET", "/api/v2/product/get_attribute_tree", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetBrandList(query easycb.AnyMap) (*GetBrandListRsp, error) {
	var result GetBrandListRsp
	err := c.doRequest("GET", "/api/v2/product/get_brand_list", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetItemLimit(query easycb.AnyMap) (*GetItemLimitRsp, error) {
	var result GetItemLimitRsp
	err := c.doRequest("GET", "/api/v2/product/get_item_limit", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetItemList(query easycb.AnyMap) (*GetItemListRsp, error) {
	var result GetItemListRsp
	err := c.doRequest("GET", "/api/v2/product/get_item_list", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetItemBaseInfo(query easycb.AnyMap) (*GetItemBaseInfoRsp, error) {
	var result GetItemBaseInfoRsp
	err := c.doRequest("GET", "/api/v2/product/get_item_base_info", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetItemExtraInfo(query easycb.AnyMap) (*GetItemExtraInfoRsp, error) {
	var result GetItemExtraInfoRsp
	err := c.doRequest("GET", "/api/v2/product/get_item_extra_info", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) AddItem(body easycb.AnyMap) (*AddItemRsp, error) {
	var result AddItemRsp
	err := c.doRequest("POST", "/api/v2/product/add_item", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateItem(body easycb.AnyMap) (*UpdateItemRsp, error) {
	var result UpdateItemRsp
	err := c.doRequest("POST", "/api/v2/product/update_item", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteItem(body easycb.AnyMap) (*DeleteItemRsp, error) {
	var result DeleteItemRsp
	err := c.doRequest("POST", "/api/v2/product/delete_item", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) InitTierVariation(body easycb.AnyMap) (*InitTierVariationRsp, error) {
	var result InitTierVariationRsp
	err := c.doRequest("POST", "/api/v2/product/init_tier_variation", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateTierVariation(body easycb.AnyMap) (*UpdateTierVariationRsp, error) {
	var result UpdateTierVariationRsp
	err := c.doRequest("POST", "/api/v2/product/update_tier_variation", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetModelList(query easycb.AnyMap) (*GetModelListRsp, error) {
	var result GetModelListRsp
	err := c.doRequest("GET", "/api/v2/product/get_model_list", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateModel(body easycb.AnyMap) (*UpdateModelRsp, error) {
	var result UpdateModelRsp
	err := c.doRequest("POST", "/api/v2/product/update_model", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteModel(body easycb.AnyMap) (*DeleteModelRsp, error) {
	var result DeleteModelRsp
	err := c.doRequest("POST", "/api/v2/product/delete_model", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SupportSizeChart(body easycb.AnyMap) (*SupportSizeChartRsp, error) {
	var result SupportSizeChartRsp
	err := c.doRequest("POST", "", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateSizeChart(body easycb.AnyMap) (*UpdateSizeChartRsp, error) {
	var result UpdateSizeChartRsp
	err := c.doRequest("POST", "", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UnlistItem(body easycb.AnyMap) (*UnlistItemRsp, error) {
	var result UnlistItemRsp
	err := c.doRequest("POST", "", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdatePrice(body easycb.AnyMap) (*UpdatePriceRsp, error) {
	var result UpdatePriceRsp
	err := c.doRequest("POST", "", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateStock(body easycb.AnyMap) (*UpdateStockRsp, error) {
	var result UpdateStockRsp
	err := c.doRequest("POST", "", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) BoostItem(body easycb.AnyMap) (*BoostItemRsp, error) {
	var result BoostItemRsp
	err := c.doRequest("POST", "", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetBoostedList(body easycb.AnyMap) (*GetBoostedListRsp, error) {
	var result GetBoostedListRsp
	err := c.doRequest("POST", "", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetItemPromotion(body easycb.AnyMap) (*GetItemPromotionRsp, error) {
	var result GetItemPromotionRsp
	err := c.doRequest("POST", "", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateSipItemPrice(body easycb.AnyMap) (*UpdateSipItemPriceRsp, error) {
	var result UpdateSipItemPriceRsp
	err := c.doRequest("POST", "", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SearchItem(body easycb.AnyMap) (*SearchItemRsp, error) {
	var result SearchItemRsp
	err := c.doRequest("POST", "", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetComment(body easycb.AnyMap) (*GetCommentRsp, error) {
	var result GetCommentRsp
	err := c.doRequest("POST", "", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ReplyComment(body easycb.AnyMap) (*ReplyCommentRsp, error) {
	var result ReplyCommentRsp
	err := c.doRequest("POST", "", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CategoryRecommend(body easycb.AnyMap) (*CategoryRecommendRsp, error) {
	var result CategoryRecommendRsp
	err := c.doRequest("POST", "", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) RegisterBrand(body easycb.AnyMap) (*RegisterBrandRsp, error) {
	var result RegisterBrandRsp
	err := c.doRequest("POST", "", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetRecommendAttribute(body easycb.AnyMap) (*GetRecommendAttributeRsp, error) {
	var result GetRecommendAttributeRsp
	err := c.doRequest("POST", "", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetProductInfo(body easycb.AnyMap) (*GetProductInfoRsp, error) {
	var result GetProductInfoRsp
	err := c.doRequest("POST", "", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetWeightRecommendation(body easycb.AnyMap) (*GetWeightRecommendationRsp, error) {
	var result GetWeightRecommendationRsp
	err := c.doRequest("POST", "", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetSizeChartList(body easycb.AnyMap) (*GetSizeChartListRsp, error) {
	var result GetSizeChartListRsp
	err := c.doRequest("POST", "", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetSizeChartDetail(body easycb.AnyMap) (*GetSizeChartDetailRsp, error) {
	var result GetSizeChartDetailRsp
	err := c.doRequest("POST", "", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetItemViolationInfo(body easycb.AnyMap) (*GetItemViolationInfoRsp, error) {
	var result GetItemViolationInfoRsp
	err := c.doRequest("POST", "", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetVariations(body easycb.AnyMap) (*GetVariationsRsp, error) {
	var result GetVariationsRsp
	err := c.doRequest("POST", "", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetAllVehicleList(body easycb.AnyMap) (*GetAllVehicleListRsp, error) {
	var result GetAllVehicleListRsp
	err := c.doRequest("POST", "", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetVehicleListByCompatibilityDetail(body easycb.AnyMap) (*GetVehicleListByCompatibilityDetailRsp, error) {
	var result GetVehicleListByCompatibilityDetailRsp
	err := c.doRequest("POST", "", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetItemContentDiagnosisResult(body easycb.AnyMap) (*GetItemContentDiagnosisResultRsp, error) {
	var result GetItemContentDiagnosisResultRsp
	err := c.doRequest("POST", "", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetItemListByContentDiagnosis(body easycb.AnyMap) (*GetItemListByContentDiagnosisRsp, error) {
	var result GetItemListByContentDiagnosisRsp
	err := c.doRequest("POST", "", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetKitItemLimit(body easycb.AnyMap) (*GetKitItemLimitRsp, error) {
	var result GetKitItemLimitRsp
	err := c.doRequest("POST", "", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) AddKitItem(body easycb.AnyMap) (*AddKitItemRsp, error) {
	var result AddKitItemRsp
	err := c.doRequest("POST", "", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateKitItem(body easycb.AnyMap) (*UpdateKitItemRsp, error) {
	var result UpdateKitItemRsp
	err := c.doRequest("POST", "", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetKitItemInfo(body easycb.AnyMap) (*GetKitItemInfoRsp, error) {
	var result GetKitItemInfoRsp
	err := c.doRequest("POST", "", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetSspList(body easycb.AnyMap) (*GetSspListRsp, error) {
	var result GetSspListRsp
	err := c.doRequest("POST", "", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetSspInfo(body easycb.AnyMap) (*GetSspInfoRsp, error) {
	var result GetSspInfoRsp
	err := c.doRequest("POST", "", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) AddSspItem(body easycb.AnyMap) (*AddSspItemRsp, error) {
	var result AddSspItemRsp
	err := c.doRequest("POST", "", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) LinkSsp(body easycb.AnyMap) (*LinkSspRsp, error) {
	var result LinkSspRsp
	err := c.doRequest("POST", "", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UnlinkSsp(body easycb.AnyMap) (*UnlinkSspRsp, error) {
	var result UnlinkSspRsp
	err := c.doRequest("POST", "", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetAitemByPitemId(body easycb.AnyMap) (*GetAitemByPitemIdRsp, error) {
	var result GetAitemByPitemIdRsp
	err := c.doRequest("POST", "", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
