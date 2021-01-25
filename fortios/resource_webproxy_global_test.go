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

func TestAccFortiOSWebProxyGlobal_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSWebProxyGlobal_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSWebProxyGlobalConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSWebProxyGlobalExists("fortios_webproxy_global.trname"),
					resource.TestCheckResourceAttr("fortios_webproxy_global.trname", "fast_policy_match", "enable"),
					resource.TestCheckResourceAttr("fortios_webproxy_global.trname", "forward_proxy_auth", "disable"),
					resource.TestCheckResourceAttr("fortios_webproxy_global.trname", "forward_server_affinity_timeout", "30"),
					resource.TestCheckResourceAttr("fortios_webproxy_global.trname", "learn_client_ip", "disable"),
					resource.TestCheckResourceAttr("fortios_webproxy_global.trname", "max_message_length", "32"),
					resource.TestCheckResourceAttr("fortios_webproxy_global.trname", "max_request_length", "4"),
					resource.TestCheckResourceAttr("fortios_webproxy_global.trname", "max_waf_body_cache_length", "32"),
					resource.TestCheckResourceAttr("fortios_webproxy_global.trname", "proxy_fqdn", "default.fqdn"),
					resource.TestCheckResourceAttr("fortios_webproxy_global.trname", "ssl_ca_cert", "Fortinet_CA_SSL"),
					resource.TestCheckResourceAttr("fortios_webproxy_global.trname", "ssl_cert", "Fortinet_Factory"),
					resource.TestCheckResourceAttr("fortios_webproxy_global.trname", "strict_web_check", "disable"),
					resource.TestCheckResourceAttr("fortios_webproxy_global.trname", "tunnel_non_http", "enable"),
					resource.TestCheckResourceAttr("fortios_webproxy_global.trname", "unknown_http_version", "best-effort"),
				),
			},
		},
	})
}

func testAccCheckFortiOSWebProxyGlobalExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found WebProxyGlobal: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No WebProxyGlobal is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadWebProxyGlobal(i)

		if err != nil {
			return fmt.Errorf("Error reading WebProxyGlobal: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating WebProxyGlobal: %s", n)
		}

		return nil
	}
}

func testAccCheckWebProxyGlobalDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_webproxy_global" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadWebProxyGlobal(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error WebProxyGlobal %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSWebProxyGlobalConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_webproxy_global" "trname" {
  fast_policy_match               = "enable"
  forward_proxy_auth              = "disable"
  forward_server_affinity_timeout = 30
  learn_client_ip                 = "disable"
  max_message_length              = 32
  max_request_length              = 4
  max_waf_body_cache_length       = 32
  proxy_fqdn                      = "default.fqdn"
  ssl_ca_cert                     = "Fortinet_CA_SSL"
  ssl_cert                        = "Fortinet_Factory"
  strict_web_check                = "disable"
  tunnel_non_http                 = "enable"
  unknown_http_version            = "best-effort"
}
`)
}
