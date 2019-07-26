package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccFortiOSFirewallObjectService_Fqdn(t *testing.T) {
	rname := "fosfq_" + acctest.RandString(12)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFirewallObjectServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallObjectServiceConfigFqdn(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallObjectServiceExists("fortios_firewall_object_service.test1"),
					resource.TestCheckResourceAttr("fortios_firewall_object_service.test1", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewall_object_service.test1", "category", "General"),
					resource.TestCheckResourceAttr("fortios_firewall_object_service.test1", "protocol", "ICMP"),
					resource.TestCheckResourceAttr("fortios_firewall_object_service.test1", "icmptype", "2"),
					resource.TestCheckResourceAttr("fortios_firewall_object_service.test1", "icmpcode", "3"),
					resource.TestCheckResourceAttr("fortios_firewall_object_service.test1", "protocol_number", "1"),
					resource.TestCheckResourceAttr("fortios_firewall_object_service.test1", "comment", "Terraform Test"),
				),
			},
		},
	})
}

func TestAccFortiOSFirewallObjectService_IPRange(t *testing.T) {
	rname := "fosipr_" + acctest.RandString(12)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFirewallObjectServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallObjectServiceConfigIPRange(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallObjectServiceExists("fortios_firewall_object_service.test2"),
					resource.TestCheckResourceAttr("fortios_firewall_object_service.test2", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewall_object_service.test2", "category", "General"),
					resource.TestCheckResourceAttr("fortios_firewall_object_service.test2", "protocol", "TCP/UDP/SCTP"),
					resource.TestCheckResourceAttr("fortios_firewall_object_service.test2", "iprange", "1.1.1.1-2.2.2.2"),
					resource.TestCheckResourceAttr("fortios_firewall_object_service.test2", "tcp_portrange", "22-33"),
					resource.TestCheckResourceAttr("fortios_firewall_object_service.test2", "udp_portrange", "44-55"),
					resource.TestCheckResourceAttr("fortios_firewall_object_service.test2", "sctp_portrange", "66-88"),
					resource.TestCheckResourceAttr("fortios_firewall_object_service.test2", "session_ttl", "1000"),
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
		o, err := c.ReadFirewallObjectService(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error Firewall Object Service %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallObjectServiceConfigFqdn(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_object_service" "test1" {
	name = "%s"
	category = "General"
	protocol = "ICMP"
	icmptype = "2"
	icmpcode = "3"
	protocol_number = "1"
	comment = "Terraform Test"
}
`, name)
}

func testAccFortiOSFirewallObjectServiceConfigIPRange(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_object_service" "test2" {
	name = "%s"
	category = "General"
	protocol = "TCP/UDP/SCTP"
	iprange = "1.1.1.1-2.2.2.2"
	tcp_portrange = "22-33"
	udp_portrange = "44-55"
	sctp_portrange = "66-88"
	session_ttl = "1000"
	comment = "Terraform Test"
}
`, name)
}
