package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFortiManagerSystemAdom(t *testing.T) {
	name := "fmg-sys-adom" + acctest.RandString(12)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckFortiManager(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFMGSystemAdomDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiManagerSystemAdomConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiManagerSystemAdomExists("fortios_fmg_system_adom.test1"),
					resource.TestCheckResourceAttr("fortios_fmg_system_adom.test1", "name", name),
					resource.TestCheckResourceAttr("fortios_fmg_system_adom.test1", "type", "FortiCarrier"),
					resource.TestCheckResourceAttr("fortios_fmg_system_adom.test1", "central_management_vpn", "false"),
					resource.TestCheckResourceAttr("fortios_fmg_system_adom.test1", "central_management_fortiap", "true"),
					resource.TestCheckResourceAttr("fortios_fmg_system_adom.test1", "central_management_sdwan", "false"),
					resource.TestCheckResourceAttr("fortios_fmg_system_adom.test1", "perform_policy_check_before_every_install", "true"),
					resource.TestCheckResourceAttr("fortios_fmg_system_adom.test1", "auto_push_policy_packages_when_device_back_online", "Enable"),
					resource.TestCheckResourceAttr("fortios_fmg_system_adom.test1", "status", "1"),
				),
			},
		},
	})
}

func testAccCheckFortiManagerSystemAdomExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found System Adom: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No System Adom is set")
		}

		c := testAccProvider.Meta().(*FortiClient).ClientFortimanager

		i := rs.Primary.ID
		o, err := c.ReadSystemAdom(i)

		if err != nil {
			return fmt.Errorf("Error reading System Adom: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating System Adom: %s", n)
		}

		return nil
	}
}

func testAccCheckFMGSystemAdomDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).ClientFortimanager

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_fmg_system_adom" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemAdom(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error System Adom %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiManagerSystemAdomConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_fmg_system_adom" "test1" {
	name = "%s"
    type = "FortiCarrier"
    central_management_vpn = false
    central_management_fortiap = true
    central_management_sdwan = false
    perform_policy_check_before_every_install = true
    auto_push_policy_packages_when_device_back_online = "Enable"
    status = 1
}
`, name)
}
