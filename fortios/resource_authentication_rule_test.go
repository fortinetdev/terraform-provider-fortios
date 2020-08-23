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

func TestAccFortiOSAuthenticationRule_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSAuthenticationRule_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSAuthenticationRuleConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSAuthenticationRuleExists("fortios_authentication_rule.trname"),
					resource.TestCheckResourceAttr("fortios_authentication_rule.trname", "ip_based", "enable"),
					resource.TestCheckResourceAttr("fortios_authentication_rule.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_authentication_rule.trname", "protocol", "ftp"),
					resource.TestCheckResourceAttr("fortios_authentication_rule.trname", "status", "enable"),
					resource.TestCheckResourceAttr("fortios_authentication_rule.trname", "transaction_based", "disable"),
					resource.TestCheckResourceAttr("fortios_authentication_rule.trname", "web_auth_cookie", "disable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSAuthenticationRuleExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found AuthenticationRule: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No AuthenticationRule is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadAuthenticationRule(i)

		if err != nil {
			return fmt.Errorf("Error reading AuthenticationRule: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating AuthenticationRule: %s", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationRuleDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_authentication_rule" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadAuthenticationRule(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error AuthenticationRule %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSAuthenticationRuleConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_authentication_rule" "trname" {
  ip_based          = "enable"
  name              = "%[1]s"
  protocol          = "ftp"
  status            = "enable"
  transaction_based = "disable"
  web_auth_cookie   = "disable"
}
`, name)
}
