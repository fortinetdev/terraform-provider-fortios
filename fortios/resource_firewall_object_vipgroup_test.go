package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccFortiOSFirewallObjectVipGroup_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFirewallObjectVipGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallObjectVipGroupConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallObjectVipGroupExists("fortios_firewall_object_vipgroup.test1"),
					resource.TestCheckResourceAttr("fortios_firewall_object_vipgroup.test1", "name", "s1"),
					resource.TestCheckResourceAttr("fortios_firewall_object_vipgroup.test1", "comments", "Terraform Test"),
					resource.TestCheckResourceAttr("fortios_firewall_object_vipgroup.test1", "interface", "port3"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallObjectVipGroupExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found Firewall Object VipGroup: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Firewall Object VipGroup is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallObjectVipGroup(i)

		if err != nil {
			return fmt.Errorf("Error reading Firewall Object VipGroup: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating Firewall Object VipGroup: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallObjectVipGroupDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_object_vipgroup" {
			continue
		}

		i := rs.Primary.ID
		_, err := c.ReadFirewallObjectVipGroup(i)

		if err == nil {
			return fmt.Errorf("Error Firewall Object VipGroup %s still exists", rs.Primary.ID)
		}

		return nil
	}

	return nil
}

const testAccFortiOSFirewallObjectVipGroupConfig = `
resource "fortios_firewall_object_vip" "viptest1" { 
	name = "vip1fortest"
	comment = "Terraform Test"
	extip = "1.1.5.0-1.1.6.0"
	mappedip = ["1.1.8.0-1.1.9.0"]
	extintf = "port3"
	portforward = "enable"
	protocol = "tcp"
	extport = "2-3"
	mappedport = "4-5"
}

resource "fortios_firewall_object_vip" "viptest2" { 
	name = "vip2fortest"
	comment = "Terraform Test"
	extip = "1.1.7.0-1.1.8.0"
	mappedip = ["1.1.10.0-1.1.11.0"]
	extintf = "port3"
	portforward = "enable"
	protocol = "tcp"
	extport = "2-3"
	mappedport = "4-5"
}


resource "fortios_firewall_object_vipgroup" "test1" {
	name = "s1"
	interface = "port3"
	comments = "Terraform Test"
	member = ["${fortios_firewall_object_vip.viptest1.name}", "${fortios_firewall_object_vip.viptest2.name}"]
}
`
