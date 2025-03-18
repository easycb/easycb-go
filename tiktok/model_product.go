package tiktok

type CheckListingPrerequisitesRsp struct {
	BaseRsp
	Data struct {
		CheckResults []struct {
			CheckItem   string   `json:"check_item"`
			FailReasons []string `json:"fail_reasons"`
			IsFailed    bool     `json:"is_failed"`
		} `json:"check_results"`
	} `json:"data"`
}

type GetCategoriesRsp struct {
	BaseRsp
	Data struct {
		Categories []struct {
			Id                 string   `json:"id"`
			IsLeaf             bool     `json:"is_leaf"`
			LocalName          string   `json:"local_name"`
			ParentId           string   `json:"parent_id"`
			PermissionStatuses []string `json:"permission_statuses"`
		} `json:"categories"`
	} `json:"data"`
}

type RecommendCategoryRsp struct {
	BaseRsp
	Data struct {
		Categories []struct {
			Id                 string   `json:"id"`
			IsLeaf             bool     `json:"is_leaf"`
			Level              int      `json:"level"`
			Name               string   `json:"name"`
			PermissionStatuses []string `json:"permission_statuses"`
		} `json:"categories"`
		LeafCategoryId string `json:"leaf_category_id"`
	} `json:"data"`
}

type GetCategoryRulesRsp struct {
	Data struct {
		Cod struct {
			IsSupported bool `json:"is_supported"`
		} `json:"cod"`
		Epr struct {
			IsRequired bool `json:"is_required"`
		} `json:"epr"`
		PackageDimension struct {
			IsRequired bool `json:"is_required"`
		} `json:"package_dimension"`
		ProductCertifications []struct {
			Id                    string `json:"id"`
			IsRequired            bool   `json:"is_required"`
			Name                  string `json:"name"`
			RequirementConditions []struct {
				AttributeId      string `json:"attribute_id"`
				AttributeValueId string `json:"attribute_value_id"`
				ConditionType    string `json:"condition_type"`
			} `json:"requirement_conditions"`
			SampleImageUrl string `json:"sample_image_url"`
		} `json:"product_certifications"`
		SizeChart struct {
			IsRequired  bool `json:"is_required"`
			IsSupported bool `json:"is_supported"`
		} `json:"size_chart"`
	} `json:"data"`
}

type GetAttributesRsp struct {
	BaseRsp
	Data struct {
		Attributes []struct {
			Id                    string `json:"id"`
			IsCustomizable        bool   `json:"is_customizable"`
			IsMultipleSelection   bool   `json:"is_multiple_selection"`
			IsRequried            bool   `json:"is_requried"`
			Name                  string `json:"name"`
			RequirementConditions []struct {
				AttributeId      string `json:"attribute_id"`
				AttributeValueId string `json:"attribute_value_id"`
				ConditionType    string `json:"condition_type"`
			} `json:"requirement_conditions"`
			Type            string `json:"type"`
			ValueDataFormat string `json:"value_data_format"`
			Values          []struct {
				Id   string `json:"id"`
				Name string `json:"name"`
			} `json:"values"`
		} `json:"attributes"`
	} `json:"data"`
}

type GetBrandsRsp struct {
	BaseRsp
	Data struct {
		Brands []struct {
			AuthorizedStatus string `json:"authorized_status"`
			BrandStatus      string `json:"brand_status"`
			Id               string `json:"id"`
			IsT1Brand        bool   `json:"is_t1_brand"`
			Name             string `json:"name"`
		} `json:"brands"`
		NextPageToken string `json:"next_page_token"`
		TotalCount    int    `json:"total_count"`
	} `json:"data"`
}

type CreateCustomBrandsRsp struct {
	BaseRsp
}

