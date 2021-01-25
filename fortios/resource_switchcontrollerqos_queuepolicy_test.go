// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

package fortios

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"log"
	"testing"
)

func TestAccFortiOSSwitchControllerQosQueuePolicy_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSwitchControllerQosQueuePolicy_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSwitchControllerQosQueuePolicyConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSwitchControllerQosQueuePolicyExists("fortios_switchcontrollerqos_queuepolicy.trname"),
					resource.TestCheckResourceAttr("fortios_switchcontrollerqos_queuepolicy.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_switchcontrollerqos_queuepolicy.trname", "rate_by", "kbps"),
					resource.TestCheckResourceAttr("fortios_switchcontrollerqos_queuepolicy.trname", "schedule", "round-robin"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSwitchControllerQosQueuePolicyExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SwitchControllerQosQueuePolicy: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SwitchControllerQosQueuePolicy is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSwitchControllerQosQueuePolicy(i)

		if err != nil {
			return fmt.Errorf("Error reading SwitchControllerQosQueuePolicy: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SwitchControllerQosQueuePolicy: %s", n)
		}

		return nil
	}
}

func testAccCheckSwitchControllerQosQueuePolicyDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_switchcontrollerqos_queuepolicy" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSwitchControllerQosQueuePolicy(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SwitchControllerQosQueuePolicy %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSwitchControllerQosQueuePolicyConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_switchcontrollerqos_queuepolicy" "trname" {
  name     = "%[1]s"
  rate_by  = "kbps"
  schedule = "round-robin"
}
`, name)
}
