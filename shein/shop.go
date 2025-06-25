package shein

import "github.com/easycb/easycb-go"

func (c *Client) QueryStoreInfo(body easycb.AnyMap) (*QueryStoreInfoRsp, error) {
	var result QueryStoreInfoRsp
	err := c.doRequest("POST", "/open-api/openapi-business-backend/query-store-info", nil, body, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetAnnoList(body easycb.AnyMap, headers map[string]string) (*GetAnnoListRsp, error) {
	var result GetAnnoListRsp
	err := c.doRequest("POST", "/open-api/ssls/announcement/get-anno-list", nil, body, headers, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetAnnoDetail(body easycb.AnyMap, headers map[string]string) (*GetAnnoDetailRsp, error) {
	var result GetAnnoDetailRsp
	err := c.doRequest("POST", "/open-api/ssls/announcement/get-anno-detail", nil, body, headers, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
