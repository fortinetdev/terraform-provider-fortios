package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// JSONSystemAdminAdministrator contains the parameters for Create and Update API function
type JSONSystemAdminAdministrator struct {
	Name        string     `json:"name"`
	Password    string     `json:"password"`
	Trusthost1  string     `json:"trusthost1,omitempty"`
	Trusthost2  string     `json:"trusthost2,omitempty"`
	Trusthost3  string     `json:"trusthost3,omitempty"`
	Trusthost4  string     `json:"trusthost4,omitempty"`
	Trusthost5  string     `json:"trusthost5,omitempty"`
	Trusthost6  string     `json:"trusthost6,omitempty"`
	Trusthost7  string     `json:"trusthost7,omitempty"`
	Trusthost8  string     `json:"trusthost8,omitempty"`
	Trusthost9  string     `json:"trusthost9,omitempty"`
	Trusthost10 string     `json:"trusthost10,omitempty"`
	Accprofile  string     `json:"accprofile"`
	Comments    string     `json:"comments"`
	Vdom        MultValues `json:"vdom,omitempty"`
}

// JSONSystemAdminAdministrator2 contains the parameters for Create and Update API function
type JSONSystemAdminAdministrator2 struct {
	Name        string     `json:"name"`
	Trusthost1  string     `json:"trusthost1,omitempty"`
	Trusthost2  string     `json:"trusthost2,omitempty"`
	Trusthost3  string     `json:"trusthost3,omitempty"`
	Trusthost4  string     `json:"trusthost4,omitempty"`
	Trusthost5  string     `json:"trusthost5,omitempty"`
	Trusthost6  string     `json:"trusthost6,omitempty"`
	Trusthost7  string     `json:"trusthost7,omitempty"`
	Trusthost8  string     `json:"trusthost8,omitempty"`
	Trusthost9  string     `json:"trusthost9,omitempty"`
	Trusthost10 string     `json:"trusthost10,omitempty"`
	Accprofile  string     `json:"accprofile"`
	Comments    string     `json:"comments"`
	Vdom        MultValues `json:"vdom,omitempty"`
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

// UpdateSystemAdminAdministrator API operation for FortiOS updates the specified administrator account.
// Returns the index value of the administrator account and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAdminAdministrator(params *JSONSystemAdminAdministrator2, mkey string) (output *JSONUpdateSystemAdminAdministratorOutput, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/admin"
	path += "/" + EscapeURLString(mkey)
	output = &JSONUpdateSystemAdminAdministratorOutput{}
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
	log.Printf("FOS-fortios response: %s", string(body))

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

// DeleteSystemAdminAdministrator API operation for FortiOS deletes the specified administrator account.
// Returns error for service API and SDK errors.
// See the system - admin chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAdminAdministrator(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/admin"
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
	log.Printf("FOS-fortios response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

	err = fortiAPIErrorFormat(result, string(body))

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
	path += "/" + EscapeURLString(mkey)

	output = &JSONSystemAdminAdministrator2{}

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
	log.Printf("FOS-fortios reading response: %s", string(body))

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
		if mapTmp["vdom"] != nil {
			member := mapTmp["vdom"].([]interface{})

			var members []MultValue
			for _, v := range member {
				c := v.(map[string]interface{})

				members = append(members,
					MultValue{
						Name: c["name"].(string),
					})
			}
			output.Vdom = members
		}

	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}
