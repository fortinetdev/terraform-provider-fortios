
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

func TestAccFortiOSSystemFortimanager_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSSystemFortimanager_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSSystemFortimanagerConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSSystemFortimanagerExists("fortios_system_fortimanager.trname"),
                    resource.TestCheckResourceAttr("fortios_system_fortimanager.trname", "central_management", "disable"),
                    resource.TestCheckResourceAttr("fortios_system_fortimanager.trname", "central_mgmt_auto_backup", "disable"),
                    resource.TestCheckResourceAttr("fortios_system_fortimanager.trname", "central_mgmt_schedule_config_restore", "disable"),
                    resource.TestCheckResourceAttr("fortios_system_fortimanager.trname", "central_mgmt_schedule_script_restore", "disable"),
                    resource.TestCheckResourceAttr("fortios_system_fortimanager.trname", "ip", "0.0.0.0"),
                    resource.TestCheckResourceAttr("fortios_system_fortimanager.trname", "ipsec", "disable"),
                    resource.TestCheckResourceAttr("fortios_system_fortimanager.trname", "vdom", "root"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSSystemFortimanagerExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemFortimanager: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemFortimanager is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemFortimanager(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemFortimanager: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemFortimanager: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemFortimanagerDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_fortimanager" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemFortimanager(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemFortimanager %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemFortimanagerConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_fortimanager" "trname" {
  central_management                   = "disable"
  central_mgmt_auto_backup             = "disable"
  central_mgmt_schedule_config_restore = "disable"
  central_mgmt_schedule_script_restore = "disable"
  ip                                   = "0.0.0.0"
  ipsec                                = "disable"
  vdom                                 = "root"
}


`)
}
