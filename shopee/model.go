package shopee

type BaseRsp struct {
	Error     string      `json:"error"`
	Message   string      `json:"message"`
	Warning   string      `json:"warning"`
	RequestId string      `json:"request_id"`
	Response  interface{} `json:"response"`
}
