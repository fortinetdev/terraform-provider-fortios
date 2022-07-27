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

func TestAccFortiOSFirewallAddress6Template_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSFirewallAddress6Template_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallAddress6TemplateConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallAddress6TemplateExists("fortios_firewall_address6template.trname"),
					resource.TestCheckResourceAttr("fortios_firewall_address6template.trname", "ip6", "2001:db8:0:b::/64"),
					resource.TestCheckResourceAttr("fortios_firewall_address6template.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewall_address6template.trname", "subnet_segment_count", "2"),
					resource.TestCheckResourceAttr("fortios_firewall_address6template.trname", "subnet_segment.0.bits", "4"),
					resource.TestCheckResourceAttr("fortios_firewall_address6template.trname", "subnet_segment.0.exclusive", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_address6template.trname", "subnet_segment.0.id", "1"),
					resource.TestCheckResourceAttr("fortios_firewall_address6template.trname", "subnet_segment.0.name", "country"),
					resource.TestCheckResourceAttr("fortios_firewall_address6template.trname", "subnet_segment.1.bits", "4"),
					resource.TestCheckResourceAttr("fortios_firewall_address6template.trname", "subnet_segment.1.exclusive", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_address6template.trname", "subnet_segment.1.id", "2"),
					resource.TestCheckResourceAttr("fortios_firewall_address6template.trname", "subnet_segment.1.name", "state"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallAddress6TemplateExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallAddress6Template: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallAddress6Template is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallAddress6Template(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallAddress6Template: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallAddress6Template: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallAddress6TemplateDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_address6template" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallAddress6Template(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallAddress6Template %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallAddress6TemplateConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_address6template" "trname" {
  ip6                  = "2001:db8:0:b::/64"
  name                 = "%[1]s"
  subnet_segment_count = 2

  subnet_segment {
    bits      = 4
    exclusive = "disable"
    id        = 1
    name      = "country"
  }
  subnet_segment {
    bits      = 4
    exclusive = "disable"
    id        = 2
    name      = "state"
  }
}
`, name)
}
