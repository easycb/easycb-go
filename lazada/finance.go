package lazada

import "github.com/easycb/easycb-go"

func (c *Client) GetPayoutStatus(query easycb.AnyMap) (*GetPayoutStatusRsp, error) {
	var result GetPayoutStatusRsp
	err := c.doRequest("GET", "/finance/payout/status/get", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) QueryAccountTransactions(body easycb.AnyMap) (*QueryAccountTransactionsRsp, error) {
	var result QueryAccountTransactionsRsp
	err := c.doRequest("POST", "/finance/transaction/accountTransactions/query", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) QueryLogisticsFeeDetail(query easycb.AnyMap) (*QueryLogisticsFeeDetailRsp, error) {
	var result QueryLogisticsFeeDetailRsp
	err := c.doRequest("GET", "/lbs/slb/queryLogisticsFeeDetail", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) QueryTransactionDetails(query easycb.AnyMap) (*QueryTransactionDetailsRsp, error) {
	var result QueryTransactionDetailsRsp
	err := c.doRequest("GET", "/finance/transaction/details/get", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
