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

func TestAccFortiOSFirewallShapingPolicy_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSFirewallShapingPolicy_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallShapingPolicyConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallShapingPolicyExists("fortios_firewall_shapingpolicy.trname"),
					resource.TestCheckResourceAttr("fortios_firewall_shapingpolicy.trname", "class_id", "0"),
					resource.TestCheckResourceAttr("fortios_firewall_shapingpolicy.trname", "diffserv_forward", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_shapingpolicy.trname", "diffserv_reverse", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_shapingpolicy.trname", "diffservcode_forward", "000000"),
					resource.TestCheckResourceAttr("fortios_firewall_shapingpolicy.trname", "diffservcode_rev", "000000"),
					resource.TestCheckResourceAttr("fortios_firewall_shapingpolicy.trname", "fosid", "1"),
					resource.TestCheckResourceAttr("fortios_firewall_shapingpolicy.trname", "internet_service", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_shapingpolicy.trname", "internet_service_src", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_shapingpolicy.trname", "ip_version", "4"),
					resource.TestCheckResourceAttr("fortios_firewall_shapingpolicy.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewall_shapingpolicy.trname", "status", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_shapingpolicy.trname", "tos", "0x00"),
					resource.TestCheckResourceAttr("fortios_firewall_shapingpolicy.trname", "tos_mask", "0x00"),
					resource.TestCheckResourceAttr("fortios_firewall_shapingpolicy.trname", "tos_negate", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_shapingpolicy.trname", "dstaddr.0.name", "all"),
					resource.TestCheckResourceAttr("fortios_firewall_shapingpolicy.trname", "dstintf.0.name", "port4"),
					resource.TestCheckResourceAttr("fortios_firewall_shapingpolicy.trname", "service.0.name", "ALL"),
					resource.TestCheckResourceAttr("fortios_firewall_shapingpolicy.trname", "srcaddr.0.name", "all"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallShapingPolicyExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallShapingPolicy: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallShapingPolicy is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallShapingPolicy(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallShapingPolicy: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallShapingPolicy: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallShapingPolicyDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_shapingpolicy" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallShapingPolicy(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallShapingPolicy %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallShapingPolicyConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_shapingpolicy" "trname" {
  class_id             = 0
  diffserv_forward     = "disable"
  diffserv_reverse     = "disable"
  diffservcode_forward = "000000"
  diffservcode_rev     = "000000"
  fosid                = 1
  internet_service     = "disable"
  internet_service_src = "disable"
  ip_version           = "4"
  name                 = "%[1]s"
  status               = "enable"
  tos                  = "0x00"
  tos_mask             = "0x00"
  tos_negate           = "disable"

  dstaddr {
    name = "all"
  }

  dstintf {
    name = "port4"
  }

  service {
    name = "ALL"
  }

  srcaddr {
    name = "all"
  }
}
`, name)
}
