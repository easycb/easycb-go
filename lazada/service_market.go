package lazada

import "github.com/easycb/easycb-go"

func (c *Client) ServiceMarketAppKeyOrderQuery(query easycb.AnyMap) (*ServiceMarketAppKeyOrderQueryRsp, error) {
	var result ServiceMarketAppKeyOrderQueryRsp
	err := c.doRequest("GET", "/service/market/order/query", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ServiceMarketAppKeySubQuery(query easycb.AnyMap) (*ServiceMarketAppKeySubQueryRsp, error) {
	var result ServiceMarketAppKeySubQueryRsp
	err := c.doRequest("GET", "/service/market/subs/query", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
