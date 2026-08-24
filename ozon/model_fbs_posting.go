package ozon

type FbsPostingProductExemplarSetRsp struct {
	BaseRsp
}

type FbsPostingProductExemplarCreateOrGetRsp struct {
	BaseRsp

	MultiBoxQty   int                                               `json:"multi_box_qty"`
	PostingNumber string                                            `json:"posting_number"`
	Products      []FbsPostingProductExemplarCreateOrGetProductItem `json:"products"`
}

type FbsPostingProductExemplarCreateOrGetProductItem struct {
	Exemplars []struct {
		ExemplarId   int64  `json:"exemplar_id"`
		Gtd          string `json:"gtd"`
		IsGtdAbsent  bool   `json:"is_gtd_absent"`
		IsRnptAbsent bool   `json:"is_rnpt_absent"`
		Marks        []struct {
			Mark     string `json:"mark"`
			MarkType string `json:"mark_type"`
		} `json:"marks"`
		Rnpt string `json:"rnpt"`
	} `json:"exemplars"`
	HasImei                 bool  `json:"has_imei"`
	IsGtdNeeded             bool  `json:"is_gtd_needed"`
	IsJwUinNeeded           bool  `json:"is_jw_uin_needed"`
	IsMandatoryMarkNeeded   bool  `json:"is_mandatory_mark_needed"`
	IsMandatoryMarkPossible bool  `json:"is_mandatory_mark_possible"`
	IsRnptNeeded            bool  `json:"is_rnpt_needed"`
	ProductId               int64 `json:"product_id"`
	Quantity                int   `json:"quantity"`
}

type FbsPostingProductExemplarStatusRsp struct {
	BaseRsp

	PostingNumber string                                       `json:"posting_number"`
	Products      []FbsPostingProductExemplarStatusProductItem `json:"products"`
	Status        string                                       `json:"status"`
}

type FbsPostingProductExemplarStatusProductItem struct {
	Exemplars []struct {
		ExemplarId     int64    `json:"exemplar_id"`
		Gtd            string   `json:"gtd"`
		GtdCheckStatus string   `json:"gtd_check_status"`
		GtdErrorCodes  []string `json:"gtd_error_codes"`
		IsGtdAbsent    bool     `json:"is_gtd_absent"`
		IsRnptAbsent   bool     `json:"is_rnpt_absent"`
		Marks          []struct {
			CheckStatus string   `json:"check_status"`
			ErrorCodes  []string `json:"error_codes"`
			Mark        string   `json:"mark"`
			MarkType    string   `json:"mark_type"`
		} `json:"marks"`
		Rnpt            string   `json:"rnpt"`
		RnptCheckStatus string   `json:"rnpt_check_status"`
		RnptErrorCodes  []string `json:"rnpt_error_codes"`
	} `json:"exemplars"`
	ProductId int64 `json:"product_id"`
}

type FbsPostingProductExemplarValidateRsp struct {
	BaseRsp

	Products []FbsPostingProductExemplarValidateProductItem `json:"products"`
}

type FbsPostingProductExemplarValidateProductItem struct {
	Error     string `json:"error"`
	Exemplars []struct {
		Errors []string `json:"errors"`
		Gtd    string   `json:"gtd"`
		Marks  []struct {
			Errors   []string `json:"errors"`
			Mark     string   `json:"mark"`
			MarkType string   `json:"mark_type"`
			Valid    bool     `json:"valid"`
		} `json:"marks"`
		Rnpt  string `json:"rnpt"`
		Valid bool   `json:"valid"`
	} `json:"exemplars"`
	ProductId int64 `json:"product_id"`
	Valid     bool  `json:"valid"`
}

type FbsPostingProductExemplarUpdateRsp struct {
	BaseRsp
}
