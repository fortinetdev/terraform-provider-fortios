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

func TestAccFortiOSRouterCommunityList_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSRouterCommunityList_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSRouterCommunityListConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSRouterCommunityListExists("fortios_router_communitylist.trname"),
					resource.TestCheckResourceAttr("fortios_router_communitylist.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_router_communitylist.trname", "type", "standard"),
					resource.TestCheckResourceAttr("fortios_router_communitylist.trname", "rule.0.action", "deny"),
					resource.TestCheckResourceAttr("fortios_router_communitylist.trname", "rule.0.match", "123:234 345:456"),
					resource.TestCheckResourceAttr("fortios_router_communitylist.trname", "rule.0.regexp", "123:234 345:456"),
				),
			},
		},
	})
}

func testAccCheckFortiOSRouterCommunityListExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found RouterCommunityList: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No RouterCommunityList is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadRouterCommunityList(i)

		if err != nil {
			return fmt.Errorf("Error reading RouterCommunityList: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating RouterCommunityList: %s", n)
		}

		return nil
	}
}

func testAccCheckRouterCommunityListDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_router_communitylist" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadRouterCommunityList(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error RouterCommunityList %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSRouterCommunityListConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_router_communitylist" "trname" {
  name = "%[1]s"
  type = "standard"

  rule {
    action = "deny"
    match  = "123:234 345:456"
    regexp = "123:234 345:456"
  }
}
`, name)
}
