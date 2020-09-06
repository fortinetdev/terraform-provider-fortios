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

func TestAccFortiOSWebProxyDebugUrl_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSWebProxyDebugUrl_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSWebProxyDebugUrlConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSWebProxyDebugUrlExists("fortios_webproxy_debugurl.trname"),
					resource.TestCheckResourceAttr("fortios_webproxy_debugurl.trname", "exact", "enable"),
					resource.TestCheckResourceAttr("fortios_webproxy_debugurl.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_webproxy_debugurl.trname", "status", "enable"),
					resource.TestCheckResourceAttr("fortios_webproxy_debugurl.trname", "url_pattern", "/examples/servlet/*Servlet"),
				),
			},
		},
	})
}

func testAccCheckFortiOSWebProxyDebugUrlExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found WebProxyDebugUrl: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No WebProxyDebugUrl is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadWebProxyDebugUrl(i)

		if err != nil {
			return fmt.Errorf("Error reading WebProxyDebugUrl: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating WebProxyDebugUrl: %s", n)
		}

		return nil
	}
}

func testAccCheckWebProxyDebugUrlDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_webproxy_debugurl" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadWebProxyDebugUrl(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error WebProxyDebugUrl %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSWebProxyDebugUrlConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_webproxy_debugurl" "trname" {
  exact       = "enable"
  name        = "%[1]s"
  status      = "enable"
  url_pattern = "/examples/servlet/*Servlet"
}
`, name)
}
