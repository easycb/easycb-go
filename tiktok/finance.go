package tiktok

import (
	"fmt"
	"github.com/easycb/easycb-go"
)

func (c *Client) GetStatements(query easycb.AnyMap) (*GetStatementsRsp, error) {
	var result GetStatementsRsp
	path := "/finance/202309/statements"
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetStatementTransactions(statementId string, query easycb.AnyMap) (*GetStatementTransactionsRsp, error) {
	var result GetStatementTransactionsRsp
	path := fmt.Sprintf("/finance/202309/statements/%s/statement_transactions", statementId)
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetStatementTransactionsV202501(statementId string, query easycb.AnyMap) (*GetStatementTransactionsV202501Rsp, error) {
	var result GetStatementTransactionsV202501Rsp
	path := fmt.Sprintf("/finance/202501/statements/%s/statement_transactions", statementId)
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetOrderStatementTransactions(orderId string) (*GetOrderStatementTransactionsRsp, error) {
	var result GetOrderStatementTransactionsRsp
	path := fmt.Sprintf("/finance/202309/orders/%s/statement_transactions", orderId)
	err := c.doRequest("GET", path, nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetOrderStatementTransactionsV202501(orderId string) (*GetOrderStatementTransactionsV202501Rsp, error) {
	var result GetOrderStatementTransactionsV202501Rsp
	path := fmt.Sprintf("/finance/202501/orders/%s/statement_transactions", orderId)
	err := c.doRequest("GET", path, nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetPayments(query easycb.AnyMap) (*GetPaymentsRsp, error) {
	var result GetPaymentsRsp
	path := "/finance/202309/payments"
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetWithdrawals(query easycb.AnyMap) (*GetWithdrawalsRsp, error) {
	var result GetWithdrawalsRsp
	path := "/finance/202309/withdrawals"
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetUnsettledTransactions(query easycb.AnyMap) (*GetUnsettledTransactionsRsp, error) {
	var result GetUnsettledTransactionsRsp
	path := "/finance/202507/orders/unsettled"
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
