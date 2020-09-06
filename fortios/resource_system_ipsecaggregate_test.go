// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

package fortios

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"log"
	"testing"
)

func TestAccFortiOSSystemIpsecAggregate_basic(t *testing.T) {
	rname := acctest.RandString(8)
	var0 := "var0" + rname
	log.Printf(var0)
	var1 := "var1" + rname
	log.Printf(var1)
	log.Printf("TestAccFortiOSSystemIpsecAggregate_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemIpsecAggregateConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemIpsecAggregateExists("fortios_system_ipsecaggregate.trname"),
					resource.TestCheckResourceAttr("fortios_system_ipsecaggregate.trname", "algorithm", "round-robin"),
					resource.TestCheckResourceAttr("fortios_system_ipsecaggregate.trname", "member.0.tunnel_name", var0),
					resource.TestCheckResourceAttr("fortios_system_ipsecaggregate.trname", "name", rname),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemIpsecAggregateExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemIpsecAggregate: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemIpsecAggregate is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemIpsecAggregate(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemIpsecAggregate: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemIpsecAggregate: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemIpsecAggregateDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_ipsecaggregate" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemIpsecAggregate(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemIpsecAggregate %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemIpsecAggregateConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_vpnipsec_phase1interface" "trname1" {
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
  name                      = "var0%[1]s"
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
  remote_gw                 = "2.2.2.2"
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

resource "fortios_vpnipsec_phase2interface" "trname1" {
  add_route                = "phase1"
  auto_discovery_forwarder = "phase1"
  auto_discovery_sender    = "phase1"
  auto_negotiate           = "disable"
  dhcp_ipsec               = "disable"
  dhgrp                    = "14 5"
  dst_addr_type            = "subnet"
  dst_end_ip6              = "::"
  dst_port                 = 0
  dst_subnet               = "0.0.0.0 0.0.0.0"
  encapsulation            = "tunnel-mode"
  keepalive                = "disable"
  keylife_type             = "seconds"
  keylifekbs               = 5120
  keylifeseconds           = 43200
  l2tp                     = "disable"
  name                     = "var1%[1]s"
  pfs                      = "enable"
  phase1name               = fortios_vpnipsec_phase1interface.trname1.name
  proposal                 = "aes128-sha1 aes256-sha1 aes128-sha256 aes256-sha256 aes128gcm aes256gcm chacha20poly1305"
  protocol                 = 0
  replay                   = "enable"
  route_overlap            = "use-new"
  single_source            = "disable"
  src_addr_type            = "subnet"
  src_end_ip6              = "::"
  src_port                 = 0
  src_subnet               = "0.0.0.0 0.0.0.0"
}

resource "fortios_system_ipsecaggregate" "trname" {
  algorithm = "round-robin"
  member {
    tunnel_name = fortios_vpnipsec_phase1interface.trname1.name
  }
  name = "%[1]s"
}
`, name)
}
