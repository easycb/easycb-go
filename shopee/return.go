package shopee

import "github.com/easycb/easycb-go"

func (c *Client) GetReturnDetail(query easycb.AnyMap) (*GetReturnDetailRsp, error) {
	var result GetReturnDetailRsp
	err := c.doRequest("GET", "/api/v2/returns/get_return_detail", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetReturnList(query easycb.AnyMap) (*GetReturnListRsp, error) {
	var result GetReturnListRsp
	err := c.doRequest("GET", "/api/v2/returns/get_return_list", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ReturnConfirm(body easycb.AnyMap) (*ReturnConfirmRsp, error) {
	var result ReturnConfirmRsp
	err := c.doRequest("POST", "/api/v2/returns/confirm", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ReturnDispute(body easycb.AnyMap) (*ReturnDisputeRsp, error) {
	var result ReturnDisputeRsp
	err := c.doRequest("POST", "/api/v2/returns/dispute", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetAvailableSolutions(query easycb.AnyMap) (*GetAvailableSolutionsRsp, error) {
	var result GetAvailableSolutionsRsp
	err := c.doRequest("GET", "/api/v2/returns/get_available_solutions", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ReturnOffer(body easycb.AnyMap) (*ReturnOfferRsp, error) {
	var result ReturnOfferRsp
	err := c.doRequest("POST", "/api/v2/returns/offer", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ReturnAcceptOffer(body easycb.AnyMap) (*ReturnAcceptOfferRsp, error) {
	var result ReturnAcceptOfferRsp
	err := c.doRequest("POST", "/api/v2/returns/accept_offer", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ConvertImage(body easycb.AnyMap, fileType string) (*ConvertImageRsp, error) {
	var result ConvertImageRsp
	err := c.doFileRequest("POST", "/api/v2/returns/convert_image", nil, body, fileType, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UploadProof(body easycb.AnyMap) (*UploadProofRsp, error) {
	var result UploadProofRsp
	err := c.doRequest("POST", "/api/v2/returns/upload_proof", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) QueryProof(query easycb.AnyMap) (*QueryProofRsp, error) {
	var result QueryProofRsp
	err := c.doRequest("GET", "/api/v2/returns/query_proof", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetReturnDisputeReason(query easycb.AnyMap) (*GetReturnDisputeReasonRsp, error) {
	var result GetReturnDisputeReasonRsp
	err := c.doRequest("GET", "/api/v2/returns/get_return_dispute_reason", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
