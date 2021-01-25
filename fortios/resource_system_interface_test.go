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

func TestAccFortiOSSystemInterface_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemInterface_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemInterfaceConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemInterfaceExists("fortios_system_interface.trname"),
					resource.TestCheckResourceAttr("fortios_system_interface.trname", "algorithm", "L4"),
					resource.TestCheckResourceAttr("fortios_system_interface.trname", "defaultgw", "enable"),
					resource.TestCheckResourceAttr("fortios_system_interface.trname", "distance", "5"),
					resource.TestCheckResourceAttr("fortios_system_interface.trname", "ip", "0.0.0.0 0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_system_interface.trname", "mtu", "1500"),
					resource.TestCheckResourceAttr("fortios_system_interface.trname", "mtu_override", "disable"),
					resource.TestCheckResourceAttr("fortios_system_interface.trname", "name", "port3"),
					resource.TestCheckResourceAttr("fortios_system_interface.trname", "type", "physical"),
					resource.TestCheckResourceAttr("fortios_system_interface.trname", "vdom", "root"),
					resource.TestCheckResourceAttr("fortios_system_interface.trname", "mode", "dhcp"),
					resource.TestCheckResourceAttr("fortios_system_interface.trname", "snmp_index", "3"),
					resource.TestCheckResourceAttr("fortios_system_interface.trname", "description", "Created by Terraform Provider for FortiOS"),
					resource.TestCheckResourceAttr("fortios_system_interface.trname", "ipv6.0.nd_mode", "basic"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemInterfaceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemInterface: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemInterface is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemInterface(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemInterface: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemInterface: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemInterfaceDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_interface" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemInterface(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemInterface %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemInterfaceConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_interface" "trname" {
  algorithm    = "L4"
  defaultgw    = "enable"
  distance     = 5
  ip           = "0.0.0.0 0.0.0.0"
  mtu          = 1500
  mtu_override = "disable"
  name         = "port3"
  type         = "physical"
  vdom         = "root"
  mode         = "dhcp"
  snmp_index   = 3
  description  = "Created by Terraform Provider for FortiOS"
  ipv6 {
    nd_mode = "basic"
  }
}
`)
}
