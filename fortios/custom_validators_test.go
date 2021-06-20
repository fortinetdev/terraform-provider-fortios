package fortios

import (
	"fmt"
	"net"
	"testing"
)

func TestFortiValidateIPv4ClassnetAny(t *testing.T) {
	var tests = []struct {
		a     interface{}
		valid bool
	}{
		{"192.168.0.1", false},
		{"192.168.0.0/24", true},
		{"192.168.0.0 255.255.255.0", true},
		{"192.168.0.0/255.255.255.0", false},
		{"192.168.2.0/16", false},
		{"192.168.2.0 255.255.0.0", false},
		{"192.168.0.1/24", false},
		{"192.168.0.1 255.255.255.0", false},
		{"192.168.0.1/255.255.255.0", false},
		{"192.168.0.255/24", false},
		{"192.168.0.255 255.255.255.0", false},
		{"192.168.0.255/255.255.255.0", false},
		{"10.10.10.80/29", true},
		{"10.10.10.80 255.255.255.248", true},
		{"10.10.10.80/255.255.255.248", false},
		{"10.10.10.76/29", false},
		{"10.10.10.76 255.255.255.248", false},
		{"10.10.10.10/32", true},
		{"10.10.10.10 255.255.255.255", true},
		{"172.16.16.4/31", true},
		{"172.16.16.4 255.255.255.254", true},
		{"172.16.16.5/31", false},
		{"172.16.16.5 255.255.255.254", false},
		{"0.0.0.0/0", true},
		{"255.255.255.255/32", true},
		{"127.0.0.1/32", true},
		{"224.0.0.1/32", true},
		{"invalidipv4address", false},
		{"192.168.256.0/24", false},
		{"192.168.0.0 255.255.256.0", false},
		{"192.168.0.0 255.0.255.0", false},
		{"192.168.0.0 255.255.0", false},
		{"192.168.0 255.255.255.0", false},
		{"192.168.0/24", false},
		{"192.168.0.0/33", false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.a)
		t.Run(testname, func(t *testing.T) {
			_, ans := fortiValidateIPv4ClassnetAny(tt.a, "testing")
			ans_bool := len(ans) == 0
			if ans_bool != tt.valid {
				t.Errorf("got %v, want %v", ans, tt.valid)
			}
		})
	}
}

func TestFortiValidateIPv4ClassnetHost(t *testing.T) {
	var tests = []struct {
		a     interface{}
		valid bool
	}{
		{"192.168.0.1", false},
		{"192.168.0.0/24", false},
		{"192.168.0.0 255.255.255.0", false},
		{"192.168.0.0/255.255.255.0", false},
		{"192.168.2.0/16", true},
		{"192.168.2.0 255.255.0.0", true},
		{"192.168.0.1/24", true},
		{"192.168.0.1 255.255.255.0", true},
		{"192.168.0.1/255.255.255.0", false},
		{"192.168.0.255/24", false},
		{"192.168.0.255 255.255.255.0", false},
		{"192.168.0.255/255.255.255.0", false},
		{"10.10.10.80/29", false},
		{"10.10.10.80 255.255.255.248", false},
		{"10.10.10.80/255.255.255.248", false},
		{"10.10.10.76/29", true},
		{"10.10.10.76 255.255.255.248", true},
		{"10.10.10.10/32", true},
		{"10.10.10.10 255.255.255.255", true},
		{"172.16.16.4/31", true},
		{"172.16.16.4 255.255.255.254", true},
		{"172.16.16.5/31", true},
		{"172.16.16.5 255.255.255.254", true},
		{"0.0.0.0/0", true},
		{"255.255.255.255/32", false},
		{"127.0.0.1/32", false},
		{"224.0.0.1/32", false},
		{"invalidipv4address", false},
		{"192.168.256.0/24", false},
		{"192.168.0.0 255.255.256.0", false},
		{"192.168.0.0 255.0.255.0", false},
		{"192.168.0.0 255.255.0", false},
		{"192.168.0 255.255.255.0", false},
		{"192.168.0/24", false},
		{"192.168.0.0/33", false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.a)
		t.Run(testname, func(t *testing.T) {
			_, ans := fortiValidateIPv4ClassnetHost(tt.a, "testing")
			ans_bool := len(ans) == 0
			if ans_bool != tt.valid {
				t.Errorf("got %v, want %v", ans, tt.valid)
			}
		})
	}
}

