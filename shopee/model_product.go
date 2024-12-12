package shopee

type GetCategoryRsp struct {
	BaseRsp
	DebugMessage string `json:"debug_message"`
	Response     struct {
		CategoryList []struct {
			CategoryId           int         `json:"category_id"`
			ParentCategoryId     int         `json:"parent_category_id"`
			OriginalCategoryName string      `json:"original_category_name"`
			DisplayCategoryName  string      `json:"display_category_name"`
			HasChildren          bool        `json:"has_children"`
			DebugMessage         interface{} `json:"debug_message"`
		} `json:"category_list"`
	} `json:"response"`
}

type GetAttributeTreeRsp struct {
	BaseRsp
	Response struct {
		List []struct {
			CategoryId    int `json:"category_id"`
			AttributeTree []struct {
				AttributeId        int    `json:"attribute_id"`
				Mandatory          bool   `json:"mandatory"`
				Name               string `json:"name"`
				AttributeValueList []struct {
					ValueId            int    `json:"value_id"`
					Name               string `json:"name"`
					ChildAttributeList []struct {
						AttributeId   int    `json:"attribute_id"`
						Mandatory     bool   `json:"mandatory"`
						Name          string `json:"name"`
						AttributeInfo struct {
							InputType           int `json:"input_type"`
							InputValidationType int `json:"input_validation_type"`
							FormatType          int `json:"format_type"`
						} `json:"attribute_info"`
						MultiLang []struct {
							Language string `json:"language"`
							Value    string `json:"value"`
						} `json:"multi_lang"`
					} `json:"child_attribute_list,omitempty"`
					MultiLang []struct {
						Language string `json:"language"`
						Value    string `json:"value"`
					} `json:"multi_lang"`
				} `json:"attribute_value_list"`
				AttributeInfo struct {
					InputType           int `json:"input_type"`
					InputValidationType int `json:"input_validation_type"`
					FormatType          int `json:"format_type"`
				} `json:"attribute_info"`
				MultiLang []struct {
					Language string `json:"language"`
					Value    string `json:"value"`
				} `json:"multi_lang"`
			} `json:"attribute_tree"`
		} `json:"list"`
	} `json:"response"`
}

type GetBrandListRsp struct {
	BaseRsp
	Response struct {
		BrandList []struct {
			BrandId           int64  `json:"brand_id"`
			OriginalBrandName string `json:"original_brand_name"`
			DisplayBrandName  string `json:"display_brand_name"`
		} `json:"brand_list"`
		HasNextPage bool   `json:"has_next_page"`
		NextOffset  int    `json:"next_offset"`
		IsMandatory bool   `json:"is_mandatory"`
		InputType   string `json:"input_type"`
	} `json:"response"`
}

type GetItemLimitRsp struct {
	BaseRsp
	Response struct {
		DimensionLimit struct {
			DimensionMandatory bool `json:"dimension_mandatory"`
		} `json:"dimension_limit"`
		DtsLimit struct {
			DaysToShipLimit struct {
				MaxLimit int `json:"max_limit"`
				MinLimit int `json:"min_limit"`
			} `json:"days_to_ship_limit"`
			NonPreOrderDaysToShip int  `json:"non_pre_order_days_to_ship"`
			SupportPreOrder       bool `json:"support_pre_order"`
		} `json:"dts_limit"`
		ExtendedDescriptionLimit struct {
			DescriptionImageAspectRatioMax float64 `json:"description_image_aspect_ratio_max"`
			DescriptionImageAspectRatioMin float64 `json:"description_image_aspect_ratio_min"`
			DescriptionImageHeightMin      int     `json:"description_image_height_min"`
			DescriptionImageNumMax         int     `json:"description_image_num_max"`
			DescriptionImageNumMin         int     `json:"description_image_num_min"`
			DescriptionImageWidthMin       int     `json:"description_image_width_min"`
			DescriptionTextLengthMax       int     `json:"description_text_length_max"`
			DescriptionTextLengthMin       int     `json:"description_text_length_min"`
		} `json:"extended_description_limit"`
		ItemCountLimit struct {
			MaxLimit int `json:"max_limit"`
		} `json:"item_count_limit"`
		ItemDescriptionLengthLimit struct {
			MaxLimit int `json:"max_limit"`
			MinLimit int `json:"min_limit"`
		} `json:"item_description_length_limit"`
		ItemImageCountLimit struct {
			MaxLimit int `json:"max_limit"`
			MinLimit int `json:"min_limit"`
		} `json:"item_image_count_limit"`
		ItemNameLengthLimit struct {
			MaxLimit int `json:"max_limit"`
			MinLimit int `json:"min_limit"`
		} `json:"item_name_length_limit"`
		PriceLimit struct {
			MaxLimit float64 `json:"max_limit"`
			MinLimit float64 `json:"min_limit"`
		} `json:"price_limit"`
		StockLimit struct {
			MaxLimit int `json:"max_limit"`
			MinLimit int `json:"min_limit"`
		} `json:"stock_limit"`
		TierVariationNameLengthLimit struct {
			MaxLimit int `json:"max_limit"`
			MinLimit int `json:"min_limit"`
		} `json:"tier_variation_name_length_limit"`
		TierVariationOptionLengthLimit struct {
			MaxLimit int `json:"max_limit"`
			MinLimit int `json:"min_limit"`
		} `json:"tier_variation_option_length_limit"`
		WeightLimit struct {
			WeightMandatory bool `json:"weight_mandatory"`
		} `json:"weight_limit"`
		WholesalePriceThresholdPercentage struct {
			MaxLimit int `json:"max_limit"`
			MinLimit int `json:"min_limit"`
		} `json:"wholesale_price_threshold_percentage"`
	} `json:"response"`
	Warning string `json:"warning"`
}

