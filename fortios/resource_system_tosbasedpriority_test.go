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

func TestAccFortiOSSystemTosBasedPriority_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemTosBasedPriority_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemTosBasedPriorityConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemTosBasedPriorityExists("fortios_system_tosbasedpriority.trname"),
					resource.TestCheckResourceAttr("fortios_system_tosbasedpriority.trname", "fosid", "1"),
					resource.TestCheckResourceAttr("fortios_system_tosbasedpriority.trname", "priority", "low"),
					resource.TestCheckResourceAttr("fortios_system_tosbasedpriority.trname", "tos", "11"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemTosBasedPriorityExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemTosBasedPriority: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemTosBasedPriority is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemTosBasedPriority(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemTosBasedPriority: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemTosBasedPriority: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemTosBasedPriorityDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_tosbasedpriority" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemTosBasedPriority(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemTosBasedPriority %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemTosBasedPriorityConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_tosbasedpriority" "trname" {
  fosid    = 1
  priority = "low"
  tos      = 11
}
`)
}
