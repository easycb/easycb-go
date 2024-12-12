package shopee

import "github.com/easycb/easycb-go"

func (c *Client) GetEscrowDetail(query easycb.AnyMap) (*GetEscrowDetailRsp, error) {
	var result GetEscrowDetailRsp
	err := c.doRequest("GET", "/api/v2/payment/get_escrow_detail", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SetShopInstallmentStatus(body easycb.AnyMap) (*SetShopInstallmentStatusRsp, error) {
	var result SetShopInstallmentStatusRsp
	err := c.doRequest("POST", "/api/v2/payment/set_shop_installment_status", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetShopInstallmentStatus() (*GetShopInstallmentStatusRsp, error) {
	var result GetShopInstallmentStatusRsp
	err := c.doRequest("GET", "/api/v2/payment/get_shop_installment_status", nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetPayoutDetail(query easycb.AnyMap) (*GetPayoutDetailRsp, error) {
	var result GetPayoutDetailRsp
	err := c.doRequest("GET", "/api/v2/payment/get_payout_detail", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SetItemInstallmentStatus(body easycb.AnyMap) (*SetItemInstallmentStatusRsp, error) {
	var result SetItemInstallmentStatusRsp
	err := c.doRequest("POST", "/api/v2/payment/set_item_installment_status", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetItemInstallmentStatus(query easycb.AnyMap) (*GetItemInstallmentStatusRsp, error) {
	var result GetItemInstallmentStatusRsp
	err := c.doRequest("GET", "/api/v2/payment/get_item_installment_status", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetPaymentMethodList() (*GetPaymentMethodListRsp, error) {
	var result GetPaymentMethodListRsp
	err := c.doRequest("GET", "/api/v2/payment/get_payment_method_list", nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetWalletTransactionList(query easycb.AnyMap) (*GetWalletTransactionListRsp, error) {
	var result GetWalletTransactionListRsp
	err := c.doRequest("GET", "/api/v2/payment/get_wallet_transaction_list", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetEscrowList(query easycb.AnyMap) (*GetEscrowListRsp, error) {
	var result GetEscrowListRsp
	err := c.doRequest("GET", "/api/v2/payment/get_escrow_list", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetPayoutInfo(query easycb.AnyMap) (*GetPayoutInfoRsp, error) {
	var result GetPayoutInfoRsp
	err := c.doRequest("GET", "/api/v2/payment/get_payout_info", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetBillingTransactionInfo(body easycb.AnyMap) (*GetBillingTransactionInfoRsp, error) {
	var result GetBillingTransactionInfoRsp
	err := c.doRequest("POST", "/api/v2/payment/get_billing_transaction_info", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetEscrowDetailBatch(body easycb.AnyMap) (*GetEscrowDetailBatchRsp, error) {
	var result GetEscrowDetailBatchRsp
	err := c.doRequest("POST", "/api/v2/payment/get_escrow_detail_batch", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
