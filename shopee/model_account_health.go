package shopee

type ShopPenaltyRsp struct {
	BaseRsp
	Response struct {
		PenaltyPoints struct {
			OverallPenaltyPoints int `json:"overall_penalty_points"`
			NonFulfillmentRate   int `json:"non_fulfillment_rate"`
			LateShipmentRate     int `json:"late_shipment_rate"`
			ListingViolations    int `json:"listing_violations"`
			Others               int `json:"others"`
		} `json:"penalty_points"`
		OngoingPunishment []struct {
			PunishmentName string `json:"punishment_name"`
			PunishmentTier int    `json:"punishment_tier"`
			DaysLeft       int    `json:"days_left"`
		} `json:"ongoing_punishment"`
	} `json:"responce"`
}

type GetShopPerformanceRsp struct {
	BaseRsp
	Response struct {
		MetricList []struct {
			CurrentPeriod  float64 `json:"current_period"`
			LastPeriod     float64 `json:"last_period"`
			MetricId       int64   `json:"metric_id"`
			MetricName     string  `json:"metric_name"`
			MetricType     int     `json:"metric_type"`
			ParentMetricId int64   `json:"parent_metric_id"`
			Target         struct {
				Comparator string  `json:"comparator"`
				Value      float64 `json:"value"`
			} `json:"target"`
			Unit int `json:"unit"`
		} `json:"metric_list"`
		OverallPerformance struct {
			CustomServiceFailed int `json:"custom_service_failed"`
			FulfillmentFailed   int `json:"fulfillment_failed"`
			ListingFailed       int `json:"listing_failed"`
			Rating              int `json:"rating"`
		} `json:"overall_performance"`
	} `json:"response"`
}

type GetMetricSourceDetailRsp struct {
	BaseRsp
	Response struct {
		NfrOrderList []struct {
			OrderSn            string `json:"order_sn"`
			NonFulfillmentType int    `json:"non_fulfillment_type"`
			DetailedReason     int    `json:"detailed_reason"`
		} `json:"nfr_order_list"`
		CancellationOrderList []struct {
			OrderSn          string `json:"order_sn"`
			CancellationType int    `json:"cancellation_type"`
			DetailedReason   int    `json:"detailed_reason"`
		} `json:"cancellation_order_list"`
		ReturnRefundOrderList []struct {
			OrderSn        string `json:"order_sn"`
			DetailedReason int    `json:"detailed_reason"`
		} `json:"return_refund_order_list"`
		LsrOrderList []struct {
			ActualShippingTime int64  `json:"actual_shipping_time"`
			LateByDays         int64  `json:"late_by_days"`
			OrderSn            string `json:"order_sn"`
			ShippingDeadline   int64  `json:"shipping_deadline"`
		} `json:"lsr_order_list"`
		FhrOrderList []struct {
			OrderSn          string `json:"order_sn"`
			ParcelId         int64  `json:"parcel_id"`
			ConfirmTime      int64  `json:"confirm_time"`
			HandoverTime     int64  `json:"handover_time"`
			HandoverDeadLine int64  `json:"handover_deadline"`
		} `json:"fhr_order_list"`
		OpfrDayDetailDataList []struct {
			Date               string `json:"date"`
			ScheduledPickupNum int    `json:"scheduled_pickup_num"`
			FailedPickupNum    int    `json:"failed_pickup_num"`
			Opfr               int    `json:"opfr"`
			Target             string `json:"target"`
		} `json:"opfr_day_detail_data_list"`
		ViolationListingList []struct {
			ItemId         int64 `json:"item_id"`
			DetailedReason int   `json:"detailed_reason"`
			UpdateTime     int64 `json:"update_time"`
		} `json:"violation_listing_list"`
		PreOrderListingViolationDataList []struct {
			Date                 string `json:"date"`
			LiveListingCount     int64  `json:"live_listing_count"`
			PreOrderListingCount int    `json:"pre_order_listing_count"`
			PreOrderListingRate  int    `json:"pre_order_listing_rate"`
			Target               string `json:"target"`
		} `json:"pre_order_listing_violation_data_list"`
		PreOrderListingList []struct {
			ItemId                int64 `json:"item_id"`
			CurrentPreOrderStatus int   `json:"current_pre_order_status"`
		} `json:"pre_order_listing_list"`
		MetricsId  int64 `json:"metrics_id"`
		TotalCount int   `json:"total_count"`
	} `json:"response"`
}

type GetPenaltyPointHistoryRsp struct {
	BaseRsp
	Response struct {
		PenaltyPointList []struct {
			IssueTime        int   `json:"issue_time"`
			LatestPointNum   int   `json:"latest_point_num"`
			OriginalPointNum int   `json:"original_point_num"`
			ReferenceId      int64 `json:"reference_id"`
			ViolationType    int   `json:"violation_type"`
		} `json:"penalty_point_list"`
		TotalCount int `json:"total_count"`
	} `json:"response"`
}

type GetPunishmentHistoryRsp struct {
	BaseRsp
	Response struct {
		PunishmentList []struct {
			EndTime        int64 `json:"end_time"`
			IssueTime      int64 `json:"issue_time"`
			PunishmentType int   `json:"punishment_type"`
			Reason         int   `json:"reason"`
			ReferenceId    int64 `json:"reference_id"`
			StartTime      int64 `json:"start_time"`
		} `json:"punishment_list"`
		TotalCount int `json:"total_count"`
	} `json:"response"`
}

type GetListingsWithIssuesRsp struct {
	BaseRsp
	Response struct {
		ListingList []struct {
			ItemId int64 `json:"item_id"`
			Reason int   `json:"reason"`
		} `json:"listing_list"`
		TotalCount int `json:"total_count"`
	} `json:"response"`
}

type GetLateOrdersRsp struct {
	BaseRsp
	Response struct {
		LateOrderList []struct {
			LateByDays       int64  `json:"late_by_days"`
			OrderSn          string `json:"order_sn"`
			ShippingDeadline int64  `json:"shipping_deadline"`
		} `json:"late_order_list"`
		TotalCount int `json:"total_count"`
	} `json:"response"`
}
