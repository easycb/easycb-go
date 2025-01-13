package lazada

import "github.com/easycb/easycb-go"

func (c *Client) AdjustSellableQuantity(body easycb.AnyMap) (*AdjustSellableQuantityRsp, error) {
	var result AdjustSellableQuantityRsp
	err := c.doRequest("POST", "/product/stock/sellable/adjust", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) BatchUpdateSizeChart(body easycb.AnyMap) (*BatchUpdateSizeChartRsp, error) {
	var result BatchUpdateSizeChartRsp
	err := c.doRequest("POST", "/size/chart/batch/update", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CreateProduct(body easycb.AnyMap) (*CreateProductRsp, error) {
	var result CreateProductRsp
	err := c.doRequest("POST", "/product/create", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeactivateProduct(body easycb.AnyMap) (*DeactivateProductRsp, error) {
	var result DeactivateProductRsp
	err := c.doRequest("POST", "/product/deactivate", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ExitExperiment(body easycb.AnyMap) (*ExitExperimentRsp, error) {
	var result ExitExperimentRsp
	err := c.doRequest("POST", "/products/experiment/exit", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetBrandByPages(query easycb.AnyMap) (*GetBrandByPagesRsp, error) {
	var result GetBrandByPagesRsp
	err := c.doRequest("GET", "/category/brands/query", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetCategoryAttributes(query easycb.AnyMap) (*GetCategoryAttributesRsp, error) {
	var result GetCategoryAttributesRsp
	err := c.doRequest("GET", "/category/attributes/get", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetCategorySuggestion(query easycb.AnyMap) (*GetCategorySuggestionRsp, error) {
	var result GetCategorySuggestionRsp
	err := c.doRequest("GET", "/product/category/suggestion/get", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetCategorySuggestionInBulk(query easycb.AnyMap) (*GetCategorySuggestionInBulkRsp, error) {
	var result GetCategorySuggestionInBulkRsp
	err := c.doRequest("GET", "/product/category/suggestion/get/batch", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetCategoryTree(query easycb.AnyMap) (*GetCategoryTreeRsp, error) {
	var result GetCategoryTreeRsp
	err := c.doRequest("GET", "/category/tree/get", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetImageExperimentData(query easycb.AnyMap) (*GetImageExperimentDataRsp, error) {
	var result GetImageExperimentDataRsp
	err := c.doRequest("GET", "/products/experiment/getdata", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetNextCascadeProp(query easycb.AnyMap) (*GetNextCascadePropRsp, error) {
	var result GetNextCascadePropRsp
	err := c.doRequest("GET", "/category/cascade/getNextCascadeProp", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetPreQcRules(query easycb.AnyMap) (*GetPreQcRulesRsp, error) {
	var result GetPreQcRulesRsp
	err := c.doRequest("GET", "/product/seller/item/getPreQcRules", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetProductContentScore(query easycb.AnyMap) (*GetProductContentScoreRsp, error) {
	var result GetProductContentScoreRsp
	err := c.doRequest("GET", "/product/content/score/get", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetProductItem(query easycb.AnyMap) (*GetProductItemRsp, error) {
	var result GetProductItemRsp
	err := c.doRequest("GET", "/product/item/get", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetProducts(query easycb.AnyMap) (*GetProductsRsp, error) {
	var result GetProductsRsp
	err := c.doRequest("GET", "/products/get", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetQCAlertProducts(query easycb.AnyMap) (*GetQCAlertProductsRsp, error) {
	var result GetQCAlertProductsRsp
	err := c.doRequest("GET", "/product/qc/alert/list", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetQcStatus(query easycb.AnyMap) (*GetQcStatusRsp, error) {
	var result GetQcStatusRsp
	err := c.doRequest("GET", "/product/qc/status/get", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetResponse(query easycb.AnyMap) (*GetResponseRsp, error) {
	var result GetResponseRsp
	err := c.doRequest("GET", "/image/response/get", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetSellerItemLimit(query easycb.AnyMap) (*GetSellerItemLimitRsp, error) {
	var result GetSellerItemLimitRsp
	err := c.doRequest("GET", "/product/seller/item/limit", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetSizeChartTemplate(query easycb.AnyMap) (*GetSizeChartTemplateRsp, error) {
	var result GetSizeChartTemplateRsp
	err := c.doRequest("GET", "/size/chart/template/get", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetUnfilledAttributeItem(query easycb.AnyMap) (*GetUnfilledAttributeItemRsp, error) {
	var result GetUnfilledAttributeItemRsp
	err := c.doRequest("GET", "/product/unfilled/attribute/get", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) JoinExperiment(body easycb.AnyMap) (*JoinExperimentRsp, error) {
	var result JoinExperimentRsp
	err := c.doRequest("POST", "/products/experiment/join", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) MigrateImage(body easycb.AnyMap) (*MigrateImageRsp, error) {
	var result MigrateImageRsp
	err := c.doRequest("POST", "/image/migrate", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) MigrateImages(body easycb.AnyMap) (*MigrateImagesRsp, error) {
	var result MigrateImagesRsp
	err := c.doRequest("POST", "/images/migrate", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ProductCheck(body easycb.AnyMap) (*ProductCheckRsp, error) {
	var result ProductCheckRsp
	err := c.doRequest("POST", "/product/pre/check", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) QueryProductExperimentConfiguration(query easycb.AnyMap) (*QueryProductExperimentConfigurationRsp, error) {
	var result QueryProductExperimentConfigurationRsp
	err := c.doRequest("GET", "/products/experiment/getconfig", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) RemoveProduct(body easycb.AnyMap) (*RemoveProductRsp, error) {
	var result RemoveProductRsp
	err := c.doRequest("POST", "/product/remove", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) RemoveSku(body easycb.AnyMap) (*RemoveSkuRsp, error) {
	var result RemoveSkuRsp
	err := c.doRequest("POST", "/product/sku/remove", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SetImages(body easycb.AnyMap) (*SetImagesRsp, error) {
	var result SetImagesRsp
	err := c.doRequest("POST", "/images/set", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdatePriceQuantity(body easycb.AnyMap) (*UpdatePriceQuantityRsp, error) {
	var result UpdatePriceQuantityRsp
	err := c.doRequest("POST", "/product/price_quantity/update", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateProduct(body easycb.AnyMap) (*UpdateProductRsp, error) {
	var result UpdateProductRsp
	err := c.doRequest("POST", "/product/update", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateSellableQuantity(body easycb.AnyMap) (*UpdateSellableQuantityRsp, error) {
	var result UpdateSellableQuantityRsp
	err := c.doRequest("POST", "/product/stock/sellable/update", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UploadImage(body easycb.AnyMap) (*UploadImageRsp, error) {
	var result UploadImageRsp
	err := c.doFileRequest("POST", "/image/upload", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
