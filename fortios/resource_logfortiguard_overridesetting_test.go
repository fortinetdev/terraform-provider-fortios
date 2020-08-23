
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

func TestAccFortiOSLogFortiguardOverrideSetting_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSLogFortiguardOverrideSetting_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSLogFortiguardOverrideSettingConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSLogFortiguardOverrideSettingExists("fortios_logfortiguard_overridesetting.trname"),
                    resource.TestCheckResourceAttr("fortios_logfortiguard_overridesetting.trname", "override", "disable"),
                    resource.TestCheckResourceAttr("fortios_logfortiguard_overridesetting.trname", "status", "disable"),
                    resource.TestCheckResourceAttr("fortios_logfortiguard_overridesetting.trname", "upload_interval", "daily"),
                    resource.TestCheckResourceAttr("fortios_logfortiguard_overridesetting.trname", "upload_option", "5-minute"),
                    resource.TestCheckResourceAttr("fortios_logfortiguard_overridesetting.trname", "upload_time", "00:00"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSLogFortiguardOverrideSettingExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found LogFortiguardOverrideSetting: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No LogFortiguardOverrideSetting is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadLogFortiguardOverrideSetting(i)

		if err != nil {
			return fmt.Errorf("Error reading LogFortiguardOverrideSetting: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating LogFortiguardOverrideSetting: %s", n)
		}

		return nil
	}
}

func testAccCheckLogFortiguardOverrideSettingDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_logfortiguard_overridesetting" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadLogFortiguardOverrideSetting(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error LogFortiguardOverrideSetting %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSLogFortiguardOverrideSettingConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_logfortiguard_overridesetting" "trname" {
  override        = "disable"
  status          = "disable"
  upload_interval = "daily"
  upload_option   = "5-minute"
  upload_time     = "00:00"
}
`)
}