func TestFortiValidateIPv4Classnet(t *testing.T) {
	var tests = []struct {
		a     interface{}
		valid bool
	}{
		{"192.168.0.1", false},
		{"192.168.0.0/24", true},
		{"192.168.0.0 255.255.255.0", true},
		{"192.168.0.0/255.255.255.0", false},
		{"192.168.2.0/16", false},
		{"192.168.2.0 255.255.0.0", false},
		{"192.168.0.1/24", false},
		{"192.168.0.1 255.255.255.0", false},
		{"192.168.0.1/255.255.255.0", false},
		{"192.168.0.255/24", false},
		{"192.168.0.255 255.255.255.0", false},
		{"192.168.0.255/255.255.255.0", false},
		{"10.10.10.80/29", true},
		{"10.10.10.80 255.255.255.248", true},
		{"10.10.10.80/255.255.255.248", false},
		{"10.10.10.76/29", false},
		{"10.10.10.76 255.255.255.248", false},
		{"10.10.10.10/32", true},
		{"10.10.10.10 255.255.255.255", true},
		{"172.16.16.4/31", true},
		{"172.16.16.4 255.255.255.254", true},
		{"172.16.16.5/31", false},
		{"172.16.16.5 255.255.255.254", false},
		{"0.0.0.0/0", true},
		{"255.255.255.255/32", false},
		{"127.0.0.1/32", false},
		{"224.0.0.1/32", false},
		{"invalidipv4address", false},
		{"192.168.256.0/24", false},
		{"192.168.0.0 255.255.256.0", false},
		{"192.168.0.0 255.0.255.0", false},
		{"192.168.0.0 255.255.0", false},
		{"192.168.0 255.255.255.0", false},
		{"192.168.0/24", false},
		{"192.168.0.0/33", false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.a)
		t.Run(testname, func(t *testing.T) {
			_, ans := fortiValidateIPv4Classnet(tt.a, "testing")
			ans_bool := len(ans) == 0
			if ans_bool != tt.valid {
				t.Errorf("got %v, want %v", ans, tt.valid)
			}
		})
	}
}

func TestFortiValidateIPv6Network(t *testing.T) {
	var tests = []struct {
		a     interface{}
		valid bool
	}{
		{"2001::1/128", true},
		{"2001::1/32", false},
		{"::/0", true},
		{"2001::1", false},
		{"fe80::1/10", false},
		{"fe80::/10", true},
		{"fe80:1234:1234::/10", false},
		{"20ac::/10", false},
		{"2001:0db8:85a3:0000:0000:8a2e:0000:7334/127", true},
		{"2001:0db8:85a3::8a2e:0000:7334/127", true},
		{"2001:0db8:85a3::8a2e:0000:7335/127", false},
		{"2001:0db8:85a3::8a2e:0000:7335/128", true},
		{"2001:0db8:85a3::8a2e::7334/128", false},
		{"invalidipv6address", false},
		{"2001::1/129", false},
		{"2001:g::1/10", false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.a)
		t.Run(testname, func(t *testing.T) {
			_, ans := fortiValidateIPv6Network(tt.a, "testing")
			ans_bool := len(ans) == 0
			if ans_bool != tt.valid {
				t.Errorf("got %v, want %v", ans, tt.valid)
			}
		})
	}
}

func TestFortiValidateIPv6Prefix(t *testing.T) {
	var tests = []struct {
		a     interface{}
		valid bool
	}{
		{"2001::1/128", true},
		{"2001::1/32", true},
		{"::/0", true},
		{"2001::1", false},
		{"fe80::1/10", true},
		{"fe80::/10", true},
		{"20ac::/10", true},
		{"2001:0db8:85a3:0000:0000:8a2e:0000:7334/127", true},
		{"2001:0db8:85a3::8a2e:0000:7334/127", true},
		{"2001:0db8:85a3::8a2e:0000:7335/127", true},
		{"2001:0db8:85a3::8a2e:0000:7335/128", true},
		{"2001:0db8:85a3::8a2e::7334/128", false},
		{"invalidipv6address", false},
		{"2001::1/129", false},
		{"2001:g::1/10", false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.a)
		t.Run(testname, func(t *testing.T) {
			_, ans := fortiValidateIPv6Prefix(tt.a, "testing")
			ans_bool := len(ans) == 0
			if ans_bool != tt.valid {
				t.Errorf("got %v, want %v", ans, tt.valid)
			}
		})
	}
}

func TestIsBroadcastAddress(t *testing.T) {
	var tests = []struct {
		a    string
		want bool
	}{
		{"192.168.0.0/24", false},
		{"192.168.0.1/24", false},
		{"192.168.0.255/24", true},
		{"10.10.20.80/29", false},
		{"10.10.20.79/29", true},
		{"10.10.20.76/29", false},
		{"8.8.8.8/32", false},
		{"172.16.16.4/31", false},
		{"172.16.16.5/31", false},
		{"0.0.0.0/0", false},
		{"255.255.255.255/32", true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.a)
		t.Run(testname, func(t *testing.T) {
			ip, network, _ := net.ParseCIDR(tt.a)
			ans := isBroadcastAddress(ip, network)
			if ans != tt.want {
				t.Errorf("got %v want %v", ans, tt.want)
			}
		})
	}
}

func TestFortiValidateIPv4Netmask(t *testing.T) {
	var tests = []struct {
		a     string
		valid bool
	}{
		{"255.255.255.0", true},
		{"255.128.0.0", true},
		{"0.0.0.0", true},
		{"255.255.255.254", true},
		{"255.255.255.255", true},
		{"255.0.255.0", false},
		{"255.255.256.0", false},
		{"10.10.10.10", false},
		{"invalidnetmask", false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.a)
		t.Run(testname, func(t *testing.T) {
			_, ans := fortiValidateIPv4Netmask(tt.a, "testing")
			ans_bool := len(ans) == 0
			if ans_bool != tt.valid {
				t.Errorf("got %v, want %v", ans, tt.valid)
			}
		})
	}
}
