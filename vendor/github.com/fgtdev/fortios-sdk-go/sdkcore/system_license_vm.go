package forticlient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"fmt"
)

// JSONSystemLicenseVM contains the parameters for Create and Update API function
type JSONSystemLicenseVM struct {
	FileContent       string `json:"file_content"`
}

// JSONCreateSystemLicenseVMOutput contains the output results for Create API function
type JSONCreateSystemLicenseVMOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// JSONUpdateSystemLicenseVMOutput contains the output results for Update API function
// Attention: Considering scalability, the previous structure and the current structure may change differently
type JSONUpdateSystemLicenseVMOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// CreateSystemLicenseVM API operation for FortiOS uploads a new VM License File.
// Returns the execution result when the request executes successfully.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) CreateSystemLicenseVM(params *JSONSystemLicenseVM) (output *JSONCreateSystemLicenseVMOutput, err error) {
	HTTPMethod := "POST"
	path := "/api/v2/monitor/system/vmlicense/upload"
	output = &JSONCreateSystemLicenseVMOutput{}
	locJSON, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return
	}

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
		if result["http_status"] != nil {
			output.HTTPStatus = result["http_status"].(float64)
		}
	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}

// UpdateSystemLicenseVM API operation for FortiOS
func (c *FortiSDKClient) UpdateSystemLicenseVM(params *JSONSystemLicenseVM, mkey string) (output *JSONUpdateSystemLicenseVMOutput, err error) {
	// HTTPMethod := "PUT"
	// path := "/api/v2/monitor/system/vmlicense/upload"
	// path += "/" + mkey
	// output = &JSONUpdateSystemLicenseVMOutput{}
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
	//		if result["status"] != "success" {
	//			err = fmt.Errorf("cannot get the right response")
	//			return
	//		}
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

// DeleteSystemLicenseVM API operation for FortiOS
func (c *FortiSDKClient) DeleteSystemLicenseVM(mkey string) (err error) {
	// HTTPMethod := "DELETE"
	// path := "/api/v2/monitor/system/vmlicense/upload"
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

// ReadSystemLicenseVM API operation for FortiOS
func (c *FortiSDKClient) ReadSystemLicenseVM(mkey string) (output *JSONSystemLicenseVM, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/monitor/license/status/select"

	output = &JSONSystemLicenseVM{}

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
			if k == "vm" {
				bFind = true;

				z := v.(map[string]interface{})
				if z == nil {
					err = fmt.Errorf("cannot get vm property from the response")
					return
				}

				if z["valid"] == nil {
					err = fmt.Errorf("cannot get vm.valid property from the response")
					return
				}

				if z["valid"] == true {
					output.FileContent = "********"
				} else {
					output = nil
					return
				}
			}
		}

		if bFind == false {
			err = fmt.Errorf("cannot get vm property from the response")
			return
		}

	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}
