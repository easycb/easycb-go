package tiktok

import (
	"fmt"
	"github.com/easycb/easycb-go"
)

func (c *Client) GetAfterSaleEligibility(orderId string, query easycb.AnyMap) (*GetAfterSaleEligibilityRsp, error) {
	var result GetAfterSaleEligibilityRsp
	path := fmt.Sprintf("/return_refund/202309/orders/%s/aftersale_eligibility", orderId)
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetRejectReasons(query easycb.AnyMap) (*GetRejectReasonsRsp, error) {
	var result GetRejectReasonsRsp
	path := "/return_refund/202309/reject_reasons"
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) CreateReturn(query easycb.AnyMap, body easycb.AnyMap) (*CreateReturnRsp, error) {
	var result CreateReturnRsp
	path := "/return_refund/202309/returns"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) ApproveReturn(returnId string, query easycb.AnyMap, body easycb.AnyMap) (*ApproveReturnRsp, error) {
	var result ApproveReturnRsp
	path := fmt.Sprintf("/return_refund/202309/returns/%s/approve", returnId)
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) RejectReturn(returnId string, query easycb.AnyMap, body easycb.AnyMap) (*RejectReturnRsp, error) {
	var result RejectReturnRsp
	path := fmt.Sprintf("/return_refund/202309/returns/%s/reject", returnId)
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) SearchReturns(query easycb.AnyMap, body easycb.AnyMap) (*SearchReturnsRsp, error) {
	var result SearchReturnsRsp
	path := "/return_refund/202309/returns/search"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetReturnRecords(returnId string, query easycb.AnyMap) (*GetReturnRecordsRsp, error) {
	var result GetReturnRecordsRsp
	path := fmt.Sprintf("/return_refund/202309/returns/%s/records", returnId)
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) CancelOrder(query easycb.AnyMap, body easycb.AnyMap) (*CancelOrderRsp, error) {
	var result CancelOrderRsp
	path := "/return_refund/202309/cancellations"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) ApproveCancellation(cancelId string, query easycb.AnyMap) (*ApproveCancellationRsp, error) {
	var result ApproveCancellationRsp
	path := fmt.Sprintf("/return_refund/202309/cancellations/%s/approve", cancelId)
	err := c.doRequest("POST", path, query, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) RejectCancellation(cancelId string, query easycb.AnyMap, body easycb.AnyMap) (*RejectCancellationRsp, error) {
	var result RejectCancellationRsp
	path := fmt.Sprintf("/return_refund/202309/cancellations/%s/reject", cancelId)
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) SearchCancellations(query easycb.AnyMap, body easycb.AnyMap) (*SearchCancellationsRsp, error) {
	var result SearchCancellationsRsp
	path := "/return_refund/202309/cancellations/search"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) CalculateRefund(query easycb.AnyMap, body easycb.AnyMap) (*CalculateRefundRsp, error) {
	var result CalculateRefundRsp
	path := "/return_refund/202309/refunds/calculate"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
