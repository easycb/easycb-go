package ozon

type Client struct {
	BaseURL string

	ClientId string
	ApiKey   string

	conf *config
}

func NewClient(clientId string, apiKey string, baseUrl string, configurers ...configurer) (*Client, error) {
	conf := &config{}
	for _, configurer := range configurers {
		configurer(conf)
	}

	if err := conf.initConfig(); err != nil {
		return nil, err
	}

	return &Client{
		BaseURL:  baseUrl,
		ClientId: clientId,
		ApiKey:   apiKey,
		conf:     conf,
	}, nil
}

func NewClientDefault(configurers ...configurer) (*Client, error) {
	conf := &config{}
	for _, configurer := range configurers {
		configurer(conf)
	}

	if err := conf.initConfig(); err != nil {
		return nil, err
	}

	return &Client{
		conf: conf,
	}, nil
}

func (c *Client) SetApiKey(apiKey string) *Client {
	c.ApiKey = apiKey
	return c
}

func (c *Client) SetClientId(clientId string) *Client {
	c.ClientId = clientId
	return c
}

func (c *Client) SetBaseurl(s string) *Client {
	c.BaseURL = s
	return c
}
