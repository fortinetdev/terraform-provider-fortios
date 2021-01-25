// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

package fortios

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"log"
	"testing"
)

func TestAccFortiOSFirewallPolicy6_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSFirewallPolicy6_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallPolicy6Config(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallPolicy6Exists("fortios_firewall_policy6.trname"),
					resource.TestCheckResourceAttr("fortios_firewall_policy6.trname", "action", "deny"),
					resource.TestCheckResourceAttr("fortios_firewall_policy6.trname", "diffserv_forward", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_policy6.trname", "diffserv_reverse", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_policy6.trname", "diffservcode_forward", "000000"),
					resource.TestCheckResourceAttr("fortios_firewall_policy6.trname", "diffservcode_rev", "000000"),
					resource.TestCheckResourceAttr("fortios_firewall_policy6.trname", "dsri", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_policy6.trname", "dstaddr_negate", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_policy6.trname", "firewall_session_dirty", "check-all"),
					resource.TestCheckResourceAttr("fortios_firewall_policy6.trname", "fixedport", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_policy6.trname", "inbound", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_policy6.trname", "ippool", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_policy6.trname", "logtraffic", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_policy6.trname", "logtraffic_start", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_policy6.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewall_policy6.trname", "nat", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_policy6.trname", "natinbound", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_policy6.trname", "natoutbound", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_policy6.trname", "outbound", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_policy6.trname", "policyid", "1"),
					resource.TestCheckResourceAttr("fortios_firewall_policy6.trname", "profile_protocol_options", "default"),
					resource.TestCheckResourceAttr("fortios_firewall_policy6.trname", "profile_type", "single"),
					resource.TestCheckResourceAttr("fortios_firewall_policy6.trname", "rsso", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_policy6.trname", "schedule", "always"),
					resource.TestCheckResourceAttr("fortios_firewall_policy6.trname", "send_deny_packet", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_policy6.trname", "service_negate", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_policy6.trname", "srcaddr_negate", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_policy6.trname", "ssl_mirror", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_policy6.trname", "status", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_policy6.trname", "tcp_mss_receiver", "0"),
					resource.TestCheckResourceAttr("fortios_firewall_policy6.trname", "tcp_mss_sender", "0"),
					resource.TestCheckResourceAttr("fortios_firewall_policy6.trname", "tcp_session_without_syn", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_policy6.trname", "timeout_send_rst", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_policy6.trname", "tos", "0x00"),
					resource.TestCheckResourceAttr("fortios_firewall_policy6.trname", "tos_mask", "0x00"),
					resource.TestCheckResourceAttr("fortios_firewall_policy6.trname", "tos_negate", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_policy6.trname", "utm_status", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_policy6.trname", "dstaddr.0.name", "all"),
					resource.TestCheckResourceAttr("fortios_firewall_policy6.trname", "dstintf.0.name", "port3"),
					resource.TestCheckResourceAttr("fortios_firewall_policy6.trname", "service.0.name", "ALL"),
					resource.TestCheckResourceAttr("fortios_firewall_policy6.trname", "srcaddr.0.name", "all"),
					resource.TestCheckResourceAttr("fortios_firewall_policy6.trname", "srcintf.0.name", "port4"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallPolicy6Exists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallPolicy6: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallPolicy6 is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallPolicy6(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallPolicy6: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallPolicy6: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallPolicy6Destroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_policy6" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallPolicy6(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallPolicy6 %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallPolicy6Config(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_policy6" "trname" {
  action                   = "deny"
  diffserv_forward         = "disable"
  diffserv_reverse         = "disable"
  diffservcode_forward     = "000000"
  diffservcode_rev         = "000000"
  dsri                     = "disable"
  dstaddr_negate           = "disable"
  firewall_session_dirty   = "check-all"
  fixedport                = "disable"
  inbound                  = "disable"
  ippool                   = "disable"
  logtraffic               = "disable"
  logtraffic_start         = "disable"
  name                     = "%[1]s"
  nat                      = "disable"
  natinbound               = "disable"
  natoutbound              = "disable"
  outbound                 = "disable"
  policyid                 = 1
  profile_protocol_options = "default"
  profile_type             = "single"
  rsso                     = "disable"
  schedule                 = "always"
  send_deny_packet         = "disable"
  service_negate           = "disable"
  srcaddr_negate           = "disable"
  ssl_mirror               = "disable"
  status                   = "enable"
  tcp_mss_receiver         = 0
  tcp_mss_sender           = 0
  tcp_session_without_syn  = "disable"
  timeout_send_rst         = "disable"
  tos                      = "0x00"
  tos_mask                 = "0x00"
  tos_negate               = "disable"
  utm_status               = "disable"

  dstaddr {
    name = "all"
  }

  dstintf {
    name = "port3"
  }

  service {
    name = "ALL"
  }

  srcaddr {
    name = "all"
  }

  srcintf {
    name = "port4"
  }
}
`, name)
}
