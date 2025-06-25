package shein

type ProductPublishOrEditRsp struct {
	BaseRsp
	Info struct {
		Success bool   `json:"success"`
		SpuName string `json:"spu_name"`
		SkcList []struct {
			SkcName string `json:"skc_name"`
			SkuList []struct {
				SkuCode     string `json:"sku_code"`
				SupplierSku string `json:"supplier_sku"`
			} `json:"sku_list"`
		} `json:"skc_list"`
		Version        string      `json:"version"`
		PreValidResult interface{} `json:"pre_valid_result"`
		MccValidResult interface{} `json:"mcc_valid_result"`
		Extra          struct {
		} `json:"extra"`
	} `json:"info"`
}

type ProductQueryDocumentStateRsp struct {
	BaseRsp
	Info struct {
		Data []ProductQueryDocumentStateItem `json:"data"`
		Meta struct {
			Count     int         `json:"count"`
			CustomObj interface{} `json:"customObj"`
		} `json:"meta"`
	} `json:"info"`
}

type ProductQueryDocumentStateItem struct {
	SpuName string `json:"spuName"`
	Version string `json:"version"`
	SkcList []struct {
		SkcName       string `json:"skcName"`
		DocumentSn    string `json:"documentSn"`
		DocumentState int    `json:"documentState"`
		FailedReason  []struct {
			Language string `json:"language"`
			Content  string `json:"content"`
		} `json:"failedReason"`
	} `json:"skcList"`
}

type ProductQueryRsp struct {
	BaseRsp
	Info struct {
		Data []ProductQueryItem `json:"data"`
	} `json:"info"`
}

type ProductQueryItem struct {
	SkcName     string   `json:"skcName"`
	SkuCodeList []string `json:"skuCodeList"`
	SpuName     string   `json:"spuName"`
}

type ProductSpuInfoRsp struct {
	BaseRsp
	Info struct {
		ProductSpuInfoItem
	} `json:"info"`
}

