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

func TestAccFortiOSUserDeviceAccessList_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSUserDeviceAccessList_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSUserDeviceAccessListConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSUserDeviceAccessListExists("fortios_user_deviceaccesslist.trname"),
					resource.TestCheckResourceAttr("fortios_user_deviceaccesslist.trname", "default_action", "accept"),
					resource.TestCheckResourceAttr("fortios_user_deviceaccesslist.trname", "name", rname),
				),
			},
		},
	})
}

func testAccCheckFortiOSUserDeviceAccessListExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found UserDeviceAccessList: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No UserDeviceAccessList is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadUserDeviceAccessList(i, "root")

		if err != nil {
			return fmt.Errorf("Error reading UserDeviceAccessList: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating UserDeviceAccessList: %s", n)
		}

		return nil
	}
}

func testAccCheckUserDeviceAccessListDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_user_deviceaccesslist" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadUserDeviceAccessList(i, "root")

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error UserDeviceAccessList %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSUserDeviceAccessListConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_user_deviceaccesslist" "trname" {
  default_action = "accept"
  name           = "%[1]s"
}
`, name)
}
