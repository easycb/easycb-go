package lazada

type AdjustSellableQuantityRsp struct {
	Code string `json:"code"`
	Data struct {
	} `json:"data"`
	RequestId string `json:"request_id"`
}

type BatchUpdateSizeChartRsp struct {
	Code string `json:"code"`
	Data struct {
	} `json:"data"`
	RequestId string `json:"request_id"`
}

type CreateProductRsp struct {
	Code string `json:"code"`
	Data struct {
		ItemId  int64 `json:"item_id"`
		SkuList []struct {
			ShopSku   string `json:"shop_sku"`
			SellerSku string `json:"seller_sku"`
			SkuId     int    `json:"sku_id"`
		} `json:"sku_list"`
	} `json:"data"`
	RequestId string `json:"request_id"`
}

type DeactivateProductRsp struct {
	Code string `json:"code"`
	Data struct {
	} `json:"data"`
	RequestId string `json:"request_id"`
}

type ExitExperimentRsp struct {
	Code      string `json:"code"`
	Data      bool   `json:"data"`
	Success   bool   `json:"success"`
	ErrorCode string `json:"errorCode"`
	RequestId string `json:"request_id"`
	ErrorMsg  string `json:"errorMsg"`
}

type GetBrandByPagesRsp struct {
	ErrorMsg string `json:"error_msg"`
	Code     string `json:"code"`
	Data     struct {
		EnableTotal bool `json:"enable_total"`
		StartRow    int  `json:"start_row"`
		PageIndex   int  `json:"page_index"`
		Module      []struct {
			Name             string `json:"name"`
			GlobalIdentifier string `json:"global_identifier"`
			NameEn           string `json:"name_en"`
			BrandId          int    `json:"brand_id"`
		} `json:"module"`
		TotalPage   int `json:"total_page"`
		PageSize    int `json:"page_size"`
		TotalRecord int `json:"total_record"`
	} `json:"data"`
	Success   bool   `json:"success"`
	ErrorCode string `json:"error_code"`
	RequestId string `json:"request_id"`
}

type GetCategoryAttributesRsp struct {
	Code string `json:"code"`
	Data []struct {
		Advanced struct {
			IsKeyProp int `json:"is_key_prop"`
		} `json:"advanced"`
		IsSaleProp int    `json:"is_sale_prop"`
		Name       string `json:"name"`
		InputType  string `json:"input_type"`
		Options    []struct {
			Name   string `json:"name"`
			EnName string `json:"en_name"`
			Id     int    `json:"id"`
		} `json:"options"`
		IsMandatory   int    `json:"is_mandatory"`
		AttributeType string `json:"attribute_type"`
		Label         string `json:"label"`
		Id            int    `json:"id"`
	} `json:"data"`
	RequestId string `json:"request_id"`
}

type GetCategorySuggestionRsp struct {
	Code string `json:"code"`
	Data struct {
		CategorySuggestions []struct {
			CategoryPath string `json:"categoryPath"`
			CategoryName string `json:"categoryName"`
			CategoryId   int    `json:"categoryId"`
		} `json:"categorySuggestions"`
	} `json:"data"`
	RequestId string `json:"request_id"`
}

type GetCategorySuggestionInBulkRsp struct {
	Code string `json:"code"`
	Data struct {
		CategorySuggestionMap struct {
			TestProduct9 []struct {
				CategoryPath string `json:"categoryPath"`
				Class        string `json:"class"`
				CategoryName string `json:"categoryName"`
				CategoryId   int    `json:"categoryId"`
			} `json:"Test Product9"`
		} `json:"categorySuggestionMap"`
	} `json:"data"`
	RequestId string `json:"request_id"`
}

type GetCategoryTreeRsp struct {
	Code string `json:"code"`
	Data []struct {
		CategoryId int `json:"category_id"`
		Children   []struct {
			CategoryId int    `json:"category_id"`
			Var        bool   `json:"var"`
			Name       string `json:"name"`
			Leaf       bool   `json:"leaf"`
		} `json:"children"`
		Var  bool   `json:"var"`
		Name string `json:"name"`
		Leaf bool   `json:"leaf"`
	} `json:"data"`
	RequestId string `json:"request_id"`
}

