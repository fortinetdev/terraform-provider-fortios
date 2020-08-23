
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

func TestAccFortiOSExtenderControllerExtender_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSExtenderControllerExtender_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSExtenderControllerExtenderConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSExtenderControllerExtenderExists("fortios_extendercontroller_extender.trname"),
                    resource.TestCheckResourceAttr("fortios_extendercontroller_extender.trname", "admin", "disable"),
                    resource.TestCheckResourceAttr("fortios_extendercontroller_extender.trname", "billing_start_day", "1"),
                    resource.TestCheckResourceAttr("fortios_extendercontroller_extender.trname", "conn_status", "0"),
                    resource.TestCheckResourceAttr("fortios_extendercontroller_extender.trname", "dial_mode", "always-connect"),
                    resource.TestCheckResourceAttr("fortios_extendercontroller_extender.trname", "dial_status", "0"),
                    resource.TestCheckResourceAttr("fortios_extendercontroller_extender.trname", "ext_name", "332"),
                    resource.TestCheckResourceAttr("fortios_extendercontroller_extender.trname", "fosid", "1"),
                    resource.TestCheckResourceAttr("fortios_extendercontroller_extender.trname", "initiated_update", "disable"),
                    resource.TestCheckResourceAttr("fortios_extendercontroller_extender.trname", "mode", "standalone"),
                    resource.TestCheckResourceAttr("fortios_extendercontroller_extender.trname", "modem_type", "gsm/lte"),
                    resource.TestCheckResourceAttr("fortios_extendercontroller_extender.trname", "multi_mode", "auto"),
                    resource.TestCheckResourceAttr("fortios_extendercontroller_extender.trname", "ppp_auth_protocol", "auto"),
                    resource.TestCheckResourceAttr("fortios_extendercontroller_extender.trname", "ppp_echo_request", "disable"),
                    resource.TestCheckResourceAttr("fortios_extendercontroller_extender.trname", "quota_limit_mb", "0"),
                    resource.TestCheckResourceAttr("fortios_extendercontroller_extender.trname", "redial", "none"),
                    resource.TestCheckResourceAttr("fortios_extendercontroller_extender.trname", "roaming", "disable"),
                    resource.TestCheckResourceAttr("fortios_extendercontroller_extender.trname", "role", "primary"),
                    resource.TestCheckResourceAttr("fortios_extendercontroller_extender.trname", "vdom", "0"),
                    resource.TestCheckResourceAttr("fortios_extendercontroller_extender.trname", "wimax_auth_protocol", "tls"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSExtenderControllerExtenderExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found ExtenderControllerExtender: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ExtenderControllerExtender is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadExtenderControllerExtender(i)

		if err != nil {
			return fmt.Errorf("Error reading ExtenderControllerExtender: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating ExtenderControllerExtender: %s", n)
		}

		return nil
	}
}

func testAccCheckExtenderControllerExtenderDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_extendercontroller_extender" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadExtenderControllerExtender(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error ExtenderControllerExtender %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSExtenderControllerExtenderConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_extendercontroller_extender" "trname" {
  admin               = "disable"
  billing_start_day   = 1
  conn_status         = 0
  dial_mode           = "always-connect"
  dial_status         = 0
  ext_name            = "332"
  fosid               = "1"
  initiated_update    = "disable"
  mode                = "standalone"
  modem_type          = "gsm/lte"
  multi_mode          = "auto"
  ppp_auth_protocol   = "auto"
  ppp_echo_request    = "disable"
  quota_limit_mb      = 0
  redial              = "none"
  roaming             = "disable"
  role                = "primary"
  vdom                = 0
  wimax_auth_protocol = "tls"
}



`)
}
