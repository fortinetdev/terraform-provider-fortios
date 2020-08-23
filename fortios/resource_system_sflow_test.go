
// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

package fortios
import (
    "fmt"
	"log"
    "testing"
    "github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
    "github.com/hashicorp/terraform-plugin-sdk/helper/resource"
    "github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccFortiOSSystemSflow_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSSystemSflow_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSSystemSflowConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSSystemSflowExists("fortios_system_sflow.trname"),
                    resource.TestCheckResourceAttr("fortios_system_sflow.trname", "collector_ip", "0.0.0.0"),
                    resource.TestCheckResourceAttr("fortios_system_sflow.trname", "collector_port", "6343"),
                    resource.TestCheckResourceAttr("fortios_system_sflow.trname", "source_ip", "0.0.0.0"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSSystemSflowExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemSflow: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemSflow is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemSflow(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemSflow: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemSflow: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemSflowDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_sflow" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemSflow(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemSflow %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemSflowConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_sflow" "trname" {
  collector_ip   = "0.0.0.0"
  collector_port = 6343
  source_ip      = "0.0.0.0"
}
`)
}