type ProductSpuInfoItem struct {
	SpuName              string `json:"spuName"`
	CategoryId           int    `json:"categoryId"`
	ProductTypeId        int    `json:"productTypeId"`
	BrandCode            string `json:"brandCode"`
	SupplierCode         string `json:"supplierCode"`
	ProductMultiNameList []struct {
		ProductName string `json:"productName"`
		Language    string `json:"language"`
	} `json:"productMultiNameList"`
	ProductMultiDescList []struct {
		ProductDesc string `json:"productDesc"`
		Language    string `json:"language"`
	} `json:"productMultiDescList"`
	ProductAttributeInfoList []struct {
		AttributeId        int `json:"attributeId"`
		AttributeMultiList []struct {
			AttributeName string `json:"attributeName"`
			Language      string `json:"language"`
		} `json:"attributeMultiList"`
		AttributeValueId        int64 `json:"attributeValueId"`
		AttributeValueMultiList []struct {
			AttributeValueName string `json:"attributeValueName"`
			Language           string `json:"language"`
		} `json:"attributeValueMultiList"`
		AttributeValue interface{} `json:"attributeValue"`
	} `json:"productAttributeInfoList"`
	DimensionAttributeInfoList []interface{} `json:"dimensionAttributeInfoList"`
	SpuImageInfoList           interface{}   `json:"spuImageInfoList"`
	SkcInfoList                []struct {
		SkcName      string `json:"skcName"`
		SupplierCode string `json:"supplierCode"`
		SampleInfo   struct {
			SampleCode        string `json:"sampleCode"`
			ReserveSampleFlag int    `json:"reserveSampleFlag"`
			SpotFlag          int    `json:"spotFlag"`
			SampleJudgeType   int    `json:"sampleJudgeType"`
		} `json:"sampleInfo"`
		ProductMultiNameList []struct {
			ProductName string `json:"productName"`
			Language    string `json:"language"`
		} `json:"productMultiNameList"`
		AttributeId        int `json:"attributeId"`
		AttributeMultiList []struct {
			AttributeName string `json:"attributeName"`
			Language      string `json:"language"`
		} `json:"attributeMultiList"`
		AttributeValueId        int64 `json:"attributeValueId"`
		AttributeValueMultiList []struct {
			AttributeValueName string `json:"attributeValueName"`
			Language           string `json:"language"`
		} `json:"attributeValueMultiList"`
		SkuInfoList []struct {
			SkuCode           string `json:"skuCode"`
			SupplierSku       string `json:"supplierSku"`
			Length            string `json:"length"`
			Width             string `json:"width"`
			Height            string `json:"height"`
			Weight            int    `json:"weight"`
			MallState         int    `json:"mallState"`
			StopPurchase      int    `json:"stopPurchase"`
			SaleAttributeList []struct {
				AttributeId             int   `json:"attributeId"`
				AttributeValueId        int64 `json:"attributeValueId"`
				AttributeValueMultiList []struct {
					AttributeValueName string `json:"attributeValueName"`
					Language           string `json:"language"`
				} `json:"attributeValueMultiList"`
			} `json:"saleAttributeList"`
			PriceInfoList []struct {
				Site         string  `json:"site"`
				BasePrice    float64 `json:"basePrice"`
				SpecialPrice float64 `json:"specialPrice"`
				Currency     string  `json:"currency"`
			} `json:"priceInfoList"`
			CostInfoList     []interface{} `json:"costInfoList"`
			SkuImageInfoList interface{}   `json:"skuImageInfoList"`
		} `json:"skuInfoList"`
		ShelfStatusInfoList []struct {
			SiteAbbr       string `json:"siteAbbr"`
			ShelfStatus    int    `json:"shelfStatus"`
			LastShelfTime  string `json:"lastShelfTime"`
			FirstShelfTime string `json:"firstShelfTime"`
			LastUpdateTime string `json:"lastUpdateTime"`
		} `json:"shelfStatusInfoList"`
		SkcImageInfoList []struct {
			GroupCode      string `json:"groupCode"`
			ImageItemId    int64  `json:"imageItemId"`
			ImageType      string `json:"imageType"`
			ImageMediumUrl string `json:"imageMediumUrl"`
			ImageSmallUrl  string `json:"imageSmallUrl"`
			ImageUrl       string `json:"imageUrl"`
			Sort           int    `json:"sort"`
		} `json:"skcImageInfoList"`
		SiteDetailImageInfoList interface{}   `json:"siteDetailImageInfoList"`
		ProofOfStockInfoList    []interface{} `json:"proofOfStockInfoList"`
		SrpPriceInfo            interface{}   `json:"srpPriceInfo"`
	} `json:"skcInfoList"`
}

type QueryCategoryTreeRsp struct {
	BaseRsp
	Info struct {
		Data []CategoryTreeItem `json:"data"`
		Meta struct {
			Count     int         `json:"count"`
			CustomObj interface{} `json:"customObj"`
		} `json:"meta"`
	} `json:"info"`
}

type CategoryTreeItem struct {
	CategoryId       int                 `json:"category_id"`
	ProductTypeId    int                 `json:"product_type_id"`
	ParentCategoryId int                 `json:"parent_category_id"`
	CategoryName     string              `json:"category_name"`
	LastCategory     bool                `json:"last_category"`
	Children         []*CategoryTreeItem `json:"children,omitempty"`
}

type QueryAttributeTemplateRsp struct {
	BaseRsp
	Info struct {
		Data []QueryAttributeTemplateItem `json:"data"`
		Meta struct {
			Count     int         `json:"count"`
			CustomObj interface{} `json:"customObj"`
		} `json:"meta"`
	} `json:"info"`
}

