package tiktok

type SellerSearchCreatorOnMarketplaceRsp struct {
	BaseRsp
	Data struct {
		Creators []struct {
			Avatar struct {
				Url string `json:"url"`
			} `json:"avatar"`
			AvgEcLiveUv         int      `json:"avg_ec_live_uv"`
			AvgEcVideoViewCount int      `json:"avg_ec_video_view_count"`
			CategoryIds         []string `json:"category_ids"`
			FollowerCount       int      `json:"follower_count"`
			Gmv                 struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"gmv"`
			GmvRange struct {
				Currency      string `json:"currency"`
				MaximumAmount string `json:"maximum_amount"`
				MinimumAmount string `json:"minimum_amount"`
			} `json:"gmv_range"`
			LiveGmv struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"live_gmv"`
			Nickname                string `json:"nickname"`
			SelectionRegion         string `json:"selection_region"`
			TopFollowerDemographics struct {
				AgeRanges   []string `json:"age_ranges"`
				MajorGender struct {
					Gender     string `json:"gender"`
					Percentage int    `json:"percentage"`
				} `json:"major_gender"`
			} `json:"top_follower_demographics"`
			UnitsSoldRange struct {
				MaximumAmount int `json:"maximum_amount"`
				MinimumAmount int `json:"minimum_amount"`
			} `json:"units_sold_range"`
			Username string `json:"username"`
			VideoGmv struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"video_gmv"`
		} `json:"creators"`
		NextPageToken string `json:"next_page_token"`
		SearchKey     string `json:"search_key"`
	} `json:"data"`
}

type GetMarketplaceCreatorPerformanceRsp struct {
	BaseRsp
	Data struct {
		Creator struct {
			Avatar struct {
				Url string `json:"url"`
			} `json:"avatar"`
			AvgCommissionRate      int `json:"avg_commission_rate"`
			AvgCommissionRateRange struct {
				MaximumAmount int `json:"maximum_amount"`
				MinimumAmount int `json:"minimum_amount"`
			} `json:"avg_commission_rate_range"`
			AvgEcLiveCommentCount  int `json:"avg_ec_live_comment_count"`
			AvgEcLiveLikeCount     int `json:"avg_ec_live_like_count"`
			AvgEcLiveShareCount    int `json:"avg_ec_live_share_count"`
			AvgEcLiveViewCount     int `json:"avg_ec_live_view_count"`
			AvgEcVideoCommentCount int `json:"avg_ec_video_comment_count"`
			AvgEcVideoLikeCount    int `json:"avg_ec_video_like_count"`
			AvgEcVideoPlayCount    int `json:"avg_ec_video_play_count"`
			AvgEcVideoShareCount   int `json:"avg_ec_video_share_count"`
			AvgGmvPerBuyer         struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"avg_gmv_per_buyer"`
			AvgGmvPerBuyerRange struct {
				Currency      string `json:"currency"`
				MaximumAmount string `json:"maximum_amount"`
				MinimumAmount string `json:"minimum_amount"`
			} `json:"avg_gmv_per_buyer_range"`
			BioDescription          string `json:"bio_description"`
			BrandCollaborationCount int    `json:"brand_collaboration_count"`
			CategoryGmvDistribution []struct {
				CategoryId string `json:"category_id"`
				Value      string `json:"value"`
			} `json:"category_gmv_distribution"`
			CategoryIds            []string `json:"category_ids"`
			ContentGmvDistribution []struct {
				ContentType string `json:"content_type"`
				Value       string `json:"value"`
			} `json:"content_gmv_distribution"`
			EcLiveCount          int    `json:"ec_live_count"`
			EcLiveEngagementRate string `json:"ec_live_engagement_rate"`
			EcVideoCount         int    `json:"ec_video_count"`
			FollowerCount        int    `json:"follower_count"`
			Gmv                  struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"gmv"`
			GmvRange struct {
				Currency      string `json:"currency"`
				MaximumAmount string `json:"maximum_amount"`
				MinimumAmount string `json:"minimum_amount"`
			} `json:"gmv_range"`
			Gpm struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"gpm"`
			GpmRange struct {
				Currency      string `json:"currency"`
				MaximumAmount string `json:"maximum_amount"`
				MinimumAmount string `json:"minimum_amount"`
			} `json:"gpm_range"`
			LiveGmv struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"live_gmv"`
			LiveGpm struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"live_gpm"`
			LiveGpmRange struct {
				Currency      string `json:"currency"`
				MaximumAmount string `json:"maximum_amount"`
				MinimumAmount string `json:"minimum_amount"`
			} `json:"live_gpm_range"`
			Nickname                  string `json:"nickname"`
			ProductOriginalPriceRange struct {
				Currency      string `json:"currency"`
				MaximumAmount string `json:"maximum_amount"`
				MinimumAmount string `json:"minimum_amount"`
			} `json:"product_original_price_range"`
			ProfileTtUri            string   `json:"profile_tt_uri"`
			PromotedProductNum      int      `json:"promoted_product_num"`
			SelectionRegion         string   `json:"selection_region"`
			TopCollaboratedBrandIds []string `json:"top_collaborated_brand_ids"`
			UnitsSold               int      `json:"units_sold"`
			UnitsSoldRange          struct {
				MaximumAmount int `json:"maximum_amount"`
				MinimumAmount int `json:"minimum_amount"`
			} `json:"units_sold_range"`
			Username string `json:"username"`
			VideoGmv struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"video_gmv"`
			VideoGpm struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"video_gpm"`
			VideoGpmRange struct {
				Currency      string `json:"currency"`
				MaximumAmount string `json:"maximum_amount"`
				MinimumAmount string `json:"minimum_amount"`
			} `json:"video_gpm_range"`
		} `json:"creator"`
	} `json:"data"`
}

