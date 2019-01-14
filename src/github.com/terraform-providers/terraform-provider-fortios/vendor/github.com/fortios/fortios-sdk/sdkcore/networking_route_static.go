package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

//JSONNetworkingRouteStatic contains ... need to comment completely
type JSONNetworkingRouteStatic struct {
	Dst       string `json:"dst"`
	Gateway   string `json:"gateway"`
	Blackhole string `json:"blackhole"`
	Distance  string `json:"distance"`
	Weight    string `json:"weight"`
	Priority  string `json:"priority"`
	Device    string `json:"device"`
	Comment   string `json:"comment"`
}

//JSONCreateNetworkingRouteStaticOutput contains ... need to comment completely
type JSONCreateNetworkingRouteStaticOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       float64 `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

//JSONUpdateNetworkingRouteStaticOutput contains ... need to comment completely
//Attention: The RESTful API changed the Mkey type from float64 in CREATE to string in UPDATE!
type JSONUpdateNetworkingRouteStaticOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

//CreateNetworkingRouteStatic will send ... need to comment completely
func (c *FortiSDKClient) CreateNetworkingRouteStatic(params *JSONNetworkingRouteStatic) (output *JSONCreateNetworkingRouteStaticOutput, err error) {
	HTTPMethod := "POST"
	path := "/api/v2/cmdb/router/static"
	output = &JSONCreateNetworkingRouteStaticOutput{}
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
			output.Mkey = result["mkey"].(float64)
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

//UpdateNetworkingRouteStatic will send ... need to comment completely
func (c *FortiSDKClient) UpdateNetworkingRouteStatic(params *JSONNetworkingRouteStatic, mkey string) (output *JSONUpdateNetworkingRouteStaticOutput, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/static"
	path += "/" + mkey
	output = &JSONUpdateNetworkingRouteStaticOutput{}
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

//DeleteNetworkingRouteStatic will send ... need to comment completely
func (c *FortiSDKClient) DeleteNetworkingRouteStatic(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/router/static"
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

//ReadNetworkingRouteStatic will send ... need to comment completely
func (c *FortiSDKClient) ReadNetworkingRouteStatic(mkey string) (output *JSONNetworkingRouteStatic, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/static"
	path += "/" + mkey

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	log.Printf("shengh.............fortios reading response: %s", string(body))

	output = &JSONNetworkingRouteStatic{}
	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

	if result != nil {
		if result["status"] == nil {
			err = fmt.Errorf("cannot get the right response")
			return
		}

		mapTmp := (result["results"].([]interface{}))[0].(map[string]interface{})

		if mapTmp == nil {
			return
		}

		if mapTmp["dst"] != nil {
			output.Dst = mapTmp["dst"].(string)
		}
		if mapTmp["gateway"] != nil {
			output.Gateway = mapTmp["gateway"].(string)
		}
		if mapTmp["blackhole"] != nil {
			output.Blackhole = mapTmp["blackhole"].(string)
		}
		if mapTmp["distance"] != nil {
			output.Distance = strconv.Itoa(int(mapTmp["distance"].(float64)))
		}
		if mapTmp["weight"] != nil {
			output.Weight = strconv.Itoa(int(mapTmp["weight"].(float64)))
		}
		if mapTmp["priority"] != nil {
			output.Priority = strconv.Itoa(int(mapTmp["priority"].(float64)))
		}
		if mapTmp["device"] != nil {
			output.Device = mapTmp["device"].(string)
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
