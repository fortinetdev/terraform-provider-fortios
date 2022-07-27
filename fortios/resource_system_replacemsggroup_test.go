// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

package fortios

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"log"
	"testing"
)

func TestAccFortiOSSystemReplacemsgGroup_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemReplacemsgGroup_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemReplacemsgGroupConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemReplacemsgGroupExists("fortios_system_replacemsggroup.trname"),
					resource.TestCheckResourceAttr("fortios_system_replacemsggroup.trname", "comment", "sgh"),
					resource.TestCheckResourceAttr("fortios_system_replacemsggroup.trname", "group_type", "utm"),
					resource.TestCheckResourceAttr("fortios_system_replacemsggroup.trname", "name", rname),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemReplacemsgGroupExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemReplacemsgGroup: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemReplacemsgGroup is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemReplacemsgGroup(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemReplacemsgGroup: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemReplacemsgGroup: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemReplacemsgGroupDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_replacemsggroup" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemReplacemsgGroup(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemReplacemsgGroup %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemReplacemsgGroupConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_replacemsggroup" "trname" {
  comment    = "sgh"
  group_type = "utm"
  name       = "%[1]s"
}
`, name)
}
