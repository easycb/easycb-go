package ozon

import "github.com/easycb/easycb-go"

func (c *Client) CarriageCreate(body easycb.AnyMap) (*CarriageCreateRsp, error) {
	var result CarriageCreateRsp
	err := c.doRequest("POST", "/v1/carriage/create", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CarriageApprove(body easycb.AnyMap) (*CarriageApproveRsp, error) {
	var result CarriageApproveRsp
	err := c.doRequest("POST", "/v1/carriage/approve", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CarriageGet(body easycb.AnyMap) (*CarriageGetRsp, error) {
	var result CarriageGetRsp
	err := c.doRequest("POST", "/v1/carriage/get", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CarriageCourierContactSet(body easycb.AnyMap) (*CarriageCourierContactSetRsp, error) {
	var result CarriageCourierContactSetRsp
	err := c.doRequest("POST", "/v1/carriage/courier-contact/set", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CarriageCourierContactGet(body easycb.AnyMap) (*CarriageCourierContactGetRsp, error) {
	var result CarriageCourierContactGetRsp
	err := c.doRequest("POST", "/v1/carriage/courier-contact/get", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PostingCarriageAvailableList(body easycb.AnyMap) (*PostingCarriageAvailableListRsp, error) {
	var result PostingCarriageAvailableListRsp
	err := c.doRequest("POST", "/v1/posting/carriage-available/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CarriageSetPostings(body easycb.AnyMap) (*CarriageSetPostingsRsp, error) {
	var result CarriageSetPostingsRsp
	err := c.doRequest("POST", "/v1/carriage/set-postings", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CarriageCancel(body easycb.AnyMap) (*CarriageCancelRsp, error) {
	var result CarriageCancelRsp
	err := c.doRequest("POST", "/v1/carriage/cancel", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PostingFbsActGetPostings(body easycb.AnyMap) (*PostingFbsActGetPostingsRsp, error) {
	var result PostingFbsActGetPostingsRsp
	err := c.doRequest("POST", "/v2/posting/fbs/act/get-postings", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PostingFbsActGetContainerLabels(body easycb.AnyMap) (*PostingFbsActGetContainerLabelsRsp, error) {
	var result PostingFbsActGetContainerLabelsRsp
	err := c.doRequest("POST", "/v2/posting/fbs/act/get-container-labels", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) AssemblyCarriagePostingList(body easycb.AnyMap) (*AssemblyCarriagePostingListRsp, error) {
	var result AssemblyCarriagePostingListRsp
	err := c.doRequest("POST", "/v1/assembly/carriage/posting/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) AssemblyCarriageProductList(body easycb.AnyMap) (*AssemblyCarriageProductListRsp, error) {
	var result AssemblyCarriageProductListRsp
	err := c.doRequest("POST", "/v1/assembly/carriage/product/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) AssemblyFbsPostingList(body easycb.AnyMap) (*AssemblyFbsPostingListRsp, error) {
	var result AssemblyFbsPostingListRsp
	err := c.doRequest("POST", "/v1/assembly/fbs/posting/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) AssemblyFbsProductList(body easycb.AnyMap) (*AssemblyFbsProductListRsp, error) {
	var result AssemblyFbsProductListRsp
	err := c.doRequest("POST", "/v1/assembly/fbs/product/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
