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

func TestAccFortiOSVpnSslWebUserBookmark_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSVpnSslWebUserBookmark_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSVpnSslWebUserBookmarkConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSVpnSslWebUserBookmarkExists("fortios_vpnsslweb_userbookmark.trname"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_userbookmark.trname", "custom_lang", "big5"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_userbookmark.trname", "name", rname),
				),
			},
		},
	})
}

func testAccCheckFortiOSVpnSslWebUserBookmarkExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found VpnSslWebUserBookmark: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No VpnSslWebUserBookmark is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadVpnSslWebUserBookmark(i)

		if err != nil {
			return fmt.Errorf("Error reading VpnSslWebUserBookmark: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating VpnSslWebUserBookmark: %s", n)
		}

		return nil
	}
}

func testAccCheckVpnSslWebUserBookmarkDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_vpnsslweb_userbookmark" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadVpnSslWebUserBookmark(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error VpnSslWebUserBookmark %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSVpnSslWebUserBookmarkConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_vpnsslweb_userbookmark" "trname" {
  custom_lang = "big5"
  name        = "%[1]s"
}
`, name)
}
