package forticlient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"fmt"
	// "strconv"
)

//JSONSystemAPIUserSetting contains ... need to comment completely
type JSONSystemAPIUserSetting struct {
	Name       string `json:"name"`
	Accprofile       string `json:"accprofile"`
	Vdom       string `json:"vdom"`
	Trusthost       APIUserMultValues `json:"trusthost"`
}

//JSONCreateSystemAPIUserSettingOutput contains ... need to comment completely
type JSONCreateSystemAPIUserSettingOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

//JSONUpdateSystemAPIUserSettingOutput contains ... need to comment completely
//Attention: The RESTful API changed the Mkey type from float64 in CREATE to string in UPDATE!
type JSONUpdateSystemAPIUserSettingOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

//APIUserMultValue contains ... need to comment completely
type APIUserMultValue struct {
	TypeF string `json:"type"`
	Ipv4TrustHost string `json:"ipv4-trusthost"`
}

//APIUserMultValues contains ... need to comment completely
type APIUserMultValues []APIUserMultValue


//CreateSystemAPIUserSetting will send ... need to comment completely
func (c *FortiSDKClient) CreateSystemAPIUserSetting(params *JSONSystemAPIUserSetting) (output *JSONCreateSystemAPIUserSettingOutput, err error) {
	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/api-user"
	output = &JSONCreateSystemAPIUserSettingOutput{}
	locJSON, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	bytes := bytes.NewBuffer(locJSON)
	req := c.NewRequest(HTTPMethod, path, nil, bytes)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)

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

//UpdateSystemAPIUserSetting will send ... need to comment completely
func (c *FortiSDKClient) UpdateSystemAPIUserSetting(params *JSONSystemAPIUserSetting, mkey string) (output *JSONUpdateSystemAPIUserSettingOutput, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/api-user"
	path += "/" + mkey
	output = &JSONUpdateSystemAPIUserSettingOutput{}
	locJSON, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	bytes := bytes.NewBuffer(locJSON)
	req := c.NewRequest(HTTPMethod, path, nil, bytes)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	log.Printf("shengh.............fortios response: %s", string(body))

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


//DeleteSystemAPIUserSetting will send ... need to comment completely
func (c *FortiSDKClient) DeleteSystemAPIUserSetting(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/api-user"
	path += "/" + mkey

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	log.Printf("shengh.............fortios response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

	if result != nil {
		if result["status"] == nil {
			err = fmt.Errorf("cannot get the right response")
			return
		}
	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}

//ReadSystemAPIUserSetting will send ... need to comment completely
func (c *FortiSDKClient) ReadSystemAPIUserSetting(mkey string) (output *JSONSystemAPIUserSetting, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/api-user"
	path += "/" + mkey

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	log.Printf("shengh.............fortios reading response: %s", string(body))

	output = &JSONSystemAPIUserSetting{}
	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

	if result != nil {
		if result["status"] == nil {
			err = fmt.Errorf("cannot get the right response")
			return
		}

		mapTmp := (result["results"].([]interface {}))[0].(map[string]interface {})

		if mapTmp == nil {
			return
		}

		if mapTmp["name"] != nil {
			output.Name = mapTmp["name"].(string)
		}
		if mapTmp["accprofile"] != nil {
			output.Accprofile = mapTmp["accprofile"].(string)
		}
		if mapTmp["vdom"] != nil {

			member := mapTmp["vdom"].([]interface{})

			output.Vdom = "root"

			for _, v := range member {
				c := v.(map[string]interface{})

				output.Vdom = c["name"].(string)

				// the break here because of the vdom in response becomes array, but vdom in request is string
				break
			}
		}

		if mapTmp["trusthost"] != nil {
			member := mapTmp["trusthost"].([]interface{})

			var members []APIUserMultValue
			for _, v := range member {
				c := v.(map[string]interface{})

				members = append(members,
					APIUserMultValue{
						TypeF: c["type"].(string),
						Ipv4TrustHost: c["ipv4-trusthost"].(string),
					})
			}
			output.Trusthost = members
		}

	} else {
		err = fmt.Errorf("cannot get the right response")
	}

	return
}