type CreateTargetCollaborationRsp struct {
	BaseRsp
	Data struct {
		TargetCollaboration struct {
			Id string `json:"id"`
		} `json:"target_collaboration"`
		TargetCollaborationConflicts []struct {
			CreatorUserId string `json:"creator_user_id"`
			ProductId     string `json:"product_id"`
		} `json:"target_collaboration_conflicts"`
	} `json:"data"`
}

type EditOpenCollaborationSettingsRsp struct {
	BaseRsp
}

type RemoveCreatorAffiliateFromCollaborationRsp struct {
	BaseRsp
}

type GenerateAffiliateProductPromotionLinkRsp struct {
	BaseRsp
	Data struct {
		ProductPromotionLink string `json:"product_promotion_link"`
	} `json:"data"`
}

type SellerSearchAffiliateOpenCollaborationProductRsp struct {
	BaseRsp
	Data struct {
		NextPageToken string `json:"next_page_token"`
		Products      []struct {
			CategoryChains []struct {
				Id        string `json:"id"`
				IsLeaf    bool   `json:"is_leaf"`
				LocalName string `json:"local_name"`
				ParentId  string `json:"parent_id"`
			} `json:"category_chains"`
			Commission struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
				Rate     int    `json:"rate"`
			} `json:"commission"`
			DetailLink    string `json:"detail_link"`
			HasInventory  bool   `json:"has_inventory"`
			Id            string `json:"id"`
			MainImageUrl  string `json:"main_image_url"`
			OriginalPrice struct {
				Currency      string `json:"currency"`
				MaximumAmount string `json:"maximum_amount"`
				MinimumAmount string `json:"minimum_amount"`
			} `json:"original_price"`
			SaleRegion string `json:"sale_region"`
			SalesPrice struct {
				Currency      string `json:"currency"`
				MaximumAmount string `json:"maximum_amount"`
				MinimumAmount string `json:"minimum_amount"`
			} `json:"sales_price"`
			Shop struct {
				Name string `json:"name"`
			} `json:"shop"`
			Title     string `json:"title"`
			UnitsSold int    `json:"units_sold"`
		} `json:"products"`
		TotalCount int `json:"total_count"`
	} `json:"data"`
}

