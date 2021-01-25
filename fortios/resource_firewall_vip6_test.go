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

func TestAccFortiOSFirewallVip6_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSFirewallVip6_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallVip6Config(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallVip6Exists("fortios_firewall_vip6.trname"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "arp_reply", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "color", "0"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "extip", "2001:1:1:12::100"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "extport", "0-65535"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "fosid", "0"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "http_cookie_age", "60"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "http_cookie_domain_from_host", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "http_cookie_generation", "0"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "http_cookie_share", "same-ip"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "http_ip_header", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "http_multiplex", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "https_cookie_secure", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "ldb_method", "static"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "mappedip", "2001:1:1:12::200"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "mappedport", "0-65535"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "max_embryonic_connections", "1000"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "outlook_web_access", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "persistence", "none"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "portforward", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "protocol", "tcp"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "ssl_algorithm", "high"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "ssl_client_fallback", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "ssl_client_renegotiation", "secure"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "ssl_client_session_state_max", "1000"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "ssl_client_session_state_timeout", "30"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "ssl_client_session_state_type", "both"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "ssl_dh_bits", "2048"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "ssl_hpkp", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "ssl_hpkp_age", "5184000"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "ssl_hpkp_include_subdomains", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "ssl_hsts", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "ssl_hsts_age", "5184000"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "ssl_hsts_include_subdomains", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "ssl_http_location_conversion", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "ssl_http_match_host", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "ssl_max_version", "tls-1.2"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "ssl_min_version", "tls-1.1"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "ssl_mode", "half"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "ssl_pfs", "require"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "ssl_send_empty_frags", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "ssl_server_algorithm", "client"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "ssl_server_max_version", "client"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "ssl_server_min_version", "client"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "ssl_server_session_state_max", "100"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "ssl_server_session_state_timeout", "60"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "ssl_server_session_state_type", "both"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "type", "static-nat"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "weblogic_server", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_vip6.trname", "websphere_server", "disable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallVip6Exists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallVip6: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallVip6 is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallVip6(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallVip6: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallVip6: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallVip6Destroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_vip6" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallVip6(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallVip6 %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallVip6Config(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_vip6" "trname" {
  arp_reply                        = "enable"
  color                            = 0
  extip                            = "2001:1:1:12::100"
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
  mappedip                         = "2001:1:1:12::200"
  mappedport                       = "0-65535"
  max_embryonic_connections        = 1000
  name                             = "%[1]s"
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
`, name)
}
