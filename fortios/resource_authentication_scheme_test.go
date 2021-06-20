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

func TestAccFortiOSAuthenticationScheme_basic(t *testing.T) {
	rname := acctest.RandString(8)
	var0 := "var0" + rname
	log.Printf(var0)
	log.Printf("TestAccFortiOSAuthenticationScheme_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSAuthenticationSchemeConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSAuthenticationSchemeExists("fortios_authentication_scheme.trname"),
					resource.TestCheckResourceAttr("fortios_authentication_scheme.trname", "fsso_agent_for_ntlm", var0),
					resource.TestCheckResourceAttr("fortios_authentication_scheme.trname", "fsso_guest", "disable"),
					resource.TestCheckResourceAttr("fortios_authentication_scheme.trname", "method", "ntlm"),
					resource.TestCheckResourceAttr("fortios_authentication_scheme.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_authentication_scheme.trname", "negotiate_ntlm", "enable"),
					resource.TestCheckResourceAttr("fortios_authentication_scheme.trname", "require_tfa", "disable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSAuthenticationSchemeExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found AuthenticationScheme: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No AuthenticationScheme is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadAuthenticationScheme(i, "root")

		if err != nil {
			return fmt.Errorf("Error reading AuthenticationScheme: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating AuthenticationScheme: %s", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationSchemeDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_authentication_scheme" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadAuthenticationScheme(i, "root")

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error AuthenticationScheme %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSAuthenticationSchemeConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_user_fsso" "trname3" {
  name       = "var0%[1]s"
  port       = 8000
  port2      = 8000
  port3      = 8000
  port4      = 8000
  port5      = 8000
  server     = "1.1.1.1"
  source_ip  = "0.0.0.0"
  source_ip6 = "::"
}

resource "fortios_authentication_scheme" "trname" {
  fsso_agent_for_ntlm = fortios_user_fsso.trname3.name
  fsso_guest          = "disable"
  method              = "ntlm"
  name                = "%[1]s"
  negotiate_ntlm      = "enable"
  require_tfa         = "disable"
}
`, name)
}
