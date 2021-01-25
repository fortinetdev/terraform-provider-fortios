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

func TestAccFortiOSVpnIpsecPhase1Interface_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSVpnIpsecPhase1Interface_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSVpnIpsecPhase1InterfaceConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSVpnIpsecPhase1InterfaceExists("fortios_vpnipsec_phase1interface.trname2"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "acct_verify", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "add_gw_route", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "add_route", "enable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "assign_ip", "enable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "assign_ip_from", "range"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "authmethod", "psk"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "auto_discovery_forwarder", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "auto_discovery_psk", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "auto_discovery_receiver", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "auto_discovery_sender", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "auto_negotiate", "enable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "cert_id_validation", "enable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "childless_ike", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "client_auto_negotiate", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "client_keep_alive", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "default_gw", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "default_gw_priority", "0"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "dhgrp", "14 5"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "digital_signature_auth", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "distance", "15"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "dns_mode", "manual"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "dpd", "on-demand"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "dpd_retrycount", "3"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "dpd_retryinterval", "20"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "eap", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "eap_identity", "use-id-payload"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "encap_local_gw4", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "encap_local_gw6", "::"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "encap_remote_gw4", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "encap_remote_gw6", "::"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "encapsulation", "none"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "encapsulation_address", "ike"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "enforce_unique_id", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "exchange_interface_ip", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "exchange_ip_addr4", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "exchange_ip_addr6", "::"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "forticlient_enforcement", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "fragmentation", "enable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "fragmentation_mtu", "1200"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "group_authentication", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "ha_sync_esp_seqno", "enable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "idle_timeout", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "idle_timeoutinterval", "15"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "ike_version", "1"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "include_local_lan", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "interface", "port3"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "ip_version", "4"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "ipv4_dns_server1", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "ipv4_dns_server2", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "ipv4_dns_server3", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "ipv4_end_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "ipv4_netmask", "255.255.255.255"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "ipv4_start_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "ipv4_wins_server1", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "ipv4_wins_server2", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "ipv6_dns_server1", "::"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "ipv6_dns_server2", "::"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "ipv6_dns_server3", "::"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "ipv6_end_ip", "::"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "ipv6_prefix", "128"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "ipv6_start_ip", "::"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "keepalive", "10"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "keylife", "86400"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "local_gw", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "local_gw6", "::"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "localid_type", "auto"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "mesh_selector_type", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "mode", "main"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "mode_cfg", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "monitor_hold_down_delay", "0"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "monitor_hold_down_time", "00:00"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "monitor_hold_down_type", "immediate"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "monitor_hold_down_weekday", "sunday"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "name", rname),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "nattraversal", "enable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "negotiate_timeout", "30"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "net_device", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "passive_mode", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "peertype", "any"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "ppk", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "priority", "0"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "proposal", "aes128-sha256 aes256-sha256 aes128-sha1 aes256-sha1"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "psksecret", "eweeeeeeeecee"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "reauth", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "rekey", "enable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "remote_gw", "102.2.2.12"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "remote_gw6", "::"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "rsa_signature_format", "pkcs1"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "save_password", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "send_cert_chain", "enable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "signature_hash_alg", "sha2-512 sha2-384 sha2-256 sha1"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "suite_b", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "tunnel_search", "selectors"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "type", "static"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "unity_support", "enable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "wizard_type", "custom"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase1interface.trname2", "xauthtype", "disable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSVpnIpsecPhase1InterfaceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found VpnIpsecPhase1Interface: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No VpnIpsecPhase1Interface is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadVpnIpsecPhase1Interface(i)

		if err != nil {
			return fmt.Errorf("Error reading VpnIpsecPhase1Interface: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating VpnIpsecPhase1Interface: %s", n)
		}

		return nil
	}
}

