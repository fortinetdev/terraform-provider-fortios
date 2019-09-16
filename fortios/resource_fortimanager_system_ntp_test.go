package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccFortiManagerSystemNTP(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckFortiManager(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFMGSystemNTPDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiManagerSystemNTPConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiManagerSystemNTPExists("fortios_fortimanager_system_ntp.test1"),
					resource.TestCheckResourceAttr("fortios_fortimanager_system_ntp.test1", "server", "ntp1.fortinet.com"),
					resource.TestCheckResourceAttr("fortios_fortimanager_system_ntp.test1", "status", "enable"),
					resource.TestCheckResourceAttr("fortios_fortimanager_system_ntp.test1", "sync_interval", "30"),
				),
			},
		},
	})
}

func testAccCheckFortiManagerSystemNTPExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found system ntp: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No system ntp is set")
		}

		c := testAccProvider.Meta().(*FortiClient).ClientFortimanager

		o, err := c.ReadSystemNTP()

		if err != nil {
			return fmt.Errorf("Error reading system ntp: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating system ntp: %s", n)
		}

		return nil
	}
}

func testAccCheckFMGSystemNTPDestroy(s *terraform.State) error {
	return nil
}

func testAccFortiManagerSystemNTPConfig() string {
	return fmt.Sprintf(`
resource "fortios_fortimanager_system_ntp" "test1" {
    server = "ntp1.fortinet.com"
    status = "enable"
    sync_interval = 30
}
`)
}
