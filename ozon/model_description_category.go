package ozon

type GetDescriptionCategoryTreeRsp struct {
	BaseRsp
	Result []GetDescriptionCategoryTreeResult `json:"result"`
}

type GetDescriptionCategoryTreeResult struct {
	DescriptionCategoryId int64  `json:"description_category_id"`
	CategoryName          string `json:"category_name"`
	Disabled              bool   `json:"disabled"`
	Children              []struct {
		DescriptionCategoryId int64                             `json:"description_category_id"`
		CategoryName          string                            `json:"category_name"`
		Disabled              bool                              `json:"disabled"`
		Children              []GetDescriptionCategoryTreeChild `json:"children"`
	} `json:"children"`
}

type GetDescriptionCategoryTreeChild struct {
	TypeName string                             `json:"type_name"`
	TypeId   int64                              `json:"type_id"`
	Disabled bool                               `json:"disabled"`
	Children []*GetDescriptionCategoryTreeChild `json:"children,omitempty"`
}

type GetDescriptionCategoryAttributeRsp struct {
	BaseRsp
	Result []GetDescriptionCategoryAttributeItem `json:"result"`
}

type GetDescriptionCategoryAttributeItem struct {
	Id                  int64  `json:"id"`
	AttributeComplexId  int64  `json:"attribute_complex_id"`
	Name                string `json:"name"`
	Description         string `json:"description"`
	Type                string `json:"type"`
	IsCollection        bool   `json:"is_collection"`
	IsRequired          bool   `json:"is_required"`
	IsAspect            bool   `json:"is_aspect"`
	MaxValueCount       int64  `json:"max_value_count"`
	GroupName           string `json:"group_name"`
	GroupId             int64  `json:"group_id"`
	DictionaryId        int64  `json:"dictionary_id"`
	CategoryDependent   bool   `json:"category_dependent"`
	ComplexIsCollection bool   `json:"complex_is_collection"`
}

type GetDescriptionCategoryAttributeValuesRsp struct {
	BaseRsp
	Result  []GetDescriptionCategoryAttributeValueItem `json:"result"`
	HasNext bool                                       `json:"has_next"`
}

type GetDescriptionCategoryAttributeValueItem struct {
	Id      int64  `json:"id"`
	Value   string `json:"value"`
	Info    string `json:"info"`
	Picture string `json:"picture"`
}

type GetDescriptionCategoryAttributeValuesSearchRsp struct {
	BaseRsp
	Result []GetDescriptionCategoryAttributeValuesSearchItem `json:"result"`
}

type GetDescriptionCategoryAttributeValuesSearchItem struct {
	Id      int64  `json:"id"`
	Info    string `json:"info"`
	Picture string `json:"picture"`
	Value   string `json:"value"`
}

type DescriptionCategoryTipsRsp struct {
	BaseRsp

	Result []DescriptionCategoryTipsItem `json:"result"`
}

type DescriptionCategoryTipsItem struct {
	ImagesUrl []string `json:"images_url"`
	InfoUrl   string   `json:"info_url"`
	TypeId    int64    `json:"type_id"`
}
