// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

package fortios

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"log"
	"testing"
)

func TestAccFortiOSWirelessControllerHotspot20AnqpIpAddressType_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSWirelessControllerHotspot20AnqpIpAddressType_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSWirelessControllerHotspot20AnqpIpAddressTypeConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSWirelessControllerHotspot20AnqpIpAddressTypeExists("fortios_wirelesscontrollerhotspot20_anqpipaddresstype.trname"),
					resource.TestCheckResourceAttr("fortios_wirelesscontrollerhotspot20_anqpipaddresstype.trname", "ipv4_address_type", "public"),
					resource.TestCheckResourceAttr("fortios_wirelesscontrollerhotspot20_anqpipaddresstype.trname", "ipv6_address_type", "not-available"),
					resource.TestCheckResourceAttr("fortios_wirelesscontrollerhotspot20_anqpipaddresstype.trname", "name", rname),
				),
			},
		},
	})
}

func testAccCheckFortiOSWirelessControllerHotspot20AnqpIpAddressTypeExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found WirelessControllerHotspot20AnqpIpAddressType: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No WirelessControllerHotspot20AnqpIpAddressType is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadWirelessControllerHotspot20AnqpIpAddressType(i)

		if err != nil {
			return fmt.Errorf("Error reading WirelessControllerHotspot20AnqpIpAddressType: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating WirelessControllerHotspot20AnqpIpAddressType: %s", n)
		}

		return nil
	}
}

func testAccCheckWirelessControllerHotspot20AnqpIpAddressTypeDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_wirelesscontrollerhotspot20_anqpipaddresstype" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadWirelessControllerHotspot20AnqpIpAddressType(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error WirelessControllerHotspot20AnqpIpAddressType %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSWirelessControllerHotspot20AnqpIpAddressTypeConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_wirelesscontrollerhotspot20_anqpipaddresstype" "trname" {
  ipv4_address_type = "public"
  ipv6_address_type = "not-available"
  name              = "%[1]s"
}
`, name)
}