type QueryAttributeTemplateItem struct {
	ProductTypeId  int `json:"product_type_id"`
	BusinessMode   int `json:"business_mode"`
	AttributeInfos []struct {
		AttributeId            int         `json:"attribute_id"`
		AttributeName          string      `json:"attribute_name"`
		AttributeNameEn        string      `json:"attribute_name_en"`
		AttributeIsShow        int         `json:"attribute_is_show"`
		AttributeSource        int         `json:"attribute_source"`
		AttributeLabel         int         `json:"attribute_label"`
		AttributeMode          int         `json:"attribute_mode"`
		DataDimension          int         `json:"data_dimension"`
		AttributeStatus        int         `json:"attribute_status"`
		AttributeType          int         `json:"attribute_type"`
		BusinessMode           int         `json:"business_mode"`
		IsSample               int         `json:"is_sample"`
		SupplierId             int         `json:"supplier_id"`
		AttributeDoc           interface{} `json:"attribute_doc"`
		AttributeDocImageList  interface{} `json:"attribute_doc_image_list"`
		AttributeValueInfoList []struct {
			AttributeValueId           int         `json:"attribute_value_id"`
			AttributeValue             string      `json:"attribute_value"`
			AttributeValueEn           string      `json:"attribute_value_en"`
			IsCustomAttributeValue     bool        `json:"is_custom_attribute_value"`
			IsShow                     int         `json:"is_show"`
			SupplierId                 int         `json:"supplier_id"`
			AttributeValueDoc          interface{} `json:"attribute_value_doc"`
			AttributeValueDocImageList interface{} `json:"attribute_value_doc_image_list"`
			AttributeValueGroupList    interface{} `json:"attribute_value_group_list"`
		} `json:"attribute_value_info_list"`
	} `json:"attribute_infos"`
	AttributeId []int `json:"attribute_id"`
}

type GetCustomAttributePermissionConfigRsp struct {
	BaseRsp
	Info struct {
		Data []struct {
			HasPermission  int `json:"has_permission"`
			LastCategoryId int `json:"last_category_id"`
			AttributeId    int `json:"attribute_id"`
		} `json:"data"`
		Meta struct {
			Count     int         `json:"count"`
			CustomObj interface{} `json:"customObj"`
		} `json:"meta"`
	} `json:"info"`
}

type AddCustomAttributeValueRsp struct {
	BaseRsp
	Info struct {
		SupplierId             int    `json:"supplier_id"`
		SupplierSource         int    `json:"supplier_source"`
		CategoryId             int    `json:"category_id"`
		AttributeId            int    `json:"attribute_id"`
		AttributeValueId       int    `json:"attribute_value_id"`
		AttributeValueName     string `json:"attribute_value_name"`
		AttributeValueMultiArr []struct {
			AttributeValueNameMulti string `json:"attribute_value_name_multi"`
			Language                string `json:"language"`
		} `json:"attribute_value_multi_arr"`
	} `json:"info"`
	Bbl interface{} `json:"bbl"`
}

type QueryPublishFillInStandardRsp struct {
	BaseRsp
	Info struct {
		FillInStandardList []struct {
			Module   string `json:"module"`
			FieldKey string `json:"field_key"`
			Required bool   `json:"required"`
			Show     bool   `json:"show"`
		} `json:"fill_in_standard_list"`
		DefaultLanguage          string        `json:"default_language"`
		PictureConfigList        []interface{} `json:"picture_config_list"`
		Currency                 string        `json:"currency"`
		SupportSaleAttributeSort interface{}   `json:"support_sale_attribute_sort"`
	} `json:"info"`
}

type ImageCategorySuggestionRsp struct {
	BaseRsp
	Data []struct {
		CategoryId int `json:"categoryId"`
		Order      int `json:"order"`
		Vote       int `json:"vote"`
	} `json:"data"`
	Meta struct {
		Count     int         `json:"count"`
		CustomObj interface{} `json:"customObj"`
	} `json:"meta"`
}

type TransformPicRsp struct {
	BaseRsp
	Info struct {
		Original      string `json:"original"`
		Transformed   string `json:"transformed"`
		FailureReason string `json:"failure_reason"`
	} `json:"info"`
}

type UploadPicRsp struct {
	BaseRsp
	Info struct {
		ImageUrl     string `json:"image_url"`
		Width        int    `json:"width"`
		Height       int    `json:"height"`
		Size         int    `json:"size"`
		ImageHexType string `json:"image_hex_type"`
	} `json:"info"`
}

