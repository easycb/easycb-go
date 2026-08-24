package ozon

import "time"

type WarehouseListRsp struct {
	BaseRsp

	Cursor     string              `json:"cursor"`
	Warehouses []WarehouseListItem `json:"warehouses"`
	HasNext    string              `json:"has_next"`
}

type WarehouseListItem struct {
	AddressInfo struct {
		Address   string  `json:"address"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Utc       string  `json:"utc"`
	} `json:"address_info"`
	CarriageLabelType string    `json:"carriage_label_type"`
	CourierComment    string    `json:"courier_comment"`
	CourierPhones     []string  `json:"courier_phones"`
	CreatedAt         time.Time `json:"created_at"`
	FirstMile         struct {
		Type                string `json:"type"`
		DropoffPointId      string `json:"dropoff_point_id"`
		TimeslotFrom        string `json:"timeslot_from"`
		TimeslotId          int    `json:"timeslot_id"`
		TimeslotTo          string `json:"timeslot_to"`
		FirstMileIsChanging bool   `json:"first_mile_is_changing"`
	} `json:"first_mile"`
	HasEntrustedAcceptance bool   `json:"has_entrusted_acceptance"`
	HasPostingsLimit       bool   `json:"has_postings_limit"`
	IsAutoAssembly         bool   `json:"is_auto_assembly"`
	IsKgt                  bool   `json:"is_kgt"`
	IsRfbs                 bool   `json:"is_rfbs"`
	IsWaybillEnabled       bool   `json:"is_waybill_enabled"`
	MinPostingsLimit       int    `json:"min_postings_limit"`
	IsComfort              bool   `json:"is_comfort"`
	IsExpress              bool   `json:"is_express"`
	WarehouseType          string `json:"warehouse_type"`
	CutInTime              int    `json:"cut_in_time"`
	Name                   string `json:"name"`
	Phone                  string `json:"phone"`
	PostingsLimit          int    `json:"postings_limit"`
	SlaCutIn               int    `json:"sla_cut_in"`
	Status                 string `json:"status"`
	Timetable              struct {
		TimetableFrom time.Time `json:"timetable_from"`
		TimetableTo   time.Time `json:"timetable_to"`
		WorkingHours  []struct {
			TimeFrom time.Time `json:"time_from"`
			TimeTo   time.Time `json:"time_to"`
		} `json:"working_hours"`
	} `json:"timetable"`
	UpdatedAt    time.Time `json:"updated_at"`
	WarehouseId  int64     `json:"warehouse_id"`
	WithItemList bool      `json:"with_item_list"`
	WorkingDays  []string  `json:"working_days"`
}

type DeliveryMethodListRsp struct {
	BaseRsp

	Cursor          string                                  `json:"cursor"`
	HasNext         bool                                    `json:"has_next"`
	DeliveryMethods []DeliveryMethodListDeliveryMethodsItem `json:"delivery_methods"`
}

type DeliveryMethodListDeliveryMethodsItem struct {
	CreatedAt       time.Time `json:"created_at"`
	Cutoff          string    `json:"cutoff"`
	Id              int64     `json:"id"`
	IsExpress       bool      `json:"is_express"`
	Name            string    `json:"name"`
	ProviderId      int64     `json:"provider_id"`
	SlaCutIn        int64     `json:"sla_cut_in"`
	Status          string    `json:"status"`
	TemplateId      int64     `json:"template_id"`
	TplDropoffPoint struct {
		Address            string `json:"address"`
		AddressCoordinates struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"address_coordinates"`
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"tpl_dropoff_point"`
	TplIntegrationType string    `json:"tpl_integration_type"`
	UpdatedAt          time.Time `json:"updated_at"`
	WarehouseId        int64     `json:"warehouse_id"`
}

type WarehouseOperationStatusRsp struct {
	BaseRsp
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
	Result struct {
		WarehouseId int64 `json:"warehouse_id"`
	} `json:"result"`
	Status string `json:"status"`
	Type   string `json:"type"`
}

type WarehouseArchiveRsp struct {
	BaseRsp

	OperationId string `json:"operation_id"`
}

type WarehouseUnArchiveRsp struct {
	BaseRsp

	OperationId string `json:"operation_id"`
}

type WarehouseInvalidProductsGetRsp struct {
	BaseRsp
	HasNext           bool                              `json:"has_next"`
	LastId            int64                             `json:"last_id"`
	ValidationResults []WarehouseInvalidProductsGetItem `json:"validation_results"`
	WarehouseId       int64                             `json:"warehouse_id"`
}

type WarehouseInvalidProductsGetItem struct {
	Item struct {
		Size struct {
			HeightMm float64 `json:"height_mm"`
			LengthMm float64 `json:"length_mm"`
			WidthMm  float64 `json:"width_mm"`
		} `json:"size"`
		Sku     int64   `json:"sku"`
		WeightG float64 `json:"weight_g"`
	} `json:"item"`
	State            string `json:"state"`
	ValidationErrors []struct {
		Characteristic   string `json:"characteristic"`
		RestrictionPrice struct {
			Currency string  `json:"currency"`
			Value    float64 `json:"value"`
		} `json:"restriction_price"`
		RestrictionVwc float64 `json:"restriction_vwc"`
		TemplateId     int     `json:"template_id"`
		Type           string  `json:"type"`
	} `json:"validation_errors"`
}

type WarehouseWarehouseWithInvalidProductsRsp struct {
	BaseRsp
	WarehouseIds []int64 `json:"warehouse_ids"`
}

type WarehouseFbsCreateDropOffListRsp struct {
	BaseRsp

	Points []struct {
		Address              string  `json:"address"`
		DiscountPercent      float64 `json:"discount_percent"`
		Id                   string  `json:"id"`
		LastTransitTimeLocal struct {
			Hours   int `json:"hours"`
			Minutes int `json:"minutes"`
			Nanos   int `json:"nanos"`
			Seconds int `json:"seconds"`
		} `json:"last_transit_time_local"`
		Coordinates struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"coordinates"`
		Type string `json:"type"`
	} `json:"points"`
}

