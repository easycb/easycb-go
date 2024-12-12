package shopee

import "github.com/easycb/easycb-go"

func (c *Client) SetAppPushConfig(query easycb.AnyMap) (*SetAppPushConfigRsp, error) {
	var result SetAppPushConfigRsp
	err := c.doRequest("GET", "/api/v2/push/set_app_push_config", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetAppPushConfig() (*GetAppPushConfigRsp, error) {
	var result GetAppPushConfigRsp
	err := c.doRequest("GET", "/api/v2/push/get_app_push_config", nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetLostPushMessage() (*GetLostPushMessageRsp, error) {
	var result GetLostPushMessageRsp
	err := c.doRequest("GET", "/api/v2/push/get_lost_push_message", nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ConfirmConsumedLostPush(body easycb.AnyMap) (*ConfirmConsumedLostPushRsp, error) {
	var result ConfirmConsumedLostPushRsp
	err := c.doRequest("POST", "/api/v2/push/confirm_consumed_lost_push_message", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
