package fortios

import (
	"crypto/tls"
	"net/http"

	"github.com/fortios/fortios-sdk/auth"
	"github.com/fortios/fortios-sdk/sdkcore"
)

// Config Basic connection information
type Config struct {
	Hostname string
	Token    string
}

// FortiClient Used for every tf modules
type FortiClient struct {
	//to sdk client
	Client *forticlient.FortiSDKClient
}

// CreateClient ...
func (c *Config) CreateClient() (interface{}, error) {
	var fClient FortiClient

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: tr,
	}

	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }
	// c.HTTPCon.Transport = tr

	auth := auth.NewAuth(c.Token)

	if auth.Token == "" {
		auth.GetEnvToken()
	}

	fc := forticlient.NewClient(c.Hostname, auth, client)

	fClient.Client = fc

	return &fClient, nil
}
