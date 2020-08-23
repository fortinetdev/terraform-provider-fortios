// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

package fortios

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"log"
	"testing"
)

func TestAccFortiOSWirelessControllerHotspot20AnqpRoamingConsortium_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSWirelessControllerHotspot20AnqpRoamingConsortium_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSWirelessControllerHotspot20AnqpRoamingConsortiumConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSWirelessControllerHotspot20AnqpRoamingConsortiumExists("fortios_wirelesscontrollerhotspot20_anqproamingconsortium.trname"),
					resource.TestCheckResourceAttr("fortios_wirelesscontrollerhotspot20_anqproamingconsortium.trname", "name", rname),
				),
			},
		},
	})
}

func testAccCheckFortiOSWirelessControllerHotspot20AnqpRoamingConsortiumExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found WirelessControllerHotspot20AnqpRoamingConsortium: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No WirelessControllerHotspot20AnqpRoamingConsortium is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadWirelessControllerHotspot20AnqpRoamingConsortium(i)

		if err != nil {
			return fmt.Errorf("Error reading WirelessControllerHotspot20AnqpRoamingConsortium: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating WirelessControllerHotspot20AnqpRoamingConsortium: %s", n)
		}

		return nil
	}
}

func testAccCheckWirelessControllerHotspot20AnqpRoamingConsortiumDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_wirelesscontrollerhotspot20_anqproamingconsortium" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadWirelessControllerHotspot20AnqpRoamingConsortium(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error WirelessControllerHotspot20AnqpRoamingConsortium %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSWirelessControllerHotspot20AnqpRoamingConsortiumConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_wirelesscontrollerhotspot20_anqproamingconsortium" "trname" {
  name = "%[1]s"
}
`, name)
}
