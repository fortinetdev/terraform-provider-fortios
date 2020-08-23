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

func TestAccFortiOSUserGroup_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSUserGroup_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSUserGroupConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSUserGroupExists("fortios_user_group.trname"),
					resource.TestCheckResourceAttr("fortios_user_group.trname", "company", "optional"),
					resource.TestCheckResourceAttr("fortios_user_group.trname", "email", "enable"),
					resource.TestCheckResourceAttr("fortios_user_group.trname", "expire", "14400"),
					resource.TestCheckResourceAttr("fortios_user_group.trname", "expire_type", "immediately"),
					resource.TestCheckResourceAttr("fortios_user_group.trname", "group_type", "firewall"),
					resource.TestCheckResourceAttr("fortios_user_group.trname", "max_accounts", "0"),
					resource.TestCheckResourceAttr("fortios_user_group.trname", "mobile_phone", "disable"),
					resource.TestCheckResourceAttr("fortios_user_group.trname", "multiple_guest_add", "disable"),
					resource.TestCheckResourceAttr("fortios_user_group.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_user_group.trname", "member.0.name", "guest"),
				),
			},
		},
	})
}

func testAccCheckFortiOSUserGroupExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found UserGroup: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No UserGroup is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadUserGroup(i)

		if err != nil {
			return fmt.Errorf("Error reading UserGroup: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating UserGroup: %s", n)
		}

		return nil
	}
}

func testAccCheckUserGroupDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_user_group" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadUserGroup(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error UserGroup %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSUserGroupConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_user_group" "trname" {
  company            = "optional"
  email              = "enable"
  expire             = 14400
  expire_type        = "immediately"
  group_type         = "firewall"
  max_accounts       = 0
  mobile_phone       = "disable"
  multiple_guest_add = "disable"
  name               = "%[1]s"

  member {
    name = "guest"
  }
}
`, name)
}
