package forticlient

import (
	"net/http"
	"bytes"
	"github.com/fortios/fortios-sdk/config"
	"github.com/fortios/fortios-sdk/auth"
	"github.com/fortios/fortios-sdk/request")

//MultValue contains ... need to comment completely
type MultValue struct {
	Name string `json:"name"`
}

//MultValues contains ... need to comment completely
type MultValues []MultValue


type FortiSDKClient struct {
	Config   config.Config
	Retries  int
}

//ExpandStringList contains ... need to comment completely
func ExpandStringList(members []MultValue) []string {
	vs := make([]string, 0, len(members))
	for _, v := range members {
		c := v.Name
		vs = append(vs, c)
	}
	return vs
}


//Create new client object
func NewClient(host string, auth *auth.Auth, client *http.Client) *FortiSDKClient{
	c := &FortiSDKClient{ }

	c.Config.Auth = auth
	c.Config.HTTPCon = client
	c.Config.FwTarget = host

	return c
}

//Build new reqeust
func (c *FortiSDKClient) NewRequest(method string, path string, params interface{}, data *bytes.Buffer) *request.Request {
	return request.New(c.Config, method, path, params, data);
}
