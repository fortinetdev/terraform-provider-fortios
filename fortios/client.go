package fortios

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	fmgclient "github.com/fgtdev/fortimanager-sdk-go/sdkcore"
	"github.com/fgtdev/fortios-sdk-go/auth"
	forticlient "github.com/fgtdev/fortios-sdk-go/sdkcore"
)

// Config gets the authentication information from the given metadata
type Config struct {
	Hostname string
	Token    string
	Insecure *bool
	CABundle string
	Vdom     string
	Username string
	Passwd   string
	Product  string
}

// FortiClient contains the basic FortiOS SDK connection information to FortiOS
// It can be used to as a client of FortiOS for the plugin
// Now FortiClient contains two kinds of clients:
// Client is for FortiGate
// Client Fottimanager is for FortiManager
type FortiClient struct {
	//to sdk client
	Client             *forticlient.FortiSDKClient
	ClientFortimanager *fmgclient.FmgSDKClient
}

// CreateClient creates a FortiClient Object with the authentication information.
// It returns the FortiClient Object for the use when the plugin is initialized.
func (c *Config) CreateClient() (interface{}, error) {
	var fClient FortiClient

	if c.Product == "" {
		c.Product = os.Getenv("FORTIOS_PRODUCT")
	}

	switch c.Product {
	case "fortimanager":
		{
			if c.Hostname == "" {
				c.Hostname = os.Getenv("FORTIOS_FMG_HOSTNAME")
			}
			if c.Username == "" {
				c.Username = os.Getenv("FORTIOS_FMG_USERNAME")
			}
			if c.Passwd == "" {
				c.Passwd = os.Getenv("FORTIOS_FMG_PASSWORD")
			}
			if c.Hostname != "" && c.Username != "" && c.Passwd != "" {
				fClient.ClientFortimanager = fmgclient.NewClient(c.Hostname, c.Username, c.Passwd)
				return &fClient, nil
			} else {
				return nil, fmt.Errorf("Error: hostname, username and passwd are needed here for fortimanager")
			}
		}
	// default is for fortios provider
	default:
		{
			config := &tls.Config{}

			auth := auth.NewAuth(c.Hostname, c.Token, c.CABundle, c.Vdom)

			if auth.Hostname == "" {
				auth.GetEnvHostname()
			}

			if auth.Token == "" {
				auth.GetEnvToken()
			}

			if auth.CABundle == "" {
				auth.GetEnvCABundle()
			}

			if auth.CABundle != "" {
				f, err := os.Open(auth.CABundle)
				if err != nil {
					return nil, fmt.Errorf("Error reading CA Bundle: %s", err)
				}
				defer f.Close()

				caBundle, err := ioutil.ReadAll(f)
				if err != nil {
					return nil, fmt.Errorf("Error reading CA Bundle: %s", err)
				}

				pool := x509.NewCertPool()
				if !pool.AppendCertsFromPEM([]byte(caBundle)) {
					return nil, fmt.Errorf("Error reading CA Bundle")
				}
				config.RootCAs = pool
			}

			if c.Insecure == nil {
				b, _ := auth.GetEnvInsecure()
				config.InsecureSkipVerify = b
			} else {
				config.InsecureSkipVerify = *c.Insecure
			}

			if config.InsecureSkipVerify == false && auth.CABundle == "" {
				return nil, fmt.Errorf("Error getting CA Bundle, CA Bundle should be set when insecure is false")
			}

			tr := &http.Transport{
				TLSClientConfig: config,
			}

			client := &http.Client{
				Transport: tr,
				Timeout:   time.Second * 10,
			}

			fc := forticlient.NewClient(auth, client)

			fClient.Client = fc

			return &fClient, nil
		}
	}
}
