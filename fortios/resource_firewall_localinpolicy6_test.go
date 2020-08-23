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

func TestAccFortiOSFirewallLocalInPolicy6_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSFirewallLocalInPolicy6_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallLocalInPolicy6Config(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallLocalInPolicy6Exists("fortios_firewall_localinpolicy6.trname"),
					resource.TestCheckResourceAttr("fortios_firewall_localinpolicy6.trname", "action", "accept"),
					resource.TestCheckResourceAttr("fortios_firewall_localinpolicy6.trname", "intf", "port4"),
					resource.TestCheckResourceAttr("fortios_firewall_localinpolicy6.trname", "policyid", "1"),
					resource.TestCheckResourceAttr("fortios_firewall_localinpolicy6.trname", "schedule", "always"),
					resource.TestCheckResourceAttr("fortios_firewall_localinpolicy6.trname", "status", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_localinpolicy6.trname", "dstaddr.0.name", "all"),
					resource.TestCheckResourceAttr("fortios_firewall_localinpolicy6.trname", "service.0.name", "ALL"),
					resource.TestCheckResourceAttr("fortios_firewall_localinpolicy6.trname", "srcaddr.0.name", "all"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallLocalInPolicy6Exists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallLocalInPolicy6: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallLocalInPolicy6 is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallLocalInPolicy6(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallLocalInPolicy6: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallLocalInPolicy6: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallLocalInPolicy6Destroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_localinpolicy6" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallLocalInPolicy6(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallLocalInPolicy6 %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallLocalInPolicy6Config(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_localinpolicy6" "trname" {
  action   = "accept"
  intf     = "port4"
  policyid = 1
  schedule = "always"
  status   = "enable"

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