func testAccCheckVpnIpsecPhase1InterfaceDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_vpnipsec_phase1interface" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadVpnIpsecPhase1Interface(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error VpnIpsecPhase1Interface %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSVpnIpsecPhase1InterfaceConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_vpnipsec_phase1interface" "trname2" {
  acct_verify               = "disable"
  add_gw_route              = "disable"
  add_route                 = "enable"
  assign_ip                 = "enable"
  assign_ip_from            = "range"
  authmethod                = "psk"
  auto_discovery_forwarder  = "disable"
  auto_discovery_psk        = "disable"
  auto_discovery_receiver   = "disable"
  auto_discovery_sender     = "disable"
  auto_negotiate            = "enable"
  cert_id_validation        = "enable"
  childless_ike             = "disable"
  client_auto_negotiate     = "disable"
  client_keep_alive         = "disable"
  default_gw                = "0.0.0.0"
  default_gw_priority       = 0
  dhgrp                     = "14 5"
  digital_signature_auth    = "disable"
  distance                  = 15
  dns_mode                  = "manual"
  dpd                       = "on-demand"
  dpd_retrycount            = 3
  dpd_retryinterval         = "20"
  eap                       = "disable"
  eap_identity              = "use-id-payload"
  encap_local_gw4           = "0.0.0.0"
  encap_local_gw6           = "::"
  encap_remote_gw4          = "0.0.0.0"
  encap_remote_gw6          = "::"
  encapsulation             = "none"
  encapsulation_address     = "ike"
  enforce_unique_id         = "disable"
  exchange_interface_ip     = "disable"
  exchange_ip_addr4         = "0.0.0.0"
  exchange_ip_addr6         = "::"
  forticlient_enforcement   = "disable"
  fragmentation             = "enable"
  fragmentation_mtu         = 1200
  group_authentication      = "disable"
  ha_sync_esp_seqno         = "enable"
  idle_timeout              = "disable"
  idle_timeoutinterval      = 15
  ike_version               = "1"
  include_local_lan         = "disable"
  interface                 = "port3"
  ip_version                = "4"
  ipv4_dns_server1          = "0.0.0.0"
  ipv4_dns_server2          = "0.0.0.0"
  ipv4_dns_server3          = "0.0.0.0"
  ipv4_end_ip               = "0.0.0.0"
  ipv4_netmask              = "255.255.255.255"
  ipv4_start_ip             = "0.0.0.0"
  ipv4_wins_server1         = "0.0.0.0"
  ipv4_wins_server2         = "0.0.0.0"
  ipv6_dns_server1          = "::"
  ipv6_dns_server2          = "::"
  ipv6_dns_server3          = "::"
  ipv6_end_ip               = "::"
  ipv6_prefix               = 128
  ipv6_start_ip             = "::"
  keepalive                 = 10
  keylife                   = 86400
  local_gw                  = "0.0.0.0"
  local_gw6                 = "::"
  localid_type              = "auto"
  mesh_selector_type        = "disable"
  mode                      = "main"
  mode_cfg                  = "disable"
  monitor_hold_down_delay   = 0
  monitor_hold_down_time    = "00:00"
  monitor_hold_down_type    = "immediate"
  monitor_hold_down_weekday = "sunday"
  name                      = "%[1]s"
  nattraversal              = "enable"
  negotiate_timeout         = 30
  net_device                = "disable"
  passive_mode              = "disable"
  peertype                  = "any"
  ppk                       = "disable"
  priority                  = 0
  proposal                  = "aes128-sha256 aes256-sha256 aes128-sha1 aes256-sha1"
  psksecret                 = "eweeeeeeeecee"
  reauth                    = "disable"
  rekey                     = "enable"
  remote_gw                 = "102.2.2.12"
  remote_gw6                = "::"
  rsa_signature_format      = "pkcs1"
  save_password             = "disable"
  send_cert_chain           = "enable"
  signature_hash_alg        = "sha2-512 sha2-384 sha2-256 sha1"
  suite_b                   = "disable"
  tunnel_search             = "selectors"
  type                      = "static"
  unity_support             = "enable"
  wizard_type               = "custom"
  xauthtype                 = "disable"
}
`, name)
}
