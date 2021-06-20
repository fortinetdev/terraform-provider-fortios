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

func TestAccFortiOSUserPeer_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSUserPeer_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSUserPeerConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSUserPeerExists("fortios_user_peer.trname1"),
					resource.TestCheckResourceAttr("fortios_user_peer.trname1", "ca", "EC-ACC"),
					resource.TestCheckResourceAttr("fortios_user_peer.trname1", "cn_type", "string"),
					resource.TestCheckResourceAttr("fortios_user_peer.trname1", "ldap_mode", "password"),
					resource.TestCheckResourceAttr("fortios_user_peer.trname1", "mandatory_ca_verify", "enable"),
					resource.TestCheckResourceAttr("fortios_user_peer.trname1", "name", rname),
					resource.TestCheckResourceAttr("fortios_user_peer.trname1", "two_factor", "disable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSUserPeerExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found UserPeer: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No UserPeer is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadUserPeer(i, "root")

		if err != nil {
			return fmt.Errorf("Error reading UserPeer: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating UserPeer: %s", n)
		}

		return nil
	}
}

func testAccCheckUserPeerDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_user_peer" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadUserPeer(i, "root")

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error UserPeer %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSUserPeerConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_user_peer" "trname1" {
  ca                  = "EC-ACC"
  cn_type             = "string"
  ldap_mode           = "password"
  mandatory_ca_verify = "enable"
  name                = "%[1]s"
  two_factor          = "disable"
}
`, name)
}