type CheckProductListingRsp struct {
	BaseRsp
	Data struct {
		CheckResult string `json:"check_result"`
		Diagnoses   []struct {
			DiagnosisResults []struct {
				Code        string `json:"code"`
				HowToSolve  string `json:"how_to_solve"`
				QualityTier string `json:"quality_tier"`
			} `json:"diagnosis_results"`
			Field       string `json:"field"`
			Suggestions struct {
				Images []struct {
					Height       int    `json:"height"`
					OptimizedUri string `json:"optimized_uri"`
					OptimizedUrl string `json:"optimized_url"`
					Uri          string `json:"uri"`
					Url          string `json:"url"`
					Width        int    `json:"width"`
				} `json:"images"`
				SeoWords []struct {
					Text string `json:"text"`
				} `json:"seo_words"`
				SmartTexts []struct {
					Text string `json:"text"`
				} `json:"smart_texts"`
			} `json:"suggestions"`
		} `json:"diagnoses"`
		FailReasons []struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"fail_reasons"`
		ListingQuality struct {
			CurrentTier              string `json:"current_tier"`
			RemainingRecommendations int    `json:"remaining_recommendations"`
		} `json:"listing_quality"`
		Warnings struct {
			Message string `json:"message"`
		} `json:"warnings"`
	} `json:"data"`
}

type UploadProductImageRsp struct {
	BaseRsp
	Data struct {
		Height  int    `json:"height"`
		Uri     string `json:"uri"`
		Url     string `json:"url"`
		UseCase string `json:"use_case"`
		Width   int    `json:"width"`
	} `json:"data"`
}

type UploadProductFileRsp struct {
	BaseRsp
	Data struct {
		Format string `json:"format"`
		Id     string `json:"id"`
		Name   string `json:"name"`
		Url    string `json:"url"`
	} `json:"data"`
}

type SearchSizeChartsRsp struct {
	BaseRsp
	Data struct {
		NextPageToken string `json:"next_page_token"`
		SizeChart     []struct {
			Images []struct {
				Locale string `json:"locale"`
				Uri    string `json:"uri"`
				Url    string `json:"url"`
			} `json:"images"`
			TemplateId   string `json:"template_id"`
			TemplateName string `json:"template_name"`
		} `json:"size_chart"`
		TotalCount int `json:"total_count"`
	} `json:"data"`
}

type CreateProductRsp struct {
	BaseRsp
	Data struct {
		ProductId string `json:"product_id"`
		Skus      []struct {
			ExternalSkuId   string `json:"external_sku_id"`
			Id              string `json:"id"`
			SalesAttributes []struct {
				Id      string `json:"id"`
				ValueId string `json:"value_id"`
			} `json:"sales_attributes"`
			SellerSku string `json:"seller_sku"`
		} `json:"skus"`
		Warnings []struct {
			Message string `json:"message"`
		} `json:"warnings"`
	} `json:"data"`
}

type PartialEditProductRsp struct {
	BaseRsp
	Data struct {
		ProductId string `json:"product_id"`
		Skus      []struct {
			ExternalSkuId   string `json:"external_sku_id"`
			Id              string `json:"id"`
			SalesAttributes []struct {
				Id      string `json:"id"`
				ValueId string `json:"value_id"`
			} `json:"sales_attributes"`
			SellerSku string `json:"seller_sku"`
		} `json:"skus"`
	} `json:"data"`
}

type EditProductRsp struct {
	BaseRsp
	Data struct {
		ProductId string `json:"product_id"`
		Skus      []struct {
			ExternalSkuId   string `json:"external_sku_id"`
			Id              string `json:"id"`
			SalesAttributes []struct {
				Id      string `json:"id"`
				ValueId string `json:"value_id"`
			} `json:"sales_attributes"`
			SellerSku string `json:"seller_sku"`
		} `json:"skus"`
		Warnings []struct {
			Message string `json:"message"`
		} `json:"warnings"`
	} `json:"data"`
}

type ActivateProductRsp struct {
	BaseRsp
	Data struct {
		Errors []struct {
			Code   int `json:"code"`
			Detail struct {
				ExtraErrors []struct {
					Code    int    `json:"code"`
					Message string `json:"message"`
				} `json:"extra_errors"`
				ProductId string `json:"product_id"`
			} `json:"detail"`
			Message string `json:"message"`
		} `json:"errors"`
	} `json:"data"`
}

type DeactivateProductsRsp struct {
	BaseRsp
	Data struct {
		Errors []struct {
			Code   int `json:"code"`
			Detail struct {
				ProductId string `json:"product_id"`
			} `json:"detail"`
			Message string `json:"message"`
		} `json:"errors"`
	} `json:"data"`
}

type DeleteProductsRsp struct {
	BaseRsp
	Data struct {
		Errors []struct {
			Code   int `json:"code"`
			Detail struct {
				ProductId string `json:"product_id"`
			} `json:"detail"`
			Message string `json:"message"`
		} `json:"errors"`
	} `json:"data"`
}

type RecoverProductsRsp struct {
	BaseRsp
	Data struct {
		Errors []struct {
			Code   int `json:"code"`
			Detail struct {
				ProductId string `json:"product_id"`
			} `json:"detail"`
			Message string `json:"message"`
		} `json:"errors"`
	} `json:"data"`
}

type GetProductRsp struct {
	BaseRsp
	Data struct {
		AuditFailedReasons []struct {
			Position    string   `json:"position"`
			Reasons     []string `json:"reasons"`
			Suggestions []string `json:"suggestions"`
		} `json:"audit_failed_reasons"`
		Brand struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"brand"`
		CategoryChains []struct {
			Id        string `json:"id"`
			IsLeaf    bool   `json:"is_leaf"`
			LocalName string `json:"local_name"`
			ParentId  string `json:"parent_id"`
		} `json:"category_chains"`
		Certifications []struct {
			Files []struct {
				Format string   `json:"format"`
				Id     string   `json:"id"`
				Name   string   `json:"name"`
				Urls   []string `json:"urls"`
			} `json:"files"`
			Id     string `json:"id"`
			Images []struct {
				Height    int      `json:"height"`
				ThumbUrls []string `json:"thumb_urls"`
				Uri       string   `json:"uri"`
				Urls      []string `json:"urls"`
				Width     int      `json:"width"`
			} `json:"images"`
			Title string `json:"title"`
		} `json:"certifications"`
		CreateTime      int `json:"create_time"`
		DeliveryOptions []struct {
			Id          string `json:"id"`
			IsAvailable bool   `json:"is_available"`
			Name        string `json:"name"`
		} `json:"delivery_options"`
		Description       string `json:"description"`
		ExternalProductId string `json:"external_product_id"`
		Id                string `json:"id"`
		IsCodAllowed      bool   `json:"is_cod_allowed"`
		MainImages        []struct {
			Height    int      `json:"height"`
			ThumbUrls []string `json:"thumb_urls"`
			Uri       string   `json:"uri"`
			Urls      []string `json:"urls"`
			Width     int      `json:"width"`
		} `json:"main_images"`
		PackageDimensions struct {
			Height string `json:"height"`
			Length string `json:"length"`
			Unit   string `json:"unit"`
			Width  string `json:"width"`
		} `json:"package_dimensions"`
		PackageWeight struct {
			Unit  string `json:"unit"`
			Value string `json:"value"`
		} `json:"package_weight"`
		ProductAttributes []struct {
			Id     string `json:"id"`
			Name   string `json:"name"`
			Values []struct {
				Id   string `json:"id"`
				Name string `json:"name"`
			} `json:"values"`
		} `json:"product_attributes"`
		SizeChart struct {
			Image struct {
				Height    int      `json:"height"`
				ThumbUrls []string `json:"thumb_urls"`
				Uri       string   `json:"uri"`
				Urls      []string `json:"urls"`
				Width     int      `json:"width"`
			} `json:"image"`
			Template struct {
				Id string `json:"id"`
			} `json:"template"`
		} `json:"size_chart"`
		Skus []struct {
			ExternalSkuId  string `json:"external_sku_id"`
			Id             string `json:"id"`
			IdentifierCode struct {
				Code string `json:"code"`
				Type string `json:"type"`
			} `json:"identifier_code"`
			Inventory []struct {
				Quantity    int    `json:"quantity"`
				WarehouseId string `json:"warehouse_id"`
			} `json:"inventory"`
			Price struct {
				Currency          string `json:"currency"`
				SalePrice         string `json:"sale_price"`
				TaxExclusivePrice string `json:"tax_exclusive_price"`
			} `json:"price"`
			SalesAttributes []struct {
				Id     string `json:"id"`
				Name   string `json:"name"`
				SkuImg struct {
					Height    int      `json:"height"`
					ThumbUrls []string `json:"thumb_urls"`
					Uri       string   `json:"uri"`
					Urls      []string `json:"urls"`
					Width     int      `json:"width"`
				} `json:"sku_img"`
				ValueId   string `json:"value_id"`
				ValueName string `json:"value_name"`
			} `json:"sales_attributes"`
			SellerSku string `json:"seller_sku"`
		} `json:"skus"`
		Status     string `json:"status"`
		Title      string `json:"title"`
		UpdateTime int    `json:"update_time"`
		Video      struct {
			CoverUrl string `json:"cover_url"`
			Format   string `json:"format"`
			Height   int    `json:"height"`
			Id       string `json:"id"`
			Size     int    `json:"size"`
			Url      string `json:"url"`
			Width    int    `json:"width"`
		} `json:"video"`
	} `json:"data"`
}

