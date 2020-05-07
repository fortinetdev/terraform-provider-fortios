package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/fgtdev/fortios-sdk-go/util"
)

// JSONNetworkingRouteStatic contains the parameters for Create and Update API function
type JSONNetworkingRouteStatic struct {
	Dst             string `json:"dst"`
	Gateway         string `json:"gateway"`
	Blackhole       string `json:"blackhole"`
	Distance        string `json:"distance"`
	Weight          string `json:"weight"`
	Priority        string `json:"priority"`
	Device          string `json:"device"`
	Comment         string `json:"comment"`
	Status          string `json:"status"`
	InternetService int    `json:"internet-service"`
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
	e, req := c.NewRequest(HTTPMethod, path, nil, bytes)
	if e != nil {
		err = fmt.Errorf("new request error %s", e)
		return
	}
	
	err = req.Send()
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %s", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

	if result != nil {
		if result["vdom"] != nil {
			output.Vdom = result["vdom"].(string)
		}
		if result["mkey"] != nil {
			output.Mkey = result["mkey"].(float64)
		}
		if result["status"] != nil {
			if result["status"] != "success" {
				if result["error"] != nil {
					err = fmt.Errorf("status is %s and error no is %.0f", result["status"], result["error"])
				} else {
					err = fmt.Errorf("status is %s and error no is not found", result["status"])
				}

				if result["http_status"] != nil {
					err = fmt.Errorf("%s, details: %s", err, util.HttpStatus2Str(int(result["http_status"].(float64))))
				} else {
					err = fmt.Errorf("%s, and http_status no is not found", err)
				}

				return
			}
			output.Status = result["status"].(string)
		} else {
			err = fmt.Errorf("cannot get status from the response")
			return
		}
		if result["http_status"] != nil {
			output.HTTPStatus = result["http_status"].(float64)
		}
	} else {
		err = fmt.Errorf("cannot get the right response")
		return
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
	e, req := c.NewRequest(HTTPMethod, path, nil, bytes)
	if e != nil {
		err = fmt.Errorf("new request error %s", e)
		return
	}
	
	err = req.Send()
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %s", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}
	log.Printf("FOS-fortios response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

	if result != nil {
		if result["vdom"] != nil {
			output.Vdom = result["vdom"].(string)
		}
		if result["mkey"] != nil {
			output.Mkey = result["mkey"].(string)
		}
		if result["status"] != nil {
			if result["status"] != "success" {
				if result["error"] != nil {
					err = fmt.Errorf("status is %s and error no is %.0f", result["status"], result["error"])
				} else {
					err = fmt.Errorf("status is %s and error no is not found", result["status"])
				}

				if result["http_status"] != nil {
					err = fmt.Errorf("%s, details: %s", err, util.HttpStatus2Str(int(result["http_status"].(float64))))
				} else {
					err = fmt.Errorf("%s, and http_status no is not found", err)
				}

				return
			}
			output.Status = result["status"].(string)
		} else {
			err = fmt.Errorf("cannot get status from the response")
			return
		}
		if result["http_status"] != nil {
			output.HTTPStatus = result["http_status"].(float64)
		}
	} else {
		err = fmt.Errorf("cannot get the right response")
		return
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

	e, req := c.NewRequest(HTTPMethod, path, nil, nil)
	if e != nil {
		err = fmt.Errorf("new request error %s", e)
		return
	}
	
	err = req.Send()
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %s", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}
	log.Printf("FOS-fortios response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

	if result != nil {
		if result["status"] == nil {
			err = fmt.Errorf("cannot get status from the response")
			return
		}

		if result["status"] != "success" {
			if result["error"] != nil {
				err = fmt.Errorf("status is %s and error no is %.0f", result["status"], result["error"])
			} else {
				err = fmt.Errorf("status is %s and error no is not found", result["status"])
			}

			if result["http_status"] != nil {
				err = fmt.Errorf("%s, details: %s", err, util.HttpStatus2Str(int(result["http_status"].(float64))))
			} else {
				err = fmt.Errorf("%s, and http_status no is not found", err)
			}

			return
		}

	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

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

	e, req := c.NewRequest(HTTPMethod, path, nil, nil)
	if e != nil {
		err = fmt.Errorf("new request error %s", e)
		return
	}
	
	err = req.Send()
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %s", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}
	log.Printf("FOS-fortios reading response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

	if result != nil {
		if result["http_status"] == nil {
			err = fmt.Errorf("cannot get http_status from the response")
			return
		}

		if result["http_status"] == 404.0 {
			output = nil
			return
		}

		if result["status"] == nil {
			err = fmt.Errorf("cannot get status from the response")
			return
		}

		if result["status"] != "success" {
			if result["error"] != nil {
				err = fmt.Errorf("status is %s and error no is %.0f", result["status"], result["error"])
			} else {
				err = fmt.Errorf("status is %s and error no is not found", result["status"])
			}

			if result["http_status"] != nil {
				err = fmt.Errorf("%s, details: %s", err, util.HttpStatus2Str(int(result["http_status"].(float64))))
			} else {
				err = fmt.Errorf("%s, and http_status no is not found", err)
			}

			return
		}

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
