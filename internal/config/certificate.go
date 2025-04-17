package config

type Certificate struct {
	// 是否开启
	Enabled *bool `default:"true" json:"enabled,omitempty"`
	// 是否跳过证书检查
	Skip *bool `default:"true" json:"skip,omitempty"`
	// 根秘钥文件路径
	Root string `json:"root,omitempty" validate:"omitempty,file"`
	// 客户端
	Clients []ClientCertificate `json:"clients,omitempty" validate:"structonly"`
}

func (c *Certificate) Enable() bool {
	return nil == c.Enabled || *c.Enabled
}
