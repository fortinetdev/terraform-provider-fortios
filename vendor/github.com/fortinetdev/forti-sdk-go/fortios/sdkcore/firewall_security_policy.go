package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// JSONFirewallSecurityPolicy contains the parameters for Create and Update API function
type JSONFirewallSecurityPolicy struct {
	Name                   string                     `json:"name"`
	Srcintf                MultValues                 `json:"srcintf"`
	Dstintf                MultValues                 `json:"dstintf"`
	Srcaddr                MultValues                 `json:"srcaddr"`
	Dstaddr                MultValues                 `json:"dstaddr"`
	InternetService        string                     `json:"internet-service,omitempty"`
	InternetServiceID      PolicyInternetIDMultValues `json:"internet-service-id,omitempty"`
	InternetServiceSrc     string                     `json:"internet-service-src,omitempty"`
	InternetServiceSrcID   PolicyInternetIDMultValues `json:"internet-service-src-id,omitempty"`
	Users                  MultValues                 `json:"users,omitempty"`
	Status                 string                     `json:"status,omitempty"`
	Action                 string                     `json:"action"`
	Schedule               string                     `json:"schedule"`
	Service                MultValues                 `json:"service"`
	UtmStatus              string                     `json:"utm-status,omitempty"`
	Logtraffic             string                     `json:"logtraffic,omitempty"`
	LogtrafficStart        string                     `json:"logtraffic-start,omitempty"`
	CapturePacket          string                     `json:"capture-packet,omitempty"`
	Ippool                 string                     `json:"ippool,omitempty"`
	Poolname               MultValues                 `json:"poolname,omitempty"`
	Groups                 MultValues                 `json:"groups,omitempty"`
	Devices                MultValues                 `json:"devices,omitempty"`
	Comments               string                     `json:"comments"`
	AvProfile              string                     `json:"av-profile,omitempty"`
	WebfilterProfile       string                     `json:"webfilter-profile,omitempty"`
	DnsfilterProfile       string                     `json:"dnsfilter-profile,omitempty"`
	IpsSensor              string                     `json:"ips-sensor,omitempty"`
	ApplicationList        string                     `json:"application-list,omitempty"`
	SslSSHProfile          string                     `json:"ssl-ssh-profile,omitempty"`
	Nat                    string                     `json:"nat,omitempty"`
	ProfileProtocolOptions string                     `json:"profile-protocol-options,omitempty"`
}

