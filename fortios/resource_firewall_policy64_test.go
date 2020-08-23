// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

package fortios

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"log"
	"testing"
)

func TestAccFortiOSFirewallPolicy64_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSFirewallPolicy64_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallPolicy64Config(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallPolicy64Exists("fortios_firewall_policy64.trname"),
					resource.TestCheckResourceAttr("fortios_firewall_policy64.trname", "action", "accept"),
					resource.TestCheckResourceAttr("fortios_firewall_policy64.trname", "dstintf", "port4"),
					resource.TestCheckResourceAttr("fortios_firewall_policy64.trname", "fixedport", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_policy64.trname", "ippool", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_policy64.trname", "logtraffic", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_policy64.trname", "permit_any_host", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_policy64.trname", "policyid", "1"),
					resource.TestCheckResourceAttr("fortios_firewall_policy64.trname", "schedule", "always"),
					resource.TestCheckResourceAttr("fortios_firewall_policy64.trname", "srcintf", "port3"),
					resource.TestCheckResourceAttr("fortios_firewall_policy64.trname", "status", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_policy64.trname", "tcp_mss_receiver", "0"),
					resource.TestCheckResourceAttr("fortios_firewall_policy64.trname", "tcp_mss_sender", "0"),
					resource.TestCheckResourceAttr("fortios_firewall_policy64.trname", "dstaddr.0.name", "all"),
					resource.TestCheckResourceAttr("fortios_firewall_policy64.trname", "service.0.name", "ALL"),
					resource.TestCheckResourceAttr("fortios_firewall_policy64.trname", "srcaddr.0.name", "all"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallPolicy64Exists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallPolicy64: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallPolicy64 is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallPolicy64(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallPolicy64: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallPolicy64: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallPolicy64Destroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_policy64" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallPolicy64(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallPolicy64 %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallPolicy64Config(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_policy64" "trname" {
  action           = "accept"
  dstintf          = "port4"
  fixedport        = "disable"
  ippool           = "disable"
  logtraffic       = "disable"
  permit_any_host  = "disable"
  policyid         = 1
  schedule         = "always"
  srcintf          = "port3"
  status           = "enable"
  tcp_mss_receiver = 0
  tcp_mss_sender   = 0

  dstaddr {
    name = "all"
  }

  service {
    name = "ALL"
  }

  srcaddr {
    name = "all"
  }
}

`)
}