type GetImageExperimentDataRsp struct {
	Code string `json:"code"`
	Data struct {
		ProductId int64 `json:"productId"`
		GroupData struct {
			Field1 struct {
				Data struct {
				} `json:"data"`
			} `json:"276518"`
			Field2 struct {
				BucketType string `json:"bucketType"`
				Data       struct {
					OrderUv    string `json:"order_uv"`
					OrderCnt   string `json:"order_cnt"`
					ExpTotalUv string `json:"exp_total_uv"`
					IpvPv      string `json:"ipv_pv"`
					IpvUv      string `json:"ipv_uv"`
				} `json:"data"`
				Sampling int `json:"sampling"`
				BucketId int `json:"bucketId"`
			} `json:"276519"`
		} `json:"groupData"`
		Venture string `json:"venture"`
	} `json:"data"`
	Success   bool   `json:"success"`
	ErrorCode string `json:"errorCode"`
	RequestId string `json:"request_id"`
	ErrorMsg  string `json:"errorMsg"`
}

type GetNextCascadePropRsp struct {
	Code string `json:"code"`
	Data struct {
		Prop struct {
			Name     string `json:"name"`
			Id       int    `json:"id"`
			Required bool   `json:"required"`
		} `json:"prop"`
		PropValue []struct {
			Name string `json:"name"`
			Id   int    `json:"id"`
			Leaf string `json:"leaf"`
		} `json:"propValue"`
	} `json:"data"`
	RequestId string `json:"request_id"`
}

type GetPreQcRulesRsp struct {
	Code   string `json:"code"`
	Values struct {
		ItemLimit         int   `json:"item_limit"`
		ItemCount         int   `json:"item_count"`
		RestrictedCateIds []int `json:"restricted_cate_ids"`
	} `json:"values"`
	RequestId string `json:"request_id"`
}

