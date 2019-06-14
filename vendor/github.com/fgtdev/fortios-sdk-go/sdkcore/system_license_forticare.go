package forticlient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"fmt"
)

// JSONSystemLicenseFortiCare contains the parameters for Create and Update API function
type JSONSystemLicenseFortiCare struct {
	RegistrationCode       string `json:"registration_code"`
}

// JSONCreateSystemLicenseFortiCareOutput contains the output results for Create API function
type JSONCreateSystemLicenseFortiCareOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// JSONUpdateSystemLicenseFortiCareOutput contains the output results for Update API function
// Attention: Considering scalability, the previous structure and the current structure may change differently
type JSONUpdateSystemLicenseFortiCareOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// CreateSystemLicenseFortiCare API operation for FortiOS commits a module registration code.
// Returns the execution result when the request executes successfully.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) CreateSystemLicenseFortiCare(params *JSONSystemLicenseFortiCare) (output *JSONCreateSystemLicenseFortiCareOutput, err error) {
	HTTPMethod := "POST"
	path := "/api/v2/monitor/registration/forticare/add-license"
	output = &JSONCreateSystemLicenseFortiCareOutput{}
	locJSON, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	// different branch for different version
	// v, err := c.GetDeviceVersion()

	// if v != "" && strings.Contains(v, "6.2"){
	// 	v = "doesn't support " + v
	// 	err = fmt.Errorf(v)
	// 	return
	// }


	bytes := bytes.NewBuffer(locJSON)
	req := c.NewRequest(HTTPMethod, path, nil, bytes)
	err = req.Send()
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request")
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body")
		return
	}
	log.Printf("FOS-fortios response(create): %s", string(body))

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
					err = fmt.Errorf("%s and http_status no is %.0f", err, result["http_status"])
				} else {
					err = fmt.Errorf("%s and and http_status no is not found", err)
				}

				return
			}
			output.Status = result["status"].(string)
		} else {
			err = fmt.Errorf("cannot get status from the response")
			return
		}

		if result["results"] == nil {
			err = fmt.Errorf("cannot get results from the response")
			return
		}

		mapTmp := result["results"].(map[string]interface {});

		if mapTmp == nil {
			err = fmt.Errorf("cannot get results from the response")
			return
		}

		if mapTmp["forticare_error"] != nil {
			s :=  mapTmp["forticare_error"].(string)
			if s != "" {
					err = fmt.Errorf("cannot get the right response " + s)
			}
			return
		}

	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}

// UpdateSystemLicenseFortiCare API operation for FortiOS
func (c *FortiSDKClient) UpdateSystemLicenseFortiCare(params *JSONSystemLicenseFortiCare, mkey string) (output *JSONUpdateSystemLicenseFortiCareOutput, err error) {
	// HTTPMethod := "PUT"
	// path := "/api/v2/monitor/registration/forticare/add-license"
	// path += "/" + mkey
	// output = &JSONUpdateSystemLicenseFortiCareOutput{}
	// locJSON, err := json.Marshal(params)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// bytes := bytes.NewBuffer(locJSON)
	// req := c.NewRequest(HTTPMethod, path, nil, bytes)
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
	//	 	if result["status"] != "success" {
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

// DeleteSystemLicenseFortiCare API operation for FortiOS
func (c *FortiSDKClient) DeleteSystemLicenseFortiCare(mkey string) (err error) {
	// HTTPMethod := "DELETE"
	// path := "/api/v2/monitor/registration/forticare/add-license"
	// path += "/" + mkey

	// req := c.NewRequest(HTTPMethod, path, nil, nil)
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
	//
	//	if result["status"] != "success" {
	//		err = fmt.Errorf("cannot get the right response")
	//		return
	//	}
 	//
	// } else {
	// 	err = fmt.Errorf("cannot get the right response")
	// 	return
	// }

	return
}

// ReadSystemLicenseFortiCare API operation for FortiOS
func (c *FortiSDKClient) ReadSystemLicenseFortiCare(mkey string) (output *JSONSystemLicenseFortiCare, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/monitor/license/status/select"

	output = &JSONSystemLicenseFortiCare{}

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.Send()
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request")
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body")
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
				err = fmt.Errorf("%s and http_status no is %.0f", err, result["http_status"])
			} else {
				err = fmt.Errorf("%s and and http_status no is not found", err)
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

		for k,v := range mapTmp {
			if k == "forticare" {
				bFind = true;

				z := v.(map[string]interface{})
				if z == nil {
					err = fmt.Errorf("cannot get forticare property from the response")
					return
				}

				if z["status"] == nil {
					err = fmt.Errorf("cannot get forticare.status property from the response")
					return
				}

				if z["status"] == "registered" {
					output.RegistrationCode = "********"
				} else {
					output = nil
					return
				}
			}
		}

		if bFind == false {
			err = fmt.Errorf("cannot get forticare property from the response")
			return
		}

	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}
