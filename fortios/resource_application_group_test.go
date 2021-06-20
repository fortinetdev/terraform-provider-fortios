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

func TestAccFortiOSApplicationGroup_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSApplicationGroup_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSApplicationGroupConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSApplicationGroupExists("fortios_application_group.trname"),
					resource.TestCheckResourceAttr("fortios_application_group.trname", "comment", "group1 test"),
					resource.TestCheckResourceAttr("fortios_application_group.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_application_group.trname", "type", "category"),
					resource.TestCheckResourceAttr("fortios_application_group.trname", "category.0.id", "2"),
				),
			},
		},
	})
}

func testAccCheckFortiOSApplicationGroupExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found ApplicationGroup: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ApplicationGroup is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadApplicationGroup(i, "root")

		if err != nil {
			return fmt.Errorf("Error reading ApplicationGroup: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating ApplicationGroup: %s", n)
		}

		return nil
	}
}

func testAccCheckApplicationGroupDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_application_group" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadApplicationGroup(i, "root")

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error ApplicationGroup %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSApplicationGroupConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_application_group" "trname" {
  comment = "group1 test"
  name    = "%[1]s"
  type    = "category"
  category {
    id = 2
  }
}
`, name)
}
