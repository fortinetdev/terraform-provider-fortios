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

func TestAccFortiOSSystemNetflow_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemNetflow_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemNetflowConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemNetflowExists("fortios_system_netflow.trname"),
					resource.TestCheckResourceAttr("fortios_system_netflow.trname", "active_flow_timeout", "30"),
					resource.TestCheckResourceAttr("fortios_system_netflow.trname", "collector_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_system_netflow.trname", "collector_port", "2055"),
					resource.TestCheckResourceAttr("fortios_system_netflow.trname", "inactive_flow_timeout", "15"),
					resource.TestCheckResourceAttr("fortios_system_netflow.trname", "source_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_system_netflow.trname", "template_tx_counter", "20"),
					resource.TestCheckResourceAttr("fortios_system_netflow.trname", "template_tx_timeout", "30"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemNetflowExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemNetflow: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemNetflow is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemNetflow(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemNetflow: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemNetflow: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemNetflowDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_netflow" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemNetflow(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemNetflow %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemNetflowConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_netflow" "trname" {
  active_flow_timeout   = 30
  collector_ip          = "0.0.0.0"
  collector_port        = 2055
  inactive_flow_timeout = 15
  source_ip             = "0.0.0.0"
  template_tx_counter   = 20
  template_tx_timeout   = 30
}
`)
}
