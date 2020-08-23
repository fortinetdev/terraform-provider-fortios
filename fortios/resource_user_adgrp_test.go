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

func TestAccFortiOSUserAdgrp_basic(t *testing.T) {
	rname := acctest.RandString(8)
	var0 := "var0" + rname
	log.Printf(var0)
	log.Printf("TestAccFortiOSUserAdgrp_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSUserAdgrpConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSUserAdgrpExists("fortios_user_adgrp.trname"),
					resource.TestCheckResourceAttr("fortios_user_adgrp.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_user_adgrp.trname", "server_name", var0),
				),
			},
		},
	})
}

func testAccCheckFortiOSUserAdgrpExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found UserAdgrp: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No UserAdgrp is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadUserAdgrp(i)

		if err != nil {
			return fmt.Errorf("Error reading UserAdgrp: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating UserAdgrp: %s", n)
		}

		return nil
	}
}

func testAccCheckUserAdgrpDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_user_adgrp" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadUserAdgrp(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error UserAdgrp %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSUserAdgrpConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_user_fsso" "trname1" {
  name       = "var0%[1]s"
  port       = 32381
  port2      = 8000
  port3      = 8000
  port4      = 8000
  port5      = 8000
  server     = "1.1.1.1"
  source_ip  = "0.0.0.0"
  source_ip6 = "::"
}

resource "fortios_user_adgrp" "trname" {
  name        = "%[1]s"
  server_name = fortios_user_fsso.trname1.name
}
`, name)
}
