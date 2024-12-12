package shopee

type RefreshAccessTokenRsp struct {
	BaseRsp
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpireIn     int64  `json:"expire_in"`
	PartnerID    int64  `json:"partner_id"`
	MerchantID   int64  `json:"merchant_id"`
	ShopID       int64  `json:"shop_id"`
}

type GetAccessTokenRsp struct {
	BaseRsp
	RefreshToken   string  `json:"refresh_token"`
	AccessToken    string  `json:"access_token"`
	ExpireIn       int     `json:"expire_in"`
	MerchantIdList []int64 `json:"merchant_id_list"`
	ShopIdList     []int64 `json:"shop_id_list"`
}
