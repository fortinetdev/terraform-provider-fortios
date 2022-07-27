// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

package fortios

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"log"
	"testing"
)

func TestAccFortiOSSystemSmsServer_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemSmsServer_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemSmsServerConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemSmsServerExists("fortios_system_smsserver.trname"),
					resource.TestCheckResourceAttr("fortios_system_smsserver.trname", "mail_server", "1.1.1.2"),
					resource.TestCheckResourceAttr("fortios_system_smsserver.trname", "name", rname),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemSmsServerExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemSmsServer: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemSmsServer is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemSmsServer(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemSmsServer: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemSmsServer: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemSmsServerDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_smsserver" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemSmsServer(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemSmsServer %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemSmsServerConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_smsserver" "trname" {
  mail_server = "1.1.1.2"
  name        = "%[1]s"
}
`, name)
}
