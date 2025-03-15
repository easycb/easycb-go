package shopee

type GlobalGetCategoryRsp struct {
	BaseRsp
	Response struct {
		CategoryList []struct {
			CategoryId           int    `json:"category_id"`
			ParentCategoryId     int    `json:"parent_category_id"`
			OriginalCategoryName string `json:"original_category_name"`
			DisplayCategoryName  string `json:"display_category_name"`
			HasChildren          bool   `json:"has_children"`
		} `json:"category_list"`
	} `json:"response"`
}

type GlobalGetAttributeTreeRsp struct {
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
							InputType           int      `json:"input_type"`
							InputValidationType int      `json:"input_validation_type"`
							FormatType          int      `json:"format_type"`
							MandatoryRegion     []string `json:"mandatory_region"`
						} `json:"attribute_info"`
					} `json:"child_attribute_list,omitempty"`
				} `json:"attribute_value_list"`
				AttributeInfo struct {
					InputType           int      `json:"input_type"`
					InputValidationType int      `json:"input_validation_type"`
					FormatType          int      `json:"format_type"`
					MandatoryRegion     []string `json:"mandatory_region"`
				} `json:"attribute_info"`
			} `json:"attribute_tree"`
		} `json:"list"`
	} `json:"response"`
}

type GlobalGetBrandListRsp struct {
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

type GetGlobalItemLimitRsp struct {
	BaseRsp
	Response struct {
		PriceLimit struct {
			MinLimit int `json:"min_limit"`
			MaxLimit int `json:"max_limit"`
		} `json:"price_limit"`
		StockLimit struct {
			MinLimit int `json:"min_limit"`
			MaxLimit int `json:"max_limit"`
		} `json:"stock_limit"`
		GlobalItemNameLengthLimit struct {
			MinLimit int `json:"min_limit"`
			MaxLimit int `json:"max_limit"`
		} `json:"global_item_name_length_limit"`
		GlobalItemImageCountLimit struct {
			MinLimit int `json:"min_limit"`
			MaxLimit int `json:"max_limit"`
		} `json:"global_item_image_count_limit"`
		GlobalItemDescriptionLengthLimit struct {
			MinLimit int `json:"min_limit"`
			MaxLimit int `json:"max_limit"`
		} `json:"global_item_description_length_limit"`
		TierVariationNameLengthLimit struct {
			MinLimit int `json:"min_limit"`
			MaxLimit int `json:"max_limit"`
		} `json:"tier_variation_name_length_limit"`
		TierVariationOptionLengthLimit struct {
			MinLimit int `json:"min_limit"`
			MaxLimit int `json:"max_limit"`
		} `json:"tier_variation_option_length_limit"`
		TextLengthMultiplier     float64 `json:"text_length_multiplier"`
		ExtendedDescriptionLimit struct {
			DescriptionTextLengthMin       int     `json:"description_text_length_min"`
			DescriptionTextLengthMax       int     `json:"description_text_length_max"`
			DescriptionImageNumMin         int     `json:"description_image_num_min"`
			DescriptionImageNumMax         int     `json:"description_image_num_max"`
			DescriptionImageWidthMin       int     `json:"description_image_width_min"`
			DescriptionImageHeightMin      int     `json:"description_image_height_min"`
			DescriptionImageAspectRatioMin float64 `json:"description_image_aspect_ratio_min"`
			DescriptionImageAspectRatioMax int     `json:"description_image_aspect_ratio_max"`
		} `json:"extended_description_limit"`
		DtsLimit struct {
			DaysToShipRangeList []struct {
				MinLimit int `json:"min_limit"`
				MaxLimit int `json:"max_limit"`
			} `json:"days_to_ship_range_list"`
		} `json:"dts_limit"`
		WeightLimit struct {
			WeightMandatory bool `json:"weight_mandatory"`
		} `json:"weight_limit"`
		DimensionLimit struct {
			DimensionMandatory bool `json:"dimension_mandatory"`
		} `json:"dimension_limit"`
	} `json:"response"`
}

type GetGlobalItemListRsp struct {
	BaseRsp
	Response struct {
		GlobalItemList []struct {
			GlobalItemId int64 `json:"global_item_id"`
			UpdateTime   int   `json:"update_time"`
		} `json:"global_item_list"`
		TotalCount  int    `json:"total_count"`
		HasNextPage bool   `json:"has_next_page"`
		Offset      string `json:"offset"`
	} `json:"response"`
}

type GetGlobalItemInfoRsp struct {
	BaseRsp
	Response struct {
		GlobalItemList []struct {
			GlobalItemId     int64  `json:"global_item_id"`
			GlobalItemName   string `json:"global_item_name"`
			Description      string `json:"description"`
			GlobalItemSku    string `json:"global_item_sku"`
			GlobalItemStatus string `json:"global_item_status"`
			CreateTime       int64  `json:"create_time"`
			UpdateTime       int64  `json:"update_time"`
			Image            struct {
				ImageIdList  []string `json:"image_id_list"`
				ImageUrlList []string `json:"image_url_list"`
			} `json:"image"`
			Weight    float64 `json:"weight"`
			Dimension struct {
				PackageLength float64 `json:"package_length"`
				PackageWidth  float64 `json:"package_width"`
				PackageHeight float64 `json:"package_height"`
			} `json:"dimension"`
			PreOrder struct {
				DaysToShip int `json:"days_to_ship"`
			} `json:"pre_order"`
			SizeChart string `json:"size_chart"`
			Condition string `json:"condition"`
			HasModel  bool   `json:"has_model"`
			StockInfo []struct {
				StockType       int    `json:"stock_type"`
				StockLocationId string `json:"stock_location_id"`
				NormalStock     int    `json:"normal_stock"`
				ReservedStock   int    `json:"reserved_stock"`
			} `json:"stock_info,omitempty"`
			PriceInfo []struct {
				Currency           string  `json:"currency"`
				OriginalPrice      float64 `json:"original_price"`
				SipItemPrice       float64 `json:"sip_item_price"`
				SipItemPriceSource string  `json:"sip_item_price_source"`
			} `json:"price_info,omitempty"`
			Brand struct {
				BrandId           int    `json:"brand_id"`
				OriginalBrandName string `json:"original_brand_name"`
			} `json:"brand,omitempty"`
			AttributeList []struct {
				AttributeId           int    `json:"attribute_id"`
				OriginalAttributeName string `json:"original_attribute_name"`
				IsMandatory           bool   `json:"is_mandatory"`
				AttributeValueList    []struct {
					ValueId           int    `json:"value_id"`
					OriginalValueName string `json:"original_value_name"`
					ValueUnit         string `json:"value_unit"`
				} `json:"attribute_value_list"`
			} `json:"attribute_list,omitempty"`
			DescriptionType string `json:"description_type,omitempty"`
			DescriptionInfo struct {
				ExtendedDescription struct {
					FieldList []struct {
						FieldType string `json:"field_type"`
						Text      string `json:"text,omitempty"`
						ImageInfo struct {
							ImageId  string `json:"image_id"`
							ImageUrl string `json:"image_url"`
						} `json:"image_info,omitempty"`
					} `json:"field_list"`
				} `json:"extended_description"`
			} `json:"description_info,omitempty"`
		} `json:"global_item_list"`
	} `json:"response"`
}

type AddGlobalItemRsp struct {
	BaseRsp
	Response struct {
		GlobalItemId int64 `json:"global_item_id"`
	} `json:"response"`
}

type UpdateGlobalItemRsp struct {
	BaseRsp
	Response struct {
		GlobalItemId int64 `json:"global_item_id"`
	} `json:"response"`
}

type DeleteGlobalItemRsp struct {
	BaseRsp
	Response struct {
		FailureDeleteItem []struct {
			ShopId int64 `json:"shop_id"`
			ItemId int64 `json:"item_id"`
		} `json:"failure_delete_item"`
	} `json:"response"`
}

type GlobalInitTierVariationRsp struct {
	BaseRsp
}

type GlobalUpdateTierVariationRsp struct {
	BaseRsp
}

type AddGlobalModelRsp struct {
	BaseRsp
}

type UpdateGlobalModelRsp struct {
	BaseRsp
}

type DeleteGlobalModelRsp struct {
	BaseRsp
	Response struct {
		GlobalModelId int64         `json:"global_model_id"`
		Failures      []interface{} `json:"failures"`
	} `json:"response"`
}

type GetGlobalModelListRsp struct {
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
		GlobalModel []struct {
			GlobalModelId  int64  `json:"global_model_id"`
			GlobalModelSku string `json:"global_model_sku"`
			PriceInfo      struct {
				OriginalPrice float64 `json:"original_price"`
				Currency      string  `json:"currency"`
			} `json:"price_info"`
			StockInfo []struct {
				StockType       int    `json:"stock_type"`
				StockLocationId string `json:"stock_location_id"`
				CurrentStock    int    `json:"current_stock"`
				NormalStock     int    `json:"normal_stock"`
				ReservedStock   int    `json:"reserved_stock"`
			} `json:"stock_info"`
			TierIndex []int  `json:"tier_index"`
			Weight    string `json:"weight"`
			Dimension struct {
				PackageLength int `json:"package_length"`
				PackageWidth  int `json:"package_width"`
				PackageHeight int `json:"package_height"`
			} `json:"dimension"`
			PreOrder struct {
				DaysToShip int `json:"days_to_ship"`
			} `json:"pre_order"`
			IsFulfillmentByShopee bool `json:"is_fulfillment_by_shopee"`
		} `json:"global_model"`
		StandardiseTierVariation []struct {
			VariationId         int    `json:"variation_id"`
			VariationName       string `json:"variation_name"`
			VariationOptionList []struct {
				VariationOptionId   int    `json:"variation_option_id"`
				VariationOptionName string `json:"variation_option_name"`
				ImageId             string `json:"image_id,omitempty"`
				ImageUrl            string `json:"image_url,omitempty"`
			} `json:"variation_option_list"`
		} `json:"standardise_tier_variation"`
	} `json:"response"`
}

