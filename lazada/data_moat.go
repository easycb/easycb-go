package lazada

import "github.com/easycb/easycb-go"

func (c *Client) DataMoatComputeRisk(body easycb.AnyMap) (*DataMoatComputeRiskRsp, error) {
	var result DataMoatComputeRiskRsp
	err := c.doRequest("POST", "/datamoat/compute_risk", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DataMoatLogin(body easycb.AnyMap) (*DataMoatLoginRsp, error) {
	var result DataMoatLoginRsp
	err := c.doRequest("POST", "/datamoat/login", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
