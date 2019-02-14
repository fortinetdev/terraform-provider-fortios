package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	//"strconv"
)

// JSONSystemAdminAdministrator contains the parameters for Create and Update API function
type JSONSystemAdminAdministrator struct {
	Name        string `json:"name"`
	Password    string `json:"password"`
	Trusthost1  string `json:"trusthost1"`
	Trusthost2  string `json:"trusthost2"`
	Trusthost3  string `json:"trusthost3"`
	Trusthost4  string `json:"trusthost4"`
	Trusthost5  string `json:"trusthost5"`
	Trusthost6  string `json:"trusthost6"`
	Trusthost7  string `json:"trusthost7"`
	Trusthost8  string `json:"trusthost8"`
	Trusthost9  string `json:"trusthost9"`
	Trusthost10 string `json:"trusthost10"`
	Accprofile  string `json:"accprofile"`
	Comments    string `json:"comments"`
}

// JSONSystemAdminAdministrator2 contains the parameters for Create and Update API function
type JSONSystemAdminAdministrator2 struct {
	Name        string `json:"name"`
	Trusthost1  string `json:"trusthost1"`
	Trusthost2  string `json:"trusthost2"`
	Trusthost3  string `json:"trusthost3"`
	Trusthost4  string `json:"trusthost4"`
	Trusthost5  string `json:"trusthost5"`
	Trusthost6  string `json:"trusthost6"`
	Trusthost7  string `json:"trusthost7"`
	Trusthost8  string `json:"trusthost8"`
	Trusthost9  string `json:"trusthost9"`
	Trusthost10 string `json:"trusthost10"`
	Accprofile  string `json:"accprofile"`
	Comments    string `json:"comments"`
}