type GetProductListRsp struct {
	BaseRsp
	Response struct {
		Item []struct {
			ItemId     int64  `json:"item_id"`
			ItemStatus string `json:"item_status"`
			UpdateTime int    `json:"update_time"`
		} `json:"item"`
		TotalCount  int    `json:"total_count"`
		HasNextPage bool   `json:"has_next_page"`
		Next        string `json:"next"`
	} `json:"response"`
}

type GetItemListRsp struct {
	BaseRsp
	Response struct {
		Item []struct {
			ItemId     int64  `json:"item_id"`
			ItemStatus string `json:"item_status"`
			Tag        struct {
				Kit bool `json:"kit"`
			} `json:"tag"`
			UpdateTime int64 `json:"update_time"`
		} `json:"item"`
		TotalCount  int  `json:"total_count"`
		HasNextPage bool `json:"has_next_page"`
		NextOffset  int  `json:"next_offset"`
	} `json:"response"`
}

type GetItemBaseInfoRsp struct {
	BaseRsp
	Response struct {
		ItemList []struct {
			ItemId        int64  `json:"item_id"`
			CategoryId    int    `json:"category_id"`
			ItemName      string `json:"item_name"`
			Description   string `json:"description"`
			ItemSku       string `json:"item_sku"`
			CreateTime    int    `json:"create_time"`
			UpdateTime    int    `json:"update_time"`
			AttributeList []struct {
				AttributeId           int    `json:"attribute_id"`
				OriginalAttributeName string `json:"original_attribute_name"`
				AttributeValueList    []struct {
					ValueId           int    `json:"value_id"`
					OriginalValueName string `json:"original_value_name"`
					ValueUnit         string `json:"value_unit"`
				} `json:"attribute_value_list"`
				IsMandatory bool `json:"is_mandatory"`
			} `json:"attribute_list,omitempty"`
			Image struct {
				ImageIdList  []string `json:"image_id_list"`
				ImageUrlList []string `json:"image_url_list"`
			} `json:"image"`
			Weight    string `json:"weight"`
			Dimension struct {
				PackageLength float64 `json:"package_length"`
				PackageWidth  float64 `json:"package_width"`
				PackageHeight float64 `json:"package_height"`
			} `json:"dimension"`
			LogisticInfo []struct {
				LogisticId           int     `json:"logistic_id"`
				LogisticName         string  `json:"logistic_name"`
				Enabled              bool    `json:"enabled"`
				SizeId               int     `json:"size_id"`
				IsFree               bool    `json:"is_free"`
				EstimatedShippingFee float64 `json:"estimated_shipping_fee,omitempty"`
			} `json:"logistic_info"`
			PreOrder struct {
				IsPreOrder bool `json:"is_pre_order"`
				DaysToShip int  `json:"days_to_ship"`
			} `json:"pre_order"`
			Condition  string `json:"condition"`
			SizeChart  string `json:"size_chart"`
			ItemStatus string `json:"item_status"`
			HasModel   bool   `json:"has_model"`
			Brand      struct {
				BrandId           int    `json:"brand_id"`
				OriginalBrandName string `json:"original_brand_name"`
			} `json:"brand"`
			ItemDangerous   int    `json:"item_dangerous"`
			DescriptionType string `json:"description_type"`
			SizeChartId     int    `json:"size_chart_id"`
			VideoInfo       []struct {
				VideoUrl     string `json:"video_url"`
				ThumbnailUrl string `json:"thumbnail_url"`
				Duration     int    `json:"duration"`
			} `json:"video_info,omitempty"`
		} `json:"item_list"`
	} `json:"response"`
}

type GetItemExtraInfoRsp struct {
	BaseRsp
	Response struct {
		ItemList []struct {
			ItemId       int64 `json:"item_id"`
			Sale         int   `json:"sale"`
			Views        int   `json:"views"`
			Likes        int   `json:"likes"`
			RatingStar   int   `json:"rating_star"`
			CommentCount int   `json:"comment_count"`
		} `json:"item_list"`
	} `json:"response"`
	DebugMessage string `json:"debug_message"`
}

