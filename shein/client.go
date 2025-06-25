package shein

import "github.com/easycb/easycb-go"

type Client struct {
	BaseURL    string
	WebHookURL string

	AppId        string
	AppSecretKey string

	OpenKeyId string
	SecretKey string

	conf *config
}

func NewClient(appId string, openKeyId string, secretKey string, baseUrl string, configurers ...configurer) (*Client, error) {
	if appId == "" || openKeyId == "" || secretKey == "" {
		return nil, easycb.MissSheinInitParamErr
	}

	conf := &config{}
	for _, configurer := range configurers {
		configurer(conf)
	}

	if err := conf.initConfig(); err != nil {
		return nil, err
	}

	return &Client{
		AppId:     appId,
		OpenKeyId: openKeyId,
		SecretKey: secretKey,
		BaseURL:   baseUrl,
		conf:      conf,
	}, nil
}

func NewClientDefault(appId string, appSecretKey string, configurers ...configurer) (*Client, error) {
	if appId == "" || appSecretKey == "" {
		return nil, easycb.MissSheinInitParamErr
	}

	conf := &config{}
	for _, configurer := range configurers {
		configurer(conf)
	}

	if err := conf.initConfig(); err != nil {
		return nil, err
	}

	return &Client{
		AppId:        appId,
		AppSecretKey: appSecretKey,
		conf:         conf,
	}, nil
}

func (c *Client) SetAppId(appId string) *Client {
	c.AppId = appId
	return c
}

func (c *Client) SetOpenKeyId(openKeyId string) *Client {
	c.OpenKeyId = openKeyId
	return c
}

func (c *Client) SetSecretKey(secretKey string) *Client {
	c.SecretKey = secretKey
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
