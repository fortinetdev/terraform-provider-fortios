package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
	req := c.NewRequest(HTTPMethod, path, nil, bytes)
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

	if result != nil && result["status"] != "success" {
		if result["error"] == -5 {
			err = fmt.Errorf("Category %s already exists.", result["mkey"])
			return 
		}
	}

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
	req := c.NewRequest(HTTPMethod, path, nil, bytes)
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

// DeleteFirewallObjectServiceCategory API operation for FortiOS deletes the specified firewall service category.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) DeleteFirewallObjectServiceCategory(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall.service/category"
	path += "/" + EscapeURLString(mkey)

	req := c.NewRequest(HTTPMethod, path, nil, nil)
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

	err = fortiAPIErrorFormat(result, string(body))

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

	req := c.NewRequest(HTTPMethod, path, nil, nil)
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
