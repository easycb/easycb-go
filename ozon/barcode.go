package ozon

import "github.com/easycb/easycb-go"

func (c *Client) BarcodeAdd(body easycb.AnyMap) (*BarcodeAddRsp, error) {
	var result BarcodeAddRsp
	err := c.doRequest("POST", "/v1/barcode/add", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) BarcodeGenerate(body easycb.AnyMap) (*BarcodeGenerateRsp, error) {
	var result BarcodeGenerateRsp
	err := c.doRequest("POST", "/v1/barcode/generate", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
