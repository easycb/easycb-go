package lazada

type AddOrUpdatePickupStopRsp struct {
	Retryable    bool   `json:"retryable"`
	Code         string `json:"code"`
	Success      bool   `json:"success"`
	ErrorMessage string `json:"errorMessage"`
	ErrorCode    string `json:"errorCode"`
	RequestId    string `json:"request_id"`
	Errors       []struct {
		Field        string `json:"field"`
		ErrorMessage string `json:"errorMessage"`
		ErrorCode    string `json:"errorCode"`
	} `json:"errors"`
}

type Create3PLStationRsp struct {
	Retryable    bool   `json:"retryable"`
	Code         string `json:"code"`
	Success      bool   `json:"success"`
	ErrorMessage string `json:"errorMessage"`
	ErrorCode    string `json:"errorCode"`
	RequestId    string `json:"request_id"`
	Errors       []struct {
		Field        string `json:"field"`
		ErrorMessage string `json:"errorMessage"`
		ErrorCode    string `json:"errorCode"`
	} `json:"errors"`
}

type GetOrderTraceRsp struct {
	Result struct {
		NotSuccess bool `json:"not_success"`
		Success    bool `json:"success"`
		Module     []struct {
			WarehouseDetailInfo   string `json:"warehouse_detail_info"`
			OfcOrderId            string `json:"ofc_order_id"`
			PackageDetailInfoList []struct {
				OrderLineInfoList      string `json:"order_line_info_list"`
				OfcPackageId           string `json:"ofc_package_id"`
				TrackingNumber         string `json:"tracking_number"`
				LogisticDetailInfoList []struct {
					PackageLocationName string        `json:"package_location_name"`
					StatusCode          string        `json:"status_code"`
					ProofImages         []interface{} `json:"proof_images"`
					DetailType          string        `json:"detail_type"`
					EventDate           string        `json:"event_date"`
					ReceiveTime         int64         `json:"receive_time"`
					Icon                string        `json:"icon"`
					Description         string        `json:"description"`
					Title               string        `json:"title"`
					EventTime           int64         `json:"event_time"`
				} `json:"logistic_detail_info_list"`
			} `json:"package_detail_info_list"`
		} `json:"module"`
		ErrorCode struct {
			DisplayMessage string `json:"displayMessage"`
		} `json:"error_code"`
		Repeated bool `json:"repeated"`
		Retry    bool `json:"retry"`
	} `json:"result"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type ScanParcelRsp struct {
	Code           string `json:"code"`
	TrackingNumber string `json:"trackingNumber"`
	RequestId      string `json:"request_id"`
}

type StationDopScanRsp struct {
	Code string `json:"code"`
	Data struct {
		TrackingNumber string `json:"trackingNumber"`
	} `json:"data"`
	Success bool `json:"success"`
	Error   struct {
		ErrorCode string `json:"errorCode"`
	} `json:"error"`
	RequestId string `json:"request_id"`
}

type Update3PLStationRsp struct {
	Retryable    bool   `json:"retryable"`
	Code         string `json:"code"`
	Success      bool   `json:"success"`
	ErrorMessage string `json:"errorMessage"`
	ErrorCode    string `json:"errorCode"`
	RequestId    string `json:"request_id"`
	Errors       []struct {
		Field        string `json:"field"`
		ErrorMessage string `json:"errorMessage"`
		ErrorCode    string `json:"errorCode"`
	} `json:"errors"`
}

type UpdatePickupTimeSlotRsp struct {
	Retryable    bool   `json:"retryable"`
	Code         string `json:"code"`
	Success      bool   `json:"success"`
	ErrorMessage string `json:"errorMessage"`
	ErrorCode    string `json:"errorCode"`
	RequestId    string `json:"request_id"`
	Errors       []struct {
		Field        string `json:"field"`
		ErrorMessage string `json:"errorMessage"`
		ErrorCode    string `json:"errorCode"`
	} `json:"errors"`
}

type CreateConsolidationServiceRsp struct {
	Code      string `json:"code"`
	Data      string `json:"data"`
	Success   bool   `json:"success"`
	ErrorCode string `json:"errorCode"`
	RequestId string `json:"request_id"`
	ErrorMsg  string `json:"errorMsg"`
}

type UpdateLastMileRsp struct {
	Code      string `json:"code"`
	Data      string `json:"data"`
	Success   bool   `json:"success"`
	ErrorCode string `json:"errorCode"`
	RequestId string `json:"request_id"`
	ErrorMsg  string `json:"errorMsg"`
}
