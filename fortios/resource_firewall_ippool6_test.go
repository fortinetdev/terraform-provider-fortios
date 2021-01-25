// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

package fortios

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"log"
	"testing"
)

func TestAccFortiOSFirewallIppool6_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSFirewallIppool6_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallIppool6Config(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallIppool6Exists("fortios_firewall_ippool6.trname"),
					resource.TestCheckResourceAttr("fortios_firewall_ippool6.trname", "endip", "2001:3ca1:10f:1a:121b::19"),
					resource.TestCheckResourceAttr("fortios_firewall_ippool6.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewall_ippool6.trname", "startip", "2001:3ca1:10f:1a:121b::10"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallIppool6Exists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallIppool6: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallIppool6 is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallIppool6(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallIppool6: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallIppool6: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallIppool6Destroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_ippool6" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallIppool6(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallIppool6 %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallIppool6Config(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_ippool6" "trname" {
  endip   = "2001:3ca1:10f:1a:121b::19"
  name    = "%[1]s"
  startip = "2001:3ca1:10f:1a:121b::10"
}
`, name)
}