type AddItemRsp struct {
	BaseRsp
	Response struct {
		Description string `json:"description"`
		Weight      int    `json:"weight"`
		PreOrder    struct {
			DaysToShip int  `json:"days_to_ship"`
			IsPreOrder bool `json:"is_pre_order"`
		} `json:"pre_order"`
		ItemName string `json:"item_name"`
		Images   struct {
			ImageIdList  []string `json:"image_id_list"`
			ImageUrlList []string `json:"image_url_list"`
		} `json:"images"`
		ItemStatus string `json:"item_status"`
		PriceInfo  struct {
			CurrentPrice  float64 `json:"current_price"`
			OriginalPrice float64 `json:"original_price"`
		} `json:"price_info"`
		LogisticInfo []struct {
			SizeId      int     `json:"size_id"`
			ShippingFee float64 `json:"shipping_fee"`
			Enabled     bool    `json:"enabled"`
			LogisticId  int     `json:"logistic_id"`
			IsFree      bool    `json:"is_free"`
		} `json:"logistic_info"`
		ItemId     int64 `json:"item_id"`
		Attributes []struct {
			AttributeId        int `json:"attribute_id"`
			AttributeValueList []struct {
				OriginalValueName string `json:"original_value_name"`
				ValueId           int    `json:"value_id"`
				ValueUnit         string `json:"value_unit"`
			} `json:"attribute_value_list"`
		} `json:"attributes"`
		CategoryId int `json:"category_id"`
		Dimension  struct {
			PackageWidth  int `json:"package_width"`
			PackageLength int `json:"package_length"`
			PackageHeight int `json:"package_height"`
		} `json:"dimension"`
		Condition string `json:"condition"`
		VideoInfo []struct {
			VideoUrl     string `json:"video_url"`
			ThumbnailUrl string `json:"thumbnail_url"`
			Duration     int    `json:"duration"`
		} `json:"video_info"`
		Wholesale []struct {
			MinCount  int     `json:"min_count"`
			MaxCount  int     `json:"max_count"`
			UnitPrice float64 `json:"unit_price"`
		} `json:"wholesale"`
		Brand struct {
			BrandId           int    `json:"brand_id"`
			OriginalBrandName string `json:"original_brand_name"`
		} `json:"brand"`
		ItemDangerous   int `json:"item_dangerous"`
		DescriptionInfo struct {
			ExtendedDescription struct {
				FieldList []struct {
					FieldType string `json:"field_type"`
					Text      string `json:"text"`
					ImageInfo struct {
						ImageId string `json:"image_id"`
					} `json:"image_info"`
				} `json:"field_list"`
			} `json:"extended_description"`
		} `json:"description_info"`
		DescriptionType string `json:"description_type"`
		ComplaintPolicy struct {
			WarrantyTime                string `json:"warranty_time"`
			ExcludeEntrepreneurWarranty bool   `json:"exclude_entrepreneur_warranty"`
			ComplaintAddressId          int    `json:"complaint_address_id"`
			AdditionalInformation       string `json:"additional_information"`
		} `json:"complaint_policy"`
		SellerStock []struct {
			LocationId string `json:"location_id"`
			Stock      int    `json:"stock"`
		} `json:"seller_stock"`
	} `json:"response"`
}

type UpdateItemRsp struct {
	BaseRsp
	Description string  `json:"description"`
	Weight      float64 `json:"weight"`
	PreOrder    struct {
		DaysToShip int  `json:"days_to_ship"`
		IsPreOrder bool `json:"is_pre_order"`
	} `json:"pre_order"`
	ItemName      string `json:"item_name"`
	AttributeList []struct {
		AttributeId        int `json:"attribute_id"`
		AttributeValueList []struct {
			ValueId           int    `json:"value_id"`
			OriginalValueName string `json:"original_value_name"`
			ValueUnit         string `json:"value_unit"`
		} `json:"attribute_value_list"`
	} `json:"attribute_list"`
	Image struct {
		ImageIdList []struct {
		} `json:"image_id_list"`
	} `json:"image"`
	ItemSku      string `json:"item_sku"`
	ItemStatus   string `json:"item_status"`
	LogisticInfo []struct {
		SizeId      int     `json:"size_id"`
		ShippingFee float64 `json:"shipping_fee"`
		Enabled     bool    `json:"enabled"`
		LogisticId  int     `json:"logistic_id"`
		IsFree      bool    `json:"is_free"`
	} `json:"logistic_info"`
	Wholesale []struct {
		MinCount  int     `json:"min_count"`
		UnitPrice float64 `json:"unit_price"`
		MaxCount  int     `json:"max_count"`
	} `json:"wholesale"`
	ItemId     int64 `json:"item_id"`
	CategoryId int   `json:"category_id"`
	Dimension  struct {
		PackageHeight int `json:"package_height"`
		PackageLength int `json:"package_length"`
		PackageWidth  int `json:"package_width"`
	} `json:"dimension"`
	Condition     string `json:"condition"`
	VideoUploadId []struct {
	} `json:"video_upload_id"`
	Brand struct {
		BrandId           int    `json:"brand_id"`
		OriginalBrandName string `json:"original_brand_name"`
	} `json:"brand"`
	ItemDangerous int `json:"item_dangerous"`
	TaxInfo       struct {
		Ncm           string `json:"ncm"`
		SameStateCfop string `json:"same_state_cfop"`
		DiffStateCfop string `json:"diff_state_cfop"`
		Csosn         string `json:"csosn"`
		Origin        string `json:"origin"`
		Cest          string `json:"cest"`
		MeasureUnit   string `json:"measure_unit"`
		InvoiceOption string `json:"invoice_option"`
		VatRate       string `json:"vat_rate"`
		HsCode        string `json:"hs_code"`
		TaxCode       string `json:"tax_code"`
	} `json:"tax_info"`
	ComplaintPolicy struct {
		WarrantyTime                string `json:"warranty_time"`
		ExcludeEntrepreneurWarranty bool   `json:"exclude_entrepreneur_warranty"`
		ComplaintAddressId          int    `json:"complaint_address_id"`
		AdditionalInformation       string `json:"additional_information"`
	} `json:"complaint_policy"`
	DescriptionInfo struct {
		ExtendedDescription struct {
			FieldList []struct {
				FieldType string `json:"field_type"`
				Text      string `json:"text"`
				ImageInfo struct {
					ImageId string `json:"image_id"`
				} `json:"image_info"`
			} `json:"field_list"`
		} `json:"extended_description"`
	} `json:"description_info"`
	DescriptionType string `json:"description_type"`
}

