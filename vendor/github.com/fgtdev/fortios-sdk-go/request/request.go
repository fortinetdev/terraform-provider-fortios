package request

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

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

// Store URL request params
// Currently only used for altering firewall security policy position
// @dstId: policy dst id
// @alterPos: before or after
func (r *Request) FillUrlParams(dstId int, alterPos string) {

	if r.HTTPRequest.Form == nil {
		r.HTTPRequest.Form = make(map[string][]string)
	}

	r.HTTPRequest.Form.Set("policy_dst_id", strconv.Itoa(dstId))
	r.HTTPRequest.Form.Set("alter_position", alterPos)
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
	u := buildURL(r)

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

func buildURL(r *Request) string {
	u := "https://"
	u += r.Config.FwTarget
	u += r.Path
	u += "?"

	if r.Config.Auth.Vdom != "" {
		u += "vdom="
		u += r.Config.Auth.Vdom
		u += "&"
	}

	if r.HTTPRequest.Form.Get("alter_position") != "" && r.HTTPRequest.Form.Get("policy_dst_id") != "" {
		u += "action=move&"
		u += r.HTTPRequest.Form.Get("alter_position")
		u += "="
		u += r.HTTPRequest.Form.Get("policy_dst_id")
		u += "&"
	}

	u += "access_token="
	u += r.Config.Auth.Token

	return u
}

// Send request data to FortiOS with special URL paramaters.
// If errors are encountered, it returns the error.
func (r *Request) SendWithSpecialParams(s string) error {
	r.HTTPRequest.Header.Set("Content-Type", "application/json")
	u := buildURL(r)

	if s != "" {
		u += "&"
		u += s
	}

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