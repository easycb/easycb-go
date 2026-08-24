package ozon

import "time"

type FbpDraftDirectSellerDlvCreateRsp struct {
	BaseRsp

	DraftId    int64  `json:"draft_id"`
	RowVersion int64  `json:"row_version"`
	SupplyId   string `json:"supply_id"`
}

type FbpDraftDirectSellerDlvEditRsp struct {
	BaseRsp

	Error struct {
		Errors string `json:"errors"`
	} `json:"error"`
	IsError    bool  `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type FbpDraftDirectTimeslotEditRsp struct {
	BaseRsp

	ErrorReasons string `json:"error_reasons"`
	RowVersion   int64  `json:"row_version"`
}

type FbpDraftDirectTimeslotGetRsp struct {
	BaseRsp

	Reasons   []string `json:"reasons"`
	Timeslots []struct {
		TimeslotEnd   time.Time `json:"timeslot_end"`
		TimeslotStart time.Time `json:"timeslot_start"`
	} `json:"timeslots"`
	WarehouseTimezoneName string `json:"warehouse_timezone_name"`
}

type FbpDraftDirectCreateRsp struct {
	BaseRsp

	DraftId    int64  `json:"draft_id"`
	RowVersion int64  `json:"row_version"`
	SupplyId   string `json:"supply_id"`
}

type FbpDraftDirectDeleteRsp struct {
	BaseRsp

	CancellationState struct {
		CancellationError struct {
			ErrorCode string `json:"error_code"`
			Message   string `json:"message"`
		} `json:"cancellation_error"`
		CancellationStatus string `json:"cancellation_status"`
	} `json:"cancellation_state"`
	RowVersion int64 `json:"row_version"`
}

type FbpDraftDirectProductValidateRsp struct {
	BaseRsp

	ApprovedItems []struct {
		Barcode  string  `json:"barcode"`
		IconName string  `json:"icon_name"`
		Name     string  `json:"name"`
		OfferId  string  `json:"offer_id"`
		Quantity int     `json:"quantity"`
		Sku      int64   `json:"sku"`
		Volume   float64 `json:"volume"`
	} `json:"approved_items"`
	BundleGenerated bool   `json:"bundle_generated"`
	BundleId        string `json:"bundle_id"`
	RejectedItems   []struct {
		Barcode          string  `json:"barcode"`
		IconName         string  `json:"icon_name"`
		Name             string  `json:"name"`
		OfferId          string  `json:"offer_id"`
		Quantity         int     `json:"quantity"`
		RejectionReasons string  `json:"rejection_reasons"`
		Sku              int64   `json:"sku"`
		Volume           float64 `json:"volume"`
	} `json:"rejected_items"`
}

type FbpDraftDirectRegistrateRsp struct {
	BaseRsp

	Error struct {
		BundleErrors []struct {
			Errors string `json:"errors"`
			Sku    int64  `json:"sku"`
		} `json:"bundle_errors"`
		OrderError string `json:"order_error"`
	} `json:"error"`
	IsError    bool  `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type FbpDraftDirectTplDlvCreateRsp struct {
	BaseRsp

	DraftId    int64  `json:"draft_id"`
	RowVersion int64  `json:"row_version"`
	SupplyId   string `json:"supply_id"`
}

type FbpDraftDirectTplDlvEditRsp struct {
	BaseRsp

	Errors     []string `json:"errors"`
	IsError    string   `json:"is_error"`
	RowVersion int64    `json:"row_version"`
}
