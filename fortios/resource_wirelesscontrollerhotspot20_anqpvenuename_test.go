
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

func TestAccFortiOSWirelessControllerHotspot20AnqpVenueName_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSWirelessControllerHotspot20AnqpVenueName_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSWirelessControllerHotspot20AnqpVenueNameConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSWirelessControllerHotspot20AnqpVenueNameExists("fortios_wirelesscontrollerhotspot20_anqpvenuename.trname"),
                    resource.TestCheckResourceAttr("fortios_wirelesscontrollerhotspot20_anqpvenuename.trname", "name", rname),
                    resource.TestCheckResourceAttr("fortios_wirelesscontrollerhotspot20_anqpvenuename.trname", "value_list.0.index", "1"),
                    resource.TestCheckResourceAttr("fortios_wirelesscontrollerhotspot20_anqpvenuename.trname", "value_list.0.lang", "CN"),
                    resource.TestCheckResourceAttr("fortios_wirelesscontrollerhotspot20_anqpvenuename.trname", "value_list.0.value", "3"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSWirelessControllerHotspot20AnqpVenueNameExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found WirelessControllerHotspot20AnqpVenueName: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No WirelessControllerHotspot20AnqpVenueName is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadWirelessControllerHotspot20AnqpVenueName(i)

		if err != nil {
			return fmt.Errorf("Error reading WirelessControllerHotspot20AnqpVenueName: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating WirelessControllerHotspot20AnqpVenueName: %s", n)
		}

		return nil
	}
}

func testAccCheckWirelessControllerHotspot20AnqpVenueNameDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_wirelesscontrollerhotspot20_anqpvenuename" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadWirelessControllerHotspot20AnqpVenueName(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error WirelessControllerHotspot20AnqpVenueName %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSWirelessControllerHotspot20AnqpVenueNameConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_wirelesscontrollerhotspot20_anqpvenuename" "trname" {
  name = "%[1]s"
  value_list {
    index = 1
    lang  = "CN"
    value = "3"
  }
}

`, name)
}
