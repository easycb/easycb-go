package ozon

import "github.com/easycb/easycb-go"

func (c *Client) GetDescriptionCategoryTree(body easycb.AnyMap) (*GetDescriptionCategoryTreeRsp, error) {
	var result GetDescriptionCategoryTreeRsp
	err := c.doRequest("POST", "/v1/description-category/tree", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetDescriptionCategoryAttribute(body easycb.AnyMap) (*GetDescriptionCategoryAttributeRsp, error) {
	var result GetDescriptionCategoryAttributeRsp
	err := c.doRequest("POST", "/v1/description-category/attribute", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetDescriptionCategoryAttributeValues(body easycb.AnyMap) (*GetDescriptionCategoryAttributeValuesRsp, error) {
	var result GetDescriptionCategoryAttributeValuesRsp
	err := c.doRequest("POST", "/v1/description-category/attribute/values", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetDescriptionCategoryAttributeValuesSearch(body easycb.AnyMap) (*GetDescriptionCategoryAttributeValuesSearchRsp, error) {
	var result GetDescriptionCategoryAttributeValuesSearchRsp
	err := c.doRequest("POST", "/v1/description-category/attribute/values/search", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DescriptionCategoryTips(body easycb.AnyMap) (*DescriptionCategoryTipsRsp, error) {
	var result DescriptionCategoryTipsRsp
	err := c.doRequest("POST", "/v1/description-category/tips", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
