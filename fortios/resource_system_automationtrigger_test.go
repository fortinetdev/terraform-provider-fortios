
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

func TestAccFortiOSSystemAutomationTrigger_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSSystemAutomationTrigger_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSSystemAutomationTriggerConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSSystemAutomationTriggerExists("fortios_system_automationtrigger.trname"),
                    resource.TestCheckResourceAttr("fortios_system_automationtrigger.trname", "event_type", "event-log"),
                    resource.TestCheckResourceAttr("fortios_system_automationtrigger.trname", "ioc_level", "high"),
                    resource.TestCheckResourceAttr("fortios_system_automationtrigger.trname", "license_type", "forticare-support"),
                    resource.TestCheckResourceAttr("fortios_system_automationtrigger.trname", "logid", "32002"),
                    resource.TestCheckResourceAttr("fortios_system_automationtrigger.trname", "name", rname),
                    resource.TestCheckResourceAttr("fortios_system_automationtrigger.trname", "trigger_frequency", "daily"),
                    resource.TestCheckResourceAttr("fortios_system_automationtrigger.trname", "trigger_minute", "60"),
                    resource.TestCheckResourceAttr("fortios_system_automationtrigger.trname", "trigger_type", "event-based"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSSystemAutomationTriggerExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemAutomationTrigger: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemAutomationTrigger is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemAutomationTrigger(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemAutomationTrigger: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemAutomationTrigger: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemAutomationTriggerDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_automationtrigger" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemAutomationTrigger(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemAutomationTrigger %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemAutomationTriggerConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_automationtrigger" "trname" {
  event_type        = "event-log"
  ioc_level         = "high"
  license_type      = "forticare-support"
  logid             = 32002
  name              = "%[1]s"
  trigger_frequency = "daily"
  trigger_minute    = 60
  trigger_type      = "event-based"
}
`, name)
}
