package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccFortiOSFirewallObjectIPPool_Overload(t *testing.T) {
	rname := "foipol_" + acctest.RandString(12)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFirewallObjectIPPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallObjectIPPoolConfigOverload(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallObjectIPPoolExists("fortios_firewall_object_ippool.test1"),
					resource.TestCheckResourceAttr("fortios_firewall_object_ippool.test1", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewall_object_ippool.test1", "type", "overload"),
					resource.TestCheckResourceAttr("fortios_firewall_object_ippool.test1", "startip", "1.2.5.0"),
					resource.TestCheckResourceAttr("fortios_firewall_object_ippool.test1", "endip", "1.2.6.0"),
					resource.TestCheckResourceAttr("fortios_firewall_object_ippool.test1", "arp_reply", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_object_ippool.test1", "comments", "Terraform Test"),
				),
			},
		},
	})
}

func TestAccFortiOSFirewallObjectIPPool_OneToOne(t *testing.T) {
	rname := "foipoto_" + acctest.RandString(12)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFirewallObjectIPPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallObjectIPPoolConfigOneToOne(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallObjectIPPoolExists("fortios_firewall_object_ippool.test2"),
					resource.TestCheckResourceAttr("fortios_firewall_object_ippool.test2", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewall_object_ippool.test2", "type", "one-to-one"),
					resource.TestCheckResourceAttr("fortios_firewall_object_ippool.test2", "startip", "3.0.0.0"),
					resource.TestCheckResourceAttr("fortios_firewall_object_ippool.test2", "endip", "4.0.0.0"),
					resource.TestCheckResourceAttr("fortios_firewall_object_ippool.test2", "arp_reply", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_object_ippool.test2", "comments", "Terraform Test"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallObjectIPPoolExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found Firewall Object IPPool: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Firewall Object IPPool is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallObjectIPPool(i)

		if err != nil {
			return fmt.Errorf("Error reading Firewall Object IPPool: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating Firewall Object IPPool: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallObjectIPPoolDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_object_ippool" {
			continue
		}

		i := rs.Primary.ID
		_, err := c.ReadFirewallObjectIPPool(i)

		if err == nil {
			return fmt.Errorf("Error Firewall Object IPPool %s still exists", rs.Primary.ID)
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallObjectIPPoolConfigOverload(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_object_ippool" "test1" {
	name = "%s"
	type = "overload"
	startip = "1.2.5.0"
	endip = "1.2.6.0"
	arp_reply = "enable"
	comments = "Terraform Test"
}
`, name)
}

func testAccFortiOSFirewallObjectIPPoolConfigOneToOne(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_object_ippool" "test2" {
	name = "%s"
	type = "one-to-one"
	startip = "3.0.0.0"
	endip = "4.0.0.0"
	arp_reply = "enable"
	comments = "Terraform Test"
}
`, name)
}
