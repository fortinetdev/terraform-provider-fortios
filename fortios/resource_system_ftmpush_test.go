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

func TestAccFortiOSSystemFtmPush_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemFtmPush_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemFtmPushConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemFtmPushExists("fortios_system_ftmpush.trname"),
					resource.TestCheckResourceAttr("fortios_system_ftmpush.trname", "server_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_system_ftmpush.trname", "server_port", "4433"),
					resource.TestCheckResourceAttr("fortios_system_ftmpush.trname", "status", "disable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemFtmPushExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemFtmPush: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemFtmPush is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemFtmPush(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemFtmPush: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemFtmPush: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemFtmPushDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_ftmpush" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemFtmPush(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemFtmPush %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemFtmPushConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_ftmpush" "trname" {
  server_ip   = "0.0.0.0"
  server_port = 4433
  status      = "disable"
}
`)
}
