package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccFortiManagerSystemAdminUser(t *testing.T) {
	userid := "fmg-sys-admin-user" + acctest.RandString(12)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckFortiManager(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFMGSystemAdminUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiManagerSystemAdminUserConfig(userid),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiManagerSystemAdminUserExists("fortios_fmg_system_admin_user.test1"),
					resource.TestCheckResourceAttr("fortios_fmg_system_admin_user.test1", "userid", userid),
					resource.TestCheckResourceAttr("fortios_fmg_system_admin_user.test1", "description", "local user"),
					resource.TestCheckResourceAttr("fortios_fmg_system_admin_user.test1", "user_type", "local"),
					resource.TestCheckResourceAttr("fortios_fmg_system_admin_user.test1", "rpc_permit", "read"),
					resource.TestCheckResourceAttr("fortios_fmg_system_admin_user.test1", "profileid", "Standard_User"),
					resource.TestCheckResourceAttr("fortios_fmg_system_admin_user.test1", "trusthost1", "1.1.1.1 255.255.255.255"),
				),
			},
		},
	})
}

func testAccCheckFortiManagerSystemAdminUserExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found System Admin User: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No System Admin User is set")
		}

		c := testAccProvider.Meta().(*FortiClient).ClientFortimanager

		i := rs.Primary.ID
		o, err := c.ReadSystemAdminUser(i)

		if err != nil {
			return fmt.Errorf("Error reading System Admin User: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating System Admin User: %s", n)
		}

		return nil
	}
}

func testAccCheckFMGSystemAdminUserDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).ClientFortimanager

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_fmg_system_admin_user" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemAdminUser(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error System Admin User %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiManagerSystemAdminUserConfig(id string) string {
	return fmt.Sprintf(`
resource "fortios_fmg_system_admin_user" "test1" {
	userid = "%s"
	password = "123"
    description = "local user"
    user_type = "local"
    rpc_permit = "read"
    profileid = "Standard_User"
    trusthost1 = "1.1.1.1 255.255.255.255"
}
`, id)
}
