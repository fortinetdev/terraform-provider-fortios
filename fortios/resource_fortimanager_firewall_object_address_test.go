package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccFortiManagerFirewallObjectAddress(t *testing.T) {
	name := "fmg-firewall-object-address" + acctest.RandString(12)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckFortiManager(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFMGFirewallObjectAddressDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiManagerFirewallObjectAddressConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiManagerFirewallObjectAddressExists("fortios_fortimanager_firewall_object_address.test1"),
					resource.TestCheckResourceAttr("fortios_fortimanager_firewall_object_address.test1", "name", name),
					resource.TestCheckResourceAttr("fortios_fortimanager_firewall_object_address.test1", "type", "fqdn"),
					resource.TestCheckResourceAttr("fortios_fortimanager_firewall_object_address.test1", "comment", "test obj address"),
					resource.TestCheckResourceAttr("fortios_fortimanager_firewall_object_address.test1", "fqdn", "fqdn.google.com"),
					resource.TestCheckResourceAttr("fortios_fortimanager_firewall_object_address.test1", "associated_intf", "any"),
					resource.TestCheckResourceAttr("fortios_fortimanager_firewall_object_address.test1", "allow_routing", "disable"),
				),
			},
		},
	})
}

func testAccCheckFortiManagerFirewallObjectAddressExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found Firewall Object Address: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Firewall Object Address is set")
		}

		c := testAccProvider.Meta().(*FortiClient).ClientFortimanager

		i := rs.Primary.ID
		o, err := c.ReadFirewallObjectAddress(i)

		if err != nil {
			return fmt.Errorf("Error reading Firewall Object Address: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating Firewall Object Address: %s", n)
		}

		return nil
	}
}

func testAccCheckFMGFirewallObjectAddressDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).ClientFortimanager

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_fortimanager_firewall_object_address" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallObjectAddress(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error Firewall Object Address %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiManagerFirewallObjectAddressConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_fortimanager_firewall_object_address" "test1" {
    name = "%s"
    type = "fqdn"
    comment = "test obj address"
    fqdn = "fqdn.google.com"
    associated_intf = "any"
	allow_routing = "disable"
}
`, name)
}
