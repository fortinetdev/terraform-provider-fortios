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

func TestAccFortiOSSystemVxlan_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemVxlan_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemVxlanConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemVxlanExists("fortios_system_vxlan.trname"),
					resource.TestCheckResourceAttr("fortios_system_vxlan.trname", "dstport", "4789"),
					resource.TestCheckResourceAttr("fortios_system_vxlan.trname", "interface", "port3"),
					resource.TestCheckResourceAttr("fortios_system_vxlan.trname", "ip_version", "ipv4-unicast"),
					resource.TestCheckResourceAttr("fortios_system_vxlan.trname", "remote_ip.0.ip", "1.1.1.1"),
					resource.TestCheckResourceAttr("fortios_system_vxlan.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_system_vxlan.trname", "vni", "3"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemVxlanExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemVxlan: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemVxlan is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemVxlan(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemVxlan: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemVxlan: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemVxlanDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_vxlan" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemVxlan(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemVxlan %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemVxlanConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_vxlan" "trname" {
  dstport    = 4789
  interface  = "port3"
  ip_version = "ipv4-unicast"
  remote_ip {
    ip = "1.1.1.1"
  }
  name = "%[1]s"
  vni  = 3
}
`, name)
}
