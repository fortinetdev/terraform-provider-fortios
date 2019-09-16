package fortimngclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

/*
// store tmp data
var Tmp = map[string][]string{}
var Tmp1 = map[string][]string{}
*/

// Request represents a JSON-RPC request sent by a client.
type Request struct {
	Id      uint64         `json:"id"`
	Method  string         `json:"method"`
	Session string         `json:"session"`
	Params  [1]interface{} `json:"params"`
}

type FortiMngClient struct {
	Ipaddr string
	User   string
	Passwd string
	Debug  string
}

func NewClient(ip, user, passwd string) *FortiMngClient {

	d := os.Getenv("TRACEDEBUG")

	return &FortiMngClient{
		Ipaddr: ip,
		User:   user,
		Passwd: passwd,
		Debug:  d,
	}
}

func (c *FortiMngClient) Execute(req *Request) (result map[string]interface{}, err error) {
	// send
	j, _ := json.Marshal(req)

	if c.Debug == "ON" {
		log.Printf("[TRACEDEBUG] +++++++++++++++ req = %s", j)
	}

	httpResp, err := http.Post("http://"+c.Ipaddr+"/jsonrpc", "application/json", bytes.NewBuffer(j))
	if err != nil {
		err = fmt.Errorf("Login failed: %s", err)
		return
	}
	defer httpResp.Body.Close()

	// check response
	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}

	if c.Debug == "ON" {
		log.Printf("[TRACEDEBUG] +++++++++++++++ result = %s", body)
	}

	result = map[string]interface{}{}
	json.Unmarshal([]byte(string(body)), &result)

	if result != nil {
		if id := uint64(result["id"].(float64)); id != req.Id {
			err = fmt.Errorf("id not match, should be 1, but is %d", id)
			return
		}

		if result["result"] != nil {
			status := (result["result"].([]interface{}))[0].(map[string]interface{})["status"].(map[string]interface{})

			c := uint(status["code"].(float64))
			m := status["message"].(string)
			if c != 0 || m != "OK" {
				err = fmt.Errorf("status not right: code is %d, message is %s", c, m)
				return
			}
		} else {
			err = fmt.Errorf("can't get response status: %s", err)
			return
		}
	}

	return
}

func (c *FortiMngClient) Login() (session string, err error) {
	params := map[string]interface{}{
		"data": map[string]string{
			"user":   c.User,
			"passwd": c.Passwd,
		},
		"url": "/sys/login/user",
	}

	req := &Request{
		Id:     1,
		Method: "exec",
		Params: [1]interface{}{params},
	}

	result, err := c.Execute(req)
	if err != nil {
		return "", fmt.Errorf("login failed:%s", err)
	}

	session = result["session"].(string)
	return
}

func (c *FortiMngClient) Logout(s string) (err error) {
	params := map[string]interface{}{
		"url": "/sys/logout",
	}

	req := &Request{
		Id:      1,
		Method:  "exec",
		Params:  [1]interface{}{params},
		Session: s,
	}

	_, err = c.Execute(req)
	if err != nil {
		err = fmt.Errorf("logout failed:%s", err)
		return
	}
	return
}

func (c *FortiMngClient) Do(method string, params map[string]interface{}) (output map[string]interface{}, err error) {
	session, err := c.Login()
	if err != nil {
		return nil, fmt.Errorf("Executing failed", err)
	}
	defer c.Logout(session)

	req := &Request{
		Id:      1,
		Method:  method,
		Params:  [1]interface{}{params},
		Session: session,
	}

	output, err = c.Execute(req)
	return
}

func (f *FortiMngClient) AccessRight2Str(ar int) (s string) {
	switch ar {
	case 0:
		s = "none"
	case 1:
		s = "read"
	case 2:
		s = "read-write"
	default:
		log.Printf("[op2Str][Warning] not support number")
	}

	return
}

func (f *FortiMngClient) RpcPermit2Str(i int) (s string) {
	switch i {
	case 0:
		s = "read-write"
	case 1:
		s = "none"
	case 2:
		s = "read"
	default:
		log.Printf("[op2Str][Warning] not support number")
	}

	return
}

func (f *FortiMngClient) UserType2Str(ut int) (s string) {
	switch ut {
	case 0:
		s = "local"
	case 1:
		s = "radius"
	case 2:
		s = "ldap"
	case 3:
		s = "tacacs-plus"
	case 4:
		s = "pki-auth"
	case 5:
		s = "group"
	default:
		log.Printf("[UserType2Str][Warning] not support number")
	}

	return
}

func (f *FortiMngClient) PolicyAction2Str(pa int) (s string) {
	switch pa {
	case 0:
		s = "deny"
	case 1:
		s = "accept"
	case 2:
		s = "ipsec"
	default:
		log.Printf("[PolicyAction2Str][Warning] not support number")
	}

	return
}

func (f *FortiMngClient) InterfaceArray2StrArray(intf []interface{}) (s []string) {
	s = make([]string, len(intf))

	for i := range intf {
		s[i] = intf[i].(string)
	}

	return
}

func (f *FortiMngClient) IntfStatus2Str(ts int) (s string) {
	switch ts {
	case 0:
		s = "down"
	case 16:
		s = "up"
	default:
		log.Printf("[IntfStatus2Str][Warning] not support number")
	}

	return
}

