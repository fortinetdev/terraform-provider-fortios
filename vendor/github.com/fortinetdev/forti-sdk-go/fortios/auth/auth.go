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
	CABundleContent string
	Vdom     string
	Insecure *bool
	Refresh  bool
	HTTPProxy string

	PeerAuth string
	CaCert string
	ClientCert string
	ClientKey string
}

// NewAuth inits Auth object with the given metadata
func NewAuth(hostname, token, cabundle, cabundlecontent, peerauth, cacert, clientcert, clientkey, vdom, httpproxy string) *Auth {
	return &Auth{
		Hostname: hostname,
		Token:    token,
		CABundle: cabundle,
		CABundleContent: cabundlecontent,
		Vdom:     vdom,
		HTTPProxy: httpproxy,

		PeerAuth:   peerauth,
		CaCert:     cacert,
		ClientCert: clientcert,
		ClientKey:  clientkey,
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

//////////////////////////////////////////////////////////////////////
// GetEnvPeerAuth gets PeerAuth from OS environment
// It returns the PeerAuth value
func (m *Auth) GetEnvPeerAuth() (string, error) {
	c := os.Getenv("FORTIOS_CA_PEERAUTH")

	if c == "" {
		return c, nil
	}

	m.PeerAuth = c

	return c, nil
}

// GetEnvCaCert gets Peer Auth Ca Cert file from OS environment
// It returns the Ca Cert file file
func (m *Auth) GetEnvCaCert() (string, error) {
	c := os.Getenv("FORTIOS_CA_CACERT")

	if c == "" {
		return c, nil
	}

	m.CaCert = c

	return c, nil
}

// GetEnvClientCert gets Peer Auth User Cert file from OS environment
// It returns the User Cert file
func (m *Auth) GetEnvClientCert() (string, error) {
	c := os.Getenv("FORTIOS_CA_CLIENTCERT")

	if c == "" {
		return c, nil
	}

	m.ClientCert = c

	return c, nil
}

// GetEnvClientKey gets Peer Auth User Key file from OS environment
// It returns the User Key file
func (m *Auth) GetEnvClientKey() (string, error) {
	c := os.Getenv("FORTIOS_CA_CLIENTKEY")

	if c == "" {
		return c, nil
	}

	m.ClientKey = c

	return c, nil
}

// GetEnvHTTPProxy gets HTTP_PROXY or HTTPS_PROXY from OS environment
// It returns the HTTP_PROXY or HTTPS_PROXY
func (m *Auth) GetEnvHTTPProxy() (string, error) {
	c := os.Getenv("HTTPS_PROXY")

	if c == "" {
		c = os.Getenv("HTTP_PROXY")
	}

	if c != "" {
		m.HTTPProxy = c
	}

	return c, nil
}

