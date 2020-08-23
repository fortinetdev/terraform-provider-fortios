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

func TestAccFortiOSVpnSslWebRealm_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSVpnSslWebRealm_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSVpnSslWebRealmConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSVpnSslWebRealmExists("fortios_vpnsslweb_realm.trname"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_realm.trname", "login_page", "1.htm"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_realm.trname", "max_concurrent_user", "33"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_realm.trname", "url_path", "1"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_realm.trname", "virtual_host", "2.2.2.2"),
				),
			},
		},
	})
}

func testAccCheckFortiOSVpnSslWebRealmExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found VpnSslWebRealm: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No VpnSslWebRealm is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadVpnSslWebRealm(i)

		if err != nil {
			return fmt.Errorf("Error reading VpnSslWebRealm: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating VpnSslWebRealm: %s", n)
		}

		return nil
	}
}

func testAccCheckVpnSslWebRealmDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_vpnsslweb_realm" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadVpnSslWebRealm(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error VpnSslWebRealm %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSVpnSslWebRealmConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_vpnsslweb_realm" "trname" {
  login_page          = "1.htm"
  max_concurrent_user = 33
  url_path            = "1"
  virtual_host        = "2.2.2.2"
}
`)
}
