package request

import (
	"net/http"
	"net/url"
	"log"
	"bytes"
	"github.com/fortios/fortios-sdk/config"
)

//Request .....
type Request struct {
	Config       config.Config
	HTTPRequest  *http.Request
	HTTPResponse *http.Response
	Path         string
	Params       interface{}
	Data         *bytes.Buffer
	//Body                   io.ReadSeeker
}

//New wil create Request object with ALL information for next send
//Need Client Structure Information!!!!!!since cannot recycle refer,
//so need to add a public module for configuration
func New(c config.Config, method string, path string, params interface{}, data *bytes.Buffer) *Request {
	log.Printf("shengh.............request New1\n")

	var h *http.Request

	if (data == nil) {
		h, _ = http.NewRequest(method, "", nil)
	} else {
		h, _ = http.NewRequest(method, "", data)
	}
	
	log.Printf("shengh.............request New2\n")

	r := &Request{
		Config:      c,
		Path:        path,
		HTTPRequest: h,
		Params:      params,
		Data:        data,
	}
	return r
}

//Build Request header

//Build Request Sign/Login Info

//Send Request to Firewall
func (r *Request) Send() error {
	//Build FortiOS
	//build Sign/Login INfo

	//httpReq.URL, err = url.Parse(clientInfo.Endpoint + operation.HTTPPath)

	r.HTTPRequest.Header.Set("Content-Type", "application/json")
	u := buildURL(r.Config.FwTarget, r.Path, r.Config.Auth.Token)

	var err error
	r.HTTPRequest.URL, err = url.Parse(u)
	if err != nil {
		log.Fatal(err)
		return err
	}

	//Send
	rsp, err := r.Config.HTTPCon.Do(r.HTTPRequest)
	r.HTTPResponse = rsp
	if err != nil {
		log.Printf("shengh.............shengh ERROR\n")
		log.Fatal(err)
	}

	return err
}

func buildURL(host string, path string, token string) string {
	u := "https://"
	u += host
	u += path
	u += "?access_token="
	u += token

	return u
}