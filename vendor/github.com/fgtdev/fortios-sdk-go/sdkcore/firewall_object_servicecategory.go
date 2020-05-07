package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/fgtdev/fortios-sdk-go/util"
)

// JSONFirewallObjectServiceCategoryItem contains the General parameters for Create and Update API function
type JSONFirewallObjectServiceCategoryItem struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

// JSONFirewallObjectServiceCategory contains the parameters for Create and Update API function
type JSONFirewallObjectServiceCategory struct {
	*JSONFirewallObjectServiceCategoryItem
}

// JSONCreateFirewallObjectServiceCategoryOutput contains the output results for Create API function
type JSONCreateFirewallObjectServiceCategoryOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// JSONUpdateFirewallObjectServiceCategoryOutput contains the output results for Update API function
type JSONUpdateFirewallObjectServiceCategoryOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// CreateFirewallObjectServiceCategory API operation for FortiOS creates a new firewall service category.
// Returns the index value of the firewall service category and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) CreateFirewallObjectServiceCategory(params *JSONFirewallObjectServiceCategory) (output *JSONCreateFirewallObjectServiceCategoryOutput, err error) {
	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall.service/category"
	output = &JSONCreateFirewallObjectServiceCategoryOutput{}
	locJSON, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Printf("POST: %s", string(locJSON))

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
	log.SetPrefix("CreateFirewallObjectServiceCategory")
	log.Printf("Path called %s", path)
	log.Printf("FortiOS response: %s", string(body))

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
					if result["error"] == -5 {
						err = fmt.Errorf("Category %s already exists.", result["mkey"])
					} else {
						err = fmt.Errorf("status is %s and error no is %.0f", result["status"], result["error"])
					}
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

// UpdateFirewallObjectServiceCategory API operation for FortiOS updates the specified firewall service category.
// Returns the index value of the firewall service and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) UpdateFirewallObjectServiceCategory(params *JSONFirewallObjectServiceCategory, mkey string) (output *JSONUpdateFirewallObjectServiceCategoryOutput, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall.service/category"
	path += "/" + EscapeURLString(mkey)
	output = &JSONUpdateFirewallObjectServiceCategoryOutput{}
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
	log.Printf("Path called %s", path)
	log.Printf("FortiOS response: %s", string(body))

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

// DeleteFirewallObjectServiceCategory API operation for FortiOS deletes the specified firewall service category.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) DeleteFirewallObjectServiceCategory(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall.service/category"
	path += "/" + EscapeURLString(mkey)

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

	log.Printf("Path called %s", path)
	log.Printf("FortiOS response: %s", string(body))

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

// ReadFirewallObjectServiceCategory API operation for FortiOS gets the firewall service category
// with the specified index value.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) ReadFirewallObjectServiceCategory(mkey string) (output *JSONFirewallObjectServiceCategory, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall.service/category"
	path += "/" + EscapeURLString(mkey)

	j1 := JSONFirewallObjectServiceCategoryItem{}

	output = &JSONFirewallObjectServiceCategory{
		JSONFirewallObjectServiceCategoryItem: &j1,
	}

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
	log.SetPrefix("ReadFirewallObjectServiceCategory")
	log.Printf("Path called %s", path)
	log.Printf("FortiOS response: %s", string(body))

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

		if mapTmp["name"] != nil {
			output.Name = mapTmp["name"].(string)
		}
		if mapTmp["comment"] != nil {
			output.Comment = mapTmp["comment"].(string)
		}

	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}
