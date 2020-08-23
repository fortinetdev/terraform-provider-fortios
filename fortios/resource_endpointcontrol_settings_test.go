
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

func TestAccFortiOSEndpointControlSettings_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSEndpointControlSettings_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSEndpointControlSettingsConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSEndpointControlSettingsExists("fortios_endpointcontrol_settings.trname"),
                    resource.TestCheckResourceAttr("fortios_endpointcontrol_settings.trname", "download_location", "fortiguard"),
                    resource.TestCheckResourceAttr("fortios_endpointcontrol_settings.trname", "forticlient_avdb_update_interval", "8"),
                    resource.TestCheckResourceAttr("fortios_endpointcontrol_settings.trname", "forticlient_dereg_unsupported_client", "enable"),
                    resource.TestCheckResourceAttr("fortios_endpointcontrol_settings.trname", "forticlient_ems_rest_api_call_timeout", "5000"),
                    resource.TestCheckResourceAttr("fortios_endpointcontrol_settings.trname", "forticlient_keepalive_interval", "60"),
                    resource.TestCheckResourceAttr("fortios_endpointcontrol_settings.trname", "forticlient_offline_grace", "disable"),
                    resource.TestCheckResourceAttr("fortios_endpointcontrol_settings.trname", "forticlient_offline_grace_interval", "120"),
                    resource.TestCheckResourceAttr("fortios_endpointcontrol_settings.trname", "forticlient_reg_key_enforce", "disable"),
                    resource.TestCheckResourceAttr("fortios_endpointcontrol_settings.trname", "forticlient_reg_timeout", "7"),
                    resource.TestCheckResourceAttr("fortios_endpointcontrol_settings.trname", "forticlient_sys_update_interval", "720"),
                    resource.TestCheckResourceAttr("fortios_endpointcontrol_settings.trname", "forticlient_user_avatar", "enable"),
                    resource.TestCheckResourceAttr("fortios_endpointcontrol_settings.trname", "forticlient_warning_interval", "1"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSEndpointControlSettingsExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found EndpointControlSettings: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No EndpointControlSettings is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadEndpointControlSettings(i)

		if err != nil {
			return fmt.Errorf("Error reading EndpointControlSettings: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating EndpointControlSettings: %s", n)
		}

		return nil
	}
}

func testAccCheckEndpointControlSettingsDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_endpointcontrol_settings" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadEndpointControlSettings(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error EndpointControlSettings %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSEndpointControlSettingsConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_endpointcontrol_settings" "trname" {
  download_location                     = "fortiguard"
  forticlient_avdb_update_interval      = 8
  forticlient_dereg_unsupported_client  = "enable"
  forticlient_ems_rest_api_call_timeout = 5000
  forticlient_keepalive_interval        = 60
  forticlient_offline_grace             = "disable"
  forticlient_offline_grace_interval    = 120
  forticlient_reg_key_enforce           = "disable"
  forticlient_reg_timeout               = 7
  forticlient_sys_update_interval       = 720
  forticlient_user_avatar               = "enable"
  forticlient_warning_interval          = 1
}
`)
}
