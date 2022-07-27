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

func TestAccFortiOSVpnIpsecPhase1_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSVpnIpsecPhase1_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSVpnIpsecPhase1Config(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSVpnIpsecPhase1Exists("fortios_vpnipsec_phase1.trnamex1"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "acct_verify", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "add_gw_route", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "add_route", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "assign_ip", "enable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "assign_ip_from", "range"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "authmethod", "psk"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "auto_negotiate", "enable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "cert_id_validation", "enable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "childless_ike", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "client_auto_negotiate", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "client_keep_alive", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "dhgrp", "14 5"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "digital_signature_auth", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "distance", "15"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "dns_mode", "manual"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "dpd", "on-demand"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "dpd_retrycount", "3"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "dpd_retryinterval", "20"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "eap", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "eap_identity", "use-id-payload"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "enforce_unique_id", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "forticlient_enforcement", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "fragmentation", "enable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "fragmentation_mtu", "1200"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "group_authentication", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "ha_sync_esp_seqno", "enable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "idle_timeout", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "idle_timeoutinterval", "15"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "ike_version", "1"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "include_local_lan", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "interface", "port4"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "ipv4_dns_server1", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "ipv4_dns_server2", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "ipv4_dns_server3", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "ipv4_end_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "ipv4_netmask", "255.255.255.255"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "ipv4_start_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "ipv4_wins_server1", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "ipv4_wins_server2", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "ipv6_dns_server1", "::"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "ipv6_dns_server2", "::"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "ipv6_dns_server3", "::"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "ipv6_end_ip", "::"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "ipv6_prefix", "128"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "ipv6_start_ip", "::"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "keepalive", "10"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "keylife", "86400"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "local_gw", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "localid_type", "auto"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "mesh_selector_type", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "mode", "main"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "mode_cfg", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "name", rname),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "nattraversal", "enable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "negotiate_timeout", "30"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "peertype", "any"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "ppk", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "priority", "0"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "proposal", "aes128-sha256 aes256-sha256 aes128-sha1 aes256-sha1"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "psksecret", "dewcEde2112"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "reauth", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "rekey", "enable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "remote_gw", "1.1.1.1"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "rsa_signature_format", "pkcs1"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "save_password", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "send_cert_chain", "enable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "signature_hash_alg", "sha2-512 sha2-384 sha2-256 sha1"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "suite_b", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "type", "static"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "unity_support", "enable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "wizard_type", "custom"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1.trnamex1", "xauthtype", "disable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSVpnIpsecPhase1Exists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found VpnIpsecPhase1: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No VpnIpsecPhase1 is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadVpnIpsecPhase1(i)

		if err != nil {
			return fmt.Errorf("Error reading VpnIpsecPhase1: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating VpnIpsecPhase1: %s", n)
		}

		return nil
	}
}

func testAccCheckVpnIpsecPhase1Destroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_vpnipsec_phase1" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadVpnIpsecPhase1(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error VpnIpsecPhase1 %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSVpnIpsecPhase1Config(name string) string {
	return fmt.Sprintf(`
resource "fortios_vpnipsec_phase1" "trnamex1" {
  acct_verify             = "disable"
  add_gw_route            = "disable"
  add_route               = "disable"
  assign_ip               = "enable"
  assign_ip_from          = "range"
  authmethod              = "psk"
  auto_negotiate          = "enable"
  cert_id_validation      = "enable"
  childless_ike           = "disable"
  client_auto_negotiate   = "disable"
  client_keep_alive       = "disable"
  dhgrp                   = "14 5"
  digital_signature_auth  = "disable"
  distance                = 15
  dns_mode                = "manual"
  dpd                     = "on-demand"
  dpd_retrycount          = 3
  dpd_retryinterval       = "20"
  eap                     = "disable"
  eap_identity            = "use-id-payload"
  enforce_unique_id       = "disable"
  forticlient_enforcement = "disable"
  fragmentation           = "enable"
  fragmentation_mtu       = 1200
  group_authentication    = "disable"
  ha_sync_esp_seqno       = "enable"
  idle_timeout            = "disable"
  idle_timeoutinterval    = 15
  ike_version             = "1"
  include_local_lan       = "disable"
  interface               = "port4"
  ipv4_dns_server1        = "0.0.0.0"
  ipv4_dns_server2        = "0.0.0.0"
  ipv4_dns_server3        = "0.0.0.0"
  ipv4_end_ip             = "0.0.0.0"
  ipv4_netmask            = "255.255.255.255"
  ipv4_start_ip           = "0.0.0.0"
  ipv4_wins_server1       = "0.0.0.0"
  ipv4_wins_server2       = "0.0.0.0"
  ipv6_dns_server1        = "::"
  ipv6_dns_server2        = "::"
  ipv6_dns_server3        = "::"
  ipv6_end_ip             = "::"
  ipv6_prefix             = 128
  ipv6_start_ip           = "::"
  keepalive               = 10
  keylife                 = 86400
  local_gw                = "0.0.0.0"
  localid_type            = "auto"
  mesh_selector_type      = "disable"
  mode                    = "main"
  mode_cfg                = "disable"
  name                    = "%[1]s"
  nattraversal            = "enable"
  negotiate_timeout       = 30
  peertype                = "any"
  ppk                     = "disable"
  priority                = 0
  proposal                = "aes128-sha256 aes256-sha256 aes128-sha1 aes256-sha1"
  psksecret               = "dewcEde2112"
  reauth                  = "disable"
  rekey                   = "enable"
  remote_gw               = "1.1.1.1"
  rsa_signature_format    = "pkcs1"
  save_password           = "disable"
  send_cert_chain         = "enable"
  signature_hash_alg      = "sha2-512 sha2-384 sha2-256 sha1"
  suite_b                 = "disable"
  type                    = "static"
  unity_support           = "enable"
  wizard_type             = "custom"
  xauthtype               = "disable"
}
`, name)
}
