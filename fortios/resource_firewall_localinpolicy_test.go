
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

func TestAccFortiOSFirewallLocalInPolicy_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSFirewallLocalInPolicy_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSFirewallLocalInPolicyConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSFirewallLocalInPolicyExists("fortios_firewall_localinpolicy.trname"),
                    resource.TestCheckResourceAttr("fortios_firewall_localinpolicy.trname", "action", "accept"),
                    resource.TestCheckResourceAttr("fortios_firewall_localinpolicy.trname", "ha_mgmt_intf_only", "disable"),
                    resource.TestCheckResourceAttr("fortios_firewall_localinpolicy.trname", "intf", "port4"),
                    resource.TestCheckResourceAttr("fortios_firewall_localinpolicy.trname", "policyid", "1"),
                    resource.TestCheckResourceAttr("fortios_firewall_localinpolicy.trname", "schedule", "always"),
                    resource.TestCheckResourceAttr("fortios_firewall_localinpolicy.trname", "status", "enable"),
                    resource.TestCheckResourceAttr("fortios_firewall_localinpolicy.trname", "dstaddr.0.name", "all"),
                    resource.TestCheckResourceAttr("fortios_firewall_localinpolicy.trname", "service.0.name", "ALL"),
                    resource.TestCheckResourceAttr("fortios_firewall_localinpolicy.trname", "srcaddr.0.name", "all"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSFirewallLocalInPolicyExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallLocalInPolicy: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallLocalInPolicy is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallLocalInPolicy(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallLocalInPolicy: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallLocalInPolicy: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallLocalInPolicyDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_localinpolicy" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallLocalInPolicy(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallLocalInPolicy %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallLocalInPolicyConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_localinpolicy" "trname" {
  action            = "accept"
  ha_mgmt_intf_only = "disable"
  intf              = "port4"
  policyid          = 1
  schedule          = "always"
  status            = "enable"

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
