package ozon

type BaseRsp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details []struct {
		TypeUrl string `json:"typeUrl"`
		Value   string `json:"value"`
	} `json:"details"`
}
