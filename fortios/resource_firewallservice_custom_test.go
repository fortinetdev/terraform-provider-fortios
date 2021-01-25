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

func TestAccFortiOSFirewallServiceCustom_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSFirewallServiceCustom_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallServiceCustomConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallServiceCustomExists("fortios_firewallservice_custom.trname"),
					resource.TestCheckResourceAttr("fortios_firewallservice_custom.trname", "app_service_type", "disable"),
					resource.TestCheckResourceAttr("fortios_firewallservice_custom.trname", "category", "General"),
					resource.TestCheckResourceAttr("fortios_firewallservice_custom.trname", "check_reset_range", "default"),
					resource.TestCheckResourceAttr("fortios_firewallservice_custom.trname", "color", "0"),
					resource.TestCheckResourceAttr("fortios_firewallservice_custom.trname", "helper", "auto"),
					resource.TestCheckResourceAttr("fortios_firewallservice_custom.trname", "iprange", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_firewallservice_custom.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewallservice_custom.trname", "protocol", "TCP/UDP/SCTP"),
					resource.TestCheckResourceAttr("fortios_firewallservice_custom.trname", "protocol_number", "6"),
					resource.TestCheckResourceAttr("fortios_firewallservice_custom.trname", "proxy", "disable"),
					resource.TestCheckResourceAttr("fortios_firewallservice_custom.trname", "tcp_halfclose_timer", "0"),
					resource.TestCheckResourceAttr("fortios_firewallservice_custom.trname", "tcp_halfopen_timer", "0"),
					resource.TestCheckResourceAttr("fortios_firewallservice_custom.trname", "tcp_portrange", "223-332"),
					resource.TestCheckResourceAttr("fortios_firewallservice_custom.trname", "tcp_timewait_timer", "0"),
					resource.TestCheckResourceAttr("fortios_firewallservice_custom.trname", "udp_idle_timer", "0"),
					resource.TestCheckResourceAttr("fortios_firewallservice_custom.trname", "visibility", "enable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallServiceCustomExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallServiceCustom: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallServiceCustom is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallServiceCustom(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallServiceCustom: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallServiceCustom: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallServiceCustomDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewallservice_custom" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallServiceCustom(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallServiceCustom %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallServiceCustomConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewallservice_custom" "trname" {
  app_service_type    = "disable"
  category            = "General"
  check_reset_range   = "default"
  color               = 0
  helper              = "auto"
  iprange             = "0.0.0.0"
  name                = "%[1]s"
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
`, name)
}
