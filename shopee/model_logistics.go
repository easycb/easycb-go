package shopee

type GetShippingParameterRsp struct {
	BaseRsp
	Response struct {
		InfoNeeded struct {
			Dropoff       []string `json:"dropoff,omitempty"`
			Pickup        []string `json:"pickup,omitempty"`
			NonIntegrated []string `json:"non_integrated,omitempty"`
		} `json:"info_needed"`
		Dropoff struct {
			BranchList struct {
				BranchId int    `json:"branch_id"`
				Region   string `json:"region"`
				State    string `json:"state"`
				City     string `json:"city"`
				District string `json:"district"`
				Town     string `json:"town"`
				Address  string `json:"address"`
				Zipcode  string `json:"zipcode"`
			} `json:"branch_list,omitempty"`
			SlugList struct {
				Slug     string `json:"slug"`
				SlugName string `json:"slug_name"`
			} `json:"slug_list,omitempty"`
		} `json:"dropoff,omitempty"`
		Pickup struct {
			AddressList []struct {
				AddressId    int      `json:"address_id"`
				Region       string   `json:"region"`
				State        string   `json:"state"`
				City         string   `json:"city"`
				District     string   `json:"district"`
				Town         string   `json:"town"`
				Address      string   `json:"address"`
				Zipcode      string   `json:"zipcode"`
				AddressFlag  []string `json:"address_flag"`
				TimeSlotList struct {
					Date         int64 `json:"date"`
					TimeText     int64 `json:"time_text"`
					PickupTimeId int64 `json:"pickup_time_id"`
				} `json:"time_slot_list"`
			} `json:"address_list"`
		} `json:"pickup,omitempty"`
	} `json:"response"`
}

type GetTrackingNumberRsp struct {
	BaseRsp
	Response struct {
		TrackingNumber          string `json:"tracking_number"`
		FirstMileTrackingNumber string `json:"first_mile_tracking_number"`
	} `json:"response"`
}

type ShipOrderRsp struct {
	BaseRsp
}

type UpdateShippingOrderRsp struct {
	BaseRsp
}

type GetShippingDocumentParameterRsp struct {
	BaseRsp
	Response struct {
		ResultList []struct {
			OrderSn                        string   `json:"order_sn"`
			PackageNumber                  string   `json:"package_number,omitempty"`
			SuggestShippingDocumentType    string   `json:"suggest_shipping_document_type,omitempty"`
			SelectableShippingDocumentType []string `json:"selectable_shipping_document_type,omitempty"`
			FailError                      string   `json:"fail_error,omitempty"`
			FailMessage                    string   `json:"fail_message,omitempty"`
		} `json:"result_list"`
	} `json:"response"`
	Warning []struct {
		OrderSn string `json:"order_sn"`
	} `json:"warning"`
}

type CreateShippingDocumentRsp struct {
	BaseRsp
	Response struct {
		ResultList []struct {
			OrderSn string `json:"order_sn"`
		} `json:"result_list"`
	} `json:"response"`
}

type DownloadShippingDocumentRsp struct {
	BaseRsp
	RawBytes []byte `json:"raw_bytes"`
}

type GetShippingDocumentResultRsp struct {
	BaseRsp
	Warning []struct {
		OrderSn       string `json:"order_sn"`
		PackageNumber string `json:"package_number"`
	} `json:"warning"`
	Response struct {
		ResultList []struct {
			Status        string `json:"status,omitempty"`
			OrderSn       string `json:"order_sn"`
			PackageNumber string `json:"package_number"`
			FailMessage   string `json:"fail_message,omitempty"`
			FailError     string `json:"fail_error,omitempty"`
		} `json:"result_list"`
	} `json:"response"`
}

type GetTrackingInfoRsp struct {
	BaseRsp
	Response struct {
		LogisticsStatus string `json:"logistics_status"`
		OrderSn         string `json:"order_sn"`
		TrackingInfo    []struct {
			UpdateTime      int64  `json:"update_time"`
			Description     string `json:"description"`
			LogisticsStatus string `json:"logistics_status"`
		} `json:"tracking_info"`
	} `json:"response"`
}

type GetAddressListRsp struct {
	BaseRsp
	Response struct {
		ShowPickupAddress bool `json:"show_pickup_address"`
		AddressList       []struct {
			AddressId   int      `json:"address_id"`
			Region      string   `json:"region"`
			State       string   `json:"state"`
			City        string   `json:"city"`
			Address     string   `json:"address"`
			Zipcode     string   `json:"zipcode"`
			District    string   `json:"district"`
			Town        string   `json:"town"`
			AddressType []string `json:"address_type"`
		} `json:"address_list"`
	} `json:"response"`
}