type DeleteItemRsp struct {
	BaseRsp
}

type InitTierVariationRsp struct {
	BaseRsp
	Response struct {
		TierVariation []struct {
			Name       string `json:"name"`
			OptionList []struct {
				Image struct {
					ImageUrl string `json:"image_url"`
				} `json:"image"`
				Option string `json:"option"`
			} `json:"option_list"`
		} `json:"tier_variation"`
		Model []struct {
			TierIndex []struct {
			} `json:"tier_index"`
			ModelSku  string `json:"model_sku"`
			PriceInfo []struct {
				OriginalPrice float64 `json:"original_price"`
			} `json:"price_info"`
			SellerStock []struct {
				LocationId string `json:"location_id"`
				Stock      int    `json:"stock"`
			} `json:"seller_stock"`
			Weight float64 `json:"weight"`
		} `json:"model"`
	} `json:"response"`
}

type UpdateTierVariationRsp struct {
	BaseRsp
}

type GetModelListRsp struct {
	BaseRsp
	Response struct {
		TierVariation []struct {
			Name       string `json:"name"`
			OptionList []struct {
				Option string `json:"option"`
				Image  struct {
					ImageUrl string `json:"image_url"`
					ImageId  string `json:"image_id"`
				} `json:"image,omitempty"`
			} `json:"option_list"`
		} `json:"tier_variation"`
		Model []struct {
			ModelId     int64  `json:"model_id"`
			ModelName   string `json:"model_name"`
			PromotionId int    `json:"promotion_id"`
			TierIndex   []int  `json:"tier_index"`
			PriceInfo   []struct {
				CurrentPrice                 float64 `json:"current_price"`
				OriginalPrice                float64 `json:"original_price"`
				InflatedPriceOfCurrentPrice  float64 `json:"inflated_price_of_current_price"`
				InflatedPriceOfOriginalPrice float64 `json:"inflated_price_of_original_price"`
				Currency                     string  `json:"currency"`
			} `json:"price_info"`
			ModelSku string `json:"model_sku"`
			PreOrder struct {
				IsPreOrder bool `json:"is_pre_order"`
				DaysToShip int  `json:"days_to_ship"`
			} `json:"pre_order"`
			StockInfoV2 struct {
				SummaryInfo struct {
					TotalReservedStock  int `json:"total_reserved_stock"`
					TotalAvailableStock int `json:"total_available_stock"`
				} `json:"summary_info"`
				SellerStock []struct {
					LocationId string `json:"location_id"`
					Stock      int    `json:"stock"`
					IfSaleable bool   `json:"if_saleable"`
				} `json:"seller_stock"`
				ShopeeStock []struct {
					LocationId string `json:"location_id"`
					Stock      int    `json:"stock"`
				} `json:"shopee_stock"`
			} `json:"stock_info_v2"`
			ModelStatus string `json:"model_status"`
		} `json:"model"`
	} `json:"response"`
}

type UpdateModelRsp struct {
	BaseRsp
}

type DeleteModelRsp struct {
	BaseRsp
}

type SupportSizeChartRsp struct {
	BaseRsp
	Response struct {
		SupportSizeChart bool `json:"support_size_chart"`
	} `json:"response"`
}

type UpdateSizeChartRsp struct {
	BaseRsp
}

type UnlistItemRsp struct {
	BaseRsp
	Response struct {
		FailureList []struct {
			ItemId       int64  `json:"item_id"`
			FailedReason string `json:"failed_reason"`
		} `json:"failure_list"`
		SuccessList []struct {
			ItemId int64 `json:"item_id"`
			Unlist bool  `json:"unlist"`
		} `json:"success_list"`
	} `json:"response"`
}

type UpdatePriceRsp struct {
	BaseRsp
	Response struct {
		FailureList []struct {
			ModelId      int    `json:"model_id"`
			FailedReason string `json:"failed_reason"`
		} `json:"failure_list"`
		SuccessList []struct {
			ModelId       int     `json:"model_id"`
			OriginalPrice float64 `json:"original_price"`
		} `json:"success_list"`
	} `json:"response"`
}

