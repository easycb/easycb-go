package shopee

type GetReturnDetailRsp struct {
	BaseRsp
	Response struct {
		Image       []string `json:"image"`
		BuyerVideos []struct {
			ThumbnailUrl string `json:"thumbnail_url"`
			VideoUrl     string `json:"video_url"`
		} `json:"buyer_videos"`
		Reason               string  `json:"reason"`
		TextReason           string  `json:"text_reason"`
		ReturnSn             string  `json:"return_sn"`
		RefundAmount         float64 `json:"refund_amount"`
		Currency             string  `json:"currency"`
		CreateTime           int64   `json:"create_time"`
		UpdateTime           int64   `json:"update_time"`
		Status               string  `json:"status"`
		DueDate              int64   `json:"due_date"`
		TrackingNumber       string  `json:"tracking_number"`
		DisputeReason        int     `json:"dispute_reason"`
		DisputeTextReason    string  `json:"dispute_text_reason"`
		NeedsLogistics       bool    `json:"needs_logistics"`
		AmountBeforeDiscount float64 `json:"amount_before_discount"`
		User                 struct {
			Username string `json:"username"`
			Email    string `json:"email"`
			Portrait string `json:"portrait"`
		} `json:"user"`
		Item []struct {
			ModelId      int      `json:"model_id"`
			Name         string   `json:"name"`
			Images       []string `json:"images"`
			Amount       float64  `json:"amount"`
			ItemPrice    float64  `json:"item_price"`
			IsAddOnDeal  bool     `json:"is_add_on_deal"`
			IsMainItem   bool     `json:"is_main_item"`
			AddOnDealId  int64    `json:"add_on_deal_id"`
			ItemId       int64    `json:"item_id"`
			ItemSku      string   `json:"item_sku"`
			VariationSku string   `json:"variation_sku"`
			RefundAmount float64  `json:"refund_amount"`
		} `json:"item"`
		OrderSn             string `json:"order_sn"`
		ReturnShipDueDate   int64  `json:"return_ship_due_date"`
		ReturnSellerDueDate int64  `json:"return_seller_due_date"`
		Activity            []struct {
			ActivityId      int64  `json:"activity_id"`
			ActivityType    string `json:"activity_type"`
			OriginalPrice   string `json:"original_price"`
			DiscountedPrice string `json:"discounted_price"`
			Items           []struct {
				ItemId            int64  `json:"item_id"`
				VariationId       int    `json:"variation_id"`
				QuantityPurchased int    `json:"quantity_purchased"`
				OriginalPrice     string `json:"original_price"`
			} `json:"items"`
			RefundAmount float64 `json:"refund_amount"`
		} `json:"activity"`
		SellerProof struct {
			SellerProofStatus      string `json:"seller_proof_status"`
			SellerEvidenceDeadline int64  `json:"seller_evidence_deadline"`
		} `json:"seller_proof"`
		SellerCompensation struct {
			SellerCompensationStatus  string  `json:"seller_compensation_status"`
			SellerCompensationDueDate int64   `json:"seller_compensation_due_date"`
			CompensationAmount        float64 `json:"compensation_amount"`
		} `json:"seller_compensation"`
		Negotiation struct {
			NegotiationStatus  string  `json:"negotiation_status"`
			LatestSolution     string  `json:"latest_solution"`
			LatestOfferAmount  float64 `json:"latest_offer_amount"`
			LatestOfferCreator string  `json:"latest_offer_creator"`
			CounterLimit       int64   `json:"counter_limit"`
			OfferDueDate       int64   `json:"offer_due_date"`
		} `json:"negotiation"`
		LogisticsStatus     string `json:"logistics_status"`
		ReturnPickupAddress struct {
			Address  string `json:"address"`
			Name     string `json:"name"`
			Phone    string `json:"phone"`
			Town     string `json:"town"`
			District string `json:"district"`
			City     string `json:"city"`
			State    string `json:"state"`
			Region   string `json:"region"`
			Zipcode  string `json:"zipcode"`
		} `json:"return_pickup_address"`
		ReturnBuyerContactInfo struct {
			VirtualContactNumber string `json:"virtual_contact_number"`
			PackageQueryNumber   string `json:"package_query_number"`
		} `json:"return_buyer_contact_info"`
		ReturnAddress struct {
			WhsId string `json:"whs_id"`
		} `json:"return_address"`
	} `json:"response"`
}

