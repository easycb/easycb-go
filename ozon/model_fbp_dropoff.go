package ozon

type FbpOrderDropOffCancelRsp struct {
	BaseRsp

	Error struct {
		OrderErrors []string `json:"order_errors"`
	} `json:"error"`
	IsError    bool  `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type FbpOrderDropOffDlvEditRsp struct {
	BaseRsp

	RowVersion int64 `json:"row_version"`
}

type FbpOrderDropOffTimetableRsp struct {
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
