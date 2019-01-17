package forticlient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"fmt"
	// "strconv"
)

//JSONSystemLicenseVM contains ... need to comment completely
type JSONSystemLicenseVM struct {
	File_content       string `json:"file_content"`
}

//JSONCreateSystemLicenseVMOutput contains ... need to comment completely
type JSONCreateSystemLicenseVMOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

//JSONUpdateSystemLicenseVMOutput contains ... need to comment completely
//Attention: The RESTful API changed the Mkey type from float64 in CREATE to string in UPDATE!
type JSONUpdateSystemLicenseVMOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

//CreateSystemLicenseVM will send ... need to comment completely
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

			if output.Status != "success" {
				err = fmt.Errorf("Error License file");
			}

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

//UpdateSystemLicenseVM will send ... need to comment completely
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


//DeleteSystemLicenseVM will send ... need to comment completely
func (c *FortiSDKClient) DeleteSystemLicenseVM(mkey string) (err error) {
	// HTTPMethod := "DELETE"
	// path := "/api/v2/monitor/system/vmlicense/upload"
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

//ReadSystemLicenseVM will send ... need to comment completely
func (c *FortiSDKClient) ReadSystemLicenseVM(mkey string) (output *JSONSystemLicenseVM, err error) {
	// HTTPMethod := "GET"
	// path := "/api/v2/monitor/system/vmlicense/upload"
	// path += "/" + mkey

	// req := c.NewRequest(HTTPMethod, path, nil, nil)
	// err = req.Send()

	// body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	// log.Printf("shengh.............fortios reading response: %s", string(body))

	// output = &JSONSystemLicenseVM{}
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
	// 		return
	// 	}

	// 	if mapTmp["file_content"] != nil {
	// 		output.File_content = mapTmp["file_content"].(string)
	// 	}

	// } else {
	// 	err = fmt.Errorf("cannot get the right response")
	// 	return
	// }

	return
}
