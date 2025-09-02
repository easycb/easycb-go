package tiktok

import (
	"fmt"
	"github.com/easycb/easycb-go"
)

func (c *Client) CheckListingPrerequisites() (*CheckListingPrerequisitesRsp, error) {
	var result CheckListingPrerequisitesRsp
	path := "/product/202312/prerequisites"
	err := c.doRequest("GET", path, nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetCategories(query easycb.AnyMap) (*GetCategoriesRsp, error) {
	var result GetCategoriesRsp
	path := "/product/202309/categories"
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) RecommendCategory(query easycb.AnyMap, body easycb.AnyMap) (*RecommendCategoryRsp, error) {
	var result RecommendCategoryRsp
	path := "/product/202309/categories/recommend"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetCategoryRules(categoryId int, query easycb.AnyMap) (*GetCategoryRulesRsp, error) {
	var result GetCategoryRulesRsp
	path := fmt.Sprintf("/product/202309/categories/%d/rules", categoryId)
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetAttributes(categoryId int, query easycb.AnyMap) (*GetAttributesRsp, error) {
	var result GetAttributesRsp
	path := fmt.Sprintf("/product/202309/categories/%d/attributes", categoryId)
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetBrands(query easycb.AnyMap) (*GetBrandsRsp, error) {
	var result GetBrandsRsp
	path := "/product/202309/brands"
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CreateCustomBrands(body easycb.AnyMap) (*CreateCustomBrandsRsp, error) {
	var result CreateCustomBrandsRsp
	path := "/product/202309/brands"
	err := c.doRequest("POST", path, nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CheckProductListing(query easycb.AnyMap, body easycb.AnyMap) (*CheckProductListingRsp, error) {
	var result CheckProductListingRsp
	path := "/product/202309/products/listing_check"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UploadProductImage(body easycb.AnyMap) (*UploadProductImageRsp, error) {
	var result UploadProductImageRsp
	path := "/product/202309/images/upload"
	err := c.doFileRequest("POST", path, nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UploadProductFile(body easycb.AnyMap) (*UploadProductFileRsp, error) {
	var result UploadProductFileRsp
	path := "/product/202309/files/upload"
	err := c.doFileRequest("POST", path, nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SearchSizeCharts(query easycb.AnyMap, body easycb.AnyMap) (*SearchSizeChartsRsp, error) {
	var result SearchSizeChartsRsp
	path := "/product/202407/sizecharts/search"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CreateProduct(query easycb.AnyMap, body easycb.AnyMap) (*CreateProductRsp, error) {
	var result CreateProductRsp
	path := "/product/202309/products"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PartialEditProduct(productId string, query easycb.AnyMap, body easycb.AnyMap) (*PartialEditProductRsp, error) {
	var result PartialEditProductRsp
	path := fmt.Sprintf("/product/202309/products/%s/partial_edit", productId)
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) EditProduct(productId string, query easycb.AnyMap, body easycb.AnyMap) (*EditProductRsp, error) {
	var result EditProductRsp
	path := fmt.Sprintf("/product/202309/products/%s", productId)
	err := c.doRequest("PUT", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ActivateProduct(query easycb.AnyMap, body easycb.AnyMap) (*ActivateProductRsp, error) {
	var result ActivateProductRsp
	path := "/product/202309/products/activate"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeactivateProducts(query easycb.AnyMap, body easycb.AnyMap) (*DeactivateProductsRsp, error) {
	var result DeactivateProductsRsp
	path := "/product/202309/products/deactivate"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteProducts(query easycb.AnyMap, body easycb.AnyMap) (*DeleteProductsRsp, error) {
	var result DeleteProductsRsp
	path := "/product/202309/products"
	err := c.doRequest("DELETE", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) RecoverProducts(query easycb.AnyMap, body easycb.AnyMap) (*RecoverProductsRsp, error) {
	var result RecoverProductsRsp
	path := "/product/202309/products/recover"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetProduct(itemId string, query easycb.AnyMap) (*GetProductRsp, error) {
	var result GetProductRsp
	path := fmt.Sprintf("/product/202309/products/%s", itemId)
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SearchProducts(query easycb.AnyMap, body easycb.AnyMap) (*SearchProductsRsp, error) {
	var result SearchProductsRsp
	path := "/product/202312/products/search"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdatePrice(productId string, query easycb.AnyMap, body easycb.AnyMap) (*UpdatePriceRsp, error) {
	var result UpdatePriceRsp
	path := fmt.Sprintf("/product/202309/products/%s/prices/update", productId)
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateInventory(productId string, query easycb.AnyMap, body easycb.AnyMap) (*UpdateInventoryRsp, error) {
	var result UpdateInventoryRsp
	path := fmt.Sprintf("/product/202309/products/%s/inventory/update", productId)
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) InventorySearch(query easycb.AnyMap, body easycb.AnyMap) (*InventorySearchRsp, error) {
	var result InventorySearchRsp
	path := "/product/202309/inventory/search"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ProductInformationIssueDiagnosis(query easycb.AnyMap) (*ProductInformationIssueDiagnosisRsp, error) {
	var result ProductInformationIssueDiagnosisRsp
	path := "/product/202405/products/diagnoses"
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetProductsSEOWords(query easycb.AnyMap) (*GetProductsSEOWordsRsp, error) {
	var result GetProductsSEOWordsRsp
	path := "/product/202405/products/seo_words"
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetRecommendedProductTitleAndDescription(query easycb.AnyMap) (*GetRecommendedProductTitleAndDescriptionRsp, error) {
	var result GetRecommendedProductTitleAndDescriptionRsp
	path := "/product/202405/products/suggestions"
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) OptimizedImages(query easycb.AnyMap, body easycb.AnyMap) (*OptimizedImagesRsp, error) {
	var result OptimizedImagesRsp
	path := "/product/202404/images/optimize"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetGlobalCategories(query easycb.AnyMap) (*GetGlobalCategoriesRsp, error) {
	var result GetGlobalCategoriesRsp
	path := "/product/202309/global_categories"
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) RecommendGlobalCategories(body easycb.AnyMap) (*RecommendGlobalCategoriesRsp, error) {
	var result RecommendGlobalCategoriesRsp
	path := "/product/202309/global_categories/recommend"
	err := c.doRequest("POST", path, nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetGlobalCategoryRules(categoryId int, query easycb.AnyMap) (*GetGlobalCategoryRulesRsp, error) {
	var result GetGlobalCategoryRulesRsp
	path := fmt.Sprintf("/product/202309/categories/%d/global_rules", categoryId)
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetGlobalAttributes(categoryId int, query easycb.AnyMap) (*GetGlobalAttributesRsp, error) {
	var result GetGlobalAttributesRsp
	path := fmt.Sprintf("/product/202309/categories/%d/global_attributes", categoryId)
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CreateGlobalProduct(body easycb.AnyMap) (*CreateGlobalProductRsp, error) {
	var result CreateGlobalProductRsp
	path := "/product/202309/global_products"
	err := c.doRequest("POST", path, nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PublishGlobalProduct(globalProductId string, body easycb.AnyMap) (*PublishGlobalProductRsp, error) {
	var result PublishGlobalProductRsp
	path := fmt.Sprintf("/product/202309/global_products/%s/publish", globalProductId)
	err := c.doRequest("POST", path, nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) EditGlobalProduct(globalProductId string, body easycb.AnyMap) (*EditGlobalProductRsp, error) {
	var result EditGlobalProductRsp
	path := fmt.Sprintf("/product/202309/global_products/%s", globalProductId)
	err := c.doRequest("PUT", path, nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteGlobalProducts(body easycb.AnyMap) (*DeleteGlobalProductsRsp, error) {
	var result DeleteGlobalProductsRsp
	path := "/product/202309/global_products"
	err := c.doRequest("DELETE", path, nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetGlobalProduct(globalProductId string) (*GetGlobalProductRsp, error) {
	var result GetGlobalProductRsp
	path := fmt.Sprintf("/product/202309/global_products/%s", globalProductId)
	err := c.doRequest("GET", path, nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SearchGlobalProducts(query easycb.AnyMap, body easycb.AnyMap) (*SearchGlobalProductsRsp, error) {
	var result SearchGlobalProductsRsp
	path := "/product/202312/global_products/search"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SearchGlobalProductsV202309(query easycb.AnyMap, body easycb.AnyMap) (*SearchGlobalProductsV202309Rsp, error) {
	var result SearchGlobalProductsV202309Rsp
	path := "/product/202309/global_products/search"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateGlobalInventory(globalProductId string, body easycb.AnyMap) (*UpdateGlobalInventoryRsp, error) {
	var result UpdateGlobalInventoryRsp
	path := fmt.Sprintf("/product/202309/global_products/%s/inventory/update", globalProductId)
	err := c.doRequest("POST", path, nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CreateManufacturer(body easycb.AnyMap) (*CreateManufacturerRsp, error) {
	var result CreateManufacturerRsp
	path := "/product/202409/compliance/manufacturers"
	err := c.doRequest("POST", path, nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PartialEditManufacturer(manufacturerId string, body easycb.AnyMap) (*PartialEditManufacturerRsp, error) {
	var result PartialEditManufacturerRsp
	path := fmt.Sprintf("/product/202409/compliance/manufacturers/%s/partial_edit", manufacturerId)
	err := c.doRequest("POST", path, nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SearchManufacturers(query easycb.AnyMap, body easycb.AnyMap) (*SearchManufacturersRsp, error) {
	var result SearchManufacturersRsp
	path := "/product/202409/compliance/manufacturers/search"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SearchManufacturersV202501(query easycb.AnyMap, body easycb.AnyMap) (*SearchManufacturersRspV202501, error) {
	var result SearchManufacturersRspV202501
	path := "/product/202501/compliance/manufacturers/search"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CreateResponsiblePerson(body easycb.AnyMap) (*CreateResponsiblePersonRsp, error) {
	var result CreateResponsiblePersonRsp
	path := "/product/202409/compliance/responsible_persons"
	err := c.doRequest("POST", path, nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PartialEditResponsiblePerson(responsiblePersonId string, body easycb.AnyMap) (*PartialEditResponsiblePersonRsp, error) {
	var result PartialEditResponsiblePersonRsp
	path := fmt.Sprintf("/product/202409/compliance/responsible_persons/%s/partial_edit", responsiblePersonId)
	err := c.doRequest("POST", path, nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SearchResponsiblePersons(query easycb.AnyMap, body easycb.AnyMap) (*SearchResponsiblePersonsRsp, error) {
	var result SearchResponsiblePersonsRsp
	path := "/product/202409/compliance/responsible_persons/search"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SearchResponsiblePersonsV202501(query easycb.AnyMap, body easycb.AnyMap) (*SearchResponsiblePersonsV202501Rsp, error) {
	var result SearchResponsiblePersonsV202501Rsp
	path := "/product/202501/compliance/responsible_persons/search"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CreateCategoryUpgradeTask(query easycb.AnyMap) (*CreateCategoryUpgradeTaskRsp, error) {
	var result CreateCategoryUpgradeTaskRsp
	path := "/product/202407/products/category_upgrade_task"
	err := c.doRequest("POST", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ListingSchemas(query easycb.AnyMap) (*ListingSchemasRsp, error) {
	var result ListingSchemasRsp
	path := "/product/202407/listing_schemas"
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DiagnoseAndOptimizeProduct(query easycb.AnyMap, body easycb.AnyMap) (*DiagnoseAndOptimizeProductRsp, error) {
	var result DiagnoseAndOptimizeProductRsp
	path := "/product/202411/products/diagnose_optimize"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetGlobalListingRules(query easycb.AnyMap) (*GetGlobalListingRulesRsp, error) {
	var result GetGlobalListingRulesRsp
	path := "/product/202507/global_listing_rules"
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ReplicateProduct(productId string, query easycb.AnyMap, body easycb.AnyMap) (*ReplicateProductRsp, error) {
	var result ReplicateProductRsp
	path := fmt.Sprintf("/product/202507/products/%s/global_replicate", productId)
	err := c.doRequest("GET", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetGlobalReplicatedProducts(productId string, query easycb.AnyMap) (*GetGlobalReplicatedProductsRsp, error) {
	var result GetGlobalReplicatedProductsRsp
	path := fmt.Sprintf("/product/202507/products/%s/replicated_products", productId)
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
