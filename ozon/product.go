package ozon

import "github.com/easycb/easycb-go"

func (c *Client) ProductImport(body easycb.AnyMap) (*ProductImportRsp, error) {
	var result ProductImportRsp
	err := c.doRequest("POST", "/v3/product/import", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ProductImportInfo(body easycb.AnyMap) (*ProductImportInfoRsp, error) {
	var result ProductImportInfoRsp
	err := c.doRequest("POST", "/v1/product/import/info", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ProductImportBySku(body easycb.AnyMap) (*ProductImportBySkuRsp, error) {
	var result ProductImportBySkuRsp
	err := c.doRequest("POST", "/v1/product/import-by-sku", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ProductAttributesUpdate(body easycb.AnyMap) (*ProductAttributesUpdateRsp, error) {
	var result ProductAttributesUpdateRsp
	err := c.doRequest("POST", "/v1/product/attributes/update", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ProductPicturesImport(body easycb.AnyMap) (*ProductPicturesImportRsp, error) {
	var result ProductPicturesImportRsp
	err := c.doRequest("POST", "/v1/product/pictures/import", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ProductList(body easycb.AnyMap) (*ProductListRsp, error) {
	var result ProductListRsp
	err := c.doRequest("POST", "/v3/product/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ProductRatingBySku(body easycb.AnyMap) (*ProductRatingBySkuRsp, error) {
	var result ProductRatingBySkuRsp
	err := c.doRequest("POST", "/v1/product/rating-by-sku", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ProductInfoList(body easycb.AnyMap) (*ProductInfoListRsp, error) {
	var result ProductInfoListRsp
	err := c.doRequest("POST", "/v3/product/info/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ProductInfoAttributes(body easycb.AnyMap) (*ProductInfoAttributesRsp, error) {
	var result ProductInfoAttributesRsp
	err := c.doRequest("POST", "/v4/product/info/attributes", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ProductInfoDescription(body easycb.AnyMap) (*ProductInfoDescriptionRsp, error) {
	var result ProductInfoDescriptionRsp
	err := c.doRequest("POST", "/v1/product/info/description", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ProductInfoLimit(body easycb.AnyMap) (*ProductInfoLimitRsp, error) {
	var result ProductInfoLimitRsp
	err := c.doRequest("POST", "/v4/product/info/limit", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ProductUpdateOfferId(body easycb.AnyMap) (*ProductUpdateOfferIdRsp, error) {
	var result ProductUpdateOfferIdRsp
	err := c.doRequest("POST", "/v1/product/update/offer-id", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ProductArchive(body easycb.AnyMap) (*ProductArchiveRsp, error) {
	var result ProductArchiveRsp
	err := c.doRequest("POST", "/v1/product/archive", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ProductUnarchive(body easycb.AnyMap) (*ProductUnarchiveRsp, error) {
	var result ProductUnarchiveRsp
	err := c.doRequest("POST", "/v1/product/unarchive", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ProductDelete(body easycb.AnyMap) (*ProductDeleteRsp, error) {
	var result ProductDeleteRsp
	err := c.doRequest("POST", "/v2/products/delete", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ProductInfoSubscription(body easycb.AnyMap) (*ProductInfoSubscriptionRsp, error) {
	var result ProductInfoSubscriptionRsp
	err := c.doRequest("POST", "/v1/product/info/subscription", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ProductRelatedSkuGet(body easycb.AnyMap) (*ProductRelatedSkuGetRsp, error) {
	var result ProductRelatedSkuGetRsp
	err := c.doRequest("POST", "/v1/product/related-sku/get", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ProductPicturesInfo(body easycb.AnyMap) (*ProductPicturesInfoRsp, error) {
	var result ProductPicturesInfoRsp
	err := c.doRequest("POST", "/v2/product/pictures/info", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ProductInfoWrongVolume(body easycb.AnyMap) (*ProductInfoWrongVolumeRsp, error) {
	var result ProductInfoWrongVolumeRsp
	err := c.doRequest("POST", "/v1/product/info/wrong-volume", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ProductStocks(body easycb.AnyMap) (*ProductStocksRsp, error) {
	var result ProductStocksRsp
	err := c.doRequest("POST", "/v2/products/stocks", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ProductInfoStock(body easycb.AnyMap) (*ProductInfoStockRsp, error) {
	var result ProductInfoStockRsp
	err := c.doRequest("POST", "/v4/product/info/stocks", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ProductInfoStocksByWarehouseFbs(body easycb.AnyMap) (*ProductInfoStocksByWarehouseFbsRsp, error) {
	var result ProductInfoStocksByWarehouseFbsRsp
	err := c.doRequest("POST", "/v2/product/info/stocks-by-warehouse/fbs", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ProductInfoStocksByWarehouseFbsV1(body easycb.AnyMap) (*ProductInfoStocksByWarehouseFbsV1Rsp, error) {
	var result ProductInfoStocksByWarehouseFbsV1Rsp
	err := c.doRequest("POST", "/v1/product/info/stocks-by-warehouse/fbs", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ProductImportPrices(body easycb.AnyMap) (*ProductImportPricesRsp, error) {
	var result ProductImportPricesRsp
	err := c.doRequest("POST", "/v1/product/import/prices", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ProductActionTimerUpdate(body easycb.AnyMap) (*ProductActionTimerUpdateRsp, error) {
	var result ProductActionTimerUpdateRsp
	err := c.doRequest("POST", "/v1/product/action/timer/update", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ProductActionTimerStatus(body easycb.AnyMap) (*ProductActionTimerStatusRsp, error) {
	var result ProductActionTimerStatusRsp
	err := c.doRequest("POST", "/v1/product/action/timer/status", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ProductInfoPrices(body easycb.AnyMap) (*ProductInfoPricesRsp, error) {
	var result ProductInfoPricesRsp
	err := c.doRequest("POST", "/v5/product/info/prices", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ProductInfoDiscounted(body easycb.AnyMap) (*ProductInfoDiscountedRsp, error) {
	var result ProductInfoDiscountedRsp
	err := c.doRequest("POST", "/v1/product/info/discounted", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ProductUpdateDiscount(body easycb.AnyMap) (*ProductUpdateDiscountRsp, error) {
	var result ProductUpdateDiscountRsp
	err := c.doRequest("POST", "/v1/product/update/discount", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ProductInfoWarehouseStocks(body easycb.AnyMap) (*ProductInfoWarehouseStocksRsp, error) {
	var result ProductInfoWarehouseStocksRsp
	err := c.doRequest("POST", "/v1/product/info/warehouse/stocks", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ProductStairwayDiscountByQuantitySet(body easycb.AnyMap) (*ProductStairwayDiscountByQuantitySetRsp, error) {
	var result ProductStairwayDiscountByQuantitySetRsp
	err := c.doRequest("POST", "/v1/product/stairway-discount/by-quantity/set", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ProductStairwayDiscountByQuantityGet(body easycb.AnyMap) (*ProductStairwayDiscountByQuantityGetRsp, error) {
	var result ProductStairwayDiscountByQuantityGetRsp
	err := c.doRequest("POST", "/v1/product/stairway-discount/by-quantity/get", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ProductVisibilitySet(body easycb.AnyMap) (*ProductVisibilitySetRsp, error) {
	var result ProductVisibilitySetRsp
	err := c.doRequest("POST", "/v1/product/visibility/set", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ProductVisibilityInfo(body easycb.AnyMap) (*ProductVisibilityInfoRsp, error) {
	var result ProductVisibilityInfoRsp
	err := c.doRequest("POST", "/v1/product/visibility/info", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
