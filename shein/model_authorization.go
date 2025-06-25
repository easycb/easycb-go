package shein

type AuthGetByTokenRsp struct {
	BaseRsp
	Info struct {
		SecretKey            string `json:"secretKey"`
		Appid                string `json:"appid"`
		OpenKeyId            string `json:"openKeyId"`
		State                string `json:"state"`
		SupplierId           int64  `json:"supplierId"`
		SupplierSource       int    `json:"supplierSource"`
		SupplierBusinessMode string `json:"supplierBusinessMode"`
	} `json:"info"`
}
