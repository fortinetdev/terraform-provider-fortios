
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

func TestAccFortiOSVpnIpsecManualkeyInterface_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSVpnIpsecManualkeyInterface_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSVpnIpsecManualkeyInterfaceConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSVpnIpsecManualkeyInterfaceExists("fortios_vpnipsec_manualkeyinterface.trname"),
                    resource.TestCheckResourceAttr("fortios_vpnipsec_manualkeyinterface.trname", "addr_type", "4"),
                    resource.TestCheckResourceAttr("fortios_vpnipsec_manualkeyinterface.trname", "auth_alg", "null"),
                    resource.TestCheckResourceAttr("fortios_vpnipsec_manualkeyinterface.trname", "auth_key", "-"),
                    resource.TestCheckResourceAttr("fortios_vpnipsec_manualkeyinterface.trname", "enc_alg", "des"),
                    resource.TestCheckResourceAttr("fortios_vpnipsec_manualkeyinterface.trname", "enc_key", "CECA2184ACADAEEF"),
                    resource.TestCheckResourceAttr("fortios_vpnipsec_manualkeyinterface.trname", "interface", "port3"),
                    resource.TestCheckResourceAttr("fortios_vpnipsec_manualkeyinterface.trname", "ip_version", "4"),
                    resource.TestCheckResourceAttr("fortios_vpnipsec_manualkeyinterface.trname", "local_gw", "0.0.0.0"),
                    resource.TestCheckResourceAttr("fortios_vpnipsec_manualkeyinterface.trname", "local_gw6", "::"),
                    resource.TestCheckResourceAttr("fortios_vpnipsec_manualkeyinterface.trname", "local_spi", "0x100"),
                    resource.TestCheckResourceAttr("fortios_vpnipsec_manualkeyinterface.trname", "name", rname),
                    resource.TestCheckResourceAttr("fortios_vpnipsec_manualkeyinterface.trname", "remote_gw", "2.2.2.2"),
                    resource.TestCheckResourceAttr("fortios_vpnipsec_manualkeyinterface.trname", "remote_gw6", "::"),
                    resource.TestCheckResourceAttr("fortios_vpnipsec_manualkeyinterface.trname", "remote_spi", "0x100"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSVpnIpsecManualkeyInterfaceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found VpnIpsecManualkeyInterface: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No VpnIpsecManualkeyInterface is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadVpnIpsecManualkeyInterface(i)

		if err != nil {
			return fmt.Errorf("Error reading VpnIpsecManualkeyInterface: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating VpnIpsecManualkeyInterface: %s", n)
		}

		return nil
	}
}

func testAccCheckVpnIpsecManualkeyInterfaceDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_vpnipsec_manualkeyinterface" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadVpnIpsecManualkeyInterface(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error VpnIpsecManualkeyInterface %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSVpnIpsecManualkeyInterfaceConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_vpnipsec_manualkeyinterface" "trname" {
  addr_type  = "4"
  auth_alg   = "null"
  auth_key   = "-"
  enc_alg    = "des"
  enc_key    = "CECA2184ACADAEEF"
  interface  = "port3"
  ip_version = "4"
  local_gw   = "0.0.0.0"
  local_gw6  = "::"
  local_spi  = "0x100"
  name       = "%[1]s"
  remote_gw  = "2.2.2.2"
  remote_gw6 = "::"
  remote_spi = "0x100"
}


`, name)
}