type GlobalSupportSizeChartRsp struct {
	BaseRsp
	Response struct {
		SupportSizeChart bool `json:"support_size_chart"`
	} `json:"response"`
}

type GlobalUpdateSizeChartRsp struct {
	BaseRsp
}

type CreatePublishTaskRsp struct {
	BaseRsp
	Response struct {
		PublishTaskId int64 `json:"publish_task_id"`
	} `json:"response"`
}

type GetPublishableShopRsp struct {
	Response struct {
		PublishableShop []struct {
			ShopId     int64  `json:"shop_id"`
			ShopRegion string `json:"shop_region"`
		} `json:"publishable_shop"`
	} `json:"response"`
}

type GetPublishTaskResultRsp struct {
	BaseRsp
	Response struct {
		PublishStatus string `json:"publish_status"`
		Success       struct {
			Region string `json:"region"`
			ShopId int64  `json:"shop_id"`
			ItemId int64  `json:"item_id"`
		} `json:"success"`
	} `json:"response"`
}

type GetPublishedListRsp struct {
	BaseRsp
	Response struct {
		PublishedItem []struct {
			ShopId     int64  `json:"shop_id"`
			ShopRegion string `json:"shop_region"`
			ItemId     int64  `json:"item_id"`
			ItemStatus int    `json:"item_status"`
		} `json:"published_item"`
	} `json:"response"`
}

