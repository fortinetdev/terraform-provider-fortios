package forticlient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/fgtdev/fortios-sdk-go/util"
)

// CreateUpdateFirewallSecurityPolicySeq API operation for FortiOS alters the specified firewall policy sequence.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) CreateUpdateFirewallSecurityPolicySeq(srcId, dstId int, alterPos string) (err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/policy"
	path += "/" + strconv.Itoa(srcId)

	e, req := c.NewRequest(HTTPMethod, path, nil, nil)
	req.FillUrlParams(dstId, alterPos)
	if e != nil {
		err = fmt.Errorf("new request error %s", e)
		return
	}
	
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

	if result != nil {
		if result["status"] == nil {
			err = fmt.Errorf("cannot get status from the response")
			return
		}

		if result["status"] != "success" {
			if result["error"] != nil {
				err = fmt.Errorf("status is %s and error no is %.0f", result["status"], result["error"])
			} else {
				err = fmt.Errorf("status is %s and error no is not found", result["status"])
			}

			if result["http_status"] != nil {
				err = fmt.Errorf("%s, details: %s", err, util.HttpStatus2Str(int(result["http_status"].(float64))))
			} else {
				err = fmt.Errorf("%s, and http_status no is not found", err)
			}

			return
		}

	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

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
