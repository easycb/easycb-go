package tiktok

type BaseWebhook struct {
	Type              int         `json:"type"`
	Timestamp         int64       `json:"timestamp"`
	TtsNotificationId string      `json:"tts_notification_id"`
	Data              interface{} `json:"data"`
}

// Affiliate Creator

type WebhookCreatorDeauthorization struct {
	BaseWebhook
	CreatorOpenId string                            `json:"creator_open_id"`
	Data          WebhookCreatorDeauthorizationData `json:"data"`
}

type WebhookCreatorDeauthorizationData struct {
	CancelTime int64 `json:"cancel_time"`
}

type WebhookShoppableContentPosting struct {
	BaseWebhook
	CreatorOpenId string                             `json:"creator_open_id"`
	Data          WebhookShoppableContentPostingData `json:"data"`
}

type WebhookShoppableContentPostingData struct {
	ProductId string `json:"product_id"`
	Event     struct {
		Type      string `json:"type"`
		Scene     string `json:"scene"`
		ContentId string `json:"content_id"`
	} `json:"event"`
}

// Customer Service

type WebhookNewConversation struct {
	BaseWebhook
	ShopId string                     `json:"shop_id"`
	Data   WebhookNewConversationData `json:"data"`
}

type WebhookNewConversationData struct {
	ConversationId string `json:"conversation_id"`
	CreateTime     int64  `json:"create_time"`
}

type WebhookNewMessage struct {
	BaseWebhook
	ShopId string                `json:"shop_id"`
	Data   WebhookNewMessageData `json:"data"`
}

type WebhookNewMessageData struct {
	Content        string `json:"content"`
	ConversationId string `json:"conversation_id"`
	CreateTime     int64  `json:"create_time"`
	IsVisible      bool   `json:"is_visible"`
	MessageId      string `json:"message_id"`
	Index          string `json:"index"`
	Type           string `json:"type"`
	Sender         struct {
		ImUserId string `json:"im_user_id"`
		Role     string `json:"role"`
	} `json:"sender"`
}

// Fullfilled by TikTok

type WebhookGoodsMatch struct {
	BaseWebhook
	SellerOpenId string                `json:"seller_open_id"`
	Data         WebhookGoodsMatchData `json:"data"`
}

type WebhookGoodsMatchData struct {
	MatchType   string `json:"match_type"`
	ProductType string `json:"product_type"`
	GoodsId     string `json:"goods_id"`
	ProductId   string `json:"product_id"`
	SkuId       string `json:"sku_id"`
	UpdateTime  int64  `json:"update_time"`
}

type WebhookInboundFBTOrderStatusChange struct {
	BaseWebhook
	SellerOpenId string                                 `json:"seller_open_id"`
	Data         WebhookInboundFBTOrderStatusChangeData `json:"data"`
}

type WebhookInboundFBTOrderStatusChangeData struct {
	InboundOrderId string `json:"inbound_order_id"`
	OrderStatus    string `json:"order_status"`
	UpdateTime     int64  `json:"update_time"`
}

type WebhookFBTInventoryUpdate struct {
	BaseWebhook
	SellerOpenId string                        `json:"seller_open_id"`
	Data         WebhookFBTInventoryUpdateData `json:"data"`
}

type WebhookFBTInventoryUpdateData struct {
	GoodsId               string                      `json:"goods_id"`
	SkuId                 string                      `json:"sku_id"`
	FbtWarehouseInventory []FbtWarehouseInventoryItem `json:"fbt_warehouse_inventory"`
	UpdateTime            int64                       `json:"update_time"`
}

type FbtWarehouseInventoryItem struct {
	FbtWarehouseId string `json:"fbt_warehouse_id"`
	OnHandDetail   struct {
		TotalQuantity     int `json:"total_quantity"`
		ReservedQuantity  int `json:"reserved_quantity"`
		AvailableQuantity int `json:"available_quantity"`
	} `json:"on_hand_detail"`
}

type WebhookFBTMerchantOnboarding struct {
	BaseWebhook
	SellerOpenId string                           `json:"seller_open_id"`
	Data         WebhookFBTMerchantOnboardingData `json:"data"`
}

type WebhookFBTMerchantOnboardingData struct {
	OnboardedRegions []FBTMerchantOnboardingRegionItem `json:"onboarded_regions"`
	UpdateTime       int64                             `json:"update_time"`
}

type FBTMerchantOnboardingRegionItem struct {
	RegionCode string `json:"region_code"`
}

// Logistics

type WebhookPackageUpdate struct {
	BaseWebhook
	ShopId string                   `json:"shop_id"`
	Data   WebhookPackageUpdateData `json:"data"`
}

type WebhookPackageUpdateData struct {
	ScType      string                     `json:"sc_type"`
	RoleType    string                     `json:"role_type"`
	PackageList []PackageUpdatePackageItem `json:"package_list"`
	UpdateTime  int64                      `json:"update_time"`
}

type PackageUpdatePackageItem struct {
	PackageId   string   `json:"package_id"`
	OrderIdList []string `json:"order_id_list"`
}

type WebhookRecipientAddressUpdate struct {
	BaseWebhook
	ShopId string                            `json:"shop_id"`
	Data   WebhookRecipientAddressUpdateData `json:"data"`
}

type WebhookRecipientAddressUpdateData struct {
	OrderId    string `json:"order_id"`
	UpdateTime int64  `json:"update_time"`
}

// Order

type WebhookCancellationStatusChange struct {
	BaseWebhook
	ShopId string                              `json:"shop_id"`
	Data   WebhookCancellationStatusChangeData `json:"data"`
}

