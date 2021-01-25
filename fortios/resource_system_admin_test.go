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

func TestAccFortiOSSystemAdmin_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemAdmin_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemAdminConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemAdminExists("fortios_system_admin.trname"),
					resource.TestCheckResourceAttr("fortios_system_admin.trname", "accprofile", "super_admin"),
					resource.TestCheckResourceAttr("fortios_system_admin.trname", "accprofile_override", "disable"),
					resource.TestCheckResourceAttr("fortios_system_admin.trname", "allow_remove_admin_session", "enable"),
					resource.TestCheckResourceAttr("fortios_system_admin.trname", "force_password_change", "disable"),
					resource.TestCheckResourceAttr("fortios_system_admin.trname", "guest_auth", "disable"),
					resource.TestCheckResourceAttr("fortios_system_admin.trname", "hidden", "0"),
					resource.TestCheckResourceAttr("fortios_system_admin.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_system_admin.trname", "password", "fdafdace"),
					resource.TestCheckResourceAttr("fortios_system_admin.trname", "password_expire", "0000-00-00 00:00:00"),
					resource.TestCheckResourceAttr("fortios_system_admin.trname", "peer_auth", "disable"),
					resource.TestCheckResourceAttr("fortios_system_admin.trname", "radius_vdom_override", "disable"),
					resource.TestCheckResourceAttr("fortios_system_admin.trname", "remote_auth", "disable"),
					resource.TestCheckResourceAttr("fortios_system_admin.trname", "two_factor", "disable"),
					resource.TestCheckResourceAttr("fortios_system_admin.trname", "wildcard", "disable"),
					resource.TestCheckResourceAttr("fortios_system_admin.trname", "vdom.0.name", "root"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemAdminExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemAdmin: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemAdmin is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemAdmin(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemAdmin: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemAdmin: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemAdminDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_admin" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemAdmin(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemAdmin %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemAdminConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_admin" "trname" {
  accprofile                 = "super_admin"
  accprofile_override        = "disable"
  allow_remove_admin_session = "enable"
  force_password_change      = "disable"
  guest_auth                 = "disable"
  hidden                     = 0
  name                       = "%[1]s"
  password                   = "fdafdace"
  password_expire            = "0000-00-00 00:00:00"
  peer_auth                  = "disable"
  radius_vdom_override       = "disable"
  remote_auth                = "disable"
  two_factor                 = "disable"
  wildcard                   = "disable"
  vdom {
    name = "root"
  }
}
`, name)
}
