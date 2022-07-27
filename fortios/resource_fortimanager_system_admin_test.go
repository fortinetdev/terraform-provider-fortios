package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFortiManagerSystemAdmin(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckFortiManager(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFMGSystemAdminDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiManagerSystemAdminConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiManagerSystemAdminExists("fortios_fmg_system_admin.test1"),
					resource.TestCheckResourceAttr("fortios_fmg_system_admin.test1", "http_port", "80"),
					resource.TestCheckResourceAttr("fortios_fmg_system_admin.test1", "idle_timeout", "20"),
				),
			},
		},
	})
}

func testAccCheckFortiManagerSystemAdminExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found system admin: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No system admin is set")
		}

		c := testAccProvider.Meta().(*FortiClient).ClientFortimanager

		o, err := c.ReadSystemAdmin()

		if err != nil {
			return fmt.Errorf("Error reading system admin: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating system admin: %s", n)
		}

		return nil
	}
}

func testAccCheckFMGSystemAdminDestroy(s *terraform.State) error {
	return nil
}

func testAccFortiManagerSystemAdminConfig() string {
	return fmt.Sprintf(`
resource "fortios_fmg_system_admin" "test1" {
    http_port = 80
    idle_timeout = 20
}
`)
}
