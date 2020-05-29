package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"github.com/fortinetdev/forti-sdk-go/fortios/request"
)

// JSONJSONGenericAPI contains the parameters for Create API function
type JSONJSONGenericAPI struct {
	Path string `json:"path"`
	Method string `json:"method"`
	Specialparams string `json:"specialparams"`
	Json string `json:"json"`
	Response string `json:"response"`
}

// CreateJSONGenericAPI API operation for FortiOS sends request to FortiGate/FortiOS APIs.
// Returns the response from FortiGate or FortiOS .
func (c *FortiSDKClient) CreateJSONGenericAPI(params *JSONJSONGenericAPI) (res string, err error) {
	HTTPMethod := params.Method
	path := params.Path
	specialparams := params.Specialparams

	var req *request.Request

	if params.Json != "" {
		locJSON := []byte(params.Json)

		var js json.RawMessage
		if json.Unmarshal([]byte(params.Json), &js) != nil {
			err = fmt.Errorf("JSON format Error")
			return
		}

		log.Printf("FOS-fortios resquest1: %s", locJSON)
		bytes := bytes.NewBuffer(locJSON)

		req = c.NewRequest(HTTPMethod, path, nil, bytes)
	} else {
		req = c.NewRequest(HTTPMethod, path, nil, nil)
	}

	err = req.SendWithSpecialParams(specialparams)
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %s", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}
	log.Printf("FOS-fortios response1: %s", string(body))

	res = string(body)
	req.HTTPResponse.Body.Close()

	return
}

// UpdateJSONGenericAPI API operation for FortiOS
func (c *FortiSDKClient) UpdateJSONGenericAPI(params *JSONJSONGenericAPI, mkey string) (res string, err error) {
	return
}

// DeleteJSONGenericAPI API operation for FortiOS
func (c *FortiSDKClient) DeleteJSONGenericAPI(mkey string) (err error) {
	return
}

// ReadJSONGenericAPI API operation for FortiOS
func (c *FortiSDKClient) ReadJSONGenericAPI(mkey string) (output *JSONJSONGenericAPI, err error) {
	return
}