// JSONCreateSystemAdminAdministratorOutput contains the output results for Create API function
type JSONCreateSystemAdminAdministratorOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// JSONUpdateSystemAdminAdministratorOutput contains the output results for Update API function
// Attention: Considering scalability, the previous structure and the current structure may change differently
type JSONUpdateSystemAdminAdministratorOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// CreateSystemAdminAdministrator API operation for FortiOS creates a new administrator account.
// Returns the index value of the administrator account and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemAdminAdministrator(params *JSONSystemAdminAdministrator) (output *JSONCreateSystemAdminAdministratorOutput, err error) {
	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/admin"
	output = &JSONCreateSystemAdminAdministratorOutput{}

	if params.Trusthost1 == "" {
		params.Trusthost1 = "0.0.0.0 0.0.0.0"
	}

	if params.Trusthost2 == "" {
		params.Trusthost2 = "0.0.0.0 0.0.0.0"
	}

	if params.Trusthost3 == "" {
		params.Trusthost3 = "0.0.0.0 0.0.0.0"
	}

	if params.Trusthost4 == "" {
		params.Trusthost4 = "0.0.0.0 0.0.0.0"
	}

	if params.Trusthost5 == "" {
		params.Trusthost5 = "0.0.0.0 0.0.0.0"
	}

	if params.Trusthost6 == "" {
		params.Trusthost6 = "0.0.0.0 0.0.0.0"
	}

	if params.Trusthost7 == "" {
		params.Trusthost7 = "0.0.0.0 0.0.0.0"
	}

	if params.Trusthost8 == "" {
		params.Trusthost8 = "0.0.0.0 0.0.0.0"
	}

	if params.Trusthost9 == "" {
		params.Trusthost9 = "0.0.0.0 0.0.0.0"
	}

	if params.Trusthost10 == "" {
		params.Trusthost10 = "0.0.0.0 0.0.0.0"
	}

	locJSON, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	bytes := bytes.NewBuffer(locJSON)
	req := c.NewRequest(HTTPMethod, path, nil, bytes)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)	
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
				err = fmt.Errorf("cannot get the right response")
				return
			}
			output.Status = result["status"].(string)
		} else {
			err = fmt.Errorf("cannot get the right response")
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

// UpdateSystemAdminAdministrator API operation for FortiOS updates the specified administrator account.
// Returns the index value of the administrator account and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAdminAdministrator(params *JSONSystemAdminAdministrator2, mkey string) (output *JSONUpdateSystemAdminAdministratorOutput, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/admin"
	path += "/" + mkey
	output = &JSONUpdateSystemAdminAdministratorOutput{}
	
	if params.Trusthost1 == "" {
		params.Trusthost1 = "0.0.0.0 0.0.0.0"
	}

	if params.Trusthost2 == "" {
		params.Trusthost2 = "0.0.0.0 0.0.0.0"
	}

	if params.Trusthost3 == "" {
		params.Trusthost3 = "0.0.0.0 0.0.0.0"
	}

	if params.Trusthost4 == "" {
		params.Trusthost4 = "0.0.0.0 0.0.0.0"
	}

	if params.Trusthost5 == "" {
		params.Trusthost5 = "0.0.0.0 0.0.0.0"
	}

	if params.Trusthost6 == "" {
		params.Trusthost6 = "0.0.0.0 0.0.0.0"
	}

	if params.Trusthost7 == "" {
		params.Trusthost7 = "0.0.0.0 0.0.0.0"
	}

	if params.Trusthost8 == "" {
		params.Trusthost8 = "0.0.0.0 0.0.0.0"
	}

	if params.Trusthost9 == "" {
		params.Trusthost9 = "0.0.0.0 0.0.0.0"
	}

	if params.Trusthost10 == "" {
		params.Trusthost10 = "0.0.0.0 0.0.0.0"
	}

	locJSON, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	bytes := bytes.NewBuffer(locJSON)
	req := c.NewRequest(HTTPMethod, path, nil, bytes)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
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
				err = fmt.Errorf("cannot get the right response")
				return
			}
			output.Status = result["status"].(string)
		} else {
			err = fmt.Errorf("cannot get the right response")
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

// DeleteSystemAdminAdministrator API operation for FortiOS deletes the specified administrator account.
// Returns error for service API and SDK errors.
// See the system - admin chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAdminAdministrator(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/admin"
	path += "/" + mkey

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	log.Printf("FOS-fortios response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

	if result != nil {
		if result["status"] == nil {
			err = fmt.Errorf("cannot get the right response")
			return
		}

		if result["status"] != "success" {
			err = fmt.Errorf("cannot get the right response")
			return
		}

	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}

// ReadSystemAdminAdministrator API operation for FortiOS gets the administrator account
// with the specified index value.
// Returns the requested administrator account value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAdminAdministrator(mkey string) (output *JSONSystemAdminAdministrator2, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/admin"
	path += "/" + mkey

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	log.Printf("FOS-fortios reading response: %s", string(body))

	output = &JSONSystemAdminAdministrator2{}
	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

	if result != nil {
		if result["status"] == nil {
			err = fmt.Errorf("cannot get the right response")
			return
		}
		
		if result["status"] != "success" {
			err = fmt.Errorf("cannot get the right response")
			return
		}

		mapTmp := (result["results"].([]interface{}))[0].(map[string]interface{})

		if mapTmp == nil {
			return
		}

		if mapTmp["name"] != nil {
			output.Name = mapTmp["name"].(string)
		}
		// if mapTmp["password"] != nil {
		// 	output.Password = mapTmp["password"].(string)
		// }
		if mapTmp["trusthost1"] != nil {
			output.Trusthost1 = mapTmp["trusthost1"].(string)
		}
		if mapTmp["trusthost2"] != nil {
			output.Trusthost2 = mapTmp["trusthost2"].(string)
		}
		if mapTmp["trusthost3"] != nil {
			output.Trusthost3 = mapTmp["trusthost3"].(string)
		}
		if mapTmp["trusthost4"] != nil {
			output.Trusthost4 = mapTmp["trusthost4"].(string)
		}
		if mapTmp["trusthost5"] != nil {
			output.Trusthost5 = mapTmp["trusthost5"].(string)
		}
		if mapTmp["trusthost6"] != nil {
			output.Trusthost6 = mapTmp["trusthost6"].(string)
		}
		if mapTmp["trusthost7"] != nil {
			output.Trusthost7 = mapTmp["trusthost7"].(string)
		}
		if mapTmp["trusthost8"] != nil {
			output.Trusthost8 = mapTmp["trusthost8"].(string)
		}
		if mapTmp["trusthost9"] != nil {
			output.Trusthost9 = mapTmp["trusthost9"].(string)
		}
		if mapTmp["trusthost10"] != nil {
			output.Trusthost10 = mapTmp["trusthost10"].(string)
		}
		if mapTmp["accprofile"] != nil {
			output.Accprofile = mapTmp["accprofile"].(string)
		}
		if mapTmp["comments"] != nil {
			output.Comments = mapTmp["comments"].(string)
		}

	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}

