// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

package fortios

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"log"
	"testing"
)

func TestAccFortiOSFirewallSslSetting_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSFirewallSslSetting_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallSslSettingConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallSslSettingExists("fortios_firewallssl_setting.trname"),
					resource.TestCheckResourceAttr("fortios_firewallssl_setting.trname", "abbreviate_handshake", "enable"),
					resource.TestCheckResourceAttr("fortios_firewallssl_setting.trname", "cert_cache_capacity", "200"),
					resource.TestCheckResourceAttr("fortios_firewallssl_setting.trname", "cert_cache_timeout", "10"),
					resource.TestCheckResourceAttr("fortios_firewallssl_setting.trname", "kxp_queue_threshold", "16"),
					resource.TestCheckResourceAttr("fortios_firewallssl_setting.trname", "no_matching_cipher_action", "bypass"),
					resource.TestCheckResourceAttr("fortios_firewallssl_setting.trname", "proxy_connect_timeout", "30"),
					resource.TestCheckResourceAttr("fortios_firewallssl_setting.trname", "session_cache_capacity", "500"),
					resource.TestCheckResourceAttr("fortios_firewallssl_setting.trname", "session_cache_timeout", "20"),
					resource.TestCheckResourceAttr("fortios_firewallssl_setting.trname", "ssl_dh_bits", "2048"),
					resource.TestCheckResourceAttr("fortios_firewallssl_setting.trname", "ssl_queue_threshold", "32"),
					resource.TestCheckResourceAttr("fortios_firewallssl_setting.trname", "ssl_send_empty_frags", "enable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallSslSettingExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallSslSetting: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallSslSetting is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallSslSetting(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallSslSetting: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallSslSetting: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallSslSettingDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewallssl_setting" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallSslSetting(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallSslSetting %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallSslSettingConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewallssl_setting" "trname" {
  abbreviate_handshake      = "enable"
  cert_cache_capacity       = 200
  cert_cache_timeout        = 10
  kxp_queue_threshold       = 16
  no_matching_cipher_action = "bypass"
  proxy_connect_timeout     = 30
  session_cache_capacity    = 500
  session_cache_timeout     = 20
  ssl_dh_bits               = "2048"
  ssl_queue_threshold       = 32
  ssl_send_empty_frags      = "enable"
}
`)
}
