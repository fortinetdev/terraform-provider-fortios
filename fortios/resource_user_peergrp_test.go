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

func TestAccFortiOSUserPeergrp_basic(t *testing.T) {
	rname := acctest.RandString(8)
	var0 := "var0" + rname
	log.Printf(var0)
	log.Printf("TestAccFortiOSUserPeergrp_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSUserPeergrpConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSUserPeergrpExists("fortios_user_peergrp.trname"),
					resource.TestCheckResourceAttr("fortios_user_peergrp.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_user_peergrp.trname", "member.0.name", var0),
				),
			},
		},
	})
}

func testAccCheckFortiOSUserPeergrpExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found UserPeergrp: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No UserPeergrp is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadUserPeergrp(i)

		if err != nil {
			return fmt.Errorf("Error reading UserPeergrp: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating UserPeergrp: %s", n)
		}

		return nil
	}
}

func testAccCheckUserPeergrpDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_user_peergrp" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadUserPeergrp(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error UserPeergrp %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSUserPeergrpConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_user_peer" "trname2" {
  ca                  = "EC-ACC"
  cn_type             = "string"
  ldap_mode           = "password"
  mandatory_ca_verify = "enable"
  name                = "var0%[1]s"
  two_factor          = "disable"
}

resource "fortios_user_peergrp" "trname" {
  name = "%[1]s"

  member {
    name = fortios_user_peer.trname2.name
  }
}
`, name)
}
