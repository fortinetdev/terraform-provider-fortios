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

func TestAccFortiOSSystemPasswordPolicy_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemPasswordPolicy_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemPasswordPolicyConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemPasswordPolicyExists("fortios_system_passwordpolicy.trname"),
					resource.TestCheckResourceAttr("fortios_system_passwordpolicy.trname", "apply_to", "admin-password"),
					resource.TestCheckResourceAttr("fortios_system_passwordpolicy.trname", "change_4_characters", "disable"),
					resource.TestCheckResourceAttr("fortios_system_passwordpolicy.trname", "expire_day", "90"),
					resource.TestCheckResourceAttr("fortios_system_passwordpolicy.trname", "expire_status", "disable"),
					resource.TestCheckResourceAttr("fortios_system_passwordpolicy.trname", "min_lower_case_letter", "0"),
					resource.TestCheckResourceAttr("fortios_system_passwordpolicy.trname", "min_non_alphanumeric", "0"),
					resource.TestCheckResourceAttr("fortios_system_passwordpolicy.trname", "min_number", "0"),
					resource.TestCheckResourceAttr("fortios_system_passwordpolicy.trname", "min_upper_case_letter", "0"),
					resource.TestCheckResourceAttr("fortios_system_passwordpolicy.trname", "minimum_length", "8"),
					resource.TestCheckResourceAttr("fortios_system_passwordpolicy.trname", "reuse_password", "enable"),
					resource.TestCheckResourceAttr("fortios_system_passwordpolicy.trname", "status", "disable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemPasswordPolicyExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemPasswordPolicy: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemPasswordPolicy is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemPasswordPolicy(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemPasswordPolicy: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemPasswordPolicy: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemPasswordPolicyDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_passwordpolicy" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemPasswordPolicy(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemPasswordPolicy %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemPasswordPolicyConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_passwordpolicy" "trname" {
  apply_to              = "admin-password"
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
