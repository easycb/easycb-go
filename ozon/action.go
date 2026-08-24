package ozon

import "github.com/easycb/easycb-go"

func (c *Client) Actions(query easycb.AnyMap) (*ActionsRsp, error) {
	var result ActionsRsp
	err := c.doRequest("GET", "/v1/actions", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ActionsCandidates(body easycb.AnyMap) (*ActionsCandidatesRsp, error) {
	var result ActionsCandidatesRsp
	err := c.doRequest("POST", "/v1/actions/candidates", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ActionsProducts(body easycb.AnyMap) (*ActionsProductsRsp, error) {
	var result ActionsProductsRsp
	err := c.doRequest("POST", "/v1/actions/products", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ActionsProductsActivate(body easycb.AnyMap) (*ActionsProductsActivateRsp, error) {
	var result ActionsProductsActivateRsp
	err := c.doRequest("POST", "/v1/actions/products/activate", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ActionsProductsDeactivate(body easycb.AnyMap) (*ActionsProductsDeactivateRsp, error) {
	var result ActionsProductsDeactivateRsp
	err := c.doRequest("POST", "/v1/actions/products/deactivate", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ActionsDiscountsTaskList(body easycb.AnyMap) (*ActionsDiscountsTaskListRsp, error) {
	var result ActionsDiscountsTaskListRsp
	err := c.doRequest("POST", "/v1/actions/discounts-task/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ActionsDiscountsTaskApprove(body easycb.AnyMap) (*ActionsDiscountsTaskApproveRsp, error) {
	var result ActionsDiscountsTaskApproveRsp
	err := c.doRequest("POST", "/v1/actions/discounts-task/approve", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ActionsDiscountsTaskDecline(body easycb.AnyMap) (*ActionsDiscountsTaskDeclineRsp, error) {
	var result ActionsDiscountsTaskDeclineRsp
	err := c.doRequest("POST", "/v1/actions/discounts-task/decline", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ActionsDiscountsTaskListV2(body easycb.AnyMap) (*ActionsDiscountsTaskListV2Rsp, error) {
	var result ActionsDiscountsTaskListV2Rsp
	err := c.doRequest("POST", "/v2/actions/discounts-task/list", nil, body, &result)

	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) ActionsAutoAddProductsList(body easycb.AnyMap) (*ActionsAutoAddProductsListRsp, error) {
	var result ActionsAutoAddProductsListRsp
	err := c.doRequest("POST", "/v1/actions/auto-add/products/list", body, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ActionsAutoAddProductsCandidates(body easycb.AnyMap) (*ActionsAutoAddProductsCandidatesRsp, error) {
	var result ActionsAutoAddProductsCandidatesRsp
	err := c.doRequest("POST", "/v1/actions/auto-add/products/candidates", body, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ActionsAutoAddProductsDelete(body easycb.AnyMap) (*ActionsAutoAddProductsDeleteRsp, error) {
	var result ActionsAutoAddProductsDeleteRsp
	err := c.doRequest("POST", "/v1/actions/auto-add/products/delete", body, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ActionsAutoAddProductsUpdate(body easycb.AnyMap) (*ActionsAutoAddProductsUpdateRsp, error) {
	var result ActionsAutoAddProductsUpdateRsp
	err := c.doRequest("POST", "/v1/actions/auto-add/products/update", body, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
