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

func TestAccFortiOSUserKrbKeytab_basic(t *testing.T) {
	rname := acctest.RandString(8)
	var0 := "var0" + rname
	log.Printf(var0)
	log.Printf("TestAccFortiOSUserKrbKeytab_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSUserKrbKeytabConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSUserKrbKeytabExists("fortios_user_krbkeytab.trname"),
					resource.TestCheckResourceAttr("fortios_user_krbkeytab.trname", "keytab", "ZXdlY2VxcmVxd3Jld3E="),
					resource.TestCheckResourceAttr("fortios_user_krbkeytab.trname", "ldap_server", var0),
					resource.TestCheckResourceAttr("fortios_user_krbkeytab.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_user_krbkeytab.trname", "principal", "testprin"),
				),
			},
		},
	})
}

func testAccCheckFortiOSUserKrbKeytabExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found UserKrbKeytab: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No UserKrbKeytab is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadUserKrbKeytab(i)

		if err != nil {
			return fmt.Errorf("Error reading UserKrbKeytab: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating UserKrbKeytab: %s", n)
		}

		return nil
	}
}

func testAccCheckUserKrbKeytabDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_user_krbkeytab" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadUserKrbKeytab(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error UserKrbKeytab %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSUserKrbKeytabConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_user_ldap" "trname2" {
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

resource "fortios_user_krbkeytab" "trname" {
  keytab      = "ZXdlY2VxcmVxd3Jld3E="
  ldap_server = fortios_user_ldap.trname2.name
  name        = "%[1]s"
  principal   = "testprin"
}
`, name)
}
