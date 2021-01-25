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

func TestAccFortiOSSystemDhcp6Server_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemDhcp6Server_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemDhcp6ServerConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemDhcp6ServerExists("fortios_systemdhcp6_server.trname"),
					resource.TestCheckResourceAttr("fortios_systemdhcp6_server.trname", "fosid", "1"),
					resource.TestCheckResourceAttr("fortios_systemdhcp6_server.trname", "interface", "port3"),
					resource.TestCheckResourceAttr("fortios_systemdhcp6_server.trname", "lease_time", "604800"),
					resource.TestCheckResourceAttr("fortios_systemdhcp6_server.trname", "rapid_commit", "disable"),
					resource.TestCheckResourceAttr("fortios_systemdhcp6_server.trname", "status", "enable"),
					resource.TestCheckResourceAttr("fortios_systemdhcp6_server.trname", "subnet", "2001:db8:1234:113::/64"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemDhcp6ServerExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemDhcp6Server: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemDhcp6Server is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemDhcp6Server(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemDhcp6Server: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemDhcp6Server: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemDhcp6ServerDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_systemdhcp6_server" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemDhcp6Server(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemDhcp6Server %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemDhcp6ServerConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_systemdhcp6_server" "trname" {
  fosid        = 1
  interface    = "port3"
  lease_time   = 604800
  rapid_commit = "disable"
  status       = "enable"
  subnet       = "2001:db8:1234:113::/64"
}
`)
}
