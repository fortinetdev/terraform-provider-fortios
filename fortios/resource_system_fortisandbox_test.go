
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

func TestAccFortiOSSystemFortisandbox_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSSystemFortisandbox_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSSystemFortisandboxConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSSystemFortisandboxExists("fortios_system_fortisandbox.trname"),
                    resource.TestCheckResourceAttr("fortios_system_fortisandbox.trname", "enc_algorithm", "default"),
                    resource.TestCheckResourceAttr("fortios_system_fortisandbox.trname", "ssl_min_proto_version", "default"),
                    resource.TestCheckResourceAttr("fortios_system_fortisandbox.trname", "status", "disable"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSSystemFortisandboxExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemFortisandbox: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemFortisandbox is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemFortisandbox(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemFortisandbox: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemFortisandbox: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemFortisandboxDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_fortisandbox" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemFortisandbox(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemFortisandbox %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemFortisandboxConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_fortisandbox" "trname" {
  enc_algorithm         = "default"
  ssl_min_proto_version = "default"
  status                = "disable"
}
`)
}
