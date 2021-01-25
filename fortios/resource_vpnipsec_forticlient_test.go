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

func TestAccFortiOSVpnIpsecForticlient_basic(t *testing.T) {
	rname := acctest.RandString(8)
	var0 := "var0" + rname
	log.Printf(var0)
	var1 := "var1" + rname
	log.Printf(var1)
	log.Printf("TestAccFortiOSVpnIpsecForticlient_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSVpnIpsecForticlientConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSVpnIpsecForticlientExists("fortios_vpnipsec_forticlient.trname"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_forticlient.trname", "phase2name", var1),
					resource.TestCheckResourceAttr("fortios_vpnipsec_forticlient.trname", "realm", "1"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_forticlient.trname", "status", "enable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_forticlient.trname", "usergroupname", "Guest-group"),
				),
			},
		},
	})
}

func testAccCheckFortiOSVpnIpsecForticlientExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found VpnIpsecForticlient: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No VpnIpsecForticlient is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadVpnIpsecForticlient(i)

		if err != nil {
			return fmt.Errorf("Error reading VpnIpsecForticlient: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating VpnIpsecForticlient: %s", n)
		}

		return nil
	}
}

func testAccCheckVpnIpsecForticlientDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_vpnipsec_forticlient" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadVpnIpsecForticlient(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error VpnIpsecForticlient %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSVpnIpsecForticlientConfig(name string) string {
	return fmt.Sprintf(`
# fortios_vpnipsec_phase1interface.trname2:
resource "fortios_vpnipsec_phase1interface" "trname4" {
  acct_verify               = "disable"
  add_gw_route              = "disable"
  add_route                 = "enable"
  assign_ip                 = "enable"
  assign_ip_from            = "range"
  authmethod                = "psk"
  authusrgrp                = "Guest-group"
  auto_discovery_forwarder  = "disable"
  auto_discovery_psk        = "disable"
  auto_discovery_receiver   = "disable"
  auto_discovery_sender     = "disable"
  auto_negotiate            = "enable"
  cert_id_validation        = "enable"
  childless_ike             = "disable"
  client_auto_negotiate     = "disable"
  client_keep_alive         = "disable"
  comments                  = "VPN: Dialup_IPsec (Created by VPN wizard)"
  default_gw                = "0.0.0.0"
  default_gw_priority       = 0
  dhgrp                     = "14 5"
  digital_signature_auth    = "disable"
  distance                  = 15
  dns_mode                  = "auto"
  dpd                       = "on-idle"
  dpd_retrycount            = 3
  dpd_retryinterval         = "60"
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
  interface                 = "port4"
  ip_version                = "4"
  ipv4_dns_server1          = "0.0.0.0"
  ipv4_dns_server2          = "0.0.0.0"
  ipv4_dns_server3          = "0.0.0.0"
  ipv4_end_ip               = "10.10.10.10"
  ipv4_netmask              = "255.255.255.192"
  ipv4_split_include        = "FIREWALL_AUTH_PORTAL_ADDRESS"
  ipv4_start_ip             = "10.10.10.1"
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
  mode                      = "aggressive"
  mode_cfg                  = "enable"
  monitor_hold_down_delay   = 0
  monitor_hold_down_time    = "00:00"
  monitor_hold_down_type    = "immediate"
  monitor_hold_down_weekday = "sunday"
  name                      = "var0%[1]s"
  nattraversal              = "enable"
  negotiate_timeout         = 30
  net_device                = "enable"
  passive_mode              = "disable"
  peertype                  = "any"
  psksecret                 = "NCIEW32930293203932"
  ppk                       = "disable"
  priority                  = 0
  proposal                  = "aes128-sha256 aes256-sha256 aes128-sha1 aes256-sha1"
  reauth                    = "disable"
  rekey                     = "enable"
  remote_gw                 = "0.0.0.0"
  remote_gw6                = "::"
  rsa_signature_format      = "pkcs1"
  save_password             = "enable"
  send_cert_chain           = "enable"
  signature_hash_alg        = "sha2-512 sha2-384 sha2-256 sha1"
  suite_b                   = "disable"
  tunnel_search             = "selectors"
  type                      = "dynamic"
  unity_support             = "enable"
  wizard_type               = "dialup-forticlient"
  xauthtype                 = "auto"
}

# fortios_vpnipsec_phase2interface.trname1:
resource "fortios_vpnipsec_phase2interface" "trname3" {
  add_route                = "phase1"
  auto_discovery_forwarder = "phase1"
  auto_discovery_sender    = "phase1"
  auto_negotiate           = "disable"
  dhcp_ipsec               = "disable"
  dhgrp                    = "14 5"
  dst_addr_type            = "subnet"
  dst_end_ip               = "0.0.0.0"
  dst_end_ip6              = "::"
  dst_port                 = 0
  dst_start_ip             = "0.0.0.0"
  dst_start_ip6            = "::"
  dst_subnet               = "0.0.0.0 0.0.0.0"
  dst_subnet6              = "::/0"
  encapsulation            = "tunnel-mode"
  keepalive                = "disable"
  keylife_type             = "seconds"
  keylifekbs               = 5120
  keylifeseconds           = 43200
  l2tp                     = "disable"
  name                     = "var1%[1]s"
  pfs                      = "enable"
  phase1name               = fortios_vpnipsec_phase1interface.trname4.name
  proposal                 = "aes128-sha1 aes256-sha1 aes128-sha256 aes256-sha256 aes128gcm aes256gcm chacha20poly1305"
  protocol                 = 0
  replay                   = "enable"
  route_overlap            = "use-new"
  single_source            = "disable"
  src_addr_type            = "subnet"
  src_end_ip               = "0.0.0.0"
  src_end_ip6              = "::"
  src_port                 = 0
  src_start_ip             = "0.0.0.0"
  src_start_ip6            = "::"
  src_subnet               = "0.0.0.0 0.0.0.0"
  src_subnet6              = "::/0"
}

resource "fortios_vpnipsec_forticlient" "trname" {
  phase2name    = fortios_vpnipsec_phase2interface.trname3.name
  realm         = "1"
  status        = "enable"
  usergroupname = "Guest-group"
}
`, name)
}