type SetAddressConfigRsp struct {
	BaseRsp
}

type DeleteAddressRsp struct {
	BaseRsp
}

type GetChannelListRsp struct {
	BaseRsp
}

type UpdateChannelRsp struct {
	BaseRsp
}

type BatchShipOrderRsp struct {
	BaseRsp
	Response struct {
		ResultList []struct {
			OrderSn string `json:"order_sn"`
		} `json:"result_list"`
	} `json:"response"`
}

type GetShippingDocumentDataInfoRsp struct {
	BaseRsp
	Response struct {
		ShippingDocumentInfo struct {
			Cod                       bool    `json:"cod"`
			CodAmount                 string  `json:"cod_amount"`
			CreateDateYmdSl           string  `json:"create_date_ymd_sl"`
			DeliverArea               string  `json:"deliver_area"`
			DeliveryHub               string  `json:"delivery_hub"`
			DestinationBaseCode       string  `json:"destination_base_code"`
			EcOrderNo                 string  `json:"ec_order_no"`
			IsLmDgBool                int     `json:"is_lm_dg_bool"`
			LastMileTrackingNumber    string  `json:"last_mile_tracking_number"`
			LastThirdDigitsBuyerPhone string  `json:"last_third_digits_buyer_phone"`
			LogisticsChannelId        int     `json:"logistics_channel_id"`
			ManufacturersName         string  `json:"manufacturers_name"`
			ManufacturersWebsite      string  `json:"manufacturers_website"`
			MutualCheck               string  `json:"mutual_check"`
			OrderWeight               float64 `json:"order_weight"`
			ParcelSize                string  `json:"parcel_size"`
			PickupHub                 string  `json:"pickup_hub"`
			PreferredDeliveryOption   int     `json:"preferred_delivery_option"`
			RecipientSortCode         struct {
				FirstRecipientSortCode  string `json:"first_recipient_sort_code"`
				SecondRecipientSortCode string `json:"second_recipient_sort_code"`
				ThirdRecipientSortCode  string `json:"third_recipient_sort_code"`
			} `json:"recipient_sort_code"`
			ReturnSortCode struct {
				ReturnFirstSortCode string `json:"return_first_sort_code"`
			} `json:"return_sort_code"`
			SenderSortCode struct {
				FirstSenderSortCode  string `json:"first_sender_sort_code"`
				SecondSenderSortCode string `json:"second_sender_sort_code"`
				ThirdSenderSortCode  string `json:"third_sender_sort_code"`
			} `json:"sender_sort_code"`
			ShippingCarrier      string `json:"shipping_carrier"`
			ShopeeTrackingNumber string `json:"shopee_tracking_number"`
			Sod                  bool   `json:"sod"`
			SpxReceiveStation    struct {
				SpxFirstReceiveStation string `json:"spx_first_receive_station"`
			} `json:"spx_receive_station"`
			SpxSubDistrict         string `json:"spx_sub_district"`
			ThirdPartyLogisticInfo struct {
				Area                    string `json:"area"`
				Barcode                 string `json:"barcode"`
				BarcodeDc               string `json:"barcode_dc"`
				BarcodeNo1              string `json:"barcode_no1"`
				BarcodeNo2              string `json:"barcode_no2"`
				BarcodeNo3              string `json:"barcode_no3"`
				BarcodeNo4              string `json:"barcode_no4"`
				BarcodeNo5              string `json:"barcode_no5"`
				BarcodePr               string `json:"barcode_pr"`
				BranchCode              string `json:"branch_code"`
				BranchName              string `json:"branch_name"`
				BuyerPreferDeliveryTime struct {
					Description string `json:"description"`
					EndTime     string `json:"end_time"`
					SlotId      string `json:"slot_id"`
					StartTime   string `json:"start_time"`
				} `json:"buyer_prefer_delivery_time"`
				CustomerCode                  string `json:"customer_code"`
				DeliverAreaTxt                string `json:"deliver_area_txt"`
				DeliverDateYmd                string `json:"deliver_date_ymd"`
				DeliverRouter                 string `json:"deliver_router"`
				EcBarCode16                   string `json:"ec_bar_code16"`
				EcBarCode9                    string `json:"ec_bar_code9"`
				EcOrderNumber                 string `json:"ec_order_number"`
				EcSupplierName                string `json:"ec_supplier_name"`
				EquipmentId                   string `json:"equipment_id"`
				EshopId                       string `json:"eshop_id"`
				FirstPickBarcode              string `json:"first_pick_barcode"`
				IsCodBool                     bool   `json:"is_cod_bool"`
				LargeLogisticsId              string `json:"large_logistics_id"`
				LastThirdDigitsRecipientPhone string `json:"last_third_digits_recipient_phone"`
				LastThirdDigitsSenderPhone    string `json:"last_third_digits_sender_phone"`
				LmhubStationCode              string `json:"lmhub_station_code"`
				LmhubStationName              string `json:"lmhub_station_name"`
				ManufacturersName             string `json:"manufacturers_name"`
				ManufacturersWebsite          string `json:"manufacturers_website"`
				MdDriverCode                  string `json:"md_driver_code"`
				OkAisleNo                     string `json:"ok_aisle_no"`
				OkGridNo                      string `json:"ok_grid_no"`
				OkMidType                     string `json:"ok_mid_type"`
				OkTrackingNumber              string `json:"ok_tracking_number"`
				OrderNo                       string `json:"order_no"`
				ParentId                      string `json:"parent_id"`
				PelicanTrackingNo             string `json:"pelican_tracking_no"`
				PickRouter                    string `json:"pick_router"`
				PrintDate                     string `json:"print_date"`
				PrintDatetime                 string `json:"print_datetime"`
				Prompt                        string `json:"prompt"`
				PurchaseTime                  int64  `json:"purchase_time"`
				PutorderStackzoneCode         string `json:"putorder_stackzone_code"`
				Pzip                          string `json:"pzip"`
				PzipC                         string `json:"pzip_c"`
				Qrcode                        string `json:"qrcode"`
				RcvStoreName                  string `json:"rcv_store_name"`
				RecipientArea                 string `json:"recipient_area"`
				ReturnCycle                   string `json:"return_cycle"`
				ReturnMode                    string `json:"return_mode"`
				ReturnTime                    string `json:"return_time"`
				RouteStep                     string `json:"route_step"`
				SdDriverCode                  string `json:"sd_driver_code"`
				SecondPickBarcode             string `json:"second_pick_barcode"`
				ServiceDescription            string `json:"service_description"`
				ShipType                      string `json:"ship_type"`
				StoreType                     string `json:"store_type"`
				Suda5Code                     string `json:"suda5_code"`
				TwLastThreeDigitsBuyerPhone   string `json:"tw_last_three_digits_buyer_phone"`
				TwStoreName                   string `json:"tw_store_name"`
				TwStoreNumber                 string `json:"tw_store_number"`
			} `json:"third_party_logistic_info"`
			TrackingNumber string `json:"tracking_number"`
			Zone           string `json:"zone"`
			ZoneCode       string `json:"zone_code"`
		} `json:"shipping_document_info"`
		RecipientAddressInfo []struct {
			Key   string `json:"key"`
			Image string `json:"image"`
		} `json:"recipient_address_info"`
	} `json:"response"`
	RequestId string `json:"request_id"`
}

