
// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

package fortios
import (
    "fmt"
	"log"
    "testing"
    "github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
    "github.com/hashicorp/terraform-plugin-sdk/helper/resource"
    "github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccFortiOSFirewallMulticastPolicy6_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSFirewallMulticastPolicy6_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSFirewallMulticastPolicy6Config(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSFirewallMulticastPolicy6Exists("fortios_firewall_multicastpolicy6.trname"),
                    resource.TestCheckResourceAttr("fortios_firewall_multicastpolicy6.trname", "action", "accept"),
                    resource.TestCheckResourceAttr("fortios_firewall_multicastpolicy6.trname", "dstintf", "port4"),
                    resource.TestCheckResourceAttr("fortios_firewall_multicastpolicy6.trname", "end_port", "65535"),
                    resource.TestCheckResourceAttr("fortios_firewall_multicastpolicy6.trname", "fosid", "1"),
                    resource.TestCheckResourceAttr("fortios_firewall_multicastpolicy6.trname", "logtraffic", "disable"),
                    resource.TestCheckResourceAttr("fortios_firewall_multicastpolicy6.trname", "protocol", "0"),
                    resource.TestCheckResourceAttr("fortios_firewall_multicastpolicy6.trname", "srcintf", "port3"),
                    resource.TestCheckResourceAttr("fortios_firewall_multicastpolicy6.trname", "start_port", "1"),
                    resource.TestCheckResourceAttr("fortios_firewall_multicastpolicy6.trname", "status", "enable"),
                    resource.TestCheckResourceAttr("fortios_firewall_multicastpolicy6.trname", "dstaddr.0.name", "all"),
                    resource.TestCheckResourceAttr("fortios_firewall_multicastpolicy6.trname", "srcaddr.0.name", "all"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSFirewallMulticastPolicy6Exists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallMulticastPolicy6: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallMulticastPolicy6 is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallMulticastPolicy6(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallMulticastPolicy6: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallMulticastPolicy6: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallMulticastPolicy6Destroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_multicastpolicy6" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallMulticastPolicy6(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallMulticastPolicy6 %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallMulticastPolicy6Config(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_multicastpolicy6" "trname" {
  action     = "accept"
  dstintf    = "port4"
  end_port   = 65535
  fosid      = 1
  logtraffic = "disable"
  protocol   = 0
  srcintf    = "port3"
  start_port = 1
  status     = "enable"

  dstaddr {
    name = "all"
  }

  srcaddr {
    name = "all"
  }
}
`)
}
