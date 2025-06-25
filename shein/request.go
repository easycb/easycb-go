package shein

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/easycb/easycb-go"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func (c *Client) doRequest(method string, apiPath string, query easycb.AnyMap, body easycb.AnyMap, headers map[string]string, result interface{}) error {
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

	for k, v := range headers {
		req.Header.Add(k, v)
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

func (c *Client) doFileRequest(method string, apiPath string, query easycb.AnyMap, body easycb.AnyMap, headers map[string]string, result interface{}) error {
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

	for k, v := range headers {
		req.Header.Add(k, v)
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
	for k, v := range query {
		values.Set(k, easycb.InterfaceToString(v))
	}

	return values, nil
}

func (c *Client) makeSignature(r *http.Request) {
	path := r.URL.Path
	if path == "/open-api/auth/get-by-token" {
		c.buildAppSignature(r)
	} else {
		c.buildSignature(r)
	}
}

func (c *Client) buildSignature(r *http.Request) {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	signStr := c.OpenKeyId + "&" + timestamp + "&" + r.URL.Path
	randomKey := easycb.GenerateRandStr(5, "")
	hashValue := easycb.GenerateSHA256([]byte(signStr), []byte(c.SecretKey+randomKey))
	base64Value := base64.StdEncoding.EncodeToString([]byte(hashValue))
	signature := randomKey + base64Value

	r.Header.Add("x-lt-appid", c.AppId)
	r.Header.Add("x-lt-openKeyId", c.OpenKeyId)
	r.Header.Add("x-lt-timestamp", timestamp)
	r.Header.Add("x-lt-signature", signature)
}

func (c *Client) buildAppSignature(r *http.Request) {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	signStr := c.AppId + "&" + timestamp + "&" + r.URL.Path
	randomKey := easycb.GenerateRandStr(5, "")
	hashValue := easycb.GenerateSHA256([]byte(signStr), []byte(c.AppSecretKey+randomKey))
	base64Value := base64.StdEncoding.EncodeToString([]byte(hashValue))
	signature := randomKey + base64Value

	r.Header.Add("x-lt-appid", c.AppId)
	r.Header.Add("x-lt-timestamp", timestamp)
	r.Header.Add("x-lt-signature", signature)
}
