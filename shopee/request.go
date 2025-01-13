package shopee

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/easycb/easycb-go"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func (c *Client) doRequest(method string, apiPath string, query easycb.AnyMap, body easycb.AnyMap, result interface{}) error {
	var err error
	if c.BaseURL == "" {
		return easycb.MissBaseUrlErr
	}

	// Build Query With System Params
	values, err := c.URLValues(query, apiPath)
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
	c.makeSignature(req)

	// Do Response
	resp, err := c.conf.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// Parser Raw Bytes
	rspContentType := resp.Header.Get("Content-Type")
	if rspContentType == "application/pdf" || rspContentType == "text/html" || strings.HasPrefix(rspContentType, "image/") {
		val := reflect.ValueOf(result)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
			fmt.Println("0000")
			if val.Kind() == reflect.Struct {
				for i := 0; i < val.NumField(); i++ {
					field := val.Type().Field(i)
					if field.Name == "RawBytes" {
						if val.Field(i).CanSet() {
							rspBody, err := ioutil.ReadAll(resp.Body)
							if err != nil {
								return err
							}
							val.Field(i).SetBytes(rspBody)
						}
						break
					}
				}
			}
		}

		return nil
	}

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
	values, err := c.URLValues(query, apiPath)
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

func (c *Client) URLValues(query easycb.AnyMap, apiPath string) (value url.Values, err error) {
	var values = url.Values{}
	values.Add("partner_id", strconv.FormatInt(c.PartnerId, 10))
	values.Add("timestamp", strconv.FormatInt(time.Now().Unix(), 10))

	if c.ShopID > 0 {
		values.Add("shop_id", fmt.Sprintf("%v", c.ShopID))
	}

	if c.MerchantID > 0 {
		values.Add("merchant_id", fmt.Sprintf("%v", c.MerchantID))
	}

	if c.AccessToken != "" {
		values.Add("access_token", c.AccessToken)
	}

	// If you want to search multiple status, please upload the url like this: item_status=NORMAL&item_status=BANNED
	if apiPath == "/api/v2/product/get_item_list" {
		if itemStatuses, ok := query["item_status"].([]string); ok {
			delete(query, "item_status")
			for _, status := range itemStatuses {
				values.Add("item_status", status)
			}
		}
	}

	if apiPath == "/api/v2/product/get_item_base_info" || apiPath == "/api/v2/product/get_item_extra_info" {
		if itemIdList, ok := query["item_id_list"].([]int64); ok {
			delete(query, "item_id_list")
			for _, itemId := range itemIdList {
				values.Add("item_id_list", strconv.FormatInt(itemId, 10))
			}
		}
	}

	if apiPath == "/api/v2/product/get_attribute_tree" {
		if itemIdList, ok := query["category_id_list"].([]int64); ok {
			delete(query, "category_id_list")
			for _, itemId := range itemIdList {
				values.Add("category_id_list", strconv.FormatInt(itemId, 10))
			}
		}
	}

	for k, v := range query {
		values.Set(k, easycb.InterfaceToString(v))
	}

	return values, nil
}

func (c *Client) makeSignature(req *http.Request) {
	queries := req.URL.Query()
	timestamp := queries.Get("timestamp")

	path := req.URL.Path
	input := ""
	if c.ShopID > 0 {
		input = fmt.Sprintf("%d%s%s%s%d", c.PartnerId, path, timestamp, c.AccessToken, c.ShopID)
	} else if c.MerchantID > 0 {
		input = fmt.Sprintf("%d%s%s%s%d", c.PartnerId, path, timestamp, c.AccessToken, c.MerchantID)
	} else {
		input = fmt.Sprintf("%d%s%s", c.PartnerId, path, timestamp)
	}

	sign := easycb.GenerateSHA256([]byte(input), []byte(c.PartnerKey))

	queries.Add("sign", strings.ToUpper(sign))

	u := req.URL
	u.RawQuery = queries.Encode()
	req.URL = u
}
