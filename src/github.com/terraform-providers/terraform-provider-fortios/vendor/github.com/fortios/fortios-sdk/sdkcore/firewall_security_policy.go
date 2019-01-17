package forticlient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"fmt"
	// "strconv"
)

//JSONFirewallSecurityPolicy contains ... need to comment completely
type JSONFirewallSecurityPolicy struct {
	Name              string                     `json:"name"`
	Srcintf           MultValues                 `json:"srcintf"`
	Dstintf           MultValues                 `json:"dstintf"`
	Srcaddr           MultValues                 `json:"srcaddr"`
	Dstaddr           MultValues                 `json:"dstaddr"`
	InternetService   string                     `json:"internet-service"`
	InternetServiceID PolicyInternetIDMultValues `json:"internet-service-id"`
	Action            string                     `json:"action"`
	Schedule          string                     `json:"schedule"`
	Service           MultValues                 `json:"service"`
	UtmStatus         string                     `json:"utm-status"`
	Logtraffic        string                     `json:"logtraffic"`
	LogtrafficStart   string                     `json:"logtraffic-start"`
	CapturePacket     string                     `json:"capture-packet"`
	Ippool            string                     `json:"ippool"`
	Poolname          MultValues                 `json:"poolname"`
	Groups            MultValues                 `json:"groups"`
	Devices           MultValues                 `json:"devices"`
	Comments          string                     `json:"comments"`
	AvProfile         string                     `json:"av-profile"`
	WebfilterProfile  string                     `json:"webfilter-profile"`
	DnsfilterProfile  string                     `json:"dnsfilter-profile"`
	IpsSensor         string                     `json:"ips-sensor"`
	ApplicationList   string                     `json:"application-list"`
	SslSSHProfile     string                     `json:"ssl-ssh-profile"`
	Nat               string                     `json:"nat"`
}

//JSONCreateFirewallSecurityPolicyOutput contains ... need to comment completely
type JSONCreateFirewallSecurityPolicyOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       float64 `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

//JSONUpdateFirewallSecurityPolicyOutput contains ... need to comment completely
//Attention: The RESTful API changed the Mkey type from float64 in CREATE to string in UPDATE!
type JSONUpdateFirewallSecurityPolicyOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

//PolicyInternetIDMultValue contains ... need to comment completely
type PolicyInternetIDMultValue struct {
	ID float64 `json:"id"`
}

//PolicyInternetIDMultValues contains ... need to comment completely
type PolicyInternetIDMultValues []PolicyInternetIDMultValue

//ExpandPolicyInternetIDList contains ... need to comment completely
func ExpandPolicyInternetIDList(members []PolicyInternetIDMultValue) []float64 {
	vs := make([]float64, 0, len(members))
	for _, v := range members {
		c := v.ID
		vs = append(vs, c)
	}
	return vs
}

//CreateFirewallSecurityPolicy will send ... need to comment completely
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

//UpdateFirewallSecurityPolicy will send ... need to comment completely
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

//DeleteFirewallSecurityPolicy will send ... need to comment completely
func (c *FortiSDKClient) DeleteFirewallSecurityPolicy(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/policy"
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

//ReadFirewallSecurityPolicy will send ... need to comment completely
func (c *FortiSDKClient) ReadFirewallSecurityPolicy(mkey string) (output *JSONFirewallSecurityPolicy, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/policy"
	path += "/" + mkey

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	log.Printf("shengh.............fortios reading response: %s", string(body))

	output = &JSONFirewallSecurityPolicy{}
	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

	if result != nil {
		if result["status"] == nil {
			err = fmt.Errorf("cannot get the right response")
			return
		}

		mapTmp := (result["results"].([]interface {}))[0].(map[string]interface {})

		if mapTmp == nil {
			return
		}

		if mapTmp["name"] != nil {
			output.Name = mapTmp["name"].(string)
		}
		if mapTmp["srcintf"] != nil {
			member := mapTmp["srcintf"].([]interface {})

			var members []MultValue
			for _, v := range member {
				c := v.(map[string]interface {})

				members = append(members,
					MultValue{
						Name: c["name"].(string),
					})
			}
			output.Srcintf = members
		}
		if mapTmp["dstintf"] != nil {
			member := mapTmp["dstintf"].([]interface {})

			var members []MultValue
			for _, v := range member {
				c := v.(map[string]interface {})

				members = append(members,
					MultValue{
						Name: c["name"].(string),
					})
			}
			output.Dstintf = members
		}
		if mapTmp["srcaddr"] != nil {
			member := mapTmp["srcaddr"].([]interface {})

			var members []MultValue
			for _, v := range member {
				c := v.(map[string]interface {})

				members = append(members,
					MultValue{
						Name: c["name"].(string),
					})
			}
			output.Srcaddr = members
		}
		if mapTmp["dstaddr"] != nil {
			member := mapTmp["dstaddr"].([]interface {})

			var members []MultValue
			for _, v := range member {
				c := v.(map[string]interface {})

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
			member := mapTmp["internet-service-id"].([]interface {})

			var members []PolicyInternetIDMultValue
			for _, v := range member {
				c := v.(map[string]interface {})

				members = append(members,
					PolicyInternetIDMultValue{
						ID: c["id"].(float64),
					})
			}
			output.InternetServiceID = members
		}
		if mapTmp["action"] != nil {
			output.Action = mapTmp["action"].(string)
		}
		if mapTmp["schedule"] != nil {
			output.Schedule = mapTmp["schedule"].(string)
		}
		if mapTmp["service"] != nil {
			member := mapTmp["service"].([]interface {})

			var members []MultValue
			for _, v := range member {
				c := v.(map[string]interface {})

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
			member := mapTmp["poolname"].([]interface {})

			var members []MultValue
			for _, v := range member {
				c := v.(map[string]interface {})

				members = append(members,
					MultValue{
						Name: c["name"].(string),
					})
			}
			output.Poolname = members
		}
		if mapTmp["groups"] != nil {
			member := mapTmp["groups"].([]interface {})

			var members []MultValue
			for _, v := range member {
				c := v.(map[string]interface {})

				members = append(members,
					MultValue{
						Name: c["name"].(string),
					})
			}
			output.Groups = members
		}
		if mapTmp["devices"] != nil {
			member := mapTmp["devices"].([]interface {})

			var members []MultValue
			for _, v := range member {
				c := v.(map[string]interface {})

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

	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}
