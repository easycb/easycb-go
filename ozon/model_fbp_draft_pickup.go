package ozon

type FbpDraftPickUpCreateRsp struct {
	BaseRsp

	DraftId    int64  `json:"draft_id"`
	RowVersion int64  `json:"row_version"`
	SupplyId   string `json:"supply_id"`
}

type FbpDraftPickUpDeleteRsp struct {
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

type FbpDraftPickUpDlvEditRsp struct {
	BaseRsp

	RowVersion int64 `json:"row_version"`
}

type FbpDraftPickUpProductValidateRsp struct {
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

type FbpDraftPickUpRegistrateRsp struct {
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
