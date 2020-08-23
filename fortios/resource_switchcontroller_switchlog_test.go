
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

func TestAccFortiOSSwitchControllerSwitchLog_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSSwitchControllerSwitchLog_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSSwitchControllerSwitchLogConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSSwitchControllerSwitchLogExists("fortios_switchcontroller_switchlog.trname"),
                    resource.TestCheckResourceAttr("fortios_switchcontroller_switchlog.trname", "severity", "critical"),
                    resource.TestCheckResourceAttr("fortios_switchcontroller_switchlog.trname", "status", "enable"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSSwitchControllerSwitchLogExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SwitchControllerSwitchLog: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SwitchControllerSwitchLog is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSwitchControllerSwitchLog(i)

		if err != nil {
			return fmt.Errorf("Error reading SwitchControllerSwitchLog: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SwitchControllerSwitchLog: %s", n)
		}

		return nil
	}
}

func testAccCheckSwitchControllerSwitchLogDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_switchcontroller_switchlog" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSwitchControllerSwitchLog(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SwitchControllerSwitchLog %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSwitchControllerSwitchLogConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_switchcontroller_switchlog" "trname" {
  severity = "critical"
  status   = "enable"
}


`)
}
