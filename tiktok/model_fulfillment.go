package tiktok

type GetOrderSplitAttributesRsp struct {
	BaseRsp
	Data struct {
		SplitAttributes []struct {
			CanSplit bool   `json:"can_split"`
			OrderId  string `json:"order_id"`
			Reason   string `json:"reason"`
		} `json:"split_attributes"`
	} `json:"data"`
}

type SplitOrdersRsp struct {
	BaseRsp
	Data struct {
		Packages []struct {
			Id                string `json:"id"`
			SplittableGroupId string `json:"splittable_group_id"`
		} `json:"packages"`
	} `json:"data"`
}

type GetEligibleShippingServiceRsp struct {
	BaseRsp
	Data struct {
		Dimension struct {
			Height string `json:"height"`
			Length string `json:"length"`
			Unit   string `json:"unit"`
			Width  string `json:"width"`
		} `json:"dimension"`
		OrderId          string   `json:"order_id"`
		OrderLineId      []string `json:"order_line_id"`
		ShippingServices []struct {
			Currency             string `json:"currency"`
			EarliestDeliveryDays int    `json:"earliest_delivery_days"`
			Id                   string `json:"id"`
			IsDefault            bool   `json:"is_default"`
			LatestDeliveryDays   int    `json:"latest_delivery_days"`
			Name                 string `json:"name"`
			Price                string `json:"price"`
			ShippingProviderId   string `json:"shipping_provider_id"`
			ShippingProviderName string `json:"shipping_provider_name"`
		} `json:"shipping_services"`
		Weight struct {
			Unit  string `json:"unit"`
			Value string `json:"value"`
		} `json:"weight"`
	} `json:"data"`
}

type CreateFirstMileBundleRsp struct {
	BaseRsp
	Data struct {
		Errors []struct {
			Code   int `json:"code"`
			Detail struct {
				OrderId string `json:"order_id"`
			} `json:"detail"`
			Message string `json:"message"`
		} `json:"errors"`
		FirstMileBundleId string `json:"first_mile_bundle_id"`
		Url               string `json:"url"`
	} `json:"data"`
}

type CreatePackagesRsp struct {
	BaseRsp
	Data struct {
		CreateTime int `json:"create_time"`
		Dimension  struct {
			Height string `json:"height"`
			Length string `json:"length"`
			Unit   string `json:"unit"`
			Width  string `json:"width"`
		} `json:"dimension"`
		OrderId             string   `json:"order_id"`
		OrderLineItemIds    []string `json:"order_line_item_ids"`
		PackageId           string   `json:"package_id"`
		ShippingServiceInfo struct {
			Currency             string `json:"currency"`
			EarliestDeliveryDays int    `json:"earliest_delivery_days"`
			Id                   string `json:"id"`
			LatestDeliveryDays   int    `json:"latest_delivery_days"`
			Name                 string `json:"name"`
			Price                string `json:"price"`
			ShippingProviderId   string `json:"shipping_provider_id"`
			ShippingProviderName string `json:"shipping_provider_name"`
		} `json:"shipping_service_info"`
		Weight struct {
			Unit  string `json:"unit"`
			Value string `json:"value"`
		} `json:"weight"`
	} `json:"data"`
}

type SearchPackageRsp struct {
	BaseRsp
	Data struct {
		NextPageToken string `json:"next_page_token"`
		Packages      []struct {
			CreateTime       int64    `json:"create_time"`
			Id               string   `json:"id"`
			OrderLineItemIds []string `json:"order_line_item_ids"`
			Orders           []struct {
				Id   string `json:"id"`
				Skus []struct {
					Id       string `json:"id"`
					ImageUrl string `json:"image_url"`
					Name     string `json:"name"`
					Quantity int    `json:"quantity"`
				} `json:"skus"`
			} `json:"orders"`
			ShippingProviderId   string `json:"shipping_provider_id"`
			ShippingProviderName string `json:"shipping_provider_name"`
			Status               string `json:"status"`
			TrackingNumber       string `json:"tracking_number"`
			UpdateTime           int64  `json:"update_time"`
		} `json:"packages"`
		TotalCount int `json:"total_count"`
	} `json:"data"`
}

type SearchCombinablePackagesRsp struct {
	BaseRsp
	Data struct {
		CombinablePackages []struct {
			Id       string   `json:"id"`
			OrderIds []string `json:"order_ids"`
		} `json:"combinable_packages"`
		NextPageToken string `json:"next_page_token"`
		TotalCount    int    `json:"total_count"`
	} `json:"data"`
}

type CombinePackageRsp struct {
	BaseRsp
	Data struct {
		Errors []struct {
			Code   int `json:"code"`
			Detail struct {
				PackageId string `json:"package_id"`
			} `json:"detail"`
			Message string `json:"message"`
		} `json:"errors"`
		Packages []struct {
			Id       string   `json:"id"`
			OrderIds []string `json:"order_ids"`
		} `json:"packages"`
	} `json:"data"`
}

type UnCombinePackagesRsp struct {
	BaseRsp
	Data struct {
		Packages []struct {
			Id       string   `json:"id"`
			OrderIds []string `json:"order_ids"`
		} `json:"packages"`
	} `json:"data"`
}

