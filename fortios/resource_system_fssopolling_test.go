// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

package fortios

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"log"
	"testing"
)

func TestAccFortiOSSystemFssoPolling_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemFssoPolling_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemFssoPollingConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemFssoPollingExists("fortios_system_fssopolling.trname"),
					resource.TestCheckResourceAttr("fortios_system_fssopolling.trname", "authentication", "disable"),
					resource.TestCheckResourceAttr("fortios_system_fssopolling.trname", "listening_port", "8000"),
					resource.TestCheckResourceAttr("fortios_system_fssopolling.trname", "status", "enable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemFssoPollingExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemFssoPolling: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemFssoPolling is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemFssoPolling(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemFssoPolling: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemFssoPolling: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemFssoPollingDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_fssopolling" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemFssoPolling(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemFssoPolling %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemFssoPollingConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_fssopolling" "trname" {
  authentication = "disable"
  listening_port = 8000
  status         = "enable"
}
`)
}
