package ozon

import "time"

type FbpOrderDirectCancelRsp struct {
	BaseRsp

	Example struct {
		Error struct {
			OrderErrors []string `json:"order_errors"`
		} `json:"error"`
		IsError    bool  `json:"is_error"`
		RowVersion int64 `json:"row_version"`
	} `json:"example"`
}

type FbpOrderDirectSellerDlvEditRsp struct {
	BaseRsp

	Error struct {
		OrderErrors string `json:"order_errors"`
	} `json:"error"`
	IsError    bool  `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type FbpOrderDirectTimeslotEditRsp struct {
	BaseRsp

	ErrorReasons []string `json:"error_reasons"`
	RowVersion   int64    `json:"row_version"`
}

type FbpOrderDirectTimeslotListRsp struct {
	BaseRsp

	Reasons   []string `json:"reasons"`
	Timeslots []struct {
		TimeslotEnd   time.Time `json:"timeslot_end"`
		TimeslotStart time.Time `json:"timeslot_start"`
	} `json:"timeslots"`
	WarehouseTimezoneName string `json:"warehouse_timezone_name"`
}
