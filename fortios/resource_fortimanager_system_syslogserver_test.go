package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccFortiManagerSystemSyslogServer(t *testing.T) {
	name := "fmg-sys-syslogsrv" + acctest.RandString(12)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckFortiManager(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFMGSystemSyslogServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiManagerSystemSyslogServerConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiManagerSystemSyslogServerExists("fortios_fortimanager_system_syslogserver.test1"),
					resource.TestCheckResourceAttr("fortios_fortimanager_system_syslogserver.test1", "name", name),
					resource.TestCheckResourceAttr("fortios_fortimanager_system_syslogserver.test1", "ip", "1.1.1.2"),
					resource.TestCheckResourceAttr("fortios_fortimanager_system_syslogserver.test1", "port", "99"),
				),
			},
		},
	})
}

func testAccCheckFortiManagerSystemSyslogServerExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found system syslogserver: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No system syslogserver is set")
		}

		c := testAccProvider.Meta().(*FortiClient).ClientFortimanager

		i := rs.Primary.ID
		o, err := c.ReadSystemSyslogServer(i)

		if err != nil {
			return fmt.Errorf("Error reading system syslogserver: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating system syslogserver: %s", n)
		}

		return nil
	}
}

func testAccCheckFMGSystemSyslogServerDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).ClientFortimanager

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_fortimanager_system_syslogserver" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemSyslogServer(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error system syslogserver %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiManagerSystemSyslogServerConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_fortimanager_system_syslogserver" "test1" {
	name = "%s"
	ip = "1.1.1.2"
	port = 99
}
`, name)
}