type UpdateTrackingStatusRsp struct {
	BaseRsp
	Response struct {
		UpdateResult string `json:"update_result"`
	} `json:"response"`
}

type GetBookingShippingParameterRsp struct {
	BaseRsp
	Response struct {
		InfoNeeded struct {
			Dropoff []interface{} `json:"dropoff"`
			Pickup  []string      `json:"pickup"`
		} `json:"info_needed"`
		Pickup struct {
			AddressList []struct {
				AddressId    int64    `json:"address_id"`
				Region       string   `json:"region"`
				State        string   `json:"state"`
				City         string   `json:"city"`
				District     string   `json:"district"`
				Town         string   `json:"town"`
				Address      string   `json:"address"`
				Zipcode      string   `json:"zipcode"`
				AddressFlag  []string `json:"address_flag"`
				TimeSlotList []struct {
					Date         int64  `json:"date"`
					TimeText     string `json:"time_text"`
					PickupTimeId string `json:"pickup_time_id"`
					Error        string `json:"error"`
					Msg          string `json:"msg"`
				} `json:"time_slot_list"`
			} `json:"address_list"`
		} `json:"pickup"`
	} `json:"response"`
}

type ShipBookingRsp struct {
	BaseRsp
}

type GetBookingTrackingNumberRsp struct {
	BaseRsp
	TrackingNumber string `json:"tracking_number"`
}

