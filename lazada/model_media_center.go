package lazada

type CompleteCreateVideoRsp struct {
	Code          string `json:"code"`
	ResultMessage string `json:"result_message"`
	Success       bool   `json:"success"`
	ResultCode    string `json:"result_code"`
	RequestId     string `json:"request_id"`
	VideoId       string `json:"video_id"`
}

type GetVideoRsp struct {
	CoverUrl      string `json:"cover_url"`
	VideoUrl      string `json:"video_url"`
	Code          string `json:"code"`
	ResultMessage string `json:"result_message"`
	Success       bool   `json:"success"`
	ResultCode    string `json:"result_code"`
	State         string `json:"state"`
	Title         string `json:"title"`
	RequestId     string `json:"request_id"`
}

type GetVideoQuotaRsp struct {
	Code          string `json:"code"`
	ResultMessage string `json:"result_message"`
	Success       bool   `json:"success"`
	ResultCode    string `json:"result_code"`
	CapacitySize  string `json:"capacity_size"`
	UsedSize      string `json:"used_size"`
	RequestId     string `json:"request_id"`
}

type InitCreateVideoRsp struct {
	UploadId      string `json:"upload_id"`
	Code          string `json:"code"`
	ResultMessage string `json:"result_message"`
	Success       bool   `json:"success"`
	ResultCode    string `json:"result_code"`
	RequestId     string `json:"request_id"`
}

type RemoveVideoRsp struct {
	Code          string `json:"code"`
	ResultMessage string `json:"result_message"`
	Success       bool   `json:"success"`
	ResultCode    string `json:"result_code"`
	RequestId     string `json:"request_id"`
}

type UploadVideoBlockRsp struct {
	Code          string `json:"code"`
	ResultMessage string `json:"result_message"`
	Success       bool   `json:"success"`
	ResultCode    string `json:"result_code"`
	ETag          string `json:"e_tag"`
	RequestId     string `json:"request_id"`
}
