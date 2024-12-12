package lazada

type RefreshAccessTokenRsp struct {
	AccessToken         string `json:"access_token"`
	Country             string `json:"country"`
	RefreshToken        string `json:"refresh_token"`
	AccountId           string `json:"account_id"`
	Code                string `json:"code"`
	CountryUserInfoList []struct {
		Country   string `json:"country"`
		UserId    string `json:"user_id"`
		SellerId  string `json:"seller_id"`
		ShortCode string `json:"short_code"`
	} `json:"country_user_info_list"`
	AccountPlatform  string `json:"account_platform"`
	RefreshExpiresIn int64  `json:"refresh_expires_in"`
	ExpiresIn        int64  `json:"expires_in"`
	RequestId        string `json:"request_id"`
	Account          string `json:"account"`
}

type GetAccessTokenRsp struct {
	AccessToken      string `json:"access_token"`
	Country          string `json:"country"`
	RefreshToken     string `json:"refresh_token"`
	AccountPlatform  string `json:"account_platform"`
	RefreshExpiresIn int64  `json:"refresh_expires_in"`
	CountryUserInfo  []struct {
		Country   string `json:"country"`
		UserId    string `json:"user_id"`
		SellerId  string `json:"seller_id"`
		ShortCode string `json:"short_code"`
	} `json:"country_user_info"`
	ExpiresIn int64  `json:"expires_in"`
	Account   string `json:"account"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type GenerateAccessTokenWithOpenIdRsp struct {
	AccessToken      string `json:"access_token"`
	Country          string `json:"country"`
	RefreshToken     string `json:"refresh_token"`
	AccountId        string `json:"account_id"`
	Code             string `json:"code"`
	AccountPlatform  string `json:"account_platform"`
	RefreshExpiresIn string `json:"refresh_expires_in"`
	CountryUserInfo  []struct {
		Country   string `json:"country"`
		UserId    string `json:"user_id"`
		SellerId  string `json:"seller_id"`
		ShortCode string `json:"short_code"`
	} `json:"country_user_info"`
	ExpiresIn string `json:"expires_in"`
	RequestId string `json:"request_id"`
	Account   string `json:"account"`
}
