package shopee

import "github.com/easycb/easycb-go"

func (c *Client) InitVideoUpload(body easycb.AnyMap) (*InitVideoUploadRsp, error) {
	var result InitVideoUploadRsp
	err := c.doRequest("POST", "/api/v2/media_space/init_video_upload", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UploadVideoPart(body easycb.AnyMap) (*UploadVideoPartRsp, error) {
	var result UploadVideoPartRsp
	err := c.doFileRequest("POST", "/api/v2/media_space/upload_video_part", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CompleteVideoUpload(body easycb.AnyMap) (*CompleteVideoUploadRsp, error) {
	var result CompleteVideoUploadRsp
	err := c.doRequest("POST", "/api/v2/media_space/upload_video_part", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetVideoUploadResult(query easycb.AnyMap) (*GetVideoUploadResultRsp, error) {
	var result GetVideoUploadResultRsp
	err := c.doRequest("GET", "/api/v2/media_space/get_video_upload_result", query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CancelVideoUpload(body easycb.AnyMap) (*CancelVideoUploadRsp, error) {
	var result CancelVideoUploadRsp
	err := c.doRequest("POST", "/api/v2/media_space/cancel_video_upload", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UploadImage(body easycb.AnyMap) (*UploadImageRsp, error) {
	var result UploadImageRsp
	err := c.doFileRequest("POST", "/api/v2/media_space/upload_image", nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
