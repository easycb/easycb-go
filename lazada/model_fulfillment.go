package lazada

type ConfirmDeliveryForDBSRsp struct {
	Result struct {
		ErrorMsg string `json:"error_msg"`
		Data     struct {
			Packages []struct {
				Msg         string `json:"msg"`
				ItemErrCode string `json:"item_err_code"`
				PackageId   string `json:"package_id"`
				Retry       bool   `json:"retry"`
			} `json:"packages"`
		} `json:"data"`
		Success   bool   `json:"success"`
		ErrorCode string `json:"error_code"`
	} `json:"result"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type DeliverDigitalRsp struct {
	Result struct {
		Data struct {
			Orders []struct {
				OrderItemList []struct {
					Msg         string `json:"msg"`
					OrderItemId int64  `json:"order_item_id"`
					ItemErrCode string `json:"item_err_code"`
					Retry       bool   `json:"retry"`
				} `json:"order_item_list"`
				OrderId int64 `json:"order_id"`
			} `json:"orders"`
		} `json:"data"`
		Success   bool   `json:"success"`
		ErrorCode string `json:"errorCode"`
		ErrorMsg  string `json:"errorMsg"`
	} `json:"result"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type FailedDeliveryForDBSRsp struct {
	Result struct {
		ErrorMsg string `json:"error_msg"`
		Data     struct {
			Packages []struct {
				Msg         string `json:"msg"`
				ItemErrCode string `json:"item_err_code"`
				PackageId   string `json:"package_id"`
				Retry       bool   `json:"retry"`
			} `json:"packages"`
		} `json:"data"`
		Success   bool   `json:"success"`
		ErrorCode string `json:"error_code"`
	} `json:"result"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type GetShipmentProviderRsp struct {
	Result struct {
		ErrorMsg string `json:"error_msg"`
		Data     struct {
			PlatformDefault   int `json:"platform_default"`
			ShipmentProviders []struct {
				Name         string `json:"name"`
				ProviderCode string `json:"provider_code"`
			} `json:"shipment_providers"`
			ShippingAllocateType string `json:"shipping_allocate_type"`
		} `json:"data"`
		Success   bool   `json:"success"`
		ErrorCode string `json:"error_code"`
	} `json:"result"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type PackRsp struct {
	Result struct {
		ErrorMsg string `json:"error_msg"`
		Data     struct {
			PackOrderList []struct {
				OrderItemList []struct {
					OrderItemId      int64  `json:"order_item_id"`
					Msg              string `json:"msg"`
					ItemErrCode      string `json:"item_err_code"`
					TrackingNumber   string `json:"tracking_number"`
					ShipmentProvider string `json:"shipment_provider"`
					PackageId        string `json:"package_id"`
					Retry            bool   `json:"retry"`
				} `json:"order_item_list"`
				OrderId int64 `json:"order_id"`
			} `json:"pack_order_list"`
		} `json:"data"`
		Success   bool   `json:"success"`
		ErrorCode string `json:"error_code"`
	} `json:"result"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type PackageStatusUpdateForDBSRsp struct {
	Code    string `json:"code"`
	Success bool   `json:"success"`
	Module  struct {
		Result bool `json:"result"`
	} `json:"module"`
	ErrorCode struct {
		DisplayMessage string `json:"displayMessage"`
	} `json:"errorCode"`
	RequestId string `json:"request_id"`
}

type PrintAWBRsp struct {
	Result struct {
		ErrorMsg string `json:"error_msg"`
		Data     struct {
			File    string `json:"file"`
			PdfUrl  string `json:"pdf_url"`
			DocType string `json:"doc_type"`
		} `json:"data"`
		Success   bool   `json:"success"`
		ErrorCode string `json:"error_code"`
	} `json:"result"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type ReadyToShipRsp struct {
	Result struct {
		ErrorMsg string `json:"error_msg"`
		Data     struct {
			Packages []struct {
				Msg         string `json:"msg"`
				ItemErrCode string `json:"item_err_code"`
				PackageId   string `json:"package_id"`
				Retry       bool   `json:"retry"`
			} `json:"packages"`
		} `json:"data"`
		Success   bool   `json:"success"`
		ErrorCode string `json:"error_code"`
	} `json:"result"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type RecreatePackageRsp struct {
	Result struct {
		ErrorMsg string `json:"error_msg"`
		Data     struct {
			Packages []struct {
				Msg         string `json:"msg"`
				ItemErrCode string `json:"item_err_code"`
				PackageId   string `json:"package_id"`
				Retry       bool   `json:"retry"`
			} `json:"packages"`
		} `json:"data"`
		Success   bool   `json:"success"`
		ErrorCode string `json:"error_code"`
	} `json:"result"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}
