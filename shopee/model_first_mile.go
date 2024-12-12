package shopee

type FirstMileGetUnbindOrderListRsp struct {
	BaseRsp
	Response struct {
		OrderList []struct {
			OrderSn         string `json:"order_sn"`
			PackageNumber   string `json:"package_number"`
			LogisticsStatus string `json:"logistics_status"`
		} `json:"order_list"`
		More       bool   `json:"more"`
		NextCursor string `json:"next_cursor"`
	} `json:"response"`
}

type FirstMileGetDetailRsp struct {
	BaseRsp
	Response struct {
		FirstMileTrackingNumber string `json:"first_mile_tracking_number"`
		ShipmentMethod          string `json:"shipment_method"`
		LogisticsChannelId      int    `json:"logistics_channel_id"`
		Status                  string `json:"status"`
		DeclareDate             string `json:"declare_date"`
		OrderList               []struct {
			OrderSn                 string `json:"order_sn"`
			PackageNumber           string `json:"package_number"`
			SlsTrackingNumber       string `json:"sls_tracking_number"`
			PickUpDone              bool   `json:"pick_up_done"`
			ArrivedTransitWarehouse bool   `json:"arrived_transit_warehouse"`
		} `json:"order_list"`
		More       bool   `json:"more"`
		NextCursor string `json:"next_cursor"`
	} `json:"response"`
}

type FirstMileGenerateFirstMileTrackingNumberRsp struct {
	BaseRsp
	Response struct {
		FirstMileTrackingNumberList []string `json:"first_mile_tracking_number_list"`
	} `json:"response"`
}

type FirstMileBindFirstMileTrackingNumberRsp struct {
	BaseRsp
	Response struct {
		FirstMileTrackingNumber string `json:"first_mile_tracking_number"`
		OrderList               []struct {
			OrderSn string `json:"order_sn"`
		} `json:"order_list"`
	} `json:"response"`
}

type FirstMileUnbindFirstMileTrackingNumberRsp struct {
	BaseRsp
	Response struct {
		FirstMileTrackingNumber string `json:"first_mile_tracking_number"`
		OrderList               []struct {
			OrderSn       string `json:"order_sn"`
			FailMessage   string `json:"fail_message"`
			FailError     string `json:"fail_error"`
			PackageNumber string `json:"package_number,omitempty"`
		} `json:"order_list"`
	} `json:"response"`
	Warning []struct {
		OrderSn       string `json:"order_sn"`
		PackageNumber string `json:"package_number,omitempty"`
	} `json:"warning"`
}

type FirstMileGetTrackingNumberListRsp struct {
	BaseRsp
	Response struct {
		FirstMileTrackingNumberList []struct {
			FirstMileTrackingNumber string `json:"first_mile_tracking_number"`
			Status                  string `json:"status"`
			DeclareDate             string `json:"declare_date"`
		} `json:"first_mile_tracking_number_list"`
		More       bool   `json:"more"`
		NextCursor string `json:"next_cursor"`
	} `json:"response"`
}

type FirstMileGetWallBillRsp struct {
	BaseRsp
	Waybill  string `json:"waybill"`
	RawBytes []byte `json:"raw_bytes"`
}
