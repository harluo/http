package config

import (
	"time"
)

type Timeout struct {
	Connection time.Duration `json:"connection,omitempty"`
	Response   time.Duration `json:"response,omitempty"`
	Idle       time.Duration `json:"idle,omitempty"`
	Handshake  time.Duration `json:"handshake,omitempty"`
}
