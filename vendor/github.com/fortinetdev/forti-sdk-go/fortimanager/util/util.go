package util

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
)

// AccessRight2Str converts access right to string
func AccessRight2Str(ar int) (s string) {
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

// RpcPermit2Str converts rpc permit to string
func RpcPermit2Str(i int) (s string) {
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

// UserType2Str converts user type to string
func UserType2Str(ut int) (s string) {
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

// PolicyAction2Str converts policy action to string
func PolicyAction2Str(pa int) (s string) {
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

// InterfaceArray2StrArray converts interface array to string array
func InterfaceArray2StrArray(intf []interface{}) (s []string) {
	s = make([]string, len(intf))

	for i := range intf {
		s[i] = intf[i].(string)
	}

	return
}

// IntfStatus2Str converts interface status to string
func IntfStatus2Str(ts int) (s string) {
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
func IntfServiceAccess2Str(isa int) (s string) {
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

// AllowAccess2int converts allow access to int
func AllowAccess2int(aa []string) int {
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

// ServiceAccess2int converts service access to int
func ServiceAccess2int(sa []string) int {
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

// FirewallObjectAddrType2Str converts firewall obj address type to string
func FirewallObjectAddrType2Str(fbj int) (s string) {
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

// FirewallObjectProtocol2Str converts firewall obj protocol to string
func FirewallObjectProtocol2Str(c int) (s string) {
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

// ControlSwitch2Str converts control switch to string
func ControlSwitch2Str(c int) (s string) {
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

// FirewallObjectIppoolType2Str converts firewall obj ippool type to string
func FirewallObjectIppoolType2Str(c int) (s string) {
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

// FirewallObjectVipType2Str converts firewall obj vip type to string
func FirewallObjectVipType2Str(c int) (s string) {
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

// ScriptTarget2Str converts script target to string
func ScriptTarget2Str(c int) (s string) {
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

// LogTraffic2Str converts log traffic to string
func LogTraffic2Str(c int) (s string) {
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

// ProfileType2Str converts profile type to string
func ProfileType2Str(c int) (s string) {
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

// RestrictedPrds2Str converts restricted prds type to string
func RestrictedPrds2Str(c int) (s string) {
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

// AdomMode2Str converts adom mode to string
func AdomMode2Str(c int) (s string) {
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

// TimeZone2Str converts time zone to string
func TimeZone2Str(c int) (s string) {
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

// InspectionMode2Str converts inspection mode to string
func InspectionMode2Str(c int) (s string) {
	switch c {
	case 0:
		s = "proxy"
	case 1:
		s = "flow"
	default:
		log.Printf("[InspectionMode2Str][Warning] not support number")
	}

	return
}

// ValidateStringIn checks the validation of input string
func ValidateStringIn(vals ...string) schema.SchemaValidateFunc {
	return func(v interface{}, k string) (ws []string, errors []error) {
		value := v.(string)
		ok := false
		for i := range vals {
			if vals[i] == value {
				ok = true
				break
			}
		}

		if !ok {
			errors = append(errors, fmt.Errorf("%q (%q) not in %#v", k, value, vals))
		}

		return
	}
}
