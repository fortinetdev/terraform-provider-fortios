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

func TestAccFortiOSSwitchControllerTrafficPolicy_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSwitchControllerTrafficPolicy_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSwitchControllerTrafficPolicyConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSwitchControllerTrafficPolicyExists("fortios_switchcontroller_trafficpolicy.trname"),
					resource.TestCheckResourceAttr("fortios_switchcontroller_trafficpolicy.trname", "guaranteed_bandwidth", "0"),
					resource.TestCheckResourceAttr("fortios_switchcontroller_trafficpolicy.trname", "guaranteed_burst", "0"),
					resource.TestCheckResourceAttr("fortios_switchcontroller_trafficpolicy.trname", "maximum_burst", "0"),
					resource.TestCheckResourceAttr("fortios_switchcontroller_trafficpolicy.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_switchcontroller_trafficpolicy.trname", "policer_status", "enable"),
					resource.TestCheckResourceAttr("fortios_switchcontroller_trafficpolicy.trname", "type", "ingress"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSwitchControllerTrafficPolicyExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SwitchControllerTrafficPolicy: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SwitchControllerTrafficPolicy is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSwitchControllerTrafficPolicy(i)

		if err != nil {
			return fmt.Errorf("Error reading SwitchControllerTrafficPolicy: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SwitchControllerTrafficPolicy: %s", n)
		}

		return nil
	}
}

func testAccCheckSwitchControllerTrafficPolicyDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_switchcontroller_trafficpolicy" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSwitchControllerTrafficPolicy(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SwitchControllerTrafficPolicy %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSwitchControllerTrafficPolicyConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_switchcontroller_trafficpolicy" "trname" {
  guaranteed_bandwidth = 0
  guaranteed_burst     = 0
  maximum_burst        = 0
  name                 = "%[1]s"
  policer_status       = "enable"
  type                 = "ingress"
}
`, name)
}
