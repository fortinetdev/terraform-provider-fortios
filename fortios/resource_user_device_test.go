// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

package fortios

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccFortiOSUserDevice_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSUserDevice_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSUserDeviceConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSUserDeviceExists("fortios_user_device.trname"),
					resource.TestCheckResourceAttr("fortios_user_device.trname", "alias", "1"),
					resource.TestCheckResourceAttr("fortios_user_device.trname", "category", "amazon-device"),
					resource.TestCheckResourceAttr("fortios_user_device.trname", "mac", "08:00:20:0a:8c:6d"),
					resource.TestCheckResourceAttr("fortios_user_device.trname", "type", "unknown"),
				),
			},
		},
	})
}

func testAccCheckFortiOSUserDeviceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found UserDevice: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No UserDevice is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadUserDevice(i, "root")

		if err != nil {
			return fmt.Errorf("Error reading UserDevice: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating UserDevice: %s", n)
		}

		return nil
	}
}

func testAccCheckUserDeviceDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_user_device" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadUserDevice(i, "root")

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error UserDevice %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSUserDeviceConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_user_device" "trname" {
  alias    = "1"
  category = "amazon-device"
  mac      = "08:00:20:0a:8c:6d"
  type     = "unknown"
}
`)
}
