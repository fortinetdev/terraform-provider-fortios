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

func TestAccFortiOSSwitchControllerQosIpDscpMap_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSwitchControllerQosIpDscpMap_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSwitchControllerQosIpDscpMapConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSwitchControllerQosIpDscpMapExists("fortios_switchcontrollerqos_ipdscpmap.trname"),
					resource.TestCheckResourceAttr("fortios_switchcontrollerqos_ipdscpmap.trname", "description", "DEIW"),
					resource.TestCheckResourceAttr("fortios_switchcontrollerqos_ipdscpmap.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_switchcontrollerqos_ipdscpmap.trname", "map.0.cos_queue", "3"),
					resource.TestCheckResourceAttr("fortios_switchcontrollerqos_ipdscpmap.trname", "map.0.diffserv", "CS0 CS1 AF11"),
					resource.TestCheckResourceAttr("fortios_switchcontrollerqos_ipdscpmap.trname", "map.0.name", "1"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSwitchControllerQosIpDscpMapExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SwitchControllerQosIpDscpMap: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SwitchControllerQosIpDscpMap is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSwitchControllerQosIpDscpMap(i)

		if err != nil {
			return fmt.Errorf("Error reading SwitchControllerQosIpDscpMap: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SwitchControllerQosIpDscpMap: %s", n)
		}

		return nil
	}
}

func testAccCheckSwitchControllerQosIpDscpMapDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_switchcontrollerqos_ipdscpmap" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSwitchControllerQosIpDscpMap(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SwitchControllerQosIpDscpMap %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSwitchControllerQosIpDscpMapConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_switchcontrollerqos_ipdscpmap" "trname" {
  description = "DEIW"
  name        = "%[1]s"

  map {
    cos_queue = 3
    diffserv  = "CS0 CS1 AF11"
    name      = "1"
  }
}

`, name)
}