type WebhookCancellationStatusChangeData struct {
	OrderId           string `json:"order_id"`
	CancellationsRole string `json:"cancellations_role"`
	CancelStatus      string `json:"cancel_status"`
	CancelId          string `json:"cancel_id"`
	CreateTime        int64  `json:"create_time"`
}

type WebhookOrderStatusChange struct {
	BaseWebhook
	ShopId string                       `json:"shop_id"`
	Data   WebhookOrderStatusChangeData `json:"data"`
}

type WebhookOrderStatusChangeData struct {
	OrderId       string `json:"order_id"`
	OrderStatus   string `json:"order_status"`
	IsOnHoldOrder bool   `json:"is_on_hold_order"`
	UpdateTime    int64  `json:"update_time"`
}

// Product

type WebhookInventoryStatusChange struct {
	BaseWebhook
	ShopId string                           `json:"shop_id"`
	Data   WebhookInventoryStatusChangeData `json:"data"`
}

type WebhookInventoryStatusChangeData struct {
	ProductId     string `json:"product_id"`
	SkuId         string `json:"sku_id"`
	TriggerReason struct {
		AlertType string `json:"alert_type"`
		LeadDays  int    `json:"lead_days"`
	} `json:"trigger_reason"`
	CurrentInventoryStatus string `json:"current_inventory_status"`
	InventoryDistribution  struct {
		TotalQuantity            int `json:"total_quantity"`
		AvailableQuantity        int `json:"available_quantity"`
		CreatorReservedQuantity  int `json:"creator_reserved_quantity"`
		CampaignReservedQuantity int `json:"campaign_reserved_quantity"`
		CommittedQuantity        int `json:"committed_quantity"`
	} `json:"inventory_distribution"`
	UpdateTime int64 `json:"update_time"`
}

type WebhookOpportunityMatchingStatusChange struct {
	BaseWebhook
	ShopId string                                     `json:"shop_id"`
	Data   WebhookOpportunityMatchingStatusChangeData `json:"data"`
}

type WebhookOpportunityMatchingStatusChangeData struct {
	ExternalProductId          string   `json:"external_product_id"`
	OpportunityMatchingStatus  string   `json:"opportunity_matching_status"`
	OpportunityIds             []string `json:"opportunity_ids"`
	OpportunityMatchingEndTime int      `json:"opportunity_matching_end_time"`
	UpdateTime                 int64    `json:"update_time"`
}

type WebhookProductCategoryChange struct {
	BaseWebhook
	ShopId string                           `json:"shop_id"`
	Data   WebhookProductCategoryChangeData `json:"data"`
}

type WebhookProductCategoryChangeData struct {
	ProductId          string `json:"product_id"`
	PreviousCategoryId string `json:"previous_category_id"`
	CurrentCategoryId  string `json:"current_category_id"`
	UpdateTime         int64  `json:"update_time"`
}

type WebhookProductCreation struct {
	BaseWebhook
	ShopId string                     `json:"shop_id"`
	Data   WebhookProductCreationData `json:"data"`
}

type WebhookProductCreationData struct {
	ProductId    string   `json:"product_id"`
	ProductTypes []string `json:"product_types"`
	UpdateTime   int64    `json:"update_time"`
}

type WebhookProductInformationChange struct {
	BaseWebhook
	ShopId string                              `json:"shop_id"`
	Data   WebhookProductInformationChangeData `json:"data"`
}

type WebhookProductInformationChangeData struct {
	ProductId     string   `json:"product_id"`
	ChangeSource  string   `json:"change_source"`
	ChangedFields []string `json:"changed_fields"`
	UpdateTime    int64    `json:"update_time"`
}

type WebhookProductStatusChange struct {
	BaseWebhook
	ShopId string                         `json:"shop_id"`
	Data   WebhookProductStatusChangeData `json:"data"`
}

type WebhookProductStatusChangeData struct {
	ProductId       int64  `json:"product_id"`
	Status          string `json:"status"`
	SuspendedReason string `json:"suspended_reason"`
	UpdateTime      int64  `json:"update_time"`
}

type WebhookSizeChartChange struct {
	BaseWebhook
	SellerOpenId string                     `json:"seller_open_id"`
	Data         WebhookSizeChartChangeData `json:"data"`
}

type WebhookSizeChartChangeData struct {
	SizeChartTemplateId   string `json:"size_chart_template_id"`
	SizeChartTemplateName string `json:"size_chart_template_name"`
	EventType             string `json:"event_type"`
	UpdateTime            int64  `json:"update_time"`
}

// Return and Refund

type WebhookReturnStatusChange struct {
	BaseWebhook
	ShopId string                        `json:"shop_id"`
	Data   WebhookReturnStatusChangeData `json:"data"`
}

type WebhookReturnStatusChangeData struct {
	OrderId      string `json:"order_id"`
	ReturnRole   string `json:"return_role"`
	ReturnType   string `json:"return_type"`
	ReturnStatus string `json:"return_status"`
	ReturnId     string `json:"return_id"`
	CreateTime   int64  `json:"create_time"`
	UpdateTime   int64  `json:"update_time"`
}

// Seller

type WebhookSellerDeauthorization struct {
	BaseWebhook
	ShopId string                           `json:"shop_id"`
	Data   WebhookSellerDeauthorizationData `json:"data"`
}

type WebhookSellerDeauthorizationData struct {
	Message string `json:"message"`
}

type WebhookUpcomingAuthorizationExpiration struct {
	BaseWebhook
	ShopId string                                     `json:"shop_id"`
	Data   WebhookUpcomingAuthorizationExpirationData `json:"data"`
}

type WebhookUpcomingAuthorizationExpirationData struct {
	Message        string `json:"message"`
	ExpirationTime string `json:"expiration_time"`
}
