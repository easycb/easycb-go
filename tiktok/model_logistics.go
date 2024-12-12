package tiktok

type GetWarehouseListRsp struct {
	BaseRsp
	Data struct {
		Warehouses []struct {
			Address struct {
				City          string `json:"city"`
				ContactPerson string `json:"contact_person"`
				Distict       string `json:"distict"`
				FullAddress   string `json:"full_address"`
				Geolocation   struct {
					Latitude  string `json:"latitude"`
					Longitude string `json:"longitude"`
				} `json:"geolocation"`
				PhoneNumber string `json:"phone_number"`
				PostalCode  string `json:"postal_code"`
				Region      string `json:"region"`
				RegionCode  string `json:"region_code"`
				State       string `json:"state"`
				Town        string `json:"town"`
			} `json:"address"`
			EffectStatus string `json:"effect_status"`
			Id           string `json:"id"`
			IsDefault    bool   `json:"is_default"`
			Name         string `json:"name"`
			SubType      string `json:"sub_type"`
			Type         string `json:"type"`
		} `json:"warehouses"`
	} `json:"data"`
}

type GetGlobalSellerWarehouseRsp struct {
	BaseRsp
	Data struct {
		GlobalWarehouses []struct {
			Id        string `json:"id"`
			Name      string `json:"name"`
			Ownership string `json:"ownership"`
		} `json:"global_warehouses"`
	} `json:"data"`
}

type GetWarehouseDeliveryOptionsRsp struct {
	BaseRsp
	Data struct {
		DeliveryOptions []struct {
			Description    string `json:"description"`
			DimensionLimit struct {
				MaxHeight int    `json:"max_height"`
				MaxLength int    `json:"max_length"`
				MaxWidth  int    `json:"max_width"`
				Unit      string `json:"unit"`
			} `json:"dimension_limit"`
			Id          string   `json:"id"`
			Name        string   `json:"name"`
			Platform    []string `json:"platform"`
			Type        string   `json:"type"`
			WeightLimit struct {
				MaxWeight int    `json:"max_weight"`
				MinWeight int    `json:"min_weight"`
				Unit      string `json:"unit"`
			} `json:"weight_limit"`
		} `json:"delivery_options"`
	} `json:"data"`
}

type GetShippingProvidersRsp struct {
	BaseRsp
	Data struct {
		ShippingProviders []struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"shipping_providers"`
	} `json:"data"`
}
