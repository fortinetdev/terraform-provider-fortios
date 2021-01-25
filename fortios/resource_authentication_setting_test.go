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

func TestAccFortiOSAuthenticationSetting_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSAuthenticationSetting_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSAuthenticationSettingConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSAuthenticationSettingExists("fortios_authentication_setting.trname"),
					resource.TestCheckResourceAttr("fortios_authentication_setting.trname", "auth_https", "enable"),
					resource.TestCheckResourceAttr("fortios_authentication_setting.trname", "captive_portal_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_authentication_setting.trname", "captive_portal_ip6", "::"),
					resource.TestCheckResourceAttr("fortios_authentication_setting.trname", "captive_portal_port", "7830"),
					resource.TestCheckResourceAttr("fortios_authentication_setting.trname", "captive_portal_ssl_port", "7831"),
					resource.TestCheckResourceAttr("fortios_authentication_setting.trname", "captive_portal_type", "fqdn"),
				),
			},
		},
	})
}

func testAccCheckFortiOSAuthenticationSettingExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found AuthenticationSetting: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No AuthenticationSetting is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadAuthenticationSetting(i)

		if err != nil {
			return fmt.Errorf("Error reading AuthenticationSetting: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating AuthenticationSetting: %s", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationSettingDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_authentication_setting" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadAuthenticationSetting(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error AuthenticationSetting %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSAuthenticationSettingConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_authentication_setting" "trname" {
  auth_https              = "enable"
  captive_portal_ip       = "0.0.0.0"
  captive_portal_ip6      = "::"
  captive_portal_port     = 7830
  captive_portal_ssl_port = 7831
  captive_portal_type     = "fqdn"
}
`)
}