type SearchSellerAffiliateOrdersRsp struct {
	BaseRsp
	Data struct {
		NextPageToken string `json:"next_page_token"`
		Orders        []struct {
			CreateTime   int64  `json:"create_time"`
			DeliveryTime int64  `json:"delivery_time"`
			Id           string `json:"id"`
			Skus         []struct {
				ActualCommissionBase struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"actual_commission_base"`
				ActualPaidCommission struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"actual_paid_commission"`
				ActualPaidShopAdsCommission struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"actual_paid_shop_ads_commission"`
				CampaignId              string `json:"campaign_id"`
				CommissionRate          string `json:"commission_rate"`
				ContentId               string `json:"content_id"`
				ContentType             string `json:"content_type"`
				CreatorUsername         string `json:"creator_username"`
				EstimatedCommissionBase struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"estimated_commission_base"`
				EstimatedPaidCommission struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"estimated_paid_commission"`
				EstimatedPaidShopAdsCommission struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"estimated_paid_shop_ads_commission"`
				OpenCollaborationId string `json:"open_collaboration_id"`
				Price               struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"price"`
				ProductId             string `json:"product_id"`
				Quantity              int    `json:"quantity"`
				RefundedQuantity      int    `json:"refunded_quantity"`
				ReturnedQuantity      int    `json:"returned_quantity"`
				ShopAdsCommissionRate string `json:"shop_ads_commission_rate"`
				TargetCollaborationId string `json:"target_collaboration_id"`
			} `json:"skus"`
			Status string `json:"status"`
		} `json:"orders"`
		TotalCount int `json:"total_count"`
	} `json:"data"`
}

type SellerSearchSampleApplicationsFulfillmentsRsp struct {
	BaseRsp
	Data struct {
		Fulfillments []struct {
			Content struct {
				CommentCount   int    `json:"comment_count"`
				CreateTime     int64  `json:"create_time"`
				Description    string `json:"description"`
				Format         string `json:"format"`
				Id             string `json:"id"`
				LikeCount      int    `json:"like_count"`
				LiveEndTime    int64  `json:"live_end_time"`
				PageLink       string `json:"page_link"`
				PaidOrderCount int    `json:"paid_order_count"`
				Url            string `json:"url"`
				ViewCount      int    `json:"view_count"`
			} `json:"content"`
			Product struct {
				Id           string `json:"id"`
				MainImageUrl string `json:"main_image_url"`
			} `json:"product"`
		} `json:"fulfillments"`
	} `json:"data"`
}

type SellerReviewSampleApplicationsRsp struct {
	BaseRsp
}

type GetOpenCollaborationSampleRulesRsp struct {
	BaseRsp
	Data struct {
		SampleRules []struct {
			AvailableQuantity     int    `json:"available_quantity"`
			EndTime               int64  `json:"end_time"`
			IsSampleTimeUnlimited bool   `json:"is_sample_time_unlimited"`
			ProductId             string `json:"product_id"`
			SampleQuota           int    `json:"sample_quota"`
			StartTime             int64  `json:"start_time"`
			Status                string `json:"status"`
			Thresholds            struct {
				AvgEcVideoViews          int      `json:"avg_ec_video_views"`
				CategoryIds              []string `json:"category_ids"`
				MinimumFollowerCount     int      `json:"minimum_follower_count"`
				MinimumGmv               int      `json:"minimum_gmv"`
				PredictedFulfillmentRank string   `json:"predicted_fulfillment_rank"`
			} `json:"thresholds"`
		} `json:"sample_rules"`
	} `json:"data"`
}

type SellerSearchSampleApplicationsRsp struct {
	BaseRsp
	Data struct {
		NextPageToken      string `json:"next_page_token"`
		SampleApplications []struct {
			ApproveExpirationTime int64  `json:"approve_expiration_time"`
			AvailableQuantity     int    `json:"available_quantity"`
			CommissionRate        string `json:"commission_rate"`
			Creator               struct {
				AvatarUrl             string `json:"avatar_url"`
				ContentCount          int    `json:"content_count"`
				EcVideoView           int    `json:"ec_video_view"`
				FollowerCount         int    `json:"follower_count"`
				FulfillmentPercentage string `json:"fulfillment_percentage"`
				Gmv                   struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"gmv"`
				Nickname string `json:"nickname"`
				UserId   string `json:"user_id"`
				Username string `json:"username"`
			} `json:"creator"`
			DisapprovableReasons []string `json:"disapprovable_reasons"`
			FulfillmentStatus    string   `json:"fulfillment_status"`
			Id                   string   `json:"id"`
			IsApprovable         bool     `json:"is_approvable"`
			OrderId              string   `json:"order_id"`
			PartnerName          string   `json:"partner_name"`
			Product              struct {
				Id          string `json:"id"`
				SkuId       string `json:"sku_id"`
				SkuImageUrl string `json:"sku_image_url"`
				SkuName     string `json:"sku_name"`
				Title       string `json:"title"`
			} `json:"product"`
			ShipmentExpirationTime int    `json:"shipment_expiration_time"`
			Status                 string `json:"status"`
			TrackingNumber         string `json:"tracking_number"`
		} `json:"sample_applications"`
		TotalCount int `json:"total_count"`
	} `json:"data"`
}

