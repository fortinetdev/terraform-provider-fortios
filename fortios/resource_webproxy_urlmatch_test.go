
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

func TestAccFortiOSWebProxyUrlMatch_basic(t *testing.T) {
    rname := acctest.RandString(8)
    var0 := "var0" + rname
    log.Printf(var0)
    log.Printf("TestAccFortiOSWebProxyUrlMatch_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSWebProxyUrlMatchConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSWebProxyUrlMatchExists("fortios_webproxy_urlmatch.trname"),
                    resource.TestCheckResourceAttr("fortios_webproxy_urlmatch.trname", "cache_exemption", "disable"),
                    resource.TestCheckResourceAttr("fortios_webproxy_urlmatch.trname", "forward_server", var0),
                    resource.TestCheckResourceAttr("fortios_webproxy_urlmatch.trname", "name", rname),
                    resource.TestCheckResourceAttr("fortios_webproxy_urlmatch.trname", "status", "enable"),
                    resource.TestCheckResourceAttr("fortios_webproxy_urlmatch.trname", "url_pattern", "/examples/servlet/*Servlet"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSWebProxyUrlMatchExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found WebProxyUrlMatch: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No WebProxyUrlMatch is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadWebProxyUrlMatch(i)

		if err != nil {
			return fmt.Errorf("Error reading WebProxyUrlMatch: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating WebProxyUrlMatch: %s", n)
		}

		return nil
	}
}

func testAccCheckWebProxyUrlMatchDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_webproxy_urlmatch" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadWebProxyUrlMatch(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error WebProxyUrlMatch %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSWebProxyUrlMatchConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_webproxy_forwardserver" "trname2" {
  addr_type          = "fqdn"
  healthcheck        = "disable"
  ip                 = "0.0.0.0"
  monitor            = "http://www.google.com"
  name               = "var0%[1]s"
  port               = 3128
  server_down_option = "block"
}

resource "fortios_webproxy_urlmatch" "trname" {
  cache_exemption = "disable"
  forward_server  = fortios_webproxy_forwardserver.trname2.name
  name            = "%[1]s"
  status          = "enable"
  url_pattern     = "/examples/servlet/*Servlet"
}


`, name)
}
