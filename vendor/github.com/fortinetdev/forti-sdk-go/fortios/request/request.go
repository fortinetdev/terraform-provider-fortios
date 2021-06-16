package request

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
	"regexp"

	"github.com/fortinetdev/forti-sdk-go/fortios/config"
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
	return r.Send2(15, false)
}

// Send2 request data to FortiOS with bIgnoreVdom.
// If errors are encountered, it returns the error.
func (r *Request) Send2(retries int, ignvdom bool) error {
	//Build FortiOS
	//build Sign/Login INfo

	//httpReq.URL, err = url.Parse(clientInfo.Endpoint + operation.HTTPPath)

	r.HTTPRequest.Header.Set("Content-Type", "application/json")
	u := ""
	if ignvdom == true {
		u = buildURLWithoutVdom(r)
	} else {
		u = buildURL(r)
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
				err = fmt.Errorf("Error found: %v", filterapikey(errdo.Error()))
				break
			}

			if retry >  retries {
				err = fmt.Errorf("lost connection to firewall with error: %v", filterapikey(errdo.Error()))
				break
			}
			time.Sleep(time.Second)
			log.Printf("Error found: %v, will resend again %s, %d", filterapikey(errdo.Error()), u, retry)

			retry++

		} else {
			break
		}
	}

	return err
}


func buildURL3(r *Request, vdomparam string) string {
	u := "https://"
	u += r.Config.FwTarget
	u += r.Path
	u += "?"

	if vdomparam != "" {
		u += "vdom="
		u += vdomparam
		u += "&"
	} else {
		if r.Config.Auth.Vdom != "" {
			u += "vdom="
			u += r.Config.Auth.Vdom
			u += "&"
		}
	}

	u += "access_token="
	u += r.Config.Auth.Token

	return u
}

// Send3 request data to FortiOS with custom vdom.
// If errors are encountered, it returns the error.
func (r *Request) Send3(vdomparam string) error {
	retries := 15

	r.HTTPRequest.Header.Set("Content-Type", "application/json")
	u :=  buildURL3(r, vdomparam)
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
				err = fmt.Errorf("Error found: %v", filterapikey(errdo.Error()))
				break
			}

			if retry >  retries {
				err = fmt.Errorf("lost connection to firewall with error: %v", filterapikey(errdo.Error()))
				break
			}
			time.Sleep(time.Second)
			log.Printf("Error found: %v, will resend again %s, %d", filterapikey(errdo.Error()), u, retry)

			retry++

		} else {
			break
		}
	}

	return err
}



func filterapikey(v string) string {
	re, _ := regexp.Compile("access_token=.*?\"");
	res := re.ReplaceAllString(v, "access_token=***************\"");

	return res
}

func buildURLWithoutVdom(r *Request) string {
	u := "https://"
	u += r.Config.FwTarget
	u += r.Path
	u += "?"

	u += "access_token="
	u += r.Config.Auth.Token

	return u
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

	u += "access_token="
	u += r.Config.Auth.Token

	return u
}

// SendWithSpecialParams sends request data to FortiOS with special URL paramaters.
// If errors are encountered, it returns the error.
func (r *Request) SendWithSpecialParams(s, vdomparam string) error {
	r.HTTPRequest.Header.Set("Content-Type", "application/json")
	u := buildURL3(r, vdomparam)

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
				err = fmt.Errorf("Error found: %v", filterapikey(errdo.Error()))
				break
			}

			if retry > 15 {
				err = fmt.Errorf("Error found: %v", filterapikey(errdo.Error()))
				break
			}
			time.Sleep(time.Duration(1) * time.Second)
			log.Printf("Error found: %v, will resend again %s, %d", filterapikey(errdo.Error()), u, retry)

			retry++

		} else {
			break
		}
	}

	return err
}
