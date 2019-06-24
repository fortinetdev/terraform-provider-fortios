package auth

import (
	"fmt"
	"os"
)

// Auth describes the authentication information for FortiOS
type Auth struct {
	Hostname string
	Token    string
	CABundle string
	Vdom     string
	Insecure *bool
	Refresh  bool
}

// NewAuth inits Auth object with the given metadata
func NewAuth(hostname string, token string, cabundle string, vdom string) *Auth {
	return &Auth{
		Hostname: hostname,
		Token:    token,
		CABundle: cabundle,
		Vdom:     vdom,
	}
}

// GetEnvToken gets token from OS environment
// It returns the token
func (m *Auth) GetEnvToken() (string, error) {
	token := os.Getenv("FORTIOS_ACCESS_TOKEN")

	if token == "" {
		return token, fmt.Errorf("GetEnvToken error")
	}

	m.Token = token

	return token, nil
}

// GetEnvHostname gets FortiOS hostname from OS environment
// It returns the hostname
func (m *Auth) GetEnvHostname() (string, error) {
	h := os.Getenv("FORTIOS_ACCESS_HOSTNAME")

	if h == "" {
		return h, fmt.Errorf("GetEnvHostname error")
	}

	m.Hostname = h

	return h, nil
}

// GetEnvCABundle gets CA Bundle file from OS environment
// It returns the CA Bundle file
func (m *Auth) GetEnvCABundle() (string, error) {
	c := os.Getenv("FORTIOS_CA_CABUNDLE")

	if c == "" {
		return c, nil
	}

	m.CABundle = c

	return c, nil
}

// GetEnvInsecure gets Insecure value from OS environment
// It returns the insecure value
func (m *Auth) GetEnvInsecure() (bool, error) {
	c := os.Getenv("FORTIOS_INSECURE")

	if c == "true" {
		return true, nil
	}

	return false, nil
}
