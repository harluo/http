package core

import (
	"github.com/goexl/http"
	"github.com/harluo/http/internal/config"
)

type Client = http.Client

func newClient(config *config.Client) *http.Client {
	builder := http.New().Payload(*config.Payload).Timeout(config.Timeout).
		Headers(config.Headers).
		Forms(config.Forms).
		Queries(config.Queries).
		Warning(config.Warning)

	if nil != config.Proxy && config.Proxy.Checked() {
		config.Proxies = append(config.Proxies, config.Proxy)
	}
	proxy := builder.Proxy()
	for _, server := range config.Proxies {
		if !server.Checked() {
			continue
		}

		proxy.Host(server.Host).Port(server.Port).Target(server.Target).Exclude(server.Exclude).
			Scheme(server.Scheme).
			Basic(server.Username, server.Password).
			Build()
	}

	if nil != config.Auth && config.Auth.Enable() {
		conf := config.Auth
		auth := builder.Auth()
		_ = auth.Scheme(conf.Scheme).Token(conf.Token).Basic(conf.Username, conf.Password).Build()
	}

	if config.Certificate.Enable() {
		conf := config.Certificate
		certificate := builder.Certificate()
		certificate = certificate.Skip(*conf.Skip).Root(conf.Root)
		for _, client := range conf.Clients {
			certificate = certificate.Client(client.Public, client.Private)
		}
		_ = certificate.Build()
	}

	return builder.Build()
}
