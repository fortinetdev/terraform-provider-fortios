
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

func TestAccFortiOSWirelessControllerHotspot20Anqp3GppCellular_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSWirelessControllerHotspot20Anqp3GppCellular_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSWirelessControllerHotspot20Anqp3GppCellularConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSWirelessControllerHotspot20Anqp3GppCellularExists("fortios_wirelesscontrollerhotspot20_anqp3gppcellular.trname"),
                    resource.TestCheckResourceAttr("fortios_wirelesscontrollerhotspot20_anqp3gppcellular.trname", "name", rname),
                ),
            },
        },
    })
}

func testAccCheckFortiOSWirelessControllerHotspot20Anqp3GppCellularExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found WirelessControllerHotspot20Anqp3GppCellular: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No WirelessControllerHotspot20Anqp3GppCellular is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadWirelessControllerHotspot20Anqp3GppCellular(i)

		if err != nil {
			return fmt.Errorf("Error reading WirelessControllerHotspot20Anqp3GppCellular: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating WirelessControllerHotspot20Anqp3GppCellular: %s", n)
		}

		return nil
	}
}

func testAccCheckWirelessControllerHotspot20Anqp3GppCellularDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_wirelesscontrollerhotspot20_anqp3gppcellular" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadWirelessControllerHotspot20Anqp3GppCellular(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error WirelessControllerHotspot20Anqp3GppCellular %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSWirelessControllerHotspot20Anqp3GppCellularConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_wirelesscontrollerhotspot20_anqp3gppcellular" "trname" {
  name = "%[1]s"
}
`, name)
}
