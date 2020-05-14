package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccFortiManagerFirewallObjectIppool(t *testing.T) {
	name := "fmg-firewall-object-ippool" + acctest.RandString(2)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckFortiManager(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFMGFirewallObjectIppoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiManagerFirewallObjectIppoolConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiManagerFirewallObjectIppoolExists("fortios_fmg_firewall_object_ippool.test1"),
					resource.TestCheckResourceAttr("fortios_fmg_firewall_object_ippool.test1", "name", name),
					resource.TestCheckResourceAttr("fortios_fmg_firewall_object_ippool.test1", "type", "one-to-one"),
					resource.TestCheckResourceAttr("fortios_fmg_firewall_object_ippool.test1", "comment", "test obj ippool"),
					resource.TestCheckResourceAttr("fortios_fmg_firewall_object_ippool.test1", "startip", "1.1.10.1"),
					resource.TestCheckResourceAttr("fortios_fmg_firewall_object_ippool.test1", "endip", "1.1.10.100"),
					resource.TestCheckResourceAttr("fortios_fmg_firewall_object_ippool.test1", "arp_intf", "any"),
					resource.TestCheckResourceAttr("fortios_fmg_firewall_object_ippool.test1", "arp_reply", "enable"),
					resource.TestCheckResourceAttr("fortios_fmg_firewall_object_ippool.test1", "associated_intf", "any"),
				),
			},
		},
	})
}

func testAccCheckFortiManagerFirewallObjectIppoolExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found Firewall Object Ippool: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Firewall Object Ippool is set")
		}

		c := testAccProvider.Meta().(*FortiClient).ClientFortimanager

		i := rs.Primary.ID
		o, err := c.ReadFirewallObjectIppool("root", i)

		if err != nil {
			return fmt.Errorf("Error reading Firewall Object Ippool: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating Firewall Object Ippool: %s", n)
		}

		return nil
	}
}

func testAccCheckFMGFirewallObjectIppoolDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).ClientFortimanager

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_fmg_firewall_object_ippool" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallObjectIppool("root", i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error Firewall Object Ippool %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiManagerFirewallObjectIppoolConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_fmg_firewall_object_ippool" "test1" {
    name = "%s"
    comment = "test obj ippool"
    type = "one-to-one"
    startip = "1.1.10.1"
    endip = "1.1.10.100"
    arp_intf = "any"
    arp_reply = "enable"
    associated_intf = "any"
}
`, name)
}
