package ozon

import "github.com/easycb/easycb-go"

func (c *Client) FbpDraftDropOffProvinceList(body easycb.AnyMap) (*FbpDraftDropOffProvinceListRsp, error) {
	var result FbpDraftDropOffProvinceListRsp
	err := c.doRequest("POST", "/v1/fbp/draft/drop-off/province/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbpDraftDropOffPointList(body easycb.AnyMap) (*FbpDraftDropOffPointListRsp, error) {
	var result FbpDraftDropOffPointListRsp
	err := c.doRequest("POST", "/v1/fbp/draft/drop-off/point/list", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbpDraftDropOffPointTimetable(body easycb.AnyMap) (*FbpDraftDropOffPointTimetableRsp, error) {
	var result FbpDraftDropOffPointTimetableRsp
	err := c.doRequest("POST", "/v1/fbp/draft/drop-off/point/timetable", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbpDraftDropOffProductValidate(body easycb.AnyMap) (*FbpDraftDropOffProductValidateRsp, error) {
	var result FbpDraftDropOffProductValidateRsp
	err := c.doRequest("POST", "/v1/fbp/draft/drop-off/product/validate", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbpDraftDropOffCreate(body easycb.AnyMap) (*FbpDraftDropOffCreateRsp, error) {
	var result FbpDraftDropOffCreateRsp
	err := c.doRequest("POST", "/v1/fbp/draft/drop-off/create", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbpDraftDropOffDelete(body easycb.AnyMap) (*FbpDraftDropOffDeleteRsp, error) {
	var result FbpDraftDropOffDeleteRsp
	err := c.doRequest("POST", "/v1/fbp/draft/drop-off/delete", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbpDraftDropOffDlvEdit(body easycb.AnyMap) (*FbpDraftDropOffDlvEditRsp, error) {
	var result FbpDraftDropOffDlvEditRsp
	err := c.doRequest("POST", "/v1/fbp/draft/drop-off/dlv/edit", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FbpDraftDropOffRegistrate(body easycb.AnyMap) (*FbpDraftDropOffRegistrateRsp, error) {
	var result FbpDraftDropOffRegistrateRsp
	err := c.doRequest("POST", "/v1/fbp/draft/drop-off/registrate", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