type GetProductContentScoreRsp struct {
	Result struct {
		Data struct {
			ProductTitle string `json:"productTitle"`
			Score        int    `json:"score"`
			Image        string `json:"image"`
			Total        int    `json:"total"`
			ProductId    int64  `json:"productId"`
			Items        []struct {
				Score      int    `json:"score"`
				Total      int    `json:"total"`
				ItemTitle  string `json:"itemTitle"`
				Label      string `json:"label"`
				Indicators []struct {
					Critical bool   `json:"critical"`
					Text     string `json:"text"`
					Key      string `json:"key"`
				} `json:"indicators"`
				ImageList []struct {
					Score      int    `json:"score"`
					ImageUrl   string `json:"imageUrl"`
					Text       string `json:"text"`
					Type       string `json:"type"`
					Indicators []struct {
						Text string `json:"text"`
						Key  string `json:"key"`
					} `json:"indicators"`
					ImageType int `json:"imageType"`
				} `json:"imageList"`
				Key    string `json:"key"`
				Group  string `json:"group"`
				Latest bool   `json:"latest"`
			} `json:"items"`
		} `json:"data"`
	} `json:"result"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type GetProductItemRsp struct {
	Data struct {
		CreatedTime string   `json:"created_time"`
		UpdatedTime string   `json:"updated_time"`
		Images      []string `json:"images"`
		Skus        []struct {
			SellerSku                 string                 `json:"SellerSku"`
			ShopSku                   string                 `json:"ShopSku"`
			Url                       string                 `json:"Url"`
			ColorFamily               string                 `json:"color_family"`
			PackageHeight             string                 `json:"package_height"`
			FblWarehouseInventories   []interface{}          `json:"fblWarehouseInventories"`
			Price                     float64                `json:"price"`
			PackageLength             string                 `json:"package_length"`
			SpecialFromDate           string                 `json:"special_from_date"`
			Available                 int                    `json:"Available"`
			SpecialToDate             string                 `json:"special_to_date"`
			Status                    string                 `json:"Status"`
			Quantity                  int                    `json:"quantity"`
			Images                    []string               `json:"Images"`
			SpecialTimeFormat         string                 `json:"special_time_format"`
			SaleProp                  map[string]interface{} `json:"saleProp"`
			MultiWarehouseInventories []struct {
				OccupyQuantity   int    `json:"occupyQuantity"`
				Quantity         int    `json:"quantity"`
				TotalQuantity    int    `json:"totalQuantity"`
				WithholdQuantity int    `json:"withholdQuantity"`
				WarehouseCode    string `json:"warehouseCode"`
				SellableQuantity int    `json:"sellableQuantity"`
			} `json:"multiWarehouseInventories"`
			PackageWidth       string        `json:"package_width"`
			SpecialToTime      string        `json:"special_to_time"`
			SpecialFromTime    string        `json:"special_from_time"`
			Size               string        `json:"size"`
			SpecialPrice       float64       `json:"special_price"`
			ChannelInventories []interface{} `json:"channelInventories"`
			PackageWeight      string        `json:"package_weight"`
			SkuId              int64         `json:"SkuId"`
		} `json:"skus"`
		ItemId    int64 `json:"item_id"`
		Variation struct {
			Variation1 struct {
				Name      string   `json:"name"`
				Label     string   `json:"label"`
				HasImage  bool     `json:"hasImage"`
				Customize bool     `json:"customize"`
				Options   []string `json:"options"`
			} `json:"Variation1"`
			Variation2 struct {
				Name      string   `json:"name"`
				Label     string   `json:"label"`
				HasImage  bool     `json:"hasImage"`
				Customize bool     `json:"customize"`
				Options   []string `json:"options"`
			} `json:"Variation2"`
		} `json:"variation"`
		TrialProduct    bool          `json:"trialProduct"`
		PrimaryCategory int           `json:"primary_category"`
		MarketImages    []interface{} `json:"marketImages"`
		Attributes      struct {
			Description      string `json:"description"`
			SkirtLength      string `json:"skirt_length"`
			SkirtStyle       string `json:"skirt_style"`
			ClothingMaterial string `json:"clothing_material"`
			WarrantyType     string `json:"warranty_type"`
			DescriptionMs    string `json:"description_ms"`
			NameMs           string `json:"name_ms"`
			Name             string `json:"name"`
			Source           string `json:"source"`
		} `json:"attributes"`
		Status string `json:"status"`
	} `json:"data"`
	Code      string `json:"code"`
	Message   string `json:"message"`
	RequestId string `json:"request_id"`
}

type GetProductsRsp struct {
	Data struct {
		TotalProducts int `json:"total_products"`
		Products      []struct {
			CreatedTime string   `json:"created_time"`
			UpdatedTime string   `json:"updated_time"`
			Images      []string `json:"images"`
			Skus        []struct {
				SellerSku                 string                 `json:"SellerSku"`
				ShopSku                   string                 `json:"ShopSku"`
				Url                       string                 `json:"Url"`
				ColorFamily               string                 `json:"color_family,omitempty"`
				PackageHeight             string                 `json:"package_height"`
				FblWarehouseInventories   []interface{}          `json:"fblWarehouseInventories"`
				Price                     float64                `json:"price"`
				PackageLength             string                 `json:"package_length"`
				SpecialFromDate           string                 `json:"special_from_date,omitempty"`
				Available                 int                    `json:"Available"`
				SpecialToDate             string                 `json:"special_to_date,omitempty"`
				Status                    string                 `json:"Status"`
				Quantity                  int                    `json:"quantity"`
				Images                    []string               `json:"Images"`
				SpecialTimeFormat         string                 `json:"special_time_format,omitempty"`
				SaleProp                  map[string]interface{} `json:"saleProp"`
				MultiWarehouseInventories []struct {
					OccupyQuantity   int    `json:"occupyQuantity"`
					Quantity         int    `json:"quantity"`
					TotalQuantity    int    `json:"totalQuantity"`
					WithholdQuantity int    `json:"withholdQuantity"`
					WarehouseCode    string `json:"warehouseCode"`
					SellableQuantity int    `json:"sellableQuantity"`
				} `json:"multiWarehouseInventories"`
				PackageWidth       string        `json:"package_width"`
				SpecialToTime      string        `json:"special_to_time,omitempty"`
				SpecialFromTime    string        `json:"special_from_time,omitempty"`
				Size               string        `json:"size,omitempty"`
				SpecialPrice       float64       `json:"special_price"`
				ChannelInventories []interface{} `json:"channelInventories"`
				PackageWeight      string        `json:"package_weight"`
				SkuId              int64         `json:"SkuId"`
				Size1              string        `json:"Size,omitempty"`
				Color              string        `json:"Color,omitempty"`
				ColorClassifi      string        `json:"color classifi,omitempty"`
				ReferenceClas      string        `json:"Reference clas,omitempty"`
				Material           string        `json:"Material,omitempty"`
				Model              string        `json:"Model,omitempty"`
				Variation2Shoe     string        `json:"Variation2Shoe,omitempty"`
				ColorClassifi1     string        `json:"Color Classifi,omitempty"`
				ChickenDrinker     string        `json:"ChickenDrinker,omitempty"`
				Combo              string        `json:"combo,omitempty"`
				ColorClassifi2     string        `json:"Color classifi,omitempty"`
				Color1             string        `json:"color,omitempty"`
				Rozmiar            string        `json:"Rozmiar,omitempty"`
				Kolor              string        `json:"Kolor,omitempty"`
				SortByColor        string        `json:"sort by color,omitempty"`
				Ukuran             string        `json:"Ukuran,omitempty"`
				Colour             string        `json:"Colour,omitempty"`
				Length             string        `json:"length,omitempty"`
				Warna              string        `json:"Warna,omitempty"`
			} `json:"skus"`
			ItemId          int64         `json:"item_id"`
			TrialProduct    bool          `json:"trialProduct"`
			PrimaryCategory int           `json:"primary_category"`
			MarketImages    []interface{} `json:"marketImages"`
			Attributes      struct {
				Description                    string `json:"description"`
				SkirtLength                    string `json:"skirt_length,omitempty"`
				SkirtStyle                     string `json:"skirt_style,omitempty"`
				ClothingMaterial               string `json:"clothing_material,omitempty"`
				WarrantyType                   string `json:"warranty_type"`
				DescriptionMs                  string `json:"description_ms,omitempty"`
				NameMs                         string `json:"name_ms,omitempty"`
				Name                           string `json:"name"`
				Source                         string `json:"source"`
				Brand                          string `json:"brand,omitempty"`
				DressLength                    string `json:"dress_length,omitempty"`
				FaPattern                      string `json:"fa_pattern,omitempty"`
				DressShape                     string `json:"dress_shape,omitempty"`
				HairAccessories                string `json:"hair_accessories,omitempty"`
				RecommendedGender              string `json:"recommended_gender,omitempty"`
				RecommendedAge                 string `json:"recommended_age,omitempty"`
				MaterialType                   string `json:"material_type,omitempty"`
				CollarType                     string `json:"collar_type,omitempty"`
				Sleeves                        string `json:"sleeves,omitempty"`
				TypeOfFasteningGluingSoldering string `json:"Type_Of_Fastening,_Gluing_&_Soldering,omitempty"`
				SockTightStyle                 string `json:"sock_tight_style,omitempty"`
				BrandCompatibility             string `json:"brand_compatibility,omitempty"`
				CaseFunction                   string `json:"case_function,omitempty"`
				HatsStyle                      string `json:"hats_style,omitempty"`
				SportsMaterial                 string `json:"Sports_Material,omitempty"`
				MTopNeckline                   string `json:"m_top_neckline,omitempty"`
				BodyJewellery                  string `json:"body_jewellery,omitempty"`
				PetSuppliesMaterial            string `json:"Pet_Supplies_Material,omitempty"`
				GlasswareMaterial              string `json:"glassware_material,omitempty"`
				Language                       string `json:"language,omitempty"`
				IsbnIssn                       string `json:"isbn_issn,omitempty"`
				ShoeAccessoryType              string `json:"shoe_accessory_type,omitempty"`
				CookwareMaterial               string `json:"cookware_material,omitempty"`
				Material                       string `json:"material,omitempty"`
				TopsType                       string `json:"tops_type,omitempty"`
				WBlouseSleevestyle             string `json:"w_blouse_sleevestyle,omitempty"`
				SleepLoungeStyles              string `json:"sleep_lounge_styles,omitempty"`
				FishingType                    string `json:"fishing_type,omitempty"`
				TopDiameter                    string `json:"Top_Diameter,omitempty"`
				LengthM                        string `json:"Length_(m),omitempty"`
				TypeOfElectricalCircuitParts   string `json:"Type_Of_Electrical_Circuit_Parts,omitempty"`
				TypeTools                      string `json:"type_tools,omitempty"`
				SkinType                       string `json:"skin_type,omitempty"`
				SkinConcern                    string `json:"skin_concern,omitempty"`
			} `json:"attributes"`
			Status string `json:"status"`
		} `json:"products"`
	} `json:"data"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type GetQCAlertProductsRsp struct {
	Code string `json:"code"`
	Data []struct {
		ProductId            int64 `json:"productId"`
		SuggestionCategories []int `json:"suggestionCategories"`
		CategoryId           int   `json:"categoryId"`
		DeactivationTime     int64 `json:"deactivationTime"`
	} `json:"data"`
	RequestId string `json:"request_id"`
}

type GetQcStatusRsp struct {
	Code string `json:"code"`
	Data []struct {
		Reason      string `json:"reason"`
		DataChanged string `json:"data_changed"`
		SellerSku   string `json:"seller_sku"`
		Status      string `json:"status"`
	} `json:"data"`
	RequestId string `json:"request_id"`
}

type GetResponseRsp struct {
	Code string `json:"code"`
	Data struct {
		Images []struct {
			HashCode string `json:"hash_code"`
			Url      string `json:"url"`
		} `json:"images"`
		Errors []struct {
			Msg         string `json:"msg"`
			Field       string `json:"field"`
			OriginalUrl string `json:"original_url"`
		} `json:"errors"`
	} `json:"data"`
	RequestId string `json:"request_id"`
}

type GetSellerItemLimitRsp struct {
	Code string `json:"code"`
	Data struct {
		PayByrCnt       int `json:"payByrCnt"`
		PayItemCnt      int `json:"payItemCnt"`
		ItemLimit       int `json:"itemLimit"`
		OnlineItemCount int `json:"onlineItemCount"`
	} `json:"data"`
	Success    bool          `json:"success"`
	ErrorCodes []interface{} `json:"errorCodes"`
	ErrorMsgs  []interface{} `json:"errorMsgs"`
	RequestId  string        `json:"request_id"`
}

type GetSizeChartTemplateRsp struct {
	Code string `json:"code"`
	Data struct {
		Total              int           `json:"total"`
		PageNo             int           `json:"pageNo"`
		TotalPage          int           `json:"totalPage"`
		PageSize           int           `json:"pageSize"`
		SizeChartResponses []interface{} `json:"sizeChartResponses"`
	} `json:"data"`
	RequestId string `json:"request_id"`
}

type GetUnfilledAttributeItemRsp struct {
	ErrorMsg      string `json:"error_msg"`
	Code          string `json:"code"`
	Success       bool   `json:"success"`
	TotalProducts int    `json:"total_products"`
	RequestId     string `json:"request_id"`
	Products      []struct {
		ItemId          int64 `json:"item_id"`
		PrimaryCategory int   `json:"primary_category"`
		Attributes      []struct {
			Advanced struct {
				IsKeyProp int `json:"is_key_prop"`
			} `json:"advanced"`
			Name      string `json:"name"`
			InputType string `json:"input_type"`
			Options   []struct {
				Name string `json:"name"`
			} `json:"options"`
			IsMandatory   int    `json:"is_mandatory"`
			AttributeType string `json:"attribute_type"`
			Label         string `json:"label"`
		} `json:"attributes"`
		SellerSkuId string `json:"seller_sku_id"`
	} `json:"products"`
}

type JoinExperimentRsp struct {
	Code      string `json:"code"`
	Data      bool   `json:"data"`
	Success   bool   `json:"success"`
	ErrorCode string `json:"errorCode"`
	RequestId string `json:"request_id"`
	ErrorMsg  string `json:"errorMsg"`
}

type MigrateImageRsp struct {
	Code string `json:"code"`
	Data struct {
		Image struct {
			HashCode string `json:"hash_code"`
			Url      string `json:"url"`
		} `json:"image"`
	} `json:"data"`
	RequestId string `json:"request_id"`
}

type MigrateImagesRsp struct {
	Code      string `json:"code"`
	BatchId   string `json:"batch_id"`
	RequestId string `json:"request_id"`
}

type ProductCheckRsp struct {
	Code string `json:"code"`
	Data struct {
	} `json:"data"`
	RequestId string `json:"request_id"`
}

type QueryProductExperimentConfigurationRsp struct {
	Code string `json:"code"`
	Data struct {
		ProductId          int64  `json:"productId"`
		EndDate            string `json:"endDate"`
		MainImageForAbTest string `json:"mainImageForAbTest"`
		Venture            string `json:"venture"`
		StartDate          string `json:"startDate"`
		ExperimentEnabled  bool   `json:"experimentEnabled"`
	} `json:"data"`
	Success   bool   `json:"success"`
	ErrorCode string `json:"errorCode"`
	RequestId string `json:"request_id"`
	ErrorMsg  string `json:"errorMsg"`
}

type RemoveProductRsp struct {
	Code string `json:"code"`
	Data struct {
	} `json:"data"`
	RequestId string `json:"request_id"`
}

type RemoveSkuRsp struct {
	Code string `json:"code"`
	Data struct {
	} `json:"data"`
	RequestId string `json:"request_id"`
}

type SetImagesRsp struct {
	Code string `json:"code"`
	Data struct {
	} `json:"data"`
	RequestId string `json:"request_id"`
}

type UpdatePriceQuantityRsp struct {
	Code string `json:"code"`
	Data struct {
	} `json:"data"`
	RequestId string `json:"request_id"`
}

type UpdateProductRsp struct {
	Code string `json:"code"`
	Data struct {
		Variation struct {
			Variation1 struct {
				HasImage  bool     `json:"has_image"`
				Name      string   `json:"name"`
				Options   []string `json:"options"`
				Customize bool     `json:"customize"`
			} `json:"Variation1"`
			Variation2 struct {
				HasImage  string   `json:"has_image"`
				Name      bool     `json:"name"`
				Options   []string `json:"options"`
				Customize bool     `json:"customize"`
			} `json:"Variation2"`
			Variation3 struct {
				HasImage  bool     `json:"has_image"`
				Name      string   `json:"name"`
				Options   []string `json:"options"`
				Customize bool     `json:"customize"`
			} `json:"Variation3"`
			Variation4 struct {
				HasImage  bool     `json:"has_image"`
				Name      string   `json:"name"`
				Options   []string `json:"options"`
				Customize bool     `json:"customize"`
			} `json:"Variation4"`
		} `json:"variation"`
	} `json:"data"`
	RequestId string `json:"request_id"`
}

type UpdateSellableQuantityRsp struct {
	Code string `json:"code"`
	Data struct {
	} `json:"data"`
	RequestId string `json:"request_id"`
}

type UploadImageRsp struct {
	Code string `json:"code"`
	Data struct {
		Image struct {
			HashCode string `json:"hash_code"`
			Url      string `json:"url"`
		} `json:"image"`
	} `json:"data"`
	RequestId string `json:"request_id"`
}
