
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

func TestAccFortiOSSwitchController8021XSettings_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSSwitchController8021XSettings_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSSwitchController8021XSettingsConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSSwitchController8021XSettingsExists("fortios_switchcontroller_8021Xsettings.trname"),
                    resource.TestCheckResourceAttr("fortios_switchcontroller_8021Xsettings.trname", "link_down_auth", "set-unauth"),
                    resource.TestCheckResourceAttr("fortios_switchcontroller_8021Xsettings.trname", "max_reauth_attempt", "3"),
                    resource.TestCheckResourceAttr("fortios_switchcontroller_8021Xsettings.trname", "reauth_period", "12"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSSwitchController8021XSettingsExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SwitchController8021XSettings: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SwitchController8021XSettings is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSwitchController8021XSettings(i)

		if err != nil {
			return fmt.Errorf("Error reading SwitchController8021XSettings: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SwitchController8021XSettings: %s", n)
		}

		return nil
	}
}

func testAccCheckSwitchController8021XSettingsDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_switchcontroller_8021Xsettings" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSwitchController8021XSettings(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SwitchController8021XSettings %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSwitchController8021XSettingsConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_switchcontroller_8021Xsettings" "trname" {
  link_down_auth     = "set-unauth"
  max_reauth_attempt = 3
  reauth_period      = 12
}
`)
}
