package shopee

import "github.com/easycb/easycb-go"

type Client struct {
	PartnerId  int64
	PartnerKey string
	BaseURL    string
	WebHookURL string

	ShopID      int64
	MerchantID  int64
	AccessToken string

	conf *config
}

func NewClient(partnerId int64, partnerKey, baseUrl string, configurers ...configurer) (*Client, error) {
	if partnerId == 0 || partnerKey == "" || baseUrl == "" {
		return nil, easycb.MissShopeeInitParamErr
	}

	conf := &config{}
	for _, configurer := range configurers {
		configurer(conf)
	}

	if err := conf.initConfig(); err != nil {
		return nil, err
	}

	return &Client{
		PartnerId:  partnerId,
		PartnerKey: partnerKey,
		BaseURL:    baseUrl,
		conf:       conf,
	}, nil
}

func NewClientDefault(partnerId int64, partnerKey string, configurers ...configurer) (*Client, error) {
	if partnerId == 0 || partnerKey == "" {
		return nil, easycb.MissShopeeInitParamErr
	}

	conf := &config{}
	for _, configurer := range configurers {
		configurer(conf)
	}

	if err := conf.initConfig(); err != nil {
		return nil, err
	}

	return &Client{
		PartnerId:  partnerId,
		PartnerKey: partnerKey,
		conf:       conf,
	}, nil
}

func (c *Client) SetShopId(n int64) *Client {
	c.ShopID = n
	return c
}

func (c *Client) SetMerchantId(n int64) *Client {
	c.MerchantID = n
	return c
}

func (c *Client) SetAccessToken(s string) *Client {
	c.AccessToken = s
	return c
}

func (c *Client) SetWebHookURL(s string) *Client {
	c.WebHookURL = s
	return c
}

func (c *Client) SetBaseurl(s string) *Client {
	c.BaseURL = s
	return c
}
