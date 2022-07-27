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

func TestAccFortiOSFirewallTtlPolicy_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSFirewallTtlPolicy_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallTtlPolicyConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallTtlPolicyExists("fortios_firewall_ttlpolicy.trname"),
					resource.TestCheckResourceAttr("fortios_firewall_ttlpolicy.trname", "action", "accept"),
					resource.TestCheckResourceAttr("fortios_firewall_ttlpolicy.trname", "fosid", "1"),
					resource.TestCheckResourceAttr("fortios_firewall_ttlpolicy.trname", "schedule", "always"),
					resource.TestCheckResourceAttr("fortios_firewall_ttlpolicy.trname", "srcintf", "port3"),
					resource.TestCheckResourceAttr("fortios_firewall_ttlpolicy.trname", "status", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_ttlpolicy.trname", "ttl", "23"),
					resource.TestCheckResourceAttr("fortios_firewall_ttlpolicy.trname", "service.0.name", "ALL"),
					resource.TestCheckResourceAttr("fortios_firewall_ttlpolicy.trname", "srcaddr.0.name", "all"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallTtlPolicyExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallTtlPolicy: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallTtlPolicy is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallTtlPolicy(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallTtlPolicy: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallTtlPolicy: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallTtlPolicyDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_ttlpolicy" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallTtlPolicy(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallTtlPolicy %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallTtlPolicyConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_ttlpolicy" "trname" {
  action   = "accept"
  fosid    = 1
  schedule = "always"
  srcintf  = "port3"
  status   = "enable"
  ttl      = "23"

  service {
    name = "ALL"
  }

  srcaddr {
    name = "all"
  }
}
`)
}
