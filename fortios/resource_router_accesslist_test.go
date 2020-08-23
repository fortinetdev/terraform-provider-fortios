
// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

package fortios
import (
    "fmt"
	"log"
    "testing"
    "github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
    "github.com/hashicorp/terraform-plugin-sdk/helper/resource"
    "github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccFortiOSRouterAccessList_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSRouterAccessList_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSRouterAccessListConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSRouterAccessListExists("fortios_router_accesslist.trname"),
                    resource.TestCheckResourceAttr("fortios_router_accesslist.trname", "comments", "test accesslist"),
                    resource.TestCheckResourceAttr("fortios_router_accesslist.trname", "name", rname),
                ),
            },
        },
    })
}

func testAccCheckFortiOSRouterAccessListExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found RouterAccessList: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No RouterAccessList is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadRouterAccessList(i)

		if err != nil {
			return fmt.Errorf("Error reading RouterAccessList: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating RouterAccessList: %s", n)
		}

		return nil
	}
}

func testAccCheckRouterAccessListDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_router_accesslist" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadRouterAccessList(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error RouterAccessList %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSRouterAccessListConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_router_accesslist" "trname" {
  comments = "test accesslist"
  name     = "%[1]s"
}
`, name)
}
