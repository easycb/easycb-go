package ozon

type SupplyOrderBundleRsp struct {
	BaseRsp

	Items      []SupplyOrderBundleItem `json:"items"`
	TotalCount int                     `json:"total_count"`
	HasNext    bool                    `json:"has_next"`
	LastId     string                  `json:"last_id"`
}

type SupplyOrderBundleItem struct {
	IconPath            string   `json:"icon_path"`
	Sku                 int      `json:"sku"`
	Name                string   `json:"name"`
	OfferId             string   `json:"offer_id"`
	Quantity            int      `json:"quantity"`
	Barcode             string   `json:"barcode"`
	ProductId           int      `json:"product_id"`
	Quant               int      `json:"quant"`
	IsQuantEditable     bool     `json:"is_quant_editable"`
	VolumeInLitres      int      `json:"volume_in_litres"`
	TotalVolumeInLitres int      `json:"total_volume_in_litres"`
	ContractorItemCode  string   `json:"contractor_item_code"`
	SfboAttribute       string   `json:"sfbo_attribute"`
	ShipmentType        string   `json:"shipment_type"`
	Tags                []string `json:"tags"`
	PlacementZone       string   `json:"placement_zone"`
}
