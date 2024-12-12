package shopee

type GetShopsByPartnerRsp struct {
	BaseRsp
	AuthedShopList []struct {
		ShopId       int64  `json:"shop_id"`
		Region       string `json:"region"`
		SipAffiShops []struct {
			AffiShopId int64  `json:"affi_shop_id"`
			Region     string `json:"region"`
		} `json:"sip_affi_shops"`
		AuthTime   int64 `json:"auth_time"`
		ExpireTime int64 `json:"expire_time"`
	} `json:"authed_shop_list"`
	More bool `json:"more"`
}

type GetMerchantsByPartnerRsp struct {
	BaseRsp
	AuthedMerchantList []struct {
		Region     string `json:"region"`
		MerchantId int64  `json:"merchant_id"`
		AuthTime   int64  `json:"auth_time"`
		ExpireTime int64  `json:"expire_time"`
	} `json:"authed_merchant_list"`
	More bool `json:"more"`
}

type GetTokenByResendCodeRsp struct {
	BaseRsp
	ShopIdList   []int64 `json:"shop_id_list"`
	RefreshToken string  `json:"refresh_token"`
	AccessToken  string  `json:"access_token"`
	ExpireIn     int64   `json:"expire_in"`
}

type GetShopeeIpRangesRsp struct {
	BaseRsp
	IpList []string `json:"ip_list"`
}