type SearchProductsRsp struct {
	BaseRsp
	Data struct {
		NextPageToken string `json:"next_page_token"`
		Products      []struct {
			Audit struct {
				Status string `json:"status"`
			} `json:"audit"`
			CreateTime                 int64  `json:"create_time"`
			Id                         string `json:"id"`
			IntegratedPlatformStatuses []struct {
				Platform string `json:"platform"`
				Status   string `json:"status"`
			} `json:"integrated_platform_statuses"`
			IsNotForSale           bool     `json:"is_not_for_sale"`
			ListingQualityTier     string   `json:"listing_quality_tier"`
			ProductSyncFailReasons []string `json:"product_sync_fail_reasons"`
			RecommendedCategories  []struct {
				Id        string `json:"id"`
				LocalName string `json:"local_name"`
			} `json:"recommended_categories"`
			SalesRegions []string `json:"sales_regions"`
			Skus         []struct {
				Id        string `json:"id"`
				Inventory []struct {
					Quantity    int    `json:"quantity"`
					WarehouseId string `json:"warehouse_id"`
				} `json:"inventory"`
				Price struct {
					Currency          string `json:"currency"`
					SalePrice         string `json:"sale_price"`
					TaxExclusivePrice string `json:"tax_exclusive_price"`
				} `json:"price"`
				SellerSku string `json:"seller_sku"`
			} `json:"skus"`
			Status     string `json:"status"`
			Title      string `json:"title"`
			UpdateTime int64  `json:"update_time"`
		} `json:"products"`
		TotalCount int `json:"total_count"`
	} `json:"data"`
}

