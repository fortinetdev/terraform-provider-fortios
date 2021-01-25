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

func TestAccFortiOSFirewallIpTranslation_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSFirewallIpTranslation_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallIpTranslationConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallIpTranslationExists("fortios_firewall_iptranslation.trname"),
					resource.TestCheckResourceAttr("fortios_firewall_iptranslation.trname", "endip", "2.2.2.2"),
					resource.TestCheckResourceAttr("fortios_firewall_iptranslation.trname", "map_startip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_firewall_iptranslation.trname", "startip", "1.1.1.1"),
					resource.TestCheckResourceAttr("fortios_firewall_iptranslation.trname", "transid", "1"),
					resource.TestCheckResourceAttr("fortios_firewall_iptranslation.trname", "type", "SCTP"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallIpTranslationExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallIpTranslation: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallIpTranslation is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallIpTranslation(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallIpTranslation: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallIpTranslation: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallIpTranslationDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_iptranslation" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallIpTranslation(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallIpTranslation %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallIpTranslationConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_iptranslation" "trname" {
  endip       = "2.2.2.2"
  map_startip = "0.0.0.0"
  startip     = "1.1.1.1"
  transid     = 1
  type        = "SCTP"
}
`)
}
