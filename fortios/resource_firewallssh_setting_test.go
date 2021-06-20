// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

package fortios

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccFortiOSFirewallSshSetting_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSFirewallSshSetting_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallSshSettingConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallSshSettingExists("fortios_firewallssh_setting.trname"),
					resource.TestCheckResourceAttr("fortios_firewallssh_setting.trname", "caname", "Fortinet_SSH_CA"),
					resource.TestCheckResourceAttr("fortios_firewallssh_setting.trname", "host_trusted_checking", "enable"),
					resource.TestCheckResourceAttr("fortios_firewallssh_setting.trname", "hostkey_dsa1024", "Fortinet_SSH_DSA1024"),
					resource.TestCheckResourceAttr("fortios_firewallssh_setting.trname", "hostkey_ecdsa256", "Fortinet_SSH_ECDSA256"),
					resource.TestCheckResourceAttr("fortios_firewallssh_setting.trname", "hostkey_ecdsa384", "Fortinet_SSH_ECDSA384"),
					resource.TestCheckResourceAttr("fortios_firewallssh_setting.trname", "hostkey_ecdsa521", "Fortinet_SSH_ECDSA521"),
					resource.TestCheckResourceAttr("fortios_firewallssh_setting.trname", "hostkey_ed25519", "Fortinet_SSH_ED25519"),
					resource.TestCheckResourceAttr("fortios_firewallssh_setting.trname", "hostkey_rsa2048", "Fortinet_SSH_RSA2048"),
					resource.TestCheckResourceAttr("fortios_firewallssh_setting.trname", "untrusted_caname", "Fortinet_SSH_CA_Untrusted"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallSshSettingExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallSshSetting: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallSshSetting is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallSshSetting(i, "root")

		if err != nil {
			return fmt.Errorf("Error reading FirewallSshSetting: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallSshSetting: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallSshSettingDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewallssh_setting" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallSshSetting(i, "root")

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallSshSetting %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallSshSettingConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewallssh_setting" "trname" {
  caname                = "Fortinet_SSH_CA"
  host_trusted_checking = "enable"
  hostkey_dsa1024       = "Fortinet_SSH_DSA1024"
  hostkey_ecdsa256      = "Fortinet_SSH_ECDSA256"
  hostkey_ecdsa384      = "Fortinet_SSH_ECDSA384"
  hostkey_ecdsa521      = "Fortinet_SSH_ECDSA521"
  hostkey_ed25519       = "Fortinet_SSH_ED25519"
  hostkey_rsa2048       = "Fortinet_SSH_RSA2048"
  untrusted_caname      = "Fortinet_SSH_CA_Untrusted"
}
`)
}
