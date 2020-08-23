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

func TestAccFortiOSLogCustomField_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSLogCustomField_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSLogCustomFieldConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSLogCustomFieldExists("fortios_log_customfield.trname"),
					resource.TestCheckResourceAttr("fortios_log_customfield.trname", "fosid", "1"),
					resource.TestCheckResourceAttr("fortios_log_customfield.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_log_customfield.trname", "value", "logteststr"),
				),
			},
		},
	})
}

func testAccCheckFortiOSLogCustomFieldExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found LogCustomField: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No LogCustomField is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadLogCustomField(i)

		if err != nil {
			return fmt.Errorf("Error reading LogCustomField: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating LogCustomField: %s", n)
		}

		return nil
	}
}

func testAccCheckLogCustomFieldDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_log_customfield" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadLogCustomField(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error LogCustomField %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSLogCustomFieldConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_log_customfield" "trname" {
  fosid = "1"
  name  = "%[1]s"
  value = "logteststr"
}

`, name)
}
