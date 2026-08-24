package ozon

type FbsPostingDeliveringRsp struct {
	BaseRsp
	Result []struct {
		Error         string `json:"error"`
		PostingNumber string `json:"posting_number"`
		Result        bool   `json:"result"`
	} `json:"result"`
}

type FbsPostingTrackingNumberSetRsp struct {
	BaseRsp

	Result []struct {
		Error         string `json:"error"`
		PostingNumber string `json:"posting_number"`
		Result        bool   `json:"result"`
	} `json:"result"`
}

type FbsPostingLastMileRsp struct {
	BaseRsp

	Result []struct {
		Error         []interface{} `json:"error"`
		PostingNumber string        `json:"posting_number"`
		Result        bool          `json:"result"`
	} `json:"result"`
}

type FbsPostingDeliveredRsp struct {
	BaseRsp

	Result []struct {
		Error         []interface{} `json:"error"`
		PostingNumber string        `json:"posting_number"`
		Result        bool          `json:"result"`
	} `json:"result"`
}

type PostingCutoffSetRsp struct {
	BaseRsp

	Result bool `json:"result"`
}
