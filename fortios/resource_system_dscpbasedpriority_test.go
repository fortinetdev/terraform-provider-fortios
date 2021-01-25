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

func TestAccFortiOSSystemDscpBasedPriority_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemDscpBasedPriority_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemDscpBasedPriorityConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemDscpBasedPriorityExists("fortios_system_dscpbasedpriority.trname"),
					resource.TestCheckResourceAttr("fortios_system_dscpbasedpriority.trname", "ds", "1"),
					resource.TestCheckResourceAttr("fortios_system_dscpbasedpriority.trname", "fosid", "1"),
					resource.TestCheckResourceAttr("fortios_system_dscpbasedpriority.trname", "priority", "low"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemDscpBasedPriorityExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemDscpBasedPriority: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemDscpBasedPriority is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemDscpBasedPriority(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemDscpBasedPriority: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemDscpBasedPriority: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemDscpBasedPriorityDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_dscpbasedpriority" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemDscpBasedPriority(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemDscpBasedPriority %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemDscpBasedPriorityConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_dscpbasedpriority" "trname" {
  ds       = 1
  fosid    = 1
  priority = "low"
}
`)
}
