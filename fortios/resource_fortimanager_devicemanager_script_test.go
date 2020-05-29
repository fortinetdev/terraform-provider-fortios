package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccFortiManagerDVMScript(t *testing.T) {
	name := "fmg-dvm-script" + acctest.RandString(12)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckFortiManager(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFMGDVMScriptDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiManagerDVMScriptConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiManagerDVMScriptExists("fortios_fmg_devicemanager_script.test1"),
					resource.TestCheckResourceAttr("fortios_fmg_devicemanager_script.test1", "name", name),
					resource.TestCheckResourceAttr("fortios_fmg_devicemanager_script.test1", "description", "description"),
					resource.TestCheckResourceAttr("fortios_fmg_devicemanager_script.test1", "content", "config system interface \n edit port3 \n\t set vdom \"root\"\n\t set ip 10.7.0.200 255.255.0.0 \n\t set allowaccess ping http https\n\t next \n end"),
					resource.TestCheckResourceAttr("fortios_fmg_devicemanager_script.test1", "target", "remote_device"),
				),
			},
		},
	})
}

func testAccCheckFortiManagerDVMScriptExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found devicemanager script: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No devicemanager script is set")
		}

		c := testAccProvider.Meta().(*FortiClient).ClientFortimanager

		i := rs.Primary.ID
		o, err := c.ReadDVMScript("root", i)

		if err != nil {
			return fmt.Errorf("Error reading Devicemanager Script: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating Devicemanager Script: %s", n)
		}

		return nil
	}
}

func testAccCheckFMGDVMScriptDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).ClientFortimanager

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_fmg_devicemanager_script" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadDVMScript("root", i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error Devicemanager Script %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiManagerDVMScriptConfig(id string) string {
	return fmt.Sprintf(`
resource "fortios_fmg_devicemanager_script" "test1" {
	name = "%s"
    description = "description"
    content = "config system interface \n edit port3 \n\t set vdom \"root\"\n\t set ip 10.7.0.200 255.255.0.0 \n\t set allowaccess ping http https\n\t next \n end"
    target  = "remote_device"
}
`, id)
}
