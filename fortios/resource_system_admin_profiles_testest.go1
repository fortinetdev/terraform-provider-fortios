package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccFortiOSSystemAdminProfiles_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSystemAdminProfilesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemAdminProfilesConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemAdminProfilesExists("fortios_system_admin_profiles.test1"),
					resource.TestCheckResourceAttr("fortios_system_admin_profiles.test1", "name", "s2"),
					resource.TestCheckResourceAttr("fortios_system_admin_profiles.test1", "scope", "vdom"),
					resource.TestCheckResourceAttr("fortios_system_admin_profiles.test1", "comments", "Terraform Test"),
					resource.TestCheckResourceAttr("fortios_system_admin_profiles.test1", "secfabgrp", "none"),
					resource.TestCheckResourceAttr("fortios_system_admin_profiles.test1", "ftviewgrp", "read"),
					resource.TestCheckResourceAttr("fortios_system_admin_profiles.test1", "authgrp", "none"),
					resource.TestCheckResourceAttr("fortios_system_admin_profiles.test1", "sysgrp", "read"),
					resource.TestCheckResourceAttr("fortios_system_admin_profiles.test1", "netgrp", "none"),
					resource.TestCheckResourceAttr("fortios_system_admin_profiles.test1", "loggrp", "none"),
					resource.TestCheckResourceAttr("fortios_system_admin_profiles.test1", "fwgrp", "none"),
					resource.TestCheckResourceAttr("fortios_system_admin_profiles.test1", "vpngrp", "none"),
					resource.TestCheckResourceAttr("fortios_system_admin_profiles.test1", "utmgrp", "none"),
					resource.TestCheckResourceAttr("fortios_system_admin_profiles.test1", "wanoptgrp", "none"),
					resource.TestCheckResourceAttr("fortios_system_admin_profiles.test1", "wifi", "none"),
					resource.TestCheckResourceAttr("fortios_system_admin_profiles.test1", "admintimeout_override", "disable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemAdminProfilesExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found System Admin Profiles: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No System Admin Profiles is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemAdminProfiles(i)

		if err != nil {
			return fmt.Errorf("Error reading System Admin Profiles: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating System Admin Profiles: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemAdminProfilesDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_admin_profiles" {
			continue
		}

		i := rs.Primary.ID
		_, err := c.ReadSystemAdminProfiles(i)

		if err == nil {
			return fmt.Errorf("Error System Admin Profiles %s still exists", rs.Primary.ID)
		}

		return nil
	}

	return nil
}

const testAccFortiOSSystemAdminProfilesConfig = `
resource "fortios_system_admin_profiles" "test1" { 
	name = "s2"
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
`
