package shein

type BaseRsp struct {
	Code    string      `json:"code"`
	Msg     string      `json:"msg"`
	Bbl     interface{} `json:"bbl"`
	TraceId string      `json:"traceId"`
	Info    interface{} `json:"info"`
}
