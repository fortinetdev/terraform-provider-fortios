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

func TestAccFortiOSWirelessControllerHotspot20AnqpNetworkAuthType_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSWirelessControllerHotspot20AnqpNetworkAuthType_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSWirelessControllerHotspot20AnqpNetworkAuthTypeConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSWirelessControllerHotspot20AnqpNetworkAuthTypeExists("fortios_wirelesscontrollerhotspot20_anqpnetworkauthtype.trname"),
					resource.TestCheckResourceAttr("fortios_wirelesscontrollerhotspot20_anqpnetworkauthtype.trname", "auth_type", "acceptance-of-terms"),
					resource.TestCheckResourceAttr("fortios_wirelesscontrollerhotspot20_anqpnetworkauthtype.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_wirelesscontrollerhotspot20_anqpnetworkauthtype.trname", "url", "www.example.com"),
				),
			},
		},
	})
}

func testAccCheckFortiOSWirelessControllerHotspot20AnqpNetworkAuthTypeExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found WirelessControllerHotspot20AnqpNetworkAuthType: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No WirelessControllerHotspot20AnqpNetworkAuthType is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadWirelessControllerHotspot20AnqpNetworkAuthType(i)

		if err != nil {
			return fmt.Errorf("Error reading WirelessControllerHotspot20AnqpNetworkAuthType: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating WirelessControllerHotspot20AnqpNetworkAuthType: %s", n)
		}

		return nil
	}
}

func testAccCheckWirelessControllerHotspot20AnqpNetworkAuthTypeDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_wirelesscontrollerhotspot20_anqpnetworkauthtype" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadWirelessControllerHotspot20AnqpNetworkAuthType(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error WirelessControllerHotspot20AnqpNetworkAuthType %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSWirelessControllerHotspot20AnqpNetworkAuthTypeConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_wirelesscontrollerhotspot20_anqpnetworkauthtype" "trname" {
  auth_type = "acceptance-of-terms"
  name      = "%[1]s"
  url       = "www.example.com"
}

`, name)
}
