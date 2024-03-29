package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

// JSONNetworkingRouteStatic contains the parameters for Create and Update API function
type JSONNetworkingRouteStatic struct {
	Dst             string `json:"dst,omitempty"`
	Gateway         string `json:"gateway"`
	Blackhole       string `json:"blackhole,omitempty"`
	Distance        string `json:"distance,omitempty"`
	Weight          string `json:"weight,omitempty"`
	Priority        string `json:"priority,omitempty"`
	Device          string `json:"device"`
	Comment         string `json:"comment"`
	Status          string `json:"status,omitempty"`
	InternetService int    `json:"internet-service,omitempty"`
}

// JSONCreateNetworkingRouteStaticOutput contains the output results for Create API function
type JSONCreateNetworkingRouteStaticOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       float64 `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// JSONUpdateNetworkingRouteStaticOutput contains the output results for Update API function
// Attention: The RESTful API changed the Mkey type from float64 in CREATE to string in UPDATE!
type JSONUpdateNetworkingRouteStaticOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// CreateNetworkingRouteStatic API operation for FortiOS creates a new static route.
// Returns the index value of the static route and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - static chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateNetworkingRouteStatic(params *JSONNetworkingRouteStatic) (output *JSONCreateNetworkingRouteStaticOutput, err error) {
	HTTPMethod := "POST"
	path := "/api/v2/cmdb/router/static"
	output = &JSONCreateNetworkingRouteStaticOutput{}
	locJSON, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	bytes := bytes.NewBuffer(locJSON)
	req := c.NewRequest(HTTPMethod, path, nil, bytes)
	err = req.Send()
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %s", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close() //#

	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	err = fortiAPIErrorFormat(result, string(body))

	if err == nil {
		if result["vdom"] != nil {
			output.Vdom = result["vdom"].(string)
		}
		if result["mkey"] != nil {
			output.Mkey = result["mkey"].(float64)
		}
	}

	return
}

// UpdateNetworkingRouteStatic API operation for FortiOS updates the specified static route.
// Returns the index value of the static route and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - static chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateNetworkingRouteStatic(params *JSONNetworkingRouteStatic, mkey string) (output *JSONUpdateNetworkingRouteStaticOutput, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/static"
	path += "/" + mkey
	output = &JSONUpdateNetworkingRouteStaticOutput{}
	locJSON, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	bytes := bytes.NewBuffer(locJSON)
	req := c.NewRequest(HTTPMethod, path, nil, bytes)
	err = req.Send()
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %s", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close() //#

	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}
	log.Printf("FOS-fortios response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	err = fortiAPIErrorFormat(result, string(body))

	if err == nil {
		if result["vdom"] != nil {
			output.Vdom = result["vdom"].(string)
		}

		if result["mkey"] != nil {
			output.Mkey = result["mkey"].(string)
		}
	}

	return
}

// DeleteNetworkingRouteStatic API operation for FortiOS deletes the specified static route.
// Returns error for service API and SDK errors.
// See the router - static chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteNetworkingRouteStatic(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/router/static"
	path += "/" + mkey

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.Send()
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %s", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close() //#

	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}
	log.Printf("FOS-fortios response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	err = fortiAPIErrorFormat(result, string(body))

	return
}

// ReadNetworkingRouteStatic API operation for FortiOS gets the static route
// with the specified index value.
// Returns the requested static route value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - static chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadNetworkingRouteStatic(mkey string) (output *JSONNetworkingRouteStatic, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/static"
	path += "/" + mkey

	output = &JSONNetworkingRouteStatic{}

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.Send()
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %s", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close() //#

	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}
	log.Printf("FOS-fortios reading response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	if fortiAPIHttpStatus404Checking(result) == true {
		output = nil
		return
	}

	err = fortiAPIErrorFormat(result, string(body))

	if err == nil {
		mapTmp := (result["results"].([]interface{}))[0].(map[string]interface{})

		if mapTmp == nil {
			err = fmt.Errorf("cannot get the results from the response")
			return
		}

		if mapTmp["dst"] != nil {
			output.Dst = mapTmp["dst"].(string)
		}
		if mapTmp["gateway"] != nil {
			output.Gateway = mapTmp["gateway"].(string)
		}
		if mapTmp["blackhole"] != nil {
			output.Blackhole = mapTmp["blackhole"].(string)
		}
		if mapTmp["distance"] != nil {
			output.Distance = strconv.Itoa(int(mapTmp["distance"].(float64)))
		}
		if mapTmp["weight"] != nil {
			output.Weight = strconv.Itoa(int(mapTmp["weight"].(float64)))
		}
		if mapTmp["priority"] != nil {
			output.Priority = strconv.Itoa(int(mapTmp["priority"].(float64)))
		}
		if mapTmp["device"] != nil {
			output.Device = mapTmp["device"].(string)
		}
		if mapTmp["comment"] != nil {
			output.Comment = mapTmp["comment"].(string)
		}
		if mapTmp["internet-service"] != nil {
			output.InternetService = int(mapTmp["internet-service"].(float64))
		}
		if mapTmp["status"] != nil {
			output.Status = mapTmp["status"].(string)
		}

	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}
