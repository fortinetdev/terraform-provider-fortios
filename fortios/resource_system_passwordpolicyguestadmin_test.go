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

func TestAccFortiOSSystemPasswordPolicyGuestAdmin_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemPasswordPolicyGuestAdmin_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemPasswordPolicyGuestAdminConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemPasswordPolicyGuestAdminExists("fortios_system_passwordpolicyguestadmin.trname"),
					resource.TestCheckResourceAttr("fortios_system_passwordpolicyguestadmin.trname", "apply_to", "guest-admin-password"),
					resource.TestCheckResourceAttr("fortios_system_passwordpolicyguestadmin.trname", "change_4_characters", "disable"),
					resource.TestCheckResourceAttr("fortios_system_passwordpolicyguestadmin.trname", "expire_day", "90"),
					resource.TestCheckResourceAttr("fortios_system_passwordpolicyguestadmin.trname", "expire_status", "disable"),
					resource.TestCheckResourceAttr("fortios_system_passwordpolicyguestadmin.trname", "min_lower_case_letter", "0"),
					resource.TestCheckResourceAttr("fortios_system_passwordpolicyguestadmin.trname", "min_non_alphanumeric", "0"),
					resource.TestCheckResourceAttr("fortios_system_passwordpolicyguestadmin.trname", "min_number", "0"),
					resource.TestCheckResourceAttr("fortios_system_passwordpolicyguestadmin.trname", "min_upper_case_letter", "0"),
					resource.TestCheckResourceAttr("fortios_system_passwordpolicyguestadmin.trname", "minimum_length", "8"),
					resource.TestCheckResourceAttr("fortios_system_passwordpolicyguestadmin.trname", "reuse_password", "enable"),
					resource.TestCheckResourceAttr("fortios_system_passwordpolicyguestadmin.trname", "status", "disable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemPasswordPolicyGuestAdminExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemPasswordPolicyGuestAdmin: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemPasswordPolicyGuestAdmin is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemPasswordPolicyGuestAdmin(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemPasswordPolicyGuestAdmin: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemPasswordPolicyGuestAdmin: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemPasswordPolicyGuestAdminDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_passwordpolicyguestadmin" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemPasswordPolicyGuestAdmin(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemPasswordPolicyGuestAdmin %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemPasswordPolicyGuestAdminConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_passwordpolicyguestadmin" "trname" {
  apply_to              = "guest-admin-password"
  change_4_characters   = "disable"
  expire_day            = 90
  expire_status         = "disable"
  min_lower_case_letter = 0
  min_non_alphanumeric  = 0
  min_number            = 0
  min_upper_case_letter = 0
  minimum_length        = 8
  reuse_password        = "enable"
  status                = "disable"
}
`)
}