type UpdatePriceRsp struct {
	BaseRsp
}

type UpdateInventoryRsp struct {
	BaseRsp
	Data struct {
		Errors []struct {
			Code   int `json:"code"`
			Detail struct {
				ExtraErrors []struct {
					Code        int    `json:"code"`
					Message     string `json:"message"`
					WarehouseId string `json:"warehouse_id"`
				} `json:"extra_errors"`
				SkuId string `json:"sku_id"`
			} `json:"detail"`
			Message string `json:"message"`
		} `json:"errors"`
	} `json:"data"`
}

type InventorySearchRsp struct {
	BaseRsp
	Data struct {
		Inventory []struct {
			ProductId string `json:"product_id"`
			Skus      []struct {
				Id                                  string `json:"id"`
				SellerSku                           string `json:"seller_sku"`
				TotalAvailableInventoryDistribution struct {
					CampaignInventory []struct {
						CampaignName string `json:"campaign_name"`
						Quantity     int    `json:"quantity"`
					} `json:"campaign_inventory"`
					CreatorInventory []struct {
						CreatorName string `json:"creator_name"`
						Quantity    int    `json:"quantity"`
					} `json:"creator_inventory"`
					InShopInventory struct {
						Quantity int `json:"quantity"`
					} `json:"in_shop_inventory"`
				} `json:"total_available_inventory_distribution"`
				TotalAvailableQuantity int `json:"total_available_quantity"`
				TotalCommittedQuantity int `json:"total_committed_quantity"`
				WarehouseInventory     []struct {
					AvailableQuantity int    `json:"available_quantity"`
					CommittedQuantity int    `json:"committed_quantity"`
					WarehouseId       string `json:"warehouse_id"`
				} `json:"warehouse_inventory"`
			} `json:"skus"`
		} `json:"inventory"`
	} `json:"data"`
}

