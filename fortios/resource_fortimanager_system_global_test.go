package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFortiManagerSystemGlobal(t *testing.T) {
	name := "fmg-sys-global" + acctest.RandString(12)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckFortiManager(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiManagerSystemGlobalConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiManagerSystemGlobalExists("fortios_fmg_system_global.test1"),
					resource.TestCheckResourceAttr("fortios_fmg_system_global.test1", "hostname", name),
					resource.TestCheckResourceAttr("fortios_fmg_system_global.test1", "fortianalyzer_status", "disable"),
					resource.TestCheckResourceAttr("fortios_fmg_system_global.test1", "adom_status", "disable"),
					resource.TestCheckResourceAttr("fortios_fmg_system_global.test1", "adom_mode", "advanced"),
					resource.TestCheckResourceAttr("fortios_fmg_system_global.test1", "timezone", "09"),
				),
			},
		},
	})
}

func testAccCheckFortiManagerSystemGlobalExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found system global: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No system global is set")
		}

		c := testAccProvider.Meta().(*FortiClient).ClientFortimanager

		o, err := c.ReadSystemGlobal()

		if err != nil {
			return fmt.Errorf("Error reading system global: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating system global: %s", n)
		}

		return nil
	}
}

func testAccFortiManagerSystemGlobalConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_fmg_system_global" "test1" {
	hostname = "%s"
    fortianalyzer_status = "disable"
    adom_status = "disable"
    adom_mode = "advanced"
    timezone = "09"
}
`, name)
}
