package lazada

type DataMoatComputeRiskRsp struct {
	Result struct {
		Msg             string `json:"msg"`
		RiskType        string `json:"riskType"`
		RiskDescription string `json:"riskDescription"`
		Success         bool   `json:"success"`
		Risk            string `json:"risk"`
	} `json:"result"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}

type DataMoatLoginRsp struct {
	Result struct {
		Msg     string `json:"msg"`
		Success bool   `json:"success"`
	} `json:"result"`
	Code      string `json:"code"`
	RequestId string `json:"request_id"`
}