type ProductInformationIssueDiagnosisRsp struct {
	BaseRsp
	Data struct {
		Products []struct {
			Diagnoses []struct {
				DiagnosisResults []struct {
					Code        string `json:"code"`
					HowToSolve  string `json:"how_to_solve"`
					QualityTier string `json:"quality_tier"`
				} `json:"diagnosis_results"`
				Field      string `json:"field"`
				Suggestion struct {
					Images []struct {
						Height       int    `json:"height"`
						OptimizedUri string `json:"optimized_uri"`
						OptimizedUrl string `json:"optimized_url"`
						Uri          string `json:"uri"`
						Url          string `json:"url"`
						Width        int    `json:"width"`
					} `json:"images"`
					SeoWords []struct {
						Text string `json:"text"`
					} `json:"seo_words"`
					SmartTexts []struct {
						Text string `json:"text"`
					} `json:"smart_texts"`
				} `json:"suggestion"`
			} `json:"diagnoses"`
			Id             string `json:"id"`
			ListingQuality struct {
				CurrentTier              string `json:"current_tier"`
				RemainingRecommendations int    `json:"remaining_recommendations"`
			} `json:"listing_quality"`
		} `json:"products"`
	} `json:"data"`
}

type GetProductsSEOWordsRsp struct {
	BaseRsp
	Data struct {
		Products []struct {
			Id       string `json:"id"`
			SeoWords []struct {
				Text string `json:"text"`
			} `json:"seo_words"`
		} `json:"products"`
	} `json:"data"`
}

type GetRecommendedProductTitleAndDescriptionRsp struct {
	BaseRsp
	Data struct {
		Products []struct {
			Id          string `json:"id"`
			Suggestions []struct {
				Field string `json:"field"`
				Items []struct {
					Text string `json:"text"`
				} `json:"items"`
			} `json:"suggestions"`
		} `json:"products"`
	} `json:"data"`
}

type OptimizedImagesRsp struct {
	BaseRsp
	Data struct {
		Images []struct {
			Height         int    `json:"height"`
			OptimizeStatus string `json:"optimize_status"`
			OptimizedUri   string `json:"optimized_uri"`
			OptimizedUrl   string `json:"optimized_url"`
			OriginalUri    string `json:"original_uri"`
			OriginalUrl    string `json:"original_url"`
			Width          int    `json:"width"`
		} `json:"images"`
	} `json:"data"`
}

type GetGlobalCategoriesRsp struct {
	BaseRsp
	Data struct {
		Categories []struct {
			Id                 string   `json:"id"`
			IsLeaf             bool     `json:"is_leaf"`
			LocalName          string   `json:"local_name"`
			ParentId           string   `json:"parent_id"`
			PermissionStatuses []string `json:"permission_statuses"`
		} `json:"categories"`
	} `json:"data"`
}

type RecommendGlobalCategoriesRsp struct {
	BaseRsp
	Data struct {
		Categories []struct {
			Id     string `json:"id"`
			IsLeaf bool   `json:"is_leaf"`
			Level  int    `json:"level"`
			Name   string `json:"name"`
		} `json:"categories"`
		LeafCategoryId string `json:"leaf_category_id"`
	} `json:"data"`
}

type GetGlobalCategoryRulesRsp struct {
	BaseRsp
	Data struct {
		ProductCertifications []struct {
			Id              string   `json:"id"`
			IsRequired      bool     `json:"is_required"`
			Name            string   `json:"name"`
			OptionalRegions []string `json:"optional_regions"`
			RequiredRegions []string `json:"required_regions"`
			SampleImageUrl  string   `json:"sample_image_url"`
		} `json:"product_certifications"`
		SizeChart struct {
			IsRequired  bool `json:"is_required"`
			IsSupported bool `json:"is_supported"`
		} `json:"size_chart"`
	} `json:"data"`
}

type GetGlobalAttributesRsp struct {
	BaseRsp
	Data struct {
		Attributes []struct {
			Id                  string `json:"id"`
			IsCustomizable      bool   `json:"is_customizable"`
			IsMultipleSelection bool   `json:"is_multiple_selection"`
			IsRequried          bool   `json:"is_requried"`
			Name                string `json:"name"`
			Type                string `json:"type"`
			Values              []struct {
				Id   string `json:"id"`
				Name string `json:"name"`
			} `json:"values"`
		} `json:"attributes"`
	} `json:"data"`
}

type CreateGlobalProductRsp struct {
	BaseRsp
	Data struct {
		GlobalProductId string `json:"global_product_id"`
		GlobalSkus      []struct {
			Id              string `json:"id"`
			SalesAttributes []struct {
				Id      string `json:"id"`
				ValueId string `json:"value_id"`
			} `json:"sales_attributes"`
			SellerSku string `json:"seller_sku"`
		} `json:"global_skus"`
	} `json:"data"`
}

