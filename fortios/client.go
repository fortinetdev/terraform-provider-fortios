package fortios

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/fgtdev/fortios-sdk-go/auth"
	"github.com/fgtdev/fortios-sdk-go/sdkcore"
)

// Config gets the authentication information from the given metadata
type Config struct {
	Hostname string
	Token    string
	Vdom     string
}

// FortiClient contains the basic FortiOS SDK connection information to FortiOS
// It can be used to as a client of FortiOS for the plugin
type FortiClient struct {
	//to sdk client
	Client *forticlient.FortiSDKClient
}

// CreateClient creates a FortiClient Object with the authentication information.
// It returns the FortiClient Object for the use when the plugin is initialized.
func (c *Config) CreateClient() (interface{}, error) {
	var fClient FortiClient

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 10,
	}

	auth := auth.NewAuth(c.Hostname, c.Token, c.Vdom)

	if auth.Hostname == "" {
		auth.GetEnvHostname()
	}

	if auth.Token == "" {
		auth.GetEnvToken()
	}

	fc := forticlient.NewClient(auth, client)

	fClient.Client = fc

	return &fClient, nil
}
