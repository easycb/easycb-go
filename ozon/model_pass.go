package ozon

import "time"

type PassListRsp struct {
	BaseRsp
	ArrivalPasses []PassListArrivalPassItem `json:"arrival_passes"`
	Cursor        string                    `json:"cursor"`
}

type PassListArrivalPassItem struct {
	ArrivalPassId       int64     `json:"arrival_pass_id"`
	ArrivalReasons      []string  `json:"arrival_reasons"`
	ArrivalTime         time.Time `json:"arrival_time"`
	DriverName          string    `json:"driver_name"`
	DriverPhone         string    `json:"driver_phone"`
	DropoffPointId      int64     `json:"dropoff_point_id"`
	IsActive            bool      `json:"is_active"`
	VehicleLicensePlate string    `json:"vehicle_license_plate"`
	VehicleModel        string    `json:"vehicle_model"`
	WarehouseId         int64     `json:"warehouse_id"`
}

type CarriagePassCreateRsp struct {
	BaseRsp
	ArrivalPassIds []int64 `json:"arrival_pass_ids"`
}

type CarriagePassUpdateRsp struct {
	BaseRsp
}

type CarriagePassDeleteRsp struct {
	BaseRsp
}

type ReturnPassCreateRsp struct {
	BaseRsp
	ArrivalPassIds []int64 `json:"arrival_pass_ids"`
}

type ReturnPassUpdateRsp struct {
	BaseRsp
}

type ReturnPassDeleteRsp struct {
	BaseRsp
}

type ReturnsCompanyFbsInfoRsp struct {
	BaseRsp

	DropOffPoints []ReturnsCompanyFbsInfoDropOffPointsItem `json:"drop_off_points"`
	HasNext       bool                                     `json:"has_next"`
}

type ReturnsCompanyFbsInfoDropOffPointsItem struct {
	Address  string `json:"address"`
	BoxCount int    `json:"box_count"`
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	PassInfo struct {
		Count      int  `json:"count"`
		IsRequired bool `json:"is_required"`
	} `json:"pass_info"`
	PlaceId       int64   `json:"place_id"`
	ReturnsCount  int     `json:"returns_count"`
	UtcOffset     string  `json:"utc_offset"`
	WarehousesIds []int64 `json:"warehouses_ids"`
}
