package tiktok

type GetFBTMerchantOnboardedRegionsRsp struct {
	BaseRsp
	Data struct {
		OnboardedRegions []struct {
			RegionCode string `json:"region_code"`
		} `json:"onboarded_regions"`
	} `json:"data"`
}

type GetFBTWarehouseListRsp struct {
	BaseRsp
	Data struct {
		Warehouses []struct {
			Addresses []struct {
				AddressLine1 string `json:"address_line_1"`
				AddressLine2 string `json:"address_line_2"`
				AddressLine3 string `json:"address_line_3"`
				City         string `json:"city"`
				District     string `json:"district"`
				PostalCode   string `json:"postal_code"`
				RegionCode   string `json:"region_code"`
				State        string `json:"state"`
			} `json:"addresses"`
			FbtWarehouseId    string `json:"fbt_warehouse_id"`
			LogisticsServices []struct {
				Id         string `json:"id"`
				Name       string `json:"name"`
				Subscribed bool   `json:"subscribed"`
			} `json:"logistics_services"`
			Name         string   `json:"name"`
			Subscribed   bool     `json:"subscribed"`
			Type         string   `json:"type"`
			WarehouseIds []string `json:"warehouse_ids"`
		} `json:"warehouses"`
	} `json:"data"`
}

type GetInboundOrderRsp struct {
	BaseRsp
	Data struct {
		InboundOrders []struct {
			ActualArriveTime int64 `json:"actual_arrive_time"`
			Carriers         []struct {
				CarrierName    string `json:"carrier_name"`
				TrackingNumber string `json:"tracking_number"`
			} `json:"carriers"`
			CreateTime       int64  `json:"create_time"`
			ExpectArriveTime int64  `json:"expect_arrive_time"`
			Id               string `json:"id"`
			Merchant         struct {
				Id   string `json:"id"`
				Name string `json:"name"`
			} `json:"merchant"`
			OrderOperationLogs []struct {
				OperateTime int64  `json:"operate_time"`
				OrderStatus string `json:"order_status"`
			} `json:"order_operation_logs"`
			PlannedGoods []struct {
				Id            string   `json:"id"`
				Name          string   `json:"name"`
				Quantity      int      `json:"quantity"`
				ReferenceCode string   `json:"reference_code"`
				SkuIds        []string `json:"sku_ids"`
			} `json:"planned_goods"`
			ReceivedBatches []struct {
				DefectiveQuantity int      `json:"defective_quantity"`
				GoodsId           string   `json:"goods_id"`
				Id                string   `json:"id"`
				NormalQuantity    int      `json:"normal_quantity"`
				ProductIds        []string `json:"product_ids"`
				SkuIds            []string `json:"sku_ids"`
				TotalQuantity     int      `json:"total_quantity"`
			} `json:"received_batches"`
			ShipTime  int    `json:"ship_time"`
			Type      string `json:"type"`
			Warehouse struct {
				FbtWarehouseId string   `json:"fbt_warehouse_id"`
				Name           string   `json:"name"`
				Type           string   `json:"type"`
				WarehouseIds   []string `json:"warehouse_ids"`
			} `json:"warehouse"`
		} `json:"inbound_orders"`
	} `json:"data"`
}

type SearchFBTInventoryRsp struct {
	BaseRsp
	Data struct {
		Inventory []struct {
			FbtWarehouseId string `json:"fbt_warehouse_id"`
			Goods          struct {
				Id            string `json:"id"`
				Name          string `json:"name"`
				ReferenceCode string `json:"reference_code"`
				Skus          []struct {
					Id           string `json:"id"`
					OnHandDetail struct {
						AvailableQuantity int `json:"available_quantity"`
						ReservedQuantity  int `json:"reserved_quantity"`
						TotalQuantity     int `json:"total_quantity"`
					} `json:"on_hand_detail"`
				} `json:"skus"`
			} `json:"goods"`
			InTransitQuantity int `json:"in_transit_quantity"`
			OnHandDetail      struct {
				AvailableQuantity     int `json:"available_quantity"`
				ReservedQuantity      int `json:"reserved_quantity"`
				TotalQuantity         int `json:"total_quantity"`
				UnfulfillableQuantity int `json:"unfulfillable_quantity"`
			} `json:"on_hand_detail"`
		} `json:"inventory"`
		NextPageToken string `json:"next_page_token"`
		TotalCount    int    `json:"total_count"`
	} `json:"data"`
}

