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

func TestAccFortiOSUserPop3_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSUserPop3_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSUserPop3Config(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSUserPop3Exists("fortios_user_pop3.trname"),
					resource.TestCheckResourceAttr("fortios_user_pop3.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_user_pop3.trname", "port", "0"),
					resource.TestCheckResourceAttr("fortios_user_pop3.trname", "secure", "pop3s"),
					resource.TestCheckResourceAttr("fortios_user_pop3.trname", "server", "1.1.1.1"),
					resource.TestCheckResourceAttr("fortios_user_pop3.trname", "ssl_min_proto_version", "default"),
				),
			},
		},
	})
}

func testAccCheckFortiOSUserPop3Exists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found UserPop3: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No UserPop3 is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadUserPop3(i)

		if err != nil {
			return fmt.Errorf("Error reading UserPop3: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating UserPop3: %s", n)
		}

		return nil
	}
}

func testAccCheckUserPop3Destroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_user_pop3" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadUserPop3(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error UserPop3 %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSUserPop3Config(name string) string {
	return fmt.Sprintf(`
resource "fortios_user_pop3" "trname" {
  name                  = "%[1]s"
  port                  = 0
  secure                = "pop3s"
  server                = "1.1.1.1"
  ssl_min_proto_version = "default"
}
`, name)
}