type GetReturnListRsp struct {
	BaseRsp
	Response []struct {
		More   bool `json:"more"`
		Return []struct {
			Image                []string `json:"image"`
			Reason               string   `json:"reason"`
			TextReason           string   `json:"text_reason"`
			ReturnSn             string   `json:"return_sn"`
			RefundAmount         float64  `json:"refund_amount"`
			Currency             string   `json:"currency"`
			CreateTime           int64    `json:"create_time"`
			UpdateTime           int64    `json:"update_time"`
			Status               string   `json:"status"`
			DueDate              int64    `json:"due_date"`
			TrackingNumber       string   `json:"tracking_number"`
			DisputeReason        []string `json:"dispute_reason"`
			DisputeTextReason    []string `json:"dispute_text_reason"`
			NeedsLogistics       bool     `json:"needs_logistics"`
			AmountBeforeDiscount float64  `json:"amount_before_discount"`
			User                 struct {
				Username string `json:"username"`
				Email    string `json:"email"`
				Portrait string `json:"portrait"`
			} `json:"user"`
			Item []struct {
				ModelId      int      `json:"model_id"`
				Name         string   `json:"name"`
				Images       []string `json:"images"`
				Amount       float64  `json:"amount"`
				ItemPrice    float64  `json:"item_price"`
				IsAddOnDeal  bool     `json:"is_add_on_deal"`
				IsMainItem   bool     `json:"is_main_item"`
				AddOnDealId  int64    `json:"add_on_deal_id"`
				ItemId       int64    `json:"item_id"`
				ItemSku      string   `json:"item_sku"`
				VariationSku string   `json:"variation_sku"`
			} `json:"item"`
			OrderSn                  string `json:"order_sn"`
			ReturnShipDueDate        int64  `json:"return_ship_due_date"`
			ReturnSellerDueDate      int64  `json:"return_seller_due_date"`
			NegotiationStatus        string `json:"negotiation_status"`
			SellerProofStatus        string `json:"seller_proof_status"`
			SellerCompensationStatus string `json:"seller_compensation_status"`
		} `json:"return"`
	} `json:"response"`
}

type ReturnConfirmRsp struct {
	BaseRsp
	Response struct {
		ReturnSn string `json:"return_sn"`
	} `json:"response"`
}

type ReturnDisputeRsp struct {
	BaseRsp
	Response struct {
		ReturnSn string `json:"return_sn"`
		Msg      string `json:"msg"`
	} `json:"response"`
}

type GetAvailableSolutionsRsp struct {
	BaseRsp
	Response struct {
		ReturnSn          string `json:"return_sn"`
		OfferReturnRefund struct {
			Eligibility            bool    `json:"eligibility"`
			RefundAmountAdjustable bool    `json:"refund_amount_adjustable"`
			MaxRefundAmount        float64 `json:"max_refund_amount"`
		} `json:"offer_return_refund"`
		OfferRefund struct {
			Eligibility            bool    `json:"eligibility"`
			RefundAmountAdjustable bool    `json:"refund_amount_adjustable"`
			MaxRefundAmount        float64 `json:"max_refund_amount"`
		} `json:"offer_refund"`
	} `json:"response"`
}

type ReturnOfferRsp struct {
	BaseRsp
	Response struct {
		ReturnSn string `json:"return_sn"`
	} `json:"response"`
}

type ReturnAcceptOfferRsp struct {
	BaseRsp
	Response struct {
		ReturnSn string `json:"return_sn"`
	} `json:"response"`
}

type ConvertImageRsp struct {
	BaseRsp
	Response struct {
		Url       string `json:"url"`
		Thumbnail string `json:"thumbnail"`
	} `json:"response"`
}

type UploadProofRsp struct {
	BaseRsp
}

type QueryProofRsp struct {
	BaseRsp
	Response struct {
		Image []struct {
			Url       string `json:"url"`
			Thumbnail string `json:"thumbnail"`
		} `json:"image"`
		Video []struct {
			Url       string `json:"url"`
			Thumbnail string `json:"thumbnail"`
		} `json:"video"`
		Description string `json:"description"`
	} `json:"response"`
}

type GetReturnDisputeReasonRsp struct {
	BaseRsp
	Response struct {
		DisputeReason []int `json:"dispute_reason"`
	} `json:"response"`
}