type SearchFBTInventoryRecordRsp struct {
	BaseRsp
	Data struct {
		InventoryRecords []struct {
			ChangedQuantity     int    `json:"changed_quantity"`
			CreateTime          int64  `json:"create_time"`
			FbtWarehouseId      string `json:"fbt_warehouse_id"`
			FinalOnHandQuantity int    `json:"final_on_hand_quantity"`
			Goods               struct {
				Id            string `json:"id"`
				Name          string `json:"name"`
				ReferenceCode string `json:"reference_code"`
			} `json:"goods"`
			InitialOnHandQuantity int    `json:"initial_on_hand_quantity"`
			InventoryGoodsType    string `json:"inventory_goods_type"`
			Order                 struct {
				Id   string `json:"id"`
				Type string `json:"type"`
			} `json:"order"`
		} `json:"inventory_records"`
		NextPageToken string `json:"next_page_token"`
		TotalCount    int    `json:"total_count"`
	} `json:"data"`
}

type SearchGoodsInfoRsp struct {
	BaseRsp
	Data struct {
		Goods []struct {
			Barcodes []struct {
				Code string `json:"code"`
				Type string `json:"type"`
			} `json:"barcodes"`
			Id                string `json:"id"`
			ImageUrl          string `json:"image_url"`
			LotExpirationInfo struct {
				Addresses struct {
					AddressLine1 string `json:"address_line_1"`
					AddressLine2 string `json:"address_line_2"`
					AddressLine3 string `json:"address_line_3"`
					City         string `json:"city"`
					District     string `json:"district"`
					Name         string `json:"name"`
					PhoneNumber  string `json:"phone_number"`
					PostalCode   string `json:"postal_code"`
					RegionCode   string `json:"region_code"`
					State        string `json:"state"`
				} `json:"addresses"`
				ExpirationAlertDays    int    `json:"expiration_alert_days"`
				HandlingMethod         string `json:"handling_method"`
				InboundCutoffDays      int    `json:"inbound_cutoff_days"`
				IsExpirationManagement bool   `json:"is_expiration_management"`
				IsLotControl           bool   `json:"is_lot_control"`
				ReturnCycle            string `json:"return_cycle"`
				SalesCutoffDays        int    `json:"sales_cutoff_days"`
				ShelfLifeDays          int    `json:"shelf_life_days"`
			} `json:"lot_expiration_info"`
			MerchantDeclarationInfo struct {
				Dimension struct {
					Height string `json:"height"`
					Length string `json:"length"`
					Unit   string `json:"unit"`
					Width  string `json:"width"`
				} `json:"dimension"`
				Weight struct {
					Unit  string `json:"unit"`
					Value string `json:"value"`
				} `json:"weight"`
			} `json:"merchant_declaration_info"`
			Name          string `json:"name"`
			ReferenceCode string `json:"reference_code"`
			Skus          []struct {
				Id       string `json:"id"`
				ImageUrl string `json:"image_url"`
				Matched  bool   `json:"matched"`
				Name     string `json:"name"`
				Product  struct {
					Id       string `json:"id"`
					ImageUrl string `json:"image_url"`
					Name     string `json:"name"`
				} `json:"product"`
				Regions []string `json:"regions"`
			} `json:"skus"`
			WarehouseConfirmationInfo struct {
				Dimension struct {
					Height string `json:"height"`
					Length string `json:"length"`
					Unit   string `json:"unit"`
					Width  string `json:"width"`
				} `json:"dimension"`
				Weight struct {
					Unit  string `json:"unit"`
					Value string `json:"value"`
				} `json:"weight"`
			} `json:"warehouse_confirmation_info"`
		} `json:"goods"`
		NextPageToken string `json:"next_page_token"`
		TotalCount    int    `json:"total_count"`
	} `json:"data"`
}
