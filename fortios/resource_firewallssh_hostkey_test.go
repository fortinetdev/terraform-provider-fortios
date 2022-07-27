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

func TestAccFortiOSFirewallSshHostKey_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSFirewallSshHostKey_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallSshHostKeyConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallSshHostKeyExists("fortios_firewallssh_hostkey.trname"),
					resource.TestCheckResourceAttr("fortios_firewallssh_hostkey.trname", "hostname", "testmachine"),
					resource.TestCheckResourceAttr("fortios_firewallssh_hostkey.trname", "ip", "1.1.1.1"),
					resource.TestCheckResourceAttr("fortios_firewallssh_hostkey.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewallssh_hostkey.trname", "port", "22"),
					resource.TestCheckResourceAttr("fortios_firewallssh_hostkey.trname", "status", "trusted"),
					resource.TestCheckResourceAttr("fortios_firewallssh_hostkey.trname", "type", "RSA"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallSshHostKeyExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallSshHostKey: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallSshHostKey is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallSshHostKey(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallSshHostKey: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallSshHostKey: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallSshHostKeyDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewallssh_hostkey" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallSshHostKey(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallSshHostKey %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallSshHostKeyConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewallssh_hostkey" "trname" {
  hostname = "testmachine"
  ip       = "1.1.1.1"
  name     = "%[1]s"
  port     = 22
  status   = "trusted"
  type     = "RSA"
}
`, name)
}
