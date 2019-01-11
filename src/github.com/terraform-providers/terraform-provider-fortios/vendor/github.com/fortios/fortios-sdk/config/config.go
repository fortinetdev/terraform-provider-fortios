package config

import (
	"net/http"

	"github.com/fortios/fortios-sdk/auth"
)

type Config struct {
	Auth     *auth.Auth
	HTTPCon  *http.Client
	FwTarget string
}

// func New(host string, auth *auth.Auth, client *http.Client) *Config {
// 	config := &Config{
// 		Auth: auth,
// 		HTTPCon: client,
// 		FwTarget: host,
// 	}
// 	return config
// }
