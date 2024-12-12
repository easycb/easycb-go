package tiktok

import (
	"github.com/easycb/easycb-go"
)

type Client struct {
	AppKey    string
	AppSecret string
	BaseURL   string

	ShopCipher  string
	AccessToken string
	conf        *config
}

func NewClient(appKey, appSecret, baseUrl string, configurers ...configurer) (*Client, error) {
	if appKey == "" || appSecret == "" || baseUrl == "" {
		return nil, easycb.MissTiktokInitParamErr
	}

	conf := &config{}
	for _, configurer := range configurers {
		configurer(conf)
	}

	if err := conf.initConfig(); err != nil {
		return nil, err
	}

	return &Client{
		AppKey:    appKey,
		AppSecret: appSecret,
		BaseURL:   baseUrl,
		conf:      conf,
	}, nil
}

func NewClientDefault(appKey, appSecret string, configurers ...configurer) (*Client, error) {
	if appKey == "" || appSecret == "" {
		return nil, easycb.MissTiktokInitParamErr
	}

	conf := &config{}
	for _, configurer := range configurers {
		configurer(conf)
	}

	if err := conf.initConfig(); err != nil {
		return nil, err
	}

	return &Client{
		AppKey:    appKey,
		AppSecret: appSecret,
		conf:      conf,
	}, nil
}

func (c *Client) SetShopCipher(s string) *Client {
	c.ShopCipher = s
	return c
}

func (c *Client) SetAccessToken(s string) *Client {
	c.AccessToken = s
	return c
}

func (c *Client) SetBaseUrl(s string) *Client {
	c.BaseURL = s
	return c
}
