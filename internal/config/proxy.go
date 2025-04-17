package config

type Proxy struct {
	// 是否开启
	Enabled *bool `default:"true" json:"enabled,omitempty"`
	// 主机
	Host string `json:"host,omitempty" validate:"required,hostname,ip"`
	// 端口
	Port int `json:"port,omitempty" validate:"omitempty,min=1,max=65535"`
	// 代理类型
	// nolint: lll
	Scheme string `default:"http" json:"scheme,omitempty" validate:"required,oneof=socks4 socks5 http https"`
	// 目标
	Target string `json:"target,omitempty"`
	// 排除
	Exclude string `json:"exclude,omitempty"`
	// 代理认证用户名
	Username string `json:"username,omitempty"`
	// 代理认证密码
	Password string `json:"password,omitempty"`
}

func (p *Proxy) Checked() bool {
	return nil != p.Enabled && *p.Enabled
}
