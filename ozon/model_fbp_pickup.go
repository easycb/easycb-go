package ozon

type FbpOrderPickUpCancelRsp struct {
	BaseRsp

	Error struct {
		OrderErrors string `json:"order_errors"`
	} `json:"error"`
	IsError    bool  `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type FbpOrderPickUpDlvEditRsp struct {
	BaseRsp

	Error struct {
		OrderErrors string `json:"order_errors"`
	} `json:"error"`
	IsError    bool  `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}
