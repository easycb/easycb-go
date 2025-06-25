package shein

import "github.com/easycb/easycb-go"

func (c *Client) ProductPublishOrEdit(body easycb.AnyMap, headers map[string]string) (*ProductPublishOrEditRsp, error) {
	var result ProductPublishOrEditRsp
	err := c.doRequest("POST", "/open-api/goods/product/publishOrEdit", nil, body, headers, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ProductQueryDocumentState(body easycb.AnyMap, headers map[string]string) (*ProductQueryDocumentStateRsp, error) {
	var result ProductQueryDocumentStateRsp
	err := c.doRequest("POST", "/open-api/goods/query-document-state", nil, body, headers, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ProductQuery(body easycb.AnyMap, headers map[string]string) (*ProductQueryRsp, error) {
	var result ProductQueryRsp
	err := c.doRequest("POST", "/open-api/openapi-business-backend/product/query", nil, body, headers, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ProductSpuInfo(body easycb.AnyMap) (*ProductSpuInfoRsp, error) {
	var result ProductSpuInfoRsp
	err := c.doRequest("POST", "/open-api/goods/spu-info", nil, body, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) QueryCategoryTree(body easycb.AnyMap, headers map[string]string) (*QueryCategoryTreeRsp, error) {
	var result QueryCategoryTreeRsp
	err := c.doRequest("POST", "/open-api/goods/query-category-tree", nil, body, headers, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) QueryAttributeTemplate(body easycb.AnyMap, headers map[string]string) (*QueryAttributeTemplateRsp, error) {
	var result QueryAttributeTemplateRsp
	err := c.doRequest("POST", "/open-api/goods/query-attribute-template", nil, body, headers, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetCustomAttributePermissionConfig(body easycb.AnyMap, headers map[string]string) (*GetCustomAttributePermissionConfigRsp, error) {
	var result GetCustomAttributePermissionConfigRsp
	err := c.doRequest("POST", "/open-api/goods/get-custom-attribute-permission-config", nil, body, headers, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) AddCustomAttributeValue(body easycb.AnyMap, headers map[string]string) (*AddCustomAttributeValueRsp, error) {
	var result AddCustomAttributeValueRsp
	err := c.doRequest("POST", "/open-api/goods/add-custom-attribute-value", nil, body, headers, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) QueryPublishFillInStandard(body easycb.AnyMap, headers map[string]string) (*QueryPublishFillInStandardRsp, error) {
	var result QueryPublishFillInStandardRsp
	err := c.doRequest("POST", "/open-api/goods/query-publish-fill-in-standard", nil, body, headers, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ImageCategorySuggestion(body easycb.AnyMap, headers map[string]string) (*ImageCategorySuggestionRsp, error) {
	var result ImageCategorySuggestionRsp
	err := c.doRequest("POST", "/open-api/goods/image-category-suggestion", nil, body, headers, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) TransformPic(body easycb.AnyMap, headers map[string]string) (*TransformPicRsp, error) {
	var result TransformPicRsp
	err := c.doRequest("POST", "/open-api/goods/transform-pic", nil, body, headers, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UploadPic(body easycb.AnyMap, headers map[string]string) (*UploadPicRsp, error) {
	var result UploadPicRsp
	err := c.doFileRequest("POST", "/open-api/goods/upload-pic", nil, body, headers, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) QuerySiteList(body easycb.AnyMap) (*QuerySiteListRsp, error) {
	var result QuerySiteListRsp
	err := c.doRequest("POST", "/open-api/goods/query-site-list", nil, body, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) QueryBrandList(body easycb.AnyMap, headers map[string]string) (*QueryBrandListRsp, error) {
	var result QueryBrandListRsp
	err := c.doRequest("POST", "/open-api/goods/query-brand-list", nil, body, headers, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ProductPriceSave(body easycb.AnyMap) (*ProductPriceSaveRsp, error) {
	var result ProductPriceSaveRsp
	err := c.doRequest("POST", "/open-api/openapi-business-backend/product/price/save", nil, body, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ProductUpdateCost(body easycb.AnyMap) (*ProductUpdateCostRsp, error) {
	var result ProductUpdateCostRsp
	err := c.doRequest("POST", "/open-api/goods/update-cost", nil, body, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ModifySkcShelf(body easycb.AnyMap, headers map[string]string) (*ModifySkcShelfRsp, error) {
	var result ModifySkcShelfRsp
	err := c.doRequest("POST", "/open-api/goods/modify-skc-shelf", nil, body, headers, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) WholeBrands(query easycb.AnyMap) (*WholeBrandsRsp, error) {
	var result WholeBrandsRsp
	err := c.doRequest("GET", "/open-api/goods-brand/whole-brands", query, nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PrintBarcode(body easycb.AnyMap, headers map[string]string) (*PrintBarcodeRsp, error) {
	var result PrintBarcodeRsp
	err := c.doRequest("POST", "/open-api/goods/print-barcode", nil, body, headers, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetCertificateRule(body easycb.AnyMap, headers map[string]string) (*GetCertificateRuleRsp, error) {
	var result GetCertificateRuleRsp
	err := c.doRequest("POST", "/open-api/goods/get-certificate-rule", nil, body, headers, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetAllCertificateTypeListV2(body easycb.AnyMap, headers map[string]string) (*GetAllCertificateTypeListV2Rsp, error) {
	var result GetAllCertificateTypeListV2Rsp
	err := c.doRequest("POST", "/open-api/goods/certificate/get-all-certificate-type-list-v2", nil, body, headers, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UploadCertificateFile(body easycb.AnyMap, headers map[string]string) (*UploadCertificateFileRsp, error) {
	var result UploadCertificateFileRsp
	err := c.doFileRequest("POST", "/open-api/goods/upload-certificate-file", nil, body, headers, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SaveOrUpdateCertificatePool(body easycb.AnyMap, headers map[string]string) (*SaveOrUpdateCertificatePoolRsp, error) {
	var result SaveOrUpdateCertificatePoolRsp
	err := c.doRequest("POST", "/open-api/goods/save-or-update-certificate-pool", nil, body, headers, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SaveOrUpdateSupplierCertificate(body easycb.AnyMap, headers map[string]string) (*SaveOrUpdateSupplierCertificateRsp, error) {
	var result SaveOrUpdateSupplierCertificateRsp
	err := c.doRequest("POST", "/open-api/goods/save-or-update-supplier-certificate", nil, body, headers, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SaveCertificatePoolSkcBind(body easycb.AnyMap, headers map[string]string) (*SaveCertificatePoolSkcBindRsp, error) {
	var result SaveCertificatePoolSkcBindRsp
	err := c.doRequest("POST", "/open-api/goods/save-certificate-pool-skc-bind", nil, body, headers, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) BatchSkcSize(body easycb.AnyMap) (*BatchSkcSizeRsp, error) {
	var result BatchSkcSizeRsp
	err := c.doRequest("POST", "/open-api/goods/batch-skc-size", nil, body, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) NumberList(query easycb.AnyMap) (*NumberListRsp, error) {
	var result NumberListRsp
	err := c.doRequest("GET", "/open-api/goods/number-list", query, nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) RevokeProduct(body easycb.AnyMap, headers map[string]string) (*RevokeProductRsp, error) {
	var result RevokeProductRsp
	err := c.doRequest("POST", "/open-api/goods/revoke-product", nil, body, headers, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
