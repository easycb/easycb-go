package lazada

import "github.com/easycb/easycb-go"

func (c *Client) GetStoreCustomPage(query easycb.AnyMap) (*GetStoreCustomPageRsp, error) {
	var result GetStoreCustomPageRsp
	err := c.doRequest("GET", "", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
