package lazada

import "github.com/easycb/easycb-go"

func (c *Client) GetReverseOrderDetail(query easycb.AnyMap) (*GetReverseOrderDetailRsp, error) {
	var result GetReverseOrderDetailRsp
	err := c.doRequest("GET", "/order/reverse/return/detail/list", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetReverseOrderHistoryList(query easycb.AnyMap) (*GetReverseOrderHistoryListRsp, error) {
	var result GetReverseOrderHistoryListRsp
	err := c.doRequest("GET", "/order/reverse/return/history/list", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetReverseOrderReasonList(query easycb.AnyMap) (*GetReverseOrderReasonListRsp, error) {
	var result GetReverseOrderReasonListRsp
	err := c.doRequest("GET", "/order/reverse/reason/list", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetReverseOrdersForSeller(query easycb.AnyMap) (*GetReverseOrdersForSellerRsp, error) {
	var result GetReverseOrdersForSellerRsp
	err := c.doRequest("GET", "/reverse/getreverseordersforseller", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) InitReverseOrderCancel(query easycb.AnyMap) (*InitReverseOrderCancelRsp, error) {
	var result InitReverseOrderCancelRsp
	err := c.doRequest("GET", "/order/reverse/cancel/create", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) InitReverseOrderCancelDecide(query easycb.AnyMap) (*InitReverseOrderCancelDecideRsp, error) {
	var result InitReverseOrderCancelDecideRsp
	err := c.doRequest("GET", "/order/reverse/cancel/seller/decide", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ReverseOrderReturnUpdate(query easycb.AnyMap) (*ReverseOrderReturnUpdateRsp, error) {
	var result ReverseOrderReturnUpdateRsp
	err := c.doRequest("GET", "/order/reverse/return/update", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
