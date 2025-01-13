package lazada

import "github.com/easycb/easycb-go"

func (c *Client) CompleteCreateVideo(body easycb.AnyMap) (*CompleteCreateVideoRsp, error) {
	var result CompleteCreateVideoRsp
	err := c.doRequest("POST", "/media/video/block/commit", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetVideo(query easycb.AnyMap) (*GetVideoRsp, error) {
	var result GetVideoRsp
	err := c.doRequest("GET", "/media/video/get", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetVideoQuota(query easycb.AnyMap) (*GetVideoQuotaRsp, error) {
	var result GetVideoQuotaRsp
	err := c.doRequest("GET", "/media/video/quota/get", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) InitCreateVideo(body easycb.AnyMap) (*InitCreateVideoRsp, error) {
	var result InitCreateVideoRsp
	err := c.doRequest("POST", "/media/video/block/create", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) RemoveVideo(body easycb.AnyMap) (*RemoveVideoRsp, error) {
	var result RemoveVideoRsp
	err := c.doRequest("POST", "/media/video/remove", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UploadVideoBlock(body easycb.AnyMap) (*UploadVideoBlockRsp, error) {
	var result UploadVideoBlockRsp
	err := c.doFileRequest("POST", "/media/video/block/upload", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
