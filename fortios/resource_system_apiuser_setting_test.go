package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccFortiOSSystemAPIUserSetting_basic(t *testing.T) {
	rname := acctest.RandString(12)
	prname := "apsbp" + rname
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSystemAPIUserSettingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemAPIUserSettingConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemAPIUserSettingExists("fortios_system_apiuser_setting.test1"),
					resource.TestCheckResourceAttr("fortios_system_apiuser_setting.test1", "name", rname),
					resource.TestCheckResourceAttr("fortios_system_apiuser_setting.test1", "accprofile", prname),
					resource.TestCheckResourceAttr("fortios_system_apiuser_setting.test1", "comments", "Terraform Test"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemAPIUserSettingExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found System APIUser Setting: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No System APIUser Setting is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemAPIUserSetting(i)

		if err != nil {
			return fmt.Errorf("Error reading System APIUser Setting: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating System APIUser Setting: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemAPIUserSettingDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_apiuser_setting" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemAPIUserSetting(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error System APIUser Setting %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemAPIUserSettingConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_admin_profiles" "test1" {
	name = "apsbp%s"
	scope = "vdom"
	comments = "Terraform Test"
	secfabgrp = "none"
	ftviewgrp = "read"
	authgrp = "none"
	sysgrp = "read"
	netgrp = "none"
	loggrp = "none"
	fwgrp = "none"
	vpngrp = "none"
	utmgrp = "none"
	wanoptgrp = "none"
	wifi = "none"
	admintimeout_override = "disable"
}

resource "fortios_system_apiuser_setting" "test1" {
	name = "%s"
	accprofile = "${fortios_system_admin_profiles.test1.name}"
	vdom = ["root"]
	trusthost {
		type = "ipv4-trusthost"
		ipv4_trusthost = "61.149.0.0 255.255.0.0"
	}

	trusthost {
		type = "ipv4-trusthost"
		ipv4_trusthost = "22.22.0.0 255.255.0.0"
	}
	comments = "Terraform Test"
}
`, name, name)
}
