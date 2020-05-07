package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// JSONLogFortiAnalyzerSetting contains the parameters for Create and Update API function
type JSONLogFortiAnalyzerSetting struct {
	Status        string `json:"status"`
	Server        string `json:"server,omitempty"`
	SourceIP      string `json:"source-ip,omitempty"`
	UploadOption  string `json:"upload-option,omitempty"`
	Reliable      string `json:"reliable,omitempty"`
	HmacAlgorithm string `json:"hmac-algorithm,omitempty"`
	EncAlgorithm  string `json:"enc-algorithm,omitempty"`
}

// JSONCreateLogFortiAnalyzerSettingOutput contains the output results for Create API function
type JSONCreateLogFortiAnalyzerSettingOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// JSONUpdateLogFortiAnalyzerSettingOutput contains the output results for Update API function
// Attention: Considering scalability, the previous structure and the current structure may change differently
type JSONUpdateLogFortiAnalyzerSettingOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// CreateLogFortiAnalyzerSetting API operation for FortiOS creates a new FortiAnalyzer log management device.
// Returns the index value of the FortiAnalyzer log management device and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log - fortianalyzer setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateLogFortiAnalyzerSetting(params *JSONLogFortiAnalyzerSetting) (output *JSONCreateLogFortiAnalyzerSettingOutput, err error) {
	HTTPMethod := "POST"
	path := "/api/v2/cmdb/log.fortianalyzer/setting"
	output = &JSONCreateLogFortiAnalyzerSettingOutput{}
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

// UpdateLogFortiAnalyzerSetting API operation for FortiOS updates the specified FortiAnalyzer log management device.
// Returns the index value of the FortiAnalyzer log management device and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log - fortianalyzer setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogFortiAnalyzerSetting(params *JSONLogFortiAnalyzerSetting, mkey string) (output *JSONUpdateLogFortiAnalyzerSettingOutput, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.fortianalyzer/setting"
	// path += "/" + mkey
	output = &JSONUpdateLogFortiAnalyzerSettingOutput{}
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

// DeleteLogFortiAnalyzerSetting API operation for FortiOS deletes the specified FortiAnalyzer log management device.
// Returns error for service API and SDK errors.
// See the log - fortianalyzer setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogFortiAnalyzerSetting(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/log.fortianalyzer/setting"
	// path += "/" + mkey

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

// ReadLogFortiAnalyzerSetting API operation for FortiOS gets the FortiAnalyzer log management device
// with the specified index value.
// Returns the requested FortiAnalyzer log management device value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log - fortianalyzer setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogFortiAnalyzerSetting(mkey string) (output *JSONLogFortiAnalyzerSetting, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.fortianalyzer/setting"
	// path += "/" + mkey

	output = &JSONLogFortiAnalyzerSetting{}

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
		mapTmp := (result["results"].(map[string]interface{}))

		if mapTmp == nil {
			err = fmt.Errorf("cannot get the results from the response")
			return
		}

		if mapTmp["status"] != nil {
			output.Status = mapTmp["status"].(string)
		}
		if mapTmp["server"] != nil {
			output.Server = mapTmp["server"].(string)
		}
		if mapTmp["source-ip"] != nil {
			output.SourceIP = mapTmp["source-ip"].(string)
		}
		if mapTmp["upload-option"] != nil {
			output.UploadOption = mapTmp["upload-option"].(string)
		}
		if mapTmp["reliable"] != nil {
			output.Reliable = mapTmp["reliable"].(string)
		}
		if mapTmp["hmac-algorithm"] != nil {
			output.HmacAlgorithm = mapTmp["hmac-algorithm"].(string)
		}
		if mapTmp["enc-algorithm"] != nil {
			output.EncAlgorithm = mapTmp["enc-algorithm"].(string)
		}

	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}
