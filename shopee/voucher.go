package shopee

import "github.com/easycb/easycb-go"

func (c *Client) GetVoucherList(query easycb.AnyMap) (*GetVoucherListRsp, error) {
	var result GetVoucherListRsp
	err := c.doRequest("GET", "/api/v2/voucher/get_voucher_list", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetVoucher(query easycb.AnyMap) (*GetVoucherRsp, error) {
	var result GetVoucherRsp
	err := c.doRequest("GET", "/api/v2/voucher/get_voucher", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateVoucher(body easycb.AnyMap) (*UpdateVoucherRsp, error) {
	var result UpdateVoucherRsp
	err := c.doRequest("POST", "/api/v2/voucher/update_voucher", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) EndVoucher(body easycb.AnyMap) (*EndVoucherRsp, error) {
	var result EndVoucherRsp
	err := c.doRequest("POST", "/api/v2/voucher/end_voucher", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteVoucher(body easycb.AnyMap) (*DeleteVoucherRsp, error) {
	var result DeleteVoucherRsp
	err := c.doRequest("POST", "/api/v2/voucher/delete_voucher", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) AddVoucher(body easycb.AnyMap) (*AddVoucherRsp, error) {
	var result AddVoucherRsp
	err := c.doRequest("POST", "/api/v2/voucher/add_voucher", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
