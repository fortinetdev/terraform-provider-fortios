
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

func TestAccFortiOSFirewallScheduleOnetime_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSFirewallScheduleOnetime_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSFirewallScheduleOnetimeConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSFirewallScheduleOnetimeExists("fortios_firewallschedule_onetime.trname"),
                    resource.TestCheckResourceAttr("fortios_firewallschedule_onetime.trname", "color", "0"),
                    resource.TestCheckResourceAttr("fortios_firewallschedule_onetime.trname", "end", "00:00 2020/12/12"),
                    resource.TestCheckResourceAttr("fortios_firewallschedule_onetime.trname", "expiration_days", "2"),
                    resource.TestCheckResourceAttr("fortios_firewallschedule_onetime.trname", "name", rname),
                    resource.TestCheckResourceAttr("fortios_firewallschedule_onetime.trname", "start", "00:00 2010/12/12"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSFirewallScheduleOnetimeExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallScheduleOnetime: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallScheduleOnetime is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallScheduleOnetime(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallScheduleOnetime: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallScheduleOnetime: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallScheduleOnetimeDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewallschedule_onetime" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallScheduleOnetime(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallScheduleOnetime %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallScheduleOnetimeConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewallschedule_onetime" "trname" {
  color           = 0
  end             = "00:00 2020/12/12"
  expiration_days = 2
  name            = "%[1]s"
  start           = "00:00 2010/12/12"
}
`, name)
}
