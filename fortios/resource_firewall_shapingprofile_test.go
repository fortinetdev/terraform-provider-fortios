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

func TestAccFortiOSFirewallShapingProfile_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSFirewallShapingProfile_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallShapingProfileConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallShapingProfileExists("fortios_firewall_shapingprofile.trname"),
					resource.TestCheckResourceAttr("fortios_firewall_shapingprofile.trname", "default_class_id", "2"),
					resource.TestCheckResourceAttr("fortios_firewall_shapingprofile.trname", "profile_name", "shapingprofiles1"),
					resource.TestCheckResourceAttr("fortios_firewall_shapingprofile.trname", "shaping_entries.0.class_id", "2"),
					resource.TestCheckResourceAttr("fortios_firewall_shapingprofile.trname", "shaping_entries.0.guaranteed_bandwidth_percentage", "33"),
					resource.TestCheckResourceAttr("fortios_firewall_shapingprofile.trname", "shaping_entries.0.id", "1"),
					resource.TestCheckResourceAttr("fortios_firewall_shapingprofile.trname", "shaping_entries.0.maximum_bandwidth_percentage", "88"),
					resource.TestCheckResourceAttr("fortios_firewall_shapingprofile.trname", "shaping_entries.0.priority", "high"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallShapingProfileExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallShapingProfile: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallShapingProfile is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallShapingProfile(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallShapingProfile: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallShapingProfile: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallShapingProfileDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_shapingprofile" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallShapingProfile(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallShapingProfile %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallShapingProfileConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_shapingprofile" "trname" {
  default_class_id = 2
  profile_name     = "shapingprofiles1"

  shaping_entries {
    class_id                        = 2
    guaranteed_bandwidth_percentage = 33
    id                              = 1
    maximum_bandwidth_percentage    = 88
    priority                        = "high"
  }
}
`)
}
