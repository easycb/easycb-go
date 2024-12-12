package lazada

import "github.com/easycb/easycb-go"

func (c *Client) SellerVoucherDeleteSelectedProductSKU(body easycb.AnyMap) (*SellerVoucheDeleteSelectedProductSKURsp, error) {
	var result SellerVoucheDeleteSelectedProductSKURsp
	err := c.doRequest("POST", "/promotion/voucher/product/sku/remove", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SellerVoucherActivate(body easycb.AnyMap) (*SellerVoucherActivateRsp, error) {
	var result SellerVoucherActivateRsp
	err := c.doRequest("POST", "/promotion/voucher/activate", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SellerVoucherAddSelectedProductSKU(body easycb.AnyMap) (*SellerVoucherAddSelectedProductSKURsp, error) {
	var result SellerVoucherAddSelectedProductSKURsp
	err := c.doRequest("POST", "/promotion/voucher/product/sku/add", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SellerVoucherCreate(body easycb.AnyMap) (*SellerVoucherCreateRsp, error) {
	var result SellerVoucherCreateRsp
	err := c.doRequest("POST", "/promotion/voucher/create", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SellerVoucherDeactivate(body easycb.AnyMap) (*SellerVoucherDeactivateRsp, error) {
	var result SellerVoucherDeactivateRsp
	err := c.doRequest("POST", "/promotion/voucher/deactivate", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SellerVoucherDetailQuery(query easycb.AnyMap) (*SellerVoucherDetailQueryRsp, error) {
	var result SellerVoucherDetailQueryRsp
	err := c.doRequest("GET", "/promotion/voucher/get", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SellerVoucherList(query easycb.AnyMap) (*SellerVoucherListRsp, error) {
	var result SellerVoucherListRsp
	err := c.doRequest("GET", "/promotion/vouchers/get", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SellerVoucherSelectedProductList(query easycb.AnyMap) (*SellerVoucherSelectedProductListRsp, error) {
	var result SellerVoucherSelectedProductListRsp
	err := c.doRequest("GET", "/promotion/voucher/products/get", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SellerVoucherUpdate(body easycb.AnyMap) (*SellerVoucherUpdateRsp, error) {
	var result SellerVoucherUpdateRsp
	err := c.doRequest("POST", "/promotion/voucher/update", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
