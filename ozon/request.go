package ozon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/easycb/easycb-go"
	"net/http"
	"net/url"
)

func (c *Client) doRequest(method string, apiPath string, query easycb.AnyMap, body easycb.AnyMap, result interface{}) error {
	var err error
	if c.BaseURL == "" {
		return easycb.MissBaseUrlErr
	}

	// Build Query With System Params
	values, err := c.URLValues(query)
	if err != nil {
		return err
	}

	// Build Body
	bodyByte := []byte{}
	if body != nil {
		bodyByte, err = json.Marshal(body)
		if err != nil {
			return err
		}
	}

	// New Request
	fullUrl := fmt.Sprintf("%s%s?%s", c.BaseURL, apiPath, values.Encode())
	req, err := http.NewRequest(method, fullUrl, bytes.NewBuffer(bodyByte))
	if err != nil {
		return err
	}

	// Build Header
	req.Header.Add("Content-Type", ContentTypeJson)
	req.Header.Add("Accept", AcceptJson)

	// Build
	c.makeHeaders(req)

	// Do Response
	resp, err := c.conf.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// Parse Response
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&result)
	return err
}

func (c *Client) URLValues(query easycb.AnyMap) (value url.Values, err error) {
	var values = url.Values{}
	for k, v := range query {
		values.Set(k, easycb.InterfaceToString(v))
	}

	return values, nil
}

func (c *Client) makeHeaders(r *http.Request) {
	r.Header.Add("Client-Id", c.ClientId)
	r.Header.Add("Api-Key", c.ApiKey)
}