type GetPackageHandoverTimeSlotsRsp struct {
	BaseRsp
	Data struct {
		CanDropOff       bool   `json:"can_drop_off"`
		CanPickup        bool   `json:"can_pickup"`
		CanVanCollection bool   `json:"can_van_collection"`
		DropOffPointUrl  string `json:"drop_off_point_url"`
		PickupSlots      []struct {
			Avaliable bool  `json:"avaliable"`
			EndTime   int64 `json:"end_time"`
			StartTime int64 `json:"start_time"`
		} `json:"pickup_slots"`
	} `json:"data"`
}

type ShipPackageRsp struct {
	BaseRsp
}

type BatchShipPackagesRsp struct {
	BaseRsp
	Data struct {
		Errors []struct {
			Code   int `json:"code"`
			Detail struct {
				PackageId string `json:"package_id"`
			} `json:"detail"`
			Message string `json:"message"`
		} `json:"errors"`
	} `json:"data"`
}

type MarkPackageAsShippedRsp struct {
	BaseRsp
	Data struct {
		OrderId          string   `json:"order_id"`
		OrderLineItemIds []string `json:"order_line_item_ids"`
		PackageId        string   `json:"package_id"`
		Warning          struct {
			Message string `json:"message"`
		} `json:"warning"`
	} `json:"data"`
}

type GetPackageShippingDocumentRsp struct {
	BaseRsp
	Data struct {
		DocUrl         string `json:"doc_url"`
		TrackingNumber string `json:"tracking_number"`
	} `json:"data"`
}

type GetPackageDetailRsp struct {
	BaseRsp
	Data struct {
		CreateTime         int64  `json:"create_time"`
		DeliveryOptionId   string `json:"delivery_option_id"`
		DeliveryOptionName string `json:"delivery_option_name"`
		Dimension          struct {
			Height string `json:"height"`
			Length string `json:"length"`
			Unit   string `json:"unit"`
			Width  string `json:"width"`
		} `json:"dimension"`
		HandoverMethod         string   `json:"handover_method"`
		HasMultiSkus           bool     `json:"has_multi_skus"`
		LastMileTrackingNumber string   `json:"last_mile_tracking_number"`
		NoteTag                string   `json:"note_tag"`
		OrderLineItemIds       []string `json:"order_line_item_ids"`
		Orders                 []struct {
			Id   string `json:"id"`
			Skus []struct {
				Id       string `json:"id"`
				ImageUrl string `json:"image_url"`
				Name     string `json:"name"`
				Quantity int    `json:"quantity"`
			} `json:"skus"`
		} `json:"orders"`
		PackageId     string `json:"package_id"`
		PackageStatus string `json:"package_status"`
		PickupSlot    struct {
			EndTime   int64 `json:"end_time"`
			StartTime int64 `json:"start_time"`
		} `json:"pickup_slot"`
		RecipientAddress struct {
			AddressDetail string `json:"address_detail"`
			AddressLine1  string `json:"address_line1"`
			AddressLine2  string `json:"address_line2"`
			AddressLine3  string `json:"address_line3"`
			AddressLine4  string `json:"address_line4"`
			FullAddress   string `json:"full_address"`
			Name          string `json:"name"`
			PhoneNumber   string `json:"phone_number"`
			PostalCode    string `json:"postal_code"`
			RegionCode    string `json:"region_code"`
		} `json:"recipient_address"`
		SenderAddress struct {
			AddressDetail string `json:"address_detail"`
			AddressLine1  string `json:"address_line1"`
			AddressLine2  string `json:"address_line2"`
			AddressLine3  string `json:"address_line3"`
			AddressLine4  string `json:"address_line4"`
			FullAddress   string `json:"full_address"`
			Name          string `json:"name"`
			PhoneNumber   string `json:"phone_number"`
			PostalCode    string `json:"postal_code"`
			RegionCode    string `json:"region_code"`
		} `json:"sender_address"`
		ShippingProviderId   string `json:"shipping_provider_id"`
		ShippingProviderName string `json:"shipping_provider_name"`
		ShippingType         string `json:"shipping_type"`
		SplitAndCombineTag   string `json:"split_and_combine_tag"`
		TrackingNumber       string `json:"tracking_number"`
		UpdateTime           int64  `json:"update_time"`
		Weight               struct {
			Unit  string `json:"unit"`
			Value string `json:"value"`
		} `json:"weight"`
	} `json:"data"`
}

type GetTrackingRsp struct {
	BaseRsp
	Data struct {
		Tracking []struct {
			Description      string `json:"description"`
			UpdateTimeMillis int64  `json:"update_time_millis"`
		} `json:"tracking"`
	} `json:"data"`
}

type UpdateShippingInfoRsp struct {
	BaseRsp
}

type UpdatePackageShippingInfoRsp struct {
	BaseRsp
}

type FulfillmentUploadDeliveryFileRsp struct {
	BaseRsp
	Data struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"data"`
}

type FulfillmentUploadDeliveryImageRsp struct {
	BaseRsp
	Data struct {
		Height int    `json:"height"`
		Url    string `json:"url"`
		Width  int    `json:"width"`
	} `json:"data"`
}

type UpdatePackageDeliveryStatusRsp struct {
	BaseRsp
	Data struct {
		Errors []struct {
			Code   int `json:"code"`
			Detail struct {
				PackageId string `json:"package_id"`
			} `json:"detail"`
			Message string `json:"message"`
		} `json:"errors"`
	} `json:"data"`
}
