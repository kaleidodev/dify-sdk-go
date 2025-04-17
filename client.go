package dify

import (
	"net/http"
	"time"

	"github.com/safejob/dify-sdk-go/base"
)

type ClientConfig struct {
	ApiServer string        // [必填]API服务器 eg: http://your.domain.com/v1 注意是包括/v1的
	ApiKey    string        // [必填]API密钥
	User      string        // 用户标识 部分接口需要传入用户标识,这里设置后,调用接口时user字段可留空
	Debug     bool          // 是否打印原始请求及响应
	Timeout   time.Duration // 超时时间,默认300秒
	Transport *http.Transport
}

const defaultTimeout = 300 * time.Second

func NewClient(config ClientConfig) (*base.Client, error) {
	var httpClient = &http.Client{}

	if config.Timeout <= 0 {
		config.Timeout = defaultTimeout
	}
	httpClient.Timeout = config.Timeout

	if config.Transport != nil {
		httpClient.Transport = config.Transport
	}

	return base.NewClient(config.ApiServer, config.ApiKey, config.User, config.Debug, config.Timeout, httpClient)
}
