
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

func TestAccFortiOSSystemEmailServer_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSSystemEmailServer_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSSystemEmailServerConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSSystemEmailServerExists("fortios_system_emailserver.trname"),
                    resource.TestCheckResourceAttr("fortios_system_emailserver.trname", "authenticate", "disable"),
                    resource.TestCheckResourceAttr("fortios_system_emailserver.trname", "port", "465"),
                    resource.TestCheckResourceAttr("fortios_system_emailserver.trname", "security", "smtps"),
                    resource.TestCheckResourceAttr("fortios_system_emailserver.trname", "server", "notification.fortinet.net"),
                    resource.TestCheckResourceAttr("fortios_system_emailserver.trname", "source_ip", "0.0.0.0"),
                    resource.TestCheckResourceAttr("fortios_system_emailserver.trname", "source_ip6", "::"),
                    resource.TestCheckResourceAttr("fortios_system_emailserver.trname", "ssl_min_proto_version", "default"),
                    resource.TestCheckResourceAttr("fortios_system_emailserver.trname", "type", "custom"),
                    resource.TestCheckResourceAttr("fortios_system_emailserver.trname", "validate_server", "disable"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSSystemEmailServerExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemEmailServer: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemEmailServer is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemEmailServer(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemEmailServer: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemEmailServer: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemEmailServerDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_emailserver" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemEmailServer(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemEmailServer %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemEmailServerConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_emailserver" "trname" {
  authenticate          = "disable"
  port                  = 465
  security              = "smtps"
  server                = "notification.fortinet.net"
  source_ip             = "0.0.0.0"
  source_ip6            = "::"
  ssl_min_proto_version = "default"
  type                  = "custom"
  validate_server       = "disable"
}
`)
}
