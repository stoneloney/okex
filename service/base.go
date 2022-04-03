package service

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"okex/define"
	"okex/helper"
	"okex/lib/httplib"
	"time"
)

type Okex struct {

}

func (ok *Okex) SendPostReq() {

}

func (ok *Okex) SendGetReq(apiuri string, params map[string]string) ([]byte, error) {
	// 构造请求参数
	if len(params) > 0 {
		var uri url.URL
		query := uri.Query()
		for k, v := range params {
			query.Add(k, v)
		}
		queryStr := query.Encode()
		apiuri += "?" + queryStr
	}

	timestamp := helper.IsoTime()
	sign, err := ok.getSign(timestamp, "GET", apiuri, "")
	if err != nil {
		return nil, err
	}

	fmt.Println(apiuri)
	apiUrl := ok.getApiUrl(apiuri)
	fmt.Println(apiUrl)

	client := httplib.Get(apiUrl)
	client.Header("OK-ACCESS-TIMESTAMP", timestamp)
	client.Header("OK-ACCESS-KEY", define.ApiKey)
	client.Header("OK-ACCESS-PASSPHRASE", define.Passphrase)
	client.Header("OK-ACCESS-SIGN", sign)
	client.SetTimeout(2*time.Second, 2*time.Second)

	resp, err := client.Response()
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (ok *Okex) getApiUrl(uri string) string {
	return define.ApiHttpHost + uri
}

func (ok *Okex) getSign(timestamp, method, uri, body string) (string, error){
	msg := fmt.Sprintf("%s%s%s%s", timestamp, method, uri, body)
	sign, err := helper.HmacSha256Base64Signer(msg, define.SecretKey)
	if err != nil {
		return "", err
	}
	return sign, nil
}