type UpdateStockRsp struct {
	BaseRsp
	Response struct {
		FailureList []struct {
			ModelId      int    `json:"model_id"`
			FailedReason string `json:"failed_reason"`
		} `json:"failure_list"`
		SuccessList []struct {
			ModelId    int    `json:"model_id"`
			LocationId string `json:"location_id"`
			Stock      int    `json:"stock"`
		} `json:"success_list"`
	} `json:"response"`
}

type BoostItemRsp struct {
	BaseRsp
	Response struct {
		FailureList []struct {
			ItemId       int64  `json:"item_id"`
			FailedReason string `json:"failed_reason"`
		} `json:"failure_list"`
		SuccessList struct {
			ItemIdList []int64 `json:"item_id_list"`
		} `json:"success_list"`
	} `json:"response"`
}

type GetBoostedListRsp struct {
	BaseRsp
	Response struct {
		ItemList []struct {
			ItemId         int64 `json:"item_id"`
			CoolDownSecond int   `json:"cool_down_second"`
		} `json:"item_list"`
	} `json:"response"`
}

type GetItemPromotionRsp struct {
	BaseRsp
	Response struct {
		SuccessList []struct {
			ItemId    int64 `json:"item_id"`
			Promotion []struct {
				PromotionType      string `json:"promotion_type"`
				PromotionId        int64  `json:"promotion_id"`
				ModelId            int    `json:"model_id"`
				StartTime          int64  `json:"start_time"`
				EndTime            int64  `json:"end_time"`
				PromotionPriceInfo []struct {
					PromotionPrice float64 `json:"promotion_price"`
				} `json:"promotion_price_info"`
				PromotionStaging     string `json:"promotion_staging"`
				PromotionStockInfoV2 struct {
					SummaryInfo struct {
						TotalReservedStock int `json:"total_reserved_stock"`
					} `json:"summary_info"`
				} `json:"promotion_stock_info_v2"`
			} `json:"promotion"`
		} `json:"success_list"`
	} `json:"response"`
}

type UpdateSipItemPriceRsp struct {
	BaseRsp
}

type SearchItemRsp struct {
	BaseRsp
	Response struct {
		ItemIdList []int  `json:"item_id_list"`
		NextOffset string `json:"next_offset"`
		TotalCount int    `json:"total_count"`
	} `json:"response"`
}

type GetCommentRsp struct {
	BaseRsp
	Response struct {
		ItemCommentList []struct {
			CommentId     int    `json:"comment_id"`
			Comment       string `json:"comment"`
			BuyerUsername string `json:"buyer_username"`
			OrderSn       string `json:"order_sn"`
			ItemId        int64  `json:"item_id"`
			ModelId       int    `json:"model_id"`
			CreateTime    int64  `json:"create_time"`
			RatingStar    int64  `json:"rating_star"`
			Editable      string `json:"editable"`
			Hidden        bool   `json:"hidden"`
			CmtReply      struct {
				Reply  string `json:"reply"`
				Hidden bool   `json:"hidden"`
			} `json:"cmt_reply"`
		} `json:"item_comment_list"`
		More       bool   `json:"more"`
		NextCursor string `json:"next_cursor"`
	} `json:"response"`
}

type ReplyCommentRsp struct {
	BaseRsp
	Response struct {
		ResultList []struct {
			CommentId int `json:"comment_id"`
		} `json:"result_list"`
	} `json:"response"`
}

type CategoryRecommendRsp struct {
	BaseRsp
	Response struct {
		CategoryId []int `json:"category_id"`
	} `json:"response"`
}

type RegisterBrandRsp struct {
	BaseRsp
	Response struct {
		BrandId           int    `json:"brand_id"`
		OriginalBrandName string `json:"original_brand_name"`
	} `json:"response"`
}

type GetRecommendAttributeRsp struct {
	BaseRsp
	Response struct {
		AttributeList []struct {
			AttributeId           int      `json:"attribute_id"`
			OriginalAttributeName string   `json:"original_attribute_name"`
			DisplayAttributeName  string   `json:"display_attribute_name"`
			IsMandatory           bool     `json:"is_mandatory"`
			InputValidationType   string   `json:"input_validation_type"`
			FormatType            string   `json:"format_type"`
			DateFormatType        string   `json:"date_format_type,omitempty"`
			InputType             string   `json:"input_type"`
			AttributeUnit         []string `json:"attribute_unit"`
			AttributeValueList    []struct {
				ValueId           int    `json:"value_id"`
				OriginalValueName string `json:"original_value_name"`
				DisplayValueName  string `json:"display_value_name"`
				ValueUnit         string `json:"value_unit,omitempty"`
			} `json:"attribute_value_list"`
		} `json:"attribute_list"`
	} `json:"response"`
}

