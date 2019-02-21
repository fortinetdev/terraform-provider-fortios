package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccFortiOSFirewallObjectVip_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFirewallObjectVipDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallObjectVipConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallObjectVipExists("fortios_firewall_object_vip.test1"),
					resource.TestCheckResourceAttr("fortios_firewall_object_vip.test1", "name", "s1"),
					resource.TestCheckResourceAttr("fortios_firewall_object_vip.test1", "comment", "Terraform Test"),
					resource.TestCheckResourceAttr("fortios_firewall_object_vip.test1", "extip", "1.1.5.0-1.1.6.0"),
					resource.TestCheckResourceAttr("fortios_firewall_object_vip.test1", "extintf", "port3"),
					resource.TestCheckResourceAttr("fortios_firewall_object_vip.test1", "portforward", "disable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallObjectVipExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found Firewall Object Vip: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Firewall Object Vip is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallObjectVip(i)

		if err != nil {
			return fmt.Errorf("Error reading Firewall Object Vip: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating Firewall Object Vip: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallObjectVipDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_object_vip" {
			continue
		}

		i := rs.Primary.ID
		_, err := c.ReadFirewallObjectVip(i)

		if err == nil {
			return fmt.Errorf("Error Firewall Object Vip %s still exists", rs.Primary.ID)
		}

		return nil
	}

	return nil
}

const testAccFortiOSFirewallObjectVipConfig = `
resource "fortios_firewall_object_vip" "test1" { 
	name = "s1"
	comment = "Terraform Test"
	extip = "1.1.5.0-1.1.6.0"
	mappedip = ["1.1.8.0-1.1.9.0"]
	extintf = "port3"
	portforward = "disable"
}
`
