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

func TestAccFortiOSFirewallServiceGroup_basic(t *testing.T) {
	rname := acctest.RandString(8)
	var0 := "var0" + rname
	log.Printf(var0)
	log.Printf("TestAccFortiOSFirewallServiceGroup_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallServiceGroupConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallServiceGroupExists("fortios_firewallservice_group.trname"),
					resource.TestCheckResourceAttr("fortios_firewallservice_group.trname", "color", "0"),
					resource.TestCheckResourceAttr("fortios_firewallservice_group.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewallservice_group.trname", "proxy", "disable"),
					resource.TestCheckResourceAttr("fortios_firewallservice_group.trname", "member.0.name", var0),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallServiceGroupExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallServiceGroup: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallServiceGroup is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallServiceGroup(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallServiceGroup: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallServiceGroup: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallServiceGroupDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewallservice_group" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallServiceGroup(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallServiceGroup %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallServiceGroupConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewallservice_custom" "trname1" {
  app_service_type    = "disable"
  category            = "General"
  check_reset_range   = "default"
  color               = 0
  helper              = "auto"
  iprange             = "0.0.0.0"
  name                = "var0%[1]s"
  protocol            = "TCP/UDP/SCTP"
  protocol_number     = 6
  proxy               = "disable"
  tcp_halfclose_timer = 0
  tcp_halfopen_timer  = 0
  tcp_portrange       = "223-332"
  tcp_timewait_timer  = 0
  udp_idle_timer      = 0
  visibility          = "enable"
}

resource "fortios_firewallservice_group" "trname" {
  color = 0
  name  = "%[1]s"
  proxy = "disable"

  member {
    name = fortios_firewallservice_custom.trname1.name
  }
}
`, name)
}
