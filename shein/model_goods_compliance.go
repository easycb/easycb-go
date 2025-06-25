package shein

type AgencyListRsp struct {
	BaseRsp
	Info []AgencyListItem `json:"info"`
}

type AgencyListItem struct {
	AgencyId            int      `json:"agencyId"`
	AgencyName          string   `json:"agencyName"`
	AgencyType          int      `json:"agencyType"`
	AgencySubType       int      `json:"agencySubType"`
	AgencyStartTime     *string  `json:"agencyStartTime"`
	AgencyEndTime       *string  `json:"agencyEndTime"`
	AgencyStatus        int      `json:"agencyStatus"`
	ApplyFailureReason  []string `json:"applyFailureReason"`
	ApplyStatus         int      `json:"applyStatus"`
	CoveredProductRange int      `json:"coveredProductRange"`
	CreateTime          string   `json:"createTime"`
	UpdateTime          string   `json:"updateTime"`
	SupplierId          int      `json:"supplierId"`
}

type SkcAgencyDetailRsp struct {
	BaseRsp
	Info []SkcAgencyDetailItem `json:"info"`
}

type SkcAgencyDetailItem struct {
	Skc         string `json:"skc"`
	ReviewState int    `json:"reviewState"`
	AgencyType  int    `json:"agencyType"`
}

type SaveSkcAgencyRsp struct {
	BaseRsp
}

type SkcLabelListRsp struct {
	BaseRsp
	Info []SkuLabelListItem `json:"info"`
}

type SkuLabelListItem struct {
	Skc              string      `json:"skc"`
	FailReasonList   interface{} `json:"failReasonList"`
	SkcShelfStatus   int         `json:"skcShelfStatus"`
	SkcLabelInfoList []struct {
		IsRequired   int      `json:"isRequired"`
		LabelId      int      `json:"labelId"`
		LabelName    string   `json:"labelName"`
		SiteList     []string `json:"siteList"`
		ReviewStatus int      `json:"reviewStatus"`
	} `json:"skcLabelInfoList"`
}

type MaterialQualityTreeV2Rsp struct {
	BaseRsp
	Info []MaterialQualityTreeV2Item `json:"info"`
}

type MaterialQualityTreeV2Item struct {
	PackageTypeId    int    `json:"packageTypeId"`
	PackageTypeName  string `json:"packageTypeName"`
	IsEnabled        int    `json:"isEnabled"`
	PackageMaterials []struct {
		PackageMaterialId   int    `json:"packageMaterialId"`
		PackageMaterialName string `json:"packageMaterialName"`
		IsEnabled           int    `json:"isEnabled"`
	} `json:"packageMaterials"`
}

type LabelPrintRsp struct {
	BaseRsp
	Info struct {
		SuccessList []struct {
			Uid         string `json:"uid"`
			LabelType   int    `json:"labelType"`
			LabelUrl    string `json:"labelUrl"`
			LabelWidth  int    `json:"labelWidth"`
			LabelLength int    `json:"labelLength"`
		} `json:"successList"`
		FailedList []struct {
			Uid       string   `json:"uid"`
			LabelType int      `json:"labelType"`
			Code      []string `json:"code"`
			Reason    []string `json:"reason"`
		} `json:"failedList"`
	} `json:"info"`
}

type UploadSkcLabelPictureRsp struct {
	BaseRsp
	Info UploadSkcLabelPictureItem `json:"info"`
}

type UploadSkcLabelPictureItem struct {
	ImageUrl string      `json:"imageUrl"`
	ImageMd5 string      `json:"imageMd5"`
	Msg      interface{} `json:"msg"`
	Code     int         `json:"code"`
}

type SkuSaveLabelRsp struct {
	BaseRsp
	Info struct {
		TotalCount   int `json:"totalCount"`
		SuccessCount int `json:"successCount"`
		FaildCount   int `json:"faildCount"`
		FaildList    []struct {
			Skc    string `json:"skc"`
			Code   string `json:"code"`
			Reason string `json:"reason"`
		} `json:"faildList"`
	} `json:"info"`
}
