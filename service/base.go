package service

import (
	"encoding/json"
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

func (ok *Okex) SendPostReq(apiUri string, reqData interface{}) ([]byte, error) {
	timestamp := helper.IsoTime()
	bodyJson, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	sign, err := ok.getSign(timestamp, "POST", apiUri, string(bodyJson))
	if err != nil {
		return nil, err
	}

	fmt.Println(apiUri)
	apiUrl := ok.getApiUrl(apiUri)
	fmt.Println(apiUrl)

	client := httplib.Post(apiUrl)
	client.SetTimeout(2*time.Second, 2*time.Second)

	client.Header("OK-ACCESS-TIMESTAMP", timestamp)
	client.Header("OK-ACCESS-KEY", define.ApiKey)
	client.Header("OK-ACCESS-PASSPHRASE", define.Passphrase)
	client.Header("OK-ACCESS-SIGN", sign)
	client.SetTimeout(2*time.Second, 2*time.Second)

	_, err = client.JSONBody(reqData)
	if err != nil {
		return nil, err
	}

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

func (ok *Okex) SendGetReq(apiUri string, params map[string]string) ([]byte, error) {
	// 构造请求参数
	if len(params) > 0 {
		var uri url.URL
		query := uri.Query()
		for k, v := range params {
			query.Add(k, v)
		}
		queryStr := query.Encode()
		apiUri += "?" + queryStr
	}

	timestamp := helper.IsoTime()
	sign, err := ok.getSign(timestamp, "GET", apiUri, "")
	if err != nil {
		return nil, err
	}

	fmt.Println(apiUri)
	apiUrl := ok.getApiUrl(apiUri)
	fmt.Println(apiUrl)

	client := httplib.Get(apiUrl)
	client.SetTimeout(2*time.Second, 2*time.Second)

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

	fmt.Println(string(body))

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
