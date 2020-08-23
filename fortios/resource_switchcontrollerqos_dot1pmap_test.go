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

func TestAccFortiOSSwitchControllerQosDot1PMap_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSwitchControllerQosDot1PMap_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSwitchControllerQosDot1PMapConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSwitchControllerQosDot1PMapExists("fortios_switchcontrollerqos_dot1pmap.trname"),
					resource.TestCheckResourceAttr("fortios_switchcontrollerqos_dot1pmap.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_switchcontrollerqos_dot1pmap.trname", "priority_0", "queue-0"),
					resource.TestCheckResourceAttr("fortios_switchcontrollerqos_dot1pmap.trname", "priority_1", "queue-0"),
					resource.TestCheckResourceAttr("fortios_switchcontrollerqos_dot1pmap.trname", "priority_2", "queue-0"),
					resource.TestCheckResourceAttr("fortios_switchcontrollerqos_dot1pmap.trname", "priority_3", "queue-0"),
					resource.TestCheckResourceAttr("fortios_switchcontrollerqos_dot1pmap.trname", "priority_4", "queue-0"),
					resource.TestCheckResourceAttr("fortios_switchcontrollerqos_dot1pmap.trname", "priority_5", "queue-0"),
					resource.TestCheckResourceAttr("fortios_switchcontrollerqos_dot1pmap.trname", "priority_6", "queue-0"),
					resource.TestCheckResourceAttr("fortios_switchcontrollerqos_dot1pmap.trname", "priority_7", "queue-0"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSwitchControllerQosDot1PMapExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SwitchControllerQosDot1PMap: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SwitchControllerQosDot1PMap is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSwitchControllerQosDot1PMap(i)

		if err != nil {
			return fmt.Errorf("Error reading SwitchControllerQosDot1PMap: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SwitchControllerQosDot1PMap: %s", n)
		}

		return nil
	}
}

func testAccCheckSwitchControllerQosDot1PMapDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_switchcontrollerqos_dot1pmap" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSwitchControllerQosDot1PMap(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SwitchControllerQosDot1PMap %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSwitchControllerQosDot1PMapConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_switchcontrollerqos_dot1pmap" "trname" {
  name       = "%[1]s"
  priority_0 = "queue-0"
  priority_1 = "queue-0"
  priority_2 = "queue-0"
  priority_3 = "queue-0"
  priority_4 = "queue-0"
  priority_5 = "queue-0"
  priority_6 = "queue-0"
  priority_7 = "queue-0"
}

`, name)
}
