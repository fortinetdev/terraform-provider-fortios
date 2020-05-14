package fortios

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
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

	FMG_Hostname string
	FMG_Username string
	FMG_Passwd   string
	FMG_Insecure *bool
	FMG_CABundle string
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

	bFOSExist := bFortiOSHostnameExist(c)
	bFMGExist := bFortiManagerHostnameExist(c)

	if bFOSExist {
		err := createFortiOSClient(&fClient, c)
		if err != nil {
			return nil, fmt.Errorf("Error create fortios client: %s", err)
		}
	}

	if bFMGExist {
		err := createFortiManagerClient(&fClient, c)
		if err != nil {
			return nil, fmt.Errorf("Error create fortimanager client: %s", err)
		}
	} else {
		fClient.ClientFortimanager = fmgclient.NewEmptyClient()
	}

	if !bFOSExist && !bFMGExist {
		return nil, fmt.Errorf("FortiOS or FortiManager, at least one of their hostnames should be set")
	}

	return &fClient, nil
}

func bFortiOSHostnameExist(c *Config) bool {
	if c.Hostname == "" {
		if os.Getenv("FORTIOS_ACCESS_HOSTNAME") == "" {
			return false
		}
	}

	return true
}

func bFortiManagerHostnameExist(c *Config) bool {
	if c.FMG_Hostname == "" {
		if os.Getenv("FORTIOS_FMG_HOSTNAME") == "" {
			return false
		}
	}

	return true
}

func createFortiOSClient(fClient *FortiClient, c *Config) error {
	config := &tls.Config{}

	auth := auth.NewAuth(c.Hostname, c.Token, c.CABundle, c.Vdom)

	if auth.Hostname == "" {
		_, err := auth.GetEnvHostname()
		if err != nil {
			return fmt.Errorf("Error reading Hostname")
		}
	}

	if auth.Token == "" {
		_, err := auth.GetEnvToken()
		if err != nil {
			return fmt.Errorf("Error reading Token")
		}
	}

	if auth.CABundle == "" {
		auth.GetEnvCABundle()
	}

	if auth.CABundle != "" {
		f, err := os.Open(auth.CABundle)
		if err != nil {
			return fmt.Errorf("Error reading CA Bundle: %s", err)
		}
		defer f.Close()

		caBundle, err := ioutil.ReadAll(f)
		if err != nil {
			return fmt.Errorf("Error reading CA Bundle: %s", err)
		}

		pool := x509.NewCertPool()
		if !pool.AppendCertsFromPEM([]byte(caBundle)) {
			return fmt.Errorf("Error reading CA Bundle")
		}
		config.RootCAs = pool
	}

	if c.Insecure == nil {
		// defaut value for Insecure is false
		b, _ := auth.GetEnvInsecure()
		config.InsecureSkipVerify = b
	} else {
		config.InsecureSkipVerify = *c.Insecure
	}

	if config.InsecureSkipVerify == false && auth.CABundle == "" {
		return fmt.Errorf("Error getting CA Bundle, CA Bundle should be set when insecure is false")
	}

	tr := &http.Transport{
		TLSClientConfig: config,
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 360,
	}

	fc := forticlient.NewClient(auth, client)

	fClient.Client = fc

	return nil
}

func createFortiManagerClient(fClient *FortiClient, c *Config) error {
	if c.FMG_Hostname == "" {
		c.FMG_Hostname = os.Getenv("FORTIOS_FMG_HOSTNAME")
	}
	if c.FMG_Username == "" {
		c.FMG_Username = os.Getenv("FORTIOS_FMG_USERNAME")
	}
	if c.FMG_Passwd == "" {
		c.FMG_Passwd = os.Getenv("FORTIOS_FMG_PASSWORD")
	}
	if c.FMG_CABundle == "" {
		c.FMG_CABundle = os.Getenv("FORTIOS_FMG_CABUNDLE")
	}
	if c.FMG_Hostname == "" || c.FMG_Username == "" || c.FMG_Passwd == "" {
		return fmt.Errorf("Error: hostname, username and passwd are needed here for fortimanager")
	}

	config := &tls.Config{}
	if c.FMG_Insecure == nil {
		insec := os.Getenv("FORTIOS_FMG_INSECURE")
		if insec == "" || insec == "false" {
			config.InsecureSkipVerify = false
		} else if insec == "true" {
			config.InsecureSkipVerify = true
		} else {
			return fmt.Errorf("Error: FORTIOS_FMG_INSECURE shoule be false or true")
		}
	} else {
		config.InsecureSkipVerify = *c.FMG_Insecure
	}

	if !config.InsecureSkipVerify {
		if c.FMG_CABundle != "" {
			f, err := os.Open(c.FMG_CABundle)
			if err != nil {
				return fmt.Errorf("Error open CA Bundle file: %s", err)
			}
			defer f.Close()

			caBundle, err := ioutil.ReadAll(f)
			if err != nil {
				return fmt.Errorf("Error reading CA Bundle: %s", err)
			}

			pool := x509.NewCertPool()
			if !pool.AppendCertsFromPEM([]byte(caBundle)) {
				return fmt.Errorf("Error appending cert from PEM")
			}
			config.RootCAs = pool
		} else {
			return fmt.Errorf("Error getting CA Bundle, CA Bundle should be set when insecure is false")
		}
	}

	tr := &http.Transport{
		TLSClientConfig: config,
	}

	client := &http.Client{
		Transport: tr,
	}
	fClient.ClientFortimanager = fmgclient.NewClient(c.FMG_Hostname, c.FMG_Username, c.FMG_Passwd, client)

	fClient.ClientFortimanager.DebugNum = 0

	session, err := fClient.ClientFortimanager.Login()
	if err != nil {
		return fmt.Errorf("FortiManager Login failed")
	}
	fClient.ClientFortimanager.SessionString = session
	fClient.ClientFortimanager.Init = true

	return nil
}
