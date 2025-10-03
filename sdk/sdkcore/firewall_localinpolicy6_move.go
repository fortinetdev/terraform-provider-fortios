package forticlient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

// CreateUpdateFirewallLocalinpolicy6Move API operation for FortiOS moves the specified item.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) CreateUpdateFirewallLocalinpolicy6Move(srcId, dstId, mv, vdomparam string) (err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/local-in-policy6"
	path += "/" + srcId

	specialparams := "action=move&"
	specialparams += mv
	specialparams += "="
	specialparams += dstId

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.SendWithSpecialParams(specialparams, vdomparam)
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %s", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close()

	if err != nil || body == nil || req.HTTPResponse.StatusCode != 200 {
		err = fmt.Errorf("cannot get response body, status %d,  %s", req.HTTPResponse.StatusCode, err)
		return
	}
	log.Printf("FOS-fortios CREATE response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	err = fortiAPIErrorFormat(result, string(body))

	return
}

// JSONMoveFirewallLocalinpolicy6Item contains the necessary parameters for each item
type JSONMoveFirewallLocalinpolicy6Item struct {
	Policyid string `json:"policyid"`
}

// GetFirewallLocalinpolicy6List API operation for FortiOS gets the list
// Returns the requested API user value when the request executes successfully.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) GetFirewallLocalinpolicy6List(vdomparam string) (out []JSONMoveFirewallLocalinpolicy6Item, err error) {

	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/local-in-policy6/"

	specialparams := "format=policyid"

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.SendWithSpecialParams(specialparams, vdomparam)
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %s", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close()

	if err != nil || body == nil || req.HTTPResponse.StatusCode != 200 {
		err = fmt.Errorf("cannot get response body, status %d, %s", req.HTTPResponse.StatusCode, err)
		return
	}
	log.Printf("FOS-fortios GET response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

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

		var members []JSONMoveFirewallLocalinpolicy6Item
		for _, v := range mapTmp {
			c := v.(map[string]interface{})

			members = append(members,
				JSONMoveFirewallLocalinpolicy6Item{
					Policyid: strconv.Itoa(int(c["policyid"].(float64))),
				})
		}

		out = members
	}

	return
}
