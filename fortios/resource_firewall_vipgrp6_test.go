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

func TestAccFortiOSFirewallVipgrp6_basic(t *testing.T) {
	rname := acctest.RandString(8)
	var0 := "var0" + rname
	log.Printf(var0)
	log.Printf("TestAccFortiOSFirewallVipgrp6_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallVipgrp6Config(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallVipgrp6Exists("fortios_firewall_vipgrp6.trname"),
					resource.TestCheckResourceAttr("fortios_firewall_vipgrp6.trname", "color", "0"),
					resource.TestCheckResourceAttr("fortios_firewall_vipgrp6.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewall_vipgrp6.trname", "member.0.name", var0),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallVipgrp6Exists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallVipgrp6: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallVipgrp6 is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallVipgrp6(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallVipgrp6: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallVipgrp6: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallVipgrp6Destroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_vipgrp6" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallVipgrp6(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallVipgrp6 %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallVipgrp6Config(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_vip6" "trname1" {
  arp_reply                        = "enable"
  color                            = 0
  extip                            = "2001:1:1:2::100"
  extport                          = "0-65535"
  fosid                            = 0
  http_cookie_age                  = 60
  http_cookie_domain_from_host     = "disable"
  http_cookie_generation           = 0
  http_cookie_share                = "same-ip"
  http_ip_header                   = "disable"
  http_multiplex                   = "disable"
  https_cookie_secure              = "disable"
  ldb_method                       = "static"
  mappedip                         = "2001:1:1:2::200"
  mappedport                       = "0-65535"
  max_embryonic_connections        = 1000
  name                             = "var0%[1]s"
  outlook_web_access               = "disable"
  persistence                      = "none"
  portforward                      = "disable"
  protocol                         = "tcp"
  ssl_algorithm                    = "high"
  ssl_client_fallback              = "enable"
  ssl_client_renegotiation         = "secure"
  ssl_client_session_state_max     = 1000
  ssl_client_session_state_timeout = 30
  ssl_client_session_state_type    = "both"
  ssl_dh_bits                      = "2048"
  ssl_hpkp                         = "disable"
  ssl_hpkp_age                     = 5184000
  ssl_hpkp_include_subdomains      = "disable"
  ssl_hsts                         = "disable"
  ssl_hsts_age                     = 5184000
  ssl_hsts_include_subdomains      = "disable"
  ssl_http_location_conversion     = "disable"
  ssl_http_match_host              = "enable"
  ssl_max_version                  = "tls-1.2"
  ssl_min_version                  = "tls-1.1"
  ssl_mode                         = "half"
  ssl_pfs                          = "require"
  ssl_send_empty_frags             = "enable"
  ssl_server_algorithm             = "client"
  ssl_server_max_version           = "client"
  ssl_server_min_version           = "client"
  ssl_server_session_state_max     = 100
  ssl_server_session_state_timeout = 60
  ssl_server_session_state_type    = "both"
  type                             = "static-nat"
  weblogic_server                  = "disable"
  websphere_server                 = "disable"
}

resource "fortios_firewall_vipgrp6" "trname" {
  color = 0
  name  = "%[1]s"

  member {
    name = fortios_firewall_vip6.trname1.name
  }
}

`, name)
}
