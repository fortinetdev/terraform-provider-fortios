package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccFortiOSFirewallObjectAddress_withIPRange(t *testing.T) {
	rname := "foair_" + acctest.RandString(12)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFirewallObjectAddressDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallObjectAddressConfigIPRange(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallObjectAddressExists("fortios_firewall_object_address.test1"),
					resource.TestCheckResourceAttr("fortios_firewall_object_address.test1", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewall_object_address.test1", "type", "iprange"),
					resource.TestCheckResourceAttr("fortios_firewall_object_address.test1", "start_ip", "1.0.0.0"),
					resource.TestCheckResourceAttr("fortios_firewall_object_address.test1", "end_ip", "2.0.0.0"),
					resource.TestCheckResourceAttr("fortios_firewall_object_address.test1", "subnet", "1.0.0.0 2.0.0.0"),
					resource.TestCheckResourceAttr("fortios_firewall_object_address.test1", "comment", "Terraform Test"),
				),
			},
		},
	})
}

func TestAccFortiOSFirewallObjectAddress_withGeography(t *testing.T) {
	rname := "foagg_" + acctest.RandString(12)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFirewallObjectAddressDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallObjectAddressConfigGeography(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallObjectAddressExists("fortios_firewall_object_address.test2"),
					resource.TestCheckResourceAttr("fortios_firewall_object_address.test2", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewall_object_address.test2", "type", "geography"),
					resource.TestCheckResourceAttr("fortios_firewall_object_address.test2", "country", "AO"),
					resource.TestCheckResourceAttr("fortios_firewall_object_address.test2", "comment", "Terraform Test"),
				),
			},
		},
	})
}

func TestAccFortiOSFirewallObjectAddress_withFqdn(t *testing.T) {
	rname := "foafq_" + acctest.RandString(12)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFirewallObjectAddressDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallObjectAddressConfigFqdn(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallObjectAddressExists("fortios_firewall_object_address.test3"),
					resource.TestCheckResourceAttr("fortios_firewall_object_address.test3", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewall_object_address.test3", "type", "fqdn"),
					resource.TestCheckResourceAttr("fortios_firewall_object_address.test3", "fqdn", "baid.com"),
					resource.TestCheckResourceAttr("fortios_firewall_object_address.test3", "comment", "Terraform Test"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallObjectAddressExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found Firewall Object Address: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Firewall Object Address is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

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

func testAccCheckFirewallObjectAddressDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_object_address" {
			continue
		}

		i := rs.Primary.ID
		_, err := c.ReadFirewallObjectAddress(i)

		if err == nil {
			return fmt.Errorf("Error Firewall Object Address %s still exists", rs.Primary.ID)
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallObjectAddressConfigIPRange(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_object_address" "test1" {
	name = "%s"
	type = "iprange"
	start_ip = "1.0.0.0"
	end_ip = "2.0.0.0"
	subnet = "1.0.0.0 2.0.0.0"
	comment = "Terraform Test"
}
`, name)
}

func testAccFortiOSFirewallObjectAddressConfigGeography(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_object_address" "test2" {
	name = "%s"
	type = "geography"
	country = "AO"
	comment = "Terraform Test"
}
`, name)
}

func testAccFortiOSFirewallObjectAddressConfigFqdn(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_object_address" "test3" {
	name = "%s"
	type = "fqdn"
	fqdn = "baid.com"
	comment = "Terraform Test"
}
`, name)
}
