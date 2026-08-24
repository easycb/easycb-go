package ozon

type PolygonCreateRsp struct {
	BaseRsp
	PolygonId int64 `json:"polygonId"`
}

type PolygonBindRsp struct {
	BaseRsp
}
