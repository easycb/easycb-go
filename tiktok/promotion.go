package tiktok

import (
	"fmt"
	"github.com/easycb/easycb-go"
)

func (c *Client) CreateActivity(body easycb.AnyMap) (*CreateActivityRsp, error) {
	var result CreateActivityRsp
	path := "/promotion/202309/activities"
	err := c.doRequest("POST", path, nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateActivity(activityId string, body easycb.AnyMap) (*UpdateActivityRsp, error) {
	var result UpdateActivityRsp
	path := fmt.Sprintf("/promotion/202309/activities/%s", activityId)
	err := c.doRequest("PUT", path, nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeactivateActivity(activityId string) (*DeactivateActivityRsp, error) {
	var result DeactivateActivityRsp
	path := fmt.Sprintf("/promotion/202309/activities/%s/deactivate", activityId)
	err := c.doRequest("POST", path, nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetActivity(activityId string) (*GetActivityRsp, error) {
	var result GetActivityRsp
	path := fmt.Sprintf("/promotion/202309/activities/%s", activityId)
	err := c.doRequest("GET", path, nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SearchActivities(body easycb.AnyMap) (*SearchActivitiesRsp, error) {
	var result SearchActivitiesRsp
	path := "/promotion/202309/activities/search"
	err := c.doRequest("POST", path, nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateActivityProduct(activityId string, body easycb.AnyMap) (*UpdateActivityProductRsp, error) {
	var result UpdateActivityProductRsp
	path := fmt.Sprintf("/promotion/202309/activities/%s/products", activityId)
	err := c.doRequest("PUT", path, nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) RemoveActivityProduct(activityId string, body easycb.AnyMap) (*RemoveActivityProductRsp, error) {
	var result RemoveActivityProductRsp
	path := fmt.Sprintf("/promotion/202309/activities/%s/products", activityId)
	err := c.doRequest("DEL", path, nil, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetCoupon(couponId string) (*GetCouponRsp, error) {
	var result GetCouponRsp
	path := fmt.Sprintf("/promotion/202406/coupons/%s", couponId)
	err := c.doRequest("GET", path, nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SearchCoupons(query easycb.AnyMap, body easycb.AnyMap) (*SearchCouponsRsp, error) {
	var result SearchCouponsRsp
	path := "/promotion/202406/coupons/search"
	err := c.doRequest("POST", path, query, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
