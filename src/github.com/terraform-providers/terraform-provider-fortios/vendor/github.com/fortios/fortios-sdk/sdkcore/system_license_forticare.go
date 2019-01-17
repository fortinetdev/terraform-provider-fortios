package forticlient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"fmt"
	// "strconv"
)

//JSONSystemLicenseFortiCare contains ... need to comment completely
type JSONSystemLicenseFortiCare struct {
	RegistrationCode       string `json:"registration_code"`
}

//JSONCreateSystemLicenseFortiCareOutput contains ... need to comment completely
type JSONCreateSystemLicenseFortiCareOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

//JSONUpdateSystemLicenseFortiCareOutput contains ... need to comment completely
//Attention: The RESTful API changed the Mkey type from float64 in CREATE to string in UPDATE!
type JSONUpdateSystemLicenseFortiCareOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

//CreateSystemLicenseFortiCare will send ... need to comment completely
func (c *FortiSDKClient) CreateSystemLicenseFortiCare(params *JSONSystemLicenseFortiCare) (output *JSONCreateSystemLicenseFortiCareOutput, err error) {
	HTTPMethod := "POST"
	path := "/api/v2/monitor/registration/forticare/add-license"
	output = &JSONCreateSystemLicenseFortiCareOutput{}
	locJSON, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	bytes := bytes.NewBuffer(locJSON)
	req := c.NewRequest(HTTPMethod, path, nil, bytes)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	log.Printf("shengh.............fortios response(create): %s", string(body))

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


		if result["results"] == nil {
			err = fmt.Errorf("cannot get the right response")
			return			
		}
				
		mapTmp := result["results"].(map[string]interface {});

		if mapTmp == nil {
			err = fmt.Errorf("cannot get the right response")
			return
		}

		if mapTmp["forticare_error"] != nil {
			s :=  mapTmp["forticare_error"].(string)
			if s != "" {
				err = fmt.Errorf("cannot get the right response" + s)
			}
			return
		}
	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}

//UpdateSystemLicenseFortiCare will send ... need to comment completely
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
	// log.Printf("shengh.............fortios response: %s", string(body))

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


//DeleteSystemLicenseFortiCare will send ... need to comment completely
func (c *FortiSDKClient) DeleteSystemLicenseFortiCare(mkey string) (err error) {
	// HTTPMethod := "DELETE"
	// path := "/api/v2/monitor/registration/forticare/add-license"
	// path += "/" + mkey

	// req := c.NewRequest(HTTPMethod, path, nil, nil)
	// err = req.Send()

	// body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	// log.Printf("shengh.............fortios response: %s", string(body))

	// var result map[string]interface{}
	// json.Unmarshal([]byte(string(body)), &result)

	// req.HTTPResponse.Body.Close()

	// if result != nil {
	// 	if result["status"] == nil {
	// 		err = fmt.Errorf("cannot get the right response")
	// 		return
	// 	}
	// } else {
	// 	err = fmt.Errorf("cannot get the right response")
	// 	return
	// }

	return
}

//ReadSystemLicenseFortiCare will send ... need to comment completely
func (c *FortiSDKClient) ReadSystemLicenseFortiCare(mkey string) (output *JSONSystemLicenseFortiCare, err error) {
	// HTTPMethod := "GET"
	// path := "/api/v2/monitor/registration/forticare/add-license"
	// path += "/" + mkey

	// req := c.NewRequest(HTTPMethod, path, nil, nil)
	// err = req.Send()

	// body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	// log.Printf("shengh.............fortios reading response: %s", string(body))

	// output = &JSONSystemLicenseFortiCare{}
	// var result map[string]interface{}
	// json.Unmarshal([]byte(string(body)), &result)

	// req.HTTPResponse.Body.Close()

	// if result != nil {
	// 	if result["status"] == nil {
	// 		err = fmt.Errorf("cannot get the right response")
	// 		return
	// 	}

	// 	mapTmp := (result["results"].([]interface {}))[0].(map[string]interface {})

	// 	if mapTmp == nil {
	// 		err = fmt.Errorf("cannot get the right response")
	// 		return
	// 	}

	// 	// if mapTmp["registration_code"] != nil {
	// 	// 	output.RegistrationCode = mapTmp["registration_code"].(string)
	// 	// }

	// 	// if mapTmp["forticare_error"] != nil {
	// 	// 	err = fmt.Errorf("cannot get the right response" + mapTmp["forticare_error"].(string))
	// 	// 	return
	// 	// }


	// } else {
	// 	err = fmt.Errorf("cannot get the right response")
	// 	return
	// }

	return
}
