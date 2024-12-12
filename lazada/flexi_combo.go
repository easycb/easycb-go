package lazada

import "github.com/easycb/easycb-go"

func (c *Client) ActivateFlexiCombo(body easycb.AnyMap) (*ActivateFlexiComboRsp, error) {
	var result ActivateFlexiComboRsp
	err := c.doRequest("POST", "/promotion/flexicombo/activate", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) AddFlexiComboProducts(body easycb.AnyMap) (*AddFlexiComboProductsRsp, error) {
	var result AddFlexiComboProductsRsp
	err := c.doRequest("POST", "/promotion/flexicombo/products/add", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CreateFlexiCombo(body easycb.AnyMap) (*CreateFlexiComboRsp, error) {
	var result CreateFlexiComboRsp
	err := c.doRequest("POST", "/promotion/flexicombo/create", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeactivateFlexiCombo(body easycb.AnyMap) (*DeactivateFlexiComboRsp, error) {
	var result DeactivateFlexiComboRsp
	err := c.doRequest("POST", "/promotion/flexicombo/deactivate", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteFlexiComboProducts(body easycb.AnyMap) (*DeleteFlexiComboProductsRsp, error) {
	var result DeleteFlexiComboProductsRsp
	err := c.doRequest("POST", "/promotion/flexicombo/products/delete", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetFlexiComboDetails(query easycb.AnyMap) (*GetFlexiComboDetailsRsp, error) {
	var result GetFlexiComboDetailsRsp
	err := c.doRequest("GET", "/promotion/flexicombo/details", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ListFlexiCombo(query easycb.AnyMap) (*ListFlexiComboRsp, error) {
	var result ListFlexiComboRsp
	err := c.doRequest("GET", "/promotion/flexicombo/list", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ListFlexiComboProducts(query easycb.AnyMap) (*ListFlexiComboProductsRsp, error) {
	var result ListFlexiComboProductsRsp
	err := c.doRequest("GET", "/promotion/flexicombo/products/list", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateFlexiCombo(body easycb.AnyMap) (*UpdateFlexiComboRsp, error) {
	var result UpdateFlexiComboRsp
	err := c.doRequest("POST", "/promotion/flexicombo/update", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
