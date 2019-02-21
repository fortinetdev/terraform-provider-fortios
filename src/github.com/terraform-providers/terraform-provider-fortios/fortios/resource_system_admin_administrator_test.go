package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccFortiOSSystemAdminAdministrator_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSystemAdminAdministratorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemAdminAdministratorConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemAdminAdministratorExists("fortios_system_admin_administrator.test1"),
					resource.TestCheckResourceAttr("fortios_system_admin_administrator.test1", "name", "s1"),
					resource.TestCheckResourceAttr("fortios_system_admin_administrator.test1", "trusthost1", "1.1.1.0 255.255.255.0"),
					resource.TestCheckResourceAttr("fortios_system_admin_administrator.test1", "trusthost2", "2.2.2.0 255.255.255.0"),
					resource.TestCheckResourceAttr("fortios_system_admin_administrator.test1", "accprofile", "accprofilefortest"),
					resource.TestCheckResourceAttr("fortios_system_admin_administrator.test1", "comments", "Terraform Test"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemAdminAdministratorExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found System Admin Administrator: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No System Admin Administrator is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemAdminAdministrator(i)

		if err != nil {
			return fmt.Errorf("Error reading System Admin Administrator: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating System Admin Administrator: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemAdminAdministratorDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_admin_administrator" {
			continue
		}

		i := rs.Primary.ID
		_, err := c.ReadSystemAdminAdministrator(i)

		if err == nil {
			return fmt.Errorf("Error System Admin Administrator %s still exists", rs.Primary.ID)
		}

		return nil
	}

	return nil
}

const testAccFortiOSSystemAdminAdministratorConfig = `
resource "fortios_system_admin_administrator" "test1" { 
	name = "s1"
	password = "cc37$1AC1"
	trusthost1 = "1.1.1.0 255.255.255.0"
	trusthost2 = "2.2.2.0 255.255.255.0"
	accprofile = "accprofilefortest"
	comments = "Terraform Test"
}
`
