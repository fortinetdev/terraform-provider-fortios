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

func TestAccFortiOSUserRadius_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSUserRadius_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSUserRadiusConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSUserRadiusExists("fortios_user_radius.trname"),
					resource.TestCheckResourceAttr("fortios_user_radius.trname", "acct_all_servers", "disable"),
					resource.TestCheckResourceAttr("fortios_user_radius.trname", "all_usergroup", "disable"),
					resource.TestCheckResourceAttr("fortios_user_radius.trname", "auth_type", "auto"),
					resource.TestCheckResourceAttr("fortios_user_radius.trname", "h3c_compatibility", "disable"),
					resource.TestCheckResourceAttr("fortios_user_radius.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_user_radius.trname", "nas_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_user_radius.trname", "password_encoding", "auto"),
					resource.TestCheckResourceAttr("fortios_user_radius.trname", "password_renewal", "disable"),
					resource.TestCheckResourceAttr("fortios_user_radius.trname", "radius_coa", "disable"),
					resource.TestCheckResourceAttr("fortios_user_radius.trname", "radius_port", "0"),
					resource.TestCheckResourceAttr("fortios_user_radius.trname", "rsso", "disable"),
					resource.TestCheckResourceAttr("fortios_user_radius.trname", "rsso_context_timeout", "28800"),
					resource.TestCheckResourceAttr("fortios_user_radius.trname", "rsso_endpoint_attribute", "Calling-Station-Id"),
					resource.TestCheckResourceAttr("fortios_user_radius.trname", "rsso_ep_one_ip_only", "disable"),
					resource.TestCheckResourceAttr("fortios_user_radius.trname", "rsso_flush_ip_session", "disable"),
					resource.TestCheckResourceAttr("fortios_user_radius.trname", "rsso_log_flags", "protocol-error profile-missing accounting-stop-missed accounting-event endpoint-block radiusd-other"),
					resource.TestCheckResourceAttr("fortios_user_radius.trname", "rsso_log_period", "0"),
					resource.TestCheckResourceAttr("fortios_user_radius.trname", "rsso_radius_response", "disable"),
					resource.TestCheckResourceAttr("fortios_user_radius.trname", "rsso_radius_server_port", "1813"),
					resource.TestCheckResourceAttr("fortios_user_radius.trname", "rsso_validate_request_secret", "disable"),
					resource.TestCheckResourceAttr("fortios_user_radius.trname", "server", "1.1.1.1"),
					resource.TestCheckResourceAttr("fortios_user_radius.trname", "secret", "FDaaewjkeiw32"),
					resource.TestCheckResourceAttr("fortios_user_radius.trname", "sso_attribute", "Class"),
					resource.TestCheckResourceAttr("fortios_user_radius.trname", "sso_attribute_value_override", "enable"),
					resource.TestCheckResourceAttr("fortios_user_radius.trname", "timeout", "5"),
					resource.TestCheckResourceAttr("fortios_user_radius.trname", "use_management_vdom", "disable"),
					resource.TestCheckResourceAttr("fortios_user_radius.trname", "username_case_sensitive", "disable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSUserRadiusExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found UserRadius: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No UserRadius is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadUserRadius(i)

		if err != nil {
			return fmt.Errorf("Error reading UserRadius: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating UserRadius: %s", n)
		}

		return nil
	}
}

func testAccCheckUserRadiusDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_user_radius" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadUserRadius(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error UserRadius %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSUserRadiusConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_user_radius" "trname" {
  acct_all_servers             = "disable"
  all_usergroup                = "disable"
  auth_type                    = "auto"
  h3c_compatibility            = "disable"
  name                         = "%[1]s"
  nas_ip                       = "0.0.0.0"
  password_encoding            = "auto"
  password_renewal             = "disable"
  radius_coa                   = "disable"
  radius_port                  = 0
  rsso                         = "disable"
  rsso_context_timeout         = 28800
  rsso_endpoint_attribute      = "Calling-Station-Id"
  rsso_ep_one_ip_only          = "disable"
  rsso_flush_ip_session        = "disable"
  rsso_log_flags               = "protocol-error profile-missing accounting-stop-missed accounting-event endpoint-block radiusd-other"
  rsso_log_period              = 0
  rsso_radius_response         = "disable"
  rsso_radius_server_port      = 1813
  rsso_validate_request_secret = "disable"
  server                       = "1.1.1.1"
  secret                       = "FDaaewjkeiw32"
  sso_attribute                = "Class"
  sso_attribute_value_override = "enable"
  timeout                      = 5
  use_management_vdom          = "disable"
  username_case_sensitive      = "disable"
}
`, name)
}
