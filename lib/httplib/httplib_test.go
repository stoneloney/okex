// beego httplib的测试用例  直接拿来用
// 测试用例还没有跑，毕竟beego已经经受了诸多用户验证了

package httplib

import (
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

func TestResponse(t *testing.T) {
	req := Get("http://httpbin.org/get")
	resp, err := req.Response()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestGet(t *testing.T) {
	req := Get("http://httpbin.org/get")
	b, err := req.Bytes()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(b)

	s, err := req.String()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s)

	if string(b) != s {
		t.Fatal("request data not match")
	}
}

func TestSimplePost(t *testing.T) {
	v := "smallfish"
	req := Post("http://httpbin.org/post")
	req.Param("username", v)

	str, err := req.String()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(str)

	n := strings.Index(str, v)
	if n == -1 {
		t.Fatal(v + " not found in post")
	}
}

//func TestPostFile(t *testing.T) {
//	v := "smallfish"
//	req := Post("http://httpbin.org/post")
//	req.Debug(true)
//	req.Param("username", v)
//	req.PostFile("uploadfile", "httplib_test.go")

//	str, err := req.String()
//	if err != nil {
//		t.Fatal(err)
//	}
//	t.Log(str)

//	n := strings.Index(str, v)
//	if n == -1 {
//		t.Fatal(v + " not found in post")
//	}
//}

func TestSimplePut(t *testing.T) {
	str, err := Put("http://httpbin.org/put").String()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(str)
}

func TestSimpleDelete(t *testing.T) {
	str, err := Delete("http://httpbin.org/delete").String()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(str)
}

func TestSimpleDeleteParam(t *testing.T) {
	str, err := Delete("http://httpbin.org/delete").Param("key", "val").String()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(str)
}

func TestWithCookie(t *testing.T) {
	v := "smallfish"
	str, err := Get("http://httpbin.org/cookies/set?k1=" + v).SetEnableCookie(true).String()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(str)

	str, err = Get("http://httpbin.org/cookies").SetEnableCookie(true).String()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(str)

	n := strings.Index(str, v)
	if n == -1 {
		t.Fatal(v + " not found in cookie")
	}
}

func TestWithBasicAuth(t *testing.T) {
	str, err := Get("http://httpbin.org/basic-auth/user/passwd").SetBasicAuth("user", "passwd").String()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(str)
	n := strings.Index(str, "authenticated")
	if n == -1 {
		t.Fatal("authenticated not found in response")
	}
}

func TestWithUserAgent(t *testing.T) {
	v := "mlol"
	str, err := Get("http://httpbin.org/headers").SetUserAgent(v).String()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(str)

	n := strings.Index(str, v)
	if n == -1 {
		t.Fatal(v + " not found in user-agent")
	}
}

func TestWithSetting(t *testing.T) {
	v := "beego"
	var setting HTTPSettings
	setting.EnableCookie = true
	setting.UserAgent = v
	setting.Transport = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          50,
		IdleConnTimeout:       90 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	setting.ReadWriteTimeout = 5 * time.Second
	SetDefaultSetting(setting)

	str, err := Get("http://httpbin.org/get").String()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(str)

	n := strings.Index(str, v)
	if n == -1 {
		t.Fatal(v + " not found in user-agent")
	}
}

func TestToJson(t *testing.T) {
	req := Get("http://httpbin.org/ip")
	resp, err := req.Response()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)

	// httpbin will return http remote addr
	type IP struct {
		Origin string `json:"origin"`
	}
	var ip IP
	err = req.ToJSON(&ip)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ip.Origin)
	ips := strings.Split(ip.Origin, ",")
	if len(ips) == 0 {
		t.Fatal("response is not valid ip")
	}
	for i := range ips {
		if net.ParseIP(strings.TrimSpace(ips[i])).To4() == nil {
			t.Fatal("response is not valid ip")
		}
	}

}

func TestToFile(t *testing.T) {
	f := "mlol_testfile"
	req := Get("http://httpbin.org/ip")
	err := req.ToFile(f)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f)
	b, err := ioutil.ReadFile(f)
	if n := strings.Index(string(b), "origin"); n == -1 {
		t.Fatal(err)
	}
}

func TestHeader(t *testing.T) {
	req := Get("http://httpbin.org/headers")
	req.Header("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.57 Safari/537.36")
	str, err := req.String()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(str)
}