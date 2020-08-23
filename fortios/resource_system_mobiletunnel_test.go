
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

func TestAccFortiOSSystemMobileTunnel_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSSystemMobileTunnel_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSSystemMobileTunnelConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSSystemMobileTunnelExists("fortios_system_mobiletunnel.trname"),
                    resource.TestCheckResourceAttr("fortios_system_mobiletunnel.trname", "hash_algorithm", "hmac-md5"),
                    resource.TestCheckResourceAttr("fortios_system_mobiletunnel.trname", "home_address", "0.0.0.0"),
                    resource.TestCheckResourceAttr("fortios_system_mobiletunnel.trname", "home_agent", "1.1.1.1"),
                    resource.TestCheckResourceAttr("fortios_system_mobiletunnel.trname", "lifetime", "65535"),
                    resource.TestCheckResourceAttr("fortios_system_mobiletunnel.trname", "n_mhae_key", "'ENC M2wyM3DcnUhqgich7vsLk5oVuPAI9LTkcFNt0c3jI1ujC6w1XBot7gsRAf2S8X5dagfUnJGhZ5LrQxw21e4y8oXuCOLp8MmaRZbCkxYCAl1wm/wVY3aNzVk2+jE='"),
                    resource.TestCheckResourceAttr("fortios_system_mobiletunnel.trname", "n_mhae_key_type", "ascii"),
                    resource.TestCheckResourceAttr("fortios_system_mobiletunnel.trname", "n_mhae_spi", "256"),
                    resource.TestCheckResourceAttr("fortios_system_mobiletunnel.trname", "name", rname),
                    resource.TestCheckResourceAttr("fortios_system_mobiletunnel.trname", "reg_interval", "5"),
                    resource.TestCheckResourceAttr("fortios_system_mobiletunnel.trname", "reg_retry", "3"),
                    resource.TestCheckResourceAttr("fortios_system_mobiletunnel.trname", "renew_interval", "60"),
                    resource.TestCheckResourceAttr("fortios_system_mobiletunnel.trname", "roaming_interface", "port3"),
                    resource.TestCheckResourceAttr("fortios_system_mobiletunnel.trname", "status", "disable"),
                    resource.TestCheckResourceAttr("fortios_system_mobiletunnel.trname", "tunnel_mode", "gre"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSSystemMobileTunnelExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemMobileTunnel: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemMobileTunnel is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemMobileTunnel(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemMobileTunnel: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemMobileTunnel: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemMobileTunnelDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_mobiletunnel" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemMobileTunnel(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemMobileTunnel %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemMobileTunnelConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_mobiletunnel" "trname" {
  hash_algorithm    = "hmac-md5"
  home_address      = "0.0.0.0"
  home_agent        = "1.1.1.1"
  lifetime          = 65535
  n_mhae_key        = "'ENC M2wyM3DcnUhqgich7vsLk5oVuPAI9LTkcFNt0c3jI1ujC6w1XBot7gsRAf2S8X5dagfUnJGhZ5LrQxw21e4y8oXuCOLp8MmaRZbCkxYCAl1wm/wVY3aNzVk2+jE='"
  n_mhae_key_type   = "ascii"
  n_mhae_spi        = 256
  name              = "%[1]s"
  reg_interval      = 5
  reg_retry         = 3
  renew_interval    = 60
  roaming_interface = "port3"
  status            = "disable"
  tunnel_mode       = "gre"
}


`, name)
}