type PublishGlobalProductRsp struct {
	BaseRsp
	Data struct {
		Products []struct {
			Id     string `json:"id"`
			Region string `json:"region"`
			ShopId string `json:"shop_id"`
			Skus   []struct {
				Id                 string `json:"id"`
				RelatedGlobalSkuId string `json:"related_global_sku_id"`
				SaleAttributes     []struct {
					Id      string `json:"id"`
					ValueId string `json:"value_id"`
				} `json:"sale_attributes"`
				SellerSku string `json:"seller_sku"`
			} `json:"skus"`
		} `json:"products"`
		PublishResult []struct {
			FailReasons []struct {
				Message string `json:"message"`
			} `json:"fail_reasons"`
			Region string `json:"region"`
			Status string `json:"status"`
		} `json:"publish_result"`
	} `json:"data"`
}

type EditGlobalProductRsp struct {
	BaseRsp
	Data struct {
		GlobalSkus []struct {
			Id              string `json:"id"`
			SalesAttributes []struct {
				Id      string `json:"id"`
				ValueId string `json:"value_id"`
			} `json:"sales_attributes"`
			SellerSku string `json:"seller_sku"`
		} `json:"global_skus"`
		PublishResults []struct {
			FailReasons []struct {
				Message string `json:"message"`
			} `json:"fail_reasons"`
			Region string `json:"region"`
			Status string `json:"status"`
		} `json:"publish_results"`
	} `json:"data"`
}

type DeleteGlobalProductsRsp struct {
	BaseRsp
	Data struct {
		Errors []struct {
			Code   int `json:"code"`
			Detail struct {
				GlobalProductId string `json:"global_product_id"`
			} `json:"detail"`
			Message string `json:"message"`
		} `json:"errors"`
	} `json:"data"`
}

type GetGlobalProductRsp struct {
	BaseRsp
	Data struct {
		Brand struct {
			Id string `json:"id"`
		} `json:"brand"`
		Category struct {
			Id string `json:"id"`
		} `json:"category"`
		Certifications []struct {
			Files []struct {
				Format string   `json:"format"`
				Id     string   `json:"id"`
				Name   string   `json:"name"`
				Urls   []string `json:"urls"`
			} `json:"files"`
			Id     string `json:"id"`
			Images []struct {
				Height int    `json:"height"`
				Uri    string `json:"uri"`
				Width  int    `json:"width"`
			} `json:"images"`
			Title string `json:"title"`
		} `json:"certifications"`
		CreateTime     int64  `json:"create_time"`
		Description    string `json:"description"`
		GlobalSellerId string `json:"global_seller_id"`
		Id             string `json:"id"`
		MainImages     []struct {
			Height int    `json:"height"`
			Uri    string `json:"uri"`
			Width  int    `json:"width"`
		} `json:"main_images"`
		Manufacturer struct {
			Address     string `json:"address"`
			Email       string `json:"email"`
			Name        string `json:"name"`
			PhoneNumber string `json:"phone_number"`
		} `json:"manufacturer"`
		PackageDimensions struct {
			Height string `json:"height"`
			Length string `json:"length"`
			Unit   string `json:"unit"`
			Width  string `json:"width"`
		} `json:"package_dimensions"`
		PackageWeight struct {
			Unit  string `json:"unit"`
			Value string `json:"value"`
		} `json:"package_weight"`
		ProductAttributes []struct {
			Id     string `json:"id"`
			Name   string `json:"name"`
			Values []struct {
				Id   string `json:"id"`
				Name string `json:"name"`
			} `json:"values"`
		} `json:"product_attributes"`
		Products []struct {
			Id     string `json:"id"`
			Region string `json:"region"`
		} `json:"products"`
		SizeChart struct {
			Image struct {
				Height int    `json:"height"`
				Uri    string `json:"uri"`
				Width  int    `json:"width"`
			} `json:"image"`
			Template struct {
				Id string `json:"id"`
			} `json:"template"`
		} `json:"size_chart"`
		Skus []struct {
			GlobalQuantity int    `json:"global_quantity"`
			Id             string `json:"id"`
			IdentifierCode struct {
				Code string `json:"code"`
				Type string `json:"type"`
			} `json:"identifier_code"`
			Inventory []struct {
				GlobalWarehouseId string `json:"global_warehouse_id"`
				Quantity          int    `json:"quantity"`
			} `json:"inventory"`
			Price struct {
				Amount    string `json:"amount"`
				Currency  string `json:"currency"`
				UnitPrice string `json:"unit_price"`
			} `json:"price"`
			SalesAttributes []struct {
				Id     string `json:"id"`
				Name   string `json:"name"`
				SkuImg struct {
					Height    int      `json:"height"`
					ThumbUrls []string `json:"thumb_urls"`
					Uri       string   `json:"uri"`
					Urls      []string `json:"urls"`
					Width     int      `json:"width"`
				} `json:"sku_img"`
				ValueId   string `json:"value_id"`
				ValueName string `json:"value_name"`
			} `json:"sales_attributes"`
			SellerSku    string `json:"seller_sku"`
			SkuUnitCount string `json:"sku_unit_count"`
		} `json:"skus"`
		Title      string `json:"title"`
		UpdateTime int64  `json:"update_time"`
		Video      struct {
			Id string `json:"id"`
		} `json:"video"`
	} `json:"data"`
}

