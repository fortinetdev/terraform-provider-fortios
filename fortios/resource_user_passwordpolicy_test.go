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

func TestAccFortiOSUserPasswordPolicy_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSUserPasswordPolicy_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSUserPasswordPolicyConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSUserPasswordPolicyExists("fortios_user_passwordpolicy.trname"),
					resource.TestCheckResourceAttr("fortios_user_passwordpolicy.trname", "expire_days", "22"),
					resource.TestCheckResourceAttr("fortios_user_passwordpolicy.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_user_passwordpolicy.trname", "warn_days", "13"),
				),
			},
		},
	})
}

func testAccCheckFortiOSUserPasswordPolicyExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found UserPasswordPolicy: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No UserPasswordPolicy is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadUserPasswordPolicy(i)

		if err != nil {
			return fmt.Errorf("Error reading UserPasswordPolicy: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating UserPasswordPolicy: %s", n)
		}

		return nil
	}
}

func testAccCheckUserPasswordPolicyDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_user_passwordpolicy" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadUserPasswordPolicy(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error UserPasswordPolicy %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSUserPasswordPolicyConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_user_passwordpolicy" "trname" {
  expire_days = 22
  name        = "%[1]s"
  warn_days   = 13
}
`, name)
}
