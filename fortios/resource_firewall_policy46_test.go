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

func TestAccFortiOSFirewallPolicy46_basic(t *testing.T) {
	rname := acctest.RandString(8)
	var0 := "var0" + rname
	log.Printf(var0)
	log.Printf("TestAccFortiOSFirewallPolicy46_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallPolicy46Config(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallPolicy46Exists("fortios_firewall_policy46.trname"),
					resource.TestCheckResourceAttr("fortios_firewall_policy46.trname", "action", "deny"),
					resource.TestCheckResourceAttr("fortios_firewall_policy46.trname", "dstintf", "port3"),
					resource.TestCheckResourceAttr("fortios_firewall_policy46.trname", "fixedport", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_policy46.trname", "ippool", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_policy46.trname", "logtraffic", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_policy46.trname", "permit_any_host", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_policy46.trname", "policyid", "2"),
					resource.TestCheckResourceAttr("fortios_firewall_policy46.trname", "schedule", "always"),
					resource.TestCheckResourceAttr("fortios_firewall_policy46.trname", "srcintf", "port2"),
					resource.TestCheckResourceAttr("fortios_firewall_policy46.trname", "status", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_policy46.trname", "tcp_mss_receiver", "0"),
					resource.TestCheckResourceAttr("fortios_firewall_policy46.trname", "tcp_mss_sender", "0"),
					resource.TestCheckResourceAttr("fortios_firewall_policy46.trname", "dstaddr.0.name", var0),
					resource.TestCheckResourceAttr("fortios_firewall_policy46.trname", "service.0.name", "ALL"),
					resource.TestCheckResourceAttr("fortios_firewall_policy46.trname", "srcaddr.0.name", "FIREWALL_AUTH_PORTAL_ADDRESS"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallPolicy46Exists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallPolicy46: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallPolicy46 is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallPolicy46(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallPolicy46: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallPolicy46: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallPolicy46Destroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_policy46" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallPolicy46(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallPolicy46 %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallPolicy46Config(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_vip46" "trname" {
  arp_reply   = "enable"
  color       = 0
  extip       = "10.1.100.55"
  extport     = "0-65535"
  fosid       = 0
  ldb_method  = "static"
  mappedip    = "2000:172:16:200::55"
  mappedport  = "0-65535"
  name        = "var0%[1]s"
  portforward = "disable"
  protocol    = "tcp"
  type        = "static-nat"
}

resource "fortios_firewall_policy46" "trname" {
  action           = "deny"
  dstintf          = "port3"
  fixedport        = "disable"
  ippool           = "disable"
  logtraffic       = "disable"
  permit_any_host  = "disable"
  policyid         = 2
  schedule         = "always"
  srcintf          = "port2"
  status           = "enable"
  tcp_mss_receiver = 0
  tcp_mss_sender   = 0

  dstaddr {
    name = fortios_firewall_vip46.trname.name
  }

  service {
    name = "ALL"
  }

  srcaddr {
    name = "FIREWALL_AUTH_PORTAL_ADDRESS"
  }
}
`, name)
}
