package request

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
	"strings"

	"github.com/fgtdev/fortios-sdk-go/config"
)

// Request describes the request to FortiOS service
type Request struct {
	Config       config.Config
	HTTPRequest  *http.Request
	HTTPResponse *http.Response
	Path         string
	Params       interface{}
	Data         *bytes.Buffer
}

// New creates reqeust object with http method, path, params and data,
// It will save the http request, path, etc. for the next operations
// such as sending data, getting response, etc.
// It returns the created request object to the gobal plugin client.
func New(c config.Config, method string, path string, params interface{}, data *bytes.Buffer) *Request {
	var h *http.Request

	if data == nil {
		h, _ = http.NewRequest(method, "", nil)
	} else {
		h, _ = http.NewRequest(method, "", data)
	}

	r := &Request{
		Config:      c,
		Path:        path,
		HTTPRequest: h,
		Params:      params,
		Data:        data,
	}
	return r
}

// Build Request header

// Build Request Sign/Login Info

// Send request data to FortiOS.
// If errors are encountered, it returns the error.
func (r *Request) Send() error {
	//Build FortiOS
	//build Sign/Login INfo

	//httpReq.URL, err = url.Parse(clientInfo.Endpoint + operation.HTTPPath)

	r.HTTPRequest.Header.Set("Content-Type", "application/json")
	u := buildURL(r.Config.FwTarget, r.Path, r.Config.Auth.Token, r.Config.Auth.Vdom)

	var err error
	r.HTTPRequest.URL, err = url.Parse(u)
	if err != nil {
		log.Fatal(err)
		return err
	}

	retry := 0
	for {
		//Send
		rsp, errdo := r.Config.HTTPCon.Do(r.HTTPRequest)
		r.HTTPResponse = rsp
		if errdo != nil {
			if strings.Contains(errdo.Error(), "x509: ") {
				err = fmt.Errorf("Error found: %s", errdo)
				break
			}

			if retry > 500 {
				err = fmt.Errorf("Error found: %s", errdo)
				break
			}
			time.Sleep(time.Duration(1) * time.Second)
			log.Printf("Error found: %s, will resend again %s, %d", errdo, u, retry)

			retry++

		} else {
			break
		}
	}

	return err
}

func buildURL(host string, path string, token string, vdom string) string {
	u := "https://"
	u += host
	u += path
	u += "?"

	if vdom != "" {
		u += "vdom="
		u += vdom
		u += "&"
	}

	u += "access_token="
	u += token

	return u
}
