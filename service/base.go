package service

import (
	"fmt"
	"io/ioutil"
	"okex/define"
	"okex/helper"
	"okex/lib/httplib"
	"strconv"
	"time"
)

type Okex struct {

}

func (ok *Okex) SendPostReq() {

}

func (ok *Okex) SendGetReq(uri string) ([]byte, error) {
	timestamp := strconv.FormatInt(time.Now().Unix(),10)
	sign, err := ok.getSign(timestamp, "GET", uri, "")
	if err != nil {
		return nil, err
	}

	client := httplib.Get(ok.getApiUrl(uri))
	client.Header("OK-ACCESS-TIMESTAMP", strconv.FormatInt(time.Now().Unix(),10))
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
