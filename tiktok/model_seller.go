package tiktok

type GetActiveShopsRsp struct {
	BaseRsp
	Data struct {
		Shops []struct {
			Id     string `json:"id"`
			Region string `json:"region"`
		} `json:"shops"`
	} `json:"data"`
}

type GetSellerPermissionsRsp struct {
	BaseRsp
	Data struct {
		Permissions []string `json:"permissions"`
	} `json:"data"`
}
