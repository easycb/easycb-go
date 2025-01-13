package lazada

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/easycb/easycb-go"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

func (c *Client) doRequest(method string, apiPath string, query easycb.AnyMap, body easycb.AnyMap, result interface{}) error {
	var err error
	// Get ServerURL By Region Or Base Url
	baseUrl := c.BaseURL
	if baseUrl == "" && c.Region != "" {
		c.BaseURL = c.getServerURL()
	}

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

	// Build Sign
	c.makeSignature(req, body)

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

func (c *Client) doFileRequest(method string, apiPath string, query easycb.AnyMap, body easycb.AnyMap, result interface{}) error {
	var err error
	// Get ServerURL By Region Or Base Url
	baseUrl := c.BaseURL
	if baseUrl == "" && c.Region != "" {
		c.BaseURL = c.getServerURL()
	}

	if c.BaseURL == "" {
		return easycb.MissBaseUrlErr
	}

	// Build Query With System Params
	values, err := c.URLValues(query)
	if err != nil {
		return err
	}

	// Build Form File
	bodyBuffer := &bytes.Buffer{}
	writer := multipart.NewWriter(bodyBuffer)
	for k, v := range body {
		// Check *os.File
		if file, ok := v.(*os.File); ok {
			filePath := file.Name()
			fileName := filepath.Base(filePath)
			part, _ := writer.CreateFormFile(k, fileName)
			if _, err := io.Copy(part, file); err != nil {
				return err
			}
		} else {
			err = writer.WriteField(k, easycb.InterfaceToString(v))
		}
	}

	if err = writer.Close(); err != nil {
		return err
	}

	// New Request
	fullUrl := fmt.Sprintf("%s%s?%s", c.BaseURL, apiPath, values.Encode())
	req, err := http.NewRequest(method, fullUrl, bodyBuffer)
	if err != nil {
		return err
	}

	// Build Header
	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.Header.Add("Accept", AcceptJson)

	// Build Sign
	c.makeSignature(req, body)

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
	sysParams := c.getSysParams()
	for k, v := range sysParams {
		values.Add(k, easycb.InterfaceToString(v))
	}

	if c.AccessToken != "" {
		values.Add("access_token", c.AccessToken)
	}

	for k, v := range query {
		values.Set(k, easycb.InterfaceToString(v))
	}

	return values, nil
}

func (c *Client) makeSignature(req *http.Request, body easycb.AnyMap) {
	queries := req.URL.Query()

	// extract all query parameters excluding sign
	keys := []string{}
	union := easycb.AnyMap{}
	for k := range queries {
		if k != "sign" {
			union[k] = queries.Get(k)
			keys = append(keys, k)
		}
	}

	mediaType, _, _ := mime.ParseMediaType(req.Header.Get("Content-type"))
	if mediaType != "multipart/form-data" {
		for k, v := range body {
			if k != "sign" {
				union[k] = v
				keys = append(keys, k)
			}
		}
	}

	// sort sys params and api params by key
	sort.Strings(keys)

	// Concatenate all the parameters in the format of {key}{value}
	var message bytes.Buffer
	message.WriteString(fmt.Sprintf("%s", strings.Replace(req.URL.Path, "/rest", "", 1)))
	for _, key := range keys {
		message.WriteString(fmt.Sprintf("%s%s", key, union[key]))
	}

	sign := easycb.GenerateSHA256(message.Bytes(), []byte(c.AppSecret))

	queries.Add("sign", strings.ToUpper(sign))

	u := req.URL
	u.RawQuery = queries.Encode()
	req.URL = u
}

func (c *Client) getSysParams() easycb.AnyMap {
	params := easycb.AnyMap{
		"app_key":     c.AppKey,
		"timestamp":   fmt.Sprintf("%d000", time.Now().Unix()),
		"sign_method": SignMethod,
		"partner_id":  Version,
	}

	return params
}

func (c *Client) getServerURL() string {
	switch c.Region {
	case "SG":
		return APIGatewaySG
	case "MY":
		return APIGatewayMY
	case "VN":
		return APIGatewayVN
	case "TH":
		return APIGatewayTH
	case "PH":
		return APIGatewayPH
	case "ID":
		return APIGatewayID
	case "AUTH":
		return ApiGatewayAuth
	case "ALL":
		return ApiGatewayAll
	}
	return ""
}
