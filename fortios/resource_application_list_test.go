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

func TestAccFortiOSApplicationList_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSApplicationList_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSApplicationListConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSApplicationListExists("fortios_application_list.trname"),
					resource.TestCheckResourceAttr("fortios_application_list.trname", "app_replacemsg", "enable"),
					resource.TestCheckResourceAttr("fortios_application_list.trname", "deep_app_inspection", "enable"),
					resource.TestCheckResourceAttr("fortios_application_list.trname", "enforce_default_app_port", "disable"),
					resource.TestCheckResourceAttr("fortios_application_list.trname", "extended_log", "disable"),
					resource.TestCheckResourceAttr("fortios_application_list.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_application_list.trname", "options", "allow-dns"),
					resource.TestCheckResourceAttr("fortios_application_list.trname", "other_application_action", "pass"),
					resource.TestCheckResourceAttr("fortios_application_list.trname", "other_application_log", "disable"),
					resource.TestCheckResourceAttr("fortios_application_list.trname", "unknown_application_action", "pass"),
					resource.TestCheckResourceAttr("fortios_application_list.trname", "unknown_application_log", "disable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSApplicationListExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found ApplicationList: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ApplicationList is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadApplicationList(i, "root")

		if err != nil {
			return fmt.Errorf("Error reading ApplicationList: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating ApplicationList: %s", n)
		}

		return nil
	}
}

func testAccCheckApplicationListDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_application_list" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadApplicationList(i, "root")

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error ApplicationList %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSApplicationListConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_application_list" "trname" {
  app_replacemsg             = "enable"
  deep_app_inspection        = "enable"
  enforce_default_app_port   = "disable"
  extended_log               = "disable"
  name                       = "%[1]s"
  options                    = "allow-dns"
  other_application_action   = "pass"
  other_application_log      = "disable"
  unknown_application_action = "pass"
  unknown_application_log    = "disable"
}
`, name)
}
