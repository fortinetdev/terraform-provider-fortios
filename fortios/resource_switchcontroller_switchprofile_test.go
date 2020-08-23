
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

func TestAccFortiOSSwitchControllerSwitchProfile_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSSwitchControllerSwitchProfile_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSSwitchControllerSwitchProfileConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSSwitchControllerSwitchProfileExists("fortios_switchcontroller_switchprofile.trname"),
                    resource.TestCheckResourceAttr("fortios_switchcontroller_switchprofile.trname", "login_passwd_override", "enable"),
                    resource.TestCheckResourceAttr("fortios_switchcontroller_switchprofile.trname", "name", rname),
                ),
            },
        },
    })
}

func testAccCheckFortiOSSwitchControllerSwitchProfileExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SwitchControllerSwitchProfile: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SwitchControllerSwitchProfile is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSwitchControllerSwitchProfile(i)

		if err != nil {
			return fmt.Errorf("Error reading SwitchControllerSwitchProfile: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SwitchControllerSwitchProfile: %s", n)
		}

		return nil
	}
}

func testAccCheckSwitchControllerSwitchProfileDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_switchcontroller_switchprofile" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSwitchControllerSwitchProfile(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SwitchControllerSwitchProfile %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSwitchControllerSwitchProfileConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_switchcontroller_switchprofile" "trname" {
  login_passwd_override = "enable"
  name                  = "%[1]s"
}
`, name)
}
