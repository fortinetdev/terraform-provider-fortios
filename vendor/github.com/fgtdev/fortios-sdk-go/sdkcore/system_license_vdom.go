package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/fgtdev/fortios-sdk-go/util"
)

// JSONSystemLicenseVDOM contains the parameters for Create and Update API function
type JSONSystemLicenseVDOM struct {
	License string `json:"license"`
}

// JSONCreateSystemLicenseVDOMOutput contains the output results for Create API function
type JSONCreateSystemLicenseVDOMOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// JSONUpdateSystemLicenseVDOMOutput contains the output results for Update API function
// Attention: The RESTful API changed the Mkey type from float64 in CREATE to string in UPDATE!
type JSONUpdateSystemLicenseVDOMOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// CreateSystemLicenseVDOM API operation for FortiOS creates a new ----------------
// Returns the index value of the ---------------- and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the ---------------- chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemLicenseVDOM(params *JSONSystemLicenseVDOM) (output *JSONCreateSystemLicenseVDOMOutput, err error) {
	HTTPMethod := "POST"
	path := "/api/v2/monitor/registration/vdom/add-license"
	output = &JSONCreateSystemLicenseVDOMOutput{}
	locJSON, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	// different branch for different version
	// v, err := c.GetDeviceVersion()

	// if v != "" && !strings.Contains(v, "6.2"){
	// 	v = "doesn't support " + v
	// 	err = fmt.Errorf(v)
	// 	return
	// }

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

// UpdateSystemLicenseVDOM API operation for FortiOS updates the specified ----------------
// Returns the index value of the ---------------- and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the ---------------- chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLicenseVDOM(params *JSONSystemLicenseVDOM, mkey string) (output *JSONUpdateSystemLicenseVDOMOutput, err error) {
	// HTTPMethod := "PUT"
	// path := "/api/v2/monitor/registration/vdom/add-license"
	// path += "/" + mkey
	// output = &JSONUpdateSystemLicenseVDOMOutput{}
	// locJSON, err := json.Marshal(params)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// bytes := bytes.NewBuffer(locJSON)
	// e, req := c.NewRequest(HTTPMethod, path, nil, bytes)
	// err = req.Send()

	// body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	// log.Printf("FOS-fortios response: %s", string(body))

	// var result map[string]interface{}
	// json.Unmarshal([]byte(string(body)), &result)

	// req.HTTPResponse.Body.Close()

	// if result != nil {
	// 	if result["vdom"] != nil {
	// 		output.Vdom = result["vdom"].(string)
	// 	}
	// 	if result["mkey"] != nil {
	// 		output.Mkey = result["mkey"].(string)
	// 	}
	// 	if result["status"] != nil {
	// 		if result["status"] != "success" {
	// 			err = fmt.Errorf("cannot get the right response")
	// 			return
	// 		}
	// 		output.Status = result["status"].(string)
	// 	} else {
	// 		err = fmt.Errorf("cannot get the right response")
	// 		return
	// 	}
	// 	if result["http_status"] != nil {
	// 		output.HTTPStatus = result["http_status"].(float64)
	// 	}
	// } else {
	// 	err = fmt.Errorf("cannot get the right response")
	// 	return
	// }

	return
}

// DeleteSystemLicenseVDOM API operation for FortiOS deletes the specified ----------------
// Returns error for service API and SDK errors.
// See the ---------------- chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLicenseVDOM(mkey string) (err error) {
	// HTTPMethod := "DELETE"
	// path := "/api/v2/monitor/registration/vdom/add-license"
	// path += "/" + mkey

	// e, req := c.NewRequest(HTTPMethod, path, nil, nil)
	// err = req.Send()

	// body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	// log.Printf("FOS-fortios response: %s", string(body))

	// var result map[string]interface{}
	// json.Unmarshal([]byte(string(body)), &result)

	// req.HTTPResponse.Body.Close()

	// if result != nil {
	// 	if result["status"] == nil {
	// 		err = fmt.Errorf("cannot get the right response")
	// 		return
	// 	}

	// 	if result["status"] != "success" {
	// 		err = fmt.Errorf("cannot get the right response")
	// 		return
	// 	}

	// } else {
	// 	err = fmt.Errorf("cannot get the right response")
	// 	return
	// }

	return
}

// ReadSystemLicenseVDOM API operation for FortiOS gets the ----------------
// with the specified index value.
// Returns the requested ---------------- when the request executes successfully.
// Returns error for service API and SDK errors.
// See the ---------------- chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLicenseVDOM(mkey string) (output *JSONSystemLicenseVDOM, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/monitor/license/status/select"

	output = &JSONSystemLicenseVDOM{}

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

		if result["results"] == nil {
			err = fmt.Errorf("cannot get results from the response")
			return
		}

		mapTmp := (result["results"].(map[string]interface{}))

		if mapTmp == nil {
			err = fmt.Errorf("cannot get results from the response")
			return
		}

		bFind := false

		for k, v := range mapTmp {
			if k == "vdom" {
				bFind = true

				z := v.(map[string]interface{})
				if z == nil {
					err = fmt.Errorf("cannot get vdom property from the response")
					return
				}

				if z["used"] == nil {
					err = fmt.Errorf("cannot get vdom.used property from the response")
					return
				}

				if z["used"] == 1.0 {
					output.License = "********"
				} else {
					output = nil
					return
				}
			}
		}

		if bFind == false {
			err = fmt.Errorf("cannot get vdom property from the response")
			return
		}

	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}
