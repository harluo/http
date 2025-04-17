package config

import (
	"net/http"
	"time"
)

type Client struct {
	// 超时
	Timeout time.Duration `json:"timeout,omitempty"`
	// 代理
	Proxy *Proxy `json:"proxy,omitempty"`
	// 代理列表
	Proxies []*Proxy `default:"[]" json:"proxies,omitempty"`
	// 授权配置
	Auth *Auth `json:"auth,omitempty"`
	// Body数据传输控制
	Payload *bool `default:"true" json:"payload,omitempty"`
	// 警告消息
	Warning bool `json:"warning,omitempty"`
	// 秘钥配置
	Certificate Certificate `json:"certificate,omitempty"`
	// 通用的查询参数
	Queries map[string]string `json:"queries,omitempty"`
	// 表单参数
	Forms map[string]string `json:"forms,omitempty"`
	// 通用头信息
	Headers map[string]string `json:"headers,omitempty"`
	// 通用Cookie
	Cookies []*http.Cookie `json:"cookies,omitempty"`
}

func newClient(http *Http) *Client {
	return http.Client
}