type QuerySiteListRsp struct {
	BaseRsp
	Info struct {
		Data []QuerySiteListItem `json:"data"`
		Meta struct {
			Count     int         `json:"count"`
			CustomObj interface{} `json:"customObj"`
		} `json:"meta"`
	} `json:"info"`
}

type QuerySiteListItem struct {
	MainSite     string `json:"main_site"`
	MainSiteName string `json:"main_site_name"`
	SubSiteList  []struct {
		SiteName    string `json:"site_name"`
		SiteAbbr    string `json:"site_abbr"`
		SiteStatus  int    `json:"site_status"`
		StoreType   int    `json:"store_type"`
		Currency    string `json:"currency"`
		SymbolLeft  string `json:"symbol_left"`
		SymbolRight string `json:"symbol_right"`
	} `json:"sub_site_list"`
}

type QueryBrandListRsp struct {
	BaseRsp
	Info struct {
		Data []QueryBrandListItem `json:"data"`
		Meta struct {
			Count     int         `json:"count"`
			CustomObj interface{} `json:"customObj"`
		} `json:"meta"`
	} `json:"info"`
}

type QueryBrandListItem struct {
	BrandCode string `json:"brand_code"`
	BrandName string `json:"brand_name"`
}

type ProductPriceSaveRsp struct {
	BaseRsp
	Info struct {
		Data []struct {
			Success     bool   `json:"success"`
			Message     string `json:"message"`
			ProductCode string `json:"productCode"`
			Site        string `json:"site"`
			Status      int    `json:"status"`
		} `json:"data"`
	} `json:"info"`
}

type ProductUpdateCostRsp struct {
	BaseRsp
}

type ModifySkcShelfRsp struct {
	BaseRsp
	Info struct {
		SuccessCount   int           `json:"success_count"`
		FailureCount   int           `json:"failure_count"`
		TotalCount     int           `json:"total_count"`
		FailureResults []interface{} `json:"failure_results"`
	} `json:"info"`
}

type WholeBrandsRsp struct {
	BaseRsp
	Info struct {
		Data []WholeBrandsItem `json:"data"`
		Meta struct {
			Count     int         `json:"count"`
			CustomObj interface{} `json:"customObj"`
		} `json:"meta"`
	} `json:"info"`
}

type WholeBrandsItem struct {
	BrandCode string `json:"brand_code"`
	BrandName string `json:"brand_name"`
}

type PrintBarcodeRsp struct {
	BaseRsp
	Info struct {
		Url            string        `json:"url"`
		ErrorData      []interface{} `json:"errorData"`
		CodingInfoList []struct {
			OrderNo          string      `json:"orderNo"`
			SupplierSku      string      `json:"supplierSku"`
			SheinSku         string      `json:"sheinSku"`
			CustomCodingList []string    `json:"customCodingList"`
			Barcode          interface{} `json:"barcode"`
		} `json:"codingInfoList"`
	} `json:"info"`
}

type GetCertificateRuleRsp struct {
	BaseRsp
	Info struct {
		Data []GetCertificateRuleItem `json:"data"`
		Meta struct {
			Count     int         `json:"count"`
			CustomObj interface{} `json:"customObj"`
		} `json:"meta"`
	} `json:"info"`
}

type GetCertificateRuleItem struct {
	CertificateTypeId int `json:"certificateTypeId"`
	MergeSiteInfoList []struct {
		MergeSiteName string   `json:"mergeSiteName"`
		SubSiteList   []string `json:"subSiteList"`
	} `json:"mergeSiteInfoList"`
	CertificateTypeValue  string `json:"certificateTypeValue"`
	FileModelUrl          string `json:"fileModelUrl"`
	CertificateDimension  int    `json:"certificateDimension"`
	CertificateLabel      int    `json:"certificateLabel"`
	CertificateMissStatus bool   `json:"certificateMissStatus"`
	IsRequired            bool   `json:"isRequired"`
	CertificatePoolList   []struct {
		CertificatePoolId       int    `json:"certificatePoolId"`
		CertificateExpireStatus int    `json:"certificateExpireStatus"`
		ExpireTime              string `json:"expireTime"`
		AuditStatus             int    `json:"auditStatus"`
		PqmsCertificateSn       string `json:"pqmsCertificateSn"`
		CertificatePoolFileList []struct {
			CertificateUrlName string `json:"certificateUrlName"`
			CertificateUrl     string `json:"certificateUrl"`
		} `json:"certificatePoolFileList"`
	} `json:"certificatePoolList"`
	OtherSourceCertInfoList []interface{} `json:"otherSourceCertInfoList"`
	SelfCertificateList     []interface{} `json:"selfCertificateList"`
}

