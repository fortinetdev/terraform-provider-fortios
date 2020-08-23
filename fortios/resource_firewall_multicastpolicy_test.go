
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

func TestAccFortiOSFirewallMulticastPolicy_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSFirewallMulticastPolicy_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSFirewallMulticastPolicyConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSFirewallMulticastPolicyExists("fortios_firewall_multicastpolicy.trname"),
                    resource.TestCheckResourceAttr("fortios_firewall_multicastpolicy.trname", "action", "accept"),
                    resource.TestCheckResourceAttr("fortios_firewall_multicastpolicy.trname", "dnat", "0.0.0.0"),
                    resource.TestCheckResourceAttr("fortios_firewall_multicastpolicy.trname", "dstintf", "port4"),
                    resource.TestCheckResourceAttr("fortios_firewall_multicastpolicy.trname", "end_port", "65535"),
                    resource.TestCheckResourceAttr("fortios_firewall_multicastpolicy.trname", "fosid", "1"),
                    resource.TestCheckResourceAttr("fortios_firewall_multicastpolicy.trname", "logtraffic", "enable"),
                    resource.TestCheckResourceAttr("fortios_firewall_multicastpolicy.trname", "protocol", "0"),
                    resource.TestCheckResourceAttr("fortios_firewall_multicastpolicy.trname", "snat", "disable"),
                    resource.TestCheckResourceAttr("fortios_firewall_multicastpolicy.trname", "snat_ip", "0.0.0.0"),
                    resource.TestCheckResourceAttr("fortios_firewall_multicastpolicy.trname", "srcintf", "port3"),
                    resource.TestCheckResourceAttr("fortios_firewall_multicastpolicy.trname", "start_port", "1"),
                    resource.TestCheckResourceAttr("fortios_firewall_multicastpolicy.trname", "status", "enable"),
                    resource.TestCheckResourceAttr("fortios_firewall_multicastpolicy.trname", "dstaddr.0.name", "all"),
                    resource.TestCheckResourceAttr("fortios_firewall_multicastpolicy.trname", "srcaddr.0.name", "all"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSFirewallMulticastPolicyExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallMulticastPolicy: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallMulticastPolicy is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallMulticastPolicy(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallMulticastPolicy: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallMulticastPolicy: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallMulticastPolicyDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_multicastpolicy" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallMulticastPolicy(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallMulticastPolicy %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallMulticastPolicyConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_multicastpolicy" "trname" {
  action     = "accept"
  dnat       = "0.0.0.0"
  dstintf    = "port4"
  end_port   = 65535
  fosid      = 1
  logtraffic = "enable"
  protocol   = 0
  snat       = "disable"
  snat_ip    = "0.0.0.0"
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
