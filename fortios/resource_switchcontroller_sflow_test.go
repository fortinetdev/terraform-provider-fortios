// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

package fortios

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccFortiOSSwitchControllerSflow_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSwitchControllerSflow_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSwitchControllerSflowConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSwitchControllerSflowExists("fortios_switchcontroller_sflow.trname"),
					resource.TestCheckResourceAttr("fortios_switchcontroller_sflow.trname", "collector_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_switchcontroller_sflow.trname", "collector_port", "6343"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSwitchControllerSflowExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SwitchControllerSflow: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SwitchControllerSflow is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSwitchControllerSflow(i, "root")

		if err != nil {
			return fmt.Errorf("Error reading SwitchControllerSflow: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SwitchControllerSflow: %s", n)
		}

		return nil
	}
}

func testAccCheckSwitchControllerSflowDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_switchcontroller_sflow" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSwitchControllerSflow(i, "root")

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SwitchControllerSflow %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSwitchControllerSflowConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_switchcontroller_sflow" "trname" {
  collector_ip   = "0.0.0.0"
  collector_port = 6343
}
`)
}