type GetAllCertificateTypeListV2Rsp struct {
	BaseRsp
	Info struct {
		Data []GetAllCertificateTypeListV2Item `json:"data"`
		Meta struct {
			Count     int         `json:"count"`
			CustomObj interface{} `json:"customObj"`
		} `json:"meta"`
	} `json:"info"`
}

type GetAllCertificateTypeListV2Item struct {
	CertificateTypeId      int         `json:"certificateTypeId"`
	CertificateLabel       int         `json:"certificateLabel"`
	CertificateType        string      `json:"certificateType"`
	CertificateDimension   int         `json:"certificateDimension"`
	FileModelUrl           string      `json:"fileModelUrl"`
	IsEnabled              int         `json:"isEnabled"`
	SrmDetectionAgencyList interface{} `json:"srmDetectionAgencyList"`
	PresetInfoList         []struct {
		PresetId        int    `json:"presetId"`
		PresetName      string `json:"presetName"`
		PresetRemark    string `json:"presetRemark"`
		InputType       int    `json:"inputType"`
		Unit            string `json:"unit"`
		IsRequired      int    `json:"isRequired"`
		PresetValueList []struct {
			PresetValueId int    `json:"presetValueId"`
			PresetValue   string `json:"presetValue"`
		} `json:"presetValueList"`
	} `json:"presetInfoList"`
	OtherPresetInfoList []struct {
		SourceFrom      string `json:"sourceFrom"`
		PresetId        int    `json:"presetId"`
		PresetName      string `json:"presetName"`
		PresetRemark    string `json:"presetRemark"`
		InputType       int    `json:"inputType"`
		Unit            string `json:"unit"`
		IsRequired      int    `json:"isRequired"`
		PresetValueList []struct {
			PresetValueId int    `json:"presetValueId"`
			PresetValue   string `json:"presetValue"`
		} `json:"presetValueList"`
	} `json:"otherPresetInfoList"`
}

type UploadCertificateFileRsp struct {
	BaseRsp
	Info struct {
		CertificateUrl string `json:"certificateUrl"`
	} `json:"info"`
}

type SaveOrUpdateCertificatePoolRsp struct {
	BaseRsp
	Info struct {
		CertificatePoolId int `json:"certificatePoolId"`
	} `json:"info"`
}

type SaveOrUpdateSupplierCertificateRsp struct {
	BaseRsp
}

type SaveCertificatePoolSkcBindRsp struct {
	BaseRsp
}

type BatchSkcSizeRsp struct {
	BaseRsp
	Info map[string]struct {
		Skc     string `json:"skc"`
		Size    string `json:"size"`
		SkuCode string `json:"sku_code"`
	} `json:"info"`
}

type NumberListRsp struct {
	BaseRsp
	Info struct {
		Page    int              `json:"page"`
		PerPage int              `json:"per_page"`
		Count   int              `json:"count"`
		List    []NumberListItem `json:"list"`
	} `json:"info"`
}

type NumberListItem struct {
	Skc         string `json:"skc"`
	SupplierSku string `json:"supplier_sku"`
	DesignCode  string `json:"design_code"`
	SkuCode     string `json:"sku_code"`
	Attribute   string `json:"attribute"`
}

type RevokeProductRsp struct {
	BaseRsp
	Info struct {
		SuccessList []struct {
			DocumentSn string `json:"documentSn"`
			SkcName    string `json:"skcName"`
		} `json:"successList"`
		FailList     []interface{} `json:"failList"`
		Total        int           `json:"total"`
		SuccessCount int           `json:"successCount"`
		FailCount    int           `json:"failCount"`
	} `json:"info"`
}
