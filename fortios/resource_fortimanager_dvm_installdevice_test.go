package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccFortiManagerDVMInstallDevice(t *testing.T) {
	name := "dvm-script" + acctest.RandString(2)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckFortiManager(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiManagerDVMInstallDeviceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiManagerDVMInstallDeviceExists("fortios_fortimanager_devicemanager_install_device.test1"),
				),
			},
		},
	})
}

func testAccCheckFortiManagerDVMInstallDeviceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found devicemanager install device config file: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No devicemanager install device config is set")
		}

		return nil
	}
}

func testAccFortiManagerDVMInstallDeviceConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_fortimanager_devicemanager_script" "test1" {
    name = "%s"
    description = "configure interface3"
    content = "config system interface \n edit port3 \n\t set vdom \"root\"\n\t set ip 10.11.0.200 255.255.0.0 \n\t set allowaccess ping http https\n\t next \n end"
    target  = "device_database"
}
resource "fortios_fortimanager_devicemanager_script_execute" "test1" {
    script_name = fortios_fortimanager_devicemanager_script.test1.name
    target_devname = "FGVM64-test"
}
resource "fortios_fortimanager_devicemanager_install_device" "test1" {
	target_devname = fortios_fortimanager_devicemanager_script_execute.test1.target_devname
}
`, name)
}