type GetBookingShippingDocumentParameterRsp struct {
	BaseRsp
	Warning []struct {
		BookingSn string `json:"booking_sn"`
	} `json:"warning"`
	Response struct {
		ResultList []struct {
			BookingSn                      string   `json:"booking_sn"`
			SuggestShippingDocumentType    string   `json:"suggest_shipping_document_type"`
			SelectableShippingDocumentType []string `json:"selectable_shipping_document_type"`
			FailError                      string   `json:"fail_error"`
			FailMessage                    string   `json:"fail_message"`
		} `json:"result_list"`
	} `json:"response"`
}

type CreateBookingShippingDocumentRsp struct {
	BaseRsp
	Warning []struct {
		BookingSn string `json:"booking_sn"`
	} `json:"warning"`
	Response struct {
		ResultList []struct {
			BookingSn   string `json:"booking_sn"`
			FailError   string `json:"fail_error"`
			FailMessage string `json:"fail_message"`
		} `json:"result_list"`
	} `json:"response"`
}

type GetBookingShippingDocumentResultRsp struct {
	BaseRsp
	Warning []struct {
		BookingSn string `json:"booking_sn"`
	} `json:"warning"`
	Response struct {
		ResultList []struct {
			BookingSn   string `json:"booking_sn"`
			Status      string `json:"status"`
			FailError   string `json:"fail_error"`
			FailMessage string `json:"fail_message"`
		} `json:"result_list"`
	} `json:"response"`
}

type DownloadBookingShippingDocumentRsp struct {
	BaseRsp
	RawBytes []byte `json:"raw_bytes"`
}

type GetBookingShippingDocumentDataInfoRsp struct {
	BaseRsp
	Response struct {
		RecipientAddressInfo struct {
			Key   string `json:"key"`
			Image string `json:"image"`
		} `json:"recipient_address_info"`
		ShippingDocumentInfo struct {
			BookingWeight      int64  `json:"booking_weight"`
			LogisticsChannelId int64  `json:"logistics_channel_id"`
			ShippingCarrier    string `json:"shipping_carrier"`
			RecipientSortCode  struct {
				FirstRecipientSortCode  string `json:"first_recipient_sort_code"`
				SecondRecipientSortCode string `json:"second_recipient_sort_code"`
				ThirdRecipientSortCode  string `json:"third_recipient_sort_code"`
			} `json:"recipient_sort_code"`
			SenderSortCode struct {
				FirstSenderSortCode  string `json:"first_sender_sort_code"`
				SecondSenderSortCode string `json:"second_sender_sort_code"`
				ThirdSenderSortCode  string `json:"third_sender_sort_code"`
			} `json:"sender_sort_code"`
			ReturnSortCode struct {
				ReturnFirstSortCode string `json:"return_first_sort_code"`
			} `json:"return_sort_code"`
			TrackingNumber       string `json:"tracking_number"`
			PickupHub            string `json:"pickup_hub"`
			DeliveryHub          string `json:"delivery_hub"`
			DeliverArea          string `json:"deliver_area"`
			EcBookingNo          string `json:"ec_booking_no"`
			CreateDateYmdSl      string `json:"create_date_ymd_sl"`
			ManufacturersName    string `json:"manufacturers_name"`
			ManufacturersWebsite string `json:"manufacturers_website"`
			IsLmDgBool           int    `json:"is_lm_dg_bool"`
			SpxSubDistrict       string `json:"spx_sub_district"`
			SpxReceiveStation    struct {
				SpxFirstReceiveStation string `json:"spx_first_receive_station"`
			} `json:"spx_receive_station"`
			Zone                string `json:"zone"`
			ZoneCode            string `json:"zone_code"`
			DestinationBaseCode string `json:"destination_base_code"`
		} `json:"shipping_document_info"`
	} `json:"response"`
}

type GetBookingTrackingInfoRsp struct {
	BaseRsp
	Response struct {
		BookingSn       string `json:"booking_sn"`
		LogisticsStatus string `json:"logistics_status"`
		TrackingInfo    []struct {
			UpdateTime      int64  `json:"update_time"`
			Description     string `json:"description"`
			LogisticsStatus string `json:"logistics_status"`
		} `json:"tracking_info"`
	} `json:"response"`
}
