package tiktok

import (
	"bytes"
	"github.com/easycb/easycb-go"
	"io/ioutil"
	"net/http"
)

func ParseVerify(req *http.Request) (sign string, body string) {
	sign = req.Header.Get("Authorization")
	if sign == "" {
		return
	}

	bodyBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}

	defer req.Body.Close()

	req.Body = ioutil.NopCloser(bytes.NewReader(bodyBytes))

	body = string(bodyBytes)
	return
}

func (c *Client) VerifySign(sign string, body string) bool {
	if sign == "" || body == "" {
		return false
	}

	input := c.AppKey + body

	signature := easycb.GenerateSHA256([]byte(input), []byte(c.AppSecret))
	if sign == signature {
		return true
	}

	return false
}
