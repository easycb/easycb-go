package shein

import "github.com/easycb/easycb-go"

func (c *Client) AgencyList(body easycb.AnyMap, headers map[string]string) (*AgencyListRsp, error) {
	var result AgencyListRsp
	err := c.doRequest("POST", "/open-api/goods-compliance/agency-list", nil, body, headers, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SkcAgencyDetail(body easycb.AnyMap, headers map[string]string) (*SkcAgencyDetailRsp, error) {
	var result SkcAgencyDetailRsp
	err := c.doRequest("POST", "/open-api/goods-compliance/skc-agency-detail", nil, body, headers, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SaveSkcAgency(body easycb.AnyMap, headers map[string]string) (*SaveSkcAgencyRsp, error) {
	var result SaveSkcAgencyRsp
	err := c.doRequest("POST", "/open-api/goods-compliance/save-skc-agency", nil, body, headers, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SkcLabelList(body easycb.AnyMap, headers map[string]string) (*SkcLabelListRsp, error) {
	var result SkcLabelListRsp
	err := c.doRequest("POST", "/open-api/goods-compliance/skc-label-list", nil, body, headers, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) MaterialQualityTreeV2(query easycb.AnyMap, headers map[string]string) (*MaterialQualityTreeV2Rsp, error) {
	var result MaterialQualityTreeV2Rsp
	err := c.doRequest("GET", "/open-api/goods-quality/environmental-label-rule/material-quality-tree-v2", query, nil, headers, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) LabelPrint(body easycb.AnyMap) (*LabelPrintRsp, error) {
	var result LabelPrintRsp
	err := c.doRequest("POST", "/open-api/goods-compliance/label-print", nil, body, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UploadSkcLabelPicture(body easycb.AnyMap, headers map[string]string) (*UploadSkcLabelPictureRsp, error) {
	var result UploadSkcLabelPictureRsp
	err := c.doFileRequest("POST", "/open-api/goods-compliance/upload-skc-label-picture", nil, body, headers, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SkuSaveLabel(body easycb.AnyMap, headers map[string]string) (*SkuSaveLabelRsp, error) {
	var result SkuSaveLabelRsp
	err := c.doRequest("POST", "/open-api/goods-compliance/skc-save-label", nil, body, headers, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
