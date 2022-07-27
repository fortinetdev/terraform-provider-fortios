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

func TestAccFortiOSFirewallPolicy_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSFirewallPolicy_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallPolicyConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallPolicyExists("fortios_firewall_policy.trname"),
					resource.TestCheckResourceAttr("fortios_firewall_policy.trname", "action", "accept"),
					resource.TestCheckResourceAttr("fortios_firewall_policy.trname", "logtraffic", "utm"),
					resource.TestCheckResourceAttr("fortios_firewall_policy.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewall_policy.trname", "policyid", "1"),
					resource.TestCheckResourceAttr("fortios_firewall_policy.trname", "schedule", "always"),
					resource.TestCheckResourceAttr("fortios_firewall_policy.trname", "wanopt", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_policy.trname", "wanopt_detection", "active"),
					resource.TestCheckResourceAttr("fortios_firewall_policy.trname", "wanopt_passive_opt", "default"),
					resource.TestCheckResourceAttr("fortios_firewall_policy.trname", "wccp", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_policy.trname", "webcache", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_policy.trname", "webcache_https", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_policy.trname", "wsso", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_policy.trname", "dstaddr.0.name", "all"),
					resource.TestCheckResourceAttr("fortios_firewall_policy.trname", "dstintf.0.name", "port4"),
					resource.TestCheckResourceAttr("fortios_firewall_policy.trname", "service.0.name", "HTTP"),
					resource.TestCheckResourceAttr("fortios_firewall_policy.trname", "srcaddr.0.name", "all"),
					resource.TestCheckResourceAttr("fortios_firewall_policy.trname", "srcintf.0.name", "port3"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallPolicyExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallPolicy: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallPolicy is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallPolicy(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallPolicy: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallPolicy: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallPolicyDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_policy" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallPolicy(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallPolicy %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallPolicyConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_policy" "trname" {
  action             = "accept"
  logtraffic         = "utm"
  name               = "%[1]s"
  policyid           = 1
  schedule           = "always"
  wanopt             = "disable"
  wanopt_detection   = "active"
  wanopt_passive_opt = "default"
  wccp               = "disable"
  webcache           = "disable"
  webcache_https     = "disable"
  wsso               = "enable"

  dstaddr {
    name = "all"
  }

  dstintf {
    name = "port4"
  }

  service {
    name = "HTTP"
  }

  srcaddr {
    name = "all"
  }

  srcintf {
    name = "port3"
  }
}
`, name)
}
