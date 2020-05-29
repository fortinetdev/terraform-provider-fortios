package config

import (
	"net/http"

	"github.com/fortinetdev/forti-sdk-go/fortios/auth"
)

// Config provides configuration to a FortiOS client instance
// It saves authentication information and a http connection
// for FortiOS Client instance to create New connction to FortiOS
// and Send data to FortiOS,  etc. (needs to be extended later.)
type Config struct {
	Auth     *auth.Auth
	HTTPCon  *http.Client
	FwTarget string
}
