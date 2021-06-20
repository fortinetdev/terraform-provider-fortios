package fortios

import (
	"encoding/binary"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func fortiValidateEnum(v []string) schema.SchemaValidateFunc {
	return validation.StringInSlice(v, false)
}

func fortiValidateEnableDisable() schema.SchemaValidateFunc {
	return fortiValidateEnum([]string{"enable", "disable"})
}

func fortiValidateIPv4ClassnetAny(i interface{}, k string) (warnings []string, errors []error) {
	v, ok := i.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected type of %q to be string", k))
		return warnings, errors
	}

	s := strings.Split(v, " ")
	if len(s) == 2 {
		ip := s[0]
		mask := s[1]
		prefix_length, _ := net.IPMask(net.ParseIP(mask).To4()).Size()
		if mask != "0.0.0.0" && prefix_length == 0 {
			errors = append(errors, fmt.Errorf("expected %s to contain a valid IPv4 netmask, got %q", k, v))
			return warnings, errors
		}
		v = ip + "/" + strconv.Itoa(prefix_length)
	}

	ip, network, err := net.ParseCIDR(v)

	if err != nil {
		errors = append(errors, fmt.Errorf("expected %s to contain a valid IPv4 value, got %q: %v", k, v, err))
		return warnings, errors
	}

	if !ip.Equal(network.IP) {
		errors = append(errors, fmt.Errorf("expected %s to contain the value of an IPv4 network address, got %q", k, v))
		return warnings, errors
	}

	return warnings, errors
}

func fortiValidateIPv4ClassnetHost(i interface{}, k string) (warnings []string, errors []error) {
	v, ok := i.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected type of %q to be string", k))
		return warnings, errors
	}

	s := strings.Split(v, " ")
	if len(s) == 2 {
		ip := s[0]
		mask := s[1]
		prefix_length, _ := net.IPMask(net.ParseIP(mask).To4()).Size()
		if mask != "0.0.0.0" && prefix_length == 0 {
			errors = append(errors, fmt.Errorf("expected %s to contain a valid IPv4 netmask, got %q", k, v))
			return warnings, errors
		}
		v = ip + "/" + strconv.Itoa(prefix_length)
	}

	ip, network, err := net.ParseCIDR(v)

	if err != nil {
		errors = append(errors, fmt.Errorf("expected %s to contain a valid IPv4 value, got %q: %v", k, v, err))
		return warnings, errors
	}

	prefix_length, _ := network.Mask.Size()
	net_exception := false
	if prefix_length == 0 || prefix_length == 31 || prefix_length == 32 {
		net_exception = true
	}

	if ip.IsMulticast() {
		errors = append(errors, fmt.Errorf("expected %s to NOT contain the value of an IPv4 multicast address, got %q", k, v))
		return warnings, errors
	}

	if ip.IsLoopback() {
		errors = append(errors, fmt.Errorf("expected %s to NOT contain the value of an IPv4 loopback address, got %q", k, v))
		return warnings, errors
	}

	if isBroadcastAddress(ip, network) || ip.Equal(net.IPv4bcast) {
		errors = append(errors, fmt.Errorf("expected %s to NOT contain the value of an IPv4 broadcast address, got %q", k, v))
		return warnings, errors
	}

	if ip.Equal(network.IP) && !net_exception {
		errors = append(errors, fmt.Errorf("expected %s to NOT contain the value of an IPv4 network address, got %q", k, v))
		return warnings, errors
	}

	return warnings, errors
}

func fortiValidateIPv4Classnet(i interface{}, k string) (warnings []string, errors []error) {
	v, ok := i.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected type of %q to be string", k))
		return warnings, errors
	}

	s := strings.Split(v, " ")
	if len(s) == 2 {
		ip := s[0]
		mask := s[1]
		prefix_length, _ := net.IPMask(net.ParseIP(mask).To4()).Size()
		if mask != "0.0.0.0" && prefix_length == 0 {
			errors = append(errors, fmt.Errorf("expected %s to contain a valid IPv4 netmask, got %q", k, v))
			return warnings, errors
		}
		v = ip + "/" + strconv.Itoa(prefix_length)
	}

	ip, network, err := net.ParseCIDR(v)

	if err != nil {
		errors = append(errors, fmt.Errorf("expected %s to contain a valid IPv4 value, got %q: %v", k, v, err))
		return warnings, errors
	}

	prefix_length, _ := network.Mask.Size()
	net_exception := false
	if prefix_length == 0 || prefix_length == 32 {
		net_exception = true
	}

	if ip.IsMulticast() {
		errors = append(errors, fmt.Errorf("expected %s to NOT contain the value of an IPv4 multicast address, got %q", k, v))
		return warnings, errors
	}

	if ip.IsLoopback() {
		errors = append(errors, fmt.Errorf("expected %s to NOT contain the value of an IPv4 loopback address, got %q", k, v))
		return warnings, errors
	}

	if ip.Equal(net.IPv4bcast) {
		errors = append(errors, fmt.Errorf("expected %s to contain the value of an IPv4 network address, got %q", k, v))
		return warnings, errors
	}

	if !ip.Equal(network.IP) && !net_exception {
		errors = append(errors, fmt.Errorf("expected %s to contain the value of an IPv4 network address, got %q", k, v))
		return warnings, errors
	}

	return warnings, errors
}

func fortiValidateIPv6Network(i interface{}, k string) (warnings []string, errors []error) {
	v, ok := i.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected type of %q to be string", k))
		return warnings, errors
	}

	ip, network, err := net.ParseCIDR(v)

	if err != nil {
		errors = append(errors, fmt.Errorf("expected %s to contain a valid IPv6 value, got %q: %v", k, v, err))
		return warnings, errors
	}

	exception := false
	prefix_length, _ := network.Mask.Size()
	if prefix_length == 0 || prefix_length == 128 {
		exception = true
	}

	if !ip.Equal(network.IP) && !exception {
		errors = append(errors, fmt.Errorf("expected %s to contain the value of an IPv6 network address, got %q", k, v))
		return warnings, errors
	}

	return warnings, errors
}

func fortiValidateIPv6Prefix(i interface{}, k string) (warnings []string, errors []error) {
	v, ok := i.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected type of %q to be string", k))
		return warnings, errors
	}

	_, _, err := net.ParseCIDR(v)
	if err != nil {
		errors = append(errors, fmt.Errorf("expected %s to contain a valid IPv6 value, got %q: %v", k, v, err))
	}

	return warnings, errors
}

func isBroadcastAddress(v1 net.IP, v2 *net.IPNet) (equal bool) {
	if v2.IP.To4() == nil {
		return false
	}

	if v1.Equal(net.IPv4bcast) {
		return true
	}

	prefix_length, _ := v2.Mask.Size()

	if prefix_length == 31 || prefix_length == 32 {
		return false
	}
	last_ip := make(net.IP, len(v2.IP.To4()))
	binary.BigEndian.PutUint32(last_ip, binary.BigEndian.Uint32(v2.IP.To4())|^binary.BigEndian.Uint32(net.IP(v2.Mask).To4()))

	return v1.Equal(last_ip)
}

func fortiValidateIPv4Netmask(i interface{}, k string) (warnings []string, errors []error) {
	v, ok := i.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected type of %q to be string", k))
		return warnings, errors
	}

	prefix_length, _ := net.IPMask(net.ParseIP(v).To4()).Size()
	if v != "0.0.0.0" && prefix_length == 0 {
		errors = append(errors, fmt.Errorf("expected %s to contain a valid IPv4 netmask, got %q", k, v))
		return warnings, errors
	}

	return warnings, errors
}
