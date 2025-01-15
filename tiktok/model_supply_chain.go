package tiktok

type ConfirmPackageShipmentRsp struct {
	BaseRsp
	Data struct {
		Errors []struct {
			Code   string `json:"code"`
			Detail struct {
				PackageId string `json:"package_id"`
			} `json:"detail"`
			Message string `json:"message"`
		} `json:"errors"`
		SuccessPackages []string `json:"success_packages"`
	} `json:"data"`
}