type SearchGlobalProductsRsp struct {
	BaseRsp
	Data struct {
		GlobalProducts []struct {
			CreateTime int    `json:"create_time"`
			Id         string `json:"id"`
			Skus       []struct {
				Id        string `json:"id"`
				SellerSku string `json:"seller_sku"`
			} `json:"skus"`
			Status     string `json:"status"`
			Title      string `json:"title"`
			UpdateTime int    `json:"update_time"`
		} `json:"global_products"`
		NextPageToken string `json:"next_page_token"`
		TotalCount    int    `json:"total_count"`
	} `json:"data"`
}

type SearchGlobalProductsV202309Rsp struct {
	BaseRsp
	Data struct {
		TotalCount     int `json:"total_count"`
		GlobalProducts []struct {
			Id     string `json:"id"`
			Title  string `json:"title"`
			Status string `json:"status"`
			Skus   []struct {
				Id        string `json:"id"`
				SellerSku string `json:"seller_sku"`
			} `json:"skus"`
			CreateTime int64 `json:"create_time"`
			UpdateTime int64 `json:"update_time"`
		} `json:"global_products"`
		NextPageToken string `json:"next_page_token"`
	} `json:"data"`
}

type UpdateGlobalInventoryRsp struct {
	BaseRsp
}

type CreateManufacturerRsp struct {
	BaseRsp
	Data struct {
		ManufacturerId string `json:"manufacturer_id"`
	} `json:"data"`
}

type PartialEditManufacturerRsp struct {
	BaseRsp
}

type SearchManufacturersRsp struct {
	BaseRsp
	Data struct {
		Manufacturers []struct {
			Address     string `json:"address"`
			Email       string `json:"email"`
			Id          string `json:"id"`
			Name        string `json:"name"`
			PhoneNumber struct {
				CountryCode string `json:"country_code"`
				LocalNumber string `json:"local_number"`
			} `json:"phone_number"`
			RegisteredTradeName string `json:"registered_trade_name"`
		} `json:"manufacturers"`
		NextPageToken string `json:"next_page_token"`
		TotalCount    int    `json:"total_count"`
	} `json:"data"`
}

type SearchManufacturersRspV202501 struct {
	BaseRsp
	Data struct {
		Manufacturers []struct {
			Id               string `json:"id"`
			RegionalProfiles []struct {
				Address     string `json:"address"`
				Email       string `json:"email"`
				Locale      string `json:"locale"`
				Name        string `json:"name"`
				PhoneNumber struct {
					CountryCode string `json:"country_code"`
					LocalNumber string `json:"local_number"`
				} `json:"phone_number"`
				RegisteredTradeName string `json:"registered_trade_name"`
			} `json:"regional_profiles"`
		} `json:"manufacturers"`
		NextPageToken string `json:"next_page_token"`
		TotalCount    int    `json:"total_count"`
	} `json:"data"`
}

type CreateResponsiblePersonRsp struct {
	BaseRsp
	Data struct {
		ResponsiblePersonId string `json:"responsible_person_id"`
	} `json:"data"`
}

