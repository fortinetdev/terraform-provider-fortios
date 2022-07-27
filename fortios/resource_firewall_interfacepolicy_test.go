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

func TestAccFortiOSFirewallInterfacePolicy_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSFirewallInterfacePolicy_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallInterfacePolicyConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallInterfacePolicyExists("fortios_firewall_interfacepolicy.trname"),
					resource.TestCheckResourceAttr("fortios_firewall_interfacepolicy.trname", "address_type", "ipv4"),
					resource.TestCheckResourceAttr("fortios_firewall_interfacepolicy.trname", "application_list_status", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_interfacepolicy.trname", "av_profile_status", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_interfacepolicy.trname", "dlp_sensor_status", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_interfacepolicy.trname", "dsri", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_interfacepolicy.trname", "interface", "port4"),
					resource.TestCheckResourceAttr("fortios_firewall_interfacepolicy.trname", "ips_sensor_status", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_interfacepolicy.trname", "logtraffic", "all"),
					resource.TestCheckResourceAttr("fortios_firewall_interfacepolicy.trname", "policyid", "1"),
					resource.TestCheckResourceAttr("fortios_firewall_interfacepolicy.trname", "scan_botnet_connections", "block"),
					resource.TestCheckResourceAttr("fortios_firewall_interfacepolicy.trname", "spamfilter_profile_status", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_interfacepolicy.trname", "status", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_interfacepolicy.trname", "webfilter_profile_status", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_interfacepolicy.trname", "dstaddr.0.name", "all"),
					resource.TestCheckResourceAttr("fortios_firewall_interfacepolicy.trname", "service.0.name", "ALL"),
					resource.TestCheckResourceAttr("fortios_firewall_interfacepolicy.trname", "srcaddr.0.name", "all"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallInterfacePolicyExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallInterfacePolicy: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallInterfacePolicy is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallInterfacePolicy(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallInterfacePolicy: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallInterfacePolicy: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallInterfacePolicyDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_interfacepolicy" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallInterfacePolicy(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallInterfacePolicy %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallInterfacePolicyConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_interfacepolicy" "trname" {
  address_type              = "ipv4"
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
