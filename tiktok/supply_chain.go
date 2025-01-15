package tiktok

import "github.com/easycb/easycb-go"

func (c *Client) ConfirmPackageShipment(body easycb.AnyMap) (*ConfirmPackageShipmentRsp, error) {
	var result ConfirmPackageShipmentRsp
	path := "/supply_chain/202309/packages/sync"
	err := c.doRequest("POST", path, nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
