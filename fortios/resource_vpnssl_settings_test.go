
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

func TestAccFortiOSVpnSslSettings_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSVpnSslSettings_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSVpnSslSettingsConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSVpnSslSettingsExists("fortios_vpnssl_settings.trname"),
                    resource.TestCheckResourceAttr("fortios_vpnssl_settings.trname", "login_attempt_limit", "2"),
                    resource.TestCheckResourceAttr("fortios_vpnssl_settings.trname", "login_block_time", "60"),
                    resource.TestCheckResourceAttr("fortios_vpnssl_settings.trname", "login_timeout", "30"),
                    resource.TestCheckResourceAttr("fortios_vpnssl_settings.trname", "port", "443"),
                    resource.TestCheckResourceAttr("fortios_vpnssl_settings.trname", "servercert", "self-sign"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSVpnSslSettingsExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found VpnSslSettings: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No VpnSslSettings is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadVpnSslSettings(i)

		if err != nil {
			return fmt.Errorf("Error reading VpnSslSettings: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating VpnSslSettings: %s", n)
		}

		return nil
	}
}

func testAccCheckVpnSslSettingsDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_vpnssl_settings" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadVpnSslSettings(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error VpnSslSettings %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSVpnSslSettingsConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_vpnssl_settings" "trname" {
  login_attempt_limit = 2
  login_block_time    = 60
  login_timeout       = 30
  port                = 443
  servercert          = "self-sign"
}
`)
}
