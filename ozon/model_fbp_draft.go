package ozon

import "time"

type FbpWarehouseListRsp struct {
	BaseRsp

	Warehouses []struct {
		AddressDetailing struct {
			City    string `json:"city"`
			Country string `json:"country"`
			House   string `json:"house"`
			Region  string `json:"region"`
			Street  string `json:"street"`
			Zipcode string `json:"zipcode"`
		} `json:"address_detailing"`
		Id           int64  `json:"id"`
		IsBonded     bool   `json:"is_bonded"`
		Name         string `json:"name"`
		PartnerName  string `json:"partner_name"`
		SupplyTypes  []int  `json:"supply_types"`
		TimezoneName string `json:"timezone_name"`
	} `json:"warehouses"`
}

type FbpDraftGetRsp struct {
	BaseRsp

	FbpDraftItem
}

type FbpDraftListRsp struct {
	BaseRsp

	HasNext bool           `json:"has_next"`
	Items   []FbpDraftItem `json:"items"`
	LastId  int            `json:"last_id"`
}

type FbpDraftItem struct {
	BundleId          string `json:"bundle_id"`
	CancellationState struct {
		CancellationError struct {
			ErrorCode string `json:"error_code"`
			Message   string `json:"message"`
		} `json:"cancellation_error"`
		CancellationStatus string `json:"cancellation_status"`
	} `json:"cancellation_state"`
	CreatedAt       time.Time `json:"created_at"`
	DeletedAt       time.Time `json:"deleted_at"`
	DeliveryDetails struct {
		DirectDetails struct {
			BySellerDetails struct {
				DriverName                string `json:"driver_name"`
				VehicleRegistrationNumber string `json:"vehicle_registration_number"`
				VehicleType               string `json:"vehicle_type"`
			} `json:"by_seller_details"`
			ByTplDetails struct {
				TrackingNumber       string `json:"tracking_number"`
				TransportCompanyName string `json:"transport_company_name"`
			} `json:"by_tpl_details"`
			TimeslotDetails struct {
				Timeslot struct {
					TimeslotEnd   time.Time `json:"timeslot_end"`
					TimeslotStart time.Time `json:"timeslot_start"`
				} `json:"timeslot"`
				TimeslotReservationId string `json:"timeslot_reservation_id"`
			} `json:"timeslot_details"`
		} `json:"direct_details"`
		DropOffPoint struct {
			Id           int64  `json:"id"`
			ProvinceUuid string `json:"province_uuid"`
			Timeslot     struct {
				TimeslotEnd   time.Time `json:"timeslot_end"`
				TimeslotStart time.Time `json:"timeslot_start"`
			} `json:"timeslot"`
		} `json:"drop_off_point"`
		PickupDetails struct {
			Address     string    `json:"address"`
			Comment     string    `json:"comment"`
			Date        time.Time `json:"date"`
			SenderName  string    `json:"sender_name"`
			SenderPhone string    `json:"sender_phone"`
		} `json:"pickup_details"`
		SupplyType string `json:"supply_type"`
	} `json:"delivery_details"`
	Editable          bool   `json:"editable"`
	Id                int64  `json:"id"`
	IsCancelable      bool   `json:"is_cancelable"`
	IsDeletable       bool   `json:"is_deletable"`
	Locked            bool   `json:"locked"`
	PackageUnitsCount int    `json:"package_units_count"`
	Status            string `json:"status"`
	SupplyId          string `json:"supply_id"`
	WarehouseId       int64  `json:"warehouse_id"`
}
