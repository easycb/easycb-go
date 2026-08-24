package ozon

import "github.com/easycb/easycb-go"

func (c *Client) SellerActionsCreateDiscount(body easycb.AnyMap) (*SellerActionsCreateDiscountRsp, error) {
	var result SellerActionsCreateDiscountRsp
	err := c.doRequest("POST", "/v1/seller-actions/create/discount", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SellerActionsCreateDiscountWithCondition(body easycb.AnyMap) (*SellerActionsCreateDiscountWithConditionRsp, error) {
	var result SellerActionsCreateDiscountWithConditionRsp
	err := c.doRequest("POST", "/v1/seller-actions/create/discount-with-condition", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SellerActionsCreateInstallment(body easycb.AnyMap) (*SellerActionsCreateInstallmentRsp, error) {
	var result SellerActionsCreateInstallmentRsp
	err := c.doRequest("POST", "/v1/seller-actions/create/installment", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SellerActionsCreateMultiLevelDiscount(body easycb.AnyMap) (*SellerActionsCreateMultiLevelDiscountRsp, error) {
	var result SellerActionsCreateMultiLevelDiscountRsp
	err := c.doRequest("POST", "/v1/seller-actions/create/multi-level-discount", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SellerActionsCreateVoucher(body easycb.AnyMap) (*SellerActionsCreateVoucherRsp, error) {
	var result SellerActionsCreateVoucherRsp
	err := c.doRequest("POST", "/v1/seller-actions/create/voucher", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SellerActionsUpdateDiscount(body easycb.AnyMap) (*SellerActionsUpdateDiscountRsp, error) {
	var result SellerActionsUpdateDiscountRsp
	err := c.doRequest("POST", "/v1/seller-actions/update/discount", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SellerActionsUpdateDiscountWithCondition(body easycb.AnyMap) (*SellerActionsUpdateDiscountWithConditionRsp, error) {
	var result SellerActionsUpdateDiscountWithConditionRsp
	err := c.doRequest("POST", "/v1/seller-actions/update/discount-with-condition", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SellerActionsUpdateInstallment(body easycb.AnyMap) (*SellerActionsUpdateInstallmentRsp, error) {
	var result SellerActionsUpdateInstallmentRsp
	err := c.doRequest("POST", "/v1/seller-actions/update/installment", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SellerActionsUpdateMultiLevelDiscount(body easycb.AnyMap) (*SellerActionsUpdateMultiLevelDiscountRsp, error) {
	var result SellerActionsUpdateMultiLevelDiscountRsp
	err := c.doRequest("POST", "/v1/seller-actions/update/multi-level-discount", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SellerActionsUpdateVoucher(body easycb.AnyMap) (*SellerActionsUpdateVoucherRsp, error) {
	var result SellerActionsUpdateVoucherRsp
	err := c.doRequest("POST", "/v1/seller-actions/update/voucher", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SellerActionsProductsAdd(body easycb.AnyMap) (*SellerActionsProductsAddRsp, error) {
	var result SellerActionsProductsAddRsp
	err := c.doRequest("POST", "/v1/seller-actions/products/add", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SellerActionsProductsCandidates(body easycb.AnyMap) (*SellerActionsProductsCandidatesRsp, error) {
	var result SellerActionsProductsCandidatesRsp
	err := c.doRequest("POST", "/v1/seller-actions/products/candidates", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SellerActionsProductsDelete(body easycb.AnyMap) (*SellerActionsProductsDeleteRsp, error) {
	var result SellerActionsProductsDeleteRsp
	err := c.doRequest("POST", "/v1/seller-actions/products/delete", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SellerActionsProductsList(body easycb.AnyMap) (*SellerActionsProductsListRsp, error) {
	var result SellerActionsProductsListRsp
	err := c.doRequest("POST", "/v1/seller-actions/products/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SellerActionsArchive(body easycb.AnyMap) (*SellerActionsArchiveRsp, error) {
	var result SellerActionsArchiveRsp
	err := c.doRequest("POST", "/v1/seller-actions/archive", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SellerActionsChangeActivity(body easycb.AnyMap) (*SellerActionsChangeActivityRsp, error) {
	var result SellerActionsChangeActivityRsp
	err := c.doRequest("POST", "/v1/seller-actions/change-activity", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SellerActionsList(body easycb.AnyMap) (*SellerActionsListRsp, error) {
	var result SellerActionsListRsp
	err := c.doRequest("POST", "/v1/seller-actions/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SellerActionsVoucherGet(body easycb.AnyMap) (*SellerActionsVoucherGetRsp, error) {
	var result SellerActionsVoucherGetRsp
	err := c.doRequest("POST", "/v1/seller-actions/voucher/get", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
