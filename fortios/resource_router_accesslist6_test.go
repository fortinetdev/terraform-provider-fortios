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

func TestAccFortiOSRouterAccessList6_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSRouterAccessList6_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSRouterAccessList6Config(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSRouterAccessList6Exists("fortios_router_accesslist6.trname"),
					resource.TestCheckResourceAttr("fortios_router_accesslist6.trname", "comments", "access-list6 test"),
					resource.TestCheckResourceAttr("fortios_router_accesslist6.trname", "name", rname),
				),
			},
		},
	})
}

func testAccCheckFortiOSRouterAccessList6Exists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found RouterAccessList6: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No RouterAccessList6 is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadRouterAccessList6(i)

		if err != nil {
			return fmt.Errorf("Error reading RouterAccessList6: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating RouterAccessList6: %s", n)
		}

		return nil
	}
}

func testAccCheckRouterAccessList6Destroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_router_accesslist6" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadRouterAccessList6(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error RouterAccessList6 %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSRouterAccessList6Config(name string) string {
	return fmt.Sprintf(`
resource "fortios_router_accesslist6" "trname" {
  comments = "access-list6 test"
  name     = "%[1]s"
}
`, name)
}
