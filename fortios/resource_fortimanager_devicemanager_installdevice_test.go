package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
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
					testAccCheckFortiManagerDVMInstallDeviceExists("fortios_fmg_devicemanager_install_device.test1"),
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
resource "fortios_fmg_devicemanager_script" "test1" {
    name = "%s"
    description = "config vlan intf"
    content = "config system interface \n edit test-intf \n\t set vdom root\n\t set interface port1\n\t set type vlan \n\t set vlanid 1 \n\t next \n end"
    target  = "device_database"
}
resource "fortios_fmg_devicemanager_script_execute" "test1" {
    script_name = fortios_fmg_devicemanager_script.test1.name
    target_devname = "myfirewall"
}
resource "fortios_fmg_devicemanager_install_device" "test1" {
	target_devname = fortios_fmg_devicemanager_script_execute.test1.target_devname
}
`, name)
}
