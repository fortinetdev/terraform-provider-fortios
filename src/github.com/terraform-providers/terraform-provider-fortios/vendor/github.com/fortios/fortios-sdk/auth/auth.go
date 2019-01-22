package auth

import (
	"fmt"
	"os"
)

// Auth describes the authentication information for FortiOS
type Auth struct {
	Token   string
	Refresh bool
}

// NewAuth inits Auth object with the given metadata
func NewAuth(token string) *Auth {
	return &Auth{
		Token: token,
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
