package config

type Auth struct {
	// 是否开启
	Enabled *bool `default:"true" json:"enabled,omitempty"`
	// 用户名
	Username string `json:"username,omitempty"`
	// 密码
	Password string `json:"password,omitempty"`
	// 授权码
	Token string `json:"token,omitempty" validate:"required_if=Type token"`
	// 身份验证方案类型
	Scheme string `json:"scheme,omitempty"`
}

func (a *Auth) Enable() bool {
	return nil == a.Enabled || *a.Enabled
}
