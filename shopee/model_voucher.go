package shopee

type GetVoucherListRsp struct {
	BaseRsp
	Response struct {
		More        bool `json:"more"`
		VoucherList []struct {
			VoucherId        int64   `json:"voucher_id"`
			VoucherCode      string  `json:"voucher_code"`
			VoucherName      string  `json:"voucher_name"`
			VoucherType      int     `json:"voucher_type"`
			RewardType       int     `json:"reward_type"`
			UsageQuantity    int     `json:"usage_quantity"`
			CurrentUsage     int     `json:"current_usage"`
			StartTime        int64   `json:"start_time"`
			EndTime          int64   `json:"end_time"`
			IsAdmin          bool    `json:"is_admin"`
			VoucherPurpose   int     `json:"voucher_purpose"`
			DiscountAmount   float64 `json:"discount_amount"`
			TargetVoucher    int     `json:"target_voucher"`
			DisplayStartTime int64   `json:"display_start_time"`
			Usecase          int     `json:"usecase"`
			Percentage       int     `json:"percentage"`
		} `json:"voucher_list"`
	} `json:"response"`
}

type GetVoucherRsp struct {
	BaseRsp
	Response struct {
		VoucherId          int64         `json:"voucher_id"`
		VoucherCode        string        `json:"voucher_code"`
		VoucherName        string        `json:"voucher_name"`
		VoucherType        int           `json:"voucher_type"`
		RewardType         int           `json:"reward_type"`
		UsageQuantity      int           `json:"usage_quantity"`
		CurrentUsage       int           `json:"current_usage"`
		StartTime          int64         `json:"start_time"`
		EndTime            int64         `json:"end_time"`
		IsAdmin            bool          `json:"is_admin"`
		VoucherPurpose     int           `json:"voucher_purpose"`
		DiscountAmount     float64       `json:"discount_amount"`
		CmtVoucherStatus   int           `json:"cmt_voucher_status"`
		MinBasketPrice     float64       `json:"min_basket_price"`
		DisplayChannelList []interface{} `json:"display_channel_list"`
		DisplayStartTime   int64         `json:"display_start_time"`
	} `json:"response"`
}

type UpdateVoucherRsp struct {
	BaseRsp
	Response struct {
		VoucherId int64 `json:"voucher_id"`
	} `json:"response"`
}

type EndVoucherRsp struct {
	BaseRsp
	Response struct {
		VoucherId int64 `json:"voucher_id"`
	} `json:"response"`
}

type AddVoucherRsp struct {
	BaseRsp
	Response struct {
		VoucherId int64 `json:"voucher_id"`
	} `json:"response"`
}

type DeleteVoucherRsp struct {
	BaseRsp
	Response struct {
		VoucherId int64 `json:"voucher_id"`
	} `json:"response"`
}
