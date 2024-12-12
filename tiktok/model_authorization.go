package tiktok

type GetAccessTokenRsp struct {
	BaseRsp
	Data struct {
		AccessToken          string `json:"access_token"`
		AccessTokenExpireIn  int64  `json:"access_token_expire_in"`
		RefreshToken         string `json:"refresh_token"`
		RefreshTokenExpireIn int64  `json:"refresh_token_expire_in"`
		OpenId               string `json:"open_id"`
		SellerName           string `json:"seller_name"`
		SellerBaseRegion     string `json:"seller_base_region"`
		UserType             int    `json:"user_type"`
	} `json:"data"`
}

type RefreshAccessTokenRsp struct {
	BaseRsp
	Data struct {
		AccessToken          string `json:"access_token"`
		AccessTokenExpireIn  int64  `json:"access_token_expire_in"`
		RefreshToken         string `json:"refresh_token"`
		RefreshTokenExpireIn int64  `json:"refresh_token_expire_in"`
		OpenId               string `json:"open_id"`
		SellerName           string `json:"seller_name"`
		SellerBaseRegion     string `json:"seller_base_region"`
		UserType             int    `json:"user_type"`
	} `json:"data"`
}

type GetAuthorizedShopRsp struct {
	BaseRsp
	Data struct {
		Shops []struct {
			Cipher     string `json:"cipher"`
			Code       string `json:"code"`
			Id         string `json:"id"`
			Name       string `json:"name"`
			Region     string `json:"region"`
			SellerType string `json:"seller_type"`
		} `json:"shops"`
	} `json:"data"`
}

type GetAuthorizedCategoryAssetsRsp struct {
	BaseRsp
	Data struct {
		CategoryAssets []struct {
			Category struct {
				Id   int64  `json:"id"`
				Name string `json:"name"`
			} `json:"category"`
			Cipher       string `json:"cipher"`
			TargetMarket string `json:"target_market"`
		} `json:"category_assets"`
	} `json:"data"`
}