type GetProductInfoRsp struct {
	BaseRsp
	Info []struct {
		ItemId int64 `json:"item_id"`
		ShopId int64 `json:"shop_id"`
		Models []struct {
			ModelId int `json:"model_id"`
			Basic   struct {
				Name        string `json:"name"`
				Status      int    `json:"status"`
				TierIndices []int  `json:"tier_indices"`
				CreateTime  int    `json:"create_time"`
				SkuType     int    `json:"sku_type"`
				SellerSku   string `json:"seller_sku"`
				IsDefault   bool   `json:"is_default"`
				Region      string `json:"region"`
				Mtime       int    `json:"mtime"`
			} `json:"basic"`
		} `json:"models"`
	} `json:"info"`
}

type GetWeightRecommendationRsp struct {
	BaseRsp
	Response struct {
		NormalWeightRange []float64 `json:"normal_weight_range"`
	} `json:"response"`
}

type GetSizeChartListRsp struct {
	BaseRsp
	Response struct {
		NextCursor    string `json:"next_cursor"`
		SizeChartList []struct {
			SizeChartId int `json:"size_chart_id"`
		} `json:"size_chart_list"`
		TotalCount int `json:"total_count"`
	} `json:"response"`
}

type GetSizeChartDetailRsp struct {
	BaseRsp
	Response struct {
		SizeChartId    int64  `json:"size_chart_id"`
		SizeChartName  string `json:"size_chart_name"`
		SizeChartTable struct {
			ColumnList []struct {
				Measurement struct {
					DisplayName string `json:"display_name"`
					InputType   string `json:"input_type"`
					Unit        string `json:"unit"`
				} `json:"measurement"`
				MeasurementValueList []struct {
					MaxValue int    `json:"max_value"`
					MinValue int    `json:"min_value"`
					Option   string `json:"option"`
					Value    int    `json:"value"`
				} `json:"measurement_value_list"`
			} `json:"column_list"`
		} `json:"size_chart_table"`
	} `json:"response"`
}

type GetItemViolationInfoRsp struct {
	BaseRsp
	Response struct {
		ItemList []struct {
			ItemId            int64  `json:"item_id"`
			ItemName          string `json:"item_name"`
			ItemStatus        string `json:"item_status"`
			Deboost           bool   `json:"deboost"`
			ItemStatusDetails []struct {
				ViolationType   string `json:"violation_type"`
				ViolationReason string `json:"violation_reason"`
				Suggestion      string `json:"suggestion"`
				UpdateTime      int64  `json:"update_time"`
				FixDeadlineTime int64  `json:"fix_deadline_time"`
			} `json:"item_status_details"`
			DeboostDetails []struct {
				ViolationType     string `json:"violation_type"`
				ViolationReason   string `json:"violation_reason"`
				Suggestion        string `json:"suggestion"`
				SuggestedCategory []struct {
					CategoryId   int    `json:"category_id"`
					CategoryName string `json:"category_name"`
				} `json:"suggested_category"`
				UpdateTime      int64 `json:"update_time"`
				FixDeadlineTime int64 `json:"fix_deadline_time"`
			} `json:"deboost_details"`
		} `json:"item_list"`
	} `json:"response"`
}

type GetVariationsRsp struct {
	BaseRsp
	Data struct {
		StandardiseVariationList []struct {
			VariationId        int    `json:"variation_id"`
			VariationName      string `json:"variation_name"`
			VariationGroupList []struct {
				VariationGroupId    int64  `json:"variation_group_id"`
				VariationGroupName  string `json:"variation_group_name"`
				VariationOptionList []struct {
					VariationOptionId   int    `json:"variation_option_id"`
					VariationOptionName string `json:"variation_option_name"`
				} `json:"variation_option_list"`
			} `json:"variation_group_list"`
		} `json:"standardise_variation_list"`
	} `json:"data"`
}

type GetAllVehicleListRsp struct {
	BaseRsp
	Response struct {
		VehicleList []struct {
			BrandId     int    `json:"brand_id"`
			BrandName   string `json:"brand_name"`
			ModelId     int    `json:"model_id"`
			ModelName   string `json:"model_name"`
			YearId      int    `json:"year_id"`
			YearName    string `json:"year_name"`
			VersionId   int    `json:"version_id"`
			VersionName string `json:"version_name"`
		} `json:"vehicle_list"`
		HasNextPage bool `json:"has_next_page"`
		NextOffset  int  `json:"next_offset"`
	} `json:"response"`
}

type GetVehicleListByCompatibilityDetailRsp struct {
	BaseRsp
	Response struct {
		VehicleList []struct {
			BrandId   int    `json:"brand_id"`
			BrandName string `json:"brand_name"`
		} `json:"vehicle_list"`
	} `json:"response"`
}

type GetItemContentDiagnosisResultRsp struct {
	BaseRsp
	Response struct {
		FailureItemList []struct {
			FailedReason string `json:"failed_reason"`
			ItemId       int64  `json:"item_id"`
		} `json:"failure_item_list"`
		SuccessItemList []struct {
			ItemId         int64 `json:"item_id"`
			QualityLevel   int   `json:"quality_level"`
			UnfinishedTask []struct {
				IssueType  int    `json:"issue_type"`
				Suggestion string `json:"suggestion"`
			} `json:"unfinished_task,omitempty"`
		} `json:"success_item_list"`
	} `json:"response"`
	Warning string `json:"warning"`
}

