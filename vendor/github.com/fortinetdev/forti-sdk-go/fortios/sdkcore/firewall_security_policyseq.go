package forticlient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

// CreateUpdateFirewallSecurityPolicySeq API operation for FortiOS alters the specified firewall policy sequence.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) CreateUpdateFirewallSecurityPolicySeq(srcId, dstId, alterPos string) (err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/policy"
	path += "/" + srcId

	specialparams := "action=move&"
	specialparams += alterPos
	specialparams += "="
	specialparams += dstId

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.SendWithSpecialParams(specialparams)
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

// Not suitable operation
func (c *FortiSDKClient) ReadFirewallSecurityPolicySeq() (err error) {

	return
}

// Not suitable operation
func (c *FortiSDKClient) DelFirewallSecurityPolicySeq() (err error) {

	return
}

// JSONSecurityPolicyItem contains the parameters for each Security Policy item
type JSONSecurityPolicyItem struct {
	PolicyID   string     `json:"policyid"`
	Name       string     `json:"name"`
	Action     string     `json:"action"`
}

// GetSecurityPolicyList API operation for FortiOS gets the Security Policy list
// Returns the requested API user value when the request executes successfully.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) GetSecurityPolicyList() (out []JSONSecurityPolicyItem, err error) {

	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/policy/"

	specialparams := "format=policyid|action|name"

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.SendWithSpecialParams(specialparams)
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

	if fortiAPIHttpStatus404Checking(result) == true {
		return
	}

	err = fortiAPIErrorFormat(result, string(body))

	if err == nil {
		mapTmp := result["results"].([]interface{}) //)[0].(map[string]interface{})

		if mapTmp == nil {
			err = fmt.Errorf("cannot get the results from the response")
			return
		}

		var members []JSONSecurityPolicyItem
		for _, v := range mapTmp {
			c := v.(map[string]interface{})

			members = append(members,
				JSONSecurityPolicyItem {
					PolicyID: strconv.Itoa(int(c["policyid"].(float64))),
					Name:     c["name"].(string),
					Action:   c["action"].(string),
				})
			}

		out = members
	}

	return
}