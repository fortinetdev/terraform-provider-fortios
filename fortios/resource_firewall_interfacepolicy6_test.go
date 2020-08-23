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

func TestAccFortiOSFirewallInterfacePolicy6_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSFirewallInterfacePolicy6_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallInterfacePolicy6Config(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallInterfacePolicy6Exists("fortios_firewall_interfacepolicy6.trname"),
					resource.TestCheckResourceAttr("fortios_firewall_interfacepolicy6.trname", "address_type", "ipv6"),
					resource.TestCheckResourceAttr("fortios_firewall_interfacepolicy6.trname", "application_list_status", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_interfacepolicy6.trname", "av_profile_status", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_interfacepolicy6.trname", "dlp_sensor_status", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_interfacepolicy6.trname", "dsri", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_interfacepolicy6.trname", "interface", "port4"),
					resource.TestCheckResourceAttr("fortios_firewall_interfacepolicy6.trname", "ips_sensor_status", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_interfacepolicy6.trname", "logtraffic", "all"),
					resource.TestCheckResourceAttr("fortios_firewall_interfacepolicy6.trname", "policyid", "1"),
					resource.TestCheckResourceAttr("fortios_firewall_interfacepolicy6.trname", "scan_botnet_connections", "block"),
					resource.TestCheckResourceAttr("fortios_firewall_interfacepolicy6.trname", "spamfilter_profile_status", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_interfacepolicy6.trname", "status", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_interfacepolicy6.trname", "webfilter_profile_status", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_interfacepolicy6.trname", "dstaddr6.0.name", "all"),
					resource.TestCheckResourceAttr("fortios_firewall_interfacepolicy6.trname", "service6.0.name", "ALL"),
					resource.TestCheckResourceAttr("fortios_firewall_interfacepolicy6.trname", "srcaddr6.0.name", "all"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallInterfacePolicy6Exists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallInterfacePolicy6: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallInterfacePolicy6 is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallInterfacePolicy6(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallInterfacePolicy6: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallInterfacePolicy6: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallInterfacePolicy6Destroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_interfacepolicy6" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallInterfacePolicy6(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallInterfacePolicy6 %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallInterfacePolicy6Config(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_interfacepolicy6" "trname" {
  address_type              = "ipv6"
  application_list_status   = "disable"
  av_profile_status         = "disable"
  dlp_sensor_status         = "disable"
  dsri                      = "disable"
  interface                 = "port4"
  ips_sensor_status         = "disable"
  logtraffic                = "all"
  policyid                  = 1
  scan_botnet_connections   = "block"
  spamfilter_profile_status = "disable"
  status                    = "enable"
  webfilter_profile_status  = "disable"

  dstaddr6 {
    name = "all"
  }

  service6 {
    name = "ALL"
  }

  srcaddr6 {
    name = "all"
  }
}
`)
}
