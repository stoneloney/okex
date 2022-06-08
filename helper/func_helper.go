package helper

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"math/big"
	"time"
)

func HmacSha256Base64Signer(message string, secretKey string) (string, error) {
	mac := hmac.New(sha256.New, []byte(secretKey))
	_, err := mac.Write([]byte(message))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(mac.Sum(nil)), nil
}

/*
 Get a iso time
  eg: 2018-03-16T18:02:48.284Z
*/
func IsoTime() string {
	utcTime := time.Now().UTC()
	iso := utcTime.String()
	isoBytes := []byte(iso)
	iso = string(isoBytes[:10]) + "T" + string(isoBytes[11:23]) + "Z"
	return iso
}

/**
 * 当前时间
 */
func TimeNowStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func InArray(val string, arr []string) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

/**
 * float转string
 */
func Float64ToString(num float64) string {
	newNum := big.NewRat(1, 1)
	newNum.SetFloat64(num)
	return newNum.FloatString(0)
}

