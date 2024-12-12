package shopee

import (
	"github.com/easycb/easycb-go"
)

func (c *Client) GetShopsByPartner(query easycb.AnyMap) (*GetShopsByPartnerRsp, error) {
	var result GetShopsByPartnerRsp
	err := c.doRequest("GET", "/api/v2/public/get_shops_by_partner", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetMerchantsByPartner(query easycb.AnyMap) (*GetMerchantsByPartnerRsp, error) {
	var result GetMerchantsByPartnerRsp
	err := c.doRequest("GET", "/api/v2/public/get_merchants_by_partner", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetTokenByResendCode(body easycb.AnyMap) (*GetTokenByResendCodeRsp, error) {
	var result GetTokenByResendCodeRsp
	err := c.doRequest("POST", "/api/v2/public/get_token_by_resend_code", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetShopeeIpRanges() (*GetShopeeIpRangesRsp, error) {
	var result GetShopeeIpRangesRsp
	err := c.doRequest("GET", "/api/v2/public/get_shopee_ip_ranges", nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
