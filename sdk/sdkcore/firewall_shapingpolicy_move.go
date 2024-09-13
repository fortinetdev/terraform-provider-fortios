package forticlient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

// CreateUpdateFirewallShapingpolicyMove API operation for FortiOS moves the specified item.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) CreateUpdateFirewallShapingpolicyMove(srcId, dstId, mv, vdomparam string) (err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/shaping-policy"
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

	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}
	log.Printf("FOS-fortios response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	err = fortiAPIErrorFormat(result, string(body))

	return
}

// JSONMoveFirewallShapingpolicyItem contains the necessary parameters for each item
type JSONMoveFirewallShapingpolicyItem struct {
	Id string `json:"id"`
}

// GetFirewallShapingpolicyList API operation for FortiOS gets the list
// Returns the requested API user value when the request executes successfully.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) GetFirewallShapingpolicyList(vdomparam string) (out []JSONMoveFirewallShapingpolicyItem, err error) {

	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/shaping-policy/"

	specialparams := "format=id"

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.SendWithSpecialParams(specialparams, vdomparam)
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %s", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close()

	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}
	log.Printf("FOS-fortios response: %s", string(body))

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

		var members []JSONMoveFirewallShapingpolicyItem
		for _, v := range mapTmp {
			c := v.(map[string]interface{})

			members = append(members,
				JSONMoveFirewallShapingpolicyItem{
					Id: strconv.Itoa(int(c["id"].(float64))),
				})
		}

		out = members
	}

	return
}