type GlobalUpdatePriceRsp struct {
	BaseRsp
}

type GlobalUpdateStockRsp struct {
	BaseRsp
}

type GlobalSetSyncFieldRsp struct {
	BaseRsp
}

type GetGlobalItemIdRsp struct {
	BaseRsp
	Response struct {
		ItemIdMap []struct {
			ItemId       int64 `json:"item_id"`
			GlobalItemId int64 `json:"global_item_id"`
		} `json:"item_id_map"`
	} `json:"response"`
}

type GlobalCategoryRecommendRsp struct {
	BaseRsp
	Response struct {
		CategoryId []int `json:"category_id"`
	} `json:"response"`
}

type GlobalGetRecommendAttributeRsp struct {
	BaseRsp
	Response struct {
		AttributeList []struct {
			AttributeId        int `json:"attribute_id"`
			AttributeValueList []struct {
				ValueId int `json:"value_id"`
			} `json:"attribute_value_list"`
		} `json:"attribute_list"`
	} `json:"response"`
}

type GlobalGetShopPublishableStatusRsp struct {
	BaseRsp
	Response struct {
		HasNextPage               bool `json:"has_next_page"`
		NextOffset                int  `json:"next_offset"`
		ShopPublishableStatusList []struct {
			Region                string `json:"region"`
			ShopId                int64  `json:"shop_id"`
			ShopPublishableStatus bool   `json:"shop_publishable_status"`
			UnpublishableReason   string `json:"unpublishable_reason"`
		} `json:"shop_publishable_status_list"`
	} `json:"response"`
}

type GlobalGetVariationRsp struct {
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

type GlobalGetSizeChartDetailRsp struct {
	BaseRsp
	Response struct {
		SizeChartId    int    `json:"size_chart_id"`
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

type GlobalGetSizeChartListRsp struct {
	BaseRsp
	Response struct {
		NextOffset    string `json:"next_offset"`
		SizeChartList []struct {
			SizeChartId int `json:"size_chart_id"`
		} `json:"size_chart_list"`
		TotalCount int `json:"total_count"`
	} `json:"response"`
}
