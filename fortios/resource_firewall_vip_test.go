
// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

package fortios
import (
    "fmt"
	"log"
    "testing"
    "github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
    "github.com/hashicorp/terraform-plugin-sdk/helper/resource"
    "github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccFortiOSFirewallVip_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSFirewallVip_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSFirewallVipConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSFirewallVipExists("fortios_firewall_vip.trname"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "arp_reply", "enable"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "color", "0"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "dns_mapping_ttl", "0"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "extintf", "any"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "extip", "1.0.0.1-1.0.0.2"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "extport", "0-65535"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "fosid", "0"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "http_cookie_age", "60"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "http_cookie_domain_from_host", "disable"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "http_cookie_generation", "0"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "http_cookie_share", "same-ip"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "http_ip_header", "disable"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "http_multiplex", "disable"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "https_cookie_secure", "disable"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "ldb_method", "static"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "mappedport", "0-65535"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "max_embryonic_connections", "1000"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "name", rname),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "nat_source_vip", "disable"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "outlook_web_access", "disable"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "persistence", "none"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "portforward", "disable"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "portmapping_type", "1-to-1"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "protocol", "tcp"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "ssl_algorithm", "high"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "ssl_client_fallback", "enable"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "ssl_client_renegotiation", "secure"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "ssl_client_session_state_max", "1000"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "ssl_client_session_state_timeout", "30"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "ssl_client_session_state_type", "both"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "ssl_dh_bits", "2048"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "ssl_hpkp", "disable"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "ssl_hpkp_age", "5184000"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "ssl_hpkp_include_subdomains", "disable"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "ssl_hsts", "disable"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "ssl_hsts_age", "5184000"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "ssl_hsts_include_subdomains", "disable"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "ssl_http_location_conversion", "disable"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "ssl_http_match_host", "enable"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "ssl_max_version", "tls-1.2"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "ssl_min_version", "tls-1.1"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "ssl_mode", "half"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "ssl_pfs", "require"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "ssl_send_empty_frags", "enable"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "ssl_server_algorithm", "client"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "ssl_server_max_version", "client"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "ssl_server_min_version", "client"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "ssl_server_session_state_max", "100"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "ssl_server_session_state_timeout", "60"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "ssl_server_session_state_type", "both"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "type", "static-nat"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "weblogic_server", "disable"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "websphere_server", "disable"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip.trname", "mappedip.0.range", "3.0.0.0-3.0.0.1"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSFirewallVipExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallVip: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallVip is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallVip(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallVip: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallVip: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallVipDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_vip" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallVip(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallVip %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallVipConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_vip" "trname" {
  arp_reply                        = "enable"
  color                            = 0
  dns_mapping_ttl                  = 0
  extintf                          = "any"
  extip                            = "1.0.0.1-1.0.0.2"
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
  mappedport                       = "0-65535"
  max_embryonic_connections        = 1000
  name                             = "%[1]s"
  nat_source_vip                   = "disable"
  outlook_web_access               = "disable"
  persistence                      = "none"
  portforward                      = "disable"
  portmapping_type                 = "1-to-1"
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

  mappedip {
    range = "3.0.0.0-3.0.0.1"
  }
}




`, name)
}