type WarehouseFbsUpdateDropOffListRsp struct {
	BaseRsp

	Points []struct {
		Address              string `json:"address"`
		DiscountPercent      int64  `json:"discount_percent"`
		Id                   string `json:"id"`
		LastTransitTimeLocal struct {
			Hours   int `json:"hours"`
			Minutes int `json:"minutes"`
			Nanos   int `json:"nanos"`
			Seconds int `json:"seconds"`
		} `json:"last_transit_time_local"`
		Coordinates struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"coordinates"`
		Type string `json:"type"`
	} `json:"points"`
}

type WarehouseFbsCreateDropOffTimeslotListRsp struct {
	BaseRsp

	Timeslots []struct {
		AcceptanceEndTimeLocal   string `json:"acceptance_end_time_local"`
		AcceptanceStartTimeLocal string `json:"acceptance_start_time_local"`
		From                     string `json:"from"`
		Id                       int64  `json:"id"`
		To                       string `json:"to"`
	} `json:"timeslots"`
}

type WarehouseFbsUpdateDropOffTimeslotListRsp struct {
	BaseRsp

	Timeslots []struct {
		AcceptanceEndTimeLocal   string `json:"acceptance_end_time_local"`
		AcceptanceStartTimeLocal string `json:"acceptance_start_time_local"`
		From                     string `json:"from"`
		Id                       int64  `json:"id"`
		To                       string `json:"to"`
	} `json:"timeslots"`
}

type WarehouseFbsCreatePickupTimeslotListRsp struct {
	BaseRsp

	IsPickupSupported bool `json:"is_pickup_supported"`
	Timeslots         []struct {
		Id   int64  `json:"id"`
		From string `json:"from"`
		To   string `json:"to"`
	} `json:"timeslots"`
}

type WarehouseFbsUpdatePickupTimeslotListRsp struct {
	BaseRsp

	Timeslots []struct {
		Id   int64  `json:"id"`
		From string `json:"from"`
		To   string `json:"to"`
	} `json:"timeslots"`
}

type WarehouseFbsCreateRsp struct {
	BaseRsp

	OperationId string `json:"operation_id"`
}

type WarehouseFbsUpdateRsp struct {
	BaseRsp

	OperationId string `json:"operation_id"`
}

type WarehouseFbsFirstMileUpdateRsp struct {
	BaseRsp

	OperationId string `json:"operation_id"`
}
