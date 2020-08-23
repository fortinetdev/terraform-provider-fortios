
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

func TestAccFortiOSSystemSessionHelper_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSSystemSessionHelper_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSSystemSessionHelperConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSSystemSessionHelperExists("fortios_system_sessionhelper.trname"),
                    resource.TestCheckResourceAttr("fortios_system_sessionhelper.trname", "fosid", "33"),
                    resource.TestCheckResourceAttr("fortios_system_sessionhelper.trname", "name", "ftp"),
                    resource.TestCheckResourceAttr("fortios_system_sessionhelper.trname", "port", "3234"),
                    resource.TestCheckResourceAttr("fortios_system_sessionhelper.trname", "protocol", "17"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSSystemSessionHelperExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemSessionHelper: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemSessionHelper is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemSessionHelper(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemSessionHelper: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemSessionHelper: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemSessionHelperDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_sessionhelper" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemSessionHelper(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemSessionHelper %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemSessionHelperConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_sessionhelper" "trname" {
  fosid    = 33
  name     = "ftp"
  port     = 3234
  protocol = 17
}
`)
}
