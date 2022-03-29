package httplib

import (
	"net"
	"net/http"
	"time"
)

const (
	// MaxIdleConns 表示连接池对所有host的最大链接数量，host也即dest-ip，默认为无穷大（0），
	// 但是通常情况下为了性能考虑都要严格限制该数目（实际使用中通常利用压测 二分得到该参数的最佳近似值）。
	// 太大容易导致客户端和服务端的socket数量剧增，导致内存吃满，文件描述符不足等问题；
	// 太小则限制了连接池的socket数量，资源利用率较低
	MaxIdleConns int = 100

	// MaxIdleConnsPerHost 表示连接池对每个host的最大链接数量，从字面意思也可以看出：
	// MaxIdleConnsPerHost <= MaxIdleConns
	// 如果客户端只需要访问一个host，那么最好将MaxIdleConnsPerHost与MaxIdleConns设置为相同，这样逻辑更加清晰
	MaxIdleConnsPerHost int = 100

	// IdleConnTimeout 空闲timeout设置，也即socket在该时间内没有交互则自动关闭连接（注意：该timeout起点是从每次空闲开始计时，若有交互则重置为0）,
	// 该参数通常设置为分钟级别，例如：90秒。
	IdleConnTimeout int = 90
)

// CreateHTTPClientPool for connection re-use
func CreateHTTPClientPool(opts ...int) *http.Client {

	maxIdleConns := MaxIdleConns
	maxIdleConnsPerHost := MaxIdleConnsPerHost
	idleConnTimeout := IdleConnTimeout

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:        maxIdleConns,
			MaxIdleConnsPerHost: maxIdleConnsPerHost,
			IdleConnTimeout:     time.Duration(idleConnTimeout) * time.Second,
		},
		Timeout: 20 * time.Second,
	}
	return client
}
