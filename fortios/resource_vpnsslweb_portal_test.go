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

func TestAccFortiOSVpnSslWebPortal_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSVpnSslWebPortal_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSVpnSslWebPortalConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSVpnSslWebPortalExists("fortios_vpnsslweb_portal.trname"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "allow_user_access", "web ftp smb sftp telnet ssh vnc rdp ping citrix portforward"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "auto_connect", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "customize_forticlient_download_url", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "display_bookmark", "enable"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "display_connection_tools", "enable"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "display_history", "enable"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "display_status", "enable"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "dns_server1", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "dns_server2", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "exclusive_routing", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "forticlient_download", "enable"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "forticlient_download_method", "direct"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "heading", "SSL-VPN Portal"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "hide_sso_credential", "enable"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "host_check", "none"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "ip_mode", "range"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "ipv6_dns_server1", "::"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "ipv6_dns_server2", "::"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "ipv6_exclusive_routing", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "ipv6_service_restriction", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "ipv6_split_tunneling", "enable"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "ipv6_tunnel_mode", "enable"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "ipv6_wins_server1", "::"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "ipv6_wins_server2", "::"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "keep_alive", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "limit_user_logins", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "mac_addr_action", "allow"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "mac_addr_check", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "os_check", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "save_password", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "service_restriction", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "skip_check_for_browser", "enable"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "skip_check_for_unsupported_os", "enable"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "smb_ntlmv1_auth", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "smbv1", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "split_tunneling", "enable"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "theme", "blue"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "tunnel_mode", "enable"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "user_bookmark", "enable"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "user_group_bookmark", "enable"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "web_mode", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "wins_server1", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "wins_server2", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "ip_pools.0.name", "SSLVPN_TUNNEL_ADDR1"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_portal.trname", "ipv6_pools.0.name", "SSLVPN_TUNNEL_IPv6_ADDR1"),
				),
			},
		},
	})
}

func testAccCheckFortiOSVpnSslWebPortalExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found VpnSslWebPortal: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No VpnSslWebPortal is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadVpnSslWebPortal(i)

		if err != nil {
			return fmt.Errorf("Error reading VpnSslWebPortal: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating VpnSslWebPortal: %s", n)
		}

		return nil
	}
}

func testAccCheckVpnSslWebPortalDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_vpnsslweb_portal" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadVpnSslWebPortal(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error VpnSslWebPortal %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSVpnSslWebPortalConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_vpnsslweb_portal" "trname" {
  allow_user_access                  = "web ftp smb sftp telnet ssh vnc rdp ping citrix portforward"
  auto_connect                       = "disable"
  customize_forticlient_download_url = "disable"
  display_bookmark                   = "enable"
  display_connection_tools           = "enable"
  display_history                    = "enable"
  display_status                     = "enable"
  dns_server1                        = "0.0.0.0"
  dns_server2                        = "0.0.0.0"
  exclusive_routing                  = "disable"
  forticlient_download               = "enable"
  forticlient_download_method        = "direct"
  heading                            = "SSL-VPN Portal"
  hide_sso_credential                = "enable"
  host_check                         = "none"
  ip_mode                            = "range"
  ipv6_dns_server1                   = "::"
  ipv6_dns_server2                   = "::"
  ipv6_exclusive_routing             = "disable"
  ipv6_service_restriction           = "disable"
  ipv6_split_tunneling               = "enable"
  ipv6_tunnel_mode                   = "enable"
  ipv6_wins_server1                  = "::"
  ipv6_wins_server2                  = "::"
  keep_alive                         = "disable"
  limit_user_logins                  = "disable"
  mac_addr_action                    = "allow"
  mac_addr_check                     = "disable"
  name                               = "%[1]s"
  os_check                           = "disable"
  save_password                      = "disable"
  service_restriction                = "disable"
  skip_check_for_browser             = "enable"
  skip_check_for_unsupported_os      = "enable"
  smb_ntlmv1_auth                    = "disable"
  smbv1                              = "disable"
  split_tunneling                    = "enable"
  theme                              = "blue"
  tunnel_mode                        = "enable"
  user_bookmark                      = "enable"
  user_group_bookmark                = "enable"
  web_mode                           = "disable"
  wins_server1                       = "0.0.0.0"
  wins_server2                       = "0.0.0.0"

  ip_pools {
    name = "SSLVPN_TUNNEL_ADDR1"
  }

  ipv6_pools {
    name = "SSLVPN_TUNNEL_IPv6_ADDR1"
  }
}
`, name)
}
