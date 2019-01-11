package auth

import (
	"fmt"
	"os"
)

type Auth struct {
	Token   string
	Refresh bool
}

// Create auth
func NewAuth(token string) *Auth {
	return &Auth{
		Token: token,
	}
}

// Get auth information from environment
func (m *Auth) GetEnvToken() (string, error) {
	token := os.Getenv("FORTIOS_ACCESS_TOKEN")

	if token == "" {
		return token, fmt.Errorf("GetEnvToken error")
	}

	m.Token = token

	return token, nil
}