type EditOpenCollaborationSampleRuleRsp struct {
	BaseRsp
}

type RemoveTargetCollaborationRsp struct {
	BaseRsp
}

type SearchTargetCollaborationsRsp struct {
	BaseRsp
	Data struct {
		NextPageToken        string `json:"next_page_token"`
		TargetCollaborations []struct {
			ContentCreatorCount  int   `json:"content_creator_count"`
			CreatorInivitedCount int   `json:"creator_inivited_count"`
			EndTime              int64 `json:"end_time"`
			FreeSampleRule       struct {
				HasFreeSample          bool `json:"has_free_sample"`
				IsSampleApprovalExempt bool `json:"is_sample_approval_exempt"`
			} `json:"free_sample_rule"`
			Id                   string `json:"id"`
			Message              string `json:"message"`
			Name                 string `json:"name"`
			ProductCount         int    `json:"product_count"`
			ShowcaseCreatorCount int    `json:"showcase_creator_count"`
			StartTime            int64  `json:"start_time"`
			Type                 string `json:"type"`
			UpdateTime           int64  `json:"update_time"`
		} `json:"target_collaborations"`
		TotalCount int `json:"total_count"`
	} `json:"data"`
}

type UpdateTargetCollaborationRsp struct {
	BaseRsp
	Data struct {
		TargetCollaborationConflicts []struct {
			CreatorUserId string `json:"creator_user_id"`
			ProductId     string `json:"product_id"`
		} `json:"target_collaboration_conflicts"`
		UpdateFailed struct {
			AddCreatorIds []string `json:"add_creator_ids"`
			AddProducts   struct {
				CommissionRate int    `json:"commission_rate"`
				Id             string `json:"id"`
			} `json:"add_products"`
			ChangeCommissions struct {
				CommissionRate int    `json:"commission_rate"`
				ProductId      string `json:"product_id"`
			} `json:"change_commissions"`
			EndTime           int64    `json:"end_time"`
			Name              string   `json:"name"`
			RemoveCreatorIds  []string `json:"remove_creator_ids"`
			RemoveProductIds  []string `json:"remove_product_ids"`
			SellerContactInfo struct {
				Email string `json:"email"`
			} `json:"seller_contact_info"`
		} `json:"update_failed"`
	} `json:"data"`
}

type GetOpenCollaborationSettingsRsp struct {
	BaseRsp
	Data struct {
		OpenCollaborationSettings struct {
			AutoAddProduct struct {
				CommissionRate int  `json:"commission_rate"`
				Enable         bool `json:"enable"`
			} `json:"auto_add_product"`
		} `json:"open_collaboration_settings"`
	} `json:"data"`
}

type RemoveOpenCollaborationRsp struct {
	BaseRsp
	Data struct {
		TerminatedEffectiveTime int64 `json:"terminated_effective_time"`
	} `json:"data"`
}

