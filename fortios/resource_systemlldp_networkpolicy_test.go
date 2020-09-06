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

func TestAccFortiOSSystemLldpNetworkPolicy_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemLldpNetworkPolicy_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemLldpNetworkPolicyConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemLldpNetworkPolicyExists("fortios_systemlldp_networkpolicy.trname"),
					resource.TestCheckResourceAttr("fortios_systemlldp_networkpolicy.trname", "comment", "test"),
					resource.TestCheckResourceAttr("fortios_systemlldp_networkpolicy.trname", "name", rname),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemLldpNetworkPolicyExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemLldpNetworkPolicy: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemLldpNetworkPolicy is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemLldpNetworkPolicy(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemLldpNetworkPolicy: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemLldpNetworkPolicy: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemLldpNetworkPolicyDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_systemlldp_networkpolicy" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemLldpNetworkPolicy(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemLldpNetworkPolicy %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemLldpNetworkPolicyConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_systemlldp_networkpolicy" "trname" {
  comment = "test"
  name    = "%[1]s"
}
`, name)
}
