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

func TestAccFortiOSSystemAutomationAction_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemAutomationAction_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemAutomationActionConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemAutomationActionExists("fortios_system_automationaction.trname"),
					resource.TestCheckResourceAttr("fortios_system_automationaction.trname", "action_type", "email"),
					resource.TestCheckResourceAttr("fortios_system_automationaction.trname", "aws_domain", "amazonaws.com"),
					resource.TestCheckResourceAttr("fortios_system_automationaction.trname", "delay", "0"),
					resource.TestCheckResourceAttr("fortios_system_automationaction.trname", "email_subject", "SUBJECT1"),
					resource.TestCheckResourceAttr("fortios_system_automationaction.trname", "method", "post"),
					resource.TestCheckResourceAttr("fortios_system_automationaction.trname", "minimum_interval", "1"),
					resource.TestCheckResourceAttr("fortios_system_automationaction.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_system_automationaction.trname", "protocol", "http"),
					resource.TestCheckResourceAttr("fortios_system_automationaction.trname", "required", "disable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemAutomationActionExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemAutomationAction: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemAutomationAction is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemAutomationAction(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemAutomationAction: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemAutomationAction: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemAutomationActionDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_automationaction" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemAutomationAction(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemAutomationAction %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemAutomationActionConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_automationaction" "trname" {
  action_type      = "email"
  aws_domain       = "amazonaws.com"
  delay            = 0
  email_subject    = "SUBJECT1"
  method           = "post"
  minimum_interval = 1
  name             = "%[1]s"
  protocol         = "http"
  required         = "disable"
}
`, name)
}
