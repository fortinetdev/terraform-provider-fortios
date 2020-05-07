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
func (c *FortiSDKClient) CreateUpdateFirewallSecurityPolicySeq(srcId, dstId int, alterPos string) (err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/policy"
	path += "/" + strconv.Itoa(srcId)

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	req.FillUrlParams(dstId, alterPos)
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

// Not suitable operation
func (c *FortiSDKClient) ReadFirewallSecurityPolicySeq() (err error) {
	return
}

// Not suitable operation
func (c *FortiSDKClient) DelFirewallSecurityPolicySeq() (err error) {
	return
}
