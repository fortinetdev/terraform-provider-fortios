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

func TestAccFortiOSVpnSslWebUserGroupBookmark_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSVpnSslWebUserGroupBookmark_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSVpnSslWebUserGroupBookmarkConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSVpnSslWebUserGroupBookmarkExists("fortios_vpnsslweb_usergroupbookmark.trname"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_usergroupbookmark.trname", "name", "Guest-group"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_usergroupbookmark.trname", "bookmarks.0.apptype", "citrix"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_usergroupbookmark.trname", "bookmarks.0.listening_port", "0"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_usergroupbookmark.trname", "bookmarks.0.name", "b1"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_usergroupbookmark.trname", "bookmarks.0.port", "0"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_usergroupbookmark.trname", "bookmarks.0.remote_port", "0"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_usergroupbookmark.trname", "bookmarks.0.security", "rdp"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_usergroupbookmark.trname", "bookmarks.0.server_layout", "en-us-qwerty"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_usergroupbookmark.trname", "bookmarks.0.show_status_window", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_usergroupbookmark.trname", "bookmarks.0.sso", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_usergroupbookmark.trname", "bookmarks.0.sso_credential", "sslvpn-login"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_usergroupbookmark.trname", "bookmarks.0.sso_credential_sent_once", "disable"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_usergroupbookmark.trname", "bookmarks.0.url", "www.aaa.com"),
				),
			},
		},
	})
}

func testAccCheckFortiOSVpnSslWebUserGroupBookmarkExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found VpnSslWebUserGroupBookmark: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No VpnSslWebUserGroupBookmark is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadVpnSslWebUserGroupBookmark(i)

		if err != nil {
			return fmt.Errorf("Error reading VpnSslWebUserGroupBookmark: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating VpnSslWebUserGroupBookmark: %s", n)
		}

		return nil
	}
}

func testAccCheckVpnSslWebUserGroupBookmarkDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_vpnsslweb_usergroupbookmark" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadVpnSslWebUserGroupBookmark(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error VpnSslWebUserGroupBookmark %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSVpnSslWebUserGroupBookmarkConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_vpnsslweb_usergroupbookmark" "trname" {
  name = "Guest-group"

  bookmarks {
    apptype                  = "citrix"
    listening_port           = 0
    name                     = "b1"
    port                     = 0
    remote_port              = 0
    security                 = "rdp"
    server_layout            = "en-us-qwerty"
    show_status_window       = "disable"
    sso                      = "disable"
    sso_credential           = "sslvpn-login"
    sso_credential_sent_once = "disable"
    url                      = "www.aaa.com"
  }
}
`)
}
