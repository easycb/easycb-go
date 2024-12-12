package easycb

import "errors"

var (
	MissLazadaInitParamErr = errors.New("missing lazada init param")
	MissShopeeInitParamErr = errors.New("missing shopee init param")
	MissTiktokInitParamErr = errors.New("missing tiktok init param")
	MissParamErr           = errors.New("missing required parameter")
	MarshalErr             = errors.New("marshal error")
	UnmarshalErr           = errors.New("unmarshal error")
	SignatureErr           = errors.New("signature error")
	MissBaseUrlErr         = errors.New("missing base url")
)
