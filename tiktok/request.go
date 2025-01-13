package tiktok

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
	"strconv"
	"time"
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
	if c.AccessToken != "" {
		req.Header.Add(AccessTokenName, c.AccessToken)
	}

	// Build Sign
	c.makeSignature(req)

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
	if c.AccessToken != "" {
		req.Header.Add(AccessTokenName, c.AccessToken)
	}

	// Build Sign
	c.makeSignature(req)

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
	values.Add("app_key", c.AppKey)
	values.Add("timestamp", strconv.FormatInt(time.Now().Unix(), 10))
	if c.ShopCipher != "" {
		values.Add("shop_cipher", c.ShopCipher)
	}

	if c.AccessToken != "" {
		values.Add("access_token", c.AccessToken)
	}

	for k, v := range query {
		values.Set(k, easycb.InterfaceToString(v))
	}

	return values, nil
}

func (c *Client) makeSignature(req *http.Request) {
	queries := req.URL.Query()

	// extract all query parameters excluding sign and access_token
	keys := make([]string, len(queries))
	idx := 0
	for k := range queries {
		// params except 'sign' & 'access_token'
		if k != "sign" && k != "access_token" {
			keys[idx] = k
			idx++
		}
	}

	// reorder the parameters' key in alphabetical order
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	// Concatenate all the parameters in the format of {key}{value}
	input := ""
	for _, key := range keys {
		input = input + key + queries.Get(key)
	}

	// append the request path
	input = req.URL.Path + input

	// if the request header Content-type is not multipart/form-data, append body to the end
	mediaType, _, _ := mime.ParseMediaType(req.Header.Get("Content-type"))
	if mediaType != "multipart/form-data" {
		body, _ := io.ReadAll(req.Body)
		input = input + string(body)

		req.Body.Close()
		// reset body after reading from the original
		req.Body = io.NopCloser(bytes.NewReader(body))
	}

	// wrap the string generated in step 5 with the App secret
	input = c.AppSecret + input + c.AppSecret
	sign := easycb.GenerateSHA256([]byte(input), []byte(c.AppSecret))

	queries.Add("sign", sign)

	u := req.URL
	u.RawQuery = queries.Encode()
	req.URL = u
}
