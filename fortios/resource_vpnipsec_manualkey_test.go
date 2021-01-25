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

func TestAccFortiOSVpnIpsecManualkey_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSVpnIpsecManualkey_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSVpnIpsecManualkeyConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSVpnIpsecManualkeyExists("fortios_vpnipsec_manualkey.trname"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_manualkey.trname", "authentication", "md5"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_manualkey.trname", "authkey", "EE32CA121ECD772A-ECACAABA212345EC"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_manualkey.trname", "enckey", "-"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_manualkey.trname", "encryption", "null"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_manualkey.trname", "interface", "port4"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_manualkey.trname", "local_gw", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_manualkey.trname", "localspi", "0x100"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_manualkey.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_vpnipsec_manualkey.trname", "remote_gw", "1.1.1.1"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_manualkey.trname", "remotespi", "0x100"),
				),
			},
		},
	})
}

func testAccCheckFortiOSVpnIpsecManualkeyExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found VpnIpsecManualkey: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No VpnIpsecManualkey is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadVpnIpsecManualkey(i)

		if err != nil {
			return fmt.Errorf("Error reading VpnIpsecManualkey: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating VpnIpsecManualkey: %s", n)
		}

		return nil
	}
}

func testAccCheckVpnIpsecManualkeyDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_vpnipsec_manualkey" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadVpnIpsecManualkey(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error VpnIpsecManualkey %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSVpnIpsecManualkeyConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_vpnipsec_manualkey" "trname" {
  authentication = "md5"
  authkey        = "EE32CA121ECD772A-ECACAABA212345EC"
  enckey         = "-"
  encryption     = "null"
  interface      = "port4"
  local_gw       = "0.0.0.0"
  localspi       = "0x100"
  name           = "%[1]s"
  remote_gw      = "1.1.1.1"
  remotespi      = "0x100"
}
`, name)
}