type GetItemListByContentDiagnosisRsp struct {
	BaseRsp
	Response struct {
		ItemList []struct {
			ItemId         int64 `json:"item_id"`
			QualityLevel   int   `json:"quality_level"`
			UnfinishedTask []struct {
				IssueType  int    `json:"issue_type"`
				Suggestion string `json:"suggestion"`
			} `json:"unfinished_task"`
		} `json:"item_list"`
		TotalCount int    `json:"total_count"`
		NextOffset string `json:"next_offset"`
	} `json:"response"`
}

type GetKitItemLimitRsp struct {
	BaseRsp
	Response struct {
		PriceLimit struct {
			MinLimit float64 `json:"min_limit"`
			MaxLimit float64 `json:"max_limit"`
		} `json:"price_limit"`
		ItemNameLengthLimit struct {
			MinLimit int `json:"min_limit"`
			MaxLimit int `json:"max_limit"`
		} `json:"item_name_length_limit"`
		ItemImageCountLimit struct {
			MinLimit int `json:"min_limit"`
			MaxLimit int `json:"max_limit"`
		} `json:"item_image_count_limit"`
		DescriptionLimit struct {
			DescriptionLengthMin           int     `json:"description_length_min"`
			DescriptionLengthMax           int     `json:"description_length_max"`
			DescriptionTextLengthMin       int     `json:"description_text_length_min"`
			DescriptionTextLengthMax       int     `json:"description_text_length_max"`
			DescriptionImageNumMin         int     `json:"description_image_num_min"`
			DescriptionImageNumMax         int     `json:"description_image_num_max"`
			DescriptionImageWidthMin       int     `json:"description_image_width_min"`
			DescriptionImageHeightMin      int     `json:"description_image_height_min"`
			DescriptionImageAspectRatioMin float64 `json:"description_image_aspect_ratio_min"`
			DescriptionImageAspectRatioMax float64 `json:"description_image_aspect_ratio_max"`
		} `json:"description_limit"`
		TierVariationNameLengthLimit struct {
			MinLimit int `json:"min_limit"`
			MaxLimit int `json:"max_limit"`
		} `json:"tier_variation_name_length_limit"`
		TierVariationOptionLengthLimit struct {
			MinLimit int `json:"min_limit"`
			MaxLimit int `json:"max_limit"`
		} `json:"tier_variation_option_length_limit"`
		WeightLimit struct {
			WeightMandatory bool `json:"weight_mandatory"`
		} `json:"weight_limit"`
		DimensionLimit struct {
			DimensionMandatory bool `json:"dimension_mandatory"`
		} `json:"dimension_limit"`
		DtsLimit struct {
			NonPreOrderDaysToShip int  `json:"non_pre_order_days_to_ship"`
			SupportPreOrder       bool `json:"support_pre_order"`
			DaysToShipLimit       struct {
				MinLimit int `json:"min_limit"`
				MaxLimit int `json:"max_limit"`
			} `json:"days_to_ship_limit"`
		} `json:"dts_limit"`
		ComponentCountLimitOfSingleModel struct {
			MinLimit int `json:"min_limit"`
			MaxLimit int `json:"max_limit"`
		} `json:"component_count_limit_of_single_model"`
	} `json:"response"`
}

type AddKitItemRsp struct {
	BaseRsp
	Response struct {
		ItemId int64 `json:"item_id"`
	} `json:"response"`
}

type UpdateKitItemRsp struct {
	BaseRsp
}

type GetKitItemInfoRsp struct {
	BaseRsp
	Response struct {
		ProductInfo struct {
			Attributes []interface{} `json:"attributes"`
			BrandInfo  struct {
				BrandId           int    `json:"brand_id"`
				OriginalBrandName string `json:"original_brand_name"`
			} `json:"brand_info"`
			CategoryId      int `json:"category_id"`
			CreateTime      int `json:"create_time"`
			DescriptionInfo struct {
				ExtendedDescription struct {
					FieldList []struct {
						FieldType string `json:"field_type"`
						Text      string `json:"text"`
					} `json:"field_list"`
				} `json:"extended_description"`
			} `json:"description_info"`
			DescriptionType string `json:"description_type"`
			Dimension       struct {
				PackageHeight int `json:"package_height"`
				PackageLength int `json:"package_length"`
				PackageWidth  int `json:"package_width"`
			} `json:"dimension"`
			Image struct {
				ImageIdList  []string `json:"image_id_list"`
				ImageRatio   string   `json:"image_ratio"`
				ImageUrlList []string `json:"image_url_list"`
			} `json:"image"`
			ItemId       int    `json:"item_id"`
			ItemName     string `json:"item_name"`
			ItemSku      string `json:"item_sku"`
			ItemStatus   string `json:"item_status"`
			LogisticInfo []struct {
				Enabled      bool   `json:"enabled"`
				IsFree       bool   `json:"is_free"`
				LogisticId   int    `json:"logistic_id"`
				LogisticName string `json:"logistic_name"`
				SizeId       int    `json:"size_id"`
			} `json:"logistic_info"`
			ModelList []struct {
				ComponentList []struct {
					ComponentItemId           int    `json:"component_item_id"`
					ComponentItemName         string `json:"component_item_name"`
					ComponentModelId          int64  `json:"component_model_id"`
					ComponentModelName        string `json:"component_model_name"`
					ComponentItemOrModelImage string `json:"component_item_or_model_image"`
					ComponentItemOrModelSku   string `json:"component_item_or_model_sku"`
					MainComponent             bool   `json:"main_component"`
					Quantity                  int    `json:"quantity"`
				} `json:"component_list"`
				ModelId       int64  `json:"model_id"`
				ModelSku      string `json:"model_sku"`
				OriginalPrice int    `json:"original_price"`
				TierIndex     []int  `json:"tier_index"`
			} `json:"model_list"`
			PreOrderInfo struct {
				DaysToShip int  `json:"days_to_ship"`
				IsPreOrder bool `json:"is_pre_order"`
			} `json:"pre_order_info"`
			SyncSetting struct {
				AutoSyncDts bool `json:"auto_sync_dts"`
			} `json:"sync_setting"`
			TierVariationList []struct {
				Name       string `json:"name"`
				OptionList []struct {
					Option string `json:"option"`
				} `json:"option_list"`
			} `json:"tier_variation_list"`
			UpdateTime int           `json:"update_time"`
			VideoList  []interface{} `json:"video_list"`
			Weight     string        `json:"weight"`
		} `json:"product_info"`
	} `json:"response"`
}

