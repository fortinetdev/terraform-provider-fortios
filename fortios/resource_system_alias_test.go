
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

func TestAccFortiOSSystemAlias_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSSystemAlias_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSSystemAliasConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSSystemAliasExists("fortios_system_alias.trname"),
                    resource.TestCheckResourceAttr("fortios_system_alias.trname", "name", rname),
                ),
            },
        },
    })
}

func testAccCheckFortiOSSystemAliasExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemAlias: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemAlias is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemAlias(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemAlias: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemAlias: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemAliasDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_alias" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemAlias(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemAlias %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemAliasConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_alias" "trname" {
  name = "%[1]s"
}
`, name)
}
