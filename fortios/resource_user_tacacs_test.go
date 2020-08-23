
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

func TestAccFortiOSUserTacacs_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSUserTacacs_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSUserTacacsConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSUserTacacsExists("fortios_user_tacacs.trname"),
                    resource.TestCheckResourceAttr("fortios_user_tacacs.trname", "authen_type", "auto"),
                    resource.TestCheckResourceAttr("fortios_user_tacacs.trname", "authorization", "disable"),
                    resource.TestCheckResourceAttr("fortios_user_tacacs.trname", "name", rname),
                    resource.TestCheckResourceAttr("fortios_user_tacacs.trname", "port", "2342"),
                    resource.TestCheckResourceAttr("fortios_user_tacacs.trname", "server", "1.1.1.1"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSUserTacacsExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found UserTacacs: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No UserTacacs is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadUserTacacs(i)

		if err != nil {
			return fmt.Errorf("Error reading UserTacacs: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating UserTacacs: %s", n)
		}

		return nil
	}
}

func testAccCheckUserTacacsDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_user_tacacs" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadUserTacacs(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error UserTacacs %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSUserTacacsConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_user_tacacs" "trname" {
  authen_type   = "auto"
  authorization = "disable"
  name          = "%[1]s"
  port          = 2342
  server        = "1.1.1.1"
}

`, name)
}