type GetSspListRsp struct {
	BaseRsp
	Response struct {
		PageInfo struct {
			HasNextPage bool   `json:"has_next_page"`
			NextOffset  string `json:"next_offset"`
		} `json:"page_info"`
		SspList []struct {
			Available bool `json:"available"`
			SspInfo   struct {
				Gtin        string `json:"gtin"`
				Oem         string `json:"oem"`
				ProductName string `json:"product_name"`
				SspId       int64  `json:"ssp_id"`
			} `json:"ssp_info"`
		} `json:"ssp_list"`
	} `json:"response"`
}

type GetSspInfoRsp struct {
	BaseRsp
	Response struct {
		SspList []struct {
			BrandInfo struct {
				BrandId           int    `json:"brand_id"`
				OriginalBrandName string `json:"original_brand_name"`
			} `json:"brand_info"`
			CategoryId        int `json:"category_id"`
			CompatibilityInfo struct {
			} `json:"compatibility_info"`
			Description struct {
				Description         string `json:"description"`
				DescriptionType     string `json:"description_type"`
				ExtendedDescription struct {
					FieldList []struct {
						Text string `json:"text"`
						Type string `json:"type"`
					} `json:"field_list"`
				} `json:"extended_description"`
			} `json:"description"`
			Dimension struct {
				PackageHeight string `json:"package_height"`
				PackageLength string `json:"package_length"`
				PackageWidth  string `json:"package_width"`
			} `json:"dimension"`
			GlobalCode struct {
				Gtin string `json:"gtin"`
				Oem  string `json:"oem"`
			} `json:"global_code"`
			Media struct {
				Image struct {
					ImageIdList  []string `json:"image_id_list"`
					ImageRatio   string   `json:"image_ratio"`
					ImageUrlList []string `json:"image_url_list"`
				} `json:"image"`
			} `json:"media"`
			ProductName string `json:"product_name"`
			SalesInfo   struct {
				ModelSettingList []struct {
					Gtin      string `json:"gtin"`
					TierIndex []int  `json:"tier_index"`
				} `json:"model_setting_list"`
				StdTierVariationList []struct {
					VariationId         int    `json:"variation_id"`
					VariationName       string `json:"variation_name"`
					VariationGroupId    int    `json:"variation_group_id"`
					VariationOptionList []struct {
						VariationOptionId   int    `json:"variation_option_id"`
						VariationOptionName string `json:"variation_option_name"`
					} `json:"variation_option_list"`
				} `json:"std_tier_variation_list"`
			} `json:"sales_info"`
			SspId  int64  `json:"ssp_id"`
			Weight string `json:"weight"`
		} `json:"ssp_list"`
	} `json:"response"`
}

type AddSspItemRsp struct {
	BaseRsp
	Response struct {
		ItemId int64 `json:"item_id"`
	} `json:"response"`
}

type LinkSspRsp struct {
	BaseRsp
	Response struct {
		ItemId int64 `json:"item_id"`
	} `json:"response"`
}

type UnlinkSspRsp struct {
	BaseRsp
	Response struct {
		ItemId int64 `json:"item_id"`
	} `json:"response"`
}

type GetAitemByPitemIdRsp struct {
	BaseRsp
	Response struct {
		AitemList []struct {
			AitemId          int64  `json:"aitem_id"`
			AshopId          int64  `json:"ashop_id"`
			AshopRegion      string `json:"ashop_region"`
			ModelMappingList []struct {
				AmodelId int `json:"amodel_id"`
				PmodelId int `json:"pmodel_id"`
			} `json:"model_mapping_list"`
		} `json:"aitem_list"`
	} `json:"response"`
}
