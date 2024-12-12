package tiktok

func (c *Client) GetActiveShops() (*GetActiveShopsRsp, error) {
	var result GetActiveShopsRsp
	path := "/seller/202309/shops"
	err := c.doRequest("GET", path, nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetSellerPermissions() (*GetSellerPermissionsRsp, error) {
	var result GetSellerPermissionsRsp
	path := "/seller/202309/permissions"
	err := c.doRequest("GET", path, nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