/*
func (f *FortiMngClient) IntfServiceAccess2Str(isa int) (s string) {
	switch isa {
	case 1:
		s = "fgtupdates"
	case 4:
		s = "webfilter"
	default:
		log.Printf("[IntfServiceAccess2Str][Warning] not support number")
	}

	return
}
*/

func (f *FortiMngClient) AllowAccess2int(aa []string) int {
	allow_access := map[string]int{
		"ping":   1,
		"https":  2,
		"ssh":    4,
		"snmp":   8,
		"telnet": 16,
		"http":   32,
		"web":    64,
	}

	sum := 0
	for i := 0; i < len(aa); i++ {
		sum += allow_access[aa[i]]
	}

	return sum
}

func (f *FortiMngClient) ServiceAccess2int(sa []string) int {
	service_access := map[string]int{
		"fgtupdates": 1,
		"webfilter":  4,
	}

	sum := 0
	for i := 0; i < len(sa); i++ {
		sum += service_access[sa[i]]
	}

	return sum
}

func (f *FortiMngClient) FirewallObjectAddrType2Str(fbj int) (s string) {
	switch fbj {
	case 0:
		s = "ipmask"
	case 1:
		s = "iprange"
	case 2:
		s = "fqdn"
	case 3:
		s = "wildcard"
	case 4:
		s = "geography"
	case 6:
		s = "wildcard fqdn"

	default:
		log.Printf("[FirewallObjectAddrType2Str][Warning] not support number")
	}

	return
}

func (f *FortiMngClient) FirewallObjectProtocol2Str(c int) (s string) {
	switch c {
	case 1:
		s = "ICMP"
	case 2:
		s = "IP"
	case 5:
		s = "TCP/UDP/SCTP"
	case 6:
		s = "ICMP6"

	default:
		log.Printf("[FirewallObjectProtocol2Str][Warning] not support number")
	}

	return
}

func (f *FortiMngClient) ControlSwitch2Str(c int) (s string) {
	switch c {
	case 0:
		s = "disable"
	case 1:
		s = "enable"

	default:
		log.Printf("[ControlSwitch2Str][Warning] not support number")
	}

	return
}

func (f *FortiMngClient) FirewallObjectIppoolType2Str(c int) (s string) {
	switch c {
	case 0:
		s = "overload"
	case 1:
		s = "one-to-one"

	default:
		log.Printf("[FirewallObjectIppoolType2Str][Warning] not support number")
	}

	return
}

func (f *FortiMngClient) FirewallObjectVipType2Str(c int) (s string) {
	switch c {
	case 0:
		s = "static-nat"
	case 4:
		s = "dns-translation"
	case 5:
		s = "fqdn"

	default:
		log.Printf("[FirewallObjectVipType2Str][Warning] not support number")
	}

	return
}
func (f *FortiMngClient) ScriptTarget2Str(c int) (s string) {
	switch c {
	case 0:
		s = "device_database"
	case 1:
		s = "remote_device"
	case 2:
		s = "adom_database"

	default:
		log.Printf("[ScriptTarget2Str][Warning] not support number")
	}

	return
}

func (f *FortiMngClient) LogTraffic2Str(c int) (s string) {
	switch c {
	case 0:
		s = "disable"
	case 2:
		s = "all"
	case 3:
		s = "utm"
	default:
		log.Printf("[LogTraffic2Str][Warning] not support number")
	}

	return
}

func (f *FortiMngClient) ProfileType2Str(c int) (s string) {
	switch c {
	case 0:
		s = "single"
	case 1:
		s = "group"
	default:
		log.Printf("[ProfileType2Str][Warning] not support number")
	}

	return
}

func (f *FortiMngClient) RestrictedPrds2Str(c int) (s string) {
	switch c {
	case 1:
		s = "FortiGate"
	case 2:
		s = "FortiCarrier"
	default:
		log.Printf("[RestrictedPrds2Str][Warning] not support number")
	}

	return
}

func (f *FortiMngClient) AdomMode2Str(c int) (s string) {
	switch c {
	case 1:
		s = "normal"
	case 2:
		s = "advanced"
	default:
		log.Printf("[AdomMode2Str][Warning] not support number")
	}

	return
}

func (f *FortiMngClient) TimeZone2Str(c int) (s string) {
	if c < 0 {
		log.Printf("[TimeZone2Str][Warning] not support number")
	}

	if c >= 0 && c < 10 {
		s = "0" + strconv.Itoa(c)
	} else {
		s = strconv.Itoa(c)
	}

	return
}

/*
func (f *FortiMngClient) SetTmp(k string, v []string) {
	Tmp[k] = v
}

func (f *FortiMngClient) SetTmp1(k string, v []string) {
	Tmp1[k] = v
}

func (f *FortiMngClient) Tmp(k string) []string {
	return Tmp[k]
}

func (f *FortiMngClient) Tmp1(k string) []string {
	return Tmp1[k]
}
*/

func (f *FortiMngClient) Trace(s string) func() {
	if f.Debug == "ON" {
		log.Printf("[TRACEDEBUG] -> Enter %s", s)
		return func() { log.Printf("[TRACEDEBUG]    Leave %s <--", s) }
	}

	return func() {}
}
