package tiktok

import (
	"fmt"
	"github.com/easycb/easycb-go"
)

func (c *Client) SellerSearchCreatorOnMarketplace(query easycb.AnyMap, body easycb.AnyMap) (*SellerSearchCreatorOnMarketplaceRsp, error) {
	var result SellerSearchCreatorOnMarketplaceRsp
	path := "/affiliate_seller/202406/marketplace_creators/search"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetMarketplaceCreatorPerformance(createUserId string, query easycb.AnyMap) (*GetMarketplaceCreatorPerformanceRsp, error) {
	var result GetMarketplaceCreatorPerformanceRsp
	path := fmt.Sprintf("/affiliate_seller/202406/marketplace_creators/%s", createUserId)
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CreateTargetCollaboration(query easycb.AnyMap, body easycb.AnyMap) (*CreateTargetCollaborationRsp, error) {
	var result CreateTargetCollaborationRsp
	path := "/affiliate_seller/202405/target_collaborations"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) EditOpenCollaborationSettings(query easycb.AnyMap, body easycb.AnyMap) (*EditOpenCollaborationSettingsRsp, error) {
	var result EditOpenCollaborationSettingsRsp
	path := "/affiliate_seller/202405/open_collaboration_settings"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) RemoveCreatorAffiliateFromCollaboration(openCollaborationId string, query easycb.AnyMap, body easycb.AnyMap) (*RemoveCreatorAffiliateFromCollaborationRsp, error) {
	var result RemoveCreatorAffiliateFromCollaborationRsp
	path := fmt.Sprintf("/affiliate_seller/202405/open_collaborations/%s/remove_creator", openCollaborationId)
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GenerateAffiliateProductPromotionLink(productId string, query easycb.AnyMap) (*GenerateAffiliateProductPromotionLinkRsp, error) {
	var result GenerateAffiliateProductPromotionLinkRsp
	path := fmt.Sprintf("/affiliate_seller/202405/products/%s/promotion_link/generate", productId)
	err := c.doRequest("POST", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SellerSearchAffiliateOpenCollaborationProduct(query easycb.AnyMap, body easycb.AnyMap) (*SellerSearchAffiliateOpenCollaborationProductRsp, error) {
	var result SellerSearchAffiliateOpenCollaborationProductRsp
	path := "/affiliate_seller/202405/open_collaborations/products/search"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SearchSellerAffiliateOrders(query easycb.AnyMap, body easycb.AnyMap) (*SearchSellerAffiliateOrdersRsp, error) {
	var result SearchSellerAffiliateOrdersRsp
	path := "/affiliate_seller/202410/orders/search"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SellerSearchSampleApplicationsFulfillments(applicationId string, query easycb.AnyMap, body easycb.AnyMap) (*SellerSearchSampleApplicationsFulfillmentsRsp, error) {
	var result SellerSearchSampleApplicationsFulfillmentsRsp
	path := fmt.Sprintf("/affiliate_seller/202409/sample_applications/%s/fulfillments/search", applicationId)
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SellerReviewSampleApplications(applicationId string, query easycb.AnyMap, body easycb.AnyMap) (*SellerReviewSampleApplicationsRsp, error) {
	var result SellerReviewSampleApplicationsRsp
	path := fmt.Sprintf("/affiliate_seller/202409/sample_applications/%s/review", applicationId)
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetOpenCollaborationSampleRules(query easycb.AnyMap) (*GetOpenCollaborationSampleRulesRsp, error) {
	var result GetOpenCollaborationSampleRulesRsp
	path := "/affiliate_seller/202410/open_collaborations/sample_rules"
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SellerSearchSampleApplications(query easycb.AnyMap, body easycb.AnyMap) (*SellerSearchSampleApplicationsRsp, error) {
	var result SellerSearchSampleApplicationsRsp
	path := "/affiliate_seller/202409/sample_applications/search"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) EditOpenCollaborationSampleRule(query easycb.AnyMap, body easycb.AnyMap) (*EditOpenCollaborationSampleRuleRsp, error) {
	var result EditOpenCollaborationSampleRuleRsp
	path := "/affiliate_seller/202410/open_collaborations/sample_rules"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) RemoveTargetCollaboration(targetCollaborationId string, query easycb.AnyMap) (*RemoveTargetCollaborationRsp, error) {
	var result RemoveTargetCollaborationRsp
	path := fmt.Sprintf("/affiliate_seller/202409/target_collaborations/%s", targetCollaborationId)
	err := c.doRequest("DELETE", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SearchTargetCollaborations(query easycb.AnyMap, body easycb.AnyMap) (*SearchTargetCollaborationsRsp, error) {
	var result SearchTargetCollaborationsRsp
	path := "/affiliate_seller/202409/target_collaborations/search"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateTargetCollaboration(targetCollaborationId string, query easycb.AnyMap, body easycb.AnyMap) (*UpdateTargetCollaborationRsp, error) {
	var result UpdateTargetCollaborationRsp
	path := fmt.Sprintf("/affiliate_seller/202409/target_collaborations/%s", targetCollaborationId)
	err := c.doRequest("PUT", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetOpenCollaborationSettings(query easycb.AnyMap) (*GetOpenCollaborationSettingsRsp, error) {
	var result GetOpenCollaborationSettingsRsp
	path := "/affiliate_seller/202409/open_collaboration_settings"
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) RemoveOpenCollaboration(productId string, query easycb.AnyMap) (*RemoveOpenCollaborationRsp, error) {
	var result RemoveOpenCollaborationRsp
	path := fmt.Sprintf("/affiliate_seller/202409/open_collaborations/products/%s", productId)
	err := c.doRequest("DELETE", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) QueryTargetCollaborationDetail(targetCollaborationId string, query easycb.AnyMap) (*QueryTargetCollaborationDetailRsp, error) {
	var result QueryTargetCollaborationDetailRsp
	path := fmt.Sprintf("/affiliate_seller/202412/target_collaborations/%s", targetCollaborationId)
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetOpenCollaborationCreatorContentDetail(query easycb.AnyMap) (*GetOpenCollaborationCreatorContentDetailRsp, error) {
	var result GetOpenCollaborationCreatorContentDetailRsp
	path := "/affiliate_seller/202412/open_collaborations/creator_content_details"
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SearchOpenCollaboration(query easycb.AnyMap, body easycb.AnyMap) (*SearchOpenCollaborationRsp, error) {
	var result SearchOpenCollaborationRsp
	path := "/affiliate_seller/202412/open_collaborations/search"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CreateOpenCollaboration(query easycb.AnyMap, body easycb.AnyMap) (*CreateOpenCollaborationRsp, error) {
	var result CreateOpenCollaborationRsp
	path := "/affiliate_seller/202412/open_collaborations"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetMessageInTheConversation(conversationId string, query easycb.AnyMap) (*GetMessageInTheConversationRsp, error) {
	var result GetMessageInTheConversationRsp
	path := fmt.Sprintf("/affiliate_seller/202412/conversation/%s/messages", conversationId)
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return nil, err
}

func (c *Client) GetConversationList(query easycb.AnyMap, body easycb.AnyMap) (*GetConversationListRsp, error) {
	var result GetConversationListRsp
	path := "/affiliate_seller/202412/conversations"
	err := c.doRequest("GET", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return nil, err
}

func (c *Client) SendIMMessage(conversationId string, query easycb.AnyMap, body easycb.AnyMap) (*SendIMMessageRsp, error) {
	var result SendIMMessageRsp
	path := fmt.Sprintf("/affiliate_seller/202412/conversations/%s/messages", conversationId)
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return nil, err
}

func (c *Client) CreateConversationWithCreator(query easycb.AnyMap, body easycb.AnyMap) (*CreateConversationWithCreatorRsp, error) {
	var result CreateConversationWithCreatorRsp
	path := "/affiliate_seller/202412/conversations"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return nil, err
}

func (c *Client) MarkConversationRead(query easycb.AnyMap, body easycb.AnyMap) (*MarkConversationReadRsp, error) {
	var result MarkConversationReadRsp
	path := "/affiliate_seller/202412/conversatons/read"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return nil, err
}

func (c *Client) GetLatestUnreadMessages(query easycb.AnyMap) (*GetLatestUnreadMessagesRsp, error) {
	var result GetLatestUnreadMessagesRsp
	path := "/affiliate_seller/202412/conversations/messages/list/newest"
	err := c.doRequest("GET", path, query, nil, &result)
	if err != nil {
		return nil, err
	}

	return nil, err
}
