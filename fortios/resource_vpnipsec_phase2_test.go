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

func TestAccFortiOSVpnIpsecPhase2_basic(t *testing.T) {
	rname := acctest.RandString(8)
	var0 := "var0" + rname
	log.Printf(var0)
	log.Printf("TestAccFortiOSVpnIpsecPhase2_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSVpnIpsecPhase2Config(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSVpnIpsecPhase2Exists("fortios_vpnipsec_phase2.trname"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase2.trname", "add_route", "phase1"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase2.trname", "auto_negotiate", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase2.trname", "dhcp_ipsec", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase2.trname", "dhgrp", "14 5"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase2.trname", "dst_addr_type", "subnet"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase2.trname", "dst_end_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase2.trname", "dst_end_ip6", "::"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase2.trname", "dst_port", "0"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase2.trname", "dst_start_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase2.trname", "dst_start_ip6", "::"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase2.trname", "dst_subnet", "0.0.0.0 0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase2.trname", "dst_subnet6", "::/0"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase2.trname", "encapsulation", "tunnel-mode"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase2.trname", "keepalive", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase2.trname", "keylife_type", "seconds"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase2.trname", "keylifekbs", "5120"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase2.trname", "keylifeseconds", "43200"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase2.trname", "l2tp", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase2.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase2.trname", "pfs", "enable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase2.trname", "phase1name", var0),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase2.trname", "proposal", "null-md5 null-sha1 null-sha256"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase2.trname", "protocol", "0"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase2.trname", "replay", "enable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase2.trname", "route_overlap", "use-new"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase2.trname", "selector_match", "auto"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase2.trname", "single_source", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase2.trname", "src_addr_type", "subnet"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase2.trname", "src_end_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase2.trname", "src_end_ip6", "::"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase2.trname", "src_port", "0"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase2.trname", "src_start_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase2.trname", "src_start_ip6", "::"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase2.trname", "src_subnet", "0.0.0.0 0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase2.trname", "src_subnet6", "::/0"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_phase2.trname", "use_natip", "disable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSVpnIpsecPhase2Exists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found VpnIpsecPhase2: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No VpnIpsecPhase2 is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadVpnIpsecPhase2(i)

		if err != nil {
			return fmt.Errorf("Error reading VpnIpsecPhase2: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating VpnIpsecPhase2: %s", n)
		}

		return nil
	}
}

func testAccCheckVpnIpsecPhase2Destroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_vpnipsec_phase2" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadVpnIpsecPhase2(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error VpnIpsecPhase2 %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSVpnIpsecPhase2Config(name string) string {
	return fmt.Sprintf(`
resource "fortios_vpnipsec_phase1" "trnamex2" {
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
  name                    = "var0%[1]s"
  nattraversal            = "enable"
  negotiate_timeout       = 30
  peertype                = "any"
  ppk                     = "disable"
  priority                = 0
  proposal                = "aes128-sha256 aes256-sha256 aes128-sha1 aes256-sha1"
  psksecret               = "dewcEde2112"
  reauth                  = "disable"
  rekey                   = "enable"
  remote_gw               = "2.1.1.1"
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

resource "fortios_vpnipsec_phase2" "trname" {
  add_route      = "phase1"
  auto_negotiate = "disable"
  dhcp_ipsec     = "disable"
  dhgrp          = "14 5"
  dst_addr_type  = "subnet"
  dst_end_ip     = "0.0.0.0"
  dst_end_ip6    = "::"
  dst_port       = 0
  dst_start_ip   = "0.0.0.0"
  dst_start_ip6  = "::"
  dst_subnet     = "0.0.0.0 0.0.0.0"
  dst_subnet6    = "::/0"
  encapsulation  = "tunnel-mode"
  keepalive      = "disable"
  keylife_type   = "seconds"
  keylifekbs     = 5120
  keylifeseconds = 43200
  l2tp           = "disable"
  name           = "%[1]s"
  pfs            = "enable"
  phase1name     = fortios_vpnipsec_phase1.trnamex2.name
  proposal       = "null-md5 null-sha1 null-sha256"
  protocol       = 0
  replay         = "enable"
  route_overlap  = "use-new"
  selector_match = "auto"
  single_source  = "disable"
  src_addr_type  = "subnet"
  src_end_ip     = "0.0.0.0"
  src_end_ip6    = "::"
  src_port       = 0
  src_start_ip   = "0.0.0.0"
  src_start_ip6  = "::"
  src_subnet     = "0.0.0.0 0.0.0.0"
  src_subnet6    = "::/0"
  use_natip      = "disable"


}

`, name)
}