// JSONCreateFirewallSecurityPolicyOutput contains the output results for Create API function
type JSONCreateFirewallSecurityPolicyOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       float64 `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// JSONUpdateFirewallSecurityPolicyOutput contains the output results for Update API function
// Attention: The RESTful API changed the Mkey type from float64 in CREATE to string in UPDATE!
type JSONUpdateFirewallSecurityPolicyOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// PolicyInternetIDMultValue contains the output results for Read API function
type PolicyInternetIDMultValue struct {
	ID float64 `json:"id"`
}

// PolicyInternetIDMultValues contains the output results for Read API function
type PolicyInternetIDMultValues []PolicyInternetIDMultValue

// ExpandPolicyInternetIDList extracts ID value from result and put them into a string array,
// and return the string array
func ExpandPolicyInternetIDList(members []PolicyInternetIDMultValue) []float64 {
	vs := make([]float64, 0, len(members))
	for _, v := range members {
		c := v.ID
		vs = append(vs, c)
	}
	return vs
}

// CreateFirewallSecurityPolicy API operation for FortiOS creates a new firewall policy.
// Returns the index value of the firewall policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallSecurityPolicy(params *JSONFirewallSecurityPolicy) (output *JSONCreateFirewallSecurityPolicyOutput, err error) {
	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/policy"
	output = &JSONCreateFirewallSecurityPolicyOutput{}
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
	req.HTTPResponse.Body.Close() //#

	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	err = fortiAPIErrorFormat(result, string(body))

	if err == nil {
		if result["vdom"] != nil {
			output.Vdom = result["vdom"].(string)
		}
		if result["mkey"] != nil {
			output.Mkey = result["mkey"].(float64)
		}
	}

	return
}

// UpdateFirewallSecurityPolicy API operation for FortiOS updates the specified firewall policy.
// Returns the index value of the firewall policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallSecurityPolicy(params *JSONFirewallSecurityPolicy, mkey string) (output *JSONUpdateFirewallSecurityPolicyOutput, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/policy"
	path += "/" + mkey
	output = &JSONUpdateFirewallSecurityPolicyOutput{}
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
	req.HTTPResponse.Body.Close() //#

	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}
	log.Printf("FOS-fortios response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

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

// DeleteFirewallSecurityPolicy API operation for FortiOS deletes the specified firewall policy.
// Returns error for service API and SDK errors.
// See the firewall - policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallSecurityPolicy(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/policy"
	path += "/" + mkey

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.Send()
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %s", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close() //#

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

// ReadFirewallSecurityPolicy API operation for FortiOS gets the firewall policy
// with the specified index value.
// Returns the requested firewall policy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallSecurityPolicy(mkey string) (output *JSONFirewallSecurityPolicy, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/policy"
	path += "/" + mkey

	output = &JSONFirewallSecurityPolicy{}

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.Send()
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %s", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close() //#

	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}
	log.Printf("FOS-fortios reading response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	if fortiAPIHttpStatus404Checking(result) == true {
		output = nil
		return
	}

	err = fortiAPIErrorFormat(result, string(body))

	if err == nil {
		mapTmp := (result["results"].([]interface{}))[0].(map[string]interface{})

		if mapTmp == nil {
			err = fmt.Errorf("cannot get the results from the response")
			return
		}

		if mapTmp["name"] != nil {
			output.Name = mapTmp["name"].(string)
		}
		if mapTmp["srcintf"] != nil {
			member := mapTmp["srcintf"].([]interface{})

			var members []MultValue
			for _, v := range member {
				c := v.(map[string]interface{})

				members = append(members,
					MultValue{
						Name: c["name"].(string),
					})
			}
			output.Srcintf = members
		}
		if mapTmp["dstintf"] != nil {
			member := mapTmp["dstintf"].([]interface{})

			var members []MultValue
			for _, v := range member {
				c := v.(map[string]interface{})

				members = append(members,
					MultValue{
						Name: c["name"].(string),
					})
			}
			output.Dstintf = members
		}
		if mapTmp["srcaddr"] != nil {
			member := mapTmp["srcaddr"].([]interface{})

			var members []MultValue
			for _, v := range member {
				c := v.(map[string]interface{})

				members = append(members,
					MultValue{
						Name: c["name"].(string),
					})
			}
			output.Srcaddr = members
		}
		if mapTmp["dstaddr"] != nil {
			member := mapTmp["dstaddr"].([]interface{})

			var members []MultValue
			for _, v := range member {
				c := v.(map[string]interface{})

				members = append(members,
					MultValue{
						Name: c["name"].(string),
					})
			}
			output.Dstaddr = members
		}
		if mapTmp["internet-service"] != nil {
			output.InternetService = mapTmp["internet-service"].(string)
		}
		if mapTmp["internet-service-id"] != nil {
			member := mapTmp["internet-service-id"].([]interface{})

			var members []PolicyInternetIDMultValue
			for _, v := range member {
				c := v.(map[string]interface{})

				members = append(members,
					PolicyInternetIDMultValue{
						ID: c["id"].(float64),
					})
			}
			output.InternetServiceID = members
		}
		if mapTmp["internet-service-src"] != nil {
			output.InternetServiceSrc = mapTmp["internet-service-src"].(string)
		}
		if mapTmp["internet-service-src-id"] != nil {
			member := mapTmp["internet-service-src-id"].([]interface{})

			var members []PolicyInternetIDMultValue
			for _, v := range member {
				c := v.(map[string]interface{})

				members = append(members,
					PolicyInternetIDMultValue{
						ID: c["id"].(float64),
					})
			}
			output.InternetServiceSrcID = members
		}

		if mapTmp["users"] != nil {
			member := mapTmp["users"].([]interface{})

			var members []MultValue
			for _, v := range member {
				c := v.(map[string]interface{})

				members = append(members,
					MultValue{
						Name: c["name"].(string),
					})
			}
			output.Users = members
		}
		if mapTmp["status"] != nil {
			output.Status = mapTmp["status"].(string)
		}
		if mapTmp["action"] != nil {
			output.Action = mapTmp["action"].(string)
		}
		if mapTmp["schedule"] != nil {
			output.Schedule = mapTmp["schedule"].(string)
		}
		if mapTmp["service"] != nil {
			member := mapTmp["service"].([]interface{})

			var members []MultValue
			for _, v := range member {
				c := v.(map[string]interface{})

				members = append(members,
					MultValue{
						Name: c["name"].(string),
					})
			}
			output.Service = members
		}
		if mapTmp["utm-status"] != nil {
			output.UtmStatus = mapTmp["utm-status"].(string)
		}
		if mapTmp["logtraffic"] != nil {
			output.Logtraffic = mapTmp["logtraffic"].(string)
		}
		if mapTmp["logtraffic-start"] != nil {
			output.LogtrafficStart = mapTmp["logtraffic-start"].(string)
		}
		if mapTmp["capture-packet"] != nil {
			output.CapturePacket = mapTmp["capture-packet"].(string)
		}
		if mapTmp["ippool"] != nil {
			output.Ippool = mapTmp["ippool"].(string)
		}
		if mapTmp["poolname"] != nil {
			member := mapTmp["poolname"].([]interface{})

			var members []MultValue
			for _, v := range member {
				c := v.(map[string]interface{})

				members = append(members,
					MultValue{
						Name: c["name"].(string),
					})
			}
			output.Poolname = members
		}
		if mapTmp["groups"] != nil {
			member := mapTmp["groups"].([]interface{})

			var members []MultValue
			for _, v := range member {
				c := v.(map[string]interface{})

				members = append(members,
					MultValue{
						Name: c["name"].(string),
					})
			}
			output.Groups = members
		}
		if mapTmp["devices"] != nil {
			member := mapTmp["devices"].([]interface{})

			var members []MultValue
			for _, v := range member {
				c := v.(map[string]interface{})

				members = append(members,
					MultValue{
						Name: c["name"].(string),
					})
			}
			output.Devices = members
		}
		if mapTmp["comments"] != nil {
			output.Comments = mapTmp["comments"].(string)
		}
		if mapTmp["av-profile"] != nil {
			output.AvProfile = mapTmp["av-profile"].(string)
		}
		if mapTmp["webfilter-profile"] != nil {
			output.WebfilterProfile = mapTmp["webfilter-profile"].(string)
		}
		if mapTmp["dnsfilter-profile"] != nil {
			output.DnsfilterProfile = mapTmp["dnsfilter-profile"].(string)
		}
		if mapTmp["ips-sensor"] != nil {
			output.IpsSensor = mapTmp["ips-sensor"].(string)
		}
		if mapTmp["application-list"] != nil {
			output.ApplicationList = mapTmp["application-list"].(string)
		}
		if mapTmp["ssl-ssh-profile"] != nil {
			output.SslSSHProfile = mapTmp["ssl-ssh-profile"].(string)
		}
		if mapTmp["nat"] != nil {
			output.Nat = mapTmp["nat"].(string)
		}
		if mapTmp["profile-protocol-options"] != nil {
			output.ProfileProtocolOptions = mapTmp["profile-protocol-options"].(string)
		}

	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}
