package ozon

import "github.com/easycb/easycb-go"

func (c *Client) FinanceRealization(body easycb.AnyMap) (*FinanceRealizationRsp, error) {
	var result FinanceRealizationRsp
	err := c.doRequest("POST", "/v2/finance/realization", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FinanceRealizationPosting(body easycb.AnyMap) (*FinanceRealizationPostingRsp, error) {
	var result FinanceRealizationPostingRsp
	err := c.doRequest("POST", "/v1/finance/realization/posting", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FinanceTransactionList(body easycb.AnyMap) (*FinanceTransactionListRsp, error) {
	var result FinanceTransactionListRsp
	err := c.doRequest("POST", "/v3/finance/transaction/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FinanceTransactionTotals(body easycb.AnyMap) (*FinanceTransactionTotalsRsp, error) {
	var result FinanceTransactionTotalsRsp
	err := c.doRequest("POST", "/v3/finance/transaction/totals", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FinanceCompensation(body easycb.AnyMap) (*FinanceCompensationRsp, error) {
	var result FinanceCompensationRsp
	err := c.doRequest("POST", "/v1/finance/compensation", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FinanceDecompensation(body easycb.AnyMap) (*FinanceDecompensationRsp, error) {
	var result FinanceDecompensationRsp
	err := c.doRequest("POST", "/v1/finance/decompensation", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FinanceBalance(body easycb.AnyMap) (*FinanceBalanceRsp, error) {
	var result FinanceBalanceRsp

	err := c.doRequest("POST", "/v1/finance/balance", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FinanceAccrualPostings(body easycb.AnyMap) (*FinanceAccrualPostingsRsp, error) {
	var result FinanceAccrualPostingsRsp

	err := c.doRequest("POST", "/v1/finance/accrual/postings", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FinanceAccrualTypes(body easycb.AnyMap) (*FinanceAccrualTypesRsp, error) {
	var result FinanceAccrualTypesRsp

	err := c.doRequest("POST", "/v1/finance/accrual/types", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FinanceAccrualByDay(body easycb.AnyMap) (*FinanceAccrualByDayRsp, error) {
	var result FinanceAccrualByDayRsp

	err := c.doRequest("POST", "/v1/finance/accrual/by-day", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
