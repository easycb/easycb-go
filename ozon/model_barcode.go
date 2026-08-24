package ozon

type BarcodeAddRsp struct {
	Errors []struct {
		Code    string `json:"code"`
		Error   string `json:"error"`
		Barcode string `json:"barcode"`
		Sku     int64  `json:"sku"`
	} `json:"errors"`
}

type BarcodeGenerateRsp struct {
	BaseRsp
	Errors []struct {
		Code      string `json:"code"`
		Error     string `json:"error"`
		Barcode   string `json:"barcode"`
		ProductId int64  `json:"product_id"`
	} `json:"errors"`
}
