package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccFortiOSSystemPasswordPolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		// CheckDestroy: testAccCheckSystemPasswordPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemPasswordPolicyConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemPasswordPolicyExists("fortios_system_password_policy.test1"),
					resource.TestCheckResourceAttr("fortios_system_password_policy.test1", "status", "enable"),
					resource.TestCheckResourceAttr("fortios_system_password_policy.test1", "apply_to", "admin-password"),
					resource.TestCheckResourceAttr("fortios_system_password_policy.test1", "minimum_length", "12"),
					resource.TestCheckResourceAttr("fortios_system_password_policy.test1", "min_lower_case_letter", "1"),
					resource.TestCheckResourceAttr("fortios_system_password_policy.test1", "min_upper_case_letter", "1"),
					resource.TestCheckResourceAttr("fortios_system_password_policy.test1", "min_non_alphanumeric", "1"),
					resource.TestCheckResourceAttr("fortios_system_password_policy.test1", "min_number", "1"),
					resource.TestCheckResourceAttr("fortios_system_password_policy.test1", "change_4_characters", "enable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemPasswordPolicyExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found System Password Policy: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No System Password Policy is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemPasswordPolicy(i)

		if err != nil {
			return fmt.Errorf("Error reading System Password Policy: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemPassword Policy: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemPasswordPolicyDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_password_policy" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemPasswordPolicy(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemPassword Policy %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

const testAccFortiOSSystemPasswordPolicyConfig = `
resource "fortios_system_password_policy" "test1" { 
    status = "enable"
    apply_to  = "admin-password"
    minimum_length = "12"
    min_lower_case_letter  = "1"
    min_upper_case_letter  = "1"
    min_non_alphanumeric  = "1"
    min_number  = "1"
    change_4_characters = "enable"
}
`