type QueryTargetCollaborationDetailRsp struct {
	BaseRsp
	Data struct {
		TargetCollaboration struct {
			ContentCreatorCount int `json:"content_creator_count"`
			CreatorInvitedCount int `json:"creator_invited_count"`
			Creators            []struct {
				Avatar struct {
					Url string `json:"url"`
				} `json:"avatar"`
				CollaborationStatus    string `json:"collaboration_status"`
				ContentProductCount    int    `json:"content_product_count"`
				Nickname               string `json:"nickname"`
				ProductEffectiveStatus string `json:"product_effective_status"`
				SelectionRegion        string `json:"selection_region"`
				ShowcaseProductCount   int    `json:"showcase_product_count"`
				Username               string `json:"username"`
			} `json:"creators"`
			EndTime        int64 `json:"end_time"`
			FreeSampleRule struct {
				HasFreeSample          bool `json:"has_free_sample"`
				IsSampleApprovalExempt bool `json:"is_sample_approval_exempt"`
			} `json:"free_sample_rule"`
			Id           string `json:"id"`
			Message      string `json:"message"`
			Name         string `json:"name"`
			ProductCount int    `json:"product_count"`
			Products     []struct {
				CollaborationStatus string `json:"collaboration_status"`
				Commission          struct {
					Currency      string `json:"currency"`
					EffectiveTime string `json:"effective_time"`
					MaximumAmount string `json:"maximum_amount"`
					MinimumAmount string `json:"minimum_amount"`
					Rate          int    `json:"rate"`
				} `json:"commission"`
				CommissionEffectiveStatus string `json:"commission_effective_status"`
				Id                        string `json:"id"`
				MainImageUrl              string `json:"main_image_url"`
				OriginalPrice             struct {
					Currency      string `json:"currency"`
					MaximumAmount string `json:"maximum_amount"`
					MinimumAmount string `json:"minimum_amount"`
				} `json:"original_price"`
				Status string `json:"status"`
				Title  string `json:"title"`
			} `json:"products"`
			SellerContactInfo struct {
				Email string `json:"email"`
			} `json:"seller_contact_info"`
			ShowcaseCreatorCount int    `json:"showcase_creator_count"`
			StartTime            int64  `json:"start_time"`
			Type                 string `json:"type"`
			UpdateTime           int64  `json:"update_time"`
		} `json:"target_collaboration"`
	} `json:"data"`
}

type GetOpenCollaborationCreatorContentDetailRsp struct {
	BaseRsp
	Data struct {
		CreatorContentDetails []struct {
			CreatorProfile struct {
				Avatar struct {
					Url string `json:"url"`
				} `json:"avatar"`
				FollowerCount int    `json:"follower_count"`
				Nickname      string `json:"nickname"`
				Username      string `json:"username"`
			} `json:"creator_profile"`
			LiveCount        int    `json:"live_count"`
			PromotionEndTime int64  `json:"promotion_end_time"`
			PromotionStatus  string `json:"promotion_status"`
			VideoCount       int    `json:"video_count"`
		} `json:"creator_content_details"`
		NextPageToken string `json:"next_page_token"`
		Product       struct {
			Id       string `json:"id"`
			ImageUrl string `json:"image_url"`
		} `json:"product"`
		TotalCount int `json:"total_count"`
	} `json:"data"`
}

type SearchOpenCollaborationRsp struct {
	BaseRsp
	Data struct {
		NextPageToken      string `json:"next_page_token"`
		OpenCollaborations []struct {
			ContentCreatorCount int `json:"content_creator_count"`
			CurrentCommission   struct {
				EndTime   int64 `json:"end_time"`
				Rate      int   `json:"rate"`
				StartTime int64 `json:"start_time"`
			} `json:"current_commission"`
			Id      string `json:"id"`
			Product struct {
				Id            string `json:"id"`
				Inventory     int    `json:"inventory"`
				MainImageUrl  string `json:"main_image_url"`
				OriginalPrice struct {
					Currency      string `json:"currency"`
					MaximumAmount string `json:"maximum_amount"`
					MinimumAmount string `json:"minimum_amount"`
				} `json:"original_price"`
				Status string `json:"status"`
				Title  string `json:"title"`
			} `json:"product"`
			ShowcaseCreatorCount int    `json:"showcase_creator_count"`
			Status               string `json:"status"`
		} `json:"open_collaborations"`
		TotalCount int `json:"total_count"`
	} `json:"data"`
}

type CreateOpenCollaborationRsp struct {
	BaseRsp
	Data struct {
		OpenCollaboration struct {
			EffectiveTime int64  `json:"effective_time"`
			Id            string `json:"id"`
			ProductId     string `json:"product_id"`
		} `json:"open_collaboration"`
	} `json:"data"`
}
