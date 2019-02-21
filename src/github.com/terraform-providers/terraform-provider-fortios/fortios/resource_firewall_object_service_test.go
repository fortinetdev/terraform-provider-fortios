package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccFortiOSFirewallObjectService_Fqdn(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFirewallObjectServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallObjectServiceConfigFqdn,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallObjectServiceExists("fortios_firewall_object_service.test1"),
					resource.TestCheckResourceAttr("fortios_firewall_object_service.test1", "name", "s1"),
					resource.TestCheckResourceAttr("fortios_firewall_object_service.test1", "category", "General"),
					resource.TestCheckResourceAttr("fortios_firewall_object_service.test1", "protocol", "TCP/UDP/SCTP"),
					resource.TestCheckResourceAttr("fortios_firewall_object_service.test1", "fqdn", "asu"),
					resource.TestCheckResourceAttr("fortios_firewall_object_service.test1", "comment", "Terraform Test"),
				),
			},
		},
	})
}

func TestAccFortiOSFirewallObjectService_IPRange(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFirewallObjectServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallObjectServiceConfigIPRange,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallObjectServiceExists("fortios_firewall_object_service.test2"),
					resource.TestCheckResourceAttr("fortios_firewall_object_service.test2", "name", "s2"),
					resource.TestCheckResourceAttr("fortios_firewall_object_service.test2", "category", "General"),
					resource.TestCheckResourceAttr("fortios_firewall_object_service.test2", "protocol", "TCP/UDP/SCTP"),
					resource.TestCheckResourceAttr("fortios_firewall_object_service.test2", "iprange", "1.1.1.1-2.2.2.2"),
					resource.TestCheckResourceAttr("fortios_firewall_object_service.test2", "comment", "Terraform Test"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallObjectServiceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found Firewall Object Service: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Firewall Object Service is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallObjectService(i)

		if err != nil {
			return fmt.Errorf("Error reading Firewall Object Service: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating Firewall Object Service: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallObjectServiceDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_object_service" {
			continue
		}

		i := rs.Primary.ID
		_, err := c.ReadFirewallObjectService(i)

		if err == nil {
			return fmt.Errorf("Error Firewall Object Service %s still exists", rs.Primary.ID)
		}

		return nil
	}

	return nil
}

const testAccFortiOSFirewallObjectServiceConfigFqdn = `
resource "fortios_firewall_object_service" "test1" {
	name = "s1"
	category = "General"
	protocol = "TCP/UDP/SCTP"
	fqdn = "asu"
	comment = "Terraform Test"
}
`

const testAccFortiOSFirewallObjectServiceConfigIPRange = `
resource "fortios_firewall_object_service" "test2" {
	name = "s2"
	category = "General"
	protocol = "TCP/UDP/SCTP"
	iprange = "1.1.1.1-2.2.2.2"
	comment = "Terraform Test"
}
`
