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

func TestAccFortiOSSwitchControllerLldpProfile_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSwitchControllerLldpProfile_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSwitchControllerLldpProfileConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSwitchControllerLldpProfileExists("fortios_switchcontroller_lldpprofile.trname"),
					resource.TestCheckResourceAttr("fortios_switchcontroller_lldpprofile.trname", "auto_isl", "enable"),
					resource.TestCheckResourceAttr("fortios_switchcontroller_lldpprofile.trname", "auto_isl_hello_timer", "3"),
					resource.TestCheckResourceAttr("fortios_switchcontroller_lldpprofile.trname", "auto_isl_port_group", "0"),
					resource.TestCheckResourceAttr("fortios_switchcontroller_lldpprofile.trname", "auto_isl_receive_timeout", "60"),
					resource.TestCheckResourceAttr("fortios_switchcontroller_lldpprofile.trname", "med_tlvs", "inventory-management network-policy"),
					resource.TestCheckResourceAttr("fortios_switchcontroller_lldpprofile.trname", "name", rname),
				),
			},
		},
	})
}

func testAccCheckFortiOSSwitchControllerLldpProfileExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SwitchControllerLldpProfile: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SwitchControllerLldpProfile is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSwitchControllerLldpProfile(i, "root")

		if err != nil {
			return fmt.Errorf("Error reading SwitchControllerLldpProfile: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SwitchControllerLldpProfile: %s", n)
		}

		return nil
	}
}

func testAccCheckSwitchControllerLldpProfileDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_switchcontroller_lldpprofile" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSwitchControllerLldpProfile(i, "root")

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SwitchControllerLldpProfile %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSwitchControllerLldpProfileConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_switchcontroller_lldpprofile" "trname" {
  auto_isl                 = "enable"
  auto_isl_hello_timer     = 3
  auto_isl_port_group      = 0
  auto_isl_receive_timeout = 60
  med_tlvs                 = "inventory-management network-policy"
  name                     = "%[1]s"
}
`, name)
}