type PartialEditResponsiblePersonRsp struct {
	BaseRsp
}

type SearchResponsiblePersonsRsp struct {
	BaseRsp
	Data struct {
		NextPageToken      string `json:"next_page_token"`
		ResponsiblePersons []struct {
			Address struct {
				City               string `json:"city"`
				Country            string `json:"country"`
				District           string `json:"district"`
				PostalCode         string `json:"postal_code"`
				Province           string `json:"province"`
				StreetAddressLine1 string `json:"street_address_line1"`
				StreetAddressLine2 string `json:"street_address_line2"`
			} `json:"address"`
			Email       string `json:"email"`
			Id          string `json:"id"`
			Name        string `json:"name"`
			PhoneNumber struct {
				CountryCode string `json:"country_code"`
				LocalNumber string `json:"local_number"`
			} `json:"phone_number"`
		} `json:"responsible_persons"`
		TotalCount int `json:"total_count"`
	} `json:"data"`
}

type SearchResponsiblePersonsV202501Rsp struct {
	BaseRsp
	Data struct {
		NextPageToken      string `json:"next_page_token"`
		ResponsiblePersons []struct {
			Id               string `json:"id"`
			RegionalProfiles []struct {
				Address struct {
					City               string `json:"city"`
					Country            string `json:"country"`
					District           string `json:"district"`
					PostalCode         string `json:"postal_code"`
					Province           string `json:"province"`
					StreetAddressLine1 string `json:"street_address_line1"`
					StreetAddressLine2 string `json:"street_address_line2"`
				} `json:"address"`
				Email       string `json:"email"`
				Locale      string `json:"locale"`
				Name        string `json:"name"`
				PhoneNumber struct {
					CountryCode string `json:"country_code"`
					LocalNumber string `json:"local_number"`
				} `json:"phone_number"`
			} `json:"regional_profiles"`
		} `json:"responsible_persons"`
		TotalCount int `json:"total_count"`
	} `json:"data"`
}

type CreateCategoryUpgradeTaskRsp struct {
	BaseRsp
}

type ListingSchemasRsp struct {
	BaseRsp
	Data struct {
		Errors []struct {
			Code   int `json:"code"`
			Detail struct {
				CategoryId int `json:"category_id"`
			} `json:"detail"`
			Message string `json:"message"`
		} `json:"errors"`
		ListingSchemas []struct {
			CategoryId int `json:"category_id"`
			Fields     []struct {
				ComplexValues []struct {
					Id      string `json:"id"`
					Name    string `json:"name"`
					Options []struct {
						Id   string `json:"id"`
						Name string `json:"name"`
					} `json:"options"`
					Rules []struct {
						Type  string `json:"type"`
						Value string `json:"value"`
					} `json:"rules"`
				} `json:"complex_values"`
				Id      string `json:"id"`
				Name    string `json:"name"`
				Options []struct {
					Id   string `json:"id"`
					Name string `json:"name"`
				} `json:"options"`
				Rules []struct {
					Type  string `json:"type"`
					Value string `json:"value"`
				} `json:"rules"`
			} `json:"fields"`
		} `json:"listing_schemas"`
	} `json:"data"`
}

type DiagnoseAndOptimizeProductRsp struct {
	BaseRsp
	Data struct {
		Diagnoses []struct {
			DiagnosisResults []struct {
				Code        string `json:"code"`
				HowToSolve  string `json:"how_to_solve"`
				QualityTier string `json:"quality_tier"`
			} `json:"diagnosis_results"`
			Field      string `json:"field"`
			Suggestion struct {
				Images []struct {
					Height       int    `json:"height"`
					OptimizedUri string `json:"optimized_uri"`
					OptimizedUrl string `json:"optimized_url"`
					Uri          string `json:"uri"`
					Url          string `json:"url"`
					Width        int    `json:"width"`
				} `json:"images"`
				SeoWords []struct {
					Text string `json:"text"`
				} `json:"seo_words"`
				SmartTexts []struct {
					Text string `json:"text"`
				} `json:"smart_texts"`
			} `json:"suggestion"`
		} `json:"diagnoses"`
		ListingQuality struct {
			CurrentTier              string `json:"current_tier"`
			RemainingRecommendations int    `json:"remaining_recommendations"`
		} `json:"listing_quality"`
	} `json:"data"`
}
