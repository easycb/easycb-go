package tiktok

import "github.com/easycb/easycb-go"

func (c *Client) GetAccessToken(query easycb.AnyMap) (*GetAccessTokenRsp, error) {
	var result GetAccessTokenRsp
	path := "/api/v2/token/get"
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) RefreshAccessToken(query easycb.AnyMap) (*RefreshAccessTokenRsp, error) {
	var result RefreshAccessTokenRsp
	path := "/api/v2/token/refresh"
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetAuthorizedShops() (*GetAuthorizedShopRsp, error) {
	var result GetAuthorizedShopRsp
	path := "/authorization/202309/shops"
	err := c.doRequest("GET", path, nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetAuthorizedCategoryAssets() (*GetAuthorizedCategoryAssetsRsp, error) {
	var result GetAuthorizedCategoryAssetsRsp
	path := "/authorization/202405/category_assets"
	err := c.doRequest("GET", path, nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
