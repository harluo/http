package config

import (
	"github.com/harluo/boot"
)

type Http struct {
	Client *Client `json:"client,omitempty" validate:"required"`
}

func newHttp(config *boot.Config) (http *Http, err error) {
	http = new(Http)
	err = config.Build().Get(&struct {
		Http *Http `json:"http,omitempty" validate:"required"`
	}{
		Http: http,
	})

	return
}
