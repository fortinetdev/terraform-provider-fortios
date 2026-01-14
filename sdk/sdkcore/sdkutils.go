package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func createUpdate(requestInput *requestInput) (output map[string]interface{}, err error) {
	c := requestInput.fortiSDKClient
	method := requestInput.method
	path := requestInput.path
	params := requestInput.params
	vdomparam := requestInput.vdomparam
	headers := requestInput.headers

	locJSON, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	bytes := bytes.NewBuffer(locJSON)

	req := c.NewRequest(method, path, nil, bytes)
	err = req.Send3(vdomparam, headers)
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("Cannot send request: %v", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close() //#

	if err != nil || body == nil {
		err = fmt.Errorf("Cannot get response body: %v", err)
		return
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)
	err = fortiAPIErrorFormat(result, string(body))

	if err == nil {
		output = make(map[string]interface{})
		if result["vdom"] != nil {
			output["vdom"] = result["vdom"]
		}

		if result["mkey"] != nil {
			output["mkey"] = result["mkey"]
		}
	}

	return
}

func delete(requestInput *requestInput) (err error) {
	c := requestInput.fortiSDKClient
	method := requestInput.method
	path := requestInput.path
	vdomparam := requestInput.vdomparam
	headers := requestInput.headers

	req := c.NewRequest(method, path, nil, nil)
	err = req.Send3(vdomparam, headers)
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("Cannot send request: %v", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close() //#

	if err != nil || body == nil {
		err = fmt.Errorf("Cannot get response body: %v", err)
		return
	}
	log.Printf("FOS-fortios response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	err = fortiAPIErrorFormat(result, string(body))

	return
}

func read(requestInput *requestInput) (mapTmp map[string]interface{}, err error) {
	c := requestInput.fortiSDKClient
	method := requestInput.method
	path := requestInput.path
	bcomplex := requestInput.bcomplex
	vdomparam := requestInput.vdomparam
	headers := requestInput.headers

	req := c.NewRequest(method, path, nil, nil)
	err = req.Send3(vdomparam, headers)
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("Cannot send request: %v", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close() //#

	if err != nil || body == nil {
		err = fmt.Errorf("Cannot get response body: %v", err)
		return
	}
	log.Printf("FOS-fortios reading response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	if fortiAPIHttpStatus404Checking(result) == true {
		mapTmp = nil
		return
	}

	err = fortiAPIErrorFormat(result, string(body))

	if err == nil {
		if bcomplex {
			mapTmp = result["results"].(map[string]interface{})
		} else {
			mapTmp = (result["results"].([]interface{}))[0].(map[string]interface{})
		}
	}

	return
}
