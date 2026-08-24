package ozon

import "time"

type FbpDraftDropOffProvinceListRsp struct {
	BaseRsp

	Provinces []struct {
		Name         string `json:"name"`
		PointsCount  int    `json:"points_count"`
		ProvinceUuid string `json:"province_uuid"`
	} `json:"provinces"`
}

type FbpDraftDropOffPointListRsp struct {
	BaseRsp

	DropOffPoints []struct {
		City               string    `json:"city"`
		DropOffPointId     int64     `json:"drop_off_point_id"`
		NearestDropOffDate time.Time `json:"nearest_drop_off_date"`
		PointAddress       string    `json:"point_address"`
		ProvinceUuid       string    `json:"province_uuid"`
	} `json:"drop_off_points"`
}

type FbpDraftDropOffPointTimetableRsp struct {
	BaseRsp

	Calendar []struct {
		CalendarItem struct {
			BreakHours struct {
				TimeslotEnd   string `json:"timeslot_end"`
				TimeslotStart string `json:"timeslot_start"`
			} `json:"break_hours"`
			IsHoliday    bool `json:"is_holiday"`
			OpeningHours struct {
				TimeslotEnd   string `json:"timeslot_end"`
				TimeslotStart string `json:"timeslot_start"`
			} `json:"opening_hours"`
		} `json:"calendar_item"`
		DayOfWeek string `json:"day_of_week"`
	} `json:"calendar"`
}

type FbpDraftDropOffProductValidateRsp struct {
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
		Barcode          string   `json:"barcode"`
		IconName         string   `json:"icon_name"`
		Name             string   `json:"name"`
		OfferId          string   `json:"offer_id"`
		Quantity         int      `json:"quantity"`
		RejectionReasons []string `json:"rejection_reasons"`
		Sku              int64    `json:"sku"`
		Volume           float64  `json:"volume"`
	} `json:"rejected_items"`
}

type FbpDraftDropOffCreateRsp struct {
	BaseRsp

	DraftId    int64  `json:"draft_id"`
	RowVersion int64  `json:"row_version"`
	SupplyId   string `json:"supply_id"`
}

type FbpDraftDropOffDeleteRsp struct {
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

type FbpDraftDropOffDlvEditRsp struct {
	BaseRsp

	RowVersion int64 `json:"row_version"`
}

type FbpDraftDropOffRegistrateRsp struct {
	BaseRsp

	Error struct {
		BundleErrors []struct {
			Errors []string `json:"errors"`
			Sku    int64    `json:"sku"`
		} `json:"bundle_errors"`
		OrderError string `json:"order_error"`
	} `json:"error"`
	IsError    bool  `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}
