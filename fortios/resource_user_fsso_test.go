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

func TestAccFortiOSUserFsso_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSUserFsso_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSUserFssoConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSUserFssoExists("fortios_user_fsso.trname"),
					resource.TestCheckResourceAttr("fortios_user_fsso.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_user_fsso.trname", "port", "32381"),
					resource.TestCheckResourceAttr("fortios_user_fsso.trname", "port2", "8000"),
					resource.TestCheckResourceAttr("fortios_user_fsso.trname", "port3", "8000"),
					resource.TestCheckResourceAttr("fortios_user_fsso.trname", "port4", "8000"),
					resource.TestCheckResourceAttr("fortios_user_fsso.trname", "port5", "8000"),
					resource.TestCheckResourceAttr("fortios_user_fsso.trname", "server", "1.1.1.1"),
					resource.TestCheckResourceAttr("fortios_user_fsso.trname", "source_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_user_fsso.trname", "source_ip6", "::"),
				),
			},
		},
	})
}

func testAccCheckFortiOSUserFssoExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found UserFsso: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No UserFsso is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadUserFsso(i)

		if err != nil {
			return fmt.Errorf("Error reading UserFsso: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating UserFsso: %s", n)
		}

		return nil
	}
}

func testAccCheckUserFssoDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_user_fsso" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadUserFsso(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error UserFsso %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSUserFssoConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_user_fsso" "trname" {
  name       = "%[1]s"
  port       = 32381
  port2      = 8000
  port3      = 8000
  port4      = 8000
  port5      = 8000
  server     = "1.1.1.1"
  source_ip  = "0.0.0.0"
  source_ip6 = "::"
}
`, name)
}
