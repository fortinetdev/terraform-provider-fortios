
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

func TestAccFortiOSSystemNtp_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSSystemNtp_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSSystemNtpConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSSystemNtpExists("fortios_system_ntp.trname"),
                    resource.TestCheckResourceAttr("fortios_system_ntp.trname", "ntpsync", "enable"),
                    resource.TestCheckResourceAttr("fortios_system_ntp.trname", "server_mode", "disable"),
                    resource.TestCheckResourceAttr("fortios_system_ntp.trname", "source_ip", "0.0.0.0"),
                    resource.TestCheckResourceAttr("fortios_system_ntp.trname", "source_ip6", "::"),
                    resource.TestCheckResourceAttr("fortios_system_ntp.trname", "syncinterval", "1"),
                    resource.TestCheckResourceAttr("fortios_system_ntp.trname", "type", "fortiguard"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSSystemNtpExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemNtp: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemNtp is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemNtp(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemNtp: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemNtp: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemNtpDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_ntp" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemNtp(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemNtp %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemNtpConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_ntp" "trname" {
  ntpsync      = "enable"
  server_mode  = "disable"
  source_ip    = "0.0.0.0"
  source_ip6   = "::"
  syncinterval = 1
  type         = "fortiguard"
}
`)
}
