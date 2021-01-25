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

func TestAccFortiOSUserLdap_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSUserLdap_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSUserLdapConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSUserLdapExists("fortios_user_ldap.trname"),
					resource.TestCheckResourceAttr("fortios_user_ldap.trname", "account_key_filter", "(&(userPrincipalName=%s)(!(UserAccountControl:1.2.840.113556.1.4.803:=2)))"),
					resource.TestCheckResourceAttr("fortios_user_ldap.trname", "account_key_processing", "same"),
					resource.TestCheckResourceAttr("fortios_user_ldap.trname", "cnid", "cn"),
					resource.TestCheckResourceAttr("fortios_user_ldap.trname", "dn", "EIWNCIEW"),
					resource.TestCheckResourceAttr("fortios_user_ldap.trname", "group_member_check", "user-attr"),
					resource.TestCheckResourceAttr("fortios_user_ldap.trname", "group_object_filter", "(&(objectcategory=group)(member=*))"),
					resource.TestCheckResourceAttr("fortios_user_ldap.trname", "member_attr", "memberOf"),
					resource.TestCheckResourceAttr("fortios_user_ldap.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_user_ldap.trname", "password_expiry_warning", "disable"),
					resource.TestCheckResourceAttr("fortios_user_ldap.trname", "password_renewal", "disable"),
					resource.TestCheckResourceAttr("fortios_user_ldap.trname", "port", "389"),
					resource.TestCheckResourceAttr("fortios_user_ldap.trname", "secure", "disable"),
					resource.TestCheckResourceAttr("fortios_user_ldap.trname", "server", "1.1.1.1"),
					resource.TestCheckResourceAttr("fortios_user_ldap.trname", "server_identity_check", "disable"),
					resource.TestCheckResourceAttr("fortios_user_ldap.trname", "source_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_user_ldap.trname", "ssl_min_proto_version", "default"),
					resource.TestCheckResourceAttr("fortios_user_ldap.trname", "type", "simple"),
				),
			},
		},
	})
}

func testAccCheckFortiOSUserLdapExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found UserLdap: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No UserLdap is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadUserLdap(i)

		if err != nil {
			return fmt.Errorf("Error reading UserLdap: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating UserLdap: %s", n)
		}

		return nil
	}
}

func testAccCheckUserLdapDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_user_ldap" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadUserLdap(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error UserLdap %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSUserLdapConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_user_ldap" "trname" {
  account_key_filter      = "(&(userPrincipalName=%s)(!(UserAccountControl:1.2.840.113556.1.4.803:=2)))"
  account_key_processing  = "same"
  cnid                    = "cn"
  dn                      = "EIWNCIEW"
  group_member_check      = "user-attr"
  group_object_filter     = "(&(objectcategory=group)(member=*))"
  member_attr             = "memberOf"
  name                    = "%[1]s"
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
`, name)
}
