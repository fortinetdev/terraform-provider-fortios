package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFortiManagerSystemAdminProfiles(t *testing.T) {
	profileid := "fmg-sys-admin-profiles" + acctest.RandString(12)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckFortiManager(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFMGSystemAdminProfilesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiManagerSystemAdminProfilesConfig(profileid),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiManagerSystemAdminProfilesExists("fortios_fmg_system_admin_profiles.test1"),
					resource.TestCheckResourceAttr("fortios_fmg_system_admin_profiles.test1", "profileid", profileid),
					resource.TestCheckResourceAttr("fortios_fmg_system_admin_profiles.test1", "description", "11"),
					resource.TestCheckResourceAttr("fortios_fmg_system_admin_profiles.test1", "system_setting", "read"),
					resource.TestCheckResourceAttr("fortios_fmg_system_admin_profiles.test1", "intf_mapping", "read-write"),
					resource.TestCheckResourceAttr("fortios_fmg_system_admin_profiles.test1", "device_forticlient", "read"),
					resource.TestCheckResourceAttr("fortios_fmg_system_admin_profiles.test1", "device_ap", "none"),
					resource.TestCheckResourceAttr("fortios_fmg_system_admin_profiles.test1", "fortiguard_center", "read"),
					resource.TestCheckResourceAttr("fortios_fmg_system_admin_profiles.test1", "device_manager", "read-write"),
					resource.TestCheckResourceAttr("fortios_fmg_system_admin_profiles.test1", "device_operation", "read"),
					resource.TestCheckResourceAttr("fortios_fmg_system_admin_profiles.test1", "config_revert", "read"),
					resource.TestCheckResourceAttr("fortios_fmg_system_admin_profiles.test1", "policy_objects", "read-write"),
					resource.TestCheckResourceAttr("fortios_fmg_system_admin_profiles.test1", "assignment", "read"),
					resource.TestCheckResourceAttr("fortios_fmg_system_admin_profiles.test1", "consistency_check", "read-write"),
				),
			},
		},
	})
}

func testAccCheckFortiManagerSystemAdminProfilesExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found System Admin Profiles: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No System Admin Profiles is set")
		}

		c := testAccProvider.Meta().(*FortiClient).ClientFortimanager

		i := rs.Primary.ID
		o, err := c.ReadSystemAdminProfiles(i)

		if err != nil {
			return fmt.Errorf("Error reading System Admin Profiles: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating System Admin Profiles: %s", n)
		}

		return nil
	}
}

func testAccCheckFMGSystemAdminProfilesDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).ClientFortimanager

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_fmg_system_admin_profiles" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemAdminProfiles(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error System Admin Profiles %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiManagerSystemAdminProfilesConfig(id string) string {
	return fmt.Sprintf(`
resource "fortios_fmg_system_admin_profiles" "test1" {
	profileid = "%s"
	description = "11"
    system_setting = "read"
    intf_mapping = "read-write"
    device_forticlient = "read"
    device_ap = "none"
    fortiguard_center = "read"
    device_manager = "read-write"
    device_operation = "read"
    config_revert = "read"
    policy_objects = "read-write"
    assignment = "read"
    consistency_check = "read-write"
}
`, id)
}
