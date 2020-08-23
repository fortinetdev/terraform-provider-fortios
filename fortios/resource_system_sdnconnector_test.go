
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

func TestAccFortiOSSystemSdnConnector_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSSystemSdnConnector_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSSystemSdnConnectorConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSSystemSdnConnectorExists("fortios_system_sdnconnector.trname"),
                    resource.TestCheckResourceAttr("fortios_system_sdnconnector.trname", "azure_region", "global"),
                    resource.TestCheckResourceAttr("fortios_system_sdnconnector.trname", "ha_status", "disable"),
                    resource.TestCheckResourceAttr("fortios_system_sdnconnector.trname", "name", rname),
                    resource.TestCheckResourceAttr("fortios_system_sdnconnector.trname", "password", "ENC -1N2ZPOCHEkvUA4pNeO78iotUQKN8="),
                    resource.TestCheckResourceAttr("fortios_system_sdnconnector.trname", "server", "1.1.1.1"),
                    resource.TestCheckResourceAttr("fortios_system_sdnconnector.trname", "server_port", "3"),
                    resource.TestCheckResourceAttr("fortios_system_sdnconnector.trname", "status", "disable"),
                    resource.TestCheckResourceAttr("fortios_system_sdnconnector.trname", "type", "aci"),
                    resource.TestCheckResourceAttr("fortios_system_sdnconnector.trname", "update_interval", "60"),
                    resource.TestCheckResourceAttr("fortios_system_sdnconnector.trname", "use_metadata_iam", "disable"),
                    resource.TestCheckResourceAttr("fortios_system_sdnconnector.trname", "username", "sg"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSSystemSdnConnectorExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemSdnConnector: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemSdnConnector is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemSdnConnector(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemSdnConnector: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemSdnConnector: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemSdnConnectorDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_sdnconnector" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemSdnConnector(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemSdnConnector %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemSdnConnectorConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_sdnconnector" "trname" {
  azure_region     = "global"
  ha_status        = "disable"
  name             = "%[1]s"
  password         = "ENC -1N2ZPOCHEkvUA4pNeO78iotUQKN8="
  server           = "1.1.1.1"
  server_port      = 3
  status           = "disable"
  type             = "aci"
  update_interval  = 60
  use_metadata_iam = "disable"
  username         = "sg"
}

`, name)
}
