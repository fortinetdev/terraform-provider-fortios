// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

package fortios

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"log"
	"testing"
)

func TestAccFortiOSFirewallProxyPolicy_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSFirewallProxyPolicy_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallProxyPolicyConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallProxyPolicyExists("fortios_firewall_proxypolicy.trname"),
					resource.TestCheckResourceAttr("fortios_firewall_proxypolicy.trname", "action", "deny"),
					resource.TestCheckResourceAttr("fortios_firewall_proxypolicy.trname", "disclaimer", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_proxypolicy.trname", "dstaddr_negate", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_proxypolicy.trname", "http_tunnel_auth", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_proxypolicy.trname", "internet_service", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_proxypolicy.trname", "internet_service_negate", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_proxypolicy.trname", "logtraffic", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_proxypolicy.trname", "logtraffic_start", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_proxypolicy.trname", "policyid", "1"),
					resource.TestCheckResourceAttr("fortios_firewall_proxypolicy.trname", "profile_protocol_options", "default"),
					resource.TestCheckResourceAttr("fortios_firewall_proxypolicy.trname", "profile_type", "single"),
					resource.TestCheckResourceAttr("fortios_firewall_proxypolicy.trname", "proxy", "transparent-web"),
					resource.TestCheckResourceAttr("fortios_firewall_proxypolicy.trname", "scan_botnet_connections", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_proxypolicy.trname", "schedule", "always"),
					resource.TestCheckResourceAttr("fortios_firewall_proxypolicy.trname", "service_negate", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_proxypolicy.trname", "srcaddr_negate", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_proxypolicy.trname", "status", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_proxypolicy.trname", "transparent", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_proxypolicy.trname", "utm_status", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_proxypolicy.trname", "webcache", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_proxypolicy.trname", "webcache_https", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_proxypolicy.trname", "dstaddr.0.name", "all"),
					resource.TestCheckResourceAttr("fortios_firewall_proxypolicy.trname", "dstintf.0.name", "port4"),
					resource.TestCheckResourceAttr("fortios_firewall_proxypolicy.trname", "service.0.name", "webproxy"),
					resource.TestCheckResourceAttr("fortios_firewall_proxypolicy.trname", "srcaddr.0.name", "all"),
					resource.TestCheckResourceAttr("fortios_firewall_proxypolicy.trname", "srcintf.0.name", "port3"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallProxyPolicyExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallProxyPolicy: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallProxyPolicy is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallProxyPolicy(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallProxyPolicy: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallProxyPolicy: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallProxyPolicyDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_proxypolicy" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallProxyPolicy(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallProxyPolicy %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallProxyPolicyConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_proxypolicy" "trname" {
  action                   = "deny"
  disclaimer               = "disable"
  dstaddr_negate           = "disable"
  http_tunnel_auth         = "disable"
  internet_service         = "disable"
  internet_service_negate  = "disable"
  logtraffic               = "disable"
  logtraffic_start         = "disable"
  policyid                 = 1
  profile_protocol_options = "default"
  profile_type             = "single"
  proxy                    = "transparent-web"
  scan_botnet_connections  = "disable"
  schedule                 = "always"
  service_negate           = "disable"
  srcaddr_negate           = "disable"
  status                   = "enable"
  transparent              = "disable"
  utm_status               = "disable"
  webcache                 = "disable"
  webcache_https           = "disable"

  dstaddr {
    name = "all"
  }

  dstintf {
    name = "port4"
  }

  service {
    name = "webproxy"
  }

  srcaddr {
    name = "all"
  }

  srcintf {
    name = "port3"
  }
}
`)
}
