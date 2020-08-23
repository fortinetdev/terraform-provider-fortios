
// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

package fortios
import (
    "fmt"
	"log"
    "testing"
    "github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
    "github.com/hashicorp/terraform-plugin-sdk/helper/resource"
    "github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccFortiOSUserDomainController_basic(t *testing.T) {
    rname := acctest.RandString(8)
    var0 := "var0" + rname
    log.Printf(var0)
    log.Printf("TestAccFortiOSUserDomainController_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSUserDomainControllerConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSUserDomainControllerExists("fortios_user_domaincontroller.trname"),
                    resource.TestCheckResourceAttr("fortios_user_domaincontroller.trname", "domain_name", "s.com"),
                    resource.TestCheckResourceAttr("fortios_user_domaincontroller.trname", "ip_address", "1.1.1.1"),
                    resource.TestCheckResourceAttr("fortios_user_domaincontroller.trname", "ldap_server", var0),
                    resource.TestCheckResourceAttr("fortios_user_domaincontroller.trname", "name", rname),
                    resource.TestCheckResourceAttr("fortios_user_domaincontroller.trname", "port", "445"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSUserDomainControllerExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found UserDomainController: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No UserDomainController is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadUserDomainController(i)

		if err != nil {
			return fmt.Errorf("Error reading UserDomainController: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating UserDomainController: %s", n)
		}

		return nil
	}
}

func testAccCheckUserDomainControllerDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_user_domaincontroller" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadUserDomainController(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error UserDomainController %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSUserDomainControllerConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_user_ldap" "trname1" {
  account_key_filter      = "(&(userPrincipalName=%s)(!(UserAccountControl:1.2.840.113556.1.4.803:=2)))"
  account_key_processing  = "same"
  cnid                    = "cn"
  dn                      = "EIWNCIEW"
  group_member_check      = "user-attr"
  group_object_filter     = "(&(objectcategory=group)(member=*))"
  member_attr             = "memberOf"
  name                    = "var0%[1]s"
  password_expiry_warning = "disable"
  password_renewal        = "disable"
  port                    = 389
  secure                  = "disable"
  server                  = "1.1.1.1"
  server_identity_check   = "disable"
  source_ip               = "0.0.0.0"
  ssl_min_proto_version   = "default"
  type                    = "simple"
}

resource "fortios_user_domaincontroller" "trname" {
  domain_name = "s.com"
  ip_address  = "1.1.1.1"
  ldap_server = fortios_user_ldap.trname1.name
  name        = "%[1]s"
  port        = 445
}
`, name)
}
