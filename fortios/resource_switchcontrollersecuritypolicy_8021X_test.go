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

func TestAccFortiOSSwitchControllerSecurityPolicy8021X_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSwitchControllerSecurityPolicy8021X_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSwitchControllerSecurityPolicy8021XConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSwitchControllerSecurityPolicy8021XExists("fortios_switchcontrollersecuritypolicy_8021X.trname"),
					resource.TestCheckResourceAttr("fortios_switchcontrollersecuritypolicy_8021X.trname", "auth_fail_vlan", "disable"),
					resource.TestCheckResourceAttr("fortios_switchcontrollersecuritypolicy_8021X.trname", "auth_fail_vlanid", "0"),
					resource.TestCheckResourceAttr("fortios_switchcontrollersecuritypolicy_8021X.trname", "eap_passthru", "disable"),
					resource.TestCheckResourceAttr("fortios_switchcontrollersecuritypolicy_8021X.trname", "framevid_apply", "enable"),
					resource.TestCheckResourceAttr("fortios_switchcontrollersecuritypolicy_8021X.trname", "guest_auth_delay", "30"),
					resource.TestCheckResourceAttr("fortios_switchcontrollersecuritypolicy_8021X.trname", "guest_vlan", "disable"),
					resource.TestCheckResourceAttr("fortios_switchcontrollersecuritypolicy_8021X.trname", "guest_vlanid", "100"),
					resource.TestCheckResourceAttr("fortios_switchcontrollersecuritypolicy_8021X.trname", "mac_auth_bypass", "disable"),
					resource.TestCheckResourceAttr("fortios_switchcontrollersecuritypolicy_8021X.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_switchcontrollersecuritypolicy_8021X.trname", "open_auth", "disable"),
					resource.TestCheckResourceAttr("fortios_switchcontrollersecuritypolicy_8021X.trname", "policy_type", "802.1X"),
					resource.TestCheckResourceAttr("fortios_switchcontrollersecuritypolicy_8021X.trname", "radius_timeout_overwrite", "disable"),
					resource.TestCheckResourceAttr("fortios_switchcontrollersecuritypolicy_8021X.trname", "security_mode", "802.1X"),
					resource.TestCheckResourceAttr("fortios_switchcontrollersecuritypolicy_8021X.trname", "user_group.0.name", "Guest-group"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSwitchControllerSecurityPolicy8021XExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SwitchControllerSecurityPolicy8021X: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SwitchControllerSecurityPolicy8021X is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSwitchControllerSecurityPolicy8021X(i)

		if err != nil {
			return fmt.Errorf("Error reading SwitchControllerSecurityPolicy8021X: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SwitchControllerSecurityPolicy8021X: %s", n)
		}

		return nil
	}
}

func testAccCheckSwitchControllerSecurityPolicy8021XDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_switchcontrollersecuritypolicy_8021X" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSwitchControllerSecurityPolicy8021X(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SwitchControllerSecurityPolicy8021X %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSwitchControllerSecurityPolicy8021XConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_switchcontrollersecuritypolicy_8021X" "trname" {
  auth_fail_vlan           = "disable"
  auth_fail_vlanid         = 0
  eap_passthru             = "disable"
  framevid_apply           = "enable"
  guest_auth_delay         = 30
  guest_vlan               = "disable"
  guest_vlanid             = 100
  mac_auth_bypass          = "disable"
  name                     = "%[1]s"
  open_auth                = "disable"
  policy_type              = "802.1X"
  radius_timeout_overwrite = "disable"
  security_mode            = "802.1X"
  user_group {
    name = "Guest-group"
  }
}
`, name)
}
