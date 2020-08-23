
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

func TestAccFortiOSFirewallLdbMonitor_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSFirewallLdbMonitor_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSFirewallLdbMonitorConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSFirewallLdbMonitorExists("fortios_firewall_ldbmonitor.trname"),
                    resource.TestCheckResourceAttr("fortios_firewall_ldbmonitor.trname", "http_max_redirects", "0"),
                    resource.TestCheckResourceAttr("fortios_firewall_ldbmonitor.trname", "interval", "10"),
                    resource.TestCheckResourceAttr("fortios_firewall_ldbmonitor.trname", "name", rname),
                    resource.TestCheckResourceAttr("fortios_firewall_ldbmonitor.trname", "port", "0"),
                    resource.TestCheckResourceAttr("fortios_firewall_ldbmonitor.trname", "retry", "3"),
                    resource.TestCheckResourceAttr("fortios_firewall_ldbmonitor.trname", "timeout", "2"),
                    resource.TestCheckResourceAttr("fortios_firewall_ldbmonitor.trname", "type", "ping"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSFirewallLdbMonitorExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallLdbMonitor: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallLdbMonitor is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallLdbMonitor(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallLdbMonitor: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallLdbMonitor: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallLdbMonitorDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_ldbmonitor" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallLdbMonitor(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallLdbMonitor %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallLdbMonitorConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_ldbmonitor" "trname" {
  http_max_redirects = 0
  interval           = 10
  name               = "%[1]s"
  port               = 0
  retry              = 3
  timeout            = 2
  type               = "ping"
}





`, name)
}
