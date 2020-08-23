
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

func TestAccFortiOSSystemCentralManagement_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSSystemCentralManagement_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSSystemCentralManagementConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSSystemCentralManagementExists("fortios_system_centralmanagement.trname"),
                    resource.TestCheckResourceAttr("fortios_system_centralmanagement.trname", "allow_monitor", "enable"),
                    resource.TestCheckResourceAttr("fortios_system_centralmanagement.trname", "allow_push_configuration", "enable"),
                    resource.TestCheckResourceAttr("fortios_system_centralmanagement.trname", "allow_push_firmware", "enable"),
                    resource.TestCheckResourceAttr("fortios_system_centralmanagement.trname", "allow_remote_firmware_upgrade", "enable"),
                    resource.TestCheckResourceAttr("fortios_system_centralmanagement.trname", "enc_algorithm", "high"),
                    resource.TestCheckResourceAttr("fortios_system_centralmanagement.trname", "fmg_source_ip", "0.0.0.0"),
                    resource.TestCheckResourceAttr("fortios_system_centralmanagement.trname", "fmg_source_ip6", "::"),
                    resource.TestCheckResourceAttr("fortios_system_centralmanagement.trname", "include_default_servers", "enable"),
                    resource.TestCheckResourceAttr("fortios_system_centralmanagement.trname", "mode", "normal"),
                    resource.TestCheckResourceAttr("fortios_system_centralmanagement.trname", "schedule_config_restore", "enable"),
                    resource.TestCheckResourceAttr("fortios_system_centralmanagement.trname", "schedule_script_restore", "enable"),
                    resource.TestCheckResourceAttr("fortios_system_centralmanagement.trname", "type", "fortimanager"),
                    resource.TestCheckResourceAttr("fortios_system_centralmanagement.trname", "vdom", "root"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSSystemCentralManagementExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemCentralManagement: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemCentralManagement is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemCentralManagement(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemCentralManagement: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemCentralManagement: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemCentralManagementDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_centralmanagement" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemCentralManagement(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemCentralManagement %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemCentralManagementConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_centralmanagement" "trname" {
  allow_monitor                 = "enable"
  allow_push_configuration      = "enable"
  allow_push_firmware           = "enable"
  allow_remote_firmware_upgrade = "enable"
  enc_algorithm                 = "high"
  fmg_source_ip                 = "0.0.0.0"
  fmg_source_ip6                = "::"
  include_default_servers       = "enable"
  mode                          = "normal"
  schedule_config_restore       = "enable"
  schedule_script_restore       = "enable"
  type                          = "fortimanager"
  vdom                          = "root"
}
`)
}
