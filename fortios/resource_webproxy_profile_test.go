// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

package fortios

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"log"
	"testing"
)

func TestAccFortiOSWebProxyProfile_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSWebProxyProfile_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSWebProxyProfileConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSWebProxyProfileExists("fortios_webproxy_profile.trname"),
					resource.TestCheckResourceAttr("fortios_webproxy_profile.trname", "header_client_ip", "pass"),
					resource.TestCheckResourceAttr("fortios_webproxy_profile.trname", "header_front_end_https", "pass"),
					resource.TestCheckResourceAttr("fortios_webproxy_profile.trname", "header_via_request", "add"),
					resource.TestCheckResourceAttr("fortios_webproxy_profile.trname", "header_via_response", "pass"),
					resource.TestCheckResourceAttr("fortios_webproxy_profile.trname", "header_x_authenticated_groups", "pass"),
					resource.TestCheckResourceAttr("fortios_webproxy_profile.trname", "header_x_authenticated_user", "pass"),
					resource.TestCheckResourceAttr("fortios_webproxy_profile.trname", "header_x_forwarded_for", "pass"),
					resource.TestCheckResourceAttr("fortios_webproxy_profile.trname", "log_header_change", "disable"),
					resource.TestCheckResourceAttr("fortios_webproxy_profile.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_webproxy_profile.trname", "strip_encoding", "disable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSWebProxyProfileExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found WebProxyProfile: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No WebProxyProfile is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadWebProxyProfile(i)

		if err != nil {
			return fmt.Errorf("Error reading WebProxyProfile: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating WebProxyProfile: %s", n)
		}

		return nil
	}
}

func testAccCheckWebProxyProfileDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_webproxy_profile" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadWebProxyProfile(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error WebProxyProfile %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSWebProxyProfileConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_webproxy_profile" "trname" {
  header_client_ip              = "pass"
  header_front_end_https        = "pass"
  header_via_request            = "add"
  header_via_response           = "pass"
  header_x_authenticated_groups = "pass"
  header_x_authenticated_user   = "pass"
  header_x_forwarded_for        = "pass"
  log_header_change             = "disable"
  name                          = "%[1]s"
  strip_encoding                = "disable"
}
`, name)
}
